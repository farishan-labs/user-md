# GUIDE — USER.md format, best practices, and migration

Companion to `USER.md` and `WORKSHEET.md`. Full research backing: [`farishan-labs/research` → `topics/ai-agent-recognition-of-human-operator`](https://github.com/farishan-labs/research/tree/main/topics/ai-agent-recognition-of-human-operator) (see `user-md-convention.md`).

## What USER.md is

A small, durable file holding **stable profile facts and preferences about the human operator**, written as directives, loaded into the system prompt at session start. It answers "who is operating me and how do they like to work" — not "who am I" (`SOUL.md`) or "what is this project" (`AGENTS.md`).

## The three-file split

| File | What it holds | Who writes it |
|---|---|---|
| `SOUL.md` | Agent's persona, voice, values | You (agent-scoped) |
| `USER.md` | **Your** profile, preferences, constraints | You (seeded), agent (curated) |
| `AGENTS.md` | Project instructions, conventions | You / team |
| `MEMORY.md` | Durable non-profile facts and decisions | Agent |

Rule: `SOUL.md` and `USER.md` are separate systems that never feed each other.

## Format (OpenClaw canonical)

```markdown
<!-- observed: YYYY-MM-DD | status: active -->

- Prefer concise progress updates during implementation work.
```

Rules:
1. Begin with an imperative: **Always**, **Never**, **Prefer**, **Avoid**.
2. Record the date the preference was **observed** (not invented).
3. Status is only `active` or `superseded`.
4. One behavioral instruction per directive.
5. Store only details that improve assistance. Not a dossier.
6. **Supersede in place** — never append contradictions.
7. Never store secrets.

## Supersede in place

When a preference changes, update its existing section:

```markdown
<!-- observed: 2026-05-10 | status: superseded -->

- Prefer detailed explanations for every code change.

<!-- observed: 2026-07-27 | status: active -->

- Prefer concise implementation summaries unless more detail is requested.
```

Keep the superseded entry next to its replacement. Rationale: preference-following degrades in long conversations (PrefEval, arXiv:2502.09597); systems often re-select originally-stated preferences after a change (HorizonBench, arXiv:2604.17283). Append-only contradictory history recreates that failure.

## Cap management

When approaching the cap, in order:
1. Remove stale superseded entries (keep the replacement next to the old one).
2. Move durable facts/lessons that don't change how you should be assisted → `MEMORY.md`.
3. Move detailed observations/running project context → `memory/YYYY-MM-DD.md`.
4. Use standing intents for event-conditioned future actions.

## Choose the right file

| Information | Store in |
|---|---|
| Stable preference or communication style | `USER.md` |
| Durable non-profile fact / decision | `MEMORY.md` |
| Detailed observation / running project context | `memory/YYYY-MM-DD.md` |
| Project instructions / conventions | `AGENTS.md` |
| Agent persona / voice / values | `SOUL.md` |

## Best practices

1. **Seed manually on day one** — five lines you write yourself beat weeks of passive inference.
2. One directive per entry, imperative voice, dated, active/superseded.
3. Supersede in place — never append contradictions.
4. Keep it compact — directive-sized, not a dossier.
5. **Audit monthly** — profiling drifts.
6. Choose the right file — don't stuff project facts or durable decisions into `USER.md`.
7. Never store secrets.
8. Use separate profiles for work / personal / public bots.
9. Know your cap (Hermes: ~1,375 chars; OpenClaw: 4,000) and consolidate above 80%.
10. **Frozen-snapshot awareness** — writes don't affect the current session; restart to pick up.

## Anti-patterns

- **Dossier creep** — turning `USER.md` into a biography.
- **Append-only contradictions** — "Prefer X" then "Prefer not-X" both active.
- **Mixed scopes** — project facts in `USER.md`, personal prefs in `AGENTS.md`.
- **Secrets in identity files.**
- **Assuming mid-session writes take effect** (frozen snapshot).
- **Conflating agent persona with user profile** (`SOUL.md` vs `USER.md`).

## Migration

- **OpenClaw → Hermes:** `hermes claw migrate` parses/merges/dedupes `USER.md` into `~/.hermes/memories/USER.md` (§ delimiter). `IDENTITY.md` is archived, not applied (merge into `SOUL.md` manually).
- **OpenClaw → OpenClaw:** imported `USER.md` files stay separate under `memory/imports/...`, indexed for `memory_search`/`memory_get`, not merged into bootstrap `USER.md`. Source files left unchanged.
- **Cross-harness (OpenCode, pi, Claude Code, Codex, Cursor, OptMem):** no standard yet. Use this repo's `USER.md` format as the portable baseline; map to each harness's global instructions file.

## Harness caps

| Harness | File | Cap |
|---|---|---|
| OpenClaw | `USER.md` (workspace root) | 4,000 chars |
| Hermes Agent | `~/.hermes/memories/USER.md` | ~1,375 chars (~500 tokens) |
| Claude Code | `~/.claude/CLAUDE.md` (global) | (context budget) |
| OptMem | memory notes | (tree-based) |

## Principles: what belongs here (Faris's judgement)

A good `USER.md` is **minimal and taste-focused**. The maintainer's own profile follows these rules:

- **Keep:** stable preferences and tastes — response style, timezone/working hours, language, commit-attribution preference, how to handle ambiguity between personal and work context, decision-making style (trade-offs over single answers), risk tolerance (reversible first), and "explain the why".
- **Move to `AGENTS.md` (machine profile):** OS, shell, editor, machine-detection scripts, SSH/GitHub account mapping, "no `~/dev/personal` on this machine".
- **Move to `AGENTS.md` / `MEMORY.md`:** project facts, durable decisions, conventions, deploy/CI details.
- **Move to harness plugins or `AGENTS.md`:** infra guards — search-guard, no-heredocs, trash-vs-rm. These are enforcement concerns, not profile facts; a profile directive is a fallback only.
- **Move to harness built-in system prompt:** "read local `AGENTS.md` first", harness-specific rules (Agensi MCP usage).
- **Move to `MEMORY.md` / OptMem:** cross-session durable facts, OptMem usage discipline.
- **Never:** secrets, tokens, passwords, API keys.

Rationale: `USER.md` is loaded into the prompt every session with a tight budget. Filling it with infra rules crowds out the tastes that actually shape how the agent talks to you. Enforce rules in code (plugins) or project docs; reserve `USER.md` for who you are and how you like to be helped.

## Recommended cap

~1,000–1,375 chars (Hermes-style) to force discipline, since durable facts live in `MEMORY.md` / OptMem. Adjust to your harness.

## Validator (planned, v1.0)

`usermd-check` will validate: parse, cap, secrets scan, contradictions, scope mixing, harness compatibility. See [`checker/README.md`](checker/README.md). Contributions welcome.