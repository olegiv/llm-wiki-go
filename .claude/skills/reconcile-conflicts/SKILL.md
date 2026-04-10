---
name: reconcile-conflicts
description: When sources disagree, record the disagreement explicitly in a page's ## Contradictions section instead of silently merging.
---

# reconcile-conflicts

Use this skill whenever two or more sources disagree about something the
wiki already records, or whenever a new source contradicts an existing
wiki claim.

## Core rule

**Never silently merge conflicting sources.** The default policy of this
repo is explicit disagreement over implicit averaging. A future reader
should be able to see that sources disagreed and decide for themselves.

## Steps

1. **Identify the affected page.** This is usually an entity or topic
   page under `wiki/entities/` or `wiki/topics/`.
2. **Read every disputing source.** Read each source page referenced in
   the affected page's `## Sources` section, plus any new source you
   are ingesting. Understand what each one actually claims.
3. **Edit the page's `## Contradictions` section.** If it does not
   exist, add it after `## Sources`. Inside it, use an H3 describing
   the topic of the disagreement, then one bullet per source:

   ```markdown
   ## Contradictions

   ### Founding year
   - sources/alice.md: Founded in 1999
   - sources/bob.md: Founded in 2001
   ```

4. **Do not overwrite the page's existing summary.** If the page's
   `## Summary` previously picked a side, rewrite that paragraph to be
   neutral and refer the reader to `## Contradictions`.
5. **Open a question if the disagreement is unresolved.** Add a bullet
   under `## Open questions` on the same page naming what further
   evidence would resolve it.
6. **Append to `wiki/log.md`** explaining what was reconciled and which
   sources disagreed.
7. **Run `wikilint`.** `go run ./cmd/wikilint -wiki ./wiki` must print
   `wikilint: OK`.

## Don'ts

- Do not delete or rewrite `raw/` content to "align" the sources.
- Do not pick the more recent source as the winner by default — freshness
  is not the same as correctness.
- Do not omit the contradiction because it feels minor; a small
  disagreement today is the breadcrumb that explains tomorrow's bug.
