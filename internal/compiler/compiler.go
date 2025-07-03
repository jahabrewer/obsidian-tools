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
	Profile      string
}

type FileConfig struct {
	OutputFilePath string              `yaml:"output_file_path"`
	GlobPatterns   []string            `yaml:"glob_patterns"`
	ListExcluded   bool                `yaml:"list_excluded"`
	ProfilesPath   string              `yaml:"profiles_path"`
	BaseProfile    *Profile            `yaml:"base"`
	Profiles       map[string]*Profile `yaml:"profiles"`
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

// Profile represents a configuration profile with specific settings
type Profile struct {
	Description  string            `yaml:"description"`
	Prompt       string            `yaml:"prompt"`
	GlobPatterns []string          `yaml:"globs"`
	Variables    map[string]string `yaml:"variables"`
}

// ProfilesConfig represents the structure of a standalone profiles configuration file
type ProfilesConfig struct {
	BaseProfile *Profile            `yaml:"base"`
	Profiles    map[string]*Profile `yaml:"profiles"`
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

	// Load profiles if available
	if err := c.loadProfiles(); err != nil {
		return fmt.Errorf("failed to load profiles: %w", err)
	}

	// Apply profile if specified
	if c.config.Profile != "" {
		if err := c.applyProfile(c.config.Profile); err != nil {
			return fmt.Errorf("failed to apply profile '%s': %w", c.config.Profile, err)
		}
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

func (c *Compiler) loadProfiles() error {
	// Load profiles from separate file if specified
	if c.fileConfig.ProfilesPath != "" {
		profilesPath := expandPath(c.fileConfig.ProfilesPath)
		if _, err := os.Stat(profilesPath); err != nil {
			if os.IsNotExist(err) {
				if c.config.Verbose {
					fmt.Printf("Profiles file not found: %s\n", profilesPath)
				}
				return nil
			}
			return fmt.Errorf("failed to access profiles file: %w", err)
		}

		data, err := os.ReadFile(profilesPath)
		if err != nil {
			return fmt.Errorf("failed to read profiles file: %w", err)
		}

		// Extract YAML from markdown file
		yamlContent, err := extractYAMLFromMarkdown(string(data))
		if err != nil {
			return fmt.Errorf("failed to extract YAML from profiles file: %w", err)
		}

		var profilesConfig ProfilesConfig
		if err := yaml.Unmarshal([]byte(yamlContent), &profilesConfig); err != nil {
			return fmt.Errorf("failed to parse profiles YAML: %w", err)
		}

		// Merge profiles from external file
		if c.fileConfig.Profiles == nil {
			c.fileConfig.Profiles = make(map[string]*Profile)
		}
		for name, profile := range profilesConfig.Profiles {
			c.fileConfig.Profiles[name] = profile
		}

		// Set base profile if not already set
		if c.fileConfig.BaseProfile == nil {
			c.fileConfig.BaseProfile = profilesConfig.BaseProfile
		}

		if c.config.Verbose {
			fmt.Printf("Loaded profiles from %s\n", profilesPath)
		}
	}

	return nil
}

// extractYAMLFromMarkdown extracts YAML content from a single YAML code block in a Markdown file
func extractYAMLFromMarkdown(content string) (string, error) {
	lines := strings.Split(content, "\n")
	var yamlLines []string
	var inYAMLBlock bool
	var inNonYAMLBlock bool
	var yamlBlockCount int
	var currentBlockStart int

	for i, line := range lines {
		trimmed := strings.TrimSpace(line)

		// Check if this line starts a code block
		if strings.HasPrefix(trimmed, "```") {
			if !inYAMLBlock && !inNonYAMLBlock {
				// Starting a new code block
				if trimmed == "```yaml" || trimmed == "```yml" || trimmed == "```" {
					inYAMLBlock = true
					yamlBlockCount++
					currentBlockStart = i
					continue
				} else {
					// This is a non-YAML code block (e.g., ```javascript)
					inNonYAMLBlock = true
					continue
				}
			} else {
				// Ending a code block
				if trimmed == "```" {
					inYAMLBlock = false
					inNonYAMLBlock = false
					continue
				}
			}
		}

		// If we're in a YAML block, collect the line
		if inYAMLBlock {
			yamlLines = append(yamlLines, line)
		}
	}

	// Validate that we found exactly one YAML block
	if yamlBlockCount == 0 {
		return "", fmt.Errorf("no YAML code blocks found in markdown file")
	}
	if yamlBlockCount > 1 {
		return "", fmt.Errorf("multiple YAML code blocks found in markdown file, expected exactly one")
	}

	// Check if we ended in a YAML block (unclosed)
	if inYAMLBlock {
		return "", fmt.Errorf("unclosed YAML code block starting at line %d", currentBlockStart+1)
	}

	yamlContent := strings.Join(yamlLines, "\n")
	if strings.TrimSpace(yamlContent) == "" {
		return "", fmt.Errorf("YAML code block is empty")
	}

	return yamlContent, nil
}

func (c *Compiler) applyProfile(profileName string) error {
	if c.fileConfig.Profiles == nil {
		return fmt.Errorf("no profiles defined")
	}

	profile, exists := c.fileConfig.Profiles[profileName]
	if !exists {
		return fmt.Errorf("profile '%s' not found", profileName)
	}

	// Apply profile settings, only if not already set by command line
	if len(c.config.GlobPatterns) == 0 {
		// Start with base profile globs if available
		var globPatterns []string
		if c.fileConfig.BaseProfile != nil && len(c.fileConfig.BaseProfile.GlobPatterns) > 0 {
			globPatterns = append(globPatterns, c.fileConfig.BaseProfile.GlobPatterns...)
		}

		// Add profile-specific globs
		if len(profile.GlobPatterns) > 0 {
			globPatterns = append(globPatterns, profile.GlobPatterns...)
		}

		c.config.GlobPatterns = globPatterns
	}

	if c.config.Verbose {
		fmt.Printf("Applied profile: %s\n", profileName)
		if profile.Description != "" {
			fmt.Printf("Description: %s\n", profile.Description)
		}
		if profile.Prompt != "" {
			fmt.Printf("Profile prompt: %s\n", profile.Prompt)
		}
	}

	return nil
}

func (c *Compiler) getProfilePrompt() string {
	if c.config.Profile == "" {
		return ""
	}

	profile, exists := c.fileConfig.Profiles[c.config.Profile]
	if !exists {
		return ""
	}

	prompt := profile.Prompt

	// Add base profile prompt if available
	if c.fileConfig.BaseProfile != nil && c.fileConfig.BaseProfile.Prompt != "" {
		basePrompt := c.fileConfig.BaseProfile.Prompt
		if prompt != "" {
			prompt = basePrompt + "\n\n" + prompt
		} else {
			prompt = basePrompt
		}
	}

	return prompt
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
	// Write profile prompt at the beginning if available
	profilePrompt := c.getProfilePrompt()
	if profilePrompt != "" {
		// Expand templates in the prompt
		expandedPrompt := c.expandTemplateWithProfile(profilePrompt)
		if _, err := fmt.Fprintf(outFile, "%s\n\n", expandedPrompt); err != nil {
			return 0, 0, err
		}
	}

	// Write context information
	if c.config.Profile != "" {
		if _, err := fmt.Fprintf(outFile, "---\nSYSTEM CONTEXT: Profile '%s'\n---\n\n", c.config.Profile); err != nil {
			return 0, 0, err
		}
	}

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
			// Normalize paths for comparison to handle cross-platform differences
			normalizedFile := filepath.ToSlash(file)
			normalizedPattern := filepath.ToSlash(excludePattern)

			// Use doublestar.Match for consistent pattern matching
			if matched, _ := doublestar.Match(normalizedPattern, normalizedFile); matched {
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

// expandTemplateWithProfile expands templates in strings with profile-specific variables
func (c *Compiler) expandTemplateWithProfile(text string) string {
	// Get home directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		homeDir = "~"
	}

	// Create template data with profile variables
	data := struct {
		TemplateData
		Vars map[string]string
	}{
		TemplateData: TemplateData{Home: homeDir},
		Vars:         make(map[string]string),
	}

	// Add profile variables if available
	if c.config.Profile != "" {
		if profile, exists := c.fileConfig.Profiles[c.config.Profile]; exists && profile.Variables != nil {
			for k, v := range profile.Variables {
				data.Vars[k] = v
			}
		}
	}

	// Parse and execute template
	tmpl, err := template.New("text").Parse(text)
	if err != nil {
		return text
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, data)
	if err != nil {
		return text
	}

	return buf.String()
}

// LoadConfigOnly loads configuration and profiles without running compilation
func (c *Compiler) LoadConfigOnly() error {
	// Load configuration
	if err := c.loadConfig(); err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	// Load profiles if available
	if err := c.loadProfiles(); err != nil {
		return fmt.Errorf("failed to load profiles: %w", err)
	}

	return nil
}

func (c *Compiler) ListAvailableProfiles() {
	if c.fileConfig.Profiles == nil || len(c.fileConfig.Profiles) == 0 {
		fmt.Println("No profiles defined.")
		return
	}

	fmt.Printf("Available profiles:\n\n")
	for name, profile := range c.fileConfig.Profiles {
		fmt.Printf("  %s", name)
		if profile.Description != "" {
			fmt.Printf(" - %s", profile.Description)
		}
		fmt.Println()
	}
}
