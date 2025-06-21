# obsidian-tools

This repository contains tools designed to enhance your Obsidian experience.

## Tools

### note-compiler

A cross-platform Go CLI tool for compiling markdown notes from multiple files into a single output file. Supports glob patterns, exclusions, YAML config, clipboard copy, and verbose output.

#### Features
- Compile markdown (`.md`) notes from multiple files into one file
- Cross-platform support (macOS, Windows, Linux)
- Supports glob pattern matching and exclusion patterns
- Verbose mode (`-v`) to list included files
- Copy output to clipboard (`-c`)
- Supports a YAML config file (`~/.note-compiler.yaml`)
- Extensive test coverage
- Single binary distribution

#### Installation

##### Using Homebrew (macOS)
```sh
brew tap jahabrewer/note-compiler
brew install note-compiler
```

##### Manual Installation
Download the latest release for your platform from the [releases page](https://github.com/jahabrewer/note-compiler/releases).

#### Usage
```sh
note-compiler [flags] <output-file> <glob-pattern>...
```

If `output_file_path` and `glob_patterns` are defined in your configuration file (see "Configuration" below), you can also run with just options:
```sh
note-compiler [flags]
```

##### Options
- `-v, --verbose`    List all files included in the compilation
- `-c, --clipboard`  Copy the resulting file to clipboard
- `-f, --config`     Specify an alternative config file (default: ~/.note-compiler.yaml)
- `-h, --help`       Help for note-compiler

##### Commands
- `note-compiler version`     Show version information
- `note-compiler completion`  Generate autocompletion script
- `note-compiler help`        Help about any command

##### Examples
```sh
note-compiler -v ~/compiled_notes/notes_$(date +%Y-%m-%d_%H%M%S).txt "**/*.md" "!.obsidian/**"
note-compiler -c output.md "*.md"
note-compiler version
```

Note on patterns: The tool uses standard glob patterns. Patterns starting with `!` (like `"!.obsidian/**"` in the example) are treated as exclusion patterns.

#### Configuration
Supports a YAML config file at `~/.note-compiler.yaml` for default options and patterns. See [sample-config.yaml](examples/sample-config.yaml) for an example.

#### Output Format
The tool formats the output file with separators and source annotations for each included file:

```markdown
---
source: /path/to/file.md
---

[Content of file.md]

---
source: /path/to/another_file.md
---

[Content of another_file.md]
```

#### Development

##### Requirements
- Go 1.21 or later

##### Building from Source
```sh
make build
```

##### Running Tests
```sh
make test
```

##### Pre-Push Checks
To ensure code quality, run all checks before pushing:
```sh
make pre-push
```

This will run:
- Code formatting (`make fmt`)
- Linting (`make lint`)
- Tests (`make test`)

##### Git Hooks Setup
Install git hooks to automatically run pre-push checks:
```sh
./scripts/install-hooks.sh
```

After installation, these checks will run automatically before every push.

##### Creating a Release
```sh
make release
```

### obsidian-backup

A shell script to back up your Obsidian vault.

## License

See [LICENSE](LICENSE) file for details.
