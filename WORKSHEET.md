# WORKSHEET — Seed your USER.md in 5 minutes

Answer the 5 questions below. Each answer becomes one or more directives in `USER.md`.

## 1. Who are you, in one line?

Role, domain, level. (e.g., "Solo dev building AI-assisted tools; Indonesian, WIB timezone.")

→ Directive(s):

## 2. How do you like to be communicated with?

Terse vs. detailed, language, tone, format preferences.

→ Directive(s):

## 3. What are your non-negotiables?

Rules the agent must always/never do. (e.g., "Never run broad filesystem searches", "Use gio trash not rm -rf".)

→ Directive(s):

## 4. What's your environment?

OS, shell, editor, timezone, working hours, work vs. personal mode rules.

→ Directive(s):

## 5. What should the agent never touch?

Secrets, off-limits paths, things requiring human approval.

→ Directive(s):

---

## Convert answers to directives

For each answer, write one directive per behavioral instruction:

```markdown
<!-- observed: YYYY-MM-DD | status: active -->

- Prefer concise, terse responses; no filler or restated questions.
```

Rules: imperative voice, dated, one instruction per directive, no secrets, no project facts (those go in `AGENTS.md` / `MEMORY.md`).

## Then

1. Paste directives into `USER.md`.
2. Check char count against your harness cap.
3. Remove the example directives in the template that don't apply.
4. Commit. Audit monthly.