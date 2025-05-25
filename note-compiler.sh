#!/bin/zsh

# Version placeholder - this will be replaced by the build process for releases
__SCRIPT_VERSION__="DEVELOPMENT_VERSION_#_NEEDS_REPLACEMENT_BY_BUILD_#"

# Default values for script logic
verbose=false
clipboard=false
config_file=~/.note-compiler.yaml

# Config file values (will be populated by read_config)
config_output_file_path=""
config_glob_patterns=()

# Variables for final determined values
output_file_path=""
typeset -a actual_glob_patterns_to_use

print_usage() {
    cat >&2 <<EOF
Usage: $0 [options] <output_file> <glob_pattern...>
       $0 [options]  # If output_file and glob_patterns are in config
       $0 --version  # Show version information

Options:
  -v, --verbose    List all files included in the compilation.
  -c, --clipboard  Copy the resulting file to clipboard.
  -f, --config F   Specify an alternative config file.
                   (Default: ~/.note-compiler.yaml)
  --version        Show version information and exit.

Example: $0 -v -f myconfig.yaml output.md "**/*.md" "!.obsidian/**"

Config file format (YAML, e.g., ~/.note-compiler.yaml):
  output_file_path: "path/to/output_\$(date +%Y-%m-%d).txt"
  glob_patterns:
    - "**/*.md"
    - "!.obsidian/**"
EOF
}

# Function to log resolved config values
log_config_values() {
    # This function is called from within read_config
    # and uses the state of variables at that point.
    echo "Config file values (from $config_file):"
    echo "  output_file_path (config): $config_output_file_path"
    echo "  glob_patterns (config):"
    if [[ ${#config_glob_patterns[@]} -gt 0 ]]; then
        for pattern in "${config_glob_patterns[@]}"; do
            echo "    $pattern"
        done
    else
        echo "    (none)"
    fi
}

# Function to read config file
read_config() {
    if [[ -f "$config_file" ]]; then
        echo "Loading configuration from $config_file"

        if command -v yq &> /dev/null; then
            # Load output_file_path from config only if yq finds it
            if yq -e '.output_file_path' "$config_file" &>/dev/null; then
                config_output_file_path=$(yq '.output_file_path' "$config_file")
            fi

            # Load glob_patterns from config only if yq finds them
            if yq -e '.glob_patterns[]' "$config_file" &>/dev/null; then
                config_glob_patterns=() # Clear to ensure fresh load
                while IFS= read -r pattern; do
                    config_glob_patterns+=("$pattern")
                done < <(yq '.glob_patterns[]' "$config_file")
            fi
        else
            echo "Warning: yq is not installed. Skipping config file loading for values other than file path if set by -f." >&2
        fi
    elif [[ "$config_file" != "$HOME/.note-compiler.yaml" ]]; then
        # If a custom config file was specified via -f but not found
        echo "Error: Specified config file '$config_file' not found." >&2
        exit 1
    fi
    # log_config_values # Call this after read_config, or it might show stale verbose/clipboard
}

# --- Argument Parsing with zparseopts ---
zmodload zsh/zutil

local -a _cli_verbose_flags
local -a _cli_clipboard_flags
local -a _cli_config_file_values # Stores value(s) from -f or --config
local -a _cli_version_flags

# -D: Deletes parsed options from $@, leaving only positional args.
# -F: Fails if an unknown option is encountered.
zparseopts -D -F -- \
  {v,-verbose}=_cli_verbose_flags \
  {c,-clipboard}=_cli_clipboard_flags \
  {f,-config}:=_cli_config_file_values \
  -version=_cli_version_flags \
  || { print_usage; exit 1; } # On failure (e.g., unknown option), print usage and exit.

# Handle --version flag first
if (( $#_cli_version_flags )); then
  if [[ "$__SCRIPT_VERSION__" != "DEVELOPMENT_VERSION_#_NEEDS_REPLACEMENT_BY_BUILD_#" ]]; then
    echo "$__SCRIPT_VERSION__"
  elif command -v git &>/dev/null; then # Check if git command exists
    if git rev-parse --is-inside-work-tree >/dev/null 2>&1; then # Check if inside a git repo
      version=$(git describe --tags --abbrev=0 2>/dev/null)
      if [[ -n "$version" ]]; then
        echo "$version (dev environment)"
      else
        commit_hash=$(git rev-parse --short HEAD 2>/dev/null)
        if [[ -n "$commit_hash" ]]; then
          echo "dev-$commit_hash (no tags found, dev environment)"
        else
          echo "unknown (git repository, but no tags or commits found, dev environment)"
        fi
      fi
    else
      echo "unknown (not a git repository, dev environment)"
    fi
  else
    echo "unknown (git command not found, dev environment - $__SCRIPT_VERSION__)"
  fi
  exit 0
fi

# Update script's main variables based on parsed CLI options
if (( $#_cli_verbose_flags )); then
  verbose=true
fi

if (( $#_cli_clipboard_flags )); then
  clipboard=true
fi

if (( $#_cli_config_file_values )); then
  # If -f/--config is used multiple times, the last one takes precedence.
  config_file="${_cli_config_file_values[-1]}"
fi
# --- End of Argument Parsing ---

# Log config values AFTER they've been fully resolved (CLI + config file)
if $verbose ; then # Show config log only if verbose is ultimately true
    # Print version information first in verbose mode
    echo "note-compiler $__SCRIPT_VERSION__"
    echo ""
fi

# Read config (CLI options like -f <new_config> will have updated `config_file` by now)
# `verbose` and `clipboard` may be further modified by read_config if not set by CLI.
read_config

# Log config values AFTER they've been fully resolved (CLI + config file)
if $verbose ; then # Show config log only if verbose is ultimately true
    log_config_values
fi

# Determine output file path
# $@ now contains only positional arguments. $1 is the potential output_file_path.
if [[ $# -ge 1 ]]; then
    output_file_path="$1"
    shift # Remove output_file_path from positional arguments, rest are globs
else
    if [[ -n "$config_output_file_path" ]]; then
        output_file_path="$config_output_file_path"
    else
        echo "Error: No output file specified and no output_file_path found in config." >&2
        echo "Please provide an output file as the first argument or set output_file_path in your config." >&2
        print_usage
        exit 1
    fi
fi

# Create output directory if it doesn't exist
output_dir=$(dirname "$output_file_path")
mkdir -p "$output_dir" || { echo "Error: Could not create directory $output_dir" >&2 ; exit 1; }

# Create/clear the output file (eval is used to allow dynamic paths like those with $(date))
output_file=$(eval echo "$output_file_path")
: > "$output_file" || { echo "Error: Could not write to output file $output_file" >&2 ; exit 1; }

# Determine glob patterns to use
# $@ at this point (after output_file_path potentially shifted) contains command-line glob patterns.
if [[ $# -gt 0 ]]; then # Command-line glob patterns were provided
    actual_glob_patterns_to_use=("$@")
else # No command-line glob patterns, try config
    if [[ ${#config_glob_patterns[@]} -gt 0 ]]; then
        actual_glob_patterns_to_use=("${config_glob_patterns[@]}")
    else
        echo "Error: No glob patterns specified and none found in config." >&2
        echo "Please provide at least one glob pattern or set glob_patterns in your config." >&2
        print_usage
        exit 1
    fi
fi

# Log resolved argument values only if verbose mode is enabled
if [[ $verbose == true ]]; then
    echo "\nResolved arguments for operation:"
    echo "  Output file: $output_file"
    echo "  Verbose mode: $verbose"
    echo "  Clipboard mode: $clipboard"
    echo "  Config file effectively used: $config_file"
    echo "  Glob patterns to be used (unexpanded):"
    for pattern in "${actual_glob_patterns_to_use[@]}"; do
        echo "    $pattern"
    done
    echo "" # Newline for better readability
fi

# Initialize counters
processed_files=0
excluded_file_count=0

# Initialize arrays for glob processing
typeset -a exclude_patterns
typeset -a matched_files

# Function to process each file
process_file() {
    local file="$1"

    if [[ $verbose == true ]]; then
        echo "Including file: $file"
    fi

    echo "---" >> "$output_file"
    echo "source: $file" >> "$output_file"
    echo "---" >> "$output_file"
    echo >> "$output_file"
    cat "$file" >> "$output_file"
    echo >> "$output_file"
    echo >> "$output_file"
    ((processed_files++))
}

# Enable extended globbing and null_glob for pattern matching
setopt extended_glob null_glob

# Process each glob pattern
for pattern in "${actual_glob_patterns_to_use[@]}"; do
    if [[ "$pattern" == !* ]]; then # Check if pattern starts with ! for exclusion
        exclude_patterns+=("${pattern#!}") # Remove the ! and add to exclusion array
    else
        # Expand the glob pattern directly
        # Using a subshell for `files=( ${~pattern} )` is safer if patterns are complex or numerous
        local current_files
        current_files=(${~pattern})
        for file in ${current_files[@]}; do
            if [[ -f $file ]]; then # Ensure it's a file
                matched_files+=("$file")
            fi
        done
    fi
done

# Deduplicate matched_files (in case patterns overlap)
typeset -U matched_files

# Process files, excluding any that match exclusion patterns
for file in "${matched_files[@]}"; do
    excluded=false
    for exclude_pattern in "${exclude_patterns[@]}"; do
        # Ensure exclude pattern is also tilde-expanded if necessary for matching
        if [[ $file == ${~exclude_pattern} ]]; then
            excluded=true
            ((excluded_file_count++))
            if [[ $verbose == true ]]; then
                echo "Excluding file (matches '$exclude_pattern'): $file"
            fi
            break
        fi
    done

    if [[ $excluded == false ]]; then
        process_file "$file"
    fi
done

echo "\nInitial files matching include patterns: ${#matched_files[@]}"
echo "Number of files excluded: $excluded_file_count"
echo "Successfully processed $processed_files files into $output_file"

# Get and display the output file size
if [[ -f "$output_file" ]]; then
    file_size=$(du -h "$output_file" | cut -f1)
    echo "Output file size: $file_size"
else
    echo "Output file $output_file not created or empty."
fi

# Copy to clipboard if requested
if [[ $clipboard == true ]]; then
    if command -v pbcopy &> /dev/null; then
        pbcopy < "$output_file"
        echo "Content copied to clipboard"
    else
        echo "Warning: pbcopy command not found. Cannot copy to clipboard." >&2
    fi
fi