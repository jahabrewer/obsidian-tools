package compiler

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	"github.com/atotto/clipboard"
	"github.com/bmatcuk/doublestar/v4"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Verbose      bool
	Clipboard    bool
	ListExcluded bool
	ConfigFile   string
	OutputFile   string
	GlobPatterns []string
}

type FileConfig struct {
	OutputFilePath string   `yaml:"output_file_path"`
	GlobPatterns   []string `yaml:"glob_patterns"`
	ListExcluded   bool     `yaml:"list_excluded"`
}

type Compiler struct {
	config     *Config
	fileConfig *FileConfig
}

// TemplateData provides variables for Go template expansion in paths
type TemplateData struct {
	Home string
}

// Date returns the current time formatted with Go's time format
func (td TemplateData) Date(format string) string {
	return time.Now().Format(format)
}

// Env returns the value of an environment variable
func (td TemplateData) Env(key string) string {
	return os.Getenv(key)
}

func New(config *Config) *Compiler {
	return &Compiler{
		config:     config,
		fileConfig: &FileConfig{},
	}
}

func (c *Compiler) Run() error {
	// Load configuration
	if err := c.loadConfig(); err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	// Apply file config values if not set via command line
	if !c.config.ListExcluded && c.fileConfig.ListExcluded {
		c.config.ListExcluded = c.fileConfig.ListExcluded
	}

	// Determine output file
	outputFile, err := c.determineOutputFile()
	if err != nil {
		return err
	}

	// Determine glob patterns
	globPatterns, err := c.determineGlobPatterns()
	if err != nil {
		return err
	}

	if c.config.Verbose {
		fmt.Printf("note-compiler (version info would be here)\n")
		fmt.Printf("Output file: %s\n", outputFile)
		fmt.Printf("Glob patterns: %v\n", globPatterns)
		fmt.Println()
	}

	// Create output directory
	if mkdirErr := os.MkdirAll(filepath.Dir(outputFile), 0755); mkdirErr != nil {
		return fmt.Errorf("failed to create output directory: %w", mkdirErr)
	}

	// Create/clear output file
	outFile, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer outFile.Close()

	// Process files
	processedCount, excludedCount, err := c.processFiles(outFile, globPatterns)
	if err != nil {
		return err
	}

	fmt.Printf("Successfully processed %d files into %s\n", processedCount, outputFile)
	fmt.Printf("Number of files excluded: %d\n", excludedCount)

	// Show file size
	if stat, err := outFile.Stat(); err == nil {
		fmt.Printf("Output file size: %d bytes\n", stat.Size())
	}

	// Copy to clipboard if requested
	if c.config.Clipboard {
		if err := c.copyToClipboard(outputFile); err != nil {
			fmt.Printf("Warning: failed to copy to clipboard: %v\n", err)
		} else {
			fmt.Println("Content copied to clipboard")
		}
	}

	return nil
}

func (c *Compiler) loadConfig() error {
	configFile := c.config.ConfigFile
	if configFile == "" {
		configFile = filepath.Join(os.Getenv("HOME"), ".note-compiler.yaml")
	}

	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		if c.config.ConfigFile != "" {
			return fmt.Errorf("specified config file not found: %s", configFile)
		}
		// Default config file doesn't exist, that's ok
		return nil
	}

	data, err := os.ReadFile(configFile)
	if err != nil {
		return fmt.Errorf("failed to read config file: %w", err)
	}

	if err := yaml.Unmarshal(data, c.fileConfig); err != nil {
		return fmt.Errorf("failed to parse config file: %w", err)
	}

	if c.config.Verbose {
		fmt.Printf("Loading configuration from %s\n", configFile)
		fmt.Printf("Config file values:\n")
		fmt.Printf("  output_file_path: %s\n", c.fileConfig.OutputFilePath)
		fmt.Printf("  glob_patterns: %v\n", c.fileConfig.GlobPatterns)
		fmt.Printf("  list_excluded: %t\n", c.fileConfig.ListExcluded)
	}

	return nil
}

func (c *Compiler) determineOutputFile() (string, error) {
	var outputPath string

	if c.config.OutputFile != "" {
		outputPath = c.config.OutputFile
	} else if c.fileConfig.OutputFilePath != "" {
		outputPath = c.fileConfig.OutputFilePath
	} else {
		return "", fmt.Errorf("no output file specified and no output_file_path found in config")
	}

	// Expand Go templates in path
	return expandPath(outputPath), nil
}

func (c *Compiler) determineGlobPatterns() ([]string, error) {
	var patterns []string

	if len(c.config.GlobPatterns) > 0 {
		patterns = c.config.GlobPatterns
	} else if len(c.fileConfig.GlobPatterns) > 0 {
		patterns = c.fileConfig.GlobPatterns
	} else {
		return nil, fmt.Errorf("no glob patterns specified and none found in config")
	}

	// Expand Go templates in glob patterns
	var expandedPatterns []string
	for _, pattern := range patterns {
		expandedPatterns = append(expandedPatterns, expandPath(pattern))
	}

	return expandedPatterns, nil
}

func (c *Compiler) processFiles(outFile io.Writer, globPatterns []string) (int, int, error) {
	var includePatterns, excludePatterns []string

	for _, pattern := range globPatterns {
		if strings.HasPrefix(pattern, "!") {
			excludePatterns = append(excludePatterns, pattern[1:])
		} else {
			includePatterns = append(includePatterns, pattern)
		}
	}

	var allFiles []string
	for _, pattern := range includePatterns {
		// Use doublestar for proper ** recursive globbing
		matches, err := doublestar.FilepathGlob(pattern)
		if err != nil {
			return 0, 0, fmt.Errorf("failed to glob pattern %s: %w", pattern, err)
		}
		allFiles = append(allFiles, matches...)
	}

	// Remove duplicates
	fileSet := make(map[string]bool)
	var uniqueFiles []string
	for _, file := range allFiles {
		if !fileSet[file] {
			fileSet[file] = true
			uniqueFiles = append(uniqueFiles, file)
		}
	}

	processedCount := 0
	excludedCount := 0

	for _, file := range uniqueFiles {
		// Check if file should be excluded using doublestar for exclusion patterns too
		excluded := false
		var matchedPattern string
		for _, excludePattern := range excludePatterns {
			// Use doublestar.Match for consistent pattern matching
			if matched, _ := doublestar.Match(excludePattern, file); matched {
				excluded = true
				matchedPattern = excludePattern
				excludedCount++
				break
			}
		}

		if excluded {
			// Show excluded files if verbose mode is on OR if ListExcluded is specifically requested
			if c.config.Verbose || c.config.ListExcluded {
				fmt.Printf("Excluding file (matches '%s'): %s\n", matchedPattern, file)
			}
		} else {
			if err := c.processFile(outFile, file); err != nil {
				return processedCount, excludedCount, err
			}
			processedCount++
		}
	}

	return processedCount, excludedCount, nil
}

func (c *Compiler) processFile(outFile io.Writer, filename string) error {
	if c.config.Verbose {
		fmt.Printf("Including file: %s\n", filename)
	}

	// Write separator and source info
	if _, err := fmt.Fprintf(outFile, "---\nsource: %s\n---\n\n", filename); err != nil {
		return err
	}

	// Read and write file content
	content, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed to read file %s: %w", filename, err)
	}

	if _, err := outFile.Write(content); err != nil {
		return err
	}

	// Add some spacing
	if _, err := outFile.Write([]byte("\n\n")); err != nil {
		return err
	}

	return nil
}

func (c *Compiler) copyToClipboard(filename string) error {
	content, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	return clipboard.WriteAll(string(content))
}

// expandPath expands Go templates in file paths
// Supports {{.Home}}, {{.Date "2006-01-02"}}, {{.Env "VAR"}}
func expandPath(path string) string {
	// Get home directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		homeDir = "~" // fallback
	}

	// Create template data
	data := TemplateData{
		Home: homeDir,
	}

	// Parse and execute template
	tmpl, err := template.New("path").Parse(path)
	if err != nil {
		// If template parsing fails, return original path
		return path
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, data)
	if err != nil {
		// If template execution fails, return original path
		return path
	}

	// Clean the path to handle cross-platform path separators
	return filepath.Clean(buf.String())
}
