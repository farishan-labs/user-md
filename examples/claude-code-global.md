# Example: Claude Code (global `~/.claude/CLAUDE.md`)

<!-- Note: Claude Code global CLAUDE.md uses plain bullets, not the OpenClaw metadata format.
     usermd-check will report "no directives found" for this file — that's expected. -->

- Be concise. No filler phrases. Show diffs, not full files.
- Default to TypeScript. Never suggest JavaScript.
- Prefer functional patterns over classes.
- Timezone: Asia/Jakarta (WIB). Working hours 08:30–17:30 Mon–Fri are work mode.
- Never run broad filesystem searches (grep/find/rg/fd/tree/ls -R) across ~, ~/dev, or /.
- Use `gio trash` instead of `rm -rf` for directories not safely on a remote.