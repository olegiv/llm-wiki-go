# AGENTS.md

Guidance for any coding agent working in this repo. Keep it skim-able and
use it as a checklist before finishing a task.

## Project purpose

`llm-wiki-go` is a Karpathy-style "LLM Wiki". The repo has two data
layers:

- `raw/` — immutable source material (documents, transcripts, notes,
  captures). Nothing in `raw/` is ever rewritten, summarized in place, or
  deleted by an agent.
- `wiki/` — the canonical compiled knowledge layer, in Obsidian-friendly
  Markdown. Every substantive fact lives here, with sources traced back to
  `raw/`.

Claude Code and other agents act as the compiler/editor: they read from
`raw/`, compile knowledge into `wiki/`, and keep the wiki internally
consistent. The `wikilint` Go CLI enforces the structural invariants.

Both `raw/` and `wiki/` are gitignored and may be real directories or
symlinks to an external location. Run `make setup` after a fresh clone
to create the directory structure and seed wiki files.

## Canonical directories

| Path               | Purpose                                             |
|--------------------|-----------------------------------------------------|
| `raw/`             | Immutable source material. **Read-only for agents**. Gitignored. |
| `wiki/`            | Canonical compiled knowledge (Markdown). Gitignored. |
| `wiki/entities/`   | Pages about discrete things (people, products, …). |
| `wiki/topics/`     | Pages that aggregate related entities and sources. |
| `wiki/sources/`    | One page per ingested item from `raw/`.            |
| `cmd/wikilint/`    | Linter CLI entry point.                             |
| `internal/wikilint/` | Linter logic, tested in-package.                   |
| `internal/wiki/`   | Helpers for ingest / answer / reconcile workflows.  |
| `.claude/skills/`  | Claude Code skills keyed to wiki workflows.         |
| `.claude/agents/`  | Focused Claude Code sub-agent definitions.          |
| `.claude/shared/`  | Git submodule: shared Claude Code support tooling.  |

## Go coding standards

- Idiomatic Go. Standard library first. Only reach for a third-party
  dependency if there is a strong reason and the user signs off.
- Small, focused functions with clear names. No premature abstractions.
- No dead code, no commented-out code, no TODOs without an owner or
  linked issue.
- Keep comments sparse. Document _why_, not _what_. Exported identifiers
  get one-line godoc comments.
- Error messages are lowercase and do not end with punctuation.
- `gofmt` must be clean before you finish.

## Package layout expectations

- Binaries live under `cmd/<name>/` with a thin `main.go` that parses
  flags and delegates to a package under `internal/`.
- Reusable logic lives under `internal/`. Nothing outside this repo should
  import `internal/...`.
- Tests live alongside the code in the same package (not `_test` suffix)
  so unexported helpers can be tested directly. Prefer table-driven tests
  when you are covering more than two cases.

## Testing expectations

- `go test ./...` must pass before you finish.
- New behavior gets a new test. Bug fixes get a regression test.
- Don't mock the filesystem — use `t.TempDir()`.
- Keep tests hermetic: no network, no fixtures outside the repo, no time-
  of-day dependencies.

## Wiki-editing constraints

- **Never modify files under `raw/`.** Not even whitespace. If you notice
  `raw/` content that looks wrong, flag it to the user instead of fixing
  it.
- Every substantive wiki page has exactly one H1 title, a `## Summary`
  section, and a `## Sources` section. `wiki/index.md` and `wiki/log.md`
  are the only pages exempt from the `## Sources` requirement.
- When sources disagree, use a `## Contradictions` section on the
  affected page. Record each position with the source path and the
  claim. Do not silently pick a winner.
- Every wiki edit updates `wiki/index.md` so the new or changed page
  stays reachable, and appends a short entry to `wiki/log.md` describing
  what changed and why.

## Small, focused changes

- One logical change per commit. If you find yourself touching unrelated
  files, split them out.
- No drive-by refactors, no unrequested reformatting of existing files.
- No backwards-compatibility shims — this project has no external users
  to preserve API compatibility for.

## Before you finish

Run the full verification chain and fix anything that fails:

```bash
make check
```

`make check` runs, in order, `gofmt -l .`, `go vet ./...`,
`go test ./...`, and `go run ./cmd/wikilint -wiki ./wiki`. The linter
step must print exactly `wikilint: OK`. If any step fails, fix the root
cause — do not bypass or skip them.

Run `make help` to see every available target, including `make fmt`,
`make lint`, and `make build`.
