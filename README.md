# user-md — Operator Profile Template for AI Agents

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A portable, harness-agnostic **USER.md** template + guide for teaching your AI agent who you are — your role, preferences, communication style, and constraints — so every session starts with that context loaded.

> **File name stays `USER.md`** (that's what OpenClaw, Hermes, Claude Code, and OptMem look for). This repo brands the concept as an *operator profile* because "user" is ambiguous — but the file is `USER.md` for compatibility.

## Why

AI agents are stateless by default. Every session starts from zero unless you give them a stable profile. `USER.md` is the convention (originating in the [OpenClaw](https://docs.openclaw.ai) workspace) for a small, durable file that holds *who the human operator is* — separate from the agent's own persona (`SOUL.md`) and from project instructions (`AGENTS.md`).

This repo gives you:
- A clean, ready-to-fill `USER.md` template.
- A `GUIDE.md` with the format spec, best practices, and anti-patterns.
- A `WORKSHEET.md` to seed your profile in 5 minutes.
- Filled examples for OpenClaw, Hermes, Claude Code (global), and OptMem.

A standalone validator CLI (`usermd-check`) is included — see [`checker/`](checker/). Build with `cd checker && make build`.

## Quick start

Install the checker CLI:

```sh
go install github.com/farishan-labs/user-md/checker@latest
```

Or download a pre-built binary from the [releases page](https://github.com/farishan-labs/user-md/releases).

1. Copy `USER.md` to your agent's workspace root (or `~/.hermes/memories/USER.md`, or `~/.claude/CLAUDE.md` global).
2. Fill in `WORKSHEET.md` (5 questions) → paste the resulting directives into `USER.md`.
3. Keep it under your harness's cap (Hermes: ~1,375 chars; OpenClaw: 4,000 chars).
4. Audit monthly.

## The three-file split

| File | What it holds | Who writes it |
|---|---|---|
| `SOUL.md` | Agent's persona, voice, values | You (or the agent, agent-scoped) |
| `USER.md` | **Your** profile, preferences, constraints | You (seeded), agent (curated) |
| `AGENTS.md` | Project instructions, conventions | You / team |
| `MEMORY.md` | Durable non-profile facts and decisions | Agent |

Rule: `SOUL.md` and `USER.md` are separate systems that never feed each other.

## Format (OpenClaw canonical)

```markdown
<!-- observed: YYYY-MM-DD | status: active -->

- Prefer concise progress updates during implementation work.
```

- One directive per entry, imperative voice (Always / Never / Prefer).
- Record the date the preference was **observed**, not invented.
- Status is only `active` or `superseded`.
- **Supersede in place** — never append contradictions.
- Keep it compact; it's not a dossier.

See [`GUIDE.md`](GUIDE.md) for the full spec, best practices, anti-patterns, and migration paths.

## Compatibility

Works with any harness that loads a user-profile file or global instructions:

- **OpenClaw** — `USER.md` at workspace root (4,000-char budget)
- **Hermes Agent** — `~/.hermes/memories/USER.md` (~1,375 chars)
- **Claude Code** — global `~/.claude/CLAUDE.md` (layered)
- **Codex / Copilot / Aider / Cursor** — via `AGENTS.md` or global rules
- **OptMem** — as a portable operator-profile note

## License

MIT. See [LICENSE](LICENSE).

## Research backing

This template comes from a survey of how AI agents recognize their human operator across OpenClaw, Hermes, Claude Code, Cursor, Aider, Codex, Continue, Cline, Roo, Windsurf, Goose, and memory frameworks (Letta, Mem0, Zep, Honcho). Full research notes (evidence ledger, 18 sources, ecosystem survey) are maintained privately; the public sources that ground each claim are cited below.

Key findings:

- **Three separate problems:** identity (who is operating), personalization (preferences), authorization (what the agent may do). Most coding agents solve only ambient identity (OS user / git author) plus project instructions. *Sources: [CSA NIST AI agent standards note](https://labs.cloudsecurityalliance.org/research/csa-research-note-nist-ai-agent-standards-initiative-2026040/), [AgenTag attribution paper](https://arxiv.org/html/2608.00966v1).*
- **Dominant transferable pattern:** the **SOUL.md vs USER.md split** — the agent's own persona (`SOUL.md`) is kept strictly separate from the human's profile (`USER.md`) and from project instructions (`AGENTS.md`). OpenClaw docs state they "never feed each other." *Sources: [OpenClaw SOUL.md guide](https://clawdocs.org/guides/soul-md), [OpenClaw memory overview](https://docs.openclaw.ai/concepts/memory), [OpenClaw USER template](https://docs.openclaw.ai/reference/templates/USER), [Hermes migration from OpenClaw](https://hermes-agent.nousresearch.com/docs/guides/migrate-from-openclaw).*
- **Coding-agent convergence:** `AGENTS.md` is the cross-tool constitution, with tool-specific layers (`CLAUDE.md` global+project, `.cursor/rules/*.mdc`). Personal preference lives in global instructions, not in project files. *Sources: [AGENTS.md vs CLAUDE.md vs .cursorrules (automatelab)](https://automatelab.tech/blog/ai-coding/claude-md-agents-md-cursorrules), [AGENTS.md vs CLAUDE.md vs .cursorrules (stackitfast)](https://stackitfast.com/insights/agents-md-vs-claude-md-vs-cursorrules), [OpenClaw AGENTS.md template](https://docs.openclaw.ai/reference/templates/AGENTS).*
- **Memory ≠ identity:** memory systems (Letta, Mem0, Zep, Honcho) recognize users via a stable `user_id` (usually from an auth provider) and retrieve scoped memories; they don't replace an identity file. *Sources: [Letta memory docs](https://docs.letta.com/letta-agent/memory), [Mem0 memory layer guide](https://mem0.ai/guide/memory-layer-user-aware-ai-agents), [Hermes memory (remoteopenclaw)](https://www.remoteopenclaw.com/blog/hermes-agent-persistent-memory-methods), [Mem0 build an agent that remembers users](https://mem0.ai/blog/build-an-ai-agent-that-remembers-your-users).*
- **Recommendation:** keep the role/identity separation, and add one small, portable operator profile file (this `USER.md`) consolidating rules scattered across `AGENTS.md` files, with OptMem as cross-harness memory. *Sources: [Hermes tips & best practices](https://hermes-agent.nousresearch.com/docs/guides/tips), [Hermes user profiling (Stanza)](https://www.stanza.dev/courses/hermes-memory/user-profile/hermes-memory-profile-what).*
- **Security:** keep commit author (agent slot) distinct from push auth (human); never store secrets in identity files. *Sources: [AgenTag attribution paper](https://arxiv.org/html/2608.00966v1), [CSA NIST AI agent standards note](https://labs.cloudsecurityalliance.org/research/csa-research-note-nist-ai-agent-standards-initiative-2026040/).*

The core claim: agents are stateless by default, and a small, curated, *separate* profile file is more reliable than stuffing preferences into `AGENTS.md` or hoping the agent remembers.
