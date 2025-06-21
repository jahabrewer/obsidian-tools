package compiler

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	config := &Config{
		Verbose:   true,
		Clipboard: false,
	}

	compiler := New(config)

	if compiler == nil {
		t.Fatal("New() returned nil")
	}

	if compiler.config.Verbose != true {
		t.Errorf("Expected verbose=true, got %v", compiler.config.Verbose)
	}
}

func TestDetermineOutputFile(t *testing.T) {
	tests := []struct {
		name           string
		configFile     string
		fileConfigPath string
		expected       string
		expectError    bool
	}{
		{
			name:        "config file path provided",
			configFile:  "output.md",
			expected:    "output.md",
			expectError: false,
		},
		{
			name:           "file config path provided",
			fileConfigPath: "config-output.md",
			expected:       "config-output.md",
			expectError:    false,
		},
		{
			name:        "no output file specified",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Compiler{
				config: &Config{
					OutputFile: tt.configFile,
				},
				fileConfig: &FileConfig{
					OutputFilePath: tt.fileConfigPath,
				},
			}

			result, err := c.determineOutputFile()

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error, but got none")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if result != tt.expected {
					t.Errorf("Expected %s, got %s", tt.expected, result)
				}
			}
		})
	}
}

func TestDetermineGlobPatterns(t *testing.T) {
	tests := []struct {
		name               string
		configPatterns     []string
		fileConfigPatterns []string
		expected           []string
		expectError        bool
	}{
		{
			name:           "config patterns provided",
			configPatterns: []string{"*.md", "!test.md"},
			expected:       []string{"*.md", "!test.md"},
			expectError:    false,
		},
		{
			name:               "file config patterns provided",
			fileConfigPatterns: []string{"**/*.md"},
			expected:           []string{"**/*.md"},
			expectError:        false,
		},
		{
			name:        "no patterns specified",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Compiler{
				config: &Config{
					GlobPatterns: tt.configPatterns,
				},
				fileConfig: &FileConfig{
					GlobPatterns: tt.fileConfigPatterns,
				},
			}

			result, err := c.determineGlobPatterns()

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error, but got none")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if len(result) != len(tt.expected) {
					t.Errorf("Expected %d patterns, got %d", len(tt.expected), len(result))
				}
			}
		})
	}
}

func TestLoadConfig(t *testing.T) {
	// Create a temporary directory for test config files
	tempDir, err := os.MkdirTemp("", "compiler_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	t.Run("config file doesn't exist", func(t *testing.T) {
		c := &Compiler{
			config: &Config{
				ConfigFile: filepath.Join(tempDir, "nonexistent.yaml"),
			},
			fileConfig: &FileConfig{},
		}

		err := c.loadConfig()
		if err == nil {
			t.Error("Expected error for nonexistent specified config file")
		}
	})

	t.Run("default config file doesn't exist", func(t *testing.T) {
		c := &Compiler{
			config: &Config{
				ConfigFile: "", // Will use default ~/.note-compiler.yaml
			},
			fileConfig: &FileConfig{},
		}

		// Temporarily set HOME to our temp dir
		originalHome := os.Getenv("HOME")
		os.Setenv("HOME", tempDir)
		defer os.Setenv("HOME", originalHome)

		err := c.loadConfig()
		if err != nil {
			t.Errorf("Unexpected error when default config doesn't exist: %v", err)
		}
	})

	t.Run("valid config file", func(t *testing.T) {
		configFile := filepath.Join(tempDir, "test-config.yaml")
		configContent := `output_file_path: "test-output.md"
glob_patterns:
  - "*.md"
  - "!excluded.md"`

		err := os.WriteFile(configFile, []byte(configContent), 0644)
		if err != nil {
			t.Fatal(err)
		}

		c := &Compiler{
			config: &Config{
				ConfigFile: configFile,
				Verbose:    false,
			},
			fileConfig: &FileConfig{},
		}

		err = c.loadConfig()
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		if c.fileConfig.OutputFilePath != "test-output.md" {
			t.Errorf("Expected output_file_path 'test-output.md', got '%s'", c.fileConfig.OutputFilePath)
		}

		if len(c.fileConfig.GlobPatterns) != 2 {
			t.Errorf("Expected 2 glob patterns, got %d", len(c.fileConfig.GlobPatterns))
		}
	})

	t.Run("invalid yaml config", func(t *testing.T) {
		configFile := filepath.Join(tempDir, "invalid-config.yaml")
		invalidContent := `output_file_path: "test
invalid yaml content`

		err := os.WriteFile(configFile, []byte(invalidContent), 0644)
		if err != nil {
			t.Fatal(err)
		}

		c := &Compiler{
			config: &Config{
				ConfigFile: configFile,
			},
			fileConfig: &FileConfig{},
		}

		err = c.loadConfig()
		if err == nil {
			t.Error("Expected error for invalid YAML")
		}
		if !strings.Contains(err.Error(), "failed to parse config file") {
			t.Errorf("Expected parse error, got: %v", err)
		}
	})
}

func TestProcessFile(t *testing.T) {
	// Create a temporary directory for test files
	tempDir, err := os.MkdirTemp("", "compiler_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	// Create a test markdown file
	testFile := filepath.Join(tempDir, "test.md")
	testContent := "# Test Header\n\nThis is test content."
	err = os.WriteFile(testFile, []byte(testContent), 0644)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("process file successfully", func(t *testing.T) {
		var buf bytes.Buffer
		c := &Compiler{
			config: &Config{
				Verbose: false,
			},
		}

		err := c.processFile(&buf, testFile)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		output := buf.String()
		if !strings.Contains(output, "source: "+testFile) {
			t.Error("Expected source annotation in output")
		}
		if !strings.Contains(output, testContent) {
			t.Error("Expected file content in output")
		}
	})

	t.Run("process file with verbose mode", func(t *testing.T) {
		var buf bytes.Buffer
		c := &Compiler{
			config: &Config{
				Verbose: true,
			},
		}

		// Capture stdout to test verbose output
		oldStdout := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		err := c.processFile(&buf, testFile)

		w.Close()
		os.Stdout = oldStdout

		verboseOutput, _ := io.ReadAll(r)

		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		if !strings.Contains(string(verboseOutput), "Including file: "+testFile) {
			t.Error("Expected verbose output about including file")
		}
	})

	t.Run("process nonexistent file", func(t *testing.T) {
		var buf bytes.Buffer
		c := &Compiler{
			config: &Config{
				Verbose: false,
			},
		}

		err := c.processFile(&buf, filepath.Join(tempDir, "nonexistent.md"))
		if err == nil {
			t.Error("Expected error for nonexistent file")
		}
	})
}

func TestProcessFiles(t *testing.T) {
	// Create a temporary directory for test files
	tempDir, err := os.MkdirTemp("", "compiler_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	// Create test files
	testFiles := []struct {
		name    string
		content string
	}{
		{"test1.md", "# Test 1\nContent 1"},
		{"test2.md", "# Test 2\nContent 2"},
		{"exclude.md", "# Exclude\nThis should be excluded"},
		{"test.txt", "Not a markdown file"},
	}

	for _, tf := range testFiles {
		err := os.WriteFile(filepath.Join(tempDir, tf.name), []byte(tf.content), 0644)
		if err != nil {
			t.Fatal(err)
		}
	}

	t.Run("process files with inclusion and exclusion", func(t *testing.T) {
		var buf bytes.Buffer
		c := &Compiler{
			config: &Config{
				Verbose: false,
			},
		}

		patterns := []string{
			filepath.Join(tempDir, "*.md"),
			"!" + filepath.Join(tempDir, "exclude.md"),
		}

		processedCount, excludedCount, err := c.processFiles(&buf, patterns)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		if processedCount != 2 {
			t.Errorf("Expected 2 processed files, got %d", processedCount)
		}

		if excludedCount != 1 {
			t.Errorf("Expected 1 excluded file, got %d", excludedCount)
		}

		output := buf.String()
		if !strings.Contains(output, "test1.md") {
			t.Error("Expected test1.md in output")
		}
		if !strings.Contains(output, "test2.md") {
			t.Error("Expected test2.md in output")
		}
		if strings.Contains(output, "exclude.md") {
			t.Error("exclude.md should not be in output")
		}
	})

	t.Run("process files with invalid glob pattern", func(t *testing.T) {
		var buf bytes.Buffer
		c := &Compiler{
			config: &Config{
				Verbose: false,
			},
		}

		patterns := []string{"[invalid glob pattern"}

		_, _, err := c.processFiles(&buf, patterns)
		if err == nil {
			t.Error("Expected error for invalid glob pattern")
		}
	})

	t.Run("process files with no matches", func(t *testing.T) {
		var buf bytes.Buffer
		c := &Compiler{
			config: &Config{
				Verbose: false,
			},
		}

		patterns := []string{filepath.Join(tempDir, "*.nonexistent")}

		processedCount, excludedCount, err := c.processFiles(&buf, patterns)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		if processedCount != 0 {
			t.Errorf("Expected 0 processed files, got %d", processedCount)
		}

		if excludedCount != 0 {
			t.Errorf("Expected 0 excluded files, got %d", excludedCount)
		}
	})
}

func TestCopyToClipboard(t *testing.T) {
	// Create a temporary file with test content
	tempDir, err := os.MkdirTemp("", "compiler_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	testFile := filepath.Join(tempDir, "test.txt")
	testContent := "Test clipboard content"
	err = os.WriteFile(testFile, []byte(testContent), 0644)
	if err != nil {
		t.Fatal(err)
	}

	c := &Compiler{
		config: &Config{},
	}

	t.Run("copy existing file to clipboard", func(t *testing.T) {
		// Note: This test may fail on systems without clipboard support
		// but it tests the file reading part at least
		err := c.copyToClipboard(testFile)
		// We don't assert on the error since clipboard might not be available
		// in all test environments, but we test that it doesn't panic
		_ = err
	})

	t.Run("copy nonexistent file to clipboard", func(t *testing.T) {
		err := c.copyToClipboard(filepath.Join(tempDir, "nonexistent.txt"))
		if err == nil {
			t.Error("Expected error for nonexistent file")
		}
	})
}

func TestRunIntegration(t *testing.T) {
	// Create a temporary directory for test files
	tempDir, err := os.MkdirTemp("", "compiler_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	// Create test files
	testFile1 := filepath.Join(tempDir, "test1.md")
	testFile2 := filepath.Join(tempDir, "test2.md")
	outputFile := filepath.Join(tempDir, "output.md")

	err = os.WriteFile(testFile1, []byte("# Test 1\nContent 1"), 0644)
	if err != nil {
		t.Fatal(err)
	}

	err = os.WriteFile(testFile2, []byte("# Test 2\nContent 2"), 0644)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("successful run", func(t *testing.T) {
		// Capture stdout
		oldStdout := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		// Temporarily set HOME to our temp dir to avoid loading real config
		originalHome := os.Getenv("HOME")
		os.Setenv("HOME", tempDir)
		defer os.Setenv("HOME", originalHome)

		c := &Compiler{
			config: &Config{
				ConfigFile:   "", // Use default path but it won't exist in tempDir
				OutputFile:   outputFile,
				GlobPatterns: []string{filepath.Join(tempDir, "*.md")},
				Verbose:      false,
				Clipboard:    false,
			},
			fileConfig: &FileConfig{},
		}

		err := c.Run()

		w.Close()
		os.Stdout = oldStdout
		output, _ := io.ReadAll(r)

		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		// Check that output file was created
		if _, err := os.Stat(outputFile); os.IsNotExist(err) {
			t.Error("Output file was not created")
		}

		// Check stdout output
		outputStr := string(output)
		if !strings.Contains(outputStr, "Successfully processed") {
			t.Errorf("Expected success message in output, got: %s", outputStr)
		}
	})

	t.Run("run with no output file", func(t *testing.T) {
		// Temporarily set HOME to our temp dir to avoid loading real config
		originalHome := os.Getenv("HOME")
		os.Setenv("HOME", tempDir)
		defer os.Setenv("HOME", originalHome)

		c := &Compiler{
			config: &Config{
				ConfigFile:   "", // Use default path but it won't exist in tempDir
				GlobPatterns: []string{filepath.Join(tempDir, "*.md")},
			},
			fileConfig: &FileConfig{},
		}

		err := c.Run()
		if err == nil {
			t.Error("Expected error when no output file specified")
		}
	})

	t.Run("run with no glob patterns", func(t *testing.T) {
		// Temporarily set HOME to our temp dir to avoid loading real config
		originalHome := os.Getenv("HOME")
		os.Setenv("HOME", tempDir)
		defer os.Setenv("HOME", originalHome)

		c := &Compiler{
			config: &Config{
				ConfigFile: "", // Use default path but it won't exist in tempDir
				OutputFile: outputFile,
			},
			fileConfig: &FileConfig{},
		}

		err := c.Run()
		if err == nil {
			t.Error("Expected error when no glob patterns specified")
		}
	})
}

func TestExpandPath(t *testing.T) {
	t.Run("home template expansion", func(t *testing.T) {
		homeDir, _ := os.UserHomeDir()
		result := expandPath("{{.Home}}/test/file.txt")
		expected := filepath.Join(homeDir, "test", "file.txt")
		if result != expected {
			t.Errorf("Expected %s, got %s", expected, result)
		}
	})

	t.Run("no template", func(t *testing.T) {
		result := expandPath("/absolute/path/file.txt")
		expected := "/absolute/path/file.txt"
		if result != expected {
			t.Errorf("Expected %s, got %s", expected, result)
		}
	})

	t.Run("date template expansion", func(t *testing.T) {
		result := expandPath(`output_{{.Date "2006-01-02"}}.txt`)
		// Should contain current date in YYYY-MM-DD format
		if !strings.Contains(result, time.Now().Format("2006-01-02")) {
			t.Errorf("Expected date template to be expanded, got: %s", result)
		}
	})

	t.Run("complex date template", func(t *testing.T) {
		result := expandPath(`notes_{{.Date "2006-01-02_150405"}}.md`)
		// Should not contain the original template
		if strings.Contains(result, "{{.Date") {
			t.Errorf("Date template not expanded: %s", result)
		}
		// Should contain current year
		if !strings.Contains(result, time.Now().Format("2006")) {
			t.Errorf("Expected current year in result: %s", result)
		}
	})

	t.Run("home and date together", func(t *testing.T) {
		homeDir, _ := os.UserHomeDir()
		result := expandPath(`{{.Home}}/notes_{{.Date "2006-01-02"}}.txt`)

		// Should expand home
		if !strings.HasPrefix(result, homeDir) {
			t.Errorf("Expected result to start with home dir, got: %s", result)
		}

		// Should expand date
		if strings.Contains(result, "{{.Date") {
			t.Errorf("Date template not expanded: %s", result)
		}
	})

	t.Run("environment variable", func(t *testing.T) {
		// Set a test environment variable
		os.Setenv("TEST_VAR", "test_value")
		defer os.Unsetenv("TEST_VAR")

		result := expandPath(`/path/{{.Env "TEST_VAR"}}/file.txt`)
		expected := "/path/test_value/file.txt"
		if result != expected {
			t.Errorf("Expected %s, got %s", expected, result)
		}
	})

	t.Run("invalid template fallback", func(t *testing.T) {
		original := "{{.InvalidTemplate"
		result := expandPath(original)
		// Should return original path when template is invalid
		if result != original {
			t.Errorf("Expected fallback to original path, got: %s", result)
		}
	})

	t.Run("real world example", func(t *testing.T) {
		homeDir, _ := os.UserHomeDir()
		result := expandPath(`{{.Home}}/compiled_notes/notes_{{.Date "2006-01-02_150405"}}.txt`)

		// Should start with home directory
		if !strings.HasPrefix(result, homeDir) {
			t.Errorf("Expected result to start with home dir")
		}

		// Should contain compiled_notes
		if !strings.Contains(result, "compiled_notes") {
			t.Errorf("Expected 'compiled_notes' in path")
		}

		// Should contain current date
		if !strings.Contains(result, time.Now().Format("2006-01-02")) {
			t.Errorf("Expected current date in path")
		}

		// Should not contain template syntax
		if strings.Contains(result, "{{") || strings.Contains(result, "}}") {
			t.Errorf("Template syntax not expanded: %s", result)
		}
	})
}

func TestTemplateData(t *testing.T) {
	homeDir, _ := os.UserHomeDir()
	data := TemplateData{
		Home: homeDir,
	}

	t.Run("Date function", func(t *testing.T) {
		result := data.Date("2006-01-02")
		expected := time.Now().Format("2006-01-02")
		if result != expected {
			t.Errorf("Expected %s, got %s", expected, result)
		}
	})

	t.Run("Env function", func(t *testing.T) {
		os.Setenv("TEST_ENV", "test_value")
		defer os.Unsetenv("TEST_ENV")

		result := data.Env("TEST_ENV")
		if result != "test_value" {
			t.Errorf("Expected 'test_value', got %s", result)
		}
	})

	t.Run("Env function with missing var", func(t *testing.T) {
		result := data.Env("NONEXISTENT_VAR")
		if result != "" {
			t.Errorf("Expected empty string for missing env var, got %s", result)
		}
	})
}
