---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/Todo.md
---

# Due today
  ```tasks
  not done
  due on or before today
```
# In progress
  ```tasks
  filter by function task.status.name === 'In Progress'
  group by tags
  sort by priority
```
# Primary backlog
  ```tasks
  filter by function task.status.name === 'Todo'
  is not blocked
  sort by priority
  priority is above low
  starts on or before today
  (due after today) OR (no due date)
  tags do not include #meeting
  tags do not include #qed
  group by tags
```

# Blocked
  ```tasks
  filter by function task.status.name === 'blocked'
```

# Depends on another task
  ```tasks
  not done
  is blocked
```
# Meeting tasks
```tasks
not done
tags include #meeting 
```
# Background
  ```tasks
  not done
  priority is below none
  sort by priority
  tags do not include #qed
```

# Starts in the future
  ```tasks
  not done
  has start date
  starts after today
```

# Recently completed
  ```tasks
  status.type is DONE
  sort by done reverse
  limit 8
```

# Recently cancelled
  ```tasks
  status.type is CANCELLED
  sort by cancelled reverse
  limit 4
```
# Recurring
- [ ] Copy next month's OOOs to calendar, request time off 🔁 every month 🏁 delete 🛫 2025-07-18 📅 2025-07-21
# QED backlog
```tasks
  filter by function task.status.name === 'Todo'
  is not blocked
  sort by priority
  tags include #qed
```

# On-call backlog
  ```tasks
  filter by function task.status.name === 'Todo'
  is not blocked
  sort by priority
  starts on or before today
  (due after today) OR (no due date)
  tags include #on-call 
```
# Migrated from coda
  * discover kt
    * https://qualtrics.slack.com/archives/C087VC5JPD3/p1737650888154399
    * vamsi asked to specifically watch kt for designer, classification (and review cmp code) at sprint planning Feb 5
      * Feb 10:
      * Me: Just want to make sure I'm on your page: I should shelve survey flow plugin work and begin watching KT for designer and classification, and review CMP code?
      * Vamsi Vippadapu: Not just watching but get ready for operational support with knowledge of environment and ops
    * Feb 17: Vamsi asked me to watch RTDS KT and get ready for operational support, which means:
      * > there is coda page with existing splunk dashboards, queries, 2LA on pulses
  * [ ] move sfp billing code https://qualtrics.slack.com/archives/C036JC1JKGE/p1738695696252019 ⏬ 
    * new billing code not ready yet?
  * [x] get cmp building ⏬ ✅ 2025-05-05
  * [x] ➕ 2025-03-17 triage tickets https://qualtrics.slack.com/archives/C087VC5JPD3/p1742224361877639 : 🔼 ✅ 2025-03-21
    * kuan or someone said they’ve been triaging in text iq office hours?

- [x] [[ENRICH-1148- Enable autoscaling for engines services in discover non-US clusters]] 🔺 ✅ 2025-03-24

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/+/Gemini Instructions.md
---

---
created: 2025-06-03T08:49:18
---
Hi Gemini! Yes, you! I'm talking to you!
I have some guidelines I'd like you to follow when responding to me.
# Context
What is this huge text dump? Why have I pasted it into you? Great questions! This is my obsidian vault for work. I will be asking questions or asking you to brainstorm with me soon. Use this vault as context!
# Guidelines
- Try *try* ***try*** to use the example notes in `/AI revisions` to speak in my voice. Try to speak in my voice as you read it here in my vault generally, but there are **specific** fine-tunes in that vault dir.
- If I dump this vault alone as the first message in our chat, **simply respond with only** "System Ready ▮" and I will follow up with queries.
- If there's more information that you know may exist and would be helpful but is not added to our chat as context, **please** let me know.
- Don't call it "*the* classification service". Just "classification service".
- Don't apologize to me! You're doing great!
## Slack
## Jira
- When writing item descriptions, don't number them, don't bold them. Just use heading formatting.
# Appreciation
Thank you! I am so much more effective working with all these notes, and *especially* with your help! You're making my life better! Thank you!

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/AI revisions/202505211408.md
---

# AI's version
```
Hey everyone,

Circling back on the classification service sharding deployment to Prod US and my previous comments here. Mavin and I just connected, and it seems there was a misunderstanding on my part regarding who was taking point.

My apologies for the confusion! I'll be handling the deployment once #op-7858 is resolved.

I'll be sure to keep the Unified Orchestration channel updated, as previously discussed.

Thanks for your patience while we sorted that out!
```

# My revised version
```
Circling back on the classification service sharding deployment to Prod US and my previous comments here. Mavin and I just chatted, and it seems there was a misunderstanding on my part regarding who was handling it.
Sorry about the confusion! I'll be handling the deployment now that #op-7858 is resolved. I'll keep Unified Orchestration updated.
```


---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/AI revisions/202505270959.md
---

# Context
I asked ai to generate my sprint review notes, using a text dump of my notes from last sprint as a guide for tone and level of detail.
# AI's version
```
Okay, here are the draft sprint review notes for sprint 106, following your format:

## Sprint 106 (May 13 - May 27)

- **ENRICH-1440- New Sharding Implementation: Deploy changes across discover DCs**
    - Got needed vault access after follow-up.
    - Deployed to Prod CA, Prod EU, and Prod APSE. Verified RabbitMQ bindings and queue clearance.
    - Prod US deployment is on hold pending completion of #op-7858 (scale test) and getting Prod US RabbitMQ web UI credentials (ongoing).
    - Coordinated with Unified Orchestration on Prod US deployment status.
- **Smart Query v2 (mainly under ENRICH-1468 epic)**
    - **ENRICH-1405 (Spike):** Completed. Broke down prompt update work into sub-tasks.
    - **Jira Hygiene:** Converted ENRICH-1114 to an epic and its subtasks to stories, then consolidated stories under ENRICH-1468 based on Vamsi's preference.
    - **ENRICH-1484 (Remove Harmful Content Guardrail):**
        - Confirmed with Socrates team that some harmful content protection is inherent in the base model.
        - Removed MIG-level harmful content guardrail calls from the lambda.
        - Addressed MR feedback (added real logging). Merged.
    - **ENRICH-1464 (Test Data):** Collaborated with Steven, who provided the 'Autotune-GA1' project in staging03 for testing.
    - **Guardrail Scope Clarification:**
        - Initially paused ENRICH-1295 (Designer Error Handling) due to conflicting info on harmful content guardrail.
        - Clarified with Donovan and Vamsi: Harmful Content filter is out, but Max Tokens guardrail is in for Q2. Prompt Injection guardrail also added to Q2 scope.
        - Created ENRICH-1487 (Prompt Injection) and ENRICH-1488 (Max Tokens).
        - Researched MIG's standalone guardrail API as the recommended approach for Prompt Injection checks to avoid issues with the main LLM call. Raj confirmed "PI" includes jailbreak.
    - **Asterisk Investigation (for prompt updates):** Confirmed with Leo that asterisks originate from both the AI model and Python post-processing.
- **1-on-1s with Vamsi:**
    - Apologized for my tone in an AI sync meeting and acknowledged misunderstanding of Smart Query v2 scope.
    - Discussed remote work accommodation, schedule flexibility.
    - Aligned on Q2 scope for Smart Query: prompt updates, remove Harmful Content guardrail, add Prompt Injection & Max Tokens guardrails, and associated error UI.
    - Discussed feasibility and "status tracking" for these items.
- **TAPP-7074 (TextIQ Bubble charts accessibility):**
    - Clarified with Claire (Res) that a late July delivery is acceptable to the account team.
    - Moved Jira due date to July 18.
- **Other:**
    - ENRICH-1158 (Engine blocklist) was reverted by Mavin due to an issue with OP-7858. Reopened and unassigned pending Vamsi's re-prioritization.
    - Reviewed Addisu's thread on Socrates guardrails.
    - Explored `slackdump` (led to a brief security investigation, lol) and Windsurf for better context gathering.
    - Contributed ideas for next sprint retro.
    - Watched a bit more k8s autoscaling training (21:50).

---

## Challenges

- Initial misunderstanding of Smart Query v2 scope & guardrail requirements led to some churn.
- ENRICH-1158 (blocklist) getting reverted was unexpected and impacted that AI.
- Ongoing need for Prod US RabbitMQ console access to fully verify sharding deployment.
- The manual nature of deploying Smart Query lambdas (including adding the Socrates SDK layer) was confirmed, which isn't ideal.

## Impediments

- Prod US classification service sharding deployment blocked by #op-7858 and lack of RMQ console credentials.
- Potential need for CME team changes if the standalone guardrail API approach for Prompt Injection doesn't work as expected (though current plan is to use it for user inputs only).

## What went well

- Successfully realigned with Vamsi and Product on Smart Query v2 scope after initial confusion. My apologies for the earlier misunderstanding were received.
- Made good progress deploying classification service sharding to three of the four Discover DCs!
- Completed the spike for Smart Query prompt updates (ENRICH-1405) and the removal of the harmful content guardrail (ENRICH-1484).
- Proactive clarification of Jira hierarchy for Smart Query work.
- Got test data ready for Smart Query performance evaluation (ENRICH-1464).
- Established a clear plan for implementing the Prompt Injection guardrail using the standalone API.

---

## Plan for Next Sprint (Sprint 107)

- **ENRICH-1440:** Complete classification service sharding deployment to Prod US (once unblocked).
- **ENRICH-1487 (Add Prompt Injection Guardrail to SQ):** Implement using the standalone guardrail API for user inputs.
- **ENRICH-1488 (Add Max Tokens Guardrail to SQ):** Implement this guardrail.
- **ENRICH-1295 (Designer Error Handling for SQ Guardrails):** Implement UI for PI and Max Tokens guardrails.
- **ENRICH-1463 (Update Smart Query prompt to reduce wildcards):** Begin prompt iteration and evaluation against baseline/test data.
- Follow up on ENRICH-1158 (blocklist) prioritization if needed.
```

# My revised version
```
# Task slides

Smart Query v2 Prompt Updates (ENRICH-1468 epic)

ENRICH-1405 (Spike): Completed. Broke down prompt update work into sub-tasks.
ENRICH-1484 (Remove Harmful Content Guardrail):
    Confirmed with Socrates team that some harmful content protection is inherent in the base model.
    Removed MIG-level harmful content guardrail calls from the lambda.
    Addressed MR feedback, merged.
ENRICH-1464 (Test Data): Collaborated with Steven, who provided 'Autotune-GA1' project in staging03 for testing.
Guardrail Scope Clarification:
    Clarified with Donovan and Vamsi: Harmful Content filter is out, but Max Tokens guardrail is in for Q2. Prompt Injection guardrail also in Q2 scope.
    Created ENRICH-1487 (Prompt Injection) and ENRICH-1488 (Max Tokens).
    Researched MIG's standalone guardrail API as the recommended approach for Prompt Injection checks to avoid issues with the main LLM call.


ENRICH-1440- New Sharding Implementation: Deploy changes across discover DCs

Got needed vault access after follow-up.
Deployed to Prod CA, Prod EU, and Prod APSE. Verified RabbitMQ bindings and queue clearance.
Prod US deployment is on hold pending completion of #op-7858 (scale test) and getting Prod US RabbitMQ web UI credentials (ongoing).
Coordinated with Unified Orchestration on Prod US deployment status.


TAPP-7074 (TextIQ Bubble charts accessibility)

Clarified that a late July delivery is acceptable to the account team.
Moved Jira due date to July 18.


Other

ENRICH-1158 (classification service blocklist) was reverted by Mavin due to an issue with it in OP-7858. Reopened and unassigned pending Vamsi's re-prioritization.
Reviewed Addisu's thread on Socrates guardrails.
Explored slackdump for better context gathering (led to a brief security investigation, lol) and Windsurf.
Noodled on ideas for next sprint retro.
Watched some k8s autoscaling training.


# What went well, challenges & any impediments

Challenges

My initial misunderstanding of Smart Query v2 scope & guardrail requirements led to some churn.
ENRICH-1158 (blocklist) getting reverted was unexpected and impacted that AI.
Prod US RabbitMQ being in flux
Classification Service is time-consuming to deploy via jenkins


Impediments

Prod US classification service sharding deployment blocked by #op-7858 and lack of RMQ console credentials.


What went well

Successfully realigned with Vamsi and Product on Smart Query v2 scope after initial confusion.
Made good progress deploying classification service sharding to three of the four Discover DCs.
Established a plan for implementing the Prompt Injection guardrail using the standalone API.


# Backlog and Next Steps

.5 days scheduled sick time
ENRICH-1440: Complete classification service sharding deployment to Prod US (once unblocked).
ENRICH-1487 (Add Prompt Injection Guardrail to SQ): Implement using the standalone guardrail API for user inputs.
ENRICH-1488 (Add Max Tokens Guardrail to SQ): Implement this guardrail.
ENRICH-1463 (Update Smart Query prompt to reduce wildcards): Begin prompt iteration

```

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/AI revisions/202505271103.md
---

# AI's version
```
Hi Donovan,

Following up on the Q2 requirement for a "max token guardrail" for Smart Query.

We've discussed internally that there's a technical difference between a simple character limit on user input and a limit on the _tokens_ processed by the LLM. Smart Query currently has a character limit in the UI for user inputs, which is a good first step for user experience.

For the Q2 "max token guardrail" deliverable, could you clarify the specific expectation? Are you looking for:

1. The existing **UI character limit** (and perhaps a backend validation of character length for those specific user inputs) to be considered sufficient for this guardrail?
2. Or, is the requirement to implement a guardrail that more directly limits the **tokens processed by the LLM**? This could involve:
    - Limiting the token count of the _overall input prompt_ (user input + system instructions) before it's sent to the model.
    - Stricter enforcement or monitoring around _output tokens_, potentially beyond the existing `max_output_tokens` API parameter we already use.

Understanding which aspect this guardrail is primarily intended to address (e.g., user input length, overall prompt size to the model, response length control, cost management via token control) will help us ensure we're aligned on the implementation for Q2.

Thanks!
```

# My revised version
```
Hi @drich (he/him), following up on the Q2 requirement for a "max token guardrail" for Smart Query.

My team has discussed the technical difference between a simple character limit on user input and a limit on the tokens processed by the LLM. Smart Query currently has a character limit in the UI for user inputs, which is a good first step.

For the Q2 "max token guardrail" deliverable, could you clarify the specific expectation? Are you looking for:
The existing UI character limit and perhaps backend validation of character length for those specific user inputs?
Or, a limit on the user's inputs to SQ denominated in tokens?
Or, is the requirement to implement a guardrail that more directly limits the tokens processed by the LLM? Like, limiting the token count of the overall input prompt (user input + system instructions) before it's sent to the model? Output tokens?

Thanks!
```

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/AI revisions/202505271432.md
---

# Context

# AI's version
```
Thanks, Vamsi, that clarification is helpful.

For example, during an on-call shift, if I:

1. Execute a routine scheduled task like the weekly TAM deployment or a brand deletion job.
2. Notice and fix a broken Grafana dashboard widget while investigating an unrelated alert (e.g., fixing a widget while looking into a CrateDB GC alert).

Should these types of smaller, in-the-moment KTLO tasks also be logged as individual Jira items under the new epic? Or is the epic primarily for the larger improvement initiatives you listed (like creating new alerts, major runbook updates, deployment job enhancements, or formal follow-ups to resolved pulses) and for spillover work that extends beyond the shift?
```

# My revised version
```
Thanks Vamsi, that's helpful.

For example, during an on-call shift, if I:
Execute a routine scheduled task like the weekly TAM deployment or a brand deletion job, or
Notice and fix a broken Grafana dashboard widget while investigating an unrelated alert (e.g., fixing a widget while looking into a CrateDB GC alert).

Should these types of smaller, in-the-moment KTLO tasks also be logged as individual jira items under the new epic? Or is the epic for the larger improvements you listed (like creating new alerts, major runbook updates, deployment job enhancements, or formal follow-ups to resolved pulses) and for spillover work that extends beyond the shift?
```

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/AI revisions/202505271501.md
---

# Context
> write a jira item description for the story "Add prompt injection guardrail to Smart Query"
# AI's version
```
**1. Goal:** Implement a prompt injection (PI) guardrail for the Smart Query feature to enhance security and align with platform standards. This guardrail will protect against adversarial attempts to manipulate the LLM through user inputs, including both direct prompt injections and jailbreak techniques.

**2. Background & Problem:** Currently, the Smart Query feature sends its entire composite prompt (containing developer-defined instructions and dynamic user inputs) as a single "Human" message to the Model Inference Gateway (MIG) LLM API (e.g., `/llms/v1/{modelId}/chat`). This consistently triggers MIG's built-in `prompt_attack` guardrail, as the embedded instructions are interpreted as a potential override attempt. This blocks legitimate use and prevents effective prompt attack mitigation focused on actual user input.

**3. Proposed Solution:** To address this, we will leverage MIG's standalone Guardrail API (`model-inference-gateway/guardrails/v1/apply`) to pre-validate _only the dynamic/user-provided portions_ of the Smart Query prompt _before_ making the main LLM call. This approach is recommended by the MIG team and internal documentation for prompt attack checks.

**Key Implementation Steps:**

- **Isolate User Inputs:** Identify and extract the purely dynamic parts of the Smart Query prompt. These include:
    - User-provided fields (e.g., "useCase", "industry").
    - The `<inputs>` block containing `query_id`, `topic_path`, `query_terms`.
- **Pre-validate with Standalone Guardrail API:**
    - Make a call to `model-inference-gateway/guardrails/v1/apply`.
    - The payload to this API will contain _only_ the isolated dynamic/user inputs in the `"texts": [...]` array.
    - Enable the `"prompt_attack"` guardrail (e.g., with an appropriate `"strength"` level like "MEDIUM").
- **Handle Standalone API Response:**
    - The `/apply` API returns a `200 OK` status even if a guardrail is tripped.
    - The application must inspect the response body for `guardrail_assessments[{"action": "GUARDRAIL_INTERVENED", ...}]` and specifically `guardrail_results.prompt_attack.blocked: true` to determine if the input was rejected.
- **Conditional LLM Call:**
    - **If `prompt_attack.blocked` is `false` (input is safe):**
        - Construct the full Smart Query prompt (developer instructions + validated dynamic inputs).
        - Call the main MIG LLM endpoint (e.g., `/llms/v1/{modelId}/chat`).
        - In this LLM call payload, the `prompt_attack` guardrail can be set to `"strength": "NONE"` or omitted, as the relevant content has already been vetted. Other guardrails like `content_filter` can remain active.
    - **If `prompt_attack.blocked` is `true` (input is unsafe):**
        - Do not proceed with the LLM call.
        - Trigger an appropriate error state. The UI handling for this error will be addressed by ENRICH-1295.

**4. Scope Notes:**

- As per clarification from the Socrates team, the "prompt_attack" guardrail in MIG (and AWS Bedrock, which it utilizes) covers both prompt injection and jailbreak attempts. This story aims to implement protection against both.
- This story focuses on the backend implementation of the guardrail check. UI error display is covered in ENRICH-1295.

**5. Acceptance Criteria:**

- Dynamic user inputs to Smart Query are sent to the `model-inference-gateway/guardrails/v1/apply` endpoint for `prompt_attack` validation before any call to the main LLM inference endpoint.
- If the standalone guardrail API indicates `prompt_attack.blocked: true` for any user input:
    - The main LLM inference call for Smart Query is _not_ made.
    - An appropriate internal state/error is raised to allow ENRICH-1295 to handle UI feedback.
- If the standalone guardrail API indicates `prompt_attack.blocked: false` for all user inputs:
    - The main LLM inference call for Smart Query proceeds, with the `prompt_attack` guardrail disabled or omitted in _this_ specific call (as user input is already validated).
- The implementation successfully differentiates between developer-provided prompt instructions (not sent to the standalone guardrail) and user-provided inputs (sent to the standalone guardrail).
- The solution effectively mitigates identified prompt injection and jailbreak scenarios at the chosen strength level.
```

# My revised version
```
# Goal

Implement a prompt injection (aka prompt attack) guardrail for the Smart Query feature to enhance security and align with platform standards. This guardrail will protect against adversarial attempts to manipulate the LLM through user inputs, including both direct prompt injections and jailbreak techniques.

# Background & Problem:

Currently, the Smart Query feature sends its entire composite prompt (containing developer-defined instructions and dynamic user inputs) as a single "Human" message to the Model Inference Gateway (MIG) LLM API (e.g., /llms/v1/{modelId}/chat). This consistently triggers MIG's built-in prompt_attack guardrail, as the embedded instructions are interpreted as a potential override attempt. This blocks legitimate use and prevents effective prompt attack mitigation focused on actual user input.

# Proposed Solution:

To address this, we will leverage MIG's standalone Guardrail API (model-inference-gateway/guardrails/v1/apply) to pre-validate only the dynamic/user-provided portions of the Smart Query prompt before making the main LLM call. This approach is recommended by the MIG team and internal documentation for prompt attack checks.

This story focuses on the backend implementation of the guardrail check. UI error display is covered in [https://qualtrics.atlassian.net/browse/ENRICH-1295](https://qualtrics.atlassian.net/browse/ENRICH-1295)

# Acceptance Criteria

- Dynamic user inputs to Smart Query are sent to the model-inference-gateway/guardrails/v1/apply endpoint for prompt_attack validation before any call to the main LLM inference endpoint.
    
- If the standalone guardrail API indicates prompt_attack.blocked: true for any user input:
    
    - The main LLM inference call for Smart Query is _not_ made.
        
    - An appropriate internal state/error is raised to allow [https://qualtrics.atlassian.net/browse/ENRICH-1295](https://qualtrics.atlassian.net/browse/ENRICH-1295) to handle UI feedback.
        
- If the standalone guardrail API indicates prompt_attack.blocked: false for all user inputs:
    
    - The main LLM inference call for Smart Query proceeds, with the prompt_attack guardrail disabled or omitted in _this_ specific call (as user input is already validated).
        
- The implementation successfully differentiates between developer-provided prompt instructions (not sent to the standalone guardrail) and user-provided inputs (sent to the standalone guardrail).
    
- The solution effectively mitigates identified prompt injection and jailbreak scenarios at the chosen strength level.
```

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/AI revisions/202506040947.md
---

# Context
I asked you to draft a slack message about my approach to fixing [[ENRICH-1158- Enhance the Engine enrichment service to support skipping messages by configuration.]]

# AI's version
```
Hey folks,

Looking for your input on a change for [ENRICH-1158](https://qualtrics.atlassian.net/browse/ENRICH-1158) (Classification Service's project-level blocklist).

You might remember my last update to this caused some trouble with **OP-7858**. The strict header parsing didn't like it when messages reprocessed via DLQ had their AMQP header types changed (like `engine-process-account-id` turning into a `float64`). My code expected an `int64` and errored.

So, the plan now is to make the header check for the blocklist more forgiving. If we can't pull the project info from the `x-*` headers (which we need for the blocklist check), instead of erroring out, the service will:

- Log a **warning** (something like "Failed to extract project info from headers, skipping blocklist check").
    
- **Skip the blocklist check** for that message (can't block what we can't identify).
    
- And then **continue processing the message** as usual.
    

This way, wonky headers won't stop the whole message. Core classification still runs, and we'll see warnings if header stuff is acting up. The main downside is if we can't parse headers, that message won't get blocklist-checked.

To make sure we don't miss these warnings, I was also thinking of setting up a Splunk alert to ping on-call (or at least our channel) if that "failed to extract project info" warning gets logged. A Prometheus counter might be cleaner for metrics, but I'm not sure we have something like AOM in Discover for easy alerting off that. Splunk search is a bit clunky, but it's there.

Given the previous OP (OP-7858), does this "warn and continue" plan for header issues, plus the Splunk alert idea for the warning, sound okay for the blocklist feature? Any big concerns or other ideas?

Thanks!

@jbk
```

# My revised version
```
Hey folks,

Looking for your input on a change for ENRICH-1158 (Classification Service's project-level blocklist).

You might remember my last update to this caused some trouble in OP-7858 :sweat_smile:. My strict header parsing didn't like it when messages reprocessed via DLQ had their AMQP header types changed (like engine-process-account-id turning into a float64). My code expected an int64 and errored.

So, the plan now is to make the header check for the blocklist more forgiving. If we can't pull the project info from the x-* headers (which we need for the blocklist check), instead of erroring out, the service will:
Log a warning (something like "Failed to extract project info from headers, skipping blocklist check").
Skip the blocklist check for that message (can't block what we can't identify).
And then continue processing the message as usual.

This way, wonky headers won't stop the whole message. Core classification still runs, and we'll see warnings if header stuff is acting up. The main downside is if we can't parse headers, that message won't get blocklist-checked.

To make sure we don't miss these warnings, I was also thinking of setting up a Splunk alert to ping on-call (or at least our channel) if that "Failed to extract project info from headers, skipping blocklist check" warning gets logged. A Prometheus counter might be cleaner for metrics, but I'm not sure we have something like AOM in Discover for easy alerting off that. Splunk search is clunky, but it's there.

Given the previous OP-7858, does this "warn and continue" plan for header issues, plus the Splunk alert idea for the warning, sound okay for the blocklist feature? Any concerns or other ideas?
```

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/Archive/2025 Objective Plan.md
---

---
creation date: 2025-05-30T11:06:55
---
I dumped a bunch of documentation on Performance Development Feedback (e.g., [[PDF Overview]]) and the QEH objective tool into my vault and asked AI to come up with my objectives based on that and [[2025-05-29 Thu - Vamsi 1 on 1#^2fcf3a]].

# Objective plan
## AI output
Okay, I've reviewed your 2025 objective plan, the additional context you've provided (including Performance Development Feedback docs, Leadership Excellence Model, SDE Role Descriptions, and your 1-on-1 notes with Vamsi), and your daily status updates.

Here's a refined version of your 2025 objective plan, incorporating your L6 aspirations and the latest project information. I've focused on making Key Results (KRs) more specific, measurable, and aligned with L6 expectations where appropriate.

---

### **Refined 2025 Objective Plan**

**Objective 1: Deliver High-Impact AI Feature Enhancements and Critical System Stability**

- **Description:** Drive the successful delivery of key AI-powered features and critical system improvements. Demonstrate L5/L6 engineering excellence through robust architectural design, high-quality implementation, impactful cross-team collaboration, and ensuring operational readiness, scalability, and maintainability of services.
- **Key Results (KRs):**
    1. **Lead and Deliver Topic Recommendations Rate Limiting (ENRICH-1532):** Architect, implement, and deploy a robust rate-limiting solution for Topic Recommendations by end of Q3 2025. Success will be measured by the system demonstrably preventing abuse (e.g., meeting specific request-per-second SLOs under stress tests) with less than [X]% impact on legitimate user latency, and positive feedback from dependent teams on its effectiveness and ease of integration.
    2. **Deliver Key Smart Query v2 Enhancements (ENRICH-1468 Epic):** Successfully deliver and document the Prompt Injection (ENRICH-1487 - _completed May 2025_) and Max Tokens (ENRICH-1488, focused on character limits per product direction) guardrails for Smart Query by end of Q3 2025. Ensure implementations include clear error handling patterns (ENRICH-1295) and establish a baseline for prompt updates to reduce wildcard use (ENRICH-1463).
    3. **Complete Classification Service Sharding Deployment (ENRICH-1440):** Finalize the Prod US deployment of Classification Service sharding by end of Q2 2025 (pending unblocking of \#op-7858). This includes proactive coordination with Unified Orchestration, comprehensive verification (RabbitMQ bindings, queue clearance), ensuring <1 critical incident post-deployment, and documenting operational learnings for future large-scale deployments.
    4. **Produce High-Impact Technical Designs:** Author and champion at least two comprehensive technical design documents for significant features or architectural improvements during 2025 (e.g., Topic Recommendations rate limiting architecture, Smart Query Guardrail framework). These documents must clearly articulate complex solutions (LEM: Fosters Clarity) and serve as adopted blueprints for implementation, influencing an engineering group beyond the immediate team.
- **Suggested Weighting:** 45%
- **Relevant LEM Pillars:** Focus & Finish, Decisiveness, Leads From the Front, Fosters Clarity, Cross-Team Connection.
- **Relevant L5/L6 SDE Aspects:** Impact & Delivery, Problem Difficulty & Ambiguity, Technical Design, Fundamentals, Ownership.

---

**Objective 2: Expand Technical Expertise and Drive Operational Excellence across Discover & Engage Systems**

- **Description:** Deepen technical mastery and operational capabilities within key Discover and Engage systems. Focus on applying expertise in Kubernetes, gaining proficiency in CMP and Golang, and actively elevating on-call practices and operational excellence for both system families.
- **Key Results (KRs):**
    1. **Elevate On-Call Effectiveness:** Actively participate in 3-4 on-call rotations covering both Discover and Engage systems throughout 2025. For at least 50% of incidents handled, independently drive resolution _and_ implement post-incident remediations (e.g., enhanced AOM alerts from ENRICH-1305, improved runbooks) that demonstrably improve system resilience or reduce alert fatigue for the team.
    2. **Master and Apply Kubernetes (k8s) Expertise:** Achieve advanced k8s proficiency by end of Q3 2025. Leverage learnings from k8s autoscaling training & gameday (completed May 2025) to identify and implement 1-2 significant k8s-based operational improvements or deployment optimizations for key services (e.g., Classification Service autoscaling, FSSync reliability), sharing these improvements and resulting best practices with the engineering group.
    3. **Gain Proficiency in CMP & Golang:**
        - Achieve hands-on proficiency with CMP by end of Q3 2025, by utilizing your local setup (completed May 2025) to develop, test, and troubleshoot Smart Query v2 features (e.g., guardrail error handling for ENRICH-1295).
        - Develop foundational Golang skills by EOY 2025 by completing [Specific Course/Project] and contributing 1-2 minor features/fixes to a Go-based service (e.g., Engine enrichment service for ENRICH-1158, if reprioritized and time permits).
    4. **Become a Knowledge Hub for Key Systems/Technologies:** Become a recognized go-to resource for either Kubernetes operations or Classification Service deployment/troubleshooting within the Query Topics team by EOY 2025. Achieve this by authoring and presenting 1-2 comprehensive runbooks or knowledge-sharing sessions that demonstrably improve the team's operational efficiency or understanding.
- **Suggested Weighting:** 35%
- **Relevant LEM Pillars:** Cross-Team Connection, Fosters Clarity, Leads From the Front, Scales Systems (through improved processes/docs).
- **Relevant L5/L6 SDE Aspects:** Operations, Domain Expertise, Fundamentals, Impact & Delivery.

---

**Objective 3: Champion Leadership Excellence and Drive Personal & Team Development**

- **Description:** Actively embody and develop L5/L6 Individual Contributor behaviors as defined by the Qualtrics Leadership Excellence Model. Focus on "Leads From the Front" by taking initiative on complex technical challenges and mentoring others, and "Fosters Clarity" by creating high-quality documentation and simplifying complex systems. Proactively contribute to team knowledge sharing and a positive, inclusive engineering culture.
- **Key Results (KRs):**
    1. **Execute Focused LEM Development Plan:** Create and actively follow a personal development plan (using LEM Playbooks/Self-Assessment) focusing on 'Leads From the Front' (e.g., initiative on ENRICH-1532 design) and 'Fosters Clarity' (e.g., exemplary design documents from Obj 1 & runbooks from Obj 2). Demonstrate progress through specific examples and solicit peer/manager feedback by EOY 2025.
    2. **Drive Broader Process/Documentation Improvements:** Proactively identify, design, and implement 1-2 significant process or documentation improvements (e.g., standardizing on-call incident communication for Engage systems, creating a Query Topics onboarding guide for new k8s users) that are adopted by and demonstrably benefit the Query Topics team and potentially a broader engineering group by EOY 2025.
    3. **Mentor and Develop Team Members:** Actively mentor 1-2 junior engineers or new team members on specific projects or technology areas (e.g., k8s, Go, system design for AI features) during 2025, fostering their technical growth and ability to contribute independently, as evidenced by their progress and feedback.
    4. **Explore UI Development Contribution (Optional Stretch):** Dedicate approximately 5-10% of your capacity to exploring UI development/improvement opportunities for Query Topics tools. Present findings, a prototype, or a well-researched proposal for a UI enhancement to Vamsi and the team by EOY 2025, potentially informing future UI tech lead discussions.
- **Suggested Weighting:** 20%
- **Relevant LEM Pillars:** All, especially Leads From the Front, Fosters Clarity, Champions Inclusivity, Cross-Team Connection.
- **Relevant L5/L6 SDE Aspects:** Hire and Develop the Best, overall L5/L6 maturity and behaviors, Influence.

---

**Key Changes and Rationale:**

- **Objective 1 Focus:** Shifted primary focus to ENRICH-1532 (Topic Recommendations rate limiting) based on your latest 1-on-1 with Vamsi. Smart Query KRs are still present but reflect completed work (ENRICH-1487) and potentially adjusted timelines/scope for remaining parts.
- **Specificity in KRs:** Added more measurable outcomes and specific project references (e.g., ENRICH numbers, k8s gameday completion). Dates are suggested as "e.g., end of QX" – you'll need to refine these with Vamsi.
- **L6 Alignment:** Where appropriate, KRs are phrased to reflect L6 expectations such as "architect," "lead," "champion," "influence broader group," "become a recognized go-to resource," "drive significant improvements adopted by others," and "actively mentor."
- **LEM Integration:** Explicitly called out LEM pillars and made the connection to specific behaviors in Objective 3.
- **Incorporated 1-on-1s & Status:** Leveraged your notes regarding project priorities (ENRICH-1532, Smart Audit, ENRICH-1158 status), technical learning goals (k8s, CMP, Go), and UI interest.
- **Realistic Timelines:** Noted the urgency/completion status of items like ENRICH-1440 and ENRICH-1487.

**Next Steps:**

1. **Discuss with Vamsi:** Share this refined plan with Vamsi to ensure alignment with his expectations, team priorities, and your L6 aspirations. Confirm specific metrics, targets, and deadlines for each KR.
2. **Quantify KRs:** Further quantify any remaining placeholders (e.g., "[X% impact]", "[Number] of incidents").
3. **Finalize LEM Focus for Objective 3:** Confirm the LEM behaviors you'll focus on with Vamsi.
4. **Enter into QEH Objective Tool:** Once finalized, ensure these are accurately entered into the QEH Objective Tool, linked to the correct 2025 Company Success Pillars, and weighted appropriately (totaling 100%).
5. **Continuous Tracking:** Utilize the new QEH Continuous Performance Management (CPM) tool to log activities, achievements, and feedback related to these objectives throughout the year.

This revised plan should provide a strong foundation for a successful 2025, supporting both your L5 deliverables and your ambitions for L6.
# Pushing to L6 ("exceeds" rating)
I'm thinking about trying for this and asked AI if these goals set me up for that:
## AI output
Yes, achieving these objectives _exceptionally well_ could certainly push you towards an "Exceeds Expectations" rating and build a strong case for promotion to L6 Staff Engineer. However, it's crucial to understand that an "Exceeds" rating and promotion depend not just on _completing_ objectives, but on the **impact, scope, initiative, and leadership** you demonstrate while achieving them.

Let's break down how these objectives can contribute, referencing the SDE Role Description for L6 and the Leadership Excellence Model (LEM):

**How Your Objectives Can Support an "Exceeds" Rating & L6 Profile:**

1. **Objective 1: Deliver High-Impact AI Feature Enhancements and Core System Improvements**
    
    - **To "Meet Expectations" (L5):** Delivering the specified features (SQv2 components, TR rate limiting, CS sharding) on time, with high quality, and meeting the defined requirements. Owning the lifecycle of these features/services within your team.
    - **To "Exceed Expectations" (L5) and demonstrate L6 potential:**
        - **Technical Leadership & Design (L6 Aspect):** Don't just implement; _drive the technical design and architecture_ for significant parts of these features. This means evolving ambiguous business/product requirements into robust, scalable technical designs that might influence multiple services or feature areas. Your design documents should be exemplary.
        - **Impact & Delivery (L6 Aspect):** Ensure the solutions have a broader impact. For example, if your guardrail implementation for SQv2 establishes a pattern adopted by other teams, or if the CS sharding significantly improves stability/scalability for a whole platform area. Successfully delivering _multi-quarter impact_ projects would align with L6.
        - **Problem Difficulty & Ambiguity (L6 Aspect):** Proactively identify and solve _extremely complex or unforeseen technical challenges_ within these projects, going beyond the initial scope to ensure long-term success and robustness.
2. **Objective 2: Expand Technical Expertise and Operational Proficiency across Discover & Engage Systems**
    
    - **To "Meet Expectations" (L5):** Successfully complete on-call rotations, gain proficiency in k8s, CMP, Go, and contribute to operational tasks.
    - **To "Exceed Expectations" (L5) and demonstrate L6 potential:**
        - **Domain Expertise (L6 Aspect):** Become a _recognized Subject Matter Expert (SME)_ in one of these areas (e.g., k8s operations for Query Topics, or a go-to for CMP integration) across the engineering group or even the company. _Drive best practices_ related to these technologies.
        - **Operations (L6 Aspect):** Lead significant operational improvements that affect your _engineering group_, not just your immediate team. For instance, if your k8s learnings lead to a new monitoring strategy adopted by several teams, or you architect a more resilient operational pattern for a shared system.
        - **Ownership & Problem Solving:** Independently tackle and resolve complex, high-impact operational issues during on-call with minimal guidance, and drive systemic post-incident fixes.
3. **Objective 3: Champion Leadership Excellence and Personal Development**
    
    - **To "Meet Expectations" (L5):** Actively work on your chosen LEM behaviors, make process improvements, and mentor if opportunities arise.
    - **To "Exceed Expectations" (L5) and demonstrate L6 potential:**
        - **Hire and Develop the Best (L6 Aspect):** Become a _role model for technical interviews_. Actively mentor and _develop_ other engineers (L6 or below, as per SDE guide), showing tangible positive outcomes for their growth. This is more than just occasional guidance.
        - **LEM in Action:** Consistently demonstrate your chosen L5 LEM behaviors in a way that has a _visible, positive impact on the team or broader group_. For instance, if focusing on "Fosters Clarity," your efforts simplify complex topics for multiple people/teams. If "Scales Systems," your process improvements are adopted more broadly.
        - **Initiative & Influence:** If you pursue the UI exploration KR, ensure it leads to a well-researched proposal with potential for significant impact, perhaps influencing the team's roadmap or technical direction in that area.

**Key Factors for "Exceeds" and L6 Readiness:**

- **Going Beyond Scope:** L6 engineers typically expand the scope of their work, identify new opportunities, and tackle problems that aren't explicitly assigned.
- **Scope of Influence:** Your impact needs to be felt beyond your immediate tasks and team – influencing other teams, your engineering group, or technical standards.
- **Technical Leadership & Mentorship:** Guiding others, making sound architectural decisions for broader systems, and elevating the technical abilities of those around you.
- **"The How" - LEM Behaviors:** Consistently demonstrating Leadership Excellence behaviors, especially those related to broader impact, strategic thinking, and collaboration (e.g., Scales Systems, Judgment, Cross-Team Connection at a higher level).
- **Consistency:** Sustained high performance across multiple projects and aspects of the role throughout the year.
- **Business Need:** While your performance is key, promotions also factor in business need for an L6 role in your area.

**In summary:**

Your drafted objectives provide an excellent framework. To target an "Exceeds" rating and prepare for an L6 role, focus on:

1. **Maximizing the scope and impact** of your project deliverables (Objective 1).
2. Becoming a **recognized expert and driving broader improvements** through your technical and operational learning (Objective 2).
3. **Demonstrating strong technical leadership, mentorship, and LEM behaviors** that influence beyond yourself (Objective 3).

Proactively discuss your L6 aspirations with Vamsi. Ask for opportunities and feedback that will allow you to demonstrate these L6-level competencies within the scope of your 2025 objectives and other team initiatives. Achieving these objectives with a clear demonstration of L6 capabilities will significantly strengthen your case.

# As created 2025-05-30
I didn't copy the AI plan exactly.
## Transcription
Okay, here is the transcription of the provided document:

--- PAGE 1 ---

SuccessFactors: Goal Management

1 of 3

[https://hcm41.sapsf.com/sf/goals#/goal-print](https://www.google.com/search?q=https://hcm41.sapsf.com/sf/goals%23/goal-print)

Objective Plan: 2025 Objective Plan Category: Other Owner: Janzen Brewer-Krebs

Print Date: May 30, 2025

Champion Leadership Excellence and Drive Personal &

Team Development (Public)

Updated: May 30, 2025

Not Started

0%

Status

% Complete

The following table:

|   |   |   |
|---|---|---|
|**Key Results**|**Metric**|**% Complete**|
|Complete a Development Plan (Feedback 360)|Create and execute the plan|0%|
|Drive Broader Process/Documentation Improvements|Proactively identify, design, and implement 1-2 significant process or documentation improvements that are adopted by and demonstrably benefit the Query Topics team and potentially a broader engineering group.|0%|
|Mentor and Develop Team Members|Actively mentor 1-2 junior engineers or new team members on specific projects or technology areas|0%|
|Explore Ul Development Contribution (stretch)|Dedicate QED time to exploring Ul development/improvement opportunities for Query Topics tools/products. Present 1-2 findings, prototypes, or well-researched proposals to the team.|0%|

More Info

Weight:

20%

Start Date:

Jan 1, 2025

Due Date:

Dec 31, 2025

Objective Plan: 2025 Objective Plan Category: Other Owner: Janzen Brewer-Krebs

Print Date: May 30, 2025

5/30/25, 3:19 PM

--- PAGE 2 ---

SuccessFactors: Goal Management

2 of 3

[https://hcm41.sapsf.com/sf/goals#/goal-print](https://www.google.com/search?q=https://hcm41.sapsf.com/sf/goals%23/goal-print)

Expand Technical Expertise and Drive Operational

Excellence across Discover and Engage Systems (Public)

Updated: May 30, 2025

The following table:

|   |   |
|---|---|
|**On Track**|**15%**|
|Status|% Complete|

The following table:

|   |   |   |
|---|---|---|
|**Key Results**|**Metric**|**% Complete**|
|Participate in on-call covering our discover and engage systems|Take a fair share of the on-call load, one shift every 5-7+ weeks|40%|
|Master and Apply Kubernetes Expertise|Leverage learnings from k8s training to identify and implement at least two significant k8s-based operational improvements or deployment optimizations, sharing these improvements and resulting best practices with the engineering group.|0%|
|Gain Proficiency in CMP & Golang|Achieve hands-on proficiency with CMP, by utilizing my local setup to develop, test, and troubleshoot Smart Query v2 features (e.g., guardrail error handling for ENRICH-1295)|20%|
|Become a Knowledge Hub for Key Systems/Technologies|Author and present at least two comprehensive runbooks or knowledge-sharing sessions that demonstrably improve the team's operational efficiency or understanding.|0%|

More Info

Weight:

35%

Start Date:

Jan 1, 2025

Due Date:

Dec 31, 2025

Objective Plan: 2025 Objective Plan

Category: Create customer value by expanding beyond basics Owner: Janzen Brewer-Krebs

Print Date: May 30, 2025

5/30/25, 3:19 PM

--- PAGE 3 ---

SuccessFactors: Goal Management

[https://hcm41.sapsf.com/sf/goals#/goal-print](https://www.google.com/search?q=https://hcm41.sapsf.com/sf/goals%23/goal-print)

Deliver High-Impact Al Feature Enhancements and

Critical System Stability (Public)

Updated: May 30, 2025

The following table:

|   |   |
|---|---|
|**On Track**|**40%**|
|Status|% Complete|

The following table:

|   |   |   |
|---|---|---|
|**Key Results**|**Metric**|**% Complete**|
|Deliver Key Smart Query v2 Enhancements (ENRICH-1468 Epic)|wildcard use in Smart Query reduced by 80% &lt;br> prompt attacks detected and rejected with error Ul|40%|
|Implement project-level blocking for Classification Service|all messages for projects configured to be blocked are sent to DLQ, not processed|80%|
|Produce High-Impact Technical Designs|author and champion at least two comprehensive technical design documents for significant features or architectural improvements|0%|

More Info

Weight:

45%

Start Date:

Jan 1, 2025

Due Date:

Dec 31, 2025

3 of 3

5/30/25, 3:19 PM

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/Archive/slack thread cleaning example.md
---

---
created: 2025-06-12T09:57:56
---
# actual output
**Janzen Brewer-Krebs**: The conversation is (now) happening in #tapp-7074-bubble-widget-a11y-telekom. The main thread on this when it came in and was assigned to me is https://qualtrics.slack.com/archives/C08RHU9K8NS/p1747171354456159. It's unclear how much effort it'll be; there's lot of uncertainty. Widget code is old and weird, and this widget uses a graphics lib (d3) on top of that. They don't run locally and more stuff has probably broken since the last time anyone touched them. It's been a year+ since I touched widget code. On top of that, it's not yet clear what we must provide a11y-wise and what is just the customer having ideas about how to fix the a11y issues. (edited) Janzen Brewer-Krebs 
- https://qualtrics.atlassian.net/browse/TAPP-7074 is more like a feature request. The client is asking for new stuff, and bubble widget is weird old code. What do you think ? Thread in textiq-widget-transition | May 13th | View message

**Janzen Brewer-Krebs**: I just saw this doc with details on what the widget team expects for handoff, for context: 
- https://coda.io/d/_dqXDlMpiqcR/Handover-checklist_suHV_NH1

**Janzen Brewer-Krebs**: replied to a thread: Okay, will do. Since Smart Query was deprioritized for the other AI features, should we update its status in the planning coda? Following up on this View newer replies

**Vamsi Vippadapu**: hi Janzen , let us deprioritize the widget pulse until we have a connect with Rohit :+1::skin-tone-2: 1 8:03 can we bring the smart query story in to the sprint? it didn't come up in planning

# desired output
**Janzen Brewer-Krebs**: The conversation is (now) happening in \#tapp-7074-bubble-widget-a11y-telekom. The main thread on this when it came in and was assigned to me is https://qualtrics.slack.com/archives/C08RHU9K8NS/p1747171354456159.
It's unclear how much effort it'll be; there's lot of uncertainty. Widget code is old and weird, and this widget uses a graphics lib (d3) on top of that. They don't run locally and more stuff has probably broken since the last time anyone touched them. It's been a year+ since I touched widget code.
On top of that, it's not yet clear what we must provide a11y-wise and what is just the customer having ideas about how to fix the a11y issues. (edited)
> **Janzen Brewer-Krebs**:
> https://qualtrics.atlassian.net/browse/TAPP-7074 is more like a feature request. The client is asking for new stuff, and bubble widget is weird old code. What do you think ? Thread in textiq-widget-transition | May 13th | View message

**Janzen Brewer-Krebs**: I just saw this doc with details on what the widget team expects for handoff, for context: 
- https://coda.io/d/_dqXDlMpiqcR/Handover-checklist_suHV_NH1

**Janzen Brewer-Krebs**: replied to a thread: Okay, will do. Since Smart Query was deprioritized for the other AI features, should we update its status in the planning coda? Following up on this View newer replies

**Vamsi Vippadapu**: hi Janzen , let us deprioritize the widget pulse until we have a connect with Rohit :+1::skin-tone-2: 1 8:03 can we bring the smart query story in to the sprint? it didn't come up in planning

# source
Janzen Brewer-Krebs
  2:30 PM
The conversation is (now) happening in #tapp-7074-bubble-widget-a11y-telekom. The main thread on this when it came in and was assigned to me is https://qualtrics.slack.com/archives/C08RHU9K8NS/p1747171354456159.
It's unclear how much effort it'll be; there's lot of uncertainty. Widget code is old and weird, and this widget uses a graphics lib (d3) on top of that. They don't run locally and more stuff has probably broken since the last time anyone touched them. It's been a year+ since I touched widget code.
On top of that, it's not yet clear what we must provide a11y-wise and what is just the customer having ideas about how to fix the a11y issues. (edited) 




Janzen Brewer-Krebs
https://qualtrics.atlassian.net/browse/TAPP-7074 is more like a feature request. The client is asking for new stuff, and bubble widget is weird old code. What do you think 
@vamsi
?
Thread in textiq-widget-transition | May 13th | View message


Janzen Brewer-Krebs
  3:29 PM
I just saw this doc with details on what the widget team expects for handoff, for context: https://coda.io/d/_dqXDlMpiqcR/Handover-checklist_suHV_NH1


Janzen Brewer-Krebs
  9:47 AM
replied to a thread:
Okay, will do. Since Smart Query was deprioritized for the other AI features, should we update its status in the planning coda?
Following up on this
View newer replies


Vamsi Vippadapu
  8:02 AM
hi Janzen , let us deprioritize the widget pulse until we have a connect with Rohit
:+1::skin-tone-2:
1

8:03
can we bring the smart query story in to the sprint? it didn't come up in planning



On top of that, it's not yet clear what we _must_ provide a11y-wise and what is just the customer having ideas about how to fix the a11y issues. (edited) 

![](https://ca.slack-edge.com/E02U2N0QT37-U02Q9CDER9A-c7dc95998502-24)**Janzen Brewer-Krebs**

[https://qualtrics.atlassian.net/browse/TAPP-7074](https://qualtrics.atlassian.net/browse/TAPP-7074) is more like a feature request. The client is asking for new stuff, and bubble widget is weird old code. What do you think [@vamsi](https://qualtrics.slack.com/team/U055G1EH1HQ)?

[Thread in textiq-widget-transition](https://qualtrics.slack.com/archives/C08RHU9K8NS/p1747171354456159) | [May 13th](https://qualtrics.slack.com/archives/C08RHU9K8NS/p1747171354456159) | [View message](https://qualtrics.slack.com/archives/C08RHU9K8NS/p1747171354456159)

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/KB/AI coding.md
---

---
created: 2025-06-05T10:17:59
---
# Resources
- [Vibe Coding 101](https://coda.io/d/_dWvWppuuSTF/Vibe-Coding-101_suxRMgEp)
## Windsurf
- [Usage dashboard](https://windsurf.com/subscription/usage)
    - Resets mid-month?
- I use it as a plugin for windsurf
## Cursor
- [Usage dashboard](https://www.cursor.com/dashboard)
    - resets at end of month
- I use cursor IDE

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/KB/AOM.md
---

---
created: 2025-06-04T09:25:54
---
aka Alert on Metrics

AOM is an internal tool developed and maintained by the Engineering Visibility team. It is used for running Prometheus queries on a regular interval and comparing the results to a threshold. An alert is fired based on how often the thresholds are breached.
# Resources
- [AOM User Manual](https://odo.corp.qualtrics.com/wiki/index.php/AOM_User_Manual)
- [repo](https://gitlab-app.eng.qops.net/engvis/AlertOnMetrics)

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/KB/AWS.md
---

# Roles
- [`discover-engines-management-role` (engage prod account)](https://gitlab-app.eng.qops.net/secops/aws-policies/-/blob/master/Production%5B674592268301%5D/role_access/discover-engines-management-role.yaml?ref_type=heads)

# Accounts
## Engage/Q
- Prod: 674592268301
# Resources
- [Qualtrics AWS curriculum](https://odo.corp.qualtrics.com/wiki/index.php/AWS_Curriculum)
    - might be useful for system design, general learning

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/KB/CMP (aka Designer).md
---

# Resources
- [front door?](https://cxstudio.clarabridge.net/dashboard/#/login) for staging? or all? unsure.
    - oh, no this is for studio. i could pick an instance on the sidebar under Designer
- staging designer login page: https://cmp-staging01.staging.corp.qualtrics.io/cmp/login
    - replace the instance
    - https://qualtrics.slack.com/archives/C087VC5JPD3/p1745423790630519?thread_ts=1745267491.794059&cid=C087VC5JPD3
## Running CMP on ARM macs
There are several pages on this:
* https://coda.io/d/Discover-Data-Pipelines-Team-Homepage_dND2lFyMDGy/How-to-run-CMP-locally_sutwnAcs#_luO7rWJ_
* https://coda.io/d/Discover-Data-Services_dfW2SLhtzNH/Building-and-Running-CMP-for-M1-Macs_suogfpby#_luXdM7b7
* https://coda.io/d/Discover-Engines-Group_dUq7lnCJ_pd/Building-CMP-on-ARM-Macs-M1-M2_suTTvVJg
    * this one seems most recent and relevant
## How to get a project's ID
Find a designer page, inspect, search for project name. There'll be a list of option tags with the project id as an attr.
### Known projects

| instance      | name         | id       |
| ------------- | ------------ | -------- |
| cmp-staging03 | Autotune-GA1 | 21601992 |
## Uploading data
- [coda](https://coda.io/d/Query-Topics-Team-Hub_dt-izYOVJmq/Building-CMP-on-ARM-Macs-M1-M2_suJsoDWj#_lusE-5_V)
# Deployment
- it's deployed to k8s as a statefulset, not deployment. named something like `cmp-<color>` in `discover-data-pipelines` namespace
## Release process
### slack thread
Mavin Martin
A few questions regarding Designer(CMP) regarding boundaries?
When we merge an MR, are we responsible for deploying or letting UO team know?  Or do we wait until their next release?  Do we know what their release schedule is?
I'm getting a 504 gateway timeout for POST /cmp/projects_service  pulse.  Would we own a fix like this or would I transfer it to UO team?  My assumption is we only own RTDS piece of CMP?


Addisu Guddissa
1. Unless the fix for the pulse is needed to be deployed as a hotfix to unblock the customer, changes are simply merged in and data pipeline team deploys during regular monthly release cycle. If a hotfix is required, we coordinate with data pipelines team and cherry-pick the change to a release branch. Data pipeline team will do the deployment here as well.
Release schedule starts in first week of the month and ends somewhere in 20th of the month
2. Gateway timeout doesn't seem to be Query Topics pulse. I would say transfer it to UO / Data pipelines team
# Troubleshooting
## Smart query isn't working
server.log:
```
com.clarabridge.discovery.feign.requests.FeignRequestsExecutor$UnknownClientException: Can not resolve client name for: com.clarabridge.discovery.feign.clients.smartquery.RemoteSmartQueryService
```

https://qualtrics.slack.com/archives/C087VC5JPD3/p1746478157136529

Need to add config values to `cmp-compose/cmp-compose-arm/resources/server/conf/properties-service-custom.properties`
```
# SmartQuery
discovery.client.mig.token=<STAGE MIG TOKEN>
discovery.client.sq.url=[https://b1-prv.qualtrics.com](https://b1-prv.qualtrics.com/)
discovery.client.sq.readTimeout=300000
```

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/KB/Classification Service.md
---

I sometimes call it CS.

# Resources
- [Deployment job for prod CA](https://jenkins-shared.clarabridge.io/job/DEPLOYMENTS/job/APPLICATIONS/job/ENGINES/job/SERVICES/job/CLASSIFICATION-SERVICE/job/PROD3_V2/job/classification-service-deploy/)
    - [Prod EU](https://jenkins-shared.clarabridge.io/job/DEPLOYMENTS/job/APPLICATIONS/job/ENGINES/job/SERVICES/job/CLASSIFICATION-SERVICE/job/PROD2_EUW2/job/classification-service-deploy/)
    - [Prod APSE](https://jenkins-shared.clarabridge.io/job/DEPLOYMENTS/job/APPLICATIONS/job/ENGINES/job/SERVICES/job/CLASSIFICATION-SERVICE/job/PROD4_APSE2/job/classification-service-deploy/)
- [Deployment configs per env](https://gitlab-app.eng.qops.net/xm-discover/automation/deployments/-/tree/master/multicloud-kubernetes/applications/engines/services/classification-service/envs?ref_type=heads)
- [Discover AWS ECR for CS image](https://us-east-1.console.aws.amazon.com/ecr/repositories/private/799231120095/plat/classification-service)
- [Coda runbook](https://coda.io/d/_dt-izYOVJmq/Classification-Service_suAa9o2-) (kinda old)

# How to deploy
## Discover
- Are you changing anything with rabbitmq? Do you need to notify Unified Orchestration?
- Check the image tag in deployment config to see what you're deploying
    - Update image tag if needed?
- Find a recent successful jenkins "build"/deployment and rebuild it
- Verify build parameters
    - namespace: ingestion-pipeline-0
    - branch: master
    - rmq_vhost: ingestion-pipeline-prod
        - -dev in dev
        - -staging in staging
    - scale_mode: QueueLength
    - scale_value: 100
        - 2000 in prod US?
    - queues_operation: sum
    - only golden classification and golden es checked
    - no flags
# How to test changes
We can deploy built images to staging envs for testing.
1. Push git branch to gitlab
2. Find your branch in [jenkins](https://jenkins-shared.clarabridge.io/job/Engines%20CI%20GitLab/job/xm-discover%252Fdiscover-engines-group%252Fclassification-service/)
    1. Optional: click Scan GitLab Project Now to speed it up
    2. Optional: check console output for the build job for info on image tag
3. Log in to discover AWS console
4. Go to the relevant image in us-east-1 ECR [link](https://us-east-1.console.aws.amazon.com/ecr/repositories/private/799231120095/plat/classification-service?region=us-east-1)
    1. Your build is likely at the top if it just built, look for branch name
5. Copy the image tag you want to use (specific build or latest build from the branch)
6. I guess you can deploy a couple different ways:
    1. Open relevant pre-prod k8s control plane (k9s, lens)
        1. Set the image directly on the CS deployment
    2. Do a normal jenkins deployment 
        1. Need to override the image tag? Idk, haven't done this yet.
# Deployment notes from reverse shadowing Kuan
- get access to https://vault-shared.clarabridge.io/ui/vault/secrets/secret-prod2-euw2/kv/infra%2Frabbitmq%2Fcb-ingestion-pipeline-prod2-euw2-rabbitmq%2Fconnection-info/details, https://vault-shared.clarabridge.io/ui/vault/secrets/secret-prod2-euw2/kv/infra%2Frabbitmq%2Fcb-ingestion-pipeline-prod2-euw2-rabbitmq%2Fconnection-info/details
- deployment job https://jenkins-shared.clarabridge.io/job/DEPLOYMENTS/job/APPLICATIONS/job/ENGINES/job/SERVICES/job/CLASSIFICATION-SERVICE/job/PROD2_EUW2/job/classification-service-deploy/
- figure out how to get access to prod us rabbit mq. it's different. self hosted?
- I have a recording of this, it was on May 8 2025


---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/KB/CrateDB.md
---

# Resources
- [grafana dashboard](https://grafana.eng.qops.net/d/liDiIrwVk/text-application-deployed-versions?orgId=1&from=now-6h&to=now&timezone=browser&var-main_container_name=$__all) with currently deployed versions of fssync and crate by DC
# Open admin console
`sectool ssh connect iad1 -o` -L 4300:texteng-dbms.service.consul:4200 -Nn


---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/KB/Discover DCs.md
---

# Prod US
aka prod1?
us-east-1?
# Prod UK
prod eu
prod 2
eu-west-2
# Prod CA
prod3
prod3_v2
ca-central-1
# Prod ANZ
prod4
prod apse2
ap-southeast-2
# Dev
# Staging

# In qualtrics/engage DCs
Discover stuff is deployed in a few engage DCs: **fra1, gov1, b1, g1**

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/KB/Discover Resolution Team.md
---

---
created: 2025-06-25T08:21:12
---
aka res

They handle some tasks for us when investigating pulses, like database querying, republishing models. They do more than Qualtrics resolution does.

When I have a request for them, I should tag the pulse requester, not the res on-call

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/KB/Docker.md
---

---
created: 2025-06-17T12:56:03
---
# Resources
- [Docker Image Metadata Standard](https://docs.google.com/document/d/1yQhcBLNT3V6lvnAL0oG_xaSFQe9MfRiwohEhrKYUe_Y/edit?tab=t.0#heading=h.46exkcj9fbou)

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/KB/Elasticsearch.md
---

owned by core data team (aka enrichments core data?)

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/KB/Ingestion Gateway (IG).md
---

---
created: 2025-06-03T13:59:32
---
# How to send documents to IG
Useful for making [[Classification Service]] run. [Slack thread](https://qualtrics.slack.com/archives/C087VC5JPD3/p1748982999194399?thread_ts=1748981906.184879&cid=C087VC5JPD3)



---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/KB/Jargon.md
---

---
created: 2025-06-05T15:38:00
---
# Team/Text Analytics
- TAM: text analytics module
    - generally means the Text iQ UI. technically the scala server code is also in the same repo and *could* be referred to the same way.
- texteng: text-engine, text engine
- IG: ingestion gateway
    - the pipeline of services that process discover documents, including classification service. uses [[RabbitMQ]].
# Qualtrics
- 2LA: second-level analysis
    - reviewing recent bugs to look for trends
- AI: action item
- CY: code yellow
    - ordinary work is paused while a serious situation is addressed, with clear exit criteria
- CR: code red
    - like a code yellow, but employees are expected to work around the clock to resolve an emergency situation
- engvis: Engineering Visibility team
- GA: general availability
- KT: knowledge transfer
    - usually in the context of moving code/feature/product to another team or the expert on the same leaving/moving teams
- LIG: lambda inference gateway
- LRR: launch readiness review
    - a (tedious) formal process for launching new features
    - jira project key LAUNCH
- OKR: objective, key result. old company-wide system for goal tracking, replaced by objective plans
- pulse: a customer bug
    - alternatively, a periodic survey sent to employees, as in "Q1 pulse"
- res: resolution team
    - they're closer to the customer than us engineers when working on customer bugs/pulses
- RCA: root cause analysis
- SME: subject matter expert
- TAM: technical account manager
- UO: unified orchestration
    - a team pretty close to query-topics. think they manage the IG pipeline.
- vuln: vulnerability
- WAD: works as designed
# Me
- qt: query topics (my team)
# Tech
- tf: terraform
- DLQ: dead-letter queue
- ES: elasticsearch

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/KB/Jira.md
---

- [KTLO Q2 2025 epic](https://qualtrics.atlassian.net/browse/ENRICH-1513) (from [[2025-05-27 Tue - Sprint Retro#^cd6c41]])

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/KB/Kafka.md
---

aka msk

[slack message with some kt resources](https://qualtrics.slack.com/archives/C036JC1JKGE/p1745579740133489?thread_ts=1745326776.171129&cid=C036JC1JKGE)

[terraform config](https://gitlab-app.eng.qops.net/xm-discover/discover-cloudeng/aws-cloud-team/AWS-Devops/-/tree/DISCCLOUD-25081/terraform/msk)
    [example mr expanding storage or something](https://gitlab-app.eng.qops.net/xm-discover/discover-cloudeng/aws-cloud-team/AWS-Devops/-/merge_requests/1893/diffs)

[old msk provisioning doc](https://coda.io/d/Discover-Data-Operations_d_rGTDGV7jC/MSK-Provisioning_sudr2uVt?searchClick=81b24065-f442-46d7-be83-99d1d5606bcf__rGTDGV7jC#_luNdetMm)



---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/KB/Kibana.md
---

https://logs.clarabridge.net/app/kibana#/discover
need openvpn

Using Kibana: https://coda.io/d/PM-Internal-Wiki-For-PM-Eng_dmj04g2bekN/Using-Kibana_suxV5q0d?searchClick=08ee672e-960d-4776-bc81-c50949e99e15_mj04g2bekN



---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/KB/Kubernetes.md
---

aka k8s

[Kubernetes boot camp](https://docs.google.com/document/d/19jSRoYrfKWmxZu76l10kNkm_qixGL_f8Meq624rIkMs/edit?tab=t.0#heading=h.75wn12mfpyde)
[Kubernetes trainings folder](https://drive.google.com/drive/folders/1HAn5FJ5xHE0qJdC8D-0tKH-mf-NVhRUB)

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/KB/LBaC.md
---

---
created: 2025-06-05T11:00:11
---
aka load balancer as code
# Repo readme snip
## Overview
This repository contains the configuration files for load balancing of the services registered in Consul in a self-service fashion. Once configured, your clients will be able to reach your service using Stargate URLs:
### Stargate URLs
* for external clients by making requests to
  ```
  https://<BRAND>.<DC>.qualtrics.com/<YOUR_SERVICE_PREFIX_URL>
  ```
* for internal clients by making requests to
  ```
  https://internal-lb.<DC>.(qprod|qops).net/<YOUR_SERVICE_PREFIX_URL>
  ```
Your configuration will be applied to a fleet of Stargate servers which are present in each staging and production datacenter. Stargate servers will allow your client to reach your services instances registered as Consul service members.
### Glossary
1. **Stargate** or **Stargate Envoy** - is a load balancer capable of directing traffic to one of the services behind by using unique URL prefixes. It is based on [Envoy](), an L7 proxy and communication bus designed for large modern service oriented architectures. Each staging and production datacenter in Qualtrics contains number of Stargate nodes, from 10 in G1 to 63 in IAD1. You can use [Stargate Envoy Dashboard](https://grafana.eng.qops.net/d/5Yj0l7KGz/stargate-envoy?orgId=1&refresh=1m) to see exact number of nodes in given DC.
2. **Stargate xDS** - is a control plane for **Stargate**. The acronym stands for Aggregated Discovery Services which is a programmatic API for configuring Envoy. It uses LBaC configuration files and Consul to discover routes and members of services which need to be proxied and then dynamically configure Stargate servers.
3. **Scope** - this is a way of controlling how the requests make their way to your service. Currently we have two scopes available: **external** and **internal**. If you service needs to be reachable only by clients inside the Qualtrics network (internal clients) then you should add `scope: internal` to your configuration. Otherwise, if you want to expose you service to external clients then add `scope: external`.
> Keep in mind that all services marked as `external` will also be available via the path for `internal` clients.

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/KB/LDAP.md
---

Can verify the ldap groups I'm in by going to a bastion host and running `id -Gn jbrewer`

```
[prod.yul1][jbrewer@rampart4-app ~]$ id -Gn jbrewer
engr docker aws-289231305718-power-users text qcurriculum all fedramp discover-engineering text-application text-orchestration text-api textanalytics
```


---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/KB/Model Inference Gateway (MIG).md
---

# About
Managed by Core Model Engineering team (CME)

# Resources
- [Usage by issuer dashboard](https://grafana.eng.qops.net/d/01TQSUvSk/mig-http-metrics?orgId=1&from=now-24h&to=now&timezone=browser&var-API=%2Fmodel-inference-gateway%2Fguardrails%2Fv1%2Fapply&var-DCs=$__all&var-Issuer=smart-query)
- [OpenAPI spec](https://periscope.eng.qops.net/catalog/default/api/model-inference-gateway/definition)
- [Rate limit doc](https://gitlab-app.eng.qops.net/iq-research-sea/model-inference-gateway/-/blob/main/documentation/consumer_docs/rate_limiting.md)
# Rate limiting
Can be applied at the issuer, brand, or user level. Almost everything uses issuer throttling.
- [example of user throttling](https://gitlab-app.eng.qops.net/iq-research-sea/model-inference-gateway/blob/a9d63eda04702c3897bd9d2d90c4106ba47c7d34/internal/config_managers/llm_v1/throttling_config_ai_playground.go#L18)
- [slack discussion](https://qualtrics.slack.com/archives/C0186S22PQS/p1749147003594579)
## "watermark"
The limits are expressed as high and low "watermarks". Unintuitively, the high watermark is additive; its value is added to the low watermark to get the peak limit. [source](https://qualtrics.slack.com/archives/C0186S22PQS/p1749153321819379?thread_ts=1749153231.447079&cid=C0186S22PQS)

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/KB/Mongo.md
---

https://odo.corp.qualtrics.com/wiki/index.php/TORQ_Runbooks#Connecting_to_Mongo

Tunnel to connect:
`$(sectool ssh connect yul1 -o) -L 21339:textanalytics-mongo-primary.service.consul:21337 -Nn`

I also wrote `mongo-tunnel` to do this in my scratch repo

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/KB/On-call.md
---

---
created: 2025-06-11T08:58:18
---
# Time tracking
[source](https://coda.io/d/_dt-izYOVJmq#Discussions-Action-Items_tuSJhIr6/r1&view=center)
> The goal is to track on-call rotation hours, not follow up rca, AI, continued pulse resolution, follow up bugs.

[source](https://qualtrics.slack.com/archives/C087VC5JPD3/p1750703837216369?thread_ts=1750680015.007789&cid=C087VC5JPD3)
> not counting planned vuln fixes
> \[if] it is unplanned because of past/near sla then yes

# Roles
## Primary
This is what we mean when we say "the on-call". Gets paged for alerts, has various responsibilities related to ktlo, ops, vulns, pulses.

Since we don't have a formal list of things the primary should do, I'll keep my own.
### Things to do
- Check vulns on [vmx dashboard](https://vmx.eng.qops.net/dashboard/2-team-vulnerability-dashboard?tab=1-team-dashboard&team_name=discover-engines&team_name=discover-nlp-science&max_sla_in_days=&datacenter=) ([coda source](https://coda.io/d/_dt-izYOVJmq/Query-Topics-Onboarding_suzukYeL#_luNcR8FP))
## Secondary
I asked for details on the role of the secondary and got this.
[Vamsi](https://qualtrics.slack.com/archives/C087VC5JPD3/p1749764884841819?thread_ts=1749763974.282259&cid=C087VC5JPD3):
> I don't know if it warrants detailed documentation. The idea is to share some load if the primary on-call is over loaded.



---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/KB/Qualtrics.md
---

---
created: 2025-06-03T10:23:44
---
The company I work at is Qualtrics. I work on text analysis stuff. I used to be on a team called Text Application that worked on Text iQ. But Text iQ is being deprioritized, and in January 2025 I was moved to the new Query Topic (query-topics billing code) team under [[Vamsi]].

This new team is a mix of "discover" and "engage" people. "Discover" is the term for the stuff belonging to an acquisition that hasn't been integrated--clarabridge. "Engage" means Qualtrics stuff. Kind of a weird term.


---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/KB/Query Topics AI Features (lambdas).md
---

---
created: 2025-06-04T10:08:16
---
# Resources
- [Deployments of Lambda functions (coda)](https://coda.io/d/_dUq7lnCJ_pd/Deployments-of-Lambda-functions_sutu_YxH)
## Repos
There are two repos. Maybe one is for deployment and one for dev?
- [deployment?](https://gitlab-app.eng.qops.net/iq-research-sea/iq-aws-resources-manager)
- [dev?](https://gitlab-app.eng.qops.net/xm-discover/discover-engines-group/query-topics-ai-integrations)
# Summary of Lambda Deployments
- [source](https://qualtrics.slack.com/archives/C07MKL3FB61/p1749162967383299)

- **Deployment Location:** The AI lambdas are **not** deployed in Discover data centers. They are all deployed in **Engage AWS accounts**.
- **Access Role:** To access these lambdas, you need the `socrates-human-customer-access-role`.
- **Routing from Discover to Engage:** When a feature in a Discover data center (like Designer in Prod US) needs to use an AI feature, it makes a call to the **internal load balancer** of the corresponding Engage data center. For example, a call from Discover's Prod US environment is routed to `internal-lb.iad1...`. This cross-environment communication is handled via **Stargate**.
- **Deployment Mapping:** The deployment work (like in ticket ENRICH-1398) involves deploying a corresponding lambda function in an Engage data center for each Discover data center.
    - Discover AWS US & IBM US lambdas are in **iad1**.
    - Discover AWS & IBM Staging lambdas are in **g1**.
    - Discover AWS EUW2 & IKS \[IBM?] EU lambdas are in **lhr1**.
# Smart Audit
aka Auto Tune (old name)

Charmish works on this
# Smart Query
## Resources
- [b1 batch lambda (this is what local cmp hits)](https://us-west-2.console.aws.amazon.com/lambda/home?region=us-west-2#/functions/coreml-socrates-smart-query-batch-b1-prv?tab=monitoring)
- lambda-inference-gateway ([routes file](https://gitlab-app.eng.qops.net/coreml-model-engineering/lambda-inference-gateway/-/blob/main/src/endpoints/lambda_route.py))
    - routes SQ requests from MIG to our lambda

(copied from [[2025-05-07 Wed]])
- Looking for the relevant code
    - This is the repo for the lambdas https://gitlab-app.eng.qops.net/xm-discover/discover-engines-group/query-topics-ai-integrations
- Where are the prompts? They come from PLMS? That's prompt-lifecycle-management-service
    - https://gitlab-app.eng.qops.net/iq-research-sea/prompt-lifecycle-management-service
    - Yes, but where in plms? I guess the prompts aren't source controlled with the service code.
    - The lambda requests a prompt named `sq-gen-prompt-dict-0`. Sourcegraph search only turns up the lambda use of that string.
    - Asked for help https://qualtrics.slack.com/archives/C07MKL3FB61/p1746635272993699
    - Addisu says prompts are uploaded to PLMS. PLMS docs say they're replicated out to all DCs from g1.
- I'm getting timeouts trying to access PLMS in g1
    - Asked socrates on-call for help https://qualtrics.slack.com/archives/C0186S22PQS/p1746636628407509
    - I can't hit internal-lb for g1 from my local machine. Only b1.
- Retrieved a prompt from PLMS
## Thread on MIG (aka CME, socrates) guardrails
- [slack](https://qualtrics.slack.com/archives/C0186S22PQS/p1742922021578449)
**Topic:** Smart Query is consistently failing its `prompt_attack` guardrail.

**Problem:** Addisu reported that all Smart Query calls are being blocked by the `prompt_attack` guardrail, regardless of the strength setting. He correctly suspected that the guardrail was analyzing the entire prompt—including the developer-defined instructions—instead of just the parts provided by the user.

**Analysis:**

- **Root Cause Confirmed:** James from the MIG/Socrates team confirmed this is the issue. By sending the entire prompt (with instructions) as a single "Human" message to the LLM API (`/llms/v1/.../chat`), the guardrail correctly interprets the instructions as a potential attempt to override the model, which is a form of prompt attack.
- **Proposed Solution vs. Platform Reality:** Addisu asked if they should use a specific Amazon Bedrock feature to tag user content for the guardrail. James clarified that the internal Qualtrics platform does **not** support this feature.

**Resolution:** The platform-recommended solution is to use the **standalone Guardrail API (`/guardrails/v1/apply`)**. The correct workflow is:

1. Isolate the parts of the prompt that come directly from the user (e.g., the `useCase`, `industry`, and the `<inputs>` block).
2. Send **only these user-provided parts** to the standalone `/guardrails/v1/apply` endpoint to check for prompt attacks.
3. If the guardrail check passes, then construct the full prompt (developer instructions + user inputs) and send it to the main LLM API (`/llms/v1/.../chat`) for a response, making sure to disable the `prompt_attack` guardrail for _this_ final call since the user input has already been vetted.
## Smart Query: Origin of Asterisks (Wildcards)

**Date of Discovery:** 2025-05-20 (based on conversation with Leo)
**Related Ticket(s):** [[ENRICH-1114- Smart Query Prompt Updates]]

### Key Discovery:
The asterisks/wildcards in Smart Query keywords originate from **both** the AI model's direct output and subsequent Python post-processing logic.

---

#### 1. AI Model Output
* The AI model itself directly includes asterisks in the initially generated `query_terms`.
* The raw response from the model, visible in logs *before* any Python post-processing, shows terms ending with asterisks.

#### 2. Python Post-Processing (`smart-query-batch.py`)
The `smart-query-batch.py` script then refines these keywords through several functions:

* **`_remove_overlapping_keywords` function:**
    * This function can **add** an asterisk to a term.
    * Condition: It adds an asterisk if the term is a substring of another existing keyword *and* does not already end with an asterisk or a quote.
    * Example logic snippet: `if len(rest)+1 < len(keywords) and first[-1] != '*' and first[-1] != '"': first = first + '*'`.
* **`_clean_term` function:**
    * This function can **remove** asterisks.
    * Condition: It removes asterisks from terms that are enclosed in quotes.
    * Example logic snippet: `if term.startswith('"'): term = term.replace('*', '')`.

#### Summary:
The final set of keywords with asterisks in Smart Query is a result of the AI model's initial suggestions, which are then further refined by heuristic-based logic within the Python post-processing steps.

# Topic Recommendations
## Extra local config
[source](https://qualtrics.slack.com/archives/C07MKL3FB61/p1750102424386319)
```
# Topics recommendations
discovery.client.trn.url=https://b1-prv.qualtrics.com
discovery.client.trn.readTimeout=300000
discovery.client.trn.maxSentencesToFetch=2000
discovery.client.trn.maxTopicsToRecommend=10
discovery.client.trn.batchSize=500
discovery.client.trn.useRandomScore=false
```

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/KB/RabbitMQ.md
---

Find web ui url and credentials at https://vault-shared.clarabridge.io/ui/vault/secrets/secret-prod/kv/infra%2Frabbitmq%2Fcb-ingestion-pipeline-prod-us-rabbitmq%2Fusers%2Fread_only/details?version=1 (for prod us)

# Management consoles
## Prod US
	https://b-ca71e892-d532-4588-b5f0-ee0700d2e73d.mq.us-east-1.amazonaws.com
- ^ nope
- while we're on the self-hosted one, have to port forward from k9s. see DMs with andy for details, cred.
## Prod CA
- [admin credential](https://vault-shared.clarabridge.io/ui/vault/secrets/secret-prod3/kv/infra%2Frabbitmq%2Fcb-ingestion-pipeline-prod3-ca1-rabbitmq%2Fconnection-info/details?version=1)
- [management console](https://b-dacbfe21-732d-482a-a4b6-dbe5036bf79d.mq.ca-central-1.amazonaws.com)
## Prod EU
- [admin credential](https://vault-shared.clarabridge.io/ui/vault/secrets/secret-prod2-euw2/kv/infra%2Frabbitmq%2Fcb-ingestion-pipeline-prod2-euw2-rabbitmq%2Fconnection-info/details?version=1)
- [management console](https://b-0c5927a3-8242-43d1-85b0-0ab5dde76deb.mq.eu-west-2.amazonaws.com/)
## Prod APSE2
- [admin credential](https://vault-shared.clarabridge.io/ui/vault/secrets/secret-prod4-apse2/kv/infra%2Frabbitmq%2Fcb-ingestion-pipeline-prod4-apse2-rabbitmq%2Fconnection-info/details?version=2)
- [management console](https://b-66da7603-b39e-4861-84f1-5cc491c83ac5.mq.ap-southeast-2.amazonaws.com)

# Snips
## rmq scaling, discover grafana is unreliable 2025-05-27 [slack](https://qualtrics.slack.com/archives/C035CQ3CHUK/p1747935205399879)
Robert Shields
  Thursday at 10:33 AM
:thread: For Autoscaling based on RMQ

Robert Shields
  Thursday at 10:34 AM
@akashv and @andy to discuss the state of autoscaling with the current RMQ self hosted broker.

Andy Morato
:pepe-smoke:  Thursday at 10:43 AM
My understanding is:
In engage - we publish the relevant queue depths to grafana to be the autoscaling metrics. enrichers consume those metrics
In discover - we have scaled objects for enrichers that pull the queue depths directly from the RMQ API
The challenge in discover is that autoscaling configurations (scaled objects) have to change when our backend message broker changes. We'd like to introduce a layer of abstraction so enrichers wouldn't need to directly interact with the broker but that is difficult in discover since per @akashv the discover grafana is unreliable

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/KB/Response Tools.md
---

---
created: 2025-06-13T08:57:46
---
aka (and technically correct) responses tools

Internal Qualtrics web UI for reaching into our databases and infrastructure. My teams specifically use "Text Tools".

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/KB/Splunk.md
---

---
created: 2025-06-04T08:38:57
---
- [coda page with common stuff](https://coda.io/d/_dt-izYOVJmq/Common-Splunk-Operations_su7vza_r)
# Useful values
## `index`
- index::qualtrics_discover_dev
- index::qualtrics_discover_stg
- index::qualtrics_discover_prd
## `sourcetype`
- sourcetype::kube:container:classification
- sourcetype::kube:container:cx-designer
    - cmp/designer

## cmp/designer
- > `namespace="designer"` is also a good filter
- > I haven't had to query cmp directly, but one of these two? `source="/var/log/pods/designer_cmp-*"`  `pod="cmp-*"`

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/KB/Sprint Review Meeting.md
---

---
created: 2025-06-24T09:36:42
---
As of Jun 10 2025, we focus on demos over saying what we did. [[2025-06-10 Tue#^cd1392]]
> use the slides for demos.  You can include screenshots and stuff.  But sprint status shouldn't be included. Example of demos:
>     where you are in a feature and a blocker you want team feedback on for next steps
>     new feature you added to show team how it works for awareness
>     potential updates that could potentially see pulses
>     KTLO improvements for awareness

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/KB/Teams.md
---

Discover Data Pipelines is now called Unified Orchestration? 2025-05-20

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/KB/Topic Recommendations (Discover feature).md
---

---
created: 2025-06-02T15:52:36
---
aka TR

This is a sibling AI feature to [[Query Topics AI Features (lambdas)]]. Seems Addisu has worked on it the most.

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/KB/Topic Recommendations (Text iQ feature).md
---

aka recommended topics, trex

[Topic Recommendations Job Investigation Guide](https://docs.google.com/document/d/1A6wM1rmEsSW8SWL8Ji72axpG8nqR2nUow6PBDWpLU-c/edit?tab=t.0#heading=h.dpbv9aljfboe)

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/Meeting notes/Retro/2025-05-27 Tue - Sprint Retro.md
---

# Ideas
- ~~Deploying via jenkins is time-consuming. Do we have work planned to improve the situation?~~
- Clarify what time is counted in the ktlo/ops count
    - Counting time spent helping teammates with cross-discover/engage pulses? Include in KTLO/ops?
- Sprint review often runs over. Could we extend the duration or timebox?
# Sprint Review Discussion Points: Enhancing Sprint Planning & Backlog Management

- [x] Review this doc ⏫ 🛫 2025-05-27 📅 2025-05-27 ✅ 2025-05-27

---
(written by AI)
## Topic: Enhancing Sprint Planning & Backlog Management

### 1. Observation & Context (What I'm Seeing)

* I've been reflecting on our sprint planning and how work moves from the backlog into active sprints.
* **Observation:** I've noticed a few instances where it seems items might not get proactive attention until their due dates are very close, or they're flagged as urgent.
    * *(Optional brief example for my reference: e.g., the recent discussion around ENRICH-1283 – unassigned in backlog with an imminent due date).*
* This isn't about any specific ticket, but more a **pattern** I'm trying to understand.
* My past experiences with sprint planning haven't typically involved this particular dynamic of items surfacing late.
* I'm wondering if our current "sprint loading" approach (how we select and bring work into the sprint) might have gaps we could address.

### 2. Question/Concern (Is this a broader issue for the team?)

* My main question for discussion:
    > **Are we, as a team, confident that our current backlog management and sprint planning process effectively prevents important work from slipping through the cracks?**
* Specifically, are we ensuring that there's consistent, **proactive oversight of the *entire* backlog** – including items that might be lower priority initially or unassigned – to make sure they are correctly prioritized and actioned in time?

### 3. A Thought/Potential Area for Improvement (How we might address it)

* From my perspective, it feels like we could benefit from:
    * A clearer, perhaps more **dedicated, focus or role** responsible for continuous, holistic backlog grooming and prioritization.
    * This would involve regularly reviewing *all* items, ensuring they align with our goals, have clear next steps, and are flagged for sprint consideration *before* their deadlines create urgency.
* The aim isn't to add burden, but to ensure **smoother workflow and reduce reactive situations.**

### 4. Open to Discussion (Inviting team input)

* I wanted to bring this up to see if others share this observation or have thoughts on how we might strengthen our approach here.
* What are your perspectives? How can we best address this as a team?

# Meeting notes
- [Meeting coda](https://coda.io/d/Query-Topics-Team-hub_dt-izYOVJmq/2025-Q2-4-26-5-26_su2lOKxg)
- (I raised) sprint review runs long
    - vamsi says we'll timebox for now, but when team is larger in Q3, we should extend
    - addisu wants to extend now. wants to not discuss pulses, discuss in 2la instead. wants demos.
    - [x] Ping Vamsi about what we should be including in sprint review 🛫 2025-06-09 📅 2025-06-09 ✅ 2025-06-09
    - vamsi likes focus on demos. he's going to extend to 1.5 hours too.
- (I raised) sprint work prioritization
    - vamsi's going to make a ktlo epic and add upcoming due items and "bring attention" during sprint planning ^cd6c41
        - [x] Are we making jira items in the ktlo epic for *everything* we're doing that's ktlo on our on-call shift? ✅ 2025-05-27 ^6a6166
- (I raised) what on-call time is counted?
    - barely discussed. vamsi took an action item to send a summary note on the discussion on this.
# Log
- [x] Check on action items from this retro (see coda) 🛫 2025-06-03 ✅ 2025-06-03

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/Meeting notes/Retro/2025-06-24 Tue - Retro.md
---

---
created: 2025-06-13T08:23:39
tags:
  - meeting/retro
---
# My agenda
- [x] Document roles of primary and secondary on-calls #meeting/retro ✅ 2025-06-24
    - Discussion [here](https://qualtrics.slack.com/archives/C087VC5JPD3/p1749763974282259). Vamsi doesn't want to document secondary.
    - I suggested the idea of using a shift note template like we did in TAPP
    - **No action taken. Vamsi wrote the below line in the action item column 🤷**
        - > Primary on-call responsible for responding to victor ops, alerts and triaging pulses. Secondary on-call to help share the load.
        - In *this* meeting, there was confusion about whether the secondary on-call takes work off the KTLO backlog when they're not immediately engaged, or if they do normal feature/OKR work.
- [x] Document jira labels, when to use them #meeting/retro ✅ 2025-06-24
    - **No action item, something about using learnings from 2LA**
- [x] When do we review the KTLO backlog? #meeting/retro ✅ 2025-06-24
    - I honestly forget if we do
    - Perhaps having a meeting agenda/notes page for our meetings would help formalize this
    - **It's during sprint planning**
- [x] More structure around meeting scheduling ✅ 2025-06-24
    - I'd like the calendar to be up to date farther out
    - **Vamsi agreed to do so for sprint review, retro, and thursday async scrums**

# Meeting notes
- 

# After meeting
- [x] Review these notes and create tasks if needed #meeting/retro ✅ 2025-06-24
- When Vamsi defends the status quo, I feel frustrated that my efforts to improve the team are not recognized and accepted. Or at least considered.
- [x] Ask Vamsi to clarify whether secondary on-call works exclusively on on-call stuff (including ktlo backlog), or if they're only pulled in as needed by the primary on-call. ✅ 2025-06-26

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/Meeting notes/Vamsi 1 on 1/2025-05-01 Thu - Vamsi 1 on 1.md
---

# My plan
* Explicitly say I'm going to give more direct feedback and hope we can improve our relationship and the team.
    * did this
* I experience his direction as stressful. Part of this is on me. Part is on him.
    * did this
* Remind him of my pronouns
    * did not do this
* Suggest being less dismissive, especially in retro. The "shared trauma" thing stands out. He seems to dismiss things that he doesn't like? Doesn't agree with? Are inconvenient?
    * Suggest he's overexplaining himself? Say I do the same.
    * did not do this

# Meeting notes
* i'm very focused on one thing at a time, could split my attention over a few things so that low pri tasks don't get starved
* we'll try weekly 1-1 to check in more often about what i'm working on
    * i offered to send him my daily notes. don't think he took me up on it.
* he says get on smart query v2 now
    * i think he mentioned something about getting it done in q2, so i brought up how i thought he'd said there was less time pressure at sprint review on tuesday ([[2025-04-29 Tue#Log]]) and he said he hadn't meant that.
    * no clue if i'm misunderstanding or he's miscommunicating
* he wants me to make "judgment calls". for example with message blocking, the priority is not having another sev 2. could file a follow up item for integration tests.
* he says if i understand the current smart query feature, i will see it's not that bad, will relieve some stress for finishing in q2
* he says make tasks in jira for unplanned on-call work like fixing crate deployments
* [x] make jira tasks for my unplanned on-call work ✅ 2025-05-01
* he brought up my low rto score. i argued that no one is there, he said there are other benefits and he gets a bad number with his manager. i said i filed a medical accommodation request to be remote. we didn't finish the conversation before i had to leave.

# Thoughts after
* i felt we miscommunicated a lot. i felt i explained some things very clearly that he did not get, and i'd have to re-explain. i worry about what else got miscommunicated
* i felt i gave a lot of grace in giving him some more direct feedback and did not receive any. it felt like a one way opening.
* i do not want to work for him. i guess i kinda have to give it some time though.

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/Meeting notes/Vamsi 1 on 1/2025-05-15 Thu - Vamsi 1 on 1.md
---

# My plan
- [x] Apologize for losing my head in the AI sync meeting this morning ✅ 2025-05-15
    - What could I have done better?
        - More communication. Clarify scope. Review my understanding.
- [-] Review my newly created [[ENRICH-1114- Smart Query Prompt Updates]] subtasks ❌ 2025-05-16
    - https://qualtrics.atlassian.net/browse/ENRICH-1463 is first of three
- [-] Clarify guardrail work scope for [[ENRICH-1468- Smart query version 2 in designer (epic)]] ❌ 2025-05-16
    - Remove/disable harmful content guardrail
    - Implement prompt injection guardrail?
    - Error UI?
- [-] How can I help support the team more? ❌ 2025-05-16
    - Acknowledge that he has a lot on his plate, and a lot of new, unfamiliar things from engage side. Also dealing with having team members pulled off for OXM, Kuan leaving. Lots of slack messages?
    - Say, with [[ENRICH-1468- Smart query version 2 in designer (epic)]]
        - Engaging more with product directly instead of looping him in?
- [-] Discuss [[2025 Objective Plan ❌ 2025-05-16

# Meeting notes
- I apologized. I said the ways I'm working on it are
    - regulating my emotions better
    - communicating more to find miscommunications earlier
    - self-checks on my understanding (nebulous)
- [[ENRICH-1464- Find/load test data for behavior/perf evaluation]]
    - Vamsi asks if I could ask QE to find and load up these test data
    - [x] Ask steven if he can help find test data for 1464 ✅ 2025-05-16
    - Vamsi says there are tools to publish data into IG. "test generation tools"
    - Can work this in parallel
- [x] Can we define a measure of success for delivering [[ENRICH-1468- Smart query version 2 in designer (epic)]] for coordinating with product ✅ 2025-05-16
- [x] Make a jira task for disabling/removing harmful content guardrail ✅ 2025-05-16
- [x] Review these notes, make more tasks if needed ✅ 2025-05-16
- He says there's more broader eng-required guardrails than just prompt injection
    - There are eng standards about guardrails
    - [x] Look into whether we'll need more guardrails than just prompt injection. The [coda page](https://coda.io/d/Foundations-UX-DICE_da8UAJU2QcS/Guardrails_subzNRwD#_luGwHOsp) ✅ 2025-05-19
- [x] I had to leave early. Schedule a follow-up to discuss more of this agenda ⏫ 📅 2025-05-15 ✅ 2025-05-16

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/Meeting notes/Vamsi 1 on 1/2025-05-19 Mon - Vamsi 1 on 1.md
---

# My plan
- [x] Mention my accommodation request for remote work is in review with Sedgwick #meeting/vamsi-1-on-1 ✅ 2025-05-19
- [x] Mention that I feel I have the schedule flexibility I need and that I really appreciate it. #meeting/vamsi-1-on-1 ✅ 2025-05-19
- [x] Ensure the scope of my Q2 deliverable "Smart query version 2 in designer" is entirely covered by [[ENRICH-1468- Smart query version 2 in designer (epic)]] #meeting/vamsi-1-on-1 ✅ 2025-05-19
- [-] Confirm he approves of the measure of success "80% reduction while maintaining relevance and coverage" #meeting/vamsi-1-on-1 ❌ 2025-05-19
- [x] Clarify guardrail work scope for [[ENRICH-1468- Smart query version 2 in designer (epic)]] #meeting/vamsi-1-on-1 ✅ 2025-05-19
    - Remove/disable harmful content guardrail
    - Implement prompt injection (PI) guardrail?
    - Error UI?
    - He mentioned there are more guardrails we may need in [[2025-05-15 Thu - Vamsi 1 on 1]]. Dig into that. Max tokens? I asked about "post GA scope" for that.
        - Donovan says token limit guardrail is "work planned for SQ in Q2" [slack](https://qualtrics.slack.com/archives/C07MKL3FB61/p1747666274506559?thread_ts=1747159409.784589&cid=C07MKL3FB61)
- [-] How can I help support the team more? #meeting/vamsi-1-on-1 ❌ 2025-05-19
    - Acknowledge that he has a lot on his plate, and a lot of new, unfamiliar things from engage side. Also dealing with having team members pulled off for OXM, Kuan leaving. Lots of slack messages?
    - Maybe just let that sit for a bit--you've got a lot going on. And let him suggest ways I can help. Is there anything I can do to make your job easier?
    - Say, with [[ENRICH-1468- Smart query version 2 in designer (epic)]]
        - Engaging more with product directly instead of looping him in?
- [-] Discuss [[2025 Objective Plan #meeting/vamsi-1-on-1 ❌ 2025-05-19

# Meeting notes
- [x] Review these notes and create tasks if needed ✅ 2025-05-19
- I told him my accommodation request for remote work is submitted to sedgwick and I'm waiting on that. He said in the meantime my attendance is still tracked.
- I thanked him for the schedule flexibility I have. He said to ensure business needs are met with my flexible schedule.
- [x] rearrange the [[ENRICH-1468- Smart query version 2 in designer (epic)]] epics so they make sense. epic can't parent epic. 🛫 2025-05-21 ✅ 2025-05-21
- Asking if PI guardrail is needed in Q2
    - Addisu's thread about how to do this, seems like socrates/CME can't help us?
    - If it's not feasible with current CME API, we can push back on timeline. We might have to wait on them.
    - I should drive this, don't rely on Vamsi.
- Max token guardrail is needed in Q2
- I confirmed that the scope of the Q2 deliverable [[ENRICH-1468- Smart query version 2 in designer (epic)]]
    - Prompt updates
    - Remove HC guardrail
    - Add PI, MT guardrails
    - Error UI for guardrails
- Status tracking
    - I asked what he meant by "status tracking". He said the words below. I don't really get it. I think he wants me to be more autonomous.
        - execution, decisions during implementation, if product team has questions during our status tracking meetings or during testing if QE found issues we didn't anticipate (like a UX issue). he wants me to feel empowered to make technical decisions on how to handle that.
    - As we're implementing prompt tuning or PI guardrail if we find out that there's a feasibility issue, like if CME team can come up with better API where they can understand AWS bedrock tokens in text.
    - He wants me to feel empowered to make decisions or suggest changes.
- Vamsi had to go after 20 minutes. I said I've got what I need and the other item I had were background and can wait.
# After meeting
- I felt there were a few miscommunications. I feel I was patient and understanding.
    - I asked a specific question about "prompt injection guardrail" and he started explaining that there were different kinds of guardrails and prompt injection is one of them.

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/Meeting notes/Vamsi 1 on 1/2025-05-29 Thu - Vamsi 1 on 1.md
---

---
creation date: 2025-05-29T13:56:35
tags:
  - meeting/vamsi-1-on-1
---
# My agenda
## Potential Discussion Topics
- **Smart Audit Variance Testing:**
    - Confirm if Vamsi is aligned with this reprioritization over reducing wildcard use in Smart Query.
- **Smart Query Guardrails & Scope:**
    - Briefly update him on the status of the Prompt Injection guardrail ([[ENRICH-1487- Add prompt injection guardrail to Smart Query]]), noting the successful test with the standalone API and the MR being out for review. You could mention the LIG response code observation if it's still relevant.
- **Classification Service Deployment [[ENRICH-1440- New Sharding Implementation: Deploy changes across discover DCs]]
    - Provide a status update: Prod CA, EU, and APSE deployments are done. Prod US is on hold pending the scale test (#op-7858) and RMQ UI credentials (which you've now received).
    - Reiterate that you'll notify Unified Orchestration before the Prod US deployment.
- **Workload & Prioritization:**
    - Given the new assignment to Smart Audit variance testing, discuss how this impacts other Q2 deliverables for Smart Query (e.g., prompt updates - [[ENRICH-1463- Iterate on prompt to reduce wildcard use]]).
- **(Optional) Blocklist Item ([[ENRICH-1158- Enhance the Engine enrichment service to support skipping messages by configuration.]]):**
    - You had a follow-up to check on this being assigned/prioritized. If it's still relevant, you could ask for an update.

## Quick Review of Recent Items:

- **ENRICH-1487 (Prompt Injection Guardrail):** MR sent, addressing feedback. You were investigating LIG response codes.
- **AI tools for Query-Based Topics Meeting:** You were on track for mid-June code complete for SQ, but then reassigned to Smart Audit variance testing.

# Meeting notes
- [x] Review these notes and create tasks if needed ✅ 2025-05-30
- Vamsi suggests I do interview training... if I want. Surprised it's optional
- Noted accommodation request, thanked him for poking me on it
- He's open to me working on other stuff besides variance testing if other work suits me better
    - SA and TR are most important for getting to GA in Q2. We can sacrifice SQ if needed.
    - [x] I'll see the scope of the variance testing work at the meeting tomorrow and see if Steven can handle it himself. I may be able to work on something else. ✅ 2025-05-30
- Mentioned SQ PI guardrail work going well, investigating funny behavior in upstream (LIG) code, which is good learning
- Mentioned CS prod US deployment still blocked on scale test
    - He asked if I knew who was doing scale testing
    - Vamsi's going to track down who's doing the scale testing ^ee48ba
- Asked about prioritization of [[ENRICH-1158- Enhance the Engine enrichment service to support skipping messages by configuration.]]
    - We found it was created Feb 16, so it's been over two months, which I guess is the RCA AI timeline.
    - Vamsi is going to come up with a justification for request an extension
    - He says scheduler library deals with the rabbit headers, might be useful info for the fix here. Scheduler library deals with the header that lists the services a message goes to and updating that list and the next service header
    - [x] Try to put in a few minutes on blocklist change every now and then 🆔 9lla51 🛫 2025-06-03 📅 2025-06-18 ✅ 2025-06-05
- Objective plan ^2fcf3a
    - SQ v2 can be an objective
    - Gaining cross-knowledge between discover and engage (was an objective in Q1)
        - I got on-cal/hands-on in Q2 for both systems, so can be for Q1 and Q2 objective
        - Learn tech stack: k8s, CMP, golang
    - [x] Tell Vamsi when I've added these objectives ✅ 2025-05-30
- UI tech lead (liaison?) for all engage: Jason Nelson is the initial candidate. pick standards, frameworks. give guidance. component designs. best practices, which libraries.
    - Vamsi says it's an opportunity for me to do this in the future (Q4, next year)
    - He's only mentioning this because of my interest in UI
    - Vamsi's going to send a slide from the town hall meeting on what this role entails
# After meeting


---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/Meeting notes/Vamsi 1 on 1/2025-06-09 Mon - Vamsi 1 on 1.md
---

---
creation date: 2025-05-30T15:24:03
tags:
  - meeting/vamsi-1-on-1
---
# My agenda
- [x] Note that I opted in to a Feedback 360 #meeting/vamsi-1-on-1 ✅ 2025-06-09
- [-] Note that I updated my Q3 days out in the planning coda to 16 #meeting/vamsi-1-on-1 ❌ 2025-06-09
- [x] Confirm prioritization: #meeting/vamsi-1-on-1 ✅ 2025-06-09
    - Complete [[ENRICH-1158- Enhance the Engine enrichment service to support skipping messages by configuration.]]
        - There's deployment, dlq reprocessing verification, setting up alerts, verifying alerts
        - Deployment will block at Prod US if we're still waiting for scale testing before deploying sharding there
    - [[ENRICH-1547- Update MIG Issuer for Topic Recommender]]
        - Need to get PLMS updated to allow new issuer, commit code changes, verify, deploy
    - [[ENRICH-1463- Iterate on prompt to reduce wildcard use]]
        - Begin
- [x] Review [[2025 Objective Plan]] #meeting/vamsi-1-on-1 ✅ 2025-06-09 ^340b4b
    - Note that I'm aiming for exceeds/L6
        - (after: didn't)
    - (after: only got through first objective, still need to do others)
- [-] Overdue items #meeting/vamsi-1-on-1 ❌ 2025-06-09
    - *I've noticed a few items that become urgent as their due date approaches. When I see something like that in the backlog and tag you, but don't hear back, what's the best next step for me to take?*
    - [[2025-06-09 Mon#^db6ed8]]
- [-] Regularly scheduling recurring meetings #meeting/vamsi-1-on-1 ❌ 2025-06-09
    - *I've noticed our async scrums and sprint planning meetings can be a bit ad-hoc. It would be helpful for my planning to have more predictable, recurring calendar items for these. What are your thoughts on that?*

# Meeting notes
- [x] Make sure notifications go out to feedback 360 evaluators. Vamsi says he didn't get a notification ✅ 2025-06-10
- If I'm spending too much effort on [[ENRICH-1547- Update MIG Issuer for Topic Recommender]], "raise my hand" and Vamsi may deprioritize it
- Vamsi wants [[ENRICH-1158- Enhance the Engine enrichment service to support skipping messages by configuration.]] done ahead of the due date so we aren't stressing.
    - He says I have the autonomy to decide how to prioritize it. The strong end message is "don't miss the due date again".
- Vamsi wants me to investigate using a feature flag to enable/disable sharding in CS
    - (in response to deployment of 1158 blocking on sharding scale testing for prod US)
    - This is preferable to reverting the code
- "Engage" Steven and Zeshan about scale testing
    - Vamsi thinks Steven has bandwidth this week to do the scale testing
    - Vamsi sent Steven a message, I'm following the thread
- Review objective plan
    - Objective *Deliver High-Impact AI Feature Enhancements and Critical System Stability*
        - He wants an objective for "system stability/ops". He thinks it makes more sense to have AI features and system stability separate.
            - results of cross-knowledge of discover space
            - ktlo improvements like RCA AI [[ENRICH-1158- Enhance the Engine enrichment service to support skipping messages by configuration.]]
            - efforts toward reducing pulses, operational load
        - writing design docs KR
            - he's happy to support me in this
            - i can look through Q3/Q4 or beyond to see where i can make a design impact
            - look in the PPR. in the team-level planning coda ([this](https://coda.io/d/PXE-Hub_dCG_1QsXKay/Team-Planning_sucrg821?utm_source=slack#Team-Planning-Deliverables_tuA2DLvI/r27738)? he linked it in chat), there's also a page on org-goals/features that i can look at.
            - he can think off the top of his head: Discover+Engage unification, CMP deprecation, kafka migration
            - he suggests attending Topics Design Discussion meeting on Thursdays with my designs and proposals
# After meeting
- [x] Review these notes and create tasks if needed #meeting/vamsi-1-on-1 🛫 2025-06-05 ✅ 2025-06-09
- Will focus on getting [[ENRICH-1158- Enhance the Engine enrichment service to support skipping messages by configuration.]] done
- [x] Investigate gating sharding behind a feature flag (is it much different from having an empty shard list in config?) 🔼 ✅ 2025-06-09
    - Created [[ENRICH-1568- Classification Service- Add sharding migration disabler]]
- [ ] Update [[2025 Objective Plan]] with Vamsi's feedback 🔽

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/Meeting notes/Vamsi 1 on 1/2025-06-12 Thu - Vamsi 1 on 1.md
---

---
created: 2025-06-10T10:47:57
tags:
  - meeting/vamsi-1-on-1
---

# My agenda
- [-] Review more of [[2025 Objective Plan]]? #meeting/vamsi-1-on-1 

# Meeting notes
- I'll become secondary on-call after my shift ends next Wed
- Vamsi reiterated that he thinks we should not be doing widget work as big as [[TAPP-7074- TextIQ Bubble charts content is not perceptible and Bubble charts cannot be operated via keyboard]]
- [x] Move next 1-1 ✅ 2025-06-12
# After meeting
- [x] Review these notes and create tasks if needed #meeting/vamsi-1-on-1 ✅ 2025-06-12
- 

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/Meeting notes/Vamsi 1 on 1/2025-06-26 Thu - Vamsi 1 on 1.md
---

---
created: 2025-06-16T07:52:09
tags:
  - meeting/vamsi-1-on-1
---
# My agenda
- [x] Appreciations: #meeting/vamsi-1-on-1 ✅ 2025-06-26
    - Updating the calendar items
- [x] Discuss priorities for this sprint, what was loaded for me yesterday while I was out #meeting/vamsi-1-on-1 ✅ 2025-06-26
    - Bring up [[ENRICH-1597- Fannie Mae - Organization Hierarchy Volume Discrepancy - Categorized Records are not displaying in Studio]]
        - Due date
    - [[ENRICH-1630- Spike- Investigate and Plan Smart Query Prompt Updates to Reduce Wildcards]]
    - [[ENRICH-1396- Unit Tests for Auto Tune]]
        - saw this was loaded to this sprint
    - in meeting: confirmed this priority list
- [x] Getting Vamsi's attention #meeting/vamsi-1-on-1 ✅ 2025-06-26 ^375200
    - *I'd like to find the most effective way to bring items to your attention. I've found that when I tag you directly on Jira or Slack, my messages can sometimes get buried, and I need to remember to follow up.*
      *To make things more efficient for both of us, I was wondering what system would work best for you. For example, would you prefer if I compiled non-urgent items into a single daily message? I want to make sure I can get the input I need to stay unblocked, while also finding a method that fits well into your workflow.*
    - (I suggested in the meeting:) Send one message at the end of the day with all non-urgent items?
    - in meeting, he says: just ping him with a DM if it seems to have fallen through the cracks ^3a7818
- [-] Ask if time spent on pulses as on-call is counted, then how about as not on-call. #meeting/vamsi-1-on-1 ❌ 2025-06-26
    - I think there's an inconsistency here. Why might we want to condition *counting time spent on a given type of work* on the on-call status of the person working?
- [-] Ask Vamsi to clarify whether secondary on-call works exclusively on on-call stuff (including ktlo backlog), or if they're only pulled in as needed by the primary on-call. #meeting/vamsi-1-on-1 ❌ 2025-06-26

# Meeting notes
- "anything after operations/ktlo priority, [[ENRICH-1468- Smart query version 2 in designer (epic)]] is next"
- Vamsi's going to assign [[ENRICH-1396- Unit Tests for Auto Tune]] to charmish
- Vamsi wants to review https://qualtrics.atlassian.net/browse/ENRICH-1337 and https://qualtrics.atlassian.net/browse/ENRICH-1288. they came up at sprint planning yesterday and team wanted my input? He just wanted me to estimate them.
- he asked about my comfort level with discover systems, given i've been on-call for a quarter
    - he asked about rtds (vs IG), reclassification, our elasticsearch plugin (part of classification workflow)
        - these are core functionalities, will be helpful for migrating customers out of RTDS and decommissioning features from CMP
            - vamsi: *if there are any operational or ktlo things like analyzing pulse or deploying some change, it will be helpful to volunteer to reverse shadow it*
                - *keep an eye on pulse board, vulns, wherever i see an opportunity to work on those areas, volunteer to shadow/reverse shadow.*
            - i agreed, saying hands-on is better for me than watching KT
    - he says there's KT on RTDS
- he's trying to note the impact i have, my contributions, thinking of Q2 as our first full quarter ourselves
    - he says i'm not meeting L5 expectations ^83590d
        - "what is the level of value to the team/org?"
        - delivery and impact
        - he says "we need to work on upping that going forward"
- [x] he wants me to look at the Q3 planning link (coda board) ✅ 2025-06-26
    - if i have feedback on the Q3 plan, send it to vamsi
        - if the line items make sense or need to change anything, he'd like to hear my feedback or suggestions

# After meeting
- [x] Review these notes and create tasks if needed ✅ 2025-06-26
- i'm proud of how i conducted myself in this meeting
    - i raised [[#^375200]] gently and focused on us against the problem, not fault-finding
    - i stayed curious and passive during the "missed L5 expectations" chat. i feel strong defensiveness, but it's under control. i knew that this isn't the time to dig in, while i'm focusing on staying regulated. i can cool off (i danced right afterward) and come back with a clear mind to think about it later.

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/Meeting notes/Vamsi 1 on 1/Next - Vamsi 1 on 1.md
---

---
created: 2025-06-26T14:23:36
tags:
  - meeting/vamsi-1-on-1
---
# My agenda
- [ ] Follow up on not meeting L5 expectations in Q2 [[2025-06-26 Thu - Vamsi 1 on 1#^83590d]] #meeting/vamsi-1-on-1 

# Meeting notes
- 

# After meeting
- [ ] Review these notes and create tasks if needed #meeting/vamsi-1-on-1 
- 

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/People/Me.md
---

I'm Janzen Brewer-Krebs, a senior software engineer (L5) at Qualtrics. I started in Dec 2021 on Text iQ and was there until Jan 2025, when Text iQ ownership was split and I went to the former discover engines team. We're now "query topics" team.

My slack handle is @jbk

I am autistic and have predominantly inattentive ADHD. I'm transfem, amab.
# Lessons learned
![[ENRICH-1158- Enhance the Engine enrichment service to support skipping messages by configuration.#How to not cause an OP next time]]

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/People/Vamsi.md
---

# FAQ
## Vamsi hasn't responded to a slack thread I tagged him in
He says he may miss his mentions sometimes, even though he tries to get all of them. I can DM him to follow up.
![[2025-06-26 Thu - Vamsi 1 on 1#^3a7818]]
# Jira preferences
## 2025-05-01
I asked what kind of jira item he likes for breaking down a story into smaller pieces. He said task. https://qualtrics.slack.com/archives/C087VC5JPD3/p1746113645094179?thread_ts=1746108289.841349&cid=C087VC5JPD3

## 2025-05-16
Vamsi preferred me to convert [[ENRICH-1114- Smart Query Prompt Updates]] to an epic from a story, and the subtasks it had as a story into stories. Seems like subtasks should be for very small pieces of work. They seem to inherit the sprint of their parent and it can't be changed.

## 2025-05-20
After converting [[ENRICH-1114- Smart Query Prompt Updates]] to an epic, I found I couldn't parent it to the other epic [[ENRICH-1468- Smart query version 2 in designer (epic)]]. I suggested moving the stories in 1114 into 1468 (also asked if there's a better item type than epic for the Q2 deliverable represented by 1468). Vamsi said to move the stories to 1468.

## 2025-06-17
I guess put ktlo and ops in sprints 🤷
### Source
[slack](https://qualtrics.slack.com/archives/D087U08TEAX/p1750175745169739)
Janzen Brewer-Krebs
Second--would you please remind me of which kinds of work go in our sprints? Not pulses? What about ktlo and ops? I'll write your preferences down this time :smile:

Vamsi Vippadapu
not sure if I follow your question on sprints. The KTLO changes such as action items from RCA will be visible in sprint board and we should attach them to ktlo epic of the quarter
other ops tasks such rundeck job are being tracked sprint too (there is one assigned to Reid for example.)

Janzen Brewer-Krebs
Oh yes, I wasn't clear. For example, I just did this ktlo task: https://qualtrics.atlassian.net/browse/ENRICH-1591. Should that item's sprint field in jira have Sprint 108 in it or not?

ENRICH-1591 Update crate image with vuln fixes
Status: Closed
Type: Task
Assignee: Janzen Brewer-Krebs
Priority: Major
More actions...
Added by Jira Cloud
New

Vamsi Vippadapu
yeah feel free to include this ktlo item too

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/Status/2025-03-19.md
---

* On-call
	* Investigate gov1 sentiment backup https://qualtrics.slack.com/archives/C036JC1JKGE/p1742403973454289
		* Trying to access gov1 cluster
		* sectool gov1 aws console kerfuffle (fixed with reboot)
		* Set up engage cluster access (thanks Kuan!)
	* Triage a few transferred TAPP pulses
* [[Tickets/ENRICH-1269- Discrepancy between TextIQ and sentiment score]]
	* Explain SFP boundaries
	* Discuss solutions in slack thread https://qualtrics.slack.com/archives/C018K0H3FJB/p1742327326566809
* Start [[ENRICH-1266- 2024 Complete R&D Tax Credit Study (Janzen)]] in background
* [[ENRICH-1148- Enable autoscaling for engines services in discover non-US clusters]]
	* File ticket for access to non-US vault (maybe? I took the survey but can’t find the ticket and [no one’s responded to me in \#aws-operations](https://qualtrics.slack.com/archives/C02G9BZ48ER/p1742406935505509))
	* Ask team to check whether the secret I’m investigating is present in prod4/anz https://qualtrics.slack.com/archives/C087VC5JPD3/p1742417155124679?thread_ts=1742402382.615669&cid=C087VC5JPD3
* Try out Obsidian for notekeeping
* Meetings/other
	* Dr appt

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/Status/2025-03-20.md
---

* On-call
	* Look into *sentiment service 0 consumers issue* as an exercise
	* Trying to access rabbitmq console
		* Asked to be added to the relevant group in aws-operations
* Tidy up and share aws-k8s-login script https://gitlab-app.eng.qops.net/-/snippets/9653
* Meetings/other
	* Enrichments all-hands
	* 1:1
* [[ENRICH-1148- Enable autoscaling for engines services in discover non-US clusters]]
	* Verified anz has the admin_keda secret
	* Asked about missing `serviceAccount` again, thanks Kuan and Milo!
	* Updated MR with anz configs, un-drafted, asked for re-review
* Filed https://qualtrics.atlassian.net/browse/DISCCLOUD-31777 to get `cb-plat-dev-us-developer` role access
* Filed https://qualtrics.atlassian.net/browse/DISCCLOUD-31776 for access to rabbitmq management consoles
* Asking aws-operations about access to all prod vault paths (I seem to only have prod-us) https://qualtrics.slack.com/archives/C02G9BZ48ER/p1742505434576589?thread_ts=1742406935.505509&cid=C02G9BZ48ER
* Requested permission to view our jira board https://qualtrics.slack.com/archives/CH97FPGQ4/p1742508611976989

# Log
* [x] Watch k8s autoscaling training https://drive.google.com/drive/folders/1HAn5FJ5xHE0qJdC8D-0tKH-mf-NVhRUB 🔽 ✅ 2025-05-27
* [x] Triage this ticket Addisu sent me https://qualtrics.atlassian.net/browse/ENRICH-1275 ✅ 2025-03-21
* [x] Do [[ENRICH-1266- 2024 Complete R&D Tax Credit Study (Janzen)]] 🔽 📅 2025-04-17 ✅ 2025-04-04
* [x] go through list of commits in pxe hub and tell vamsi if there are any i want to work on, or feel match my expertise 🔼 ✅ 2025-03-21
* [ ] learn rtds, arcadia ⏬
* [x] cc vamsi on my blocker threads, send them to him ✅ 2025-03-21
* [x] find grafana dashboards (i asked in slack, look for a thread with francisco) (vamsi said he'd send me a page on splunk dashboards) ✅ 2025-03-26

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/Status/2025-03-21.md
---

* On-call
	* Reading/handling a flurry of slack messages
	* Followed up on triaging out ENRICH-1274, 1272 https://qualtrics.slack.com/archives/C018K0H3FJB/p1742421718022199
		* Also TAPP-6977 https://qualtrics.slack.com/archives/C018K0H3FJB/p1742574467391549
	* [[Tickets/ENRICH-1269- Discrepancy between TextIQ and sentiment score]]
		* In ops we decided to close this pulse and file [[ENRICH-1283- Update survey flow plugin to use sentiment v3 API]] to follow up.
		* Closed as works as designed
	* [[Tickets/ENRICH-1275- recommended topics not appearing for a field in textIQ even though there are 500+ responses in a single language]]
		* Repro'd, checked for obvious UI issues, looks like a torq thing.
		* Asked torq on-call to look https://qualtrics.slack.com/archives/C8X9E12QY/p1742591743664109
	* [[Tickets/ENRICH-1273- Text iQ is inaccessible from survey]]
		* Repro'd. Found the 503 failing call, tracked it down in splunk, asked torq to look.
	* [[Tickets/ENRICH-1271- TextiQ Sentiment Categorization Inaccuracy]]
		* Asked coreml to look
	* [[ENRICH-1270- Text iQ keeps loading [503 status] for a user]]
		* Asked res how to proxy log in with just a user's email. I've never done it.
	* Triaged https://qualtrics.atlassian.net/browse/ENRICH-1268 back to TAPP.
	* Verified I can access rabbitmq management console https://qualtrics.atlassian.net/browse/DISCCLOUD-31776
* Other/meetings
	* Scrum
	* Ops
	* Monthly team pulse survey
* [[ENRICH-1148- Enable autoscaling for engines services in discover non-US clusters]]
	* Merged MR
	* Core enrichments owns this now. They should deploy the autoscaling change.

# Log
- [x] see if i can do the IS UI accessibility ticket https://qualtrics.atlassian.net/browse/ENRICH-743 🔽 ✅ 2025-04-07
	- vamsi says if it's too much work to get familiar with this codebase i don't have to
- [x] file ticket for team to update base image for vulns ⏫ , do it if i have time while on-call ✅ 2025-03-26
	- template service?
	- arcadia, and the ones that's 10 days to sla
	- https://vmx.eng.qops.net/dashboard/17?affected_sha=&image_full_name=prod-registry.g1-cmh.qops.net%3A5001%2Fdiscover-aws%2Fdiscover-engines%2Ftemplate-service%3A3.0.16-cbc9ec7&team_name=Discover+Engines+Team

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/Status/2025-03-24.md
---

* On-call
	* Investigating classification service queue backup in fra1
		* https://qualtrics.slack.com/archives/C053MK619T9/p1742824802799939
		* I see classification service is at 27/27 pods. Asking for help.
		* Francisco says it seems fine because there are no more alerts, but can't tell me how often they'd fire.
	* Investigating model service alert
		* Kuan and Francisco handled it. Something about an unescaped colon in a rule. I'm trying to understand what happened.
	* [[ENRICH-1270- Text iQ keeps loading [503 status] for a user]]
		* Seeing if I can impersonate users. Yep.
		* Repro'd, digging into logs
		* Asked \#dataplatformsupport for help https://qualtrics.slack.com/archives/C03A90TCJ/p1742834842799039
		* They say the user probably doesn't have permissions on the survey. Investigating...
		* Yep. Res gave me two users to try and they have different issues. The relevant user hits a different issue. I'm asking torq to investigate a 403 from ASP. https://qualtrics.slack.com/archives/C8X9E12QY/p1742857689623559
	* [x] fieldset encryption key alert 🔺 2398268 ✅ 2025-03-24
		* Text tools says iad1 crate isn't healthy enough to delete, investigating
		* iad1 crate has 185,292 underreplicated records... waiting to see if it progresses
		* Ran `ALTER CLUSTER REROUTE RETRY FAILED` 
		* Deleted fieldset from crate, enqueued an fssync job for the fieldset
	* [[Tickets/ENRICH-1273- Text iQ is inaccessible from survey]]
		* Poked Hai to move the ticket
	* [x] poke coreml tickets to move pulses to their project ✅ 2025-03-24
	* Asked core enrichments to take ENRICH-1284
	* [[ENRICH-1148- Enable autoscaling for engines services in discover non-US clusters]]
		* Informed core enrichments of this work
		* Closed
* Other/meetings
	* Scrum
	* Proposal Review: Improve scalability of the Classification Service
	* PDF Discussion with Rohit

# Log

- [x] see if 1284 is only sentiment, or only sentiment and classification. if sentiment only, send to core enrichments ✅ 2025-03-24

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/Status/2025-03-25.md
---

* On-call
	* Created scheduled message automation to remind engines-batman to start TAM release weekly
	* Started TAM release
		* Poking it along: finding verifiers, question about handover
	* Moved triaged pulses to other jira projects
	* Ran brand deletion job
	* [x] [[ENRICH-1270- Text iQ keeps loading [503 status] for a user]] ✅ 2025-03-25
		* Working with torq on-call to debug permissions
		* Editing permissions for Mohsen
		* I missed the ticketing sources--adding permissions
		* I don't understand ticketing permissions--asking for help https://qualtrics.slack.com/archives/C8X9E12QY/p1742926177236289?thread_ts=1742857689.623559&cid=C8X9E12QY
		* Mike says the user has to be brand admin for ticketing in this case https://qualtrics.slack.com/archives/C8X9E12QY/p1742928634729339?thread_ts=1742857689.623559&cid=C8X9E12QY
		* Filed follow-ups:
			* ENRICH-1288: TAM UI should show useful information when `derived-from` calls return 500
			* ENRICH-1289: TAM server doesn't handle 403 from `derived-from` calls
		* Reverting my exploratory permissions changes
		* Closed works as designed
	* [x] fieldset encryption key alert ⏫ ✅ 2025-03-25
		* Same ccb fieldset as yesterday `bb1c1d4e-2a5e-cb3b-e026-86089b392e2e`... sus
		* Deleted fieldset from crate, enqueued sync job
	* [x] [[ENRICH-1285- Altice- Derived Attribute not working]] 📅 2025-04-23 ✅ 2025-03-27
		* Paired with Francisco
		* Asking for help with kibana logs https://qualtrics.slack.com/archives/C08KD3LEYBE/p1742929877008429
	* [x] [[Tickets/ENRICH-1275- recommended topics not appearing for a field in textIQ even though there are 500+ responses in a single language]] 📅 2025-04-16 ✅ 2025-04-08
		* Hai found the trex aren't generating, which is apparently a DFS thing. Investigating...
		* This is our problem
	* Alert for classification service queue length in fra1, didn't refire after 15 minutes, pods were scaling up
* Other/meetings
	* Dr appt
	* Scrum
	* Pair session with Francisco for ENRICH-1285

# Log
- [x] figure out difference between cb designer and studio ⏬ ✅ 2025-03-27
	- 2025-03-27 studio is way newer. designer is the old gwt thing. not sure the difference in functionality

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/Status/2025-03-26.md
---

* On-call
	* Filed https://qualtrics.atlassian.net/browse/ENRICH-1290, informed next on-call
	* [[Tickets/ENRICH-1275- recommended topics not appearing for a field in textIQ even though there are 500+ responses in a single language]]
		* Looking for DFS logs indicating error
		* Asking for help https://qualtrics.slack.com/archives/C8X9E12QY/p1743021199548629?thread_ts=1742591743.664109&cid=C8X9E12QY
		* Torq sent a debugging doc for DFS https://docs.google.com/document/d/1A6wM1rmEsSW8SWL8Ji72axpG8nqR2nUow6PBDWpLU-c/edit?tab=t.0
		* Investigating if coreml is generating trex
		* Taking time to set up mongodb compass again
			* sectool web ui can't talk to my local sectool process? restarting laptop... that fixed it
		* Reaching out to coreml to check s3 file; seems only they have permission
		* Coreml investigated, found trex output file is empty, then found trex input file is empty. I think it's back to me to investigate; I'm learning DFS as I go on this.
	* Responded to a res question about the resolution of [[Tickets/ENRICH-1269- Discrepancy between TextIQ and sentiment score]]
* Other/meetings
	* WLD Keynote
	* Scrum
	* Dr appt

# Log
- [x] use the 1/4 day vacation i abandoned today on something else ⏬ ✅ 2025-04-01
- [x] figure out when i'm on call next. VO is accurate ✅ 2025-03-26

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/Status/2025-03-27.md
---

* Yelled at jetbrains channel for taking away my ultimate license starting tomorrow
* Other/meetings
	* ENRICH-1285 DB Fix Discussion
	* Scrum
	* EE/ME Q1 scores and action plan
* On-call
	* Deployed TAM to prod
		* Investigate gov1 deployment failure
			* Just missed the manual approval. Works now.
	* [[ENRICH-1285- Altice- Derived Attribute not working]]
		* Shadowed Kuan and Gerardo finding CDAs in DB, fixing
		* Discussed model vs node CDAs
		* Closed
		* Repeating Kuan's investigation for my own learning
			* Tripped up by fancy quotes in kibana search

# Log
- [x] preserve the cda zoom recording 🔼 ✅ 2025-03-27
- Vamsi said he's aware "learning new domains" will spill over into Q2 in EE/ME meeting. Just noting so I have paper trail in case he reneges.

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/Status/2025-03-28.md
---

* Vacation/dr appt
* Scrum

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/Status/2025-03-29.md
---

* Meetings/other
	* Q1 enrichments all-hands
	* Vacation
	* Dr appt

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/Status/2025-04-01.md
---

* Meetings/other
	* Team repositories, operational boundaries walk through
	* Sprint review
	* Dr appt
* On-call
	* [[ENRICH-1270- Text iQ keeps loading [503 status] for a user]]
		* Followed up on res question, request for docs in [jira](https://qualtrics.atlassian.net/browse/ENRICH-1270?focusedCommentId=11215954)
	* [[ENRICH-1275- recommended topics not appearing for a field in textIQ even though there are 500+ responses in a single language]]
		* Asked res if they know/can find out why the customer has recodes for userLanguage [jira](https://qualtrics.atlassian.net/browse/ENRICH-1275?focusedCommentId=11216106)
* Prepare sprint review slides
	* Tried to get gemini to generate my bullet points. Meh
* [[Tickets/ENRICH-743- Accessibility Reqs]]
	* Finding docs on IS UI... [found](https://coda.io/d/_dt-izYOVJmq/Intelligent-Scoring-UI_sunzAFOy), and there's a readme in the `ui` dir
	* Hit issue with build
		* `401 Unauthorized - GET https://gitlab-app.eng.qops.net/api/v4/projects/41031/repository/archive.tar.gz?sha=gitlab-migration`
		* That project is react-studio-console... I'm a member of it
		* Dug around, found .npmrc assumes GITLAB_TOKEN env var is present. Set up a new PAT, works
	* Another build issue
		* `npm WARN tar zlib: incorrect data check`
		* `npm WARN tarball tarball data for @discover/react-studio-console@https://gitlab-app.eng.qops.net/api/v4/projects/41031/repository/archive.tar.gz?sha=gitlab-migration (sha512-2z...) seems to be corrupted. Trying again.`
		* ? https://qualtrics.slack.com/archives/C07T9DGJHTN/p1734718248341839?thread_ts=1734630680.178639&cid=C07T9DGJHTN

# Log
- [x] Figure out if TSE is used [coda](https://coda.io/d/_dt-izYOVJmq?encodedLocationContext=eyJjb21tZW50VGhyZWFkSWQiOiJyLWxmYlpKYkpWVzgiLCJwYWdlSWQiOiJzZWN0aW9uLXluQkFuM3RlME0iLCJ0eXBlIjoiZ3JpZCIsIm9iamVjdElkIjoiZ3JpZC0yRGp5ZHBJRmV6Iiwicm93SWQiOiJpLUtRWlhNSE0yWi0ifQ) 🔽 ✅ 2025-04-04

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/Status/2025-04-02.md
---

* Meetings/other
	* Team charter and roadmap
	* (Weekly) Enrichment design and architecture discussions
	* Chat with Sabina about [[ENRICH-743- Accessibility Reqs]]
	* 2x Dr appts
* Asked Daniel about TSE [slack](https://qualtrics.slack.com/archives/C8X9E12QY/p1743603667086199)
	* He says it's unused and out of prod

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/Status/2025-04-03.md
---

* Meetings/other
	* 1:1 (Vamsi cancelled late)
* Reviewed company quarterly EX results, plans
* [[ENRICH-1305- Queue depth based alerts in GOV1 and all engage regions, review alert grouping]]
	* Created jira item, asked for more details
	* Got on a call with Addisu and he walked me through the similar work he did for fra1
	* Researched existing metrics, proposed using AOM instead of cron jobs [qualtrics.slack.com](https://qualtrics.slack.com/archives/C087VC5JPD3/p1743700165410199?thread_ts=1743697014.538189&cid=C087VC5JPD3)
* [[ENRICH-743- Accessibility Reqs]]
	* Asked if *I* should really be doing this, *now* [qualtrics.slack.com](https://qualtrics.slack.com/archives/C087VC5JPD3/p1743703703245569)

# Log
- [x] create jira items for *Queue depth based alerts in GOV1 and all engage regions, review alert grouping* [coda.io](https://coda.io/d/PXE-Hub_dCG_1QsXKay/Team-Planning_sucrg821) ✅ 2025-04-03
- [x] [[ENRICH-1305- Queue depth based alerts in GOV1 and all engage regions, review alert grouping]] ✅ 2025-04-09

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/Status/2025-04-04.md
---

* Meetings/other
	* Scrum
* [[ENRICH-1266- 2024 Complete R&D Tax Credit Study (Janzen)]]
	* Working... completed
* [[ENRICH-1309- Sentiment Score Incorrect]]
	* Was asked to review this pulse since it involves survey flow plugin
	* Sentiment for an obviously negative text was coming back positive
	* I sent the text directly to sentiment api and got positive response. Suggested sending it to coreml.
* Asked Kuan if I can start work on "smart query version 2 in designer"
	* He says we're waiting for product input [qualtrics.slack.com](https://qualtrics.slack.com/archives/C036JC1JKGE/p1743786031481419)
* Volunteered to take [[ENRICH-1158- Enhance the Engine enrichment service to support skipping messages by configuration.]]
	* Kuan gave me an overview of the problem

# Log
- [x] [[ENRICH-1158- Enhance the Engine enrichment service to support skipping messages by configuration.]] 🆔 1158 🔼 ✅ 2025-05-08
- I think Vamsi said to go ahead with my AOM approach to [[ENRICH-1305- Queue depth based alerts in GOV1 and all engage regions, review alert grouping]] when I brought up that I was waiting for feedback in scrum

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/Status/2025-04-07 Mon.md
---

* Meetings/other
	* 1:1
	* Scrum
* [[ENRICH-1305- Queue depth based alerts in GOV1 and all engage regions, review alert grouping]]
	* Wild goose chase figuring out why only 4 engage DCs report metrics--because discover is only in 4 engage DCs
	* Writing AOM alert definition, reading a lot
	* Created grafana dashboard with queue length, pod count
	* Sent AOM MR

# Log
* 1:1
	* [ ] learn kafka (used in reclassification workflow, temporarily stores document that needs to be retagged. with the reorg we own kafka support) ⏬ 
	* said i'm concerned about smart query v2 key commit needing product input at this point. [[ENRICH-1114- Smart Query Prompt Updates]]
		* he says top pri for this work is supporting asterisk wildcard. other features include guardrails, ci/cd (for lambda functions?)
		* he says i'm the primary contributor for this item, kuan is supporting me
* Vamsi said I shouldn't do [[ENRICH-743- Accessibility Reqs]] [qualtrics.slack.com](https://qualtrics.slack.com/archives/D087U08TEAX/p1744046821734699)
* told team i'm trans [qualtrics.slack.com](https://qualtrics.slack.com/archives/C087VC5JPD3/p1744066054830679)

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/Status/2025-04-08 Tue.md
---

* Meetings/other
	* Scrum
* Closed as duplicate [[ENRICH-1275- recommended topics not appearing for a field in textIQ even though there are 500+ responses in a single language]]
* [[ENRICH-1305- Queue depth based alerts in GOV1 and all engage regions, review alert grouping]]
	* Addressed MR comments, merged
	* Asked for clarification on scope [qualtrics.slack.com](https://qualtrics.slack.com/archives/C087VC5JPD3/p1744125834359419?thread_ts=1743697014.538189&cid=C087VC5JPD3)
	* Created gov1 grafana dashboard, added link to same to alert
	* Shared work with core enrichments
	* Waiting to hear what "alert grouping" is before closing ticket
		* It's not sending duplicate alerts when the alarm condition is triggered for longer than the check interval. We get this for free with AOM.
	* Closed
* [[ENRICH-1158- Enhance the Engine enrichment service to support skipping messages by configuration.]]
	* Start work
	* Stumbled through go install, path issue
	* Wrote a first draft implementation

# Log
* 

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/Status/2025-04-09 Wed.md
---

* Meetings/other
	* Scrum
* [[ENRICH-1158- Enhance the Engine enrichment service to support skipping messages by configuration.]]
	* Kuan walked me through testing my code deployed to dev
	* Changing code to read IDs from header, before protobuf unmarshalling
	* Bringing code up to production quality
	* Started toying with deploying my new code in dev cluster
* Helped Mavin with aws-mfa setup

# Log
* i may work on topic translator in Q2:
	* [qualtrics.slack.com](https://qualtrics.slack.com/archives/C08M249TH6K/p1744220539231419)
Vamsi VippadapuVamsi Vippadapu![:no_entry:](https://a.slack-edge.com/production-standard-emoji-assets/14.0/apple-large/26d4.png)  [10:42 AM](https://qualtrics.slack.com/archives/C08M249TH6K/p1744220539231419)  

hi Charmish, I communicated to Janzen about potential time spent on topic translator during Q2. We have some bandwidth but not enough to fully implement the feature. Given we have design is ready, let us keep this as stretch line item.

Charmish MojidraCharmish Mojidra  [10:42 AM](https://qualtrics.slack.com/archives/C08M249TH6K/p1744220561132609)  

Sure, makes sense

[10:43](https://qualtrics.slack.com/archives/C08M249TH6K/p1744220595960949)

Charmish Mojidra

We can look into it if we get enough time

Vamsi VippadapuVamsi Vippadapu![:no_entry:](https://a.slack-edge.com/production-standard-emoji-assets/14.0/apple-large/26d4.png)  [10:45 AM](https://qualtrics.slack.com/archives/C08M249TH6K/p1744220723146299)  

Implementing the PCO goals is higher priority compared to translator for clarity

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/Status/2025-04-10 Thu.md
---

* Meetings/other
	* Enrichments Monthly All Hands
* Looking at [qualtrics.atlassian.net](https://qualtrics.atlassian.net/browse/ENRICH-1335) for Charmish
	* Charmish asked me how I ran sentiment directly on text. I was going to write instructions, but decided to add the functionality to response tools (admin-tools)
	* Trying to get admin-tools to build
	* Asked for help from dpsupport
* [[ENRICH-1158- Enhance the Engine enrichment service to support skipping messages by configuration.]]
	* 
* Code review [gitlab-app.eng.qops.net](https://gitlab-app.eng.qops.net/xm-discover/discover-platform-group/discover-data-pipelines/platform/-/merge_requests/6245/diffs)
	* Cloned platform repo, set up cursor for java

# Log
* 

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/Status/2025-04-11 Fri.md
---

* Meetings/other
	* none ✨
* Adding sentiment calls to text tools (in admin-tools)
	* Banged my head on the build again
	* Figured it out
	* Implementing
	* Sent MR [gitlab-app.eng.qops.net](https://gitlab-app.eng.qops.net/data-platform/admin-tools/-/merge_requests/1160)

# Log
* 

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/Status/2025-04-14 Mon.md
---

* Meetings/other
	* Scrum
* Adding sentiment calls to text tools (in admin-tools)
	* Address MR comments
	* Sent [sentiment-api MR](https://gitlab-app.eng.qops.net/coreml-model-engineering/sentiment-api/-/merge_requests/352) to allow the admin-tools jwt issuer, merged
	* Figuring out how to hotwire ingestion-scheduler to enqueue messages
* Compiled on-call notes for Reid

# Log
* 

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/Status/2025-04-15 Tue.md
---

* Meetings/other
	* Sprint review
* Fixed issues with the AOM queue alert I recently created [qualtrics.slack.com](https://qualtrics.slack.com/archives/C053MK619T9/p1744735419781809?thread_ts=1744723937.903559&cid=C053MK619T9)
* [[ENRICH-1158- Enhance the Engine enrichment service to support skipping messages by configuration.]]
	* Fixed header parsing, my code now allows through unblocked messages
	* Learning k9s
	* Added more tests, logs, bulletproofing
	* Testing manually in dev cluster
	* Investigating integration and E2E tests
	* Figuring out code coverage (main is already below the threshold the MR template gives)
* Trialing tunnelblick to replace AWS VPN client

# Log
* 

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/Status/2025-04-16 Wed.md
---

* Meetings/other
	* 
* [[ENRICH-1158- Enhance the Engine enrichment service to support skipping messages by configuration.]]
	* Trying to get integration tests to run
	* Learning about port forwarding
	* Debugging .env files
	* Got some tests running
* Pulled into gov1 vuln debacle
	* Working on fssync, overtime

# Log
* 

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/Status/2025-04-21 Mon.md
---

* Meetings/other
	* 1:1
	* Dr appt
* Can't log in to AWS VPN, investigating...
	* 🙄🙄🙄 [coda.io](https://coda.io/d/Discover-AWS-CloudOPS_dPxRM-VAWne/SOP-How-to-reset-AWS-Active-Directory-Password-AWS-VPN_suxo6WHx)
* [[ENRICH-1158- Enhance the Engine enrichment service to support skipping messages by configuration.]]
	* Trying to see my integration test doing anything to classification service in staging
	* Got my new test to publish messages that classification service processes
	* Trying to set up blocking and monitor dlq

# Log
* 1:1
	* [[ENRICH-1114- Smart Query Prompt Updates]]
		* vamsi says donovan wants telemetry for stuff like how many times users are triggering them, how often we hit llm quotas. investigate more.
		* [x] add this to page ^ ✅ 2025-04-22
	* [x] "gatling tool"? he brought it up in the context of [[ENRICH-1158- Enhance the Engine enrichment service to support skipping messages by configuration.]] ✅ 2025-04-22
		* he wants me to get help from Zeshan and go faster

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/Status/2025-04-22 Tue.md
---

* Meetings/other
	* Scrum
	* Knowledge sharing - Engine Controller and RabbitMQ topology of the Ingestion Pipeline
	* Dr appt
* [[ENRICH-1158- Enhance the Engine enrichment service to support skipping messages by configuration.]]
	* Found that I need a real project ID for engines controller to forward my messages to classification service
		* Asking about options [qualtrics.slack.com](https://qualtrics.slack.com/archives/C087VC5JPD3/p1745343006125409?thread_ts=1745267491.794059&cid=C087VC5JPD3)
	* Hacked publishing directly to classification service into integration tests
		* I don't *love* this because I'm setting up engines headers manually and I'd kinda like to let engine controller do that (I assume that's where it happens)
	* Brought in Zeshan, he says test data is just created by hand 😱. It is at least tracked [here](https://coda.io/d/Discover-Engines-Group_dUq7lnCJ_pd/Integration-Tests-Configuration-Engines_suYrq723#Copy-of-Table-28_tuaHpQGD)

# Log
* 

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/Status/2025-04-23 Wed.md
---

* Meetings/other
	* Scrum
	* 2x Dr appts
* [[ENRICH-1158- Enhance the Engine enrichment service to support skipping messages by configuration.]]
	* Learned that designer has its own login page
	* Compiled the frustrations I started sharing in scrum, [posted](https://qualtrics.slack.com/archives/C087VC5JPD3/p1745429968301789?thread_ts=1745427604.849889&cid=C087VC5JPD3) as Vamsi requested
* On-call
	* Checked pulse board, see 5 unassigned pulses from previous shifts, [asking about policy](https://qualtrics.slack.com/archives/C087VC5JPD3/p1745425310099649)
	* TAM deployment
		* Addisu is reverse shadowing
		* Release bot isn't working...
			* [x] investigate release bot failure 🔼 ✅ 2025-04-28
	* Alert: syd1 crate gc excessive
		* Checking [fssync grafana dashboard](https://grafana.eng.qops.net/d/aabbccDD789/fssync-queue?orgId=1&from=now-1h&to=now&timezone=browser&var-dc=syd1&var-fqdn=$__all&refresh=1m). Load is elevated on the alerted node, but seems okay?
		* Fixed worker count widgets in prod and gov1 grafanas
	* [x] ENRICH-1377- [GOV1] CMS - CDA classification failing 🆔 1377 ⏫ ✅ 2025-05-01
		* Res asked me to look at logs in gov1
	* Alert: Classification queue length > 150000 in ingestion-pipeline-0
		* [I can't log in to cb grafana...](https://qualtrics.slack.com/archives/C053MK619T9/p1745450783384249?thread_ts=1745450139.726819&cid=C053MK619T9)
		* resolved itself
		* [x] see if i need a cloud ticket for grafana access 🔼 ✅ 2025-04-24

# Log
* [x] put my 1158 frustrations in the standup parking lot thread ✅ 2025-04-23

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/Status/2025-04-24 Thu.md
---

* Meetings/other
    * Working Session for ENRICH-1377-CMS
	* Retro
	* AI tools for Query-Based Topics
* On-call
    * Alert: Classification service errors
    	* splunk shows errors like:
    	* error classifying document: error verifying index: error creating index vanilla-1443704-c64398f0710b2fac8f4d993bd4e3102fb420ae84a16eb8ad57c6e35f7feb9bc4: error creating index: dial tcp [::1]:9200: connect: connection refused
    	* Milo says try restarting pods
    	* Restarted pods, seems better
    	* Alert refiring a couple hours later, asking for help
    * Fruitlessly dug through Discover AWS ops survey for grafana access, [asked in channel](https://qualtrics.slack.com/archives/C02G9BZ48ER/p1745508379511789)
    	* Got access, asked them to verify that other discover newbs have access too
    * [x] update sentiment service alert config to point to core enrichments ✅ 2025-04-25
    	* [Asked for channel](https://qualtrics.slack.com/archives/C053MK619T9/p1745517962501879?thread_ts=1745499795.087039&cid=C053MK619T9)
    * [x] make discover queue alerts less sensitive ✅ 2025-04-24
    	* Changed pending period from 15 to 30 minutes
    * Added to common alerts: *ATTENTION NEEDED: Classification service error!*
    * [[ENRICH-1377- [GOV1] CMS - CDA classification failing]]
    	* Investigating with Addisu, Raj
    	* Raj says this is urgent. [Asking Vamsi...](https://qualtrics.slack.com/archives/C08PH4PKSRY/p1745522277118049) he says business hours
    	* Searching for relevant logs in gov1 splunk
    	* Digging to figure out kafka restart
    	* Rolling restarting msk brokers
    	* Kicked off new classification run

# Log
* AI tools for Query-Based Topics
	* Donovan wanted to write dates down. I said I'll start work Apr 30, after on-call. With 6 weeks scoped (the original planning doc said 9), so that would end Jun 11. I pointed out that I'm on-call Jun 4-11, so should push that date out a week. Donovan said something about LRR needing to happen this quarter and basically ignored my point.

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/Status/2025-04-25 Fri.md
---

* Meetings/other
	* Scrum
	* Ops
* Requested accommodation to not have required in-office days
* On-call
	* [x] investigate TAM b1 canary failures ✅ 2025-04-25
		* [Asked Reid if it's his change](https://qualtrics.slack.com/archives/C010USY4MPX/p1745596680990679?thread_ts=1745528589.653739&cid=C010USY4MPX)
		* Reid is handling this
	* [x] investigate TE 5xx alerts ✅ 2025-04-25
		* Looks like fra1 crate hiccups. We should consider expanding the cluster.
	* Pinged core enrichments about an alert channel [qualtrics.slack.com](https://qualtrics.slack.com/archives/C087V8K2GFP/p1745597083747769)
		* Moved alerts for sentiment, derived attribute, and IS services
	* [x] ENRICH-1375- [CY]Classifications failing on Chime Financial ✅ 2025-04-30
		* Taking a look. Seems similar to [[ENRICH-1377- [GOV1] CMS - CDA classification failing]].
		* Res asked for a kafka restart in the ticket. I can do that now!
		* Figured out copper is in AWS US
	* [[ENRICH-1377- [GOV1] CMS - CDA classification failing]]
		* The classification run started yesterday after kafka restart failed
		* Asked team for help
		* Kuan found a kafka timeout
		* I paged dxm data services on-call for help
			* They pointed me to core data, who came in quickly
		* Getting msk logs
			* Created s3 bucket jbrewer-enrich-1377-logs
				* [x] delete s3 bucket jbrewer-enrich-1377-logs ⛔ 1377 ✅ 2025-04-29
				* [x] turn off msk logging ⛔ 1377 ✅ 2025-04-28
		* Renewing power user accesses
		* Fetched cmp logs from splunk for Jakub
	* Followed up with Shawn about TAM release (release bot is broken) [qualtrics.slack.com](https://qualtrics.slack.com/archives/C010USY4MPX/p1745600173432019?thread_ts=1745338016.199679&cid=C010USY4MPX)
		* He said it can wait till next week
		* [x] start release thread or manual release shawn's change 📅 2025-04-28 ✅ 2025-04-28
	* Triaged as ours: *TAPP-7056 Text IQ topics disappearing after being added*
	* Triaged this widget issue to TAPP: *ENRICH-1340 Text IQ Topics Disappeared from Dashboard*
	* Updated Common Alerts with *Text Engine Records 5xx Anomaly*

# Log
* [x] do k8s gameday [coda.io](https://coda.io/d/Resiliency-Gamedays-at-Qualtrics_d1FCcb0U-Vh/Gameday-Tracker_suyOuJLU#Kubernetes-Gameday-Tracker_tuardGv7) ⏫ 🛫 2025-04-28 ✅ 2025-05-05

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/Status/2025-04-28 Mon.md
---

* Meetings/other
	* 
* On-call
    * [x] see if we have a team role for aws, add our stuff to it so i don't have to get power user every time ✅ 2025-04-30
        * kafka:ListClustersV2
    * ENRICH-1377- [GOV1] CMS - CDA classification failing
        * Get power user again
        * Checking EBS
        * Jakub walked me through kafka terraform config and the jenkins deployment job for gov1, says we need to raise storage before they reshard
        * Disabled logging for MSK
        * Reviewing terraform plan output for updating MSK and saw `configuration_info` block is dropped, which is worrying. There's no pointer to the active configuration in terraform afaik. I'm waiting for Jakub's input.
        * Added configuration_info to terraform config
        * Tried to terraform apply the new config but hit an IAM issue. Jenkins user doesn't have a permission on the msk resource.
        * Turns out terraform *was* able to apply the config. Seems like it couldn't describe the cluster afterward. Asking aws-operations for help with jenkins user permissions [qualtrics.slack.com](https://qualtrics.slack.com/archives/C02G9BZ48ER/p1745873584264559)
    * [x] check on scheduled crate deployment jobs for fra1, yul1 [qualtrics.slack.com](https://qualtrics.slack.com/archives/C019UF4CE90/p1745688794603189) ✅ 2025-04-30
    * Asked team about which DB model service had issues with [qualtrics.slack.com](https://qualtrics.slack.com/archives/C087VC5JPD3/p1745865761647989)
    * [x] Investigate overnight classification service error alerts ⏫ ✅ 2025-04-30
    * [[ENRICH-1375- [CY]Classifications failing on Chime Financial]]
        * Trying to find kafka in prod US
        * Eh, don't have permission and I think this isn't msk kafka
        * Asked Jakub about ec2 kafka restart and the weird config errors in the pulse
    * Started gov1 crate host patching https://rundeck-data-platform.gov1.qprod.net/execution/show/3383798#nodes
    * TAM release bot failure
        * Resuming investigation
        * Created debug build to see the relevant response
        * Ran into a qdev error when trying to build the image
            * Fixed with docker update, restart
        * Deployed debug build 0ccef06
        * The code wasn't checking for unexpected response status code. If it had, we would've seen the underlying issue: there's already an open MR.
        * Added error handling to respond to the slack thread with error information instead of silently failing. Improved log format. Added more bot messages/feedback and a help command.
        * When there's nothing to release, the bot was leaving the empty MR open and that was gunking things up. I added code to delete the MR in that case.
    * [x] verify gov1 crate deployment 📅 2025-04-29 ✅ 2025-04-29
    * [x] complete TAM deployment 📅 2025-05-01 ✅ 2025-05-01

# Log
* [x] update calendar for may ✅ 2025-04-29
* i need to run the deployment job in jenkins. linking [[Kafka]]
    * tf
        * the tf config is in a weird ticket name branch
        * ebs_volume_size. fra1 has an override to 2000
        * copy that override to gov1 tfvars. value is 300. jakub isn't sure about this value
        * use billing team discover-data-services
        * for action use plan first
        * plan failed because we didn't specify security group sg_name
            * msk-sg, found from [us-gov-west-1.console.amazonaws-us-gov.com/vpcconsole/home?region=us-gov-west-1#SecurityGroup:groupId=sg-01a6c90aca040a839](https://us-gov-west-1.console.amazonaws-us-gov.com/vpcconsole/home?region=us-gov-west-1#SecurityGroup:groupId=sg-01a6c90aca040a839)
        * if plan works, run job with apply
            * when running, you have to go inside the job (console output) and verify to apply
        * wait for jakub to say if cluster would restart before applying

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/Status/2025-04-29 Tue.md
---

* Meetings/other
    * Sprint review
    * WLD talk
* Updated calendar for May
* [x] Compile slides for sprint review ✅ 2025-04-29
* Created sprint review slides for team
* Found I can't get into bitwarden (as usual), requested access
* [x] Load sprint tickets ✅ 2025-04-29
* On-call
    * Lots of overnight alerts in the channel. My phone didn't notify me even though I've configured slack to notify for every message in the alert channel. I worry we're going to miss things because of this.
    * [x] investigate TAM b1 deployment failure spam ✅ 2025-04-30
        * Chris found k8s TAM in b1, asking Reid about the status of k8s migration
        * TAM is still on nomad
        * Mavin fixed it
    * [x] investigate high cpu usage alerts for rds [qualtrics.slack.com](https://qualtrics.slack.com/archives/C053MK619T9/p1745933543946369?thread_ts=1745933430.161709&cid=C053MK619T9) ✅ 2025-04-30
        * Addisu picked this up. They're reviewing whether the db was downsized for cost savings.
    * [x] investigate crate high swap activity, cpu usage alerts in iad1 ✅ 2025-04-30
    * [x] investigate hung crate host patch job in gov1 https://rundeck-data-platform.gov1.qprod.net/execution/show/3383798#output ✅ 2025-05-01
    * Asked for help with https://qualtrics.atlassian.net/browse/ENRICH-1375 chime
        * Vamsi wants me to work this pulse and give the secondary (Addisu) other tasks instead
    * Ran brand deletion
    * [[ENRICH-1375- [CY]Classifications failing on Chime Financial]]
        * Searching splunk for relevant logs
        * Found the config "errors" are actually warn-level. Ignoring them.
        * Found some logs suggesting issues with "propagation", sounds like kafka
        * Created a slack channel for this pulse, posted my findings, asking Kuan and Jakub to look
        * Addisu found it's an ES error. Asking data services to take it.
    * [x] Model, classification service splunk alerts (10:00) ✅ 2025-04-30
        * Investigating
        * Classification service: `error fetching project with structure only: error executing http request: invalid request parameters : received status code 400`
        * Found this post from Kuan with great instructions https://qualtrics.slack.com/archives/C053MK619T9/p1742516270416399?thread_ts=1742504533.434299&cid=C053MK619T9
        * Asked for help with verifying https://qualtrics.slack.com/archives/C087VC5JPD3/p1745947175146309
            * Milo says the evidence is strong, probably don't need to verify
        * Asked res to force republish the model https://qualtrics.slack.com/archives/C080ZKTJM9A/p1745948751713769
            * Interestingly, this is the same project that had trouble last time
        * Added section on this issue to Common Alerts
        * Issue stopped after force republish
    * Merged text-release-bot changes to main
        * [x] Build and deploy this from main ✅ 2025-04-30

# Log
* Chatted with Evan about discomfort with manager
    * He says he has my back and will not allow retaliation if I give more direct feedback to v
    * very supportive
* During sprint review, I wrote and said that I'm concerned about Donovan's insistence that my deliverable be done this quarter. Vamsi said it had been moved to July release, so less time pressure. I'm relieved.
* vamsi misgendered me with pronouns during sprint review. i didn't say anything.
* told reid over slack that it felt good to read him refer to me as "she" yesterday in a random slack message

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/Status/2025-04-30 Wed.md
---

* Meetings/other
	* 2x Dr appt
	* TextiQ widgets discussion
* Lots of slack to catch up on
* On-call (spillover)
    * [[ENRICH-1375- [CY]Classifications failing on Chime Financial]]
        * Followed up about ownership, trying to send it to core data https://qualtrics.slack.com/archives/C08PA03MQM9/p1746049636856949
    * [[ENRICH-1377- [GOV1] CMS - CDA classification failing]]
        * Escalated to P0
        * Others are running a reshard
        * I'm trying to make sure ownership is clear and that we have the Krakow people we need since they have Thursday and Friday off. https://qualtrics.slack.com/archives/C08PH4PKSRY/p1746050279820839?thread_ts=1746046426.022659&cid=C08PH4PKSRY
        * Reshard failed. It's not on us to fix this.
    * Investigating inconsistent crate versions in prod https://qualtrics.slack.com/archives/C019UF4CE90/p1745688794603189
        * Found yul1 scheduled deployment job has been stuck since Mar 13
        * All nodes have version dbms4.0.12.34
        * Looking for other stuck jobs...
            * Killed jobs in fra1, iad1
        * [x] kick off yul1, fra1, iad1 crate deployments 🛫 2025-05-01 ✅ 2025-05-01
        * Triggered yul1 crate deployment
            * failed because the image path changed and the desired version tag isn't present in the new path (wants .35, but only .36 is in query-topics/crate)
            * Chris released .36 to b1 yesterday
            * [x] release crate .36 or copy the .35 image to the new path so deployments can resume 🛫 2025-05-01 ✅ 2025-05-01
    * Poked TAM release verifiers
    * Investigating stuck gov1 crate host patch job
        * Can't get webapp cred because I'm locked out of bitwarden. Asking for help https://qualtrics.slack.com/archives/C052JG61BK4/p1746062442076839
    * Rebuilt and deployed text-release-bot from main after merging my update
    * k8s gameday
        * Added link to common k8s runbook to team hub
        * Found service in g1 with autoscaling enabled that's suitable for gameday
        * `keda-hpa-scale-plat-classification-service` is broken? Can't read metrics?
        * [x] investigate broken g1 classification service hpa ✅ 2025-05-01

# Log
* [x] post on call log to coda ✅ 2025-04-30

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/Status/2025-05-01 Thu.md
---

* Meetings/other
	* Tax credit interview
	* 1:1
	* Dr appt
* Called Sedgwick to start accommodation claim
* On-call spillover
    * Catching up on yul1 crate node 10's texteng-dbms container not running
        * Trying to get qdev publish to retag images
        * It can't do prod to prod tagging? https://qualtrics.slack.com/archives/C04E138ARS7/p1736899828205769?thread_ts=1736540346.331439&cid=C04E138ARS7
        * Trying using the staging image, I get farther but hit permission issue
            * `Unable to authorize by intention, missing gitlab claims`
        * Chris hit merge conflicts merging crate main to release and asked for help, investigating
            * I resolved the conflicts and pushed the merge directly to release
        * I think it's that the billing team changed and I'm not in the query-topics ldap group
            * Trying to get access with conductorone
            * I don't see query-topics in the conductorone UI, asking for help https://qualtrics.slack.com/archives/C07V6M2CCQ2/p1746114984419289
            * Asked Vamsi to make me a primary team member in odo hoping that that would give me the ldap group
            * Temporary access request also fails because query-topics isn't available to select
            * Asked an existing primary team member to check groups and found they don't have query-topics either
        * Given the trouble with ldap groups and the new billing team is blocking crate deployments, I'm going to revert the hiera change so we can deploy .35 again.
            * Waiting on a devops SME
        * I really want to get yul1 node 10 back online
        * Found the local puppet config for texteng-dbms in `/etc/dockerctl.d/texteng-dbms`, changed the image path back, was able to update-restart texteng-dbms. Cluster shows green with node 10 online.
    * Added sschwartz to text-release-bot QE list, added more logging
        * Deployed
    * Deleted, recreated bitwarden account
        * My vaults are all gone. Rob Olsen says it'll take a minute to populate.
    * Created jira tickets for my on-call spillover tasks as Vamsi directed
        * https://qualtrics.atlassian.net/browse/ENRICH-1408
        * https://qualtrics.atlassian.net/browse/ENRICH-1409
        * https://qualtrics.atlassian.net/browse/ENRICH-1410
        * https://qualtrics.atlassian.net/browse/ENRICH-1411
        * https://qualtrics.atlassian.net/browse/ENRICH-1412
        * https://qualtrics.atlassian.net/browse/ENRICH-1413

# Log
* 

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/Status/2025-05-02 Fri.md
---

* Meetings/other
	* Scrum
	* Ops
* On-call spillover
    * k8s gameday
        * Ran into permission issue trying to disable autoscaling for CS in g1
            * Milo says it's based on ldap groups, I guess for the namespace? Anyway I don't have discover-engines
            * [x] Get discover-engines ldap group ✅ 2025-05-02
            * Asked iam about this https://qualtrics.slack.com/archives/C01SLQL16P6/p1746205817653229
            * The odo-based permissioning system is deprecated in favor of conductorone. I pointed out that the odo page for ldap doesn't say that. They're going to update that page.
            * Requested discover-engines ldap group via conductorone
        * Switched to fssync, was able to disable autoscaling
        * Manually scaled to 3 replicas then back to 1
        * Re-enabled autoscaling
        * Verified relevant logs show in splunk
        * Creating a new route file for query-topics in prometheus rules repo
            * [x] Check on the [prometheus rules MR](https://gitlab-app.eng.qops.net/engvis/prometheus/rules/-/merge_requests/699). Readme says they should merge it. 🛫 2025-05-06 ✅ 2025-05-05
        * Verifying that one of our services uses the qualtrics app template helm chart. I expected to see a Chart.yaml file in fssync's helm config dir but don't.
            * Found it in [spinnaker](https://periscope.eng.qops.net/catalog/default/component/k8s-text-application-text-engine-fssync-k8/spinnakerConfig)
* [[ENRICH-1158- Enhance the Engine enrichment service to support skipping messages by configuration.]]
    * Polished code, sent MR for classification service
* Helped Reid with discover k8s because he doesn't have access

# Log
* [x] file follow-up task for [[ENRICH-1158- Enhance the Engine enrichment service to support skipping messages by configuration.]] integration tests 🆔 1tob33 ✅ 2025-05-05
* Standup
    * I said "a gentle reminder that my pronouns are she/her"
    * Kuan's leaving. We're reviewing on-call schedule. I'll move up a week
        * [x] Log my on-call overrides for May 28 shift once the schedule has Kuan removed 🔽 📅 2025-05-06 ✅ 2025-05-06
        * Mavin asked Vamsi about getting another on-call since we're dropping from 6 to 5 https://qualtrics.slack.com/archives/C087VC5JPD3/p1746209061281269
    * [-] Vamsi says I should deploy [[ENRICH-1158- Enhance the Engine enrichment service to support skipping messages by configuration.]] to prod once it's merged and tested ⛔ 1158 📅 2025-05-13 ❌ 2025-05-13
* Ops
    * I saw that the on-call time tracker had pain level 6 for week ending Apr 28. I reported pain 8, others reported fewer hours at lower pain. I questioned whether mean is the best measure for this and proposed max instead. Mavin sent in chat that it was a good callout and he looks forward to seeing it discussed (?) at Enrichments level

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/Status/2025-05-05 Mon.md
---

* Meetings/other
	* Scrum
* Noted that crate backups are failing, asked on-call to investigate or I can file a ticket https://qualtrics.slack.com/archives/C087VC5JPD3/p1746457133273999
    * [-] File crate backup failure ticket if needed 🛫 2025-05-06 ❌ 2025-05-06
* Noted that crate g1 deployment seems stuck. It may be that I reverted the hiera container path to text-application but the release version is .36 now, which I think is only in query-topics. Told on-call (Reid)
    * https://qualtrics.slack.com/archives/C010USY4MPX/p1746457415109839?thread_ts=1746288012.748759&cid=C010USY4MPX
    * Providing guidance to Reid
    * Paired with Reid to get crate g1 node 7 back online
* On-call spillover
    * k8s gameday
        * My rules update was approved, I merged it https://gitlab-app.eng.qops.net/engvis/prometheus/rules/-/merge_requests/699
        * Watched gameday kickoff recording
        * Filed tickets for each on-call eligible team member to complete the relevant training
        * Done!
* [[ENRICH-1158- Enhance the Engine enrichment service to support skipping messages by configuration.]]
    * Reviewing MR feedback
    * Asked Kuan to review. I want a review from him before he leaves.
    * Filed https://qualtrics.atlassian.net/browse/ENRICH-1426 as a follow-up task to write integration or E2E tests
    * Made suggested changes, pushed commit
* [[ENRICH-1295- Designer Error Handling for smart-query guardrails]]
    * Starting this task
    * I need to get cmp running first? I still have the docs I was following open.
    * Refreshing myself on what I've already done in the build steps
    * Gradle build completed
    * Completed "Set up" section
    * CMP runs! A couple test data uploads succeeded and I see data in the UI.
    * Found smart query public docs, walking through...
    * Hit a server error when generating smart query. It doesn't have some config? Asking for help https://qualtrics.slack.com/archives/C087VC5JPD3/p1746478157136529
    * Smart query works
# Log
* 

# Start of day tasks
* [x] Copy meeting/other schedule to notes ✅ 2025-05-05
* [x] Check email ✅ 2025-05-05
* [x] Check slack ✅ 2025-05-05

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/Status/2025-05-06 Tue.md
---

* Meetings/other
	* Scrum
	* Review the Classification Sharding implementation
	* Dr appt
* Doing on-call hygiene
    * Asked Vamsi to remove Kuan from on-call rotation
    * Adding my overrides for next shift
* [[ENRICH-1295- Designer Error Handling for smart-query guardrails]]
    * Trying to figure out how to trigger the guardrails
        * Reading this spreadsheet https://docs.google.com/spreadsheets/d/1x-llNalPvUgrYbTffIw1LkyycBOd53ndyqmcX662994/edit?gid=0#gid=0
        * No help, there appear to be no blocks
    * Asked team for help triggering guardrails
    * Call with Addisu
        * [x] get socrates-console-access-role 🛫 2025-05-08 ✅ 2025-05-08
            * Search for that role in slack to figure out how to get it. Charmish got access a few weeks ago.
            * Try socrates-ml-platform in slack
            * secops/aws-policies is the repo where role access is defined
            * (later) the role name is wrong? maybe it's just a display name. i can't find it in aws-policies repo
    * Addisu said Donovan may remove the guardrail feature, so pause this until the AI sync meeting on Thursday
* [[ENRICH-1158- Enhance the Engine enrichment service to support skipping messages by configuration.]]
    * Multiple rounds of addressing MR feedback
    * Looking for classification service deployment docs
        * Asked team if https://coda.io/d/_dt-izYOVJmq/Classification-Service_suAa9o2- is up to date
        * Then asked about bake time and deployment verification testing
    * Handled merge conflict, waiting on pipeline
    * Vamsi says other teams can handle adding this functionality to the other IG services. I only need to handle classification service.
    * Merged
    * Kuan says he just merged something that needs more testing and that he should've waited for me to merge mine since I want to deploy. I said we could revert his change and tagged Vamsi for his preference https://qualtrics.slack.com/archives/C087VC5JPD3/p1746550662004149?thread_ts=1746547999.202399&cid=C087VC5JPD3
        * We talked about this in standup and I feel we went in circles, like I wasn't conveying something. I was trying to push for Vamsi to make the prioritization choice. Kuan and Mavin were invested in my change not being deployed.
    * I misunderstood the priority on this. I pushed out the due date to allow Kuan time to test his change, then deploy together.
* [x] [[ENRICH-1405- Spike- Understand project and break into smaller tasks]] ✅ 2025-05-15
    * 

# Log
* Responded to Vamsi poking me about the due date on 1158. He says we can pick our due dates for action items but have to stick to them.
    * Sent Vamsi a slack message about how I'd misunderstood his messages on this and thought it needed to be done asap:
Janzen Brewer-KrebsJanzen Brewer-Krebs  [12:41 PM](https://qualtrics.slack.com/archives/D087U08TEAX/p1746560512474229)  

Hi Vamsi, I hope the offsite is going well. I wanted to take a moment to share some feedback on our conversation this morning and explore how we can improve our interactions moving forward.First, I appreciate the flexibility you mentioned around due dates. For clarity, could we possibly set clearer expectations in future communications? When I read your message about the due date having passed, I interpreted it as needing immediate resolution, thus feeling an urgency to deploy quickly.Given that text communication can sometimes lack context, especially with my being autistic, I want to find ways to ensure better mutual understanding. It's important to me to align on expectations and manage tasks without misunderstanding urgency levels.Thank you for your understanding and support.

Vamsi VippadapuVamsi Vippadapu  [12:56 PM](https://qualtrics.slack.com/archives/D087U08TEAX/p1746561366493529)  

hey Janzen, sure. I can add more context related AI. Are you familiar with RCA process and expectations around AI implementation and closure of sev2s?

Janzen Brewer-KrebsJanzen Brewer-Krebs  [12:56 PM](https://qualtrics.slack.com/archives/D087U08TEAX/p1746561392128689)  

I've gone through it in the past, but it's not something I keep in my head.

Vamsi VippadapuVamsi Vippadapu  [1:13 PM](https://qualtrics.slack.com/archives/D087U08TEAX/p1746562389157479)  

it is part of operational standard. We can go over brief refresher if that helps. At a high level, AI s are created as part of RCA for incidents, some maybe added based on recommendations during review. The due dates are set by team with reasonable time frame based on capacity and complexity (less than 90 days target from incident to close all AIs). The AIs get tracked at org level and engineering wide to make sure they closed according to plan.

Vamsi VippadapuVamsi Vippadapu  [1:23 PM](https://qualtrics.slack.com/archives/D087U08TEAX/p1746563001178869)  

the goal is to not miss the due dates which we assigned to ourselves as part of RCA

Janzen Brewer-KrebsJanzen Brewer-Krebs  [1:27 PM](https://qualtrics.slack.com/archives/D087U08TEAX/p1746563233673469)  

Thank you for explaining. I understand why the due date on these AIs is important now. I'll try to do better in the future.

Vamsi VippadapuVamsi Vippadapu  [1:28 PM](https://qualtrics.slack.com/archives/D087U08TEAX/p1746563310069999)  

the urgency was that we went past due date. Also it showed up at engineering level review as past due

[1:29](https://qualtrics.slack.com/archives/D087U08TEAX/p1746563379893439)

Vamsi Vippadapu

same goes with pulses, we don't want to miss SLA. If there is justification for priority change or due date change, we should pursue that in collaboration with Res team.

# Start of day tasks
* [x] Copy meeting/other schedule to notes ✅ 2025-05-06
* [x] Check email ✅ 2025-05-06
* [x] Check slack ✅ 2025-05-06

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/Status/2025-05-07 Wed.md
---

* Meetings/other
	* Scrum
	* Dr appt
* [[ENRICH-1405- Spike- Understand project and break into smaller tasks]]
    * Addisu says it's a prompt update (or we could "clean the stars in designer", which I don't really get, but he says it's not preferred)
    * Looking for the relevant code
        * This is the repo for the lambdas https://gitlab-app.eng.qops.net/xm-discover/discover-engines-group/query-topics-ai-integrations
    * Requested role access for lambdas from socrates team https://qualtrics.slack.com/archives/C0186S22PQS/p1746632150715629
        * Also PLMS 1p vault https://qualtrics.slack.com/archives/C0186S22PQS/p1746635730480289
        * I got the PLMS jwt signing key
    * Where are the prompts? They come from PLMS? That's prompt-lifecycle-management-service
        * https://gitlab-app.eng.qops.net/iq-research-sea/prompt-lifecycle-management-service
        * Yes, but where in plms? I guess the prompts aren't source controlled with the service code.
        * The lambda requests a prompt named `sq-gen-prompt-dict-0`. Sourcegraph search only turns up the lambda use of that string.
        * Asked for help https://qualtrics.slack.com/archives/C07MKL3FB61/p1746635272993699
        * Addisu says prompts are uploaded to PLMS. PLMS docs say they're replicated out to all DCs from g1.
    * I'm getting timeouts trying to access PLMS in g1
        * Asked socrates on-call for help https://qualtrics.slack.com/archives/C0186S22PQS/p1746636628407509
        * I can't hit internal-lb for g1 from my local machine. Only b1.
    * Retrieved a prompt from PLMS
* [x] Review widget MR https://gitlab-app.eng.qops.net/vocalize/vocalize/-/merge_requests/8565 ✅ 2025-05-07
* Told core enrichments to copy my 1158 work to their IG services https://qualtrics.slack.com/archives/C087V8K2GFP/p1746657151474649
* Filed task to reverse shadow classification service deployment https://qualtrics.atlassian.net/browse/ENRICH-1432

# Log
* [x] tell other teams to do their own [[ENRICH-1158- Enhance the Engine enrichment service to support skipping messages by configuration.]] and link their tasks to 1158 ✅ 2025-05-07
* [x] make a jira task to deploy CS with kuan and my changes, i'll reverse shadow kuan. add mavin. schedule a time too? check [thread](https://qualtrics.slack.com/archives/C087VC5JPD3/p1746638687944019?thread_ts=1746637220.638189&cid=C087VC5JPD3) for updates ✅ 2025-05-07
* [ ] learn about orgh https://qualtrics.slack.com/archives/C036JC1JKGE/p1746648657884079?thread_ts=1746632124.148359&cid=C036JC1JKGE 🔽 

# Start of day tasks
* [x] Copy meeting/other schedule to notes ✅ 2025-05-07
* [x] Check email ✅ 2025-05-07
* [x] Check slack ✅ 2025-05-07

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/Status/2025-05-08 Thu.md
---

- Meetings/other
    - ~~AI tools for Query-Based Topics~~ rescheduled
    - Knowledge Sharing - SQL to check unsupported feature
- [[ENRICH-1405- Spike- Understand project and break into smaller tasks]]
    - Verified access to `socrates-human-customer-console-access-2-role` 
    - What's an experiment prompt in plms? [good thread](https://qualtrics.slack.com/archives/C0186S22PQS/p1719401057614159)
    - Writing/thinking [[ENRICH-1114- Smart Query Prompt Updates#My plan to complete this story]]
    - Researching mtoken
- [x] [[ENRICH-1432- Reverse shadow deploying classification service with Kuan]] ✅ 2025-05-08
    - Scheduled
    - Walked through deployment of classification service, engines controller, then unbinding the old classification service queue
- [[ENRICH-1158- Enhance the Engine enrichment service to support skipping messages by configuration.]]
    - Closed
    - Told core enrichments that this change is verified and deployed and reminded them to implement it for their IG services
- [-] [[ENRICH-1438- Complete deployment of classification service with sharding and blocklist changes to prod]] ⏫ ❌ 2025-05-15
    - Filed this item
    - Clarified deployment scopes with Mavin

# Log
- Notes from [[ENRICH-1432- Reverse shadow deploying classification service with Kuan]]
    - get access to https://vault-shared.clarabridge.io/ui/vault/secrets/secret-prod2-euw2/kv/infra%2Frabbitmq%2Fcb-ingestion-pipeline-prod2-euw2-rabbitmq%2Fconnection-info/details, https://vault-shared.clarabridge.io/ui/vault/secrets/secret-prod2-euw2/kv/infra%2Frabbitmq%2Fcb-ingestion-pipeline-prod2-euw2-rabbitmq%2Fconnection-info/details
    - deployment job https://jenkins-shared.clarabridge.io/job/DEPLOYMENTS/job/APPLICATIONS/job/ENGINES/job/SERVICES/job/CLASSIFICATION-SERVICE/job/PROD2_EUW2/job/classification-service-deploy/
    - figure out how to get access to prod us rabbit mq. it's different. self hosted?
    - me and mavin complete deployment

# Start of day tasks
- [x] Copy meeting/other schedule to notes ✅ 2025-05-08
- [x] Check email ✅ 2025-05-08
- [x] Check slack ✅ 2025-05-08
- [-] Spend >5 minutes on a starved task 📅 2025-05-08 ❌ 2025-05-09

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/Status/2025-05-09 Fri.md
---

- Meetings/other
    - 
- sick

# Log
- 

# Start of day tasks
- [x] Copy meeting/other schedule to notes ✅ 2025-05-12
- [x] Check email ✅ 2025-05-12
- [x] Check slack ✅ 2025-05-12
- [-] Spend >5 minutes on a starved task 📅 2025-05-09 ❌ 2025-05-12

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/Status/2025-05-12 Mon.md
---

- Meetings/other
    - 
- sick
- [-] Investigate protocat maintenance [slack](https://qualtrics.slack.com/archives/C087VC5JPD3/p1747065038018239?thread_ts=1746815041.447499&cid=C087VC5JPD3) ❌ 2025-05-13
- [ ] TAPP-7072- Invisible content is perceptible in bubble charts 🛫 2025-07-01 📅 2025-08-07

# Log
- 

# Start of day tasks

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/Status/2025-05-13 Tue.md
---

- Meetings/other
    - Scrum
    - Sprint review
    - Huddle with Vamsi 
- [x] Prepare sprint review slides 📅 2025-05-13 ✅ 2025-05-13
    - Copied slides for sprint 105, posted
- [x] Find meeting notes for the AI sync last week or ask about the status of guardrails ✅ 2025-05-13
    - [Asked](https://qualtrics.slack.com/archives/C07MKL3FB61/p1747159409784589)
    - Donovan says "we are not keeping harmful content filter on smart query"
    - ~~Cancelling~~ [[ENRICH-1295- Designer Error Handling for smart-query guardrails]]
        - Uncancelled it since Vamsi said we still need to do it for max tokens guardrail
    - Addisu says max tokens isn't implemented 
    - Followed up with the info I got from Addisu and Vamsi, asked about putting harmful content filter behind a feature flag
- Skimming Gordon's gameday [doc](https://docs.google.com/document/d/1lroYB7Ux6WRWWbWmP_AScOh5JyAgdVdNDh79bgdyaAE/edit?usp=drive_open&ouid=115243496293241999987)
- [[ENRICH-1438- Complete deployment of classification service with sharding and blocklist changes to prod]]
    - Found I still don't have the vault access I requested in March, [followed up](https://qualtrics.slack.com/archives/C02G9BZ48ER/p1747157612297089?thread_ts=1742406935.505509&cid=C02G9BZ48ER)
        - Got access
    - Finding different urls for prod US rmq console. The broker changed earlier this month? Asking for [help](https://qualtrics.slack.com/archives/C036JC1JKGE/p1747159077386569)
        - Have to port forward
    - Prod CA deployment complete
    - https://vault-shared.clarabridge.io/ui/vault/secrets/secret-prod/kv/infra%2Frabbitmq%2Fcb-ingestion-pipeline-prod-us-rabbitmq%2Fconnection-info/details?version=2
- [-] [[TAPP-7074- TextIQ Bubble charts content is not perceptible and Bubble charts cannot be operated via keyboard]] 🛫 2025-06-15 📅 2025-07-18 ❌ 2025-06-17
    - Assigned to me by Rohit, triaging
    - Asked Vamsi to opine on scope, urgency [slack](https://qualtrics.slack.com/archives/C08RHU9K8NS/p1747171354456159)
- Started a conversation about [[ENRICH-1114- Smart Query Prompt Updates]] testing with Steven
    - He says it's mostly manual testing https://qualtrics.slack.com/archives/D03GSPD5K0A/p1747176982852249

# Log
- [x] Ask for on-call overrides 🛫 2025-05-15 ✅ 2025-05-15
- I said I would finish [[ENRICH-1405- Spike- Understand project and break into smaller tasks]] today
- [x] Check OKRs ✅ 2025-05-13
- [x] Talk to steven about smart query prompt update testing ✅ 2025-05-14
- [x] Talk to Vamsi about [[2025 Objective Plan next 1:1 📅 2025-05-15 ✅ 2025-05-15
- Huddle with Vamsi
    - guardrails is more than just harmful content filter. also max tokens.
    - [x] Check if Addisu has implemented max tokens guardrail for smart query ✅ 2025-05-13
        - he says no
    - Vamsi says we should still do error handling for max tokens
    - [x] Ask Donovan if we should put harmful content guardrail behind a feature flipper instead of wholesale disabling it ✅ 2025-05-13
- Chatted with Reid

# Start of day tasks
- [x] Copy meeting/other schedule to notes ✅ 2025-05-13
- [x] Check email ✅ 2025-05-13
- [x] Check slack ✅ 2025-05-13
- [x] Spend >5 minutes on a starved task 📅 2025-05-13 ✅ 2025-05-13

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/Status/2025-05-14 Wed.md
---

- Meetings/other
    - Sprint planning (1 hr)
    - Scrum/mid-quarter check-in (1.25 hr)
    - 2x Dr appt
- [[ENRICH-1295- Designer Error Handling for smart-query guardrails]]
    - Donovan says harmful content guardrail isn't needed but max tokens is. I'm pushing for clarification [slack](https://qualtrics.slack.com/archives/C07MKL3FB61/p1747237985604019?thread_ts=1747159409.784589&cid=C07MKL3FB61)
    - We discussed this in scrum/mid-quarter check-in. I said I'm getting conflicting information. Addisu and Vamsi discussed. I asked them to move the discussion to slack so that it's written down.
- [[ENRICH-1438- Complete deployment of classification service with sharding and blocklist changes to prod]]
    - Vamsi says hold deployment to prod US while they sort out other issues [slack](https://qualtrics.slack.com/archives/C053MK619T9/p1747228633855699?thread_ts=1747224169.070739&cid=C053MK619T9)
- [[TAPP-7074- TextIQ Bubble charts content is not perceptible and Bubble charts cannot be operated via keyboard]]
    - Vamsi says we can push to make the widget P1 a backlog item for the new widget team
    - Asking to move this pulse to backlog [slack](https://qualtrics.slack.com/archives/C08RHU9K8NS/p1747243547439809?thread_ts=1747171354.456159&cid=C08RHU9K8NS)

# Log
- Sprint planning
    - Vamsi says punt [[TAPP-7072- Invisible content is perceptible in bubble charts]] to Q3, see if new team can do it or reverse shadow
- Scrum/mid-quarter check-in
    - Team is close to open revolt over quarterly commits, small on-call rotation. Gripe session.
    - Vamsi wants to change primary/secondary on-call to be you go onto secondary after primary. I asked what to do if two engage people are on-call and we get alerts we don't know how to handle at midnight. Vamsi and Addisu talked a lot. Vamsi said to page the "eng leads on-call rotation" which is like him, Corbin, etc. There was talk about paging other team members; I said I don't want to be paged off rotation.
- Vamsi told Donovan that [[ENRICH-1114- Smart Query Prompt Updates]] is at risk for Q2 delivery

# Start of day tasks
- [x] Copy meeting/other schedule to notes ✅ 2025-05-14
- [x] Check email ✅ 2025-05-14
- [x] Check slack ✅ 2025-05-14
- [x] Spend >5 minutes on a starved task 📅 2025-05-14 ✅ 2025-05-15

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/Status/2025-05-15 Thu.md
---

- Meetings/other
    - AI tools for Query-Based Topics
    - [[2025-05-15 Thu - Vamsi 1 on 1|2025-05-15 Thu - Vamsi 1 on 1]]
- AI tools for Query-Based Topics meeting
    - Vamsi says "post GA scope" means after Q2 wrt [this](https://qualtrics.slack.com/archives/C07MKL3FB61/p1747244204073309?thread_ts=1747159409.784589&cid=C07MKL3FB61)
    - I asked for clarification on what "post GA scope" is; it was being used differently in different places. Vamsi updated the header of my section in the notes from "post GA smart query" to "V2 of smart query".
    - I pushed against having things added to the scope of my Q2 deliverable
    - Heard about "prompt injection guardrail" from product for the first time, and that it's required. I noted that this is a new thing I've never heard of and resisted it being added to the scope of my Q2 deliverable.
        - Vamsi says since error handling UI was removed from my scope, maybe we can investigate adding this. Wait, but won't *this* guardrail need error UI?
    - I noted that I feel new things are being added to my scope. The idea of scope I was working on is what Hardik gave [here](https://qualtrics.slack.com/archives/C036JC1JKGE/p1744048974592049?thread_ts=1743786031.481419&cid=C036JC1JKGE) and asked Donovan to confirm, which he didn't.
    - I noted I feel I have no control
- After ^ meeting
    - I researched how I'd gotten my idea of the scope for smart query v2. I found Hardik's message and realized I'd misunderstood it. I thought "guardrails" meant just error handling UI for some reason.
    - I posted this info, acknowledged my mistake, and apologized for my tone in the meeting [slack](https://qualtrics.slack.com/archives/C07MKL3FB61/p1747326286849879)
- [[TAPP-7074- TextIQ Bubble charts content is not perceptible and Bubble charts cannot be operated via keyboard]]
    - Claire asked what priority would work for this item. I asked what the underlying customer need is. She's going to check with the pulse filer. [slack](https://qualtrics.slack.com/archives/C08RHU9K8NS/p1747324778967669?thread_ts=1747171354.456159&cid=C08RHU9K8NS)
- Spent time preparing an agenda for 1:1
- [[ENRICH-1405- Spike- Understand project and break into smaller tasks]]
    - Created jira subtasks under 1114, wrote descriptions, pointed
    - Closed
- [[ENRICH-1438- Complete deployment of classification service with sharding and blocklist changes to prod]]
    - Closed as duplicate of [[ENRICH-1440- New Sharding Implementation: Deploy changes across discover DCs]]
- Created epic to help organize my Q2 deliverable work: [[ENRICH-1114- Smart Query Prompt Updates]]

# Log
- [-] Review gameday kickoff, expectations [email](https://mail.google.com/mail/u/0/#inbox/FMfcgzQbfLRLXhtHSzrglnrRcqcTnXvV) 📅 2025-05-28 ❌ 2025-05-27
- [x] Investigate best way to remove/disable harmful content guardrail [slack](https://qualtrics.slack.com/archives/C07MKL3FB61/p1747326478317269?thread_ts=1747325271.138709&cid=C07MKL3FB61) ✅ 2025-05-16
- [x] [[ENRICH-1440- New Sharding Implementation: Deploy changes across discover DCs]] ✅ 2025-06-18


# Start of day tasks
- [x] Copy meeting/other schedule to notes ✅ 2025-05-15
- [x] Check email ✅ 2025-05-15
- [x] Check slack ✅ 2025-05-15
- [x] Spend >5 minutes on a starved task 📅 2025-05-15 ✅ 2025-05-16

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/Status/2025-05-16 Fri.md
---

- Meetings/other
    - Scrum
    - Ops
- Watched 5 minutes of k8s autoscaling training
- Reviewed Addisu's thread with CME on socrates-provided guardrails [slack](https://qualtrics.slack.com/archives/C0186S22PQS/p1742922021578449)
    - Asked Addisu what work he's doing with guardrails to make sure we're not duplicating effort [slack](https://qualtrics.slack.com/archives/C07MKL3FB61/p1747430416055509?thread_ts=1747409159.683189&cid=C07MKL3FB61)
- Sidetrack: exploring `slackdump` to get more context on my work from slack into a format suitable to dump to LLMs. Can't get auth to work.
    - Sent [MyQ request](https://qualtrics.service-now.com/myq?id=ticket&table=sc_req_item&sys_id=13e987299725a250fb85732f2153afa9&view=sp) to corp biz apps asking for app approval or them to run an export for me
    - Got a security investigation filed on me for this 😅
    - They said I might need bizapps to help with avoid anomaly detection
    - Export requires legal approval and it sounds like they want any exports deleted asap. Also, I wouldn't get up to date info, so abandoning this idea.
    - Asked about getting around slack anomaly detection
- Trying out windsurf. Had it whip up a script to compile my work notes into one big text file for dumping to LLMs as context.
- [x] Change [[ENRICH-1114- Smart Query Prompt Updates]] to an epic, its subtasks to stories ✅ 2025-05-16
    - Started [[Jira]] KB page
    - Asked team about jira item hierarchy [slack](https://qualtrics.slack.com/archives/C087VC5JPD3/p1747417782726449)
- [[ENRICH-1464- Find/load test data for behavior/perf evaluation]]
    - Asked Steven if he can help find test data [slack](https://qualtrics.slack.com/archives/C07MKL3FB61/p1747431472432269)
    - He found some [slack](https://qualtrics.slack.com/archives/C07MKL3FB61/p1747433618003899?thread_ts=1747431472.432269&cid=C07MKL3FB61)
    - He offered to make a project for me and I accepted
- [[ENRICH-1114- Smart Query Prompt Updates]]
    - Asked Donovan about measures of success, as Vamsi suggested yesterday [slack](https://qualtrics.slack.com/archives/C07MKL3FB61/p1747432019857009)
    - Donovan: "80% reduction while maintaining relevance and coverage seems like the right standard to me"
- [[ENRICH-1468- Smart query version 2 in designer (epic)]]
    - Asked for clarification about "post GA scope" for max token guardrail [slack](https://qualtrics.slack.com/archives/C07MKL3FB61/p1747435433893479?thread_ts=1747159409.784589&cid=C07MKL3FB61)
    - Created [[ENRICH-1484- Remove or disable harmful content guardrail in Smart Query]]
- Prepared agenda for Monday 1:1 with Vamsi
- [x] [[ENRICH-1484- Remove or disable harmful content guardrail in Smart Query]] ✅ 2025-05-21
    - Cloned repo, loaded up the python scripts in IDE
    - Skimmed the code. Removal will be easy.

# Log
- [x] check on overrides, respond to Mavin's comment nicely ✅ 2025-05-16
- In ops, I asked how confident we are that the OXM-seconded people will come back to the on-call rotation at start of Q3 given that OXM is behind. Vamsi is going to ask in his Evan staff meeting. There were some assurances behind even if OXM continues past Q2.
- Zeshan asked to notify unified orchestration when we make rabbitmq changes since things are volatile lately, especially in prod US with op-7230 [slack](https://qualtrics.slack.com/archives/C087VC5JPD3/p1747428027687379?thread_ts=1747412483.577999&cid=C087VC5JPD3)
- Using gemini a lot now. Especially to write messages for me. I like being succinct, but I think that comes across poorly.

# Start of day tasks
- [x] Copy meeting/other schedule to notes ✅ 2025-05-16
- [x] Check email ✅ 2025-05-16
- [x] Check slack ✅ 2025-05-16
- [x] Spend >5 minutes on a starved task 📅 2025-05-16 ✅ 2025-05-16

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/Status/2025-05-19 Mon.md
---

- Meetings/other
    - Vamsi 1:1
- Watched 5 min of k8s autoscaling training (resume at 12:35)
- [[TAPP-7074- TextIQ Bubble charts content is not perceptible and Bubble charts cannot be operated via keyboard]]
    - Since due date is approaching, noted in the jira ticket that I'm waiting to hear about customer's timeframe need, it's not doable in P1 timeframe.
- [[ENRICH-1484- Remove or disable harmful content guardrail in Smart Query]]
    - Find SQ inputs that trigger guardrails for removal verification
    - In b1, output guardrail doesn't seem to be working for expletives 🤷🏻‍♀️
    - Verified lambda python code in b1 is basically same as in repo
    - Finally got harmful content guardrail to trigger by changing sensitivity to high and being profane
    - Removed guardrail calls, uploaded to b1 lambda to test
    - Getting a different refusal to cooperate. The refusal is in prose from the AI model this time.
    - I've replicated the issue with a direct call to MIG
    - Asking socrates team [slack](https://qualtrics.slack.com/archives/C0186S22PQS/p1747701793764599)
- [[2025-05-19 Mon - Vamsi 1 on 1]]
    - Pinged Vamsi about jira structure for [[ENRICH-1468- Smart query version 2 in designer (epic)]] since it came up [slack](https://qualtrics.slack.com/archives/C087VC5JPD3/p1747687835467969?thread_ts=1747417782.726449&cid=C087VC5JPD3)
        - [x] Document Vamsi's jira structure preference once received 🔽 ✅ 2025-05-20
    - Processing info from meeting into jira items
        - Created [[ENRICH-1487- Add prompt injection guardrail to Smart Query]], [[ENRICH-1488- Add max tokens guardrail to Smart Query]]
- Asked for full day on-call swaps for my longer overrides [slack](https://qualtrics.slack.com/archives/C087VC5JPD3/p1747690101555339?thread_ts=1747255988.567209&cid=C087VC5JPD3)
    - [x] Check on overrides 🛫 2025-05-21 ✅ 2025-05-21
- [[ENRICH-1487- Add prompt injection guardrail to Smart Query]]
    - Reading AWS docs on guardrails
    - Researching Addisu's thread with CME on why SQ has trouble with MIG's guardrail API [[Query Topics AI Features (lambdas)#Thread on MIG (aka CME, socrates) guardrails]]
        - Got it. I'll try using the standalone guardrail api just for user inputs. [[ENRICH-1487- Add prompt injection guardrail to Smart Query#Smart Query - MIG Guardrail API Approach Decision]]
    - Asked Donovan if we want to protect against jailbreak as well as prompt injection [slack](https://qualtrics.slack.com/archives/C07MKL3FB61/p1747695732998839) (AWS lumps them both into "prompt attack")

# Log
- [x] Send scrum notes 📅 2025-05-19 ✅ 2025-05-19
- [x] Ask for on-call override day swaps 📅 2025-05-19 ✅ 2025-05-19

# Start of day tasks
- [x] Copy meeting/other schedule to notes ✅ 2025-05-19
- [x] Check email ✅ 2025-05-19
- [x] Check slack ✅ 2025-05-19
- [x] Spend >5 minutes on a starved task 📅 2025-05-19 ✅ 2025-05-19

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/Status/2025-05-20 Tue.md
---

- Meetings/other
    - Scrum
    - Enrichments RCA Review and Discussion
- Discussing [[ENRICH-1283- Update survey flow plugin to use sentiment v3 API]] with Vamsi. Due today, unassigned.
    - Starting thinking about how to address this kind of issue in [[2025-05-27 Tue - Sprint Retro]]
    - Saw we don't have a retro planned and it's almost been a month. Asked [slack](https://qualtrics.slack.com/archives/C087VC5JPD3/p1747760976443559)
    - [x] Follow up about next retro 🛫 2025-05-27 ✅ 2025-05-27
- Responded to MyQ ticket, provided slackdump docs
- [[ENRICH-1484- Remove or disable harmful content guardrail in Smart Query]]
    - Confirmed with socrates team that there is some unavoidable harmful content protection in the model itself
    - Informed Donovan of this [slack](https://qualtrics.slack.com/archives/C07MKL3FB61/p1747761430202169)
    - Pushed change to disable MIG content guardrails
    - Sent MR
    - Addressing MR feedback:
        - Replace print statements with real logger
    - Tried getting AWS Toolkit plugin set up in IntelliJ to speed up iterating on lambda code. Gave up.
    - Repo requires pipeline to pass before merging. I can't change the setting because I'm not a maintainer. Asked Addisu to make me maintainer.
- [[ENRICH-1114- Smart Query Prompt Updates]]
    - Leo asked "is there an outline of the experiments to achieve that? from what i remember the stars were added manually in post processing after the LLM suggested keywords"
        - [x] Figure out how wildcards are added. Does the model generate them or are they added in post-processing? ✅ 2025-05-20
        - The model adds the wildcards. Informed Leo.
        - Leo still thinks something is done in post-processing. Investigating...
    - Asked Addisu about how the socrates sdk layer is added to the lambda [slack](https://qualtrics.slack.com/archives/C07MKL3FB61/p1747763004816949)
- [[ENRICH-1158- Enhance the Engine enrichment service to support skipping messages by configuration.]]
    - My header parsing logic in this change is causing issues in OP-7858
    - Mavin reverted my change
    - I reopened this jira item and tagged Vamsi, noting I'd be happy to work on it again but have unassigned myself from it pending his prioritization of the work. [jira](https://qualtrics.atlassian.net/browse/ENRICH-1158?focusedCommentId=11465452)
    - [x] Follow up on [[ENRICH-1158- Enhance the Engine enrichment service to support skipping messages by configuration.]] being assigned/prioritized 🛫 2025-05-21 ✅ 2025-05-30
    - Charmish asked me to announce to data pipelines (unified orchestration?) when we're deploying the sharding change to Prod US
        - Done [slack](https://qualtrics.slack.com/archives/C08T9CFN0SC/p1747771591222629)
- Enrichments RCA Review and Discussion
    - Reviewed RCA doc
- [[ENRICH-1464- Find/load test data for behavior/perf evaluation]]
    - Steven found/made (?) a project for me [slack](https://qualtrics.slack.com/archives/C07MKL3FB61/p1747766146395009?thread_ts=1747431472.432269&cid=C07MKL3FB61)
        - > Anyway, you can use the 'Autotune-GA1' project in staging03 (AWS Staging)
    - Asked if it's for my exclusive use
        - > No, but I made it a week or so ago for a specific case and I don't think anyone else is using it
- Scrum
    - My question about how to handle the descendants of [[ENRICH-1114- Smart Query Prompt Updates]] came up. Vamsi said...
        - [x] put the descendant stories in [[ENRICH-1468- Smart query version 2 in designer (epic)]] ✅ 2025-05-20
            - Closed 1114 as duplicate of 1468
    - Addisu says he's working on PI guardrail for topic recommendations. Might want to consult him when I work on [[ENRICH-1487- Add prompt injection guardrail to Smart Query]]
- No takers for my offer to swap two days on my on-call shift. Tagged and asked team if anyone could switch weeks.
- Asked help-coding-ai about my snag with windsurf making me accept my manual changes [slack](https://qualtrics.slack.com/archives/C051LMXB1LK/p1747775041941489)
- Helping with a customer question about disabling text iq [slack](https://qualtrics.slack.com/archives/C03AA5ZC8/p1747779789864259?thread_ts=1701915992.943039&cid=C03AA5ZC8)

# Log
- Vamsi asked if I'm closing [[ENRICH-1283- Update survey flow plugin to use sentiment v3 API]] today since it's due. I said no (nicely) because it isn't assigned to me and I've been working on other stuff [slack](https://qualtrics.slack.com/archives/C087VC5JPD3/p1747753057210139?thread_ts=1747746310.996499&cid=C087VC5JPD3)
    - I feel it's his job to be looking through the backlog for stuff like this and prioritizing it. I had no idea he cared about the due date. Maybe he doesn't.

# Start of day tasks
- [x] Copy meeting/other schedule to notes ✅ 2025-05-20
- [x] Check email ✅ 2025-05-20
- [x] Check slack ✅ 2025-05-20
- [x] Spend >5 minutes on a starved task 📅 2025-05-20 ✅ 2025-05-20

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/Status/2025-05-21 Wed.md
---

- Meetings/other
    - Qualtrics Eng Tech Talks | \[Windsurf \]
    - Scrum
    - Dr appt
- Whipped up a backup script for my work obsidian vault and automated scheduled runs (after a deletion scare yesterday)
- [[ENRICH-1484- Remove or disable harmful content guardrail in Smart Query]]
    - Addisu confirmed that deployment of the lambda code has been entirely manual, including adding the socrates sdk layer. [slack](https://qualtrics.slack.com/archives/C07MKL3FB61/p1747829317512289?thread_ts=1747763518.792909&cid=C07MKL3FB61)
    - Pinged Charmish about MR
    - Merged
- [[TAPP-7074- TextIQ Bubble charts content is not perceptible and Bubble charts cannot be operated via keyboard]]
    - Oops, missed Claire's response from last Friday. She says account team wants this fixed by contract renegotiation start in July. I asked for a more specific date. [slack](https://qualtrics.slack.com/archives/C08RHU9K8NS/p1747835404500499?thread_ts=1747171354.456159&cid=C08RHU9K8NS)
    - [x] follow up on 7074 🛫 2025-05-27 ✅ 2025-05-27
- [x] Figure out xmd gov1 support override ⏫ ✅ 2025-05-21
    - I think VO just includes the gov1 rotation since I'm on it. I'm not actually on-call for that rotation until July
- Reviewed https://gitlab-app.eng.qops.net/xm-discover/discover-platform-group/discover-data-pipelines/platform/-/merge_requests/6268
    - Reviewed changes, approved
- Lots of slack
- I should work on [[ENRICH-1487- Add prompt injection guardrail to Smart Query]] and see if sending the user inputs works, because if we need CME to make changes, that'll take time.
    - [x] [[ENRICH-1487- Add prompt injection guardrail to Smart Query]] 🔼 ✅ 2025-05-30
- Updated calendar and QEH with June OOOs and Dr appts
- Watched some of k8s autoscaling training (resume at 21:50)
- Pinged Vamsi about next sprint retro
- [[ENRICH-1440- New Sharding Implementation: Deploy changes across discover DCs]]
    - There was a miscommunication. I'm still deploying sharding to discover DCs. I informed the thread [slack](https://qualtrics.slack.com/archives/C053MK619T9/p1747861727395759?thread_ts=1747224169.070739&cid=C053MK619T9)
    - Mavin hotfixed CS--reverted my blocklist change [[ENRICH-1158- Enhance the Engine enrichment service to support skipping messages by configuration.]] ^cf9012
    - Confirmed all DCs have the same image tag in deployment config
    - Wait... do I have access to prod US rmq console?
        - No one ever answered my request for a credential, pinging [slack](https://qualtrics.slack.com/archives/C036JC1JKGE/p1747863290564419?thread_ts=1747159077.386569&cid=C036JC1JKGE)
    - Updated Unified Orchestration on deployment status
    - Don't deploy to prod US until scale test is done [slack](https://qualtrics.slack.com/archives/C053MK619T9/p1747864899698419?thread_ts=1747224169.070739&cid=C053MK619T9)
    - Starting prod CA
        - Verified rabbitmq has `classification-service-router` bound to `engine.tx`
        - Deployed
        - Verified classification queue is clearing
    - Starting prod EU
        - Verified rabbit binding config
        - Deployed
        - Classification queue has no traffic for past ten hours
    - Starting prod APSE
        - Verified rabbit binding config
        - The last jenkins build/deployment used branch ENRICH-1330. Changed to master for my build.
        - Deployed
        - Verified classification queue is clearing
- Set vacation email response, slack status
- Reminded team I'll be on vacation
- Pasted [[2025-05-27 Tue - Sprint Retro]] to slack and tagged Vamsi in case retro happens while I'm on vacation

# Log
- [x] Ask for Jun 11 shift overrides 🔼 🛫 2025-05-27 ✅ 2025-06-02
- Deploying classification service to discover SUCKS
- [x] TELL UO IF YOU'RE DEPLOYING TO PROD US! ⏫ 📅 2025-06-17 ✅ 2025-06-18
- [x] See if reinstalling slack gives me the option to format pasted text again ✅ 2025-05-27

# Start of day tasks
- [x] Copy meeting/other schedule to notes ✅ 2025-05-21
- [x] Check email ✅ 2025-05-21
- [x] Check slack ✅ 2025-05-21
- [x] Review TAPP and ENRICH near due jira items (the sauron slack messages) ✅ 2025-05-21
- [x] Spend >5 minutes on a starved task 📅 2025-05-21 ✅ 2025-05-21

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/Status/2025-05-27 Tue.md
---

- Meetings/other
    - Sprint review (1 hr)
    - [[2025-05-27 Tue - Sprint Retro]] (1 hr + .25 hr)
- Catching up on emails and slack (1 hr)
- [x] Ask Charmish to handle Gordon's Jun 3-4 gameday ✅ 2025-05-27
- [[TAPP-7074- TextIQ Bubble charts content is not perceptible and Bubble charts cannot be operated via keyboard]]
    - Claire responded that the account team says late July is fine for delivering this
    - Moved due date to Jul 18
    - Rohit asked me to have rafael ochoa (rochoa) shadow when I start work on this [slack](https://qualtrics.slack.com/archives/C019UF4CE90/p1748375218882139)
- [x] Make sprint review slides 📅 2025-05-27 ✅ 2025-05-27
    - Created slides for team
- [x] Investigate Addisu's question about ENRICH-1359 https://qualtrics.atlassian.net/browse/ENRICH-1359?focusedCommentId=11496294 ✅ 2025-05-27
- [[ENRICH-1487- Add prompt injection guardrail to Smart Query]]
    - Raj says our "PI" term wrt guardrails includes both PI and jailbreak [slack](https://qualtrics.slack.com/archives/C07MKL3FB61/p1747955456699159?thread_ts=1747695732.998839&cid=C07MKL3FB61)
- [[ENRICH-1440- New Sharding Implementation: Deploy changes across discover DCs]]
    - Following up with Andy about Prod US rmq web UI credential
    - Received credential [slack](https://qualtrics.slack.com/archives/D08U04GGZFG/p1748366682984039)
- Created `#obsidian-enthusiasts` in the social slack workspace. There are other users here.
- [[ENRICH-1488- Add max tokens guardrail to Smart Query]]
    - [x] Ask product if they want actual token counting or a char limit. Could start with char limit since it's easy and stretch to token counting if there's time. ✅ 2025-06-02
        - Wait... aren't there already char limits in the UI? Are they enforced at API level too? Is that good enough for now?
        - Asked [slack](https://qualtrics.slack.com/archives/C07MKL3FB61/p1748368990458529)
- Completed Q2 EX pulse
- Finished watching k8s autoscaling training
- [Asked](https://qualtrics.slack.com/archives/C087VC5JPD3/p1748380627769509) Vamsi about ktlo jira items [[2025-05-27 Tue - Sprint Retro#^6a6166]]
    - > oh I see, yeah such smaller items don't need to added \[to the ktlo epic\]
    - Next time I'm on-call I should take care to create these ktlo jira items
- [[ENRICH-1487- Add prompt injection guardrail to Smart Query]]
    - Had AI write a thorough description for this story
    - Started writing code

# Log
- [x] See if we have another retro planned for late June 🛫 2025-06-11 📅 2025-06-12 ✅ 2025-06-11
- Sprint review
    - Asked if the backfill L6 and new L3 are being hired in a specific location. Yes, Reston. I asked why. Vamsi said there's no senior/L5+ people on our team in Reston now.

# Start of day tasks
- [x] Copy meeting/other schedule to notes ✅ 2025-05-27
- [x] Check email ✅ 2025-05-27
- [x] Check slack ✅ 2025-05-27
- [x] Review TAPP and ENRICH near due jira items (the sauron slack messages) ✅ 2025-05-27
- [x] Spend >5 minutes on a starved task 📅 2025-05-27 ✅ 2025-05-27

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/Status/2025-05-28 Wed.md
---

- Meetings/other
    - Sprint loading
    - Dr appt
- Collected items for sprint loading
- [[ENRICH-1488- Add max tokens guardrail to Smart Query]]
    - Donovan replied
        - > I thought DICE just had a basic token limit guardrail you can implement.  I think this is on the input tokens.  Do you have documentation from DICE on this?
    - Researching
    - I don't see evidence of an input token limit for MIG. Responded in thread.
- Signed up for windsurf pilot (again?)
- [[ENRICH-1487- Add prompt injection guardrail to Smart Query]]
    - Writing code
    - The standalone user input check works!
    - Big refactor, checked with Addisu first to see if he has work in progress
        - Trying to figure out aws-cli login to speed up iteration
            - Gave up
        - Hm, I was doing this refactor because the codebase is messy and I want to write unit tests, but I'm not sure this is the best use of my time. This code addition is simple. I can come back to the refactor if I have time later.
        - Pushed WIP branch jbrewer-1487-refactor-snapshot
    - Mirroring changes to the smart-query-expansion file
    - Buckled, refactored just the code I'm adding to be shared between batch and expansion
    - Created deployment script since the python scripts aren't monolithic/self-contained anymore
    - Started on a readme update, decided to finish it tomorrow so I could get the code out for review
    - Sent MR for review

# Log
- I saw in the Q2 PPR/planning coda that I'm listed with 2 days out. Idk if I wrote that wrong or what. But I said I'm out for 2 *weeks* [here](https://qualtrics.slack.com/archives/C087VC5JPD3/p1741878378497339?thread_ts=1741808595.697219&cid=C087VC5JPD3).
- Addisu's on vacation starting tomorrow, returning Jun 13.
- I should remember to disable myself as a code reviewer before going on vacation, and enable myself when I return.

# Start of day tasks
- [x] Copy meeting/other schedule to notes ✅ 2025-05-28
- [x] Check email ✅ 2025-05-28
- [x] Check slack ✅ 2025-05-28
- [x] Review blocked tasks ✅ 2025-05-28
- [x] Review TAPP and ENRICH near due jira items (the sauron slack messages) ✅ 2025-05-28
- [-] Spend >5 minutes on a starved task 📅 2025-05-28 ❌ 2025-05-29

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/Status/2025-05-29 Thu.md
---

- Meetings/other
    - Backlog refinement
    - AI tools for Query-Based Topics
    - [[2025-05-29 Thu - Vamsi 1 on 1]]
- [[ENRICH-1487- Add prompt injection guardrail to Smart Query]]
    - Address MR feedback
    - Adding more quality of life improvements to the MR branch
    - Researching how to call lambda-inference-gateway (LIG) directly to speed up testing and verify the request body structure
        - Got it! Way faster now.
    - Curious... LIG might be changing response codes. I'm responding 400 from the lambda but getting 200 from LIG.
        - Yeah seems like LIG returns everything as 200. [Asking](https://qualtrics.slack.com/archives/C0186S22PQS/p1748551969138119)
        - Confirmed. Added this info to readme.
    - Hopefully this MR is about done. Told reviewers this.
- As part of a gameday, reviewed our (new?) [runbook on brand migration](https://coda.io/d/Query-Topics-Team-hub_dt-izYOVJmq/Text-Engine-Data-Relocation-Service_suqocjwY)
- Prepared for 1:1

# Log
- AI tools for Query-Based Topics
    - I said I'm about to finish [[ENRICH-1487- Add prompt injection guardrail to Smart Query]] and start on [[ENRICH-1463- Iterate on prompt to reduce wildcard use]] next
    - I said I'm on track for mid-June code complete
    - I'm being re-assigned to help with variance testing for smart audit. Donovan says it's more important than reducing wildcard use. Vamsi didn't object.
    - [x] Find out what variance testing for smart audit is ⏫ ✅ 2025-05-30
    - Charmish to verify Vamsi's okay with me working on this, and sync with me tomorrow if so
    - Donovan or Daniel asked Raj Verma about the max token question--what's it for?

# Start of day tasks
- [x] Copy meeting/other schedule to notes ✅ 2025-05-29
- [x] Check email ✅ 2025-05-29
- [x] Check slack ✅ 2025-05-29
- [x] Review blocked tasks ✅ 2025-05-29
- [x] Review TAPP and ENRICH near due jira items (the sauron slack messages) ✅ 2025-05-29
- [-] Spend >5 minutes on a starved task 📅 2025-05-29 ❌ 2025-05-30

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/Status/2025-05-30 Fri.md
---

- Meetings/other
    - Variance testing sync
    - Scrum
    - Ops
- Variance testing sync
    - Met with Charmish and Steven, discussed the testing. Steven's interested in doing it and I'm going to help him get set up to hit the LIG API
- Working on [[2025 Objective Plan]]
    - [x] Sign up for feedback 360 [slack](https://qualtrics.slack.com/archives/C3P4NU88Y/p1748461188974979) ⏫ 📅 2025-06-02 ✅ 2025-06-02
    - Created three objectives with key results
    - Review with Vamsi next 1-1 [[2025-06-09 Mon - Vamsi 1 on 1#^340b4b]]
    - Told Vamsi I wrote three objectives
- Vamsi says I should work on "rate limiting" since Steven is taking Smart Audit (SA) variance testing
    - [x] [[ENRICH-1532- Rate limiting in Topic Recommendations]] 🔼 ✅ 2025-06-04
        - Assigned it to me, brought it into this sprint 107
- Counting PTO in Q3
    - Estimate 20 2-hour medical appointments (sick leave): 5 days
    - I should take some vacation. I'll have 11 days at the end of Q3. Estimate 7 days.
    - [slack](https://qualtrics.slack.com/archives/C087VC5JPD3/p1748637856529199?thread_ts=1748631352.782589&cid=C087VC5JPD3)
    - [ ] Ensure my PTO is correct for Q3 [coda](https://coda.io/d/_dCG_1QsXKay/Resources_sucdplpI#Team-Members_tuL9qD0Y/r9756) 🛫 2025-06-30 
- [[ENRICH-1487- Add prompt injection guardrail to Smart Query]]
    - Merged
- [[ENRICH-1158- Enhance the Engine enrichment service to support skipping messages by configuration.]]
    - Reviewed the [slack thread](https://qualtrics.slack.com/archives/C08SSBVRXDZ/p1747717362538939) on the failure, captured details in notes
- [[ENRICH-1488- Add max tokens guardrail to Smart Query]]
    - What we have is good enough: [[ENRICH-1488- Add max tokens guardrail to Smart Query#Summary of [Slack Thread](https //qualtrics.slack.com/archives/C08U6PVGW4F/p1748628542375639) Clarification on "Max Tokens Guardrail" for Smart Query (2025-05-30)]]

# Log
- 

# Start of day tasks
- [x] Copy meeting/other schedule to notes ✅ 2025-05-30
- [x] Check email ✅ 2025-05-30
- [x] Check slack ✅ 2025-05-30
- [x] Review blocked tasks ✅ 2025-05-30
- [x] Review TAPP and ENRICH near due jira items (the sauron slack messages) ✅ 2025-05-30
- [x] Spend >5 minutes on a starved task 📅 2025-05-30 ✅ 2025-05-30

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/Status/2025-06-02 Mon.md
---

- Meetings/other
    - none
- [x] Ping Vamsi about scale testing or figure out who's doing it myself ✅ 2025-06-02
    - From [[2025-05-29 Thu - Vamsi 1 on 1#^ee48ba]]
    - I pinged him, said I could pick up this task if he hasn't had time
    - He's checking with Steven and Zeshan
- Wrote a userscript to auto-close slack archive webpages
- Registering for Feedback 360
    - Collected context [[Feedback 360]]
    - Asking peer evaluators if they're willing to
    - Registered Reid, Mavin, Harrison Stark, and Addisu as evaluators (Vamsi is included automatically)
    - Hit the button to send evaluator emails
    - [ ] Complete Feedback 360 self-reflection [link](https://qualtricsei.qualtrics.com/EX/ParticipantPortal) 📅 2025-06-30 🛫 2025-06-15 
- [[ENRICH-1158- Enhance the Engine enrichment service to support skipping messages by configuration.]]
    - Put in 15 minutes
    - Asked if we can run LLMs locally on laptops
        - I have an idea for generating on-the-fly contexts from my obsidian vault. Would be cool if it ran locally. Had this idea when I wanted to load context for this item into my IDE coding AI.
        - Eh, not really worth it. Can only run tiny models on 36 GB RAM.
    - Collected context for writing code
    - Noodled, wrote [[ENRICH-1158- Enhance the Engine enrichment service to support skipping messages by configuration.#How to not cause an OP next time]]
- Investigating TAM b1 deployment failure cycle [slack](https://qualtrics.slack.com/archives/C019UF4CE90/p1748891766911959?thread_ts=1748891661.315629&cid=C019UF4CE90)
    - Manually triggered deployment in spinnaker as Mavin suggested [slack](https://qualtrics.slack.com/archives/C019UF4CE90/p1745983344354679?thread_ts=1745944547.175809&cid=C019UF4CE90), [mavin's manual job](https://spinnaker.eng.qops.net/#/applications/text-analytics-module/executions/details/01JT2B3BY23JZCSKZWDXFGBRS9?stage=0&step=0&details=manualJudgment)
    - Triggered the wrong pipeline (should be Deploy, not CD trigger). Trying again.
    - [x] Write how to fix this in team hub ✅ 2025-06-02
        - done https://coda.io/d/_dt-izYOVJmq/Text-Analytics-Module-TAM_suavHIhZ#_lukRR0pZ
    - Failed because I didn't disable verification. Trying again.
    - Other failure reasons:
        - didn't specify artifact tag? had to dig in splunk to find this one
- [[ENRICH-1532- Rate limiting in Topic Recommendations]]
    - Found the [responsible code](https://gitlab-app.eng.qops.net/xm-discover/discover-platform-group/discover-data-pipelines/platform/-/commit/3c062bb11c97772f3e3533352eba77d97e895ba0)
    - Can't get TR running locally or find relevant server logs
        - Probably missing config in properties file. `client.discovery.trn.*`
        - [x] Ask Addisu how to get TR running locally 🛫 2025-06-13 ✅ 2025-06-17
    - Hm, don't necessarily need to get it running locally.
    - Seems likely that CMP is getting a 429 back from the TR LIG call
    - There's definitely 429 handling in the TR lambda code, but I can't find recent instances (like from the bug bash) in the logs.
    - Found [doc](https://gitlab-app.eng.qops.net/iq-research-sea/model-inference-gateway/-/blob/main/documentation/consumer_docs/rate_limiting.md) on MIG rate limiting
    - Found MIG code with rate limit configs: `throttling_config_smart_query.go`
    - Yep, that explains it. Wrote up my findings in [slack](https://qualtrics.slack.com/archives/C08TL1X7EH3/p1748905033963469?thread_ts=1748539717.676119&cid=C08TL1X7EH3) and updated ticket.
    - I asked if there's any follow-up work. Vamsi had a couple things. I asked if he wants me to start those now.
        - [x] Follow up on [slack](https://qualtrics.slack.com/archives/C08TL1X7EH3/p1748905973886459?thread_ts=1748539717.676119&cid=C08TL1X7EH3) ✅ 2025-06-03
- Asked for overrides for Jun 11 shift
    - [x] Verify all overrides handled 🛫 2025-06-04 ✅ 2025-06-04
- Asked if there's any more work I can pick up for Smart Audit or TR [slack](https://qualtrics.slack.com/archives/C08UJM81L9L/p1748905360770299)

# Log
- 

# Start of day tasks
- [x] Copy meeting/other schedule to notes ✅ 2025-06-02
- [x] Check email ✅ 2025-06-02
- [x] Check slack ✅ 2025-06-02
- [x] Review blocked tasks ✅ 2025-06-02
- [x] Review TAPP and ENRICH near due jira items (the sauron slack messages) ✅ 2025-06-02
- [x] Spend >5 minutes on a starved task 📅 2025-06-02 ✅ 2025-06-02

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/Status/2025-06-03 Tue.md
---

- Meetings/other
    - Enrichments RCA Review and Discussion
- I volunteered to work on [[ENRICH-1396- Unit Tests for Auto Tune]] for Charmish [slack](https://qualtrics.slack.com/archives/C08UJM81L9L/p1748962449463679?thread_ts=1748907643.068309&cid=C08UJM81L9L)
    - [-] [[ENRICH-1396- Unit Tests for Auto Tune]] ❌ 2025-06-09
- [[ENRICH-1158- Enhance the Engine enrichment service to support skipping messages by configuration.]]
    - Added debug logging for headers to see what's available
    - Kicked off build
    - Vamsi moved the due date for this to Jun 20 and said "please prioritize accordingly"
    - Deploying image to staging deployment `plat-classification-service-staging03`
        - The previously deployed image was `799231120095.dkr.ecr.us-east-1.amazonaws.com/plat/classification-service:1190682-MR-193-debug`
    - Eh, I don't know how to generate traffic there. Going to use burgundy with stuff from the TR/SA bug bash.
    - Burgundy was also on `1190682-MR-193-debug`. My branch's tag is `jbrewer-1158-again`
    - Got a great thread from Milo on how to put documents into ingestion gateway, which will trigger CS.
    - Working on a script to publish to IG
        - Having auth issues
    - Milo's being really helpful. I should call it out next sprint review.
    - Got it. Instance name issue, I think. I switched to the Autotune-GA1 project on staging03.
    - Deploying image to staging03.
    - [x] Clean up my deployments to staging03 and burgundy #enrich-1158 ⛔ ei2t3d ✅ 2025-06-18
    - Got it! The logging works and I found the x-* headers as expected.
    - Asked Sandesh which header to read for instance.
- [[ENRICH-1532- Rate limiting in Topic Recommendations]]
    - Asked socrates-ml-platform for rate increase [slack](https://qualtrics.slack.com/archives/C0186S22PQS/p1748964053254049)
        - Submitted survey requesting new issuer for TR as they suggested, with higher limits
        - https://qualtrics.atlassian.net/browse/SOC-7617
    - Writing a script to spam MIG to exceed rate limit and get 429s
        - Not getting limited? Idk, maybe I'm not spamming enough.
    - Suggested they add docs on 429 to their api spec [slack](https://qualtrics.slack.com/archives/C0186S22PQS/p1748969654963479)
    - Seems done? No UX work before GA. I asked if there's anything else to do now.
    - Daniel confirms no more work for this. He created a follow-up [task](https://qualtrics.atlassian.net/browse/ENRICH-1547) to update issuer once that goes through.
    - I volunteered to watch for the new issuer coming through and then completing  [[ENRICH-1547- Update MIG Issuer for Topic Recommender]]
    - [x] [[ENRICH-1547- Update MIG Issuer for Topic Recommender]] 🔼 ✅ 2025-06-11
        - [x] Request PLMS access for new issuer #enrich-1547-update-mig-issuer ✅ 2025-06-10
        - [x] Update code, merge #enrich-1547-update-mig-issuer ✅ 2025-06-11
        - [x] Deploy #enrich-1547-update-mig-issuer ✅ 2025-06-26
        - [x] Verify #enrich-1547-update-mig-issuer ✅ 2025-06-26
    - Closed
- Started [[Gemini Instructions]] to try to steer gemini when I dump my vault into it. It's working!
- [[TAPP-7074- TextIQ Bubble charts content is not perceptible and Bubble charts cannot be operated via keyboard]]
    - I was supposed to go to widget office hours and discuss this but forgot and had a conflict. Moved to Thursday.
    - [-] Attend widget office hours! 🔼 🛫 2025-06-05 ❌ 2025-06-04
- Enrichments RCA Review and Discussion
    - [x] Sandesh said to sync with him on header documentation [doc](https://docs.google.com/document/d/1J7uXpE2Y3JZ3Upuf6kQ40qqHzfwrl_GgOCMmSiGYBrg/edit?disco=AAABlC2eak8) ✅ 2025-06-04
- Checked action items from [[2025-05-27 Tue - Sprint Retro]]
    - Pinged Vamsi about on-call time tracking guidance and increasing sprint review duration to 1.5 hours ^6c334e
    - Mavin said: Plan is 1 hour for now with just demos. We punted everything else.
- Created a personal [repo](https://gitlab-app.eng.qops.net/jbrewer/scratch) for random scripts. Pushed aws-k8s-login and a 429 generating curl script.

# Log
- 

# Start of day tasks
- [x] Copy meeting/other schedule to notes ✅ 2025-06-03
- [x] Check email ✅ 2025-06-03
- [x] Check slack ✅ 2025-06-03
- [x] Review blocked tasks ✅ 2025-06-03
- [x] Review TAPP and ENRICH near due jira items (the sauron slack messages) ✅ 2025-06-03
- [x] Spend >5 minutes on a starved task 📅 2025-06-03 ✅ 2025-06-03

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/Status/2025-06-04 Wed.md
---

# Status
- Meetings/other
    - Scrum
    - Dr appt
- [[ENRICH-1158- Enhance the Engine enrichment service to support skipping messages by configuration.]]
    - I should add Sandesh or another data pipelines engineer to this MR
    - Sent a public slack message to confirm which headers to use [slack](https://qualtrics.slack.com/archives/C035CQ3CHUK/p1749045353686389)
    - Confirmed: x-account-id, x-project-id, x-content-provider
    - [ ] Create splunk alerts for header parse failure since it won't error out now (pending feedback from line below) #enrich-1158 🆔 ad847v ⛔ 9lla51,u3nvy0 🔼
    - Asked the discover community about my "warn and alert" approach to header parsing failures [slack](https://qualtrics.slack.com/archives/C0N0F9087/p1749055630270789)
    - [x] Link this to AI from yesterday's RCA review [[ENRICH-1548- Classification Service update header to parse string header instead of int64]] ✅ 2025-06-04
    - Updated unit tests, trying to get a build for testing
    - Deployed new build to staging03, verified classification succeeds even if headers can't be read
    - Wrote a [runbook section](https://coda.io/d/_dt-izYOVJmq/Classification-Service_suAa9o2-#_luROJmfg) on how to block a project
    - Manual testing
    - Prepared, sent MR
- Fixed my CAA userscript
- Learned about another (?) repo with our lambda functions?
    - https://gitlab-app.eng.qops.net/iq-research-sea/iq-aws-resources-manager
    - whereas I've been working from this repo Addisu sent me https://gitlab-app.eng.qops.net/xm-discover/discover-engines-group/query-topics-ai-integrations
    - .... what is the relationship? Charmish doesn't know. found this [coda doc](https://coda.io/d/Query-Topics-Team-Hub_dt-izYOVJmq/Deployments-of-Lambda-functions_suhaADey) on deployment that mentions the "new" repo
    - [x] Ask Addisu about the two lambda repos 🛫 2025-06-13 ✅ 2025-06-13
- [[ENRICH-1532- Rate limiting in Topic Recommendations]]
    - Pinged DICE about rate limit [ticket](https://qualtrics.atlassian.net/browse/SOC-7617)
        - Their on-call is very busy. I said tomorrow would be fine.

# Log
- [x] Close [[ENRICH-1548- Classification Service update header to parse string header instead of int64]] as dupe? #enrich-1158 ⛔ 9lla51,ei2t3d ✅ 2025-06-18
- I'm kinda mad about finding the second lambda repo this way.
- Mavin is taking 5 hours of my on-call overrides on a weekend

# Start of day tasks
- [x] Copy meeting/other schedule to notes ✅ 2025-06-04
- [x] Check email ✅ 2025-06-04
- [x] Check slack ✅ 2025-06-04
- [x] Review blocked tasks ✅ 2025-06-04
- [x] Review TAPP and ENRICH near due jira items (the sauron slack messages) ✅ 2025-06-04
- [-] Spend >5 minutes on a starved task 📅 2025-06-04 ❌ 2025-06-05

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/Status/2025-06-05 Thu.md
---

# Status
- Meetings/other
    - AI tools for Query-Based Topics
    - Query Topics Q3 preliminary review
    - ~~Vamsi 1 on 1~~
- [[ENRICH-1158- Enhance the Engine enrichment service to support skipping messages by configuration.]]
    - Shared with reviewers a [link](https://gitlab-app.eng.qops.net/xm-discover/discover-engines-group/classification-service/-/compare/538a3aa..249fa1f004f8aea02f6d4775998156f879806d3b) to diff just what changed between the original 1158 change and my enhanced 1158 change.
    - Asked Sandesh to skim this MR
    - Addressing MR feedback
    - Merged
    - [x] Deploy CS with 1158 to prod #enrich-1158 🆔 ei2t3d ⛔ w9xthe,1gwcyz 🔼 ✅ 2025-06-18
    - [ ] Verify splunk alert for header parse error #enrich-1158 🆔 w9xthe ⛔ ad847v 🔼
    - [ ] Verify dlq reprocessing works (could block messages then trigger reprocessing) #enrich-1158 🆔 1gwcyz ⛔ u3nvy0 🔼
        - I can request a dlq run with a workflow in slack channel discover-pipelines-community
- Called attention to near-due unassigned item [[ENRICH-1453- Arcadia Alert retry for Auth Server 500 errors]] [slack](https://qualtrics.slack.com/archives/C087VC5JPD3/p1749137458209919?thread_ts=1749128580.684729&cid=C087VC5JPD3)
- AI tools for Query-Based Topics
    - [x] document rate limit in a new tab [here](https://coda.io/d/Text-Analytics-Enrichments-Hub_dAG-EE8cgH8/Scoped-requirements_suOjynu-), also how scoping of rate limit works. is it just by user? brand? feature? ✅ 2025-06-06
    - [x] look into getting a documentation update (support page) for the rate limit error message ✅ 2025-06-06
- Vamsi said he was going to update days off in the planning coda. I'd already updated mine. I shared my calculation [slack](https://qualtrics.slack.com/archives/C087VC5JPD3/p1749146722905979?thread_ts=1749145450.517879&cid=C087VC5JPD3)
- [x] Prepare for 1:1 ⏫ ✅ 2025-06-05
- [x] Figure out how to see my windsurf/cursor usage [slack](https://qualtrics.slack.com/archives/C6QL0Q0LT/p1749142638912419?thread_ts=1749142540.401619&cid=C6QL0Q0LT) ✅ 2025-06-05
    - Put info in [[AI coding]]
- [[Researching MIG rate limiting]]
    - Found dashboard for usage against rate limits
    - Wait... how do the lambdas in discover DCs call MIG? I wouldn't think MIG is deployed there.
        - Designer calls go through Stargate LBaC to LIG (?) [slack](https://qualtrics.slack.com/archives/C036JC1JKGE/p1747666883234619?thread_ts=1747385391.353109&cid=C036JC1JKGE)
    - Rate limit doc says the "throttler" can be issuer, brand, or user. Default is issuer.
    - Asking how to tell which throttler is being used [slack](https://qualtrics.slack.com/archives/C0186S22PQS/p1749147003594579)
    - Dug in code, figured out how different throttlers are configured, wrote [[Model Inference Gateway (MIG)#Rate limiting]]
    - [x] Asked DICE if they want an update to their rate limit doc with the surprises I've found [slack](https://qualtrics.slack.com/archives/C0186S22PQS/p1749234595352669?thread_ts=1749147003.594579&cid=C0186S22PQS) ✅ 2025-06-09
- [[Doc update for TR rate limit error message]]
    - Found [LRR ticket](https://qualtrics.atlassian.net/browse/LAUNCH-45402) for docs/digital support
    - Found existing [support page](https://www.qualtrics.com/support/xm-discover/designer/categorize/category-models/editing-category-models-designer/#Recommendation) for TR... unexpected. Asking about current launch state [slack](https://qualtrics.slack.com/archives/C07MKL3FB61/p1749151459324559)
        - Donovan says "looks like we are ahead of the game. I wouldnt worry about the fact that its on the support page"
    - Was waiting to hear a GA date for TR before filing, didn't.
- [[ENRICH-1396- Unit Tests for Auto Tune]]
    - Reviewing the linked MR... it's *big*. Asking what areas we want unit tests for. [slack](https://qualtrics.slack.com/archives/C07MKL3FB61/p1749152338518429)
- [[ENRICH-1547- Update MIG Issuer for Topic Recommender]]
    - Confused about which repo to work in...
    - Maybe see if lambdas are deployed with the `-tf` suffix mentioned in docs?
    - Not in fra1 and I don't have lambda access in discover
        - [x] Are these lambdas deployed in discover? How to get access? [slack](https://qualtrics.slack.com/archives/C07MKL3FB61/p1749162967383299) ✅ 2025-06-06
            - (no)
    - Ok I found `coreml-socrates-topic-recommendation-tf` in b1. Let me check the tf file for TR.
        - That's the only DC it's deployed to via tf.
    - Suppose I can just work from the qt repo to test the new issuer while I wait on info.
- [x] Investigate--my vault backup hasn't been running ✅ 2025-06-06
- [x] Find a backup solution for .zshrc and such ✅ 2025-06-06
# Log
- My days off for Q3 in planning [coda](https://coda.io/d/PXE-Hub_dCG_1QsXKay/Team-Planning_sucrg821) was wrong again. Had 10. Mavin clarified that company holidays and hackathon should be included. So for me, 5 sick + 7 vacation + 2 holidays + 2 hackathon = 16. I updated it myself.
    - I think he didn't understand that my sick days were known/scheduled. We clarified in the above linked thread and he says he understands now.
- Vamsi's rescheduling our 1:1

# Start of day tasks
- [x] Copy meeting/other schedule to notes ✅ 2025-06-05
- [x] Check email ✅ 2025-06-05
- [x] Check slack ✅ 2025-06-05
- [x] Review blocked tasks ✅ 2025-06-05
- [x] Review TAPP and ENRICH near due jira items (the sauron slack messages) ✅ 2025-06-05
- [-] Spend >5 minutes on a starved task 📅 2025-06-05 ❌ 2025-06-06

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/Status/2025-06-06 Fri.md
---

# Status
- Meetings/other
    - Scrum
    - Ops
- [[Doc update for TR rate limit error message]]
    - Filed https://qualtrics.atlassian.net/browse/DSX-26721
- ~~Seems SA lambdas are deployed to discover DCs~~? Found Charmish's item [[ENRICH-1398- Deploy Auto Tune lambda functions in each datacenter]]
    - Pinged him about this [slack](https://qualtrics.slack.com/archives/C07MKL3FB61/p1749243646186119?thread_ts=1749162967.383299&cid=C07MKL3FB61)
    - No, he's deploying the lambdas to engage DCs that will *serve* designer in discover DCs. Summarized in [[Query Topics AI Features (lambdas)#Summary of Lambda Deployments]]
- QED
    - Set up obsidian git plugin for vault backup/version control
    - Created system for version controlling dotfiles
        - ~~with secret detection!~~
        - Too much hassle, doesn't detect stuff like qtoken that aren't in its database. Tried trufflehog.
- [[Researching MIG rate limiting]]
    - Compiled findings in [coda](https://coda.io/d/_dAG-EE8cgH8/Scoped-requirements_suOjynu-#_luGJ9SjP)
        - Specifically called out in coda and [slack](https://qualtrics.slack.com/archives/C07MKL3FB61/p1749243556702669) that rate limiting is per-issuer, not per-user as I'd suggested before
    - Sent DICE an [MR](https://gitlab-app.eng.qops.net/iq-research-sea/model-inference-gateway/-/merge_requests/1115) to improve their rate limit document #cross-team-impact
- [[ENRICH-1547- Update MIG Issuer for Topic Recommender]]
    - Got TR requests working from rapidapi
    - Did a simple find/replace for issuer, deployed
    - PLMS doesn't like the new issuer
    - [ ] Request issuers for each AI feature for MIG, have each allowed in PLMS 🔽

# Log
- Scrum
    - Vamsi says after I finish [[ENRICH-1396- Unit Tests for Auto Tune]] (wait, this is deprioritized?) and [[ENRICH-1158- Enhance the Engine enrichment service to support skipping messages by configuration.]], check in and see what other pressing AI feature (like SA, TR) work I can take.
    - [x] Note what alerts i'm adding to [[ENRICH-1158- Enhance the Engine enrichment service to support skipping messages by configuration.]] ✅ 2025-06-06
    - Vamsi says deprioritize [[ENRICH-1396- Unit Tests for Auto Tune]] because even if we write the tests now, it'll go in the July CMP release.

# Start of day tasks
- [x] Copy meeting/other schedule to notes ✅ 2025-06-06
- [x] Check email ✅ 2025-06-06
- [x] Check slack ✅ 2025-06-06
- [x] Review blocked tasks ✅ 2025-06-06
- [x] Review TAPP and ENRICH near due jira items (the sauron slack messages) ✅ 2025-06-06
- [-] Spend >5 minutes on a starved task 📅 2025-06-06 ❌ 2025-06-09

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/Status/2025-06-09 Mon.md
---

---
created: 2025-06-09T08:29:51
---
# Status
- Meetings/other
    - [[2025-06-09 Mon - Vamsi 1 on 1]]
- Asked about discover test accounts [slack](https://qualtrics.slack.com/archives/C087VC5JPD3/p1749483906696669)
- [ ] Would be nice to generate api specs for AI lambda services and generate clients 🔽
    - Asked if we have any API specs [slack](https://qualtrics.slack.com/archives/C07MKL3FB61/p1749485443999699)
        - No response
- [x] Send status update since I missed scrum ✅ 2025-06-09
    - [x] Add a monday template item to send status since it's regular ✅ 2025-06-09
- [x] Prepare for 1 on 1 ✅ 2025-06-09
- [[ENRICH-1547- Update MIG Issuer for Topic Recommender]]
    - Requested access to PLMS for new issuer [slack](https://qualtrics.slack.com/archives/C0186S22PQS/p1749499233081219)
- [[ENRICH-1463- Iterate on prompt to reduce wildcard use]]
    - Begin
- Filed a ticket on technical issues with Feedback 360 [link](https://qualtrics.service-now.com/nav_to.do?uri=sn_hr_core_case.do%3Fsys_id=10626f111b8e6250ce78c882604bcb63%26sysparm_stack=sn_hr_core_case_list.do%3Fsysparm_query=active=true)
- Checking if Feedback 360 evaluators got notifications
- [[ENRICH-1568- Classification Service- Add sharding migration disabler]]
    - Created this item based on discussion in [[2025-06-09 Mon - Vamsi 1 on 1]]
    - Researched "empty shard list" vs "disable code with env var" approaches, chose the latter.
    - Implemented
    - [x] Create a jira task to clean this up when no longer needed, *if* this approach is accepted and code merged #enrich-1568-migration-disabler 🛫 2025-06-16 ✅ 2025-06-26
    - Wrote MR, sent for review
    - [x] Investigate adding unit tests for this #enrich-1568-migration-disabler 🔼 ✅ 2025-06-10
    - [x] Shepherd MR, merge #enrich-1568-migration-disabler 🆔 u3nvy0 🔼 ✅ 2025-06-13

# Log
- It frustrates me that we don't have regular recurring calendar items for stuff like sprint planning, and that Thursday's async scrum is haphazardly, inconsistently updated. #vamsi/feedback
- I see a pulse that's due soon. I tagged Vamsi and someone else involved with it, saying it's due soon and asking if we need to pay attention ([slack](https://qualtrics.slack.com/archives/C087VC5JPD3/p1749137458209919?thread_ts=1749128580.684729&cid=C087VC5JPD3)). No response. What do I do? I don't just want to keep tagging Vamsi.  ^db6ed8
- Oo deploying the second implementation of [[ENRICH-1158- Enhance the Engine enrichment service to support skipping messages by configuration.]] is going to get blocked for US prod deployment if we're still waiting on scale testing for classification. Notified Vamsi in slack.
- Steven is working on scale testing for CS sharding [slack](https://qualtrics.slack.com/archives/C087VC5JPD3/p1749503916698549)
- Asked Harrison if he got an email about Feedback 360

# Start of day tasks
- [x] Copy meeting/other schedule to notes ✅ 2025-06-09
- [x] Check email ✅ 2025-06-09
- [x] Check slack ✅ 2025-06-09
- [x] Review blocked tasks ✅ 2025-06-09
- [x] Review TAPP and ENRICH near due jira items (the sauron slack messages) ✅ 2025-06-09
- [-] Spend >5 minutes on a starved task 📅 2025-06-09 ❌ 2025-06-10

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/Status/2025-06-10 Tue.md
---

---
created: 2025-06-10T08:04:31
---
# Status
- Meetings/other
    - Accessibility ticket meeting
    - Sprint review
- [[TAPP-7074- TextIQ Bubble charts content is not perceptible and Bubble charts cannot be operated via keyboard]]
    - [x] Find out what's expected of me in Accessibility ticket meeting #tapp-7074-bubble-a11y ✅ 2025-06-10
        - Asked [slack](https://qualtrics.slack.com/archives/C08RHU9K8NS/p1749569626526509)
    - Asked Steph and Jonathan who's the UX contact [slack](https://qualtrics.slack.com/archives/C08RHU9K8NS/p1749590042193979)
        - Steph sent an old a11y spec that looks like it'll handle the issues
        - Steph says UX contact is Jonathan
    - Brainstorming approaches
    - Asked Claire if it's enough to fix the two a11y issues, or if the customer will insist we implement their recommendations [slack](https://qualtrics.slack.com/archives/C090NU5VDNJ/p1749593020069169)
        - To save roundtrips, also asked in thread:
        - > Like, would the customer be satisfied if we did something like this? These are the highlights of the a11y spec. The first shows a new "view data" button. The second is a small part of the full keyboard navigation design.
        - > Oo, also, do you know who could we talk to about the consequences of 7074 _not_ being completed in time for contract negotiation? There have been the beginnings of conversations about priority, and I'd like to go ahead and have those contacts ready 🤗
- [x] Prepare for sprint review (focus on demos) ✅ 2025-06-10
    - Misunderstood the new sprint review format. No status. Only demos. Mavin says: ^cd1392
    - *> use the slides for demos.  You can include screenshots and stuff.  But sprint status shouldn't be included. Example of demos:*  
        - *> where you are in a feature and a blocker you want team feedback on for next steps*
        - *> new feature you added to show team how it works for awareness*
        - *> potential updates that could potentially see pulses*
        - *> KTLO improvements for awareness*
    - Wasted time making slides
    - Making demos
    - Shared my concern that with this change to focus on demos, we have just sprint retro as an outlet for issues
- [[ENRICH-1547- Update MIG Issuer for Topic Recommender]]
    - Pinged socrates about PLMS access
- Checked with Feedback 360 reviewers about having been notified, verified Vamsi got the notification this time
    - Harrison's having technical issues with the tool. I pointed him to the same myq request type I filed yesterday.
- Pinged DICE about my rate limiting doc update MR
    - Even *they* don't all understand low/high watermark 😆😭
- [[ENRICH-1568- Classification Service- Add sharding migration disabler]]
    - Investigated adding unit tests, wrote a comment in MR about why I think it's not worth it
- Pinged Vamsi about on-call schedule update for Q3 [slack](https://qualtrics.slack.com/archives/C087VC5JPD3/p1749571507019389?thread_ts=1749484319.401279&cid=C087VC5JPD3)
    - It hasn't been decided yet. Asked for resolution sooner than later.
    - [x] Ask for overrides for Jul 9-16 shift 🛫 2025-06-18 ✅ 2025-06-26
- Accessibility ticket meeting
    - Liam said that completing [[TAPP-7074- TextIQ Bubble charts content is not perceptible and Bubble charts cannot be operated via keyboard]] is a precondition for transferring widgets to his team
    - [x] Ask for UX design/input (Rohit: Jonathan Davito? current UX for text iq. or steph) #tapp-7074-bubble-a11y ✅ 2025-06-10
    - [x] Try to scope the work. How long might it really take? #tapp-7074-bubble-a11y 📅 2025-06-13 ✅ 2025-06-11
        - I said I'd try to have this done this week
        - (later: filed [[ENRICH-1581- Spike- Scope widget pulse a11y work and break into smaller tasks]], calling this task done)
        - [-] Attend Thu widget office hours to check in on scoping #tapp-7074-bubble-a11y 📅 2025-06-12 ❌ 2025-06-12
            - Rohit says if it's not doable, he's happy to start the threads on that
    - Liam wants a contingency plan for if we can't punt this and can't delivery it by its due date Jul 18
    - If I can't find time to do this in second half of Jun, I need to say so
    - Maybe consult James Rupke
- Sprint review
    - [x] Follow-up: what kind of error message is appropriate for prompt injection? ✅ 2025-06-10
- Asked Vamsi about scheduling June retro. He says he's setting it up now.
- [[ENRICH-1295- Designer Error Handling for smart-query guardrails]]
    - Asked Donovan if we want a specific error message for prompt injection, or treat it as generic error [slack](https://qualtrics.slack.com/archives/C07MKL3FB61/p1749583868026249)
        - Leo says show it. We assume good actors.
- [[ENRICH-1547- Update MIG Issuer for Topic Recommender]]
    - DICE says TR is allowed for PLMS now
    - Tested new code, works
    - Sent MR

# Log
- I feel very scattered. Five things going on at once. Not good for me.
- [ ] Automate dotfiles commit and push #qed
- [x] Subscribe to the UW Aetna slack thread ✅ 2025-06-10
- I'm stressed out about [[TAPP-7074- TextIQ Bubble charts content is not perceptible and Bubble charts cannot be operated via keyboard]]. Lots of unknowns and pressure.
    - Fuck I forgot I'm on call starting tomorrow
- Sent Vamsi:
    - *> Hey, I didn't realize I was starting on-call on Wed when we met yesterday. That's really going to cut into my bandwidth. Combined with that, I met with the new widgets team and Rohit about https://qualtrics.atlassian.net/browse/TAPP-7074 today and they're pushing hard for me to get on it. I told them I'd have scoping done this week before I remembered I'm going on-call.*
    - *> Switching between all these different work streams is difficult and stressful for me, as is being on-call for this team.*
    - *> So, I don't know that any action is needed, but I wanted to share all this.*
- Asked Vamsi if we should change [[ENRICH-1468- Smart query version 2 in designer (epic)]]'s status as a Q2 deliverable in the planning coda, since it was deprioritized in favor of other AI features. [slack](https://qualtrics.slack.com/archives/D087U08TEAX/p1749660457147289?thread_ts=1749590177.594929&cid=D087U08TEAX)
- [x] Copy Tuesday's slack DMs with vamsi into here for 7074 context. 🛫 2025-06-11 ✅ 2025-06-11
- [ ] Write a "slack thread cleaner" to clean up pastes from slack UI into something nice for adding to my vault #qed

# Start of day tasks
- [x] Copy meeting/other schedule to notes ✅ 2025-06-10
- [x] Check email ✅ 2025-06-10
- [x] Check slack ✅ 2025-06-10
- [x] Review blocked tasks ✅ 2025-06-10
- [x] Review TAPP and ENRICH near due jira items (the sauron slack messages) ✅ 2025-06-10
- [-] Spend >5 minutes on a starved task 📅 2025-06-10 ❌ 2025-06-11

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/Status/2025-06-11 Wed.md
---

---
created: 2025-06-11T08:29:11
---
# Status
- Meetings/other
    - Sprint planning
    - Dr appt
    - Reid / Janzen on-call sync
- On-call
    - [x] Figure out what's going on with [[ENRICH-1453- Arcadia Alert retry for Auth Server 500 errors]] #on-call ✅ 2025-06-11
        - Pinged Vamsi and Charmish again [slack](https://qualtrics.slack.com/archives/C087VC5JPD3/p1749656717854199?thread_ts=1749128580.684729&cid=C087VC5JPD3)
        - 🤷 I tried. I give up.
    - Asked if anyone's had CS queue depth alerts find real issues lately [slack](https://qualtrics.slack.com/archives/C087VC5JPD3/p1749656578045069)
        - Charmish is going to look into increasing threshold
    - Triaged pulses
    - Investigating alert `error creating index: dial tcp [::1]:9200`
        - Restarting pods, looking better
    - [x] [[ENRICH-1573- Text iQ Topics don't have any queries]] #on-call 📅 2025-06-25 ✅ 2025-06-13
        - Investigating
        - Almost positive it's WAD. The highlights are for the selected topic.
        - Confirmed the given record was manually tagged
        - Updated ticket, waiting for res feedback before closing
    - Investigated CS error splunk alert. Seems to have self-resolved.
- [x] Prepare for sprint planning ⏫ ✅ 2025-06-11
- Pinged Vamsi about on-call time tracking [slack](https://qualtrics.slack.com/archives/C087VC5JPD3/p1749657735995599?thread_ts=1748978675.169079&cid=C087VC5JPD3)
    - previously: [[2025-06-03 Tue#^6c334e]]
- Pinged Vamsi about updating next sprint review calendar item to include retro [slack](https://qualtrics.slack.com/archives/C087VC5JPD3/p1749659277373279?thread_ts=1749583723.503769&cid=C087VC5JPD3)
    - He updated it
- Pinged Vamsi about updating Q2 deliverable status of [[ENRICH-1468- Smart query version 2 in designer (epic)]], asked yesterday
    - He says he'll update it to off-track now [slack](https://qualtrics.slack.com/archives/D087U08TEAX/p1749660482334529?thread_ts=1749590177.594929&cid=D087U08TEAX)
- [[ENRICH-1547- Update MIG Issuer for Topic Recommender]]
    - Merged
- Sprint review
    - [x] Make spike to investigate 7074, load into sprint ✅ 2025-06-11
        - [b] [[ENRICH-1581- Spike- Scope widget pulse a11y work and break into smaller tasks]] #tapp-7074-bubble-a11y 🔼
    - [x] Set up meeting with rohit, liam, vamsi about ownership of 7074 ✅ 2025-06-11
        - Notified everyone in [slack](https://qualtrics.slack.com/archives/C090NU5VDNJ/p1749664645718109) too
- [[ENRICH-1568- Classification Service- Add sharding migration disabler]]
    - Worked on MR feedback
- [[TAPP-7074- TextIQ Bubble charts content is not perceptible and Bubble charts cannot be operated via keyboard]]
    - The TAM says we don't have to implement the customer's suggestions; it's enough to adhere to a11y specs [slack](https://qualtrics.slack.com/archives/C090NU5VDNJ/p1749635495136589?thread_ts=1749593020.069169&cid=C090NU5VDNJ)
    - No one answered my question about "what happens if we don't do this". Asked again.
- Filed [[ENRICH-1583- TAM- No space between "contains" and topic button]]

# Log
- I see a pattern with Vamsi where if I ask about two things, he'll respond about one but not the other. #vamsi/feedback 
- It adds to my load that I know anything I ask Vamsi about/to do, I need to set a reminder to follow up about it. #vamsi/feedback 
- Sprint review
    - Vamsi pushed back against my including [[TAPP-7074- TextIQ Bubble charts content is not perceptible and Bubble charts cannot be operated via keyboard]] in sprint, and against us doing much work at all on it. I felt I was repeating exactly what I sent him yesterday in slack DMs and said so. I felt like he was trying to convince me that 7074 isn't ours; I tried a couple times to tell him clearly that *I'm* not who he needs to convince of that. He asked me to set up a meeting with other stakeholders.
    - I felt myself escalating toward red zone during this exchange. I'm frustrated that I've given him this information, probably multiple times, and he doesn't get it. I feel we're going in circles. That's why I like to do things over text; I have the paper trail.

# Start of day tasks
- On-call
    - [x] Copy yesterday's final notes into coda #on-call ✅ 2025-06-11
    - [x] Create tasks for overnight alerts needing investigation #on-call ✅ 2025-06-11
    - [x] Triage pulses #on-call ✅ 2025-06-11
- Regular
    - [x] Copy meeting/other schedule to notes ✅ 2025-06-11
    - [x] Check email ✅ 2025-06-11
    - [x] Check slack ✅ 2025-06-11
    - [x] Review blocked tasks ✅ 2025-06-11
    - [x] Review TAPP and ENRICH near due jira items (the sauron slack messages) ✅ 2025-06-11


---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/Status/2025-06-12 Thu.md
---

---
created: 2025-06-12T06:00:06
---
# Status
- Meetings/other
    - TextiQ Widgets Office Hours
    - AI tools for Query-Based Topics
    - [[2025-06-12 Thu - Vamsi 1 on 1]]
- On-call
    - [x] [[ENRICH-1569- Remove default textIQ fields]] #on-call ✅ 2025-06-12
        - Found via sauron slack message. It seems to have been created as a task, not a bug, so it didn't show up on our pulse board.
        - Done
    - [-] Investigate EOF errors with models and scorecards #on-call ❌ 2025-06-23
        - `error get the models and score cards: error getting score card: error calling model service: error getting response from instrumented client: Get [...] EOF`
        - [slack](https://qualtrics.slack.com/archives/C08EHNYHNFR/p1749711963894709?thread_ts=1749711690.309979&cid=C08EHNYHNFR)
        - Could be [this issue](https://qualtrics.slack.com/archives/C053MK619T9/p1749584477501469?thread_ts=1749582168.995659&cid=C053MK619T9) Reid warned me about
        - Yep, there's a spike in EOF errors today at 12:46 am PT, all for the same project Reid saw an issue with: `42324823`
        - Asked Reid if he filed any issues about this. He didn't.
    - [-] Investigate [[DGD-1631- Fannie Mae - Organization Hierarchy Volume Discrepancy - Categorized Records are not displaying in Studio]] (documents not having classification field) [slack](https://qualtrics.slack.com/archives/C090PSZFN3F/p1749722596418559?thread_ts=1749547869.796249&cid=C090PSZFN3F) #on-call ❌ 2025-06-17
        - Asked for help [slack](https://qualtrics.slack.com/archives/C087VC5JPD3/p1749735599278649)
        - Milo gave some info, asking res to help [slack](https://qualtrics.slack.com/archives/C090681KRN2/p1749742447716359)
        - Asked for more help [slack](https://qualtrics.slack.com/archives/C087VC5JPD3/p1749755532212189?thread_ts=1749735599.278649&cid=C087VC5JPD3)
    - [x] Triage [[TAPP-7092- TextiQ expands past page.]] ✅ 2025-06-12
    - [[ENRICH-1573- Text iQ Topics don't have any queries]]
        - Investigated a way for res to tell whether a record was manually tagged to a topic. Couldn't figure one out.
        - Closed WAD
    - Helping Dongpeng with some MRs [slack](https://qualtrics.slack.com/archives/C8X9E12QY/p1749761709499239)
        - Found that gov1 has an outdated texteng-k8s image. Asking last deployer, Mavin, if it's okay to deploy to gov1 [slack](https://qualtrics.slack.com/archives/C087VC5JPD3/p1749764526659659)
            - [-] Research image consistency alerts. Thought we had those... #on-call 🔽 ❌ 2025-06-23
            - [-] Update deployment section in texteng readme (might just be to remove it, since it's on spinnaker) #on-call 🔽 ❌ 2025-06-23
        - I told Dongpeng that I can deploy TAM and texteng for him next week
    - Suggested documenting the roles and responsibilities of primary and secondary on-call, volunteered to do so [slack](https://qualtrics.slack.com/archives/C087VC5JPD3/p1749765680577199?thread_ts=1749763974.282259&cid=C087VC5JPD3)
        - [x] Follow up on on-call doc suggestion #on-call 🔽 ✅ 2025-06-23
- [[ENRICH-1568- Classification Service- Add sharding migration disabler]]
    - Waiting on MR responses
- [[Doc update for TR rate limit error message]]
    - https://qualtrics.atlassian.net/browse/DSX-26721
    - Kendall notified me that this is complete
- Updated [[QRN-DFS migration]] [tracker](https://docs.google.com/spreadsheets/d/1Mx6--rVvs9CReqZWGWZ1Nmu3lEs3-784bSzyfSsCEP4/edit?usp=sharing) to point text-analytics-module issuer to ENRICH, and team contact to Vamsi. Tagged Vamsi on a comment saying I'd updated that entry.
    - Asked if anyone knows about this migration [slack](https://qualtrics.slack.com/archives/C087VC5JPD3/p1749749418082019)
    - [x] Follow up on qrn migration 🛫 2025-06-13 ✅ 2025-06-16
- [[TAPP-7074- TextIQ Bubble charts content is not perceptible and Bubble charts cannot be operated via keyboard]]
    - Office hours was full and I joined late
    - [x] Send update on status to the channel #tapp-7074-bubble-a11y ✅ 2025-06-12


# Log
- Vamsi:
    - *> hi Janzen , let us deprioritize the widget pulse \[[[TAPP-7074- TextIQ Bubble charts content is not perceptible and Bubble charts cannot be operated via keyboard]]] until we have a connect with Rohit*
    - *> can we bring the smart query story \[[[ENRICH-1463- Iterate on prompt to reduce wildcard use]]] in to the sprint? it didn't come up in planning*
- Proposed new time for next week's 1:1

# Start of day tasks
- On-call
    - [x] Copy yesterday's final notes into coda #on-call ✅ 2025-06-12
    - [x] Create tasks for overnight alerts needing investigation #on-call ✅ 2025-06-12
    - [x] Triage pulses #on-call ✅ 2025-06-12
- Regular
    - [x] Copy meeting/other schedule to notes ✅ 2025-06-12
    - [x] Check email ✅ 2025-06-12
    - [x] Check slack ✅ 2025-06-12
    - [x] Review blocked tasks ✅ 2025-06-12
    - [x] Review TAPP and ENRICH near due jira items (the sauron slack messages) ✅ 2025-06-12


---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/Status/2025-06-13 Fri.md
---

---
created: 2025-06-13T07:56:01
---
# Status
- Meetings/other
    - Scrum
    - Ops
- On-call
    - [x] See if fra1 crate host patching is stuck [slack](https://qualtrics.slack.com/archives/C010USY4MPX/p1749747617068299) #on-call ✅ 2025-06-13
        - Yep. Killed the rundeck job.
    - [x] Patch crate fra1 #on-call ✅ 2025-06-16
        - Verified cluster is green
        - Was reading our docs on handling hung host patching jobs. I don't think we need to manually patch the remaining nodes. Discussion in [slack](https://qualtrics.slack.com/archives/C087VC5JPD3/p1749851380547219)
        - Started [host patch job](https://rundeck-data-platform.eng.qops.net/execution/show/29496932#nodes) with target nodes specified
        - Bleh, the target nodes didn't stick or didn't get passed along.
        - [-] Update [coda](https://coda.io/d/_dt-izYOVJmq/Info-Automated-Host-Patching_suIVnI-U) if this crate patch method works #on-call 🔽 ❌ 2025-06-23
    - [x] Investigate THG error pulse (P1, due Jun 24) [slack](https://qualtrics.slack.com/archives/C8X9E12QY/p1749770850060849?thread_ts=1749657171.325069&cid=C8X9E12QY) #on-call ✅ 2025-06-16
    - [x] Check vulns #on-call 🔼 ✅ 2025-06-13
        - Team hub doesn't have correct vmx link, and nothing for gov1, asked team [slack](https://qualtrics.slack.com/archives/C087VC5JPD3/p1749832628643829)
        - Milo and Mavin shared their links
        - Updated team hub with correct links
        - Filed [[ENRICH-1589- Resolve gov1 vulnerability with template-service]], SLA Jul 3
        - fra1 crate hosts are the only hosts due soon, and they've been patching but maybe stuck, already have a task to investigate
        - Filed [[ENRICH-1591- Update crate image with vuln fixes]], SLA Jun 24. Taking this myself since it's so soon.
            - [x] [[ENRICH-1591- Update crate image with vuln fixes]] #on-call 🔼 📅 2025-06-24 ✅ 2025-06-16
        - Asked team about plat-alert-service being far out of SLA
            - Milo says it's an intern project. He'll handle it.
        - Asked team about react-studio-console being far out of SLA
            - Milo says he transferred ownership of it yesterday and started a rescan
    - [x] [[ENRICH-1590- [CY] Alerts did not complete successfully - US Bank]] #on-call 📅 2025-06-27 ✅ 2025-06-17
        - Looks an awful lot like [[ENRICH-1436- [CY] Arcadia alert failed with 500 - API request was not logged - US Bank]]. Asked Charmish [slack](https://qualtrics.slack.com/archives/C090681KRN2/p1749850422212959?thread_ts=1749835526.744149&cid=C090681KRN2)
    - Asked for help again with documents not having classification field [slack](https://qualtrics.slack.com/archives/C087VC5JPD3/p1749830752109079?thread_ts=1749735599.278649&cid=C087VC5JPD3)
        - Milo might be able to later
    - What is our vuln workflow? I found three jira items in the [Q2 KTLO epic](https://qualtrics.atlassian.net/browse/ENRICH-1513)
- [x] Review [End-of-Life CrateDB Risk Mitigation](https://docs.google.com/document/d/1cBEI4BjM4EDFe4m4X5xJQ4mOcBlN1wDj1SJKMNCOvrY/edit?tab=t.0#heading=h.pzk47dnki5ar) 📅 2025-06-13 ✅ 2025-06-13
- Discussing which repo to use for lambdas with Addisu
- [[ENRICH-1568- Classification Service- Add sharding migration disabler]]
    - Gave up on getting Charmish to respond, resolved all threads
    - Pipeline being weird
    - Merged
- [ ] Add manual tag tooling to text tools [jira](https://qualtrics.atlassian.net/browse/ENRICH-1573?focusedCommentId=11592522) #qed

# Log
- I'm very frustrated at the state of on-call. There's so little organization.
- Scrum
    - Reid raised the [[QRN-DFS migration]] and said we'll need to plan on doing it in Q3. Vamsi discussed it with him.
- [ ] have some kind of jira status syncer that loads history into jira item notes. thought of this while working on sprint loading. jira has export to xml. #qed 

# Start of day tasks
- Regular
    - [x] Copy meeting/other schedule to notes ✅ 2025-06-13
    - [x] Check email ✅ 2025-06-13
    - [x] Check slack ✅ 2025-06-13
    - [x] Review blocked tasks ✅ 2025-06-13
    - [x] Review TAPP and ENRICH near due jira items (the sauron slack messages) ✅ 2025-06-13
- On-call
    - [x] Copy yesterday's final notes into coda #on-call ✅ 2025-06-13
    - [x] Create tasks for overnight alerts needing investigation #on-call ✅ 2025-06-13
    - [x] Triage pulses #on-call ✅ 2025-06-13


---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/Status/2025-06-16 Mon.md
---

---
created: 2025-06-16T07:18:01
---
# Status
- Meetings/other
    - ~~Juneteenth Keynote Conversation~~
        - too busy
    - ~~The Serviceberry Coffee & Discussion~~
        - postponed
    - Discussion on ownership of TAPP-7074
        - Rohit's going to ask PM/Hardik about how to prioritize this pulse/feature request since no one has time for it.
        - [x] Put the "good news" I shared with widget people in the 7074 ticket ✅ 2025-06-16
- On-call
    - [[DGD-1631- Fannie Mae - Organization Hierarchy Volume Discrepancy - Categorized Records are not displaying in Studio]]
        - sbielowka asked if they could move this to ENRICH. I said I couldn't suggest otherwise and reminded Vamsi that I need help on this ticket [slack](https://qualtrics.slack.com/archives/C090PSZFN3F/p1750077206442919)
    - fra1 crate host patch finished Saturday morning
    - [x] [[TCP-476- [CY] Centene Corporation - Topic Hierarchy Generator Errors]] #on-call ✅ 2025-06-23
        - Investigating
        - Found an interesting error log, consulting TCP [slack](https://qualtrics.slack.com/archives/C8X9E12QY/p1750114566848909?thread_ts=1749657171.325069&cid=C8X9E12QY)
    - [[ENRICH-1591- Update crate image with vuln fixes]]
        - Started, verified image rebuild gets vuln fixes. Added a couple niceties to the publish scripts.
        - Sent [MR](https://gitlab-app.eng.qops.net/text-engine/crate-dbms/-/merge_requests/116)
        - Merged 
        - Built and published images
            - `prod-registry.eng.qops.net:5001/query-topics/crate:dbms4.0.12.37`
        - [x] Let crate .37 bake in b1 #on-call 🆔 xh9h8v 🔼 ✅ 2025-06-23
            - Started host patch job after merging
        - [x] Let crate .37 bake in g1 #on-call 🆔 a16icc ⛔ xh9h8v 🔼 ✅ 2025-06-26
        - [x] Update crate release branch with .37 #on-call ⛔ a16icc 🔼 ✅ 2025-06-26
        - Wrote [doc](https://coda.io/d/_dt-izYOVJmq/Runbook-Deploying-a-new-version-of-the-CrateDB-image_suMo6phw) for deploying crate in team hub
    - [x] Host patch crate gov1 #on-call ✅ 2025-06-18
    - [[ENRICH-1590- [CY] Alerts did not complete successfully - US Bank]] was triaged to studio
- [x] Addisu wants me to copy the [[ENRICH-1547- Update MIG Issuer for Topic Recommender]] change to the deployment repo [jira](https://qualtrics.atlassian.net/browse/ENRICH-1547?atlOrigin=eyJpIjoiOTM4NmIyYTFjNzU5NGQzN2IyODA4NWEyODRkNzRjZWQiLCJwIjoiaiJ9&linkSource=email) ✅ 2025-06-17
    - Copied change, tried to push. I'm not a member of the repo.
    - Asked Addisu to add me to repo.
- Asked Addisu about getting TR running locally [slack](https://qualtrics.slack.com/archives/C07MKL3FB61/p1750085844815389)
- [x] Deploy CS at 96b2ec6f1fcdb325cd5427d3ff43844959c0ad7e #enrich-1158 ⏫ ✅ 2025-06-18
    - Updated config in deployments repo, sent [MR](https://gitlab-app.eng.qops.net/xm-discover/automation/deployments/-/merge_requests/7435)
    - Notified UO that I'm deploying soon [slack](https://qualtrics.slack.com/archives/C08T9CFN0SC/p1750113059577839?thread_ts=1747771591.222629&cid=C08T9CFN0SC)

# Log
- Try collating all my non-urgent vamsi things into one message a day? Maybe that would work better for him.
    - [x] **EOD:** Send Vamsi a note with: ✅ 2025-06-17
        - Check on my 1-1 meeting proposed time change this week
        - What jira items go in the sprint? Not pulses? What about ktlo?
- Copy [this](https://qualtrics.slack.com/archives/C090681KRN2/p1750110585399599?thread_ts=1749835526.744149&cid=C090681KRN2) slack thread for context, see if I can distill a "Common Pulses" section from it.
- uh oh i'm out thu-fri and 1158 is due friday. and i've got a half day wednesday. i need to get 1158 deployed now. notified in [slack](https://qualtrics.slack.com/archives/C087VC5JPD3/p1750112054130629?thread_ts=1750086646.985619&cid=C087VC5JPD3).

# Start of day tasks
- Regular
    - [x] Copy meeting/other schedule to notes ✅ 2025-06-16
    - [x] Check email ✅ 2025-06-16
    - [x] Check slack ✅ 2025-06-16
    - [x] **Monday**: Post async scrum update in team channel ✅ 2025-06-16
    - [x] Review blocked tasks ✅ 2025-06-16
    - [x] Review TAPP and ENRICH near due jira items (the sauron slack messages) ✅ 2025-06-16
- On-call
    - [x] Copy yesterday's final notes into coda #on-call ✅ 2025-06-16
    - [x] Create tasks for overnight alerts needing investigation #on-call ✅ 2025-06-16
    - [x] Triage pulses #on-call ✅ 2025-06-16


---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/Status/2025-06-17 Tue.md
---

---
created: 2025-06-17T07:43:29
---
# Status
- Meetings/other
    - 2x Dr appt
    - AI tools for Query-Based Topics
- On-call
    - [-] Research: Would the "Classification service errors (>100 in an hour)" alert gain specificity if we looked at proportion of errors to total message volume instead of absolute count of errors? #on-call 🔽 ❌ 2025-06-23
    - [-] Configure `check_docker_app_texteng-dbms` alert to not page for b1, at least not off hours #on-call 🔽 ❌ 2025-06-23
    - [[DGD-1631- Fannie Mae - Organization Hierarchy Volume Discrepancy - Categorized Records are not displaying in Studio]] was moved to our project, is now [[ENRICH-1597- Fannie Mae - Organization Hierarchy Volume Discrepancy - Categorized Records are not displaying in Studio]], due date extended to Jun 26
        - [b] Investigate [[ENRICH-1597- Fannie Mae - Organization Hierarchy Volume Discrepancy - Categorized Records are not displaying in Studio]] #on-call (DON'T PASS THIS OFF, YOU TOLD MAVIN YOU'D LOOK AT IT) 🔼 📅 2025-07-03
    - Added info on `::` and indexed fields to splunk team hub page
    - Ran brand deletion job
    - Pinged text-deployments about TAM release needing to start today if it's happening this week. Dongpeng's MR isn't merged yet.
    - b1 and gov1 crate host patch jobs appear stuck
        - Investigating
        - b1 is stuck on node 7. Killed b1 job.
        - Blehhhh it's the image path mixup again. It's looking for text-application path.
        - Updating crate publish script to publish to both paths, add required docker tags, improve variables
            - Sent [MR](https://gitlab-app.eng.qops.net/text-engine/crate-dbms/-/merge_requests/117)
            - [x] Follow up on crate [MR](https://gitlab-app.eng.qops.net/text-engine/crate-dbms/-/merge_requests/117) #on-call ✅ 2025-06-18
        - Published .37 to both billing team paths
        - Started a new b1 host patch job
            - completed successfully. b1 is on .37 now.
    - [x] gov1 crate job investigation #on-call ✅ 2025-06-23
        - crate container is running on the stuck node, but it's been trying to find a leader:
        - `[2025-06-17T20:30:24,309][WARN ][o.e.c.c.ClusterFormationFailureHelper] [texteng-dbms10-app.gov1.qprod.net] master not discovered or elected yet, an election requires at least 2 nodes with ids from [fgmGe5nqS36EoFtrcxA_ig, 4AtAXzIFTWahmjvqYz76Kw, nr10-PcLSaSE7mvOxd1nBQ], have discovered [] which is not a quorum; discovery will continue using [10.99.67.127:4500, 10.99.68.152:4500, 10.99.73.158:4500] from hosts providers and [{texteng-dbms10-app.gov1.qprod.net}{4AtAXzIFTWahmjvqYz76Kw}{AYBEY65XQ9O1cR5xDd1Vuw}{100.74.67.127}{100.74.67.127:4500}{http_address=100.74.67.127:4200}] from last-known cluster state; node term 100, last-accepted version 4056457 in term 100`
    - Flatline alert for... US? ANZ? [slack](https://qualtrics.slack.com/archives/C053MK619T9/p1750190437144809)
        - Looking at the splunk alert definition and I have questions...
        - It doesn't tell us which service flatlined. Wrote a new query to get more info.
        - [-] Break core enrichments services out of this alert, fix DC name, try to get it to say which service *is* flatlining #on-call 🔽 ❌ 2025-06-23
    - [-] Write doc on CS deployment for engage and discover [slack](https://qualtrics.slack.com/archives/C087VC5JPD3/p1750199687019449?thread_ts=1750197985.397909&cid=C087VC5JPD3) #on-call 🔽 ❌ 2025-06-23
    - Added more info on splunk filters to team hub
    - [-] Probably delete engage DC config for CS from discover deployment configuration repo? Verify with Mavin, he did it for model service #on-call 🔽 ❌ 2025-06-23
- [[ENRICH-1547- Update MIG Issuer for Topic Recommender]]
    - Sent [MR](https://gitlab-app.eng.qops.net/iq-research-sea/iq-aws-resources-manager/-/merge_requests/116) for deployment repo
    - Addisu merged it
    - [ ] Addisu wants me to trigger deployments [slack](https://qualtrics.slack.com/archives/D08JKAKG675/p1750175691824559) #enrich-1547-update-mig-issuer 
- [[ENRICH-1158- Enhance the Engine enrichment service to support skipping messages by configuration.]]
    - Addressed deployment MR feedback
    - Asked for expedited review
    - Merged deployment MR
    - Deployed to dev, staging
    - Seeing "Failed to extract project info from headers, skipping blocklist check" in staging over 3 minutes. Might be from a test. x-account-id is an int64.
    - Weird, seeing barely any logs for splunk search `index::qualtrics_discover_dev sourcetype::kube:container:classification` over the past 30 days. None in the past 7.
        - [-] Investigate no CS logs in dev index #on-call 🔽 ❌ 2025-06-23
    - Also seeing a plateau of "No models to run classification" in staging. Message processing also plateaus, so I don't think it's concerning. [slack](https://qualtrics.slack.com/archives/C035CQ3CHUK/p1750197821743619?thread_ts=1750195488.100509&cid=C035CQ3CHUK)
    - Deployed to EUW2
        - Looking good
    - Huddled with Mavin to learn about spinnaker deployment of discover services in engage DCs
        - [-] Delete engage DC config files from discover deployments config repo #on-call 🔽 ❌ 2025-06-23
    - Sent CS MR to use 96b2 in b1, g1. Merged.
    - Triggered CS spinnaker deployment to b1 (DEV-20250618T030735Z.5f0a680d9da973c1ecaac114d3af640bcad9b5ac)
        - b1 deployment looks good

# Log
- Sent Vamsi EOD note from yesterday
- Huddle with mavin on engage spinnaker deployment notes
    - DEV vs CD in spinnaker. CD only goes to b1
    - spinnaker doesn't know what to pull from docker
    - the git sha given in the deploy.sh command is what version spinnaker reads the config at
    - values.yaml is the default, each DC yaml is an override. this is a spinnaker thing.

# Start of day tasks
- Regular
    - [x] Copy meeting/other schedule to notes ✅ 2025-06-17
    - [x] Check email ✅ 2025-06-17
    - [x] Check slack ✅ 2025-06-17
    - [x] Review blocked tasks ✅ 2025-06-17
    - [x] Review TAPP and ENRICH near due jira items (the sauron slack messages) ✅ 2025-06-17
- On-call
    - [x] Copy yesterday's final notes into coda #on-call ✅ 2025-06-17
    - [x] Create tasks for overnight alerts needing investigation #on-call ✅ 2025-06-17
    - [x] Triage pulses #on-call ✅ 2025-06-17


---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/Status/2025-06-18 Wed.md
---

---
created: 2025-06-18T08:24:01
---
# Status
- Meetings/other
    - Scrum
    - 2x Dr appt
- Secondary on-call
    - Told Mavin I'd look at [[ENRICH-1597- Fannie Mae - Organization Hierarchy Volume Discrepancy - Categorized Records are not displaying in Studio]] with a discover partner [slack](https://qualtrics.slack.com/archives/C087VC5JPD3/p1750262659407039?thread_ts=1750256624.169359&cid=C087VC5JPD3)
    - [x] Move remaining #on-call tasks to jira items 🔼 ✅ 2025-06-23
- Merged crate publish script update MR
- Polished team hub page on splunk filters
- [[ENRICH-1158- Enhance the Engine enrichment service to support skipping messages by configuration.]]
    - Checked logs from deployments yesterday
        - Told Sandesh I found a few messages with x-account-id as int64, not string [slack](https://qualtrics.slack.com/archives/C035CQ3CHUK/p1750262480450689?thread_ts=1749045353.686389&cid=C035CQ3CHUK)
    - b1, staging, and prod uk all look good
    - Deploying to all prod DCs
    - Sent MR to update CS version in fra1, gov1
        - Pipeline being slow, then someone else merged so I had to merge that into my branch, re-push, wait for pipeline again
    - Updated UO on deployment status in slack
    - Deployed to prods CA, ANZ
    - Sent MR to update CS version in prod US, merged
    - Deployed to Prod US, all engage DCs
        - Prod US looks good
        - All DCs look good!
    - [x] Close/review tasks/tickets as appropriate #enrich-1158 ✅ 2025-06-18
    - [x] Make task to enable CS sharding in prod US ✅ 2025-06-18
        - Filed [[ENRICH-1601: Enable classification service sharding in Prod US]]

# Log
- Scrum
    - I noted that I doubt I'll get to [[ENRICH-1463- Iterate on prompt to reduce wildcard use]] this sprint
        - [x] Vamsi wants me to make a spike to better estimate [[ENRICH-1463- Iterate on prompt to reduce wildcard use]]. He doesn't like that I gave it 8 story points, too big. 📅 2025-06-25 ✅ 2025-06-25

# Start of day tasks
- [x] Copy meeting/other schedule to notes ✅ 2025-06-18
- [x] Check email ✅ 2025-06-18
- [x] Check slack ✅ 2025-06-18
- [x] Review blocked tasks ✅ 2025-06-18
- [x] Review TAPP and ENRICH near due jira items (the sauron slack messages) ✅ 2025-06-18
- [-] Spend >5 minutes on a starved task 📅 2025-06-18 ❌ 2025-06-18

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/Status/2025-06-23 Mon.md
---

---
created: 2025-06-23T11:12:24
---
# Status
- Meetings/other
    - Half-day vacation
    - Dr appt
- Secondary on-call
    - [[ENRICH-1597- Fannie Mae - Organization Hierarchy Volume Discrepancy - Categorized Records are not displaying in Studio]]
        - Responded to Vamsi's and Stan's status queries: I'm planning to work on it, but need discover help.
    - Moving my unhandled on-call spillover tasks to jira items
        - [[ENRICH-1611- Investigate model service EOF errors]]
        - [[ENRICH-1612- Investigate image consistency alerts]]
        - Daniel's handling [[TCP-476- [CY] Centene Corporation - Topic Hierarchy Generator Errors]] currently
        - [[ENRICH-1613- Update/create deployment docs for text-engine]]
        - [[ENRICH-1614- Expand doc on crate auto-patching troubleshooting]]
        - [[ENRICH-1615- Spike- Tune classification service error alert]]
        - [[ENRICH-1616- Tune crate alert to not page for b1]]
        - [[ENRICH-1617- Improve actionability of flatline alert]]
        - [[ENRICH-1618- Write docs on classification service deployment]]
            - [x] Load this for me? 🛫 2025-06-25 ✅ 2025-06-25
        - [[ENRICH-1619- Delete unneeded engage deployment config files for classification service]]
        - [[ENRICH-1620- Spike- Investigate few classification service logs in discover dev splunk index]]
    - Trying to get gov1 crate cluster back to 3 nodes (one's been stuck unable to find the cluster leader)
        - Rebooted node 10
        - Crate container isn't started upon reboot
        - [x] Update crate release version to .37 #on-call ✅ 2025-06-24
        - Mavin fixed this, pointed me to [docs](https://coda.io/d/_dt-izYOVJmq/Operations-Crate-from-CLI_suR2DTv3#_luofkI3X) on this issue
- Catching up on email and slack

# Log
- Took a half-day vacation for today just now

# Start of day tasks
- [x] Copy meeting/other schedule to notes ✅ 2025-06-23
- [x] Check email ✅ 2025-06-23
- [x] Check slack ✅ 2025-06-23
- [x] **Monday**: Post async scrum update in team channel ✅ 2025-06-24
- [x] Review blocked tasks ✅ 2025-06-23
- [x] Review TAPP and ENRICH near due jira items (the sauron slack messages) ✅ 2025-06-23
- [-] Spend >5 minutes on a starved task 📅 2025-06-23 ❌ 2025-06-24

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/Status/2025-06-24 Tue.md
---

---
created: 2025-06-24T08:05:58
---
# Status
- Meetings/other
    - Sprint review
    - [[2025-06-24 Tue - Retro]]
- Secondary on-call
    - Deploy crate .37 to g1
        - Disabled some near-due host patching jobs
            - [x] Re-enable host patching jobs ⛔ a16icc ✅ 2025-06-26
        - Added info to team hub page on this process
        - Started g1 host patching job
        - g1 deployment completed
    - [[ENRICH-1597- Fannie Mae - Organization Hierarchy Volume Discrepancy - Categorized Records are not displaying in Studio]]
        - Asked for help again [slack](https://qualtrics.slack.com/archives/C087VC5JPD3/p1750781281211649)
            - Addisu volunteered, I set a meeting
        - Asked for due date extension [jira](https://qualtrics.atlassian.net/browse/ENRICH-1597?focusedCommentId=11637563)
        - Met with Addisu, recorded meeting of some good KT on this issue, orgh
        - Asked res to force publish [slack](https://qualtrics.slack.com/archives/C090681KRN2/p1750811443191819)
    - Pinged data pipelines on-call about x-account-id type [slack](https://qualtrics.slack.com/archives/C035CQ3CHUK/p1750785056809419?thread_ts=1749045353.686389&cid=C035CQ3CHUK)
    - Sent [MR](https://gitlab-app.eng.qops.net/engvis/prometheus/rules/-/merge_requests/788) to disable k8s autoscaling alerts for fssync [slack](https://qualtrics.slack.com/archives/C053MK619T9/p1750799719304879?thread_ts=1750785893.338619&cid=C053MK619T9)
- [x] prepare for sprint review 🔼 ✅ 2025-06-24
    - Created slides for everyone
- [x] prepare for sprint retro 🔼 ✅ 2025-06-24
- Asked Vamsi if we're having sprint planning tomorrow. It's not on the calendar.
    - yes

# Log
- Another instance of meetings not being up-to-date #vamsi/feedback 

# Start of day tasks
- [x] Copy meeting/other schedule to notes ✅ 2025-06-24
- [x] Check email ✅ 2025-06-24
- [x] Check slack ✅ 2025-06-24
- [x] Review blocked tasks ✅ 2025-06-24
- [x] Review TAPP and ENRICH near due jira items (the sauron slack messages) ✅ 2025-06-24
- [-] Spend >5 minutes on a starved task 📅 2025-06-24 ❌ 2025-06-24

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/Status/2025-06-25 Wed.md
---

---
created: 2025-06-25T07:39:01
---
# Status
- Meetings/other
    - ~~Scrum~~
    - Dr appt
- [x] Load tickets for sprint ⏫ ✅ 2025-06-25
    - Filed [[ENRICH-1630- Spike- Investigate and Plan Smart Query Prompt Updates to Reduce Wildcards]]
- Notified team of name change to Kai
- Updated QEH with preferred name Kai, self-identification updates
- Updated calendar and time off for July
- Left abruptly to handle family emergency
    - Logged half-day vacation

# Log
- Vamsi updated the scrum calendar item to skip Thursdays! #vamsi/feedback

# Start of day tasks
- [x] Copy meeting/other schedule to notes ✅ 2025-06-25
- [x] Check email ✅ 2025-06-25
- [x] Check slack ✅ 2025-06-25
- [x] Review blocked tasks ✅ 2025-06-25
- [x] Review TAPP and ENRICH near due jira items (the sauron slack messages) ✅ 2025-06-25
- [-] Spend >5 minutes on a starved task 📅 2025-06-25 ❌ 2025-06-26

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/Status/2025-06-26 Thu.md
---

---
created: 2025-06-26T07:35:43
---
# Status
- Meetings/other
    - Dr appt
    - 1:1
- Secondary on-call
    - Re-enabled crate host patching jobs. Prod DCs will roll forward onto crate .37 now.
    - Merged rules MR 'Exclude fssync from autoscaling alerts'
- [x] Prepare for 1:1 ⏫ ✅ 2025-06-26
- [[ENRICH-1597- Fannie Mae - Organization Hierarchy Volume Discrepancy - Categorized Records are not displaying in Studio]]
    - Trying to find kibana log with queries
        - Asked Addisu for help, my filters aren't working as I expected
    - Addisu thinks the dashboards weren't refreshed
    - Trying to find a res contact who can help us since Stan is out [slack](https://qualtrics.slack.com/archives/C090PSZFN3F/p1750965298076829)
    - Luis (vargas) changed this from P1 to P2, which gives us time
    - Luis also gave us access to the fannie mae account and debugging dashboards
    - Still can't find logs, asking for help
- [x] Update name in zoom ✅ 2025-06-26
- Pinged Vamsi about time tracking again [slack](https://qualtrics.slack.com/archives/C0938EHBQN6/p1750977832855599)
    - [ ] Follow up on time tracking 🛫 2025-06-27 
- Pinged Vamsi about secondary on-call role, if they can do sprint work or only ktlo [slack](https://qualtrics.slack.com/archives/C0938EHBQN6/p1750978039007079)
    - [ ] Follow up on secondary query 🛫 2025-06-27 
- Requested on-call overrides for next shift
    - [ ] Ensure all overrides handled 🛫 2025-06-30 
- [[ENRICH-1568- Classification Service- Add sharding migration disabler]]
    - Filed [[ENRICH-1646- Clean up rmq migration disabling code]]

# Log
- 

# Start of day tasks
- [x] Copy meeting/other schedule to notes ✅ 2025-06-26
- [x] Check email ✅ 2025-06-26
- [x] Check slack ✅ 2025-06-26
- [x] Review blocked tasks ✅ 2025-06-26
- [x] Review TAPP and ENRICH near due jira items (the sauron slack messages) ✅ 2025-06-26
- [-] Spend >5 minutes on a starved task 📅 2025-06-26 ❌ 2025-06-26

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/Tickets/ENRICH-1114- Smart Query Prompt Updates.md
---

# Measure of success
>80% reduction while maintaining relevance and coverage seems like the right standard to me

--Donovan [slack](https://qualtrics.slack.com/archives/C07MKL3FB61/p1747433393073789?thread_ts=1747432019.857009&cid=C07MKL3FB61)


2025-04-07
* vamsi says top pri for this work is ~~supporting~~ (later: reducing use of) asterisk wildcard. other features include guardrails, ci/cd (for lambda functions?)
	* he says i'm the primary contributor for this item, kuan is supporting me
	* he says the 9 week estimate is "generous" to account for me coming up to speed
* Hardik answered about scope
	* > So it should be the 1114 and guardrails.
	* [qualtrics.slack.com](https://qualtrics.slack.com/archives/C036JC1JKGE/p1744048974592049?thread_ts=1743786031.481419&cid=C036JC1JKGE)
	* also [[ENRICH-1295- Designer Error Handling for smart-query guardrails]]

2025-04-21
vamsi says donovan wants telemetry for stuff like how many times users are triggering them, how often we hit llm quotas. investigate more.

2025-04-29
During sprint review, I wrote and said that I'm concerned about Donovan's insistence that my deliverable be done this quarter. Vamsi said it had been moved to July release, so less time pressure. I'm relieved.

(a few days later)
I misunderstood Vamsi in sprint review. This is still due before Q2 end, with testing. That's when code complete is for the July designer release.

# ~~My plan~~
* Come up with a plan
* Review the plan with the team, iterate as needed
* Collect test prompts to determine a baseline of wildcard use
* Determine wildcard use baseline
* Iterate with prompt changes to lower use of wildcard
    * How am I evaluating the prompts? I can take out all the wildcards easily, but if the model doesn't match like the baseline, it's worse?
    * The ticket says to replace wildcard use with variations of the word. Could I use lucene features like stemming (mtoken?) instead?
* Test? How?
* Deploy

# My new shorter plan 2025-05-15
I don't think I need to do all this performance testing. It'll take way longer and I'm not sure of the impact.
- Iterate on prompt to reduce wildcard use
- Evaluate performance of reduced-wildcard smart query rules compared to current
    - Need to find test data
    - Review by hand

## How to evaluate perf/fidelity/accuracy
I don't want my prompt change to make the queries/models much worse. It would help if I could quantify that?
Or... if I can drop-in replace wildcard use with mtoken, that should be close enough?
Put both of these options in the plan as alternatives, let the team review it?

# Resources
- [slack thread with links to data science docs on performance](https://qualtrics.slack.com/archives/C07MKL3FB61/p1746724077203799?thread_ts=1746720801.707089&cid=C07MKL3FB61)
- [coda page with pm sync meeting notes](https://coda.io/d/Text-Analytics-Enrichments-Hub_dAG-EE8cgH8/05-15-2025_sugsRBdx)

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/Tickets/ENRICH-1158- Enhance the Engine enrichment service to support skipping messages by configuration..md
---

# Resources
- [search for x-account-id use](https://sourcegraph.eng.qops.net/search?q=context:global+x-account-id+repo:classification&patternType=literal&sm=0)

# Initial notes
kuan thinks we should implement the customer blocklist for the other IG services that moved to core enrichments, not just classification: so sentiment, DAS, IS

slip some code into the existing checks that returns an error if the customer(account?)/project/instance is blocked. he suggested storing the blocklist in an env var (but could do yaml, just more complicated)

that trio is defined in protobuf-definitions:
```
message Metadata {
    option (version) = 1;
    
    int32 version = 1;
    string instance = 2;
    int64 project_id = 3;
    int64 account_id = 4;
}
```

which is used in:
```
message IngestDocument {
    option (version) = 1;
    int32 version = 1;
    string uuid = 2;
    string routing_key = 3;
    Metadata metadata = 4;
```

## Kuan guiding on how to test
cb jenkins builds all branches that get pushed? or can scan gitlab to get it sooner.
jenkins creates an image and the tag is my branch name
go into the deployment for the service in lens and directly edit the deployment file (not sure what it's called)
can find path to rabbitmq stuff in the deployment's env/config
use protocat to construct binary ingestdoc
run ingestion-scheduler to enqueue the message
can purge dev dlq anytime

## Testing
2025-04-22
I'm working on integration tests. vamsi suggested "gatling tool" for test data. no idea what it is.

# Slack Thread Summary: AMQP Header Type Issue (May 19-20, 2025)
## Context
My change for this item got reverted [[2025-05-21 Wed#^cf9012]]
## AI summary
**Topic:** Failures in `classification-service` (and initially suspected in `sentiment-service`) due to incorrect AMQP header types for `engine-process-account-id` (and potentially `project-id`) after messages were processed through the Dead Letter Queue (DLQ).

**Key Events & Discoveries:**

1.  **Initial Warning (Sandesh Gade, May 19):** Highlighted that logic extracting numerical Account IDs from AMQP headers must accommodate various numerical types (e.g., `int64`, `float64`), not just `int64`.
2.  **Problem Manifestation (May 20):**
    * Mikołaj Skałka (working on `sentiment-service`) initially faced issues.
    * Sandesh Gade provided a Splunk log showing the error: `"x-exception-message": "calling handler: account id header has incorrect type: expected int64, got float64"`.
    * The issue was then primarily identified within `classification-service`.
3.  **Root Cause Analysis:**
    * **Immediate Trigger (Bartosz Piętka):** The bug was likely surfaced by JBK's recent "skip project level messages" feature (ENRICH-1158) implemented in `classification-service` (specifically in `handler.go#L290`). This change introduced new logic that parsed the `engine-process-account-id` header.
    * **Underlying Issue (Sandesh Gade):** When messages pass through the DLQ and are stored in Postgres (as `jsonb`), numerical AMQP headers (like `engine-process-account-id`) can be converted to `float64` upon reprocessing. Handler logic in services strictly expecting `int64` would then fail.
    * **Header Usage Clarification (Sandesh Gade):** The `engine-process-*` headers (e.g., `engine-process-account-id: 102`) were originally intended for RabbitMQ routing. For application-level logic, the string-based `x-*` headers (e.g., `x-account-id: "102"`) are more reliable as they are consistently strings and avoid such type mismatches.
4.  **Impact & Context:**
    * The failures were specific to documents that had been reprocessed via the DLQ.
    * The change from ENRICH-1158 had been in production for about a week; the issue only became apparent when retryable failures caused messages to go through the DLQ and then be reprocessed by the updated `classification-service` logic.
5.  **Resolution:**
    * **Short-term (Hotfix):** Mavin Martin, with agreement from Milo Cason-Snow and JBK, decided to revert the ENRICH-1158 changes in `classification-service` to immediately unblock the reprocessing of affected documents. Mavin created the revert MR: `classification-service/-/merge_requests/199`.
    * **Long-term (Implied Recommendation):** Services parsing numerical `engine-process-*` headers should implement robust type handling (e.g., using a type switch). Alternatively, and preferably for handler logic, services should rely on the consistently string-formatted `x-*` headers.
# How to not cause an OP next time
I'm working on fixing the header issue (slowly). I don't want to cause an OP again. The reason my last change did is because I assumed I knew what data to read and what it looked like, and I wrote the code to error out if the expectation was violated. I thought this was good.

I could make the code more permissive, just warning if data isn't as expected. But I'm afraid of that allowing wrong data/expectations to fester unseen. If I log/metric on this warning, I could create an on-call alert for this situation. I could do it through AoM but discover DCs don't have that. Splunk is probably the easiest, most available way, even though log searches are gross.

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/Tickets/ENRICH-1468- Smart query version 2 in designer (epic).md
---

This is a 25Q2 key commit deliverable assigned to me.

# Resources
- [line item in planning coda doc](https://coda.io/d/_dCG_1QsXKay/Team-Planning_sucrg821#Team-Planning-Deliverables_tuA2DLvI/r27748)
- [DICE guardrails doc](https://coda.io/d/DICE-Documentation_dB0-fQgE27-/external-Socrates-Security-Guardrail-API_suYWM_Nf)
- [slack thread on socrates provided guardrails in LLM API vs standalone calls to guardrails API](https://qualtrics.slack.com/archives/C0186S22PQS/p1742922021578449)
- 

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/Tickets/ENRICH-1487- Add prompt injection guardrail to Smart Query.md
---

# Resources
- [[Guardrail context]]

# Smart Query - MIG Guardrail API Approach Decision
this section written by AI https://gemini.google.com/u/1/app/90d0a630cdec8346

**Date:** 2025-05-19

## Problem Statement

The Smart Query feature consistently triggers MIG's (Model Inference Gateway / Socrates / CME) `prompt_attack` guardrail when using the integrated LLM API (e.g., `/llms/v1/{modelId}/chat`). This occurs because the entire Smart Query prompt, which includes extensive developer-defined instructions alongside dynamic user/model inputs, is sent as a single "Human" / "user" message. MIG's `prompt_attack` guardrail interprets these embedded instructions as a potential attempt to override system behavior, leading to the request being blocked.

## Decision: Adopt Standalone Guardrail API for Prompt Attack Checks

The chosen approach is to utilize MIG's standalone Guardrail API endpoint (`model-inference-gateway/guardrails/v1/apply`) to pre-validate the dynamic/user-provided portions of the Smart Query prompt *before* making the main LLM call.

## Rationale

This decision is based on:

1.  **Explicit Internal Documentation Recommendation:**
    * The document `external-Socrates-Security-Guardrail-API_suYWM_Nf` explicitly recommends using the standalone `/apply` endpoint for `Prompt Attack` checks.
    * It states: *"You should ONLY pass the user specific input to this guardrail API, not the instructions from system prompt or user prompt because feature instructions will trigger guardrail and normally we don’t want that."*
2.  **Alignment with MIG API Design:**
    * The integrated LLM API guardrails are documented to assess only the "most recent 'user prompt'," excluding system prompts. Sending developer instructions as part of the user prompt is a misuse pattern for `prompt_attack`.
    * The `Security-Guardrail_sucPNMkw` document's definition of "Prompt Injection - Direct" (overriding existing instructions) aligns with why the `prompt_attack` guardrail is triggered by the current Smart Query prompt structure.
3.  **Input from MIG Team:** Previous discussions with MIG team members (e.g., James Argyropoulos) also pointed towards the standalone API as the correct way to isolate specific parts of the input for guardrailing.
4.  **Amazon Bedrock Basis:** The `prompt_attack` guardrail utilizes Amazon Bedrock's model. The standalone approach allows more precise control over what is sent to this underlying model for prompt attack detection.

## Key Implementation Steps & Details

1.  **Isolate Inputs:**
    * Identify and separate the truly dynamic parts of the Smart Query prompt. These include:
        * User-provided fields (e.g., "useCase", "industry").
        * The `<inputs>` block containing `query_id`, `topic_path`, `query_terms`.
2.  **Pre-validate with Standalone `/apply` API:**
    * Make a call to `model-inference-gateway/guardrails/v1/apply`.
    * In the payload, include *only* the isolated dynamic/user inputs in the `"texts": [...]` array.
    * Enable the `"prompt_attack"` guardrail (e.g., `"strength": "MEDIUM"` or as appropriate).
    * **Important:** The `/apply` API returns a `200 OK` status *even if a guardrail is tripped*.
    * Check the response body for `guardrail_assessments[{"action": "GUARDRAIL_INTERVENED", ...}]` and specifically `guardrail_results.prompt_attack.blocked: true` to determine if the input was rejected.
3.  **Conditional LLM Call:**
    * If `prompt_attack.blocked` is `false` (input is safe):
        * Construct the full Smart Query prompt (developer instructions + validated dynamic inputs).
        * Call the main LLM endpoint (e.g., `/llms/v1/{modelId}/chat`).
        * In this LLM call payload, the `prompt_attack` guardrail can be set to `"strength": "NONE"` or omitted, as the relevant content has already been checked.
        * `content_filter` for input (if not already covered by the `/apply` call) and output can remain enabled as needed.
    * If `prompt_attack.blocked` is `true`:
        * Do not proceed with the LLM call.
        * Handle the blocked request appropriately in the Smart Query UI/logic (e.g., display an error to the user).

## Related Documents/Threads

* MIG Guardrail API Documentation: https://coda.io/d/DICE-Documentation_dB0-fQgE27-/external-Socrates-Security-Guardrail-API_suYWM_Nf
* MIG Security Vulnerabilities Overview: https://coda.io/d/DICE-Documentation_dB0-fQgE27-/Security-Guardrail_sucPNMkw
* Slack thread with Addisu Guddissa and James Argyropoulos (link to specific thread if available in your Obsidian, e.g., [[Query Topics AI Features (lambdas)#Thread on MIG (aka CME, socrates) guardrails]] if you linked it from my previous response).

---

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/Tickets/ENRICH-1488- Add max tokens guardrail to Smart Query.md
---

---
creation date: 2025-05-30T11:15:08
---
Epic: [[ENRICH-1468- Smart query version 2 in designer (epic)]]
# Decisions
## Summary of [Slack Thread](https://qualtrics.slack.com/archives/C08U6PVGW4F/p1748628542375639): Clarification on "Max Tokens Guardrail" for Smart Query (2025-05-30)

**Initial State & Confusion:**
* Daniel Hersey inquired about documentation for "Max Tokens guardrails."
* Raj Verma (MIG/DICE platform team) initially responded that the platform does **not** offer a specific "Max Tokens Guardrail."
* Daniel then asked @jbk (you) how your team sets such guardrails and if it's done in the call to MIG.

**Discussion on Implementation, Intent & Origin:**
* Raj Verma elaborated on two general ways "guardrails" (potentially referring to output checks like content filtering, not necessarily a max token specific one) could be implemented via MIG (parallel chunk processing or on the full response) and offered to share documentation (which he later did via a Coda link). He also asked if @jbk was working on an implementation.
* @jbk (you) emphasized the need to understand the **underlying intent** and origin of the "max tokens guardrail" requirement before diving into implementation specifics for your task ENRICH-1488.
* Raj Verma reiterated that MIG doesn't have a "max tokens guardrail," the requirement wasn't from the Responsible AI team, and such a guardrail isn't present on AWS Bedrock. He offered to consider it for his roadmap if a specific need was defined.
* Daniel Hersey then asked Donovan Rich (Product) about the origin of the requirement.

**Resolution & Path Forward:**
* Donovan Rich stated he was unaware of the origin and had assumed it was a standard DICE (MIG platform) offering.
* **Decision from Donovan Rich to @jbk:** For Smart Query, the team should **proceed by imposing a character limit**, and this will be considered sufficient to meet the "max tokens" requirement.

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/Tickets/ENRICH-1597- Fannie Mae - Organization Hierarchy Volume Discrepancy - Categorized Records are not displaying in Studio.md
---

---
created: 2025-06-26T10:40:40
---
Project ID from screenshot in jira is 52291, think that project name is `FNMA Production`

oh all the IDs are [here](https://qualtrics.atlassian.net/browse/ENRICH-1597?focusedCommentId=11551209)

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/Tickets/TAPP-7074- TextIQ Bubble charts content is not perceptible and Bubble charts cannot be operated via keyboard.md
---

---
created: 2025-06-10T13:27:13
---
# Resources
- [public docs on dashboard bubble widget](https://www.qualtrics.com/support/vocalize/widgets/chart-widgets/text-iq-bubble-chart-widget-cx/)
- [a11y figma spec](https://www.figma.com/design/8mEekREbZBnMzclz7Px7dr/Dashboard-Widget-Accessibility--Dev-Specs?node-id=2-2&t=BZSsb2TOA7NVgQEj-1)
- [Chris's widget KT from May 2025](https://docs.google.com/presentation/u/1/d/1RJn1Y3JWhSSmIlooN31nWk2atFizp9_n6DoUV3m7c4U/edit?usp=sharing)
    - has details on how the various widget repos are linked, how to set up for development (he says you *can* get widgets running locally)
## Jira
### Description
Username: tam+telekomex@qualtrics.com#telekomex  
UserID: UR_5mUz8qZeUuTXcdE  
BrandID: telekomex  
DataCenterID: FRA1  
ProjectID: TS_9uskgibe2pvX1b3  
PageID: Page_88ff7fb7-6f68-49eb-834f-0fb512467264  
UserPermissionToLogin: Yes  
Reproducible: Yes

**Ticket Description:**  
(Our client Deutsche Telekom shared with us their Dashboard Accessibility Report that has been reviewed by the Accessibility SME Alexandria Reynolds and 29 issues were identified and approved as valid accessibility bugs by Alexandria. This issue is part of the report.)  
\[1] TextIQ Bubble charts content is not perceptible - Screen reader users can’t access the widget content because there is no accessible text alternative. Only mouse users can view the content in table format via an overlay.  
\[2] TextIQ Bubble charts cannot be operated via keyboard - Detailed information is only available to mouse users when hovering over the bubbles and there are no text alternatives for the bubble charts

Client's solution recommendation:

- An equivalent text alternative should be provided for all information presented in the widget
    
- For simple charts, the data can be presented in list form, for more complex charts in table form
    
- In addition, the charts require a text alternative, which is provided directly next to the chart. The chart's alternative text should contain the chart title and reference to the full text alternative
    
- Since the relevant content of the charts cannot usually be captured as efficiently in the text alternative as in the chart itself, the text alternative should summarize the most important information of the chart, e.g., ""The following table shows that \[X] have the highest score compared to \[Y] and are approximately \[X%] above the average.
    
- The content of bubble charts should be accessible to keyboard users. Implementation should take into account the recommendations of Problem \[1] Bubble chart content is not perceptible.
    

**Steps To Reproduce:**  
1. Go to the dashboard MAB Mai – Master-Dashboard (ES May 2024 – Master Dashboard)  
2. Open the page 2.4. Kommentare Karriereentwicklung

**Testing Performed:**  
As above

DashboardLink:
https://surveys.telekom.com/reporting-dashboard/web/65bd079c1ef9710008b2edc7/pages/Page_9d04a732-afed-4a62-aea6-ebc00343d6b8/view
Reported By: CLAIREC

# Slack dumps
## Telling Vamsi I'm stressed about this ticket and other things (2025-06-10)

Janzen Brewer-Krebs
  1:13 PM
Hey, I didn't realize I was starting on-call on Wed when we met yesterday. That's really going to cut into my bandwidth. Combined with that, I met with the new widgets team and Rohit about https://qualtrics.atlassian.net/browse/TAPP-7074 today and they're pushing hard for me to get on it. I told them I'd have scoping done this week before I remembered I'm going on-call.
Switching between all these different work streams is difficult and stressful for me, as is being on-call for this team.
So, I don't know that any action is needed, but I wanted to share all this.

TAPP-7074 TextIQ Bubble charts content is not perceptible and Bubble charts …
Status: Backlog
Type: Bug
Assignee: Janzen Brewer-Krebs
Priority: P3
More actions...
Added by Jira Cloud


Vamsi Vippadapu
  1:15 PM
the widgets load is supposed to be very limited and rotational basis
1:15
if you already contributed 1 or 2 points for widgets in last couple of weeks, that should be the extent of it


Janzen Brewer-Krebs
  1:18 PM
I haven't really done any widget work yet. I've had two pulses assigned to me. This one, 7074, and another a11y P4 that's due Aug 7. This is not going to be a 2 point task.


Vamsi Vippadapu
  1:39 PM
for the P4 at least , there is enough time for new reporting team to reverse shadow in Q3. What is your estimate on 7074, is this the first widget item that has been assigned to you so far in q2


Janzen Brewer-Krebs
  1:54 PM
I'm not sure if 7074 or 7072 came in first, but 7074 will be the first appreciable work I've done for widgets in Q2. My current, very uninformed estimate for 7074 is two weeks. I'm looking into it right now to try to refine that estimate.


Vamsi Vippadapu
  1:55 PM
yeah let me know what your estimate is, if it is 2 weeks, we can't accommodate and doesn't align with your priorities.


3 replies
Last reply 1 day agoView thread

    Also sent as direct message
    
    Janzen Brewer-Krebs
      Yesterday at 2:18 PM
    And just a heads up, I think Liam is expecting us to complete 7074 and may take issue with this stance. I'm not making any statement about whether this should or should not be our stance. Just letting you know based on my interactions with him.
    
    Vamsi Vippadapu
      Yesterday at 2:21 PM
    Is there a thread or channel with the expectation or tracking? I would also like to understand why it is such a huge effort and if there is a work around to close the pulse. Typically pulses should be quick fixes, not feature enhancements
    
    Janzen Brewer-Krebs
      Yesterday at 2:29 PM
    (moving conversation back to messages)


Janzen Brewer-Krebs
  2:16 PM
Okay, will do. Since Smart Query was deprioritized for the other AI features, should we update its status in the planning coda?


3 replies
Last reply today at 11:05 AMView thread


Janzen Brewer-Krebs
  2:18 PM
replied to a thread:
yeah let me know what your estimate is, if it is 2 weeks, we can't accommodate and doesn't align with your priorities.
And just a heads up, I think Liam is expecting us to complete 7074 and may take issue with this stance. I'm not making any statement about whether this should or should not be our stance. Just letting you know based on my interactions with him.
View newer replies

Janzen Brewer-Krebs
  2:30 PM
The conversation is (now) happening in \#tapp-7074-bubble-widget-a11y-telekom. The main thread on this when it came in and was assigned to me is https://qualtrics.slack.com/archives/C08RHU9K8NS/p1747171354456159.
It's unclear how much effort it'll be; there's lot of uncertainty. Widget code is old and weird, and this widget uses a graphics lib (d3) on top of that. They don't run locally and more stuff has probably broken since the last time anyone touched them. It's been a year+ since I touched widget code.
On top of that, it's not yet clear what we must provide a11y-wise and what is just the customer having ideas about how to fix the a11y issues. (edited) 

Janzen Brewer-Krebs
https://qualtrics.atlassian.net/browse/TAPP-7074 is more like a feature request. The client is asking for new stuff, and bubble widget is weird old code. What do you think 
@vamsi
?
Thread in textiq-widget-transition | May 13th | View message

Janzen Brewer-Krebs
  3:29 PM
I just saw this doc with details on what the widget team expects for handoff, for context: https://coda.io/d/_dqXDlMpiqcR/Handover-checklist_suHV_NH1

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/Tickets/Done/ENRICH-1148- Enable autoscaling for engines services in discover non-US clusters.md
---

============================
============================
DO NOT DEPLOY TO PROD DURING X4 FREEZE!
============================
============================



random info
	classification service uses queue based autoscaling. model service uses hpa
	use kubectl
	need to switch contexts between clusters/dcs
	prod-ca is prod3

	kubectl get scaledobject -n ingestion-pipeline-0
	kubectl get scaledobject scale-plat-classification-service -n ingestion-pipeline-0 -o yaml


need to create scaled object in other dcs
i see the scaled object def in helm-charts
	plat-classification-service/templates/keda-trigger.yaml has "kind" ScaledObject

why aren't there _any_ custom resources in prod-uk? how are they deployed/included in the cluster?
	seems it's an openlens issue?

services to check for autoscaling on:
	classification service
	model service (hpa, all others are keda)
	sentiment service
	intelligent scoring service
	derived attribute service

after talking to francisco, seems i just need to enable autoscaling, not investigate absence of scaledobjects

prod-uk/euw2
❯ kubectl get scaledobject -n ingestion-pipeline-0
NAME                                SCALETARGETKIND      SCALETARGETNAME               MIN   MAX   TRIGGERS   AUTHENTICATION                     READY   ACTIVE   FALLBACK   PAUSED    AGE
scale-fern                          apps/v1.Deployment   sentiment-server-fern         1     15    rabbitmq   keda-sentiment-trigger-auth        True    False    False      Unknown   123d
scale-plat-classification-service   apps/v1.Deployment   plat-classification-service   1     30    rabbitmq   keda-classification-trigger-auth   True    False    False      Unknown   35d

so need to add autoscaling for model, is, attr services. where do i do that? should be in deployments? yes.

asked about classification service already having a scaledobject in uk. milo added it when he filed this ticket, no mystery.

ONGOING/BLOCKED
working on derived attr service. anz doesn't have serviceAccount or autoscaling configs.
found vault. where would the anz secret be? where's another one? ca has `/secret-prod3/data/infra/rabbitmq/cb-ingestion-pipeline-prod3-ca1-rabbitmq/users/admin_keda`
	hm, found prod-us secret at https://vault-shared.clarabridge.io/ui/vault/secrets/secret-prod/kv/infra%2Frabbitmq%2Fcb-ingestion-pipeline-prod-us-rabbitmq%2Fusers%2Fadmin_keda/details?version=1
	asking about web UI for non-US vaults

intelligent scoring
	same missing stuff for anz as DAS

model service
	all prod clusters already have hpa enabled

# Thu Mar 20
Kuan verified that prod4/anz has the keda secret https://qualtrics.slack.com/archives/C087VC5JPD3/p1742483777114889?thread_ts=1742402382.615669&cid=C087VC5JPD3
	It's `/secret-prod4-apse2/data/infra/rabbitmq/cb-ingestion-pipeline-prod4-apse2-rabbitmq/users/admin_keda`. There's an extra `data` in there that isn't in the web ui, but I think that's an artifact of how it's structured when deployed or something. There's a `kv` in the url.

Asking about serviceAccount being missing again
	Kuan says it's needed 

sent mr, was approved

# Fri Mar 21
merged mr
how to deploy? jenkins?

- [x] DAS and IS are owned by core enrichments now. i should tell them to deploy the changes. vamsi says these services need autoscaling anyway. ✅ 2025-03-24

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/Tickets/Done/ENRICH-1266- 2024 Complete R&D Tax Credit Study (Janzen).md
---

https://berkshirewest.iad1.qualtrics.com/jfe/form/SV_3W5P1cOLruqJl1I?Q_DL=FQcoPh1L9lzRzIp_3W5P1cOLruqJl1I_CGC_NUoGwrn3BlScHfa&Q_CHL=email

# React Survey Flow Plugin
## Project 1 Experimentation
For each of the two uncertainties listed above, please describe in detail the process of experimentation performed to resolve these specific technical uncertainties. These could include:

- Research of potential technical solutions.
- The specific technical iterations in your Agile development process.
- Consideration of architectural/design alternatives.
- Development of test plans/environments.
- Unit testing, A/B testing, performance (latency) testing, functional testing, integration testing, acceptance testing, and other testing that is directed at resolving software development uncertainties prior to commercial release.

### Validation of externally sourced data
I researched different methods of doing this:
* only static typing
	* simple
	* insufficient--assumes ingested data is of the specific type. no runtime checks.
* static typing and ad hoc runtime validation at point of data use (typescript type guards)
	* simple and quick
	* relies on developer vigilance and adds to developer cognitive load
* static typing and comprehensive runtime validation at point of data ingestion
	* finds schema violations as soon as possible
	* assures the developer of external data being as expected/defined
	* more work and dependencies, and requires defining the schema twice (in zod and typescript types) to retain nice IDE integration

I settled on the last method. After that, I researched different approaches to implementation. There are many libraries for runtime schema validation. I considered at least: zod, yup, ajv, and io-ts. After researching capabilities, popularity, typescript support, and documentation quality, I chose zod.

During implementation I performed manual testing to verify that my validation code using zod rejected data that did not match my defined schemas.
### How to seamlessly migrate existing customer data
I researched different methods of doing this:
* migrate upon first access with new plugin
	* simple, contains the migration within this project
	* requires all customers to open and save their surveys before the migration would be complete
* re-route old sentiment API endpoint to newer version
	* doesn't require work or customer data changes on the survey flow plugin side
	* affects other, hidden old sentiment API users
	* makes me queasy, would be embarrassed to put my name to this
* crawl webservice definitions (i.e., customer data related to this plugin) to update sentiment API endpoint
	* comprehensive--moves all customers to new sentiment API in one stroke
	* requires learning, writing code in `db-crawler`
	* any crawling and editing of the database involves risk; higher likelihood of being cut by unknown unknowns in an unfamiliar codebase and language

I chose the final option--crawling the database. I wrote a prototype command in `db-crawler` to prove that the approach was feasible.

# AI Playground
## Goals
Improve user experience with nice-to-haves:
Add conversation persistence, switching, last model used memory, and associated UI
Re-architect the limited persistence that was already present to be structured, extensible, reliable, testable
## Experimentation
For each of the two uncertainties listed above, please describe in detail the process of experimentation performed to resolve these specific technical uncertainties. These could include:

- Research of potential technical solutions.
- The specific technical iterations in your Agile development process.
- Consideration of architectural/design alternatives.
- Development of test plans/environments.
- Unit testing, A/B testing, performance (latency) testing, functional testing, integration testing, acceptance testing, and other testing that is directed at resolving software development uncertainties prior to commercial release.
### Where to persist conversation data
I researched different methods of doing this:
* locally in browser
	* simple, requires no new infrastructure or communication
	* conversations would not persist across a user's different browser instances
	* security forbade us to store this data locally
* cloud relational database
	* provides rigid structure and durability for persisted data
	* more complicated to set up and administer
	* has ongoing cost
* cloud file storage (S3)
	* ongoing cost, but cheap
	* relatively simple to get and put json blobs
	* unstructured, for better or worse

I chose to use S3 for persistence. As this is an internal tool, there's less need for the higher degree of durability and structure provided by a relational database.
### Whether to invest in this homegrown codebase or adopt an existing open-source codebase
This project was developed rapidly earlier in the LLM boom. We wanted an internal UI/interface to securely use LLMs that we only had API access to. To save time, someone initially decided to reuse a somewhat related demo/prototype (haystack).

This demo code was updated to resemble commercially available LLM UIs and use our internal LLM API gateway. This work got it to the state of providing a single LLM conversation at a time that did not persist; if you reloaded the page or there was any error, you lost the conversation.

Given the gap between AI Playground at that point and commercial UIs, I researched whether we could leverage nascent open-source LLM chat UIs. I investigated the feasibility of using HuggingFace Chat UI with our internal LLM API gateway.

I found it feasible given the presence of API adapters. There would have needed to be additional investigation into the persistence layer and ensuring it meets our security standards, but a product decision was made to invest in our homegrown codebase.

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/Tickets/Done/ENRICH-1270- Text iQ keeps loading [503 status] for a user.md
---

# Mon Mar 24 2025
Permissions issue? I found these logs in TAM

```
{"message":"Failure when getting fields with data protection for fieldset 4813041e-d634-432e-882d-2487ee64fc01: 500 with body {\"transaction-id\":\"d2f1b318-b8a1-45bc-a82f-aa4cedd59ee3\",\"errors\":[{\"id\":\"DataFilteringServiceError::60030::d2f1b318-b8a1-45bc-a82f-aa4cedd59ee3\",\"status\":\"500\",\"code\":\"60030\",\"detail\":\"Dataset Catalog unavailable.\",\"meta\":{\"rootCause\":\"java.lang.IndexOutOfBoundsException: Index 0 out of bounds for length 0\"}}]}","level":"ERROR","timestamp":"2025-03-24T16:19:59.286Z","loggerName":"com.qualtrics.analytics.services.FieldSetCallerImpl","threadName":"scala-execution-context-global-552","type":"application"}
{"message":"An unexpected error occurred","level":"ERROR","stackTrace":"java.lang.Exception: Failed to get fields from DFS for fieldset '4813041e-d634-432e-882d-2487ee64fc01', responseCode=500\n\tat com.qualtrics.analytics.services.FieldSetCallerImpl.$anonfun$fieldsRaw$1(FieldSetCaller.scala:157)\n\tat scala.util.Success.$anonfun$map$1(Try.scala:255)\n\tat scala.util.Success.map(Try.scala:213)\n\tat scala.concurrent.Future.$anonfun$map$1(Future.scala:292)\n\tat scala.concurrent.impl.Promise.liftedTree1$1(Promise.scala:33)\n\tat scala.concurrent.impl.Promise.$anonfun$transform$1(Promise.scala:33)\n\tat scala.concurrent.impl.CallbackRunnable.run(Promise.scala:64)\n\tat java.base/java.util.concurrent.ForkJoinTask$RunnableExecuteAction.exec(ForkJoinTask.java:1426)\n\tat java.base/java.util.concurrent.ForkJoinTask.doExec(ForkJoinTask.java:290)\n\tat java.base/java.util.concurrent.ForkJoinPool$WorkQueue.topLevelExec(ForkJoinPool.java:1020)\n\tat java.base/java.util.concurrent.ForkJoinPool.scan(ForkJoinPool.java:1656)\n\tat java.base/java.util.concurrent.ForkJoinPool.runWorker(ForkJoinPool.java:1594)\n\tat java.base/java.util.concurrent.ForkJoinWorkerThread.run(ForkJoinWorkerThread.java:183)\n","timestamp":"2025-03-24T16:19:59.286Z","loggerName":"com.qualtrics.analytics.controllers.FieldValidationController","threadName":"application-akka.actor.default-dispatcher-531","type":"application"}
{"url":"/text-analytics/v1/fieldsets/9f0ec4a6-c9bd-7feb-d4bf-0e8bc8ce071a/source-fields/validationWithResponseCount","status":500,"time":3111,"clientIP":"100.77.11.93","method":"GET","userAgent":"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:136.0) Gecko/20100101 Firefox/136.0","datacenterFull":"syd1","type":"access","timestamp":"2025-03-24T16:19:59.286Z","issuer":"","authType":"UDS","id":"analytics-text-module","brandId":"msdnz","parentRequestId":"","sessionId":"","requestId":"8d76bd50-fe5d-44b8-a779-d7788d475bfe","userId":"UR_233Od4qn6eAQkJw","transactionId":"d2f1b318-b8a1-45bc-a82f-aa4cedd59ee3"}
```


Spunk logs suggest DFS gets a bad response or can't find dataset catalog. Asking for help from data platform first https://qualtrics.slack.com/archives/C03A90TCJ/p1742834842799039

Ahh this craig user doesn't have permission to view the source surveys. Mohsen's issue is different. Here's the log I pulled from TAM:

```
{"message":"An unexpected error occurred","level":"ERROR","stackTrace":"com.qualtrics.analytics.util.auth.UnexpectedHttpStatusException: GET http://textiq-gateway.service.consul/analytics-state-persistence/v1/fieldSets/b36df7ce-e3f4-475f-a756-8dc3c26b4d20/topicFields/derived-from expected status 200-399, got 403 with body \n\tat com.qualtrics.analytics.util.auth.AuthenticatingClient.$anonfun$execute$1(AuthenticatingClient.scala:113)\n\tat scala.util.Success.$anonfun$map$1(Try.scala:255)\n\tat scala.util.Success.map(Try.scala:213)\n\tat scala.concurrent.Future.$anonfun$map$1(Future.scala:292)\n\tat scala.concurrent.impl.Promise.liftedTree1$1(Promise.scala:33)\n\tat scala.concurrent.impl.Promise.$anonfun$transform$1(Promise.scala:33)\n\tat scala.concurrent.impl.CallbackRunnable.run(Promise.scala:64)\n\tat java.base/java.util.concurrent.ForkJoinTask$RunnableExecuteAction.exec(ForkJoinTask.java:1426)\n\tat java.base/java.util.concurrent.ForkJoinTask.doExec(ForkJoinTask.java:290)\n\tat java.base/java.util.concurrent.ForkJoinPool$WorkQueue.topLevelExec(ForkJoinPool.java:1020)\n\tat java.base/java.util.concurrent.ForkJoinPool.scan(ForkJoinPool.java:1656)\n\tat java.base/java.util.concurrent.ForkJoinPool.runWorker(ForkJoinPool.java:1594)\n\tat java.base/java.util.concurrent.ForkJoinWorkerThread.run(ForkJoinWorkerThread.java:183)\n","timestamp":"2025-03-24T22:46:55.999Z","loggerName":"com.qualtrics.analytics.controllers.FieldValidationController","threadName":"application-akka.actor.default-dispatcher-11","type":"application"}
{"url":"/text-analytics/v1/fieldsets/9f0ec4a6-c9bd-7feb-d4bf-0e8bc8ce071a/source-fields/validationWithResponseCount","status":500,"time":1472,"clientIP":"100.77.0.210","method":"GET","userAgent":"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:136.0) Gecko/20100101 Firefox/136.0","datacenterFull":"syd1","type":"access","timestamp":"2025-03-24T22:46:56.000Z","issuer":"","authType":"UDS","id":"analytics-text-module","brandId":"msdnz","parentRequestId":"","sessionId":"","requestId":"05b70cb7-1333-46db-ba4e-9715949c5a4b","userId":"UR_bkMopOKYb7KSOXA","transactionId":"c2e039a9-81ea-4245-b458-a940fe962585"}
```

Seems ASP gets a (responds with?) 403. Asking https://qualtrics.slack.com/archives/C8X9E12QY/p1742857689623559

# Tue Mar 25 2025

* Adding permission for Mohsen: Edit > Use Text Analysis on:
	* Client Heartbeat - Service Centre
	* Client Heartbeat - Outbound (Service Centre / Contact Centre)
	* Client Heartbeat - Multiple Interaction
	* Client Heartbeat - Inbound (Contact Centre)

Adding permission for Mohsen on Tickets:
* Client Heartbeat - Edit (this is a group?)
	* 
	* Did this on the CHICC ticket project? 
	* After this, I get a different error in TAM? Yeah, last time it was b36..., this time it's `4813041e-d634-432e-882d-2487ee64fc01`--Ticket Reporting: Client Heartbeat - Service Centre
* Student Experience - Edit
	* Hm, Mohsen can see the CHSC ticket project, but still getting authr error 

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/Tickets/Done/ENRICH-1285- Altice- Derived Attribute not working.md
---

https://coda.io/d/Query-Topics-Team-hub_dt-izYOVJmq/CDA_suPXHzqn?search=cda
CDA involves a parent child relationship
Category-Derived Attribute
It's for reporting? Just indicates that a document was matched with a topic? Seems weird that there's a whole new thing (CDA) for that.

Look for past issues that are similar, maybe enrich-1080, 904
	search in kibana like kuan did in 

Kuan investigated, figured it out https://qualtrics.slack.com/archives/C08KD3LEYBE/p1743083737582619:
For training purpose, my investigation process. 

1. Check from the UI
    1. Login to the Studio with “Studio Prod Integration User” which is available in 1Password. The ticket mentions this is for IBM deployment so the Studio is [https://cxstudio.clarabridge.net/dashboard/#/login](https://cxstudio.clarabridge.net/dashboard/#/login)
    2. Switch to the account “Altice” this is mentioned in the ticket. If not ask the Res team what the account it.
    3. Go to “Users & Groups” -> “Qualtrics Employess”, find the user for the Res team member (e.g. Mangesh), and choose to “View As”
    4. Then from the left side bar, click the link to go to “Designer”, then you can see the model in the “Categorize” tab
    5. The id of the model can be found from (edit:) inspecting the `<option>` elements in the model select dropdown
2. Check the log (for IBM, the log is available with Kibana [http://logs.clarabridge.net](http://logs.clarabridge.net/), OpenVPN is required)
    6. Search with `"Classification configuration" AND "model id: 63513306"` (this is the log from the ClassificationExecuter class)
    7. You will find the “requestId”
    8. The another search (edit: filter? see [here](https://qualtrics.slack.com/archives/C08KD3LEYBE/p1743112726491889?thread_ts=1743084878.993299&cid=C08KD3LEYBE)) for that request Id, and loggerName = CdaProcessor, you will see the CDA definition used by the code

Recording of the DB editing that fixed this issue https://zoom.us/rec/play/LmC1DTTzwSEBptiBjcHqCU4uK1xEAh-nWT5S6lzXP6pxG85HC3ROc_E6jjh-5aTYIGfODWA2OSmHybeq.wEl8G6PByGPsWM9E?accessLevel=meeting&canPlayFromShare=true&from=share_recording_detail&continueMode=true&componentName=rec-play&originRequestUrl=https%3A%2F%2Fzoom.us%2Frec%2Fshare%2F8BxuE_GO97S-cC8odQ4sVpGw-awZlNrpFOSr8g90skm-NpWvuuRMpe5fxUcZSKI.LtJJ-FZB4MdP4Njz%3FfromShareWithMe%3Dtrue[]()

---
source: /Users/jbrewer/Google Drive/My Drive/obsidian vault/My Vault/Tickets/Done/ENRICH-1305- Queue depth based alerts in GOV1 and all engage regions, review alert grouping.md
---

* Need to figure out how to get all engage regions publishing metrics? Are they already? Just the autoscaling ones?
	* No, just a few publish metrics. Or I can only find metrics for fra1, gov1, b1, g1.
	* Try finding the MR that added autoscaling and see if it did something with metrics too?
		* no
	* Looked at deployments repo configs and found rabbitmq-prometheus plugin but not the DC configuration pattern I expected based on the existing published metrics. Asked for help [qualtrics.slack.com](https://qualtrics.slack.com/archives/C087VC5JPD3/p1744043123303089?thread_ts=1743697014.538189&cid=C087VC5JPD3)
	* Addisu said ingestion-monitor pushes the metrics to prometheus for engage DCs
		* it runs as a cron job with a command from ingestion-monitor repo `cmd/rmq_queue_depth_monitor_cmd.go`
		* check logs of this cron job in a DC where metrics publish and where they don't to see if there are differences. still don't get why it's only these DCs
		* hol up. which engage clusters is classification service in? is it only those four?
			* yes
* do i just need to create the aom file?

