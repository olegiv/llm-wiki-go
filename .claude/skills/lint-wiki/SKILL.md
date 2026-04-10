---
name: lint-wiki
description: Run the wikilint Go CLI and fix every reported wiki issue in the wiki (not the linter).
---

# lint-wiki

Use this skill whenever the user asks you to lint the wiki or to fix
wiki structural issues.

## Steps

1. **Run the linter.**

   ```bash
   go run ./cmd/wikilint -wiki ./wiki
   ```

2. **Parse the output.** Each failing line is `<relative-path>: <message>`.
   The linter checks:
   - exactly one H1 title per page,
   - a non-empty `## Sources` section on every substantive page (all
     pages except `wiki/index.md` and `wiki/log.md`),
   - non-empty `## Contradictions` and `## Open questions` sections when
     present,
   - Markdown links resolve to existing `.md` files inside `wiki/`,
   - Obsidian-style `[[wikilinks]]` resolve (and are not ambiguous),
   - no duplicate entity pages by normalized title slug,
   - every page is reachable from `wiki/index.md`.
3. **Fix in the wiki, not in the linter.** For each issue, edit the
   relevant wiki page to satisfy the invariant. If the linter itself
   looks wrong, stop and flag it to the user — do not weaken checks to
   make output pass.
4. **Re-run** `go run ./cmd/wikilint -wiki ./wiki` until it prints
   exactly `wikilint: OK`.
5. **Append to `wiki/log.md`** if the fixes were substantive (not just a
   typo), noting what was fixed and why.

## Don'ts

- Do not delete pages to silence orphan warnings. Link them from
  `wiki/index.md` or a relevant topic page instead.
- Do not remove `## Sources` entries to silence broken-link warnings.
  Repair the link.
- Do not rename pages to dodge duplicate-slug warnings unless you first
  understand which of the duplicates should actually survive.
