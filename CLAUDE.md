# CLAUDE.md

@AGENTS.md

## Repo goal

`llm-wiki-go` is a Karpathy-style "LLM Wiki". Treat `raw/` as immutable
source material and `wiki/` as the canonical compiled knowledge layer in
Obsidian-friendly Markdown. You are the compiler/editor: read from
`raw/`, compile facts into `wiki/`, and keep the wiki internally
consistent.

## Canonical directories

- `raw/` — immutable source material. **Read-only.** Tracked in git.
- `wiki/` — canonical compiled knowledge in Markdown. Tracked in git.
- `wiki/entities/` — pages about discrete things.
- `wiki/topics/` — pages aggregating related entities and sources.
- `wiki/sources/` — one page per ingested item from `raw/`.
- `cmd/wikilint/` + `internal/wikilint/` — the linter.
- `internal/wiki/` — helpers for ingest / answer / reconcile workflows.
- `.claude/` — skills, agents, and the shared submodule.

`raw/` ships as a directory of symlinks into a sibling source checkout
(the default example is `../ocms-go.core`). `wiki/` ships as the
compiled Markdown. Both are tracked in git — no `make setup` needed for
the default `ocms-go.core` example, provided `ocms-go.core` is cloned as
a sibling of this repo. `make setup` remains available for bootstrapping
against a different source repo.

## The `raw/` rule

Never write to, rename, or delete anything under `raw/`. Not whitespace,
not encoding fixes, not reformatting. If a raw source looks wrong, flag
it to the user and wait for instructions.

## Answer-from-wiki-first rule

When the user asks a knowledge question, answer from `wiki/` first.
Only consult `raw/` when the wiki genuinely does not contain the answer.
After consulting `raw/`, compile what you learned back into `wiki/` so
the next question can be answered from the wiki alone.

## Required `## Sources` section

Every substantive wiki page has exactly one H1 title, a `## Summary`, and
a `## Sources` section. Only `wiki/index.md` and `wiki/log.md` are exempt
from the `## Sources` requirement. Every non-trivial claim on a page
should be traceable through `## Sources` back to something in `raw/` or
to another wiki page that cites `raw/`.

## Handling contradictions

When sources disagree, do not silently pick a winner and do not
average-out the facts. On the affected page, use a `## Contradictions`
section and record each position as a bullet with the source path and
the claim:

```markdown
## Contradictions

### Founding year
- sources/alice.md: Founded in 1999
- sources/bob.md: Founded in 2001
```

If the disagreement cannot be resolved yet, also add an entry to the
page's `## Open questions` section.

## Required updates on every wiki edit

Every wiki change updates two bookkeeping pages:

1. **`wiki/index.md`** — keep new and renamed pages reachable from the
   index. The `wikilint` orphan check fails if a page is not reachable.
2. **`wiki/log.md`** — append a dated entry describing what changed and
   why. Most recent entries at the top.

## Run `wikilint` before finishing a wiki edit

Before you consider a wiki change done, run:

```bash
make lint-wiki
```

(equivalent to `go run ./cmd/wikilint -wiki ./wiki`). It must print
exactly `wikilint: OK`. If it reports any issue, fix the wiki content —
not the linter — and re-run until it is clean.

## Before finishing any code change

Run the full verification chain:

```bash
make check
```

That runs `gofumpt -l .`, `go vet ./...`, `golangci-lint run ./...`,
`wikilint`, and `go test ./...`. All steps must succeed before you hand
work back to the user. Run `make help` to list every target, including
`make build-prod`, `make build-linux-amd64`, and `make build-darwin-arm64`.
