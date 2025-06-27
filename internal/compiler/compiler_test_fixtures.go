package compiler

import (
	"embed"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/bmatcuk/doublestar/v4"
)

//go:embed testdata
var testdataFS embed.FS

// TestGlobPatternMatching tests glob pattern logic without filesystem operations
func TestGlobPatternMatching(t *testing.T) {
	tests := []struct {
		name        string
		pattern     string
		testPath    string
		shouldMatch bool
	}{
		{
			name:        "simple recursive pattern",
			pattern:     "**/level1/*.md",
			testPath:    "some/path/level1/file.md",
			shouldMatch: true,
		},
		{
			name:        "deep recursive pattern",
			pattern:     "**/*.md",
			testPath:    "level1/level2/level3/file.md",
			shouldMatch: true,
		},
		{
			name:        "exclusion pattern",
			pattern:     "**/_resources/**",
			testPath:    "vault/_resources/template.md",
			shouldMatch: true,
		},
		{
			name:        "exclusion pattern deep",
			pattern:     "**/_resources/**",
			testPath:    "vault/_resources/nested/template.md",
			shouldMatch: true,
		},
		{
			name:        "non-matching extension",
			pattern:     "**/*.md",
			testPath:    "level1/file.txt",
			shouldMatch: false,
		},
		{
			name:        "single vs double asterisk",
			pattern:     "*/level2/*.md",
			testPath:    "level1/level2/level3/deep.md",
			shouldMatch: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matched, err := doublestar.Match(tt.pattern, tt.testPath)
			if err != nil {
				t.Fatalf("Error matching pattern: %v", err)
			}

			if matched != tt.shouldMatch {
				t.Errorf("Pattern %q against %q: expected %t, got %t",
					tt.pattern, tt.testPath, tt.shouldMatch, matched)
			}
		})
	}
}

// TestRecursiveGlobbingWithFixtures uses embedded testdata for predictable testing
func TestRecursiveGlobbingWithFixtures(t *testing.T) {
	// Copy embedded files to temp directory for testing
	tempDir := t.TempDir()

	err := copyEmbeddedFS(testdataFS, "testdata", tempDir)
	if err != nil {
		t.Fatalf("Failed to copy test fixtures: %v", err)
	}

	tests := []struct {
		name              string
		patterns          []string
		expectedFiles     []string
		expectedProcessed int
		expectedExcluded  int
	}{
		{
			name:     "recursive glob finds all markdown files",
			patterns: []string{filepath.Join(tempDir, "**/*.md")},
			expectedFiles: []string{
				"root.md", "file1.md", "file2.md", "deep1.md",
				"deepest.md", "excluded.md", "template.md",
			},
			expectedProcessed: 7, // All .md files
			expectedExcluded:  0,
		},
		{
			name: "recursive glob with exclusions",
			patterns: []string{
				filepath.Join(tempDir, "**/*.md"),
				"!" + filepath.Join(tempDir, "exclude_dir/**"),
				"!" + filepath.Join(tempDir, "**/_resources/**"),
			},
			expectedFiles: []string{
				"root.md", "file1.md", "file2.md", "deep1.md", "deepest.md",
			},
			expectedProcessed: 5,
			expectedExcluded:  2, // excluded.md and template.md
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			c := &Compiler{
				config:     &Config{Verbose: false},
				fileConfig: &FileConfig{},
			}

			processedCount, excludedCount, err := c.processFiles(&buf, tt.patterns)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			if processedCount != tt.expectedProcessed {
				t.Errorf("Expected %d processed files, got %d", tt.expectedProcessed, processedCount)
			}

			if excludedCount != tt.expectedExcluded {
				t.Errorf("Expected %d excluded files, got %d", tt.expectedExcluded, excludedCount)
			}

			output := buf.String()
			for _, expectedFile := range tt.expectedFiles {
				if !strings.Contains(output, expectedFile) {
					t.Errorf("Expected %s in output", expectedFile)
				}
			}
		})
	}
}

// copyEmbeddedFS copies an embedded filesystem to a real directory
func copyEmbeddedFS(fsys embed.FS, srcDir, destDir string) error {
	return fs.WalkDir(fsys, srcDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Skip the root directory
		if path == srcDir {
			return nil
		}

		// Calculate relative path
		relPath, err := filepath.Rel(srcDir, path)
		if err != nil {
			return err
		}

		destPath := filepath.Join(destDir, relPath)

		if d.IsDir() {
			return os.MkdirAll(destPath, 0755)
		}

		// Copy file
		data, err := fs.ReadFile(fsys, path)
		if err != nil {
			return err
		}

		return os.WriteFile(destPath, data, 0644)
	})
}
