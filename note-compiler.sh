#!/bin/zsh

# Check if at least one argument is provided
if [[ $# -eq 0 ]]; then
    echo "Usage: $0 <output_file> <glob_pattern...>"
    echo "Example: $0 output.md \"**/*.md\" \"!.obsidian/**\""
    exit 1
fi

# First argument is the output file
output_file="$1"
shift

# Create output directory if it doesn't exist
output_dir=$(dirname "$output_file")
mkdir -p "$output_dir"

# Create/clear the output file
: > "$output_file"

# Initialize counters
total_matches=0
processed_files=0

# Initialize arrays
typeset -a exclude_patterns
typeset -a matched_files

# Function to process each file
process_file() {
    local file="$1"
    
    echo "---" >> "$output_file"
    echo "source: ${file#./}" >> "$output_file"
    echo "---" >> "$output_file"
    echo >> "$output_file"
    cat "$file" >> "$output_file"
    echo >> "$output_file"
    echo >> "$output_file"
    ((processed_files++))
}

# Enable extended globbing and null_glob
setopt extended_glob null_glob

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
