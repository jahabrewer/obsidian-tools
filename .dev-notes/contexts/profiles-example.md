# Note Compiler Profiles Configuration

This file defines the compilation profiles for the note-compiler tool. It should be stored in your Obsidian vault for portability and version control.

You can edit this file directly in Obsidian with full syntax highlighting and preview support.

```yaml
# Base profile inherited by all other profiles
base:
  prompt: "Use these notes as context to answer questions and provide helpful responses."
  globs:
    - "{{.Home}}/vault/People/AI Instructions.md"
    - "!**/_resources/**"  # Universal exclude

# Profile definitions
profiles:
  # Default profile if none is specified
  default:
    description: "The standard full-vault context, excluding archives."
    globs:
      - "{{.Home}}/vault/**/*.md"
      - "!{{.Home}}/vault/Archive/**/*.md"
      - "!{{.Home}}/vault/Tickets/xml/**"

  # Profile for 1-on-1 meeting preparation
  one-on-one-prep:
    description: "Prepares context for my 1-on-1 with my manager."
    prompt: "Summarize my recent work, blockers, and talking points for my upcoming 1-on-1. What questions should I ask my manager based on recent feedback?"
    globs:
      - "{{.Home}}/vault/Meeting notes/Manager 1 on 1/*.md"
      - "{{.Home}}/vault/Todo.md"
      - "{{.Home}}/vault/Objective Plan.md"
      - "{{.Home}}/vault/Status/{{.Date \"2006-07\"}}*.md"

  # Profile for on-call preparation
  on-call-prep:
    description: "Review my active on-call tasks and provide relevant documentation links for troubleshooting."
    prompt: "Review my active on-call tasks and provide relevant documentation links for troubleshooting."
    globs:
      - "{{.Home}}/vault/Todo.md"
      - "{{.Home}}/vault/KB/On-call.md"
      - "{{.Home}}/vault/KB/Monitoring.md"
      - "{{.Home}}/vault/KB/Message-Queue.md"
      - "{{.Home}}/vault/KB/Database.md"
      - "{{.Home}}/vault/Status/{{.Date \"2006-01\"}}*.md"
      - "{{.Home}}/vault/Status/{{.Date \"2006-02\"}}*.md"

  # Dynamic profile for feature deep dives
  feature-dive:
    description: "Context for a specific feature. Use with variables like feature_name."
    prompt: "Give me a complete overview of the {{.Vars.feature_name}} feature, including requirements, related tickets, and architectural decisions."
    globs:
      - "{{.Home}}/vault/Tickets/*{{.Vars.feature_name}}*.md"
      - "{{.Home}}/vault/KB/*{{.Vars.feature_name}}*.md"
      - "{{.Home}}/vault/AI revisions/*.md"
    variables:
      feature_name: "Search Feature"  # Default value

  # Profile for self-assessment
  self-assessment:
    description: "I'm writing my self-assessment. Based on my notes, what are 3 key strengths and 3 areas for improvement from the last quarter?"
    prompt: "I'm writing my self-assessment. Based on my notes, what are 3 key strengths and 3 areas for improvement from the last quarter?"
    globs:
      - "{{.Home}}/vault/Archive/Performance Review/Self-assessment.md"
      - "{{.Home}}/vault/Status/*.md"
      - "{{.Home}}/vault/Objective Plan.md"

  # Profile for knowledge base only
  kb-only:
    description: "Only includes knowledge base articles"
    prompt: "Based on my knowledge base, help me answer technical questions and provide detailed explanations."
    globs:
      - "{{.Home}}/vault/KB/**/*.md"

  # Profile for personal development
  personal-dev:
    description: "Review notes with personal development tags and suggest discussion topics."
    prompt: "Review notes from the past week with #development tags and suggest topics for discussion."
    globs:
      - "{{.Home}}/vault/**/*.md"
      - "!{{.Home}}/vault/Archive/**/*.md"
      # Note: This would benefit from date filtering, but glob patterns are limited
      # In a real implementation, you might want to use date-based filtering

  # Profile for recent work only (last 30 days)
  recent-work:
    description: "Focus on recent work and status updates"
    prompt: "Summarize my recent work and current priorities based on my latest notes."
    globs:
      - "{{.Home}}/vault/Status/{{.Date \"2006-01\"}}*.md"
      - "{{.Home}}/vault/Status/{{.Date \"2006-02\"}}*.md"
      - "{{.Home}}/vault/Todo.md"
      - "{{.Home}}/vault/Objective Plan.md"
```

## How to Use

1. Save this file in your Obsidian vault (e.g., `Configs/compiler-profiles.md`)
2. Update your main config to point to this file:
   ```yaml
   profiles_path: "{{.Home}}/vault/Configs/compiler-profiles.md"
   ```
3. Use profiles with the CLI:
   ```bash
   note-compiler --profile one-on-one-prep
   note-compiler --list-profiles
   ```

## Benefits of Markdown Format

- ✅ **Obsidian Integration**: Edit directly in Obsidian with syntax highlighting
- ✅ **Documentation**: Add context and comments around the configuration
- ✅ **Preview**: See rendered output in Obsidian's preview mode
- ✅ **Version Control**: Track changes with meaningful commit messages
- ✅ **Linking**: Reference other notes and create connections 