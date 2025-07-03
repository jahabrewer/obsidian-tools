# Profiles Feature Documentation

## Overview

The profiles feature allows you to define reusable configurations for different compilation contexts. Instead of manually specifying glob patterns and output files each time, you can create named profiles that encapsulate common use cases.

## Key Features

- **Named Profiles**: Define reusable configurations with descriptive names
- **Profile Inheritance**: Base profiles that are inherited by all other profiles
- **Custom Prompts**: Each profile can specify a custom prompt for AI interactions
- **Template Variables**: Use variables in prompts and paths that are expanded at runtime
- **In-Vault Configuration**: Store profiles in your Obsidian vault as Markdown files for portability and version control
- **Command-Line Integration**: Easy profile selection via command-line flags
- **Obsidian-Native Editing**: Edit profiles directly in Obsidian with syntax highlighting and documentation

## Configuration

### Option 1: External Profiles File (Recommended)

Create a Markdown file with a YAML code block in your vault and reference it in your main config:

**~/.note-compiler.yaml**
```yaml
profiles_path: "{{.Home}}/vault/Configs/compiler-profiles.md"
```

**In your vault: Configs/compiler-profiles.md**
```markdown
# Profile Configuration

This is my profiles configuration for note-compiler.

```yaml
# Base profile inherited by all others
base:
  prompt: "Use these notes as context to answer questions and provide helpful responses."
  globs:
    - "{{.Home}}/vault/People/AI Instructions.md"
    - "!**/_resources/**"

profiles:
  one-on-one-prep:
    description: "Prepares context for my 1-on-1 with my manager"
    prompt: "Summarize my recent work, blockers, and talking points for my upcoming 1-on-1."
    globs:
      - "{{.Home}}/vault/Meeting notes/Manager 1 on 1/*.md"
      - "{{.Home}}/vault/Todo.md"
```​

You can add documentation, notes, and context around the YAML configuration.
```

### Option 2: Inline Profiles

Define profiles directly in your main configuration file:

**~/.note-compiler.yaml**
```yaml
base:
  prompt: "Use these notes as context to answer questions and provide helpful responses."
  globs:
    - "{{.Home}}/vault/People/AI Instructions.md"

profiles:
  kb-only:
    description: "Knowledge base only"
    prompt: "Based on my knowledge base, help me answer technical questions."
    globs:
      - "{{.Home}}/vault/KB/**/*.md"
```

## Markdown File Requirements

When using external profiles files (Option 1), the Markdown file must contain:

- **Exactly one YAML code block**: The file is parsed to find YAML code blocks (marked with ` ```yaml `, ` ```yml `, or ` ``` `)
- **Valid YAML syntax**: The content within the code block must be valid YAML
- **Single configuration**: Multiple YAML code blocks will cause an error

### Valid Examples

✅ **With language specifier:**
```markdown
# My Profiles

```yaml
profiles:
  test:
    description: "Test profile"
```​
```

✅ **Without language specifier:**
```markdown
# My Profiles

```
profiles:
  test:
    description: "Test profile"
```​
```

✅ **With documentation:**
```markdown
# Profile Configuration

This file defines my compilation contexts.

## Available Profiles

```yaml
base:
  prompt: "Use these notes as context to answer questions and provide helpful responses."
profiles:
  kb-only:
    description: "Knowledge base articles only"
    globs:
      - "{{.Home}}/vault/KB/**/*.md"
```​

The `kb-only` profile focuses on technical documentation.
```

### Invalid Examples

❌ **Multiple YAML blocks:**
```markdown
```yaml
base:
  prompt: "System Ready"
```​

```yaml
profiles:
  test:
    description: "Test"
```​
```

❌ **No YAML blocks:**
```markdown
# Just text with no code blocks
```

❌ **Empty YAML block:**
```markdown
```yaml
```​
```

## Usage Examples

### Basic Profile Usage

```bash
# Use a specific profile
note-compiler --profile one-on-one-prep

# Use a profile with verbose output
note-compiler --profile kb-only -v

# List all available profiles
note-compiler --list-profiles
```

### Profile with Custom Output

```bash
# Use profile's globs but specify custom output file
note-compiler --profile kb-only ~/custom-output.txt
```

### Override Profile Settings

```bash
# Command-line arguments override profile settings
note-compiler --profile kb-only ~/output.txt "custom/*.md"
```

## Profile Structure

Each profile can contain the following fields:

```yaml
profile-name:
  description: "Human-readable description of the profile"  # Optional
  prompt: "Custom prompt for AI interactions"               # Optional
  globs: ["list", "of", "glob", "patterns"]                # Optional
  variables:                                                # Optional
    var_name: "value"
    feature_name: "Smart Query"
```

## Template Variables

Profiles support template variables that are expanded at runtime:

### Standard Variables

- `{{.Home}}` - User's home directory
- `{{.Date "format"}}` - Current date/time with Go time format
- `{{.Env "VAR"}}` - Environment variable

### Profile-Specific Variables

- `{{.Vars.variable_name}}` - Custom variables defined in the profile

### Example with Variables

```yaml
profiles:
  feature-dive:
    description: "Deep dive into a specific feature"
    prompt: "Analyze the {{.Vars.feature_name}} feature in detail"
    globs:
      - "{{.Home}}/vault/Tickets/*{{.Vars.feature_name}}*.md"
      - "{{.Home}}/vault/KB/*{{.Vars.feature_name}}*.md"
    variables:
      feature_name: "Search Feature"  # Default value
```

## Profile Inheritance

The `base` profile is automatically inherited by all other profiles:

```yaml
base:
  prompt: "Use these notes as context to answer questions and provide helpful responses."
  globs:
    - "{{.Home}}/vault/People/AI Instructions.md"
    - "!**/_resources/**"

profiles:
  my-profile:
    prompt: "Additional prompt text"
    globs:
      - "{{.Home}}/vault/specific/*.md"
```

When using `my-profile`, the final configuration will be:
- **Prompt**: "Use these notes as context to answer questions and provide helpful responses.\n\nAdditional prompt text"
- **Globs**: All globs from base + profile-specific globs

## Output Format

When using profiles, the compiled output includes:

1. **Profile Prompt**: The combined base and profile prompts at the beginning
2. **Context Information**: System context showing which profile was used
3. **File Contents**: The actual compiled files

Example output:
```
Use these notes as context to answer questions and provide helpful responses.

Summarize my recent work, blockers, and talking points for my upcoming 1-on-1.

---
SYSTEM CONTEXT: Profile 'one-on-one-prep'
---

---
source: /path/to/file.md
---
[File contents...]
```

## Best Practices

### 1. Use Descriptive Names

```yaml
# Good
one-on-one-prep:
  description: "Prepares context for my 1-on-1 with my manager"

# Avoid
prep1:
  description: "Some preparation"
```

### 2. Leverage the Base Profile

Put common settings in the base profile to avoid repetition:

```yaml
base:
  globs:
    - "{{.Home}}/vault/People/AI Instructions.md"
    - "!**/_resources/**"
    - "!{{.Home}}/vault/Archive/**"
```

### 3. Use Template Variables for Flexibility

```yaml
profiles:
  monthly-review:
    globs:
      - "{{.Home}}/vault/Status/{{.Date \"2006-01\"}}*.md"
      - "{{.Home}}/vault/Status/{{.Date \"2006-02\"}}*.md"
```

### 4. Store Profiles in Your Vault as Markdown

Use the `profiles_path` approach with Markdown files:

```yaml
profiles_path: "{{.Home}}/vault/Configs/compiler-profiles.md"
```

Benefits:
- ✅ **Portability**: Profiles move with your vault
- ✅ **Version Control**: Profiles are tracked in Git
- ✅ **Easy Editing**: Edit profiles directly in Obsidian with syntax highlighting
- ✅ **Documentation**: Add context and notes around the configuration
- ✅ **Preview**: See rendered output in Obsidian's preview mode

### 5. Document Your Profiles

Use Markdown to document your profiles:

```markdown
# Profile Configuration

This file defines my compilation contexts for different workflows.

## Profile Usage

- **one-on-one-prep**: Use before weekly 1-on-1 meetings
- **on-call-prep**: Use when starting on-call rotation
- **feature-dive**: Use when analyzing specific features

```yaml
# Configuration goes here
```​

## Notes

Remember to update the feature-dive variables when working on new features.
```

## Error Handling

The compiler provides clear error messages for common issues:

```bash
# Profile not found
$ note-compiler --profile nonexistent
Error: failed to apply profile 'nonexistent': profile 'nonexistent' not found

# No profiles defined
$ note-compiler --list-profiles
No profiles defined.

# No YAML blocks in Markdown file
$ note-compiler --profile test
Error: failed to extract YAML from profiles file: no YAML code blocks found in markdown file

# Multiple YAML blocks
$ note-compiler --profile test
Error: failed to extract YAML from profiles file: multiple YAML code blocks found in markdown file, expected exactly one

# Invalid YAML syntax
$ note-compiler --profile test
Error: failed to load profiles: failed to parse profiles YAML: yaml: line 5: ...
```

## Integration with Existing Workflow

The profiles feature is fully backward compatible:

```bash
# Still works exactly as before
note-compiler ~/output.txt "**/*.md" "!Archive/**"

# New profile-based approach
note-compiler --profile default
```

## Advanced Examples

### Conditional Profile Usage

You can create profiles for different scenarios:

```yaml
profiles:
  # For quick reviews
  recent-only:
    description: "Recent work only"
    globs:
      - "{{.Home}}/vault/Status/{{.Date \"2006-01\"}}*.md"
      - "{{.Home}}/vault/Todo.md"
  
  # For comprehensive analysis
  full-context:
    description: "Complete vault context"
    globs:
      - "{{.Home}}/vault/**/*.md"
      - "!{{.Home}}/vault/Archive/**"
  
  # For specific topics
  technical-deep-dive:
    description: "Technical documentation and tickets"
    globs:
      - "{{.Home}}/vault/KB/**/*.md"
      - "{{.Home}}/vault/Tickets/**/*.md"
```

### Dynamic Date-Based Profiles

```yaml
profiles:
  this-week:
    description: "This week's work"
    globs:
      - "{{.Home}}/vault/Status/{{.Date \"2006-01-02\"}}*.md"
      # Note: This gives you today's date - you might want to adjust the logic
  
  this-month:
    description: "This month's work"
    globs:
      - "{{.Home}}/vault/Status/{{.Date \"2006-01\"}}*.md"
```

## Troubleshooting

### Profile Not Loading

1. Check the profiles file path in your config
2. Verify the Markdown file has exactly one YAML code block
3. Use `-v` flag to see detailed loading information

### YAML Extraction Errors

1. Ensure you have exactly one YAML code block in your Markdown file
2. Check that the code block uses ` ```yaml `, ` ```yml `, or ` ``` `
3. Verify the YAML code block is properly closed with ` ``` `

### Templates Not Expanding

1. Verify template syntax: `{{.Home}}`, `{{.Date "format"}}`, `{{.Vars.name}}`
2. Check that variable names match exactly
3. Ensure variables are defined in the profile

### Unexpected File Inclusion

1. Check base profile globs - they're included in all profiles
2. Verify exclusion patterns start with `!`
3. Use `-v` flag to see which files are included/excluded 