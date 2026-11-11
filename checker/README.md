# checker — `usermd-check` (planned, v1.0)

A standalone, cross-platform Go binary that validates a `USER.md` file.

## Planned checks

- **Parse** — valid markdown, well-formed metadata lines, status ∈ {active, superseded}.
- **Cap** — char count vs budget (`--budget 1375|4000`, default 1375).
- **Secrets** — regex scan for tokens, API keys, emails, passwords.
- **Contradictions** — two `active` directives that conflict (heuristic).
- **Scope mixing** — directives that look like project facts / durable decisions → suggest `MEMORY.md`.
- **Harness compat** — report which harnesses would accept the file (OpenClaw 4k, Hermes 1.375k, Claude Code global, OptMem).

## Output

Human-readable by default; `--json` for CI.

## Status

Not yet implemented. Contributions welcome. Tracking: see the private ops repo (maintainer only).