---
name: ingest-source
description: Read files in raw/, extract entities/topics/claims, and compile them into wiki/ pages with full provenance.
---

# ingest-source

Use this skill to compile new knowledge from `raw/` into `wiki/`.

## Preconditions

- The source material already exists under `raw/`. You do not create it.
- You have read `CLAUDE.md` and `AGENTS.md` for this repo.

## Steps

1. **List raw files.** Walk `raw/` and pick the file (or files) the user
   asked about. Never modify anything under `raw/`.
2. **Read.** Read each target file in full. Do not paraphrase from the
   filename; read the contents.
3. **Extract.** Identify the entities (discrete things), topics (aggregating
   subjects), and concrete claims present in the source.
4. **Create or update source page.** For each ingested raw file, ensure a
   page exists under `wiki/sources/` with:
   - an H1 title naming the source,
   - a `## Summary` section describing what the source contains,
   - a `## Sources` section whose first bullet is `Origin: \`raw/<path>\``.
5. **Create or update entity/topic pages.** For each entity or topic the
   source informs, create or update the matching page under
   `wiki/entities/` or `wiki/topics/`. Every substantive page needs an
   H1, a `## Summary`, and a `## Sources` section citing the source page
   you created in step 4.
6. **Update `wiki/index.md`.** Add links to any brand-new pages so the
   orphan check stays clean.
7. **Append to `wiki/log.md`.** Add a dated entry at the top describing
   what was ingested and which pages were touched.
8. **Reconcile contradictions.** If the new source contradicts something
   already in the wiki, invoke the `reconcile-conflicts` skill instead
   of silently rewriting existing claims.
9. **Run the linter.** `go run ./cmd/wikilint -wiki ./wiki`. It must
   print exactly `wikilint: OK`. Fix any reported issue in the wiki.

## Don'ts

- Do not write anything under `raw/`.
- Do not summarize a source page away — keep `wiki/sources/<name>.md` as
  the permanent anchor for citations.
- Do not skip the log entry, even for small ingests.
