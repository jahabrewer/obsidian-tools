package compiler

import (
	"testing"

	"github.com/bmatcuk/doublestar/v4"
)

// matchesAnyPattern checks if a file path matches any of the given patterns
func matchesAnyPattern(filePath string, patterns []string) bool {
	for _, pattern := range patterns {
		if matched, _ := doublestar.Match(pattern, filePath); matched {
			return true
		}
	}
	return false
}

// categorizeFiles separates files into included and excluded based on patterns
func categorizeFiles(filePaths, includePatterns, excludePatterns []string) (included, excluded []string) {
	for _, filePath := range filePaths {
		if !matchesAnyPattern(filePath, includePatterns) {
			continue // File doesn't match include patterns
		}

		if matchesAnyPattern(filePath, excludePatterns) {
			excluded = append(excluded, filePath)
		} else {
			included = append(included, filePath)
		}
	}
	return included, excluded
}

// verifyIncludedFiles checks that all expected files are in the included list
func verifyIncludedFiles(t *testing.T, included, expectedIncluded []string) {
	if len(included) != len(expectedIncluded) {
		t.Errorf("Expected %d included files, got %d. Expected: %v, Got: %v",
			len(expectedIncluded), len(included), expectedIncluded, included)
	}

	for _, expectedFile := range expectedIncluded {
		found := false
		for _, actualFile := range included {
			if actualFile == expectedFile {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected file %s to be included, but it wasn't", expectedFile)
		}
	}
}

// verifyExcludedFiles checks that all expected files are in the excluded list
func verifyExcludedFiles(t *testing.T, excluded, expectedExcluded []string) {
	for _, expectedFile := range expectedExcluded {
		found := false
		for _, actualFile := range excluded {
			if actualFile == expectedFile {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected file %s to be excluded, but it wasn't", expectedFile)
		}
	}
}

// TestGlobLogicOnly tests the glob matching logic without any filesystem operations
func TestGlobLogicOnly(t *testing.T) {
	tests := []struct {
		name             string
		includePatterns  []string
		excludePatterns  []string
		filePaths        []string
		expectedIncluded []string
		expectedExcluded []string
	}{
		{
			name:            "recursive pattern includes all levels",
			includePatterns: []string{"**/vault/**/*.md"},
			excludePatterns: []string{},
			filePaths: []string{
				"vault/root.md",
				"vault/level1/file.md",
				"vault/level1/level2/deep.md",
				"vault/level1/level2/level3/deepest.md",
				"vault/other.txt",
			},
			expectedIncluded: []string{
				"vault/root.md",
				"vault/level1/file.md",
				"vault/level1/level2/deep.md",
				"vault/level1/level2/level3/deepest.md",
			},
			expectedExcluded: []string{},
		},
		{
			name:            "exclusion patterns work recursively",
			includePatterns: []string{"**/vault/**/*.md"},
			excludePatterns: []string{"**/vault/_resources/**", "**/vault/archive/**"},
			filePaths: []string{
				"vault/note.md",
				"vault/_resources/template.md",
				"vault/_resources/nested/template2.md",
				"vault/archive/old.md",
				"vault/kb/knowledge.md",
			},
			expectedIncluded: []string{
				"vault/note.md",
				"vault/kb/knowledge.md",
			},
			expectedExcluded: []string{
				"vault/_resources/template.md",
				"vault/_resources/nested/template2.md",
				"vault/archive/old.md",
			},
		},
		{
			name:            "single asterisk vs double asterisk",
			includePatterns: []string{"vault/*/file.md"}, // Single asterisk - only one level
			excludePatterns: []string{},
			filePaths: []string{
				"vault/level1/file.md",        // Should match (1 level)
				"vault/level1/level2/file.md", // Should NOT match (2 levels)
			},
			expectedIncluded: []string{
				"vault/level1/file.md",
			},
			expectedExcluded: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			included, excluded := categorizeFiles(tt.filePaths, tt.includePatterns, tt.excludePatterns)
			verifyIncludedFiles(t, included, tt.expectedIncluded)
			verifyExcludedFiles(t, excluded, tt.expectedExcluded)
		})
	}
}

// BenchmarkGlobMatching benchmarks different glob patterns
func BenchmarkGlobMatching(b *testing.B) {
	testPath := "vault/level1/level2/level3/deep.md"

	benchmarks := []struct {
		name    string
		pattern string
	}{
		{"SingleAsterisk", "vault/*/deep.md"},
		{"DoubleAsterisk", "vault/**/*.md"},
		{"ComplexPattern", "**/vault/**/level*/*.md"},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = doublestar.Match(bm.pattern, testPath)
			}
		})
	}
}
