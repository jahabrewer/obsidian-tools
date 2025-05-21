#!/bin/zsh


# Default values
verbose=false
clipboard=false
config_file=~/.note-compiler.yaml

# Config file values (set via ~/.note-compiler.config)
config_output_file_path=""
config_glob_patterns=()

# Command-line values (set via arguments)
output_file_path=""
glob_patterns=()

# Function to log resolved config values
log_config_values() {
    echo "Config file values:"
    echo "  verbose: $verbose"
    echo "  clipboard: $clipboard"
    echo "  output_file_path: $config_output_file_path"
    echo "  glob_patterns:"
    for pattern in "${config_glob_patterns[@]}"; do
        echo "    $pattern"
    done
}

# Function to read config file
read_config() {
    if [[ -f "$config_file" ]]; then
        echo "Loading configuration from $config_file"
        
        # Only use yq if it's available
        if command -v yq &> /dev/null; then
            if yq -e '.verbose' "$config_file" &>/dev/null; then
                local config_verbose=$(yq '.verbose' "$config_file")
                [[ "$config_verbose" == "true" ]] && verbose=true
            fi
            
            if yq -e '.clipboard' "$config_file" &>/dev/null; then
                local config_clipboard=$(yq '.clipboard' "$config_file")
                [[ "$config_clipboard" == "true" ]] && clipboard=true
            fi
            
            if yq -e '.output_file_path' "$config_file" &>/dev/null; then
                config_output_file_path=$(yq '.output_file_path' "$config_file")
            fi
            
            if yq -e '.glob_patterns[]' "$config_file" &>/dev/null; then
                config_glob_patterns=()
                while IFS= read -r pattern; do
                    config_glob_patterns+=("$pattern")
                done < <(yq '.glob_patterns[]' "$config_file")
            fi
        else
            echo "Warning: yq is not installed, skipping config file loading."
        fi
    fi
    log_config_values
}

# Check if at least one argument is provided
if [[ $# -eq 0 ]]; then
    echo "Usage: $0 [options] <output_file> <glob_pattern...>"
    echo "Options:"
    echo "  -v, --verbose    List all files included in the compilation"
    echo "  -c, --clipboard  Copy the resulting file to clipboard"
    echo "  -f, --config     Specify an alternative config file (default: ~/.note-compiler.config)"
    echo
    echo "Example: $0 [-cv] output.md \"**/*.md\" \"!.obsidian/**\""
    echo
    echo "Config file (~/.note-compiler.yaml) format (YAML):"
    echo "  verbose: true|false"
    echo "  clipboard: true|false"
    echo "  output_file_path: \"path/to/output_\$(date +%Y-%m-%d).txt\""
    echo "  glob_patterns:"
    echo "    - \"**/*.md\""
    echo "    - \"!.obsidian/**\""
    exit 1
fi

# Parse options
while [[ "$1" == -* ]]; do
    case "$1" in
        -v|--verbose)
            verbose=true
            shift
            ;;
        -c|--clipboard)
            clipboard=true
            shift
            ;;
        -f|--config)
            config_file="$2"
            shift 2
            ;;
        *)
            echo "Unknown option: $1"
            exit 1
            ;;
    esac
done

# Read config (CLI options will override these)
read_config

# Determine output file 
if [[ $# -ge 1 ]]; then
    # Use provided output file from command line
    output_file_path="$1"
    shift
else
    # Use output file path from config if available
    if [[ -n "$config_output_file_path" ]]; then
        output_file_path="$config_output_file_path"
    else
        echo "Error: No output file specified and no output_file_path found in config."
        echo "Please provide an output file as the first argument or set output_file_path in your config."
        exit 1
    fi
fi

# Create output directory if it doesn't exist
output_dir=$(dirname "$output_file_path")
mkdir -p "$output_dir" || exit 1

# Create/clear the output file
output_file=$(eval echo "$output_file_path")
: > "$output_file"

# Log resolved argument values
echo "Resolved arguments:"
echo "Output file: $output_file"
echo "Verbose mode: $verbose"
echo "Clipboard mode: $clipboard"

# Log glob patterns
echo "Glob patterns:"
# Get the final patterns that will be used
if [[ $# -gt 0 ]]; then
    # Use command-line patterns if provided
    for pattern in "$@"; do
        echo "  $pattern"
    done
else
    # Use config patterns if no command-line patterns

    # Use config patterns if no command-line patterns
    for pattern in "${config_glob_patterns[@]}"; do
        echo "  $pattern"
    done
fi

# Initialize counters
total_matches=0
processed_files=0

# Initialize arrays
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

# Enable extended globbing and null_glob
setopt extended_glob null_glob

# Use passed glob patterns or config patterns if none provided
if [[ $# -eq 0 ]]; then
    if [[ ${#config_glob_patterns[@]} -gt 0 ]]; then
        set -- "${config_glob_patterns[@]}"
    else
        echo "Error: No glob patterns specified and none found in config."
        echo "Please provide at least one glob pattern or set glob_patterns in your config."
        exit 1
    fi
fi

# For each remaining argument (glob pattern)
for pattern in "$@"; do
    # Check if pattern starts with ! for exclusion
    if [[ "$pattern" = !* ]]; then
        # Remove the ! and add it to exclusion array
        exclude_patterns+=("${pattern#!}")
    else
        # Expand the glob pattern directly
        files=(${~pattern})
        for file in $files; do
            if [[ -f $file ]]; then
                matched_files+=($file)
            fi
        done
    fi
done

# Process files, excluding any that match exclusion patterns
for file in $matched_files; do
    excluded=false
    for exclude in $exclude_patterns; do
        if [[ $file = ${~exclude} ]]; then
            excluded=true
            break
        fi
    done

    if [[ $excluded == false ]]; then
        process_file "$file"
        ((total_matches++))
    fi
done

echo "Found $total_matches matching files"
echo "Processed $processed_files files into $output_file"

# Get and display the output file size
file_size=$(du -h "$output_file" | cut -f1)
echo "Output file size: $file_size"

# Copy to clipboard if requested
if [[ $clipboard == true ]]; then
    pbcopy < "$output_file"
    echo "Content copied to clipboard"
fi
