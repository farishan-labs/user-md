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

## Related

- Research: [`farishan-labs/research` → `topics/ai-agent-recognition-of-human-operator`](https://github.com/farishan-labs/research/tree/main/topics/ai-agent-recognition-of-human-operator)
- Deep dive: [`user-md-convention.md`](https://github.com/farishan-labs/research/blob/main/topics/ai-agent-recognition-of-human-operator/user-md-convention.md)
- Companion template file: [`templates/USER.md`](https://github.com/farishan-labs/research/blob/main/topics/ai-agent-recognition-of-human-operator/templates/USER.md)