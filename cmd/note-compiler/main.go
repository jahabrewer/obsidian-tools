package main

import (
	"fmt"
	"os"

	"github.com/jahabrewer/note-compiler/internal/compiler"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "note-compiler [flags] <output-file> <glob-pattern>...",
		Short: "Compile markdown notes from multiple files into a single output file",
		Long: `note-compiler compiles markdown notes from multiple files into a single output file.
Supports glob patterns, exclusions, YAML config, clipboard copy, and verbose output.`,
		Example: `  note-compiler -v ~/compiled_notes/notes_$(date +%Y-%m-%d_%H%M%S).txt "**/*.md" "!.obsidian/**"
  note-compiler -c output.md "*.md"`,
		Args: cobra.MinimumNArgs(0), // Allow 0 args for config-only mode
		RunE: runCompiler,
	}

	// Flags
	rootCmd.Flags().BoolP("verbose", "v", false, "list all files included in the compilation")
	rootCmd.Flags().BoolP("clipboard", "c", false, "copy the resulting file to clipboard")
	rootCmd.Flags().BoolP("list-excluded", "e", false, "list files excluded from compilation")
	rootCmd.Flags().StringP("config", "f", "", "specify an alternative config file (default: ~/.note-compiler.yaml)")

	// Version command
	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Show version information",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("note-compiler %s (commit: %s, built: %s)\n", version, commit, date)
		},
	}
	rootCmd.AddCommand(versionCmd)

	// Bind flags to viper (ignore errors as they are not critical for flag binding)
	_ = viper.BindPFlag("verbose", rootCmd.Flags().Lookup("verbose"))
	_ = viper.BindPFlag("clipboard", rootCmd.Flags().Lookup("clipboard"))
	_ = viper.BindPFlag("list-excluded", rootCmd.Flags().Lookup("list-excluded"))
	_ = viper.BindPFlag("config", rootCmd.Flags().Lookup("config"))

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func runCompiler(cmd *cobra.Command, args []string) error {
	config := &compiler.Config{
		Verbose:      viper.GetBool("verbose"),
		Clipboard:    viper.GetBool("clipboard"),
		ListExcluded: viper.GetBool("list-excluded"),
		ConfigFile:   viper.GetString("config"),
	}

	// Parse arguments
	if len(args) > 0 {
		config.OutputFile = args[0]
		config.GlobPatterns = args[1:]
	}

	compiler := compiler.New(config)
	return compiler.Run()
}
