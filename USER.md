# USER.md — Operator Profile Template

<!--
  USER.md — stable profile facts and preferences about the human operator.
  Loaded into the system prompt at session start.

  Format (OpenClaw canonical): one directive per entry.
    <!-- observed: YYYY-MM-DD | status: active -->
    - <Imperative directive: Always / Never / Prefer ...>

  Rules:
    1. Begin with an imperative (Always / Never / Prefer / Avoid).
    2. Record the date the preference was OBSERVED (not invented).
    3. Status is only "active" or "superseded".
    4. One behavioral instruction per directive.
    5. Store only details that improve assistance. Not a dossier.
    6. Supersede in place: keep the old (superseded) entry next to the new (active) one.
    7. Never store secrets here.
    8. Choose the right file:
         stable preference / comm style  -> USER.md (this file)
         durable non-profile fact/decision -> MEMORY.md (or OptMem)
         detailed observation / project   -> memory/YYYY-MM-DD.md
         project instructions/conventions  -> AGENTS.md
         agent persona/voice/values       -> SOUL.md

  Cap (recommended): ~1,375 chars (Hermes-style) to force discipline,
  since durable facts live in OptMem. Adjust to your harness budget.
-->

<!-- observed: 2026-11-11 | status: active -->

- Prefer concise, terse responses; no filler or restated questions.

<!-- observed: 2026-11-11 | status: active -->

- Timezone: Asia/Jakarta (WIB, UTC+7); working hours 08:30–17:30 Mon–Fri are work mode.

<!-- observed: 2026-11-11 | status: active -->

- Never run broad filesystem searches (grep/find/rg/fd/tree/ls -R) across ~, ~/dev, or /; ask to narrow scope instead.

<!-- observed: 2026-11-11 | status: active -->

- Use `gio trash` instead of `rm -rf` for directories not safely on a remote.

<!-- observed: 2026-11-11 | status: active -->

- Communicate in Indonesian or English; match the user's language.

<!-- Example of superseding in place:
<!-- observed: 2026-10-01 | status: superseded -->

- Prefer detailed explanations for every code change.

<!-- observed: 2026-11-11 | status: active -->

- Prefer concise implementation summaries unless more detail is requested.
-->