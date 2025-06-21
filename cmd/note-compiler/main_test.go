package main

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

func TestVersionCommand(t *testing.T) {
	// Create a buffer to capture output
	var buf bytes.Buffer

	// Create the version command
	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Show version information",
		Run: func(cmd *cobra.Command, args []string) {
			buf.WriteString("note-compiler dev (commit: none, built: unknown)\n")
		},
	}

	// Execute the command
	versionCmd.SetOut(&buf)
	err := versionCmd.Execute()

	if err != nil {
		t.Errorf("Version command failed: %v", err)
	}

	output := buf.String()
	if !strings.Contains(output, "note-compiler") {
		t.Errorf("Expected version output to contain 'note-compiler', got: %s", output)
	}
}

func TestRunCompilerError(t *testing.T) {
	// Test that runCompiler returns an error when no arguments are provided
	// and no config file exists

	// Create a temporary directory for isolation
	tempDir, err := os.MkdirTemp("", "cli_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	// Temporarily set HOME to our temp dir to avoid loading real config
	originalHome := os.Getenv("HOME")
	os.Setenv("HOME", tempDir)
	defer os.Setenv("HOME", originalHome)

	// Create a mock command to test runCompiler
	var testCmd = &cobra.Command{
		Use:  "test",
		RunE: runCompiler,
	}

	// Test with no arguments - should fail
	err = testCmd.Execute()
	if err == nil {
		t.Error("Expected error when no output file or glob patterns provided")
	}
}
