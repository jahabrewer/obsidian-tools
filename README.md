# obsidian-tools

This repository contains a collection of shell scripts designed to enhance your Obsidian experience.

## Scripts

### note-compiler.sh

A shell script for compiling markdown notes from multiple files into a single output file. Supports glob patterns, exclusions, YAML config, clipboard copy, and verbose output.

#### Features
- Compile markdown (`.md`) notes from multiple files into one file
- Supports glob pattern matching and exclusion patterns
- Verbose mode (`-v`) to list included files
- Copy output to clipboard (`-c`)
- Supports a YAML config file (`~/.note-compiler.yaml`) parsed with `yq`

#### Requirements
- Zsh shell
- [`yq`](https://github.com/mikefarah/yq) for YAML config parsing (only required if using a config file)
- (Optional) `pbcopy` for clipboard support (macOS)

#### Installation
Install `note-compiler` using Homebrew:

```sh
brew tap jahabrewer/obsidian-tools
brew install note-compiler
```

This will install the `note-compiler` script and make it available in your `$PATH`.

#### Usage
```sh
note-compiler.sh [options] <output-file> <source-glob> [<exclude-glob> ...]
```

If `output_file_path` and `glob_patterns` are defined in your configuration file (see "Configuration" below), you can also run the script with just options:
```sh
note-compiler.sh [options]
```

##### Options
- `-v, --verbose`    Log extra information (e.g., included files, config values, config sources)
- `-c, --clipboard`  Copy output to clipboard (macOS only)
- `-f, --config`     Specify an alternative config file (default: ~/.note-compiler.yaml)
- `--version`        Show version information and exit.

##### Example
```sh
note-compiler.sh -v ~/compiled_notes/notes_$(date +%Y-%m-%d_%H%M%S).txt "**/*.md" "!.obsidian/**"
```

Note on patterns: The script uses Zsh's extended globbing. Patterns starting with `!` (like `"!.obsidian/**"` in the example) are treated as exclusion patterns.

#### Configuration
Supports a YAML config file at `~/.note-compiler.yaml` for default options and patterns. See [sample-config.yaml](examples/sample-config.yaml) for an example.

#### Output Format
The script formats the output file with separators and source annotations for each included file:

```
---
source: /path/to/file.md
---

[Content of file.md]
```

### obsidian-backup.sh

A shell script to back up your Obsidian vault.
