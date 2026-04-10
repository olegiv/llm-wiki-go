---
name: answer-from-wiki
description: Answer user questions from wiki/ first, consult raw/ only when needed, then compile missing knowledge back into wiki/.
---

# answer-from-wiki

Use this skill whenever the user asks a knowledge question about anything
the wiki is responsible for.

## Preconditions

- You have read `CLAUDE.md` and `AGENTS.md` for this repo.
- The question is a knowledge question, not a coding or tooling request.

## Steps

1. **Search `wiki/` first.** Glob and grep under `wiki/` for the
   relevant terms. Read every page that could contain the answer.
2. **If the wiki answers the question**, answer the user, cite the
   specific pages you relied on (slash-separated relative paths like
   `wiki/entities/foo.md`), and stop.
3. **If the wiki only partially answers the question**, answer the part
   the wiki covers, then continue to step 4 for the rest.
4. **Consult `raw/` when necessary.** Search `raw/` for the missing
   piece. Read only the specific raw files you actually need. Never
   modify `raw/`.
5. **Answer the user** with the combined information, citing both wiki
   pages and the relevant `raw/` file(s).
6. **Compile the new knowledge back into `wiki/`.** Follow the
   `ingest-source` skill to turn what you just learned into durable
   wiki pages (including a `wiki/sources/` page citing the raw file),
   and update `wiki/index.md` and `wiki/log.md`.
7. **Run `wikilint`.** `go run ./cmd/wikilint -wiki ./wiki` must print
   `wikilint: OK` before you hand control back to the user.

## Don'ts

- Do not skip step 6. The point of this wiki is that the next question
  should be answerable from the wiki alone.
- Do not paraphrase from memory when you have the source file available
  — quote or cite precisely.
- Do not write to `raw/`, ever.
