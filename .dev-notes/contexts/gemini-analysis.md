# Analysis of Profiles Feature
This is a great feature idea that would make your `note-compiler` tool much more powerful. Based on typical vault structures, here are some augmentations and concrete examples.

## Concrete Profile Examples

You can define profiles that map directly to your recurring workflows. These could be defined in a single YAML or Markdown file within your vault (e.g., `/Configs/compiler-profiles.yaml`).

| Profile Name | Included Globs / Key Files | Suggested Initial Prompt |
| :--- | :--- | :--- |
| **`on-call-prep`** | `Todo.md` (On-call backlog), `KB/On-call.md`, `KB/Monitoring.md`, `KB/Message-Queue.md`, `KB/Database.md`, recent `Status/*.md` | "Review my active on-call tasks and provide relevant documentation links for troubleshooting." |
| **`1-on-1-prep`** | `Meeting notes/Manager 1 on 1/*.md`, `Todo.md`, `Status/` from the last two weeks, `Objective Plan.md` | "Summarize my recent accomplishments, blockers, and talking points for my next 1-on-1 with my manager. Reference recent feedback." |
| **`feature-deep-dive`** | `Tickets/FEATURE-123*.md`, `KB/Search Feature*.md`, `KB/API Gateway.md`, `AI revisions/*.md` | "I'm working on the Search Feature. Summarize the key requirements, architectural decisions, and recent work on it." |
| **`self-assessment`** | `Archive/Performance Review/Self-assessment.md`, `Status/*.md`, `Objective Plan.md` | "I'm writing my self-assessment. Based on my notes, what are 3 key strengths and 3 areas for improvement from the last quarter?" |

## Implementation Details & Suggestions

### 1\. Profile Definition File

You could create a file like `vault/Configs/compiler-profiles.yaml` to define these contexts. This approach is powerful because the profiles themselves are part of your vault.

```yaml
# vault/Configs/compiler-profiles.yaml

# The 'base' profile is inherited by all other profiles.
# Perfect for your AI Instructions.
base:
  prompt: "Use these notes as context to answer questions and provide helpful responses."
  globs:
    - "{{.Home}}/vault/People/AI Instructions.md"
    - "!**/_resources/**" # Universal exclude

profiles:
  # Default profile if none is specified
  default:
    description: "The standard full-vault context, excluding archives."
    globs:
      - "{{.Home}}/vault/**/*.md"
      - "!{{.Home}}/vault/Archive/**/*.md"
      - "!{{.Home}}/vault/Tickets/xml/**"

  # Profile for your 1-on-1 meetings
  one-on-one-prep:
    description: "Prepares context for my 1-on-1 with my manager."
    prompt: "Summarize my recent work, blockers, and talking points for my upcoming 1-on-1. What questions should I ask my manager based on recent feedback?"
    globs:
      # Include specific, high-value directories
      - "{{.Home}}/vault/Meeting notes/Manager 1 on 1/*.md"
      - "{{.Home}}/vault/Todo.md"
      - "{{.Home}}/vault/Objective Plan.md"
      # Include only recent status updates
      - "{{.Home}}/vault/Status/{{.Date \"2006-07\"}}*.md"

  # A dynamic, parameterized profile for deep-diving on a feature
  feature-dive:
    description: "Context for a specific feature. Requires a 'feature_name' variable."
    prompt: "Give me a complete overview of the {{.Vars.feature_name}} feature, including requirements, related tickets, and architectural decisions."
    # Use variables to make the profile reusable
    globs:
      - "{{.Home}}/vault/Tickets/*{{.Vars.feature_name}}*.md"
      - "{{.Home}}/vault/KB/*{{.Vars.feature_name}}*.md"
```

### 2\. Augmenting Your Ideas

  * **Profile Composition/Inheritance:** As shown above, having a `base` profile that all other profiles inherit from is a clean way to ensure your core `AI Instructions.md` file is always included without repetition.
  * **Dynamic Prompts & Status:** The initial prompt can be dynamically generated. When you run `note-compiler --profile one-on-one-prep`, the first line sent to the AI could be the pre-defined prompt for that profile, followed by a status line.
      * *Example Output:*
        ```
        Summarize my recent work, blockers, and talking points for my upcoming 1-on-1. What questions should I ask my manager based on recent feedback?

        ---
        SYSTEM CONTEXT: Profile 'one-on-one-prep' | 42 files included, 1897 excluded.
        ---
        source: /Users/user/vault/Meeting notes/Manager 1 on 1/2025-06-26 Thu - Manager 1 on 1.md
        ---
        ...
        ```
  * **Templating in Prompts:** Your `note-compiler` already uses Go templates for file paths. You could extend this to allow simple date/time logic in your prompts. For personal development, you could define a prompt like: `"Review notes from {{.DateOffset \"-7d\" \"2006-01-02\"}} to today with the #development tag and suggest topics for discussion."` This would require the tool to resolve the template before passing it to the AI.

# On where to put the profiles

Here is an analysis of the two approaches for storing your profile configuration.

For most use cases, **storing the profile configuration inside your Obsidian vault is the superior option**. It aligns better with existing workflows and provides significant long-term benefits.

***

### Option A: Storing Profiles **Inside** the Vault

This involves creating a dedicated file like `/Configs/compiler-profiles.yaml` within your vault. The main `~/.note-compiler.yaml` would then just need a path pointing to this file.

**Pros:**
* **✅ Portability & Self-Contained:** Your entire context-generation system (notes + profiles) lives in one place. If you move your vault to a new computer or a different directory, the profiles come with it automatically.
* **✅ Version Control:** If you use Git with your vault, your profiles history will be tracked. You can see when you created a profile, what changes you made, and revert if needed. This is a major advantage.
* **✅ Ease of Editing:** You can edit your profiles directly within Obsidian, your primary tool, without needing to open a separate file in your home directory.
* **✅ AI Context-Awareness:** When you compile notes for an AI prompt, you can easily include the profile definition file itself, allowing the AI to understand the *intent* behind the context it's receiving.

**Cons:**
* **❌ Initial Setup:** It requires one extra line in your `~/.note-compiler.yaml` to specify the path to the profiles file (e.g., `profiles_path: "{{.Home}}/vault/Configs/compiler-profiles.yaml"`).

***

### Option B: Storing Profiles in the Home Directory (`~/.note-compiler.yaml`)

This involves adding a `profiles:` section directly into your existing config file.

**Pros:**
* **✅ Simple Discovery:** The tool already knows where to find the file. This is the most straightforward implementation with no extra configuration needed.
* **✅ Separation of Concerns:** It maintains a clean, traditional separation between your configuration and your content.

**Cons:**
* **❌ Not Portable:** If you move your vault or set it up on a new machine, you have to remember to copy your `~/.note-compiler.yaml` file. The configuration is not tied to the content.
* **❌ No Version Control:** Your home directory is not in Git, so any changes to your profiles—adding new ones, tweaking globs—will not be versioned. This is a significant loss given typical setups.
* **❌ Harder to Share:** It's more difficult to share a specific profile configuration with a colleague or your future self on another machine.

## Recommendation

Go with **Option A (inside the vault)**. The small, one-time cost of adding a `profiles_path` to your main config file is heavily outweighed by the benefits of portability and integrated version control.
