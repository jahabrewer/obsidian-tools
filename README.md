# note-compiler

A shell script for compiling markdown notes from multiple files into a single output file. Supports glob patterns, exclusions, YAML config, clipboard copy, and verbose output.

## Features
- Compile markdown (`.md`) notes from multiple files into one file
- Supports glob pattern matching and exclusion patterns
- Verbose mode (`-v`) to list included files
- Copy output to clipboard (`-c`)
- Supports a YAML config file (`~/.note-compiler.yaml`) parsed with `yq`

## Requirements
- Zsh shell
- [`yq`](https://github.com/mikefarah/yq) for YAML config parsing (only required if using a config file)
- (Optional) `pbcopy` for clipboard support (macOS)

## Installation
1. Download or clone this repository:
   ```sh
   git clone https://github.com/yourusername/note-compiler.git
   cd note-compiler
   chmod +x note-compiler.sh
   ```
2. (Optional) Move to a directory in your `$PATH`:
   ```sh
   mv note-compiler.sh /usr/local/bin/note-compiler
   ```

## Usage
```sh
note-compiler.sh [options] <output-file> <source-glob> [<exclude-glob> ...]
```

### Options
- `-v, --verbose`    List all files included in the compilation
- `-c, --clipboard`  Copy output to clipboard (macOS only)
- `-f, --config`     Specify an alternative config file (default: ~/.note-compiler.yaml)

### Example
```zsh
note-compiler.sh -v ~/compiled_notes/notes_$(date +%Y-%m-%d_%H%M%S).txt "**/*.md" "!.obsidian/**"
```

## Configuration
Supports a YAML config file at `~/.note-compiler.yaml` for default options and patterns. See [sample-config.yaml](sample-config.yaml) for an example.

## Output Format
The script formats the output file with separators and source annotations for each included file:

```
---
source: /path/to/file.md
---

[Content of file.md]
