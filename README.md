# llm-wiki-go

[![Go](https://github.com/olegiv/llm-wiki-go/actions/workflows/go.yml/badge.svg)](https://github.com/olegiv/llm-wiki-go/actions/workflows/go.yml) [![CodeQL](https://github.com/olegiv/llm-wiki-go/actions/workflows/github-code-scanning/codeql/badge.svg)](https://github.com/olegiv/llm-wiki-go/actions/workflows/github-code-scanning/codeql)

A [Karpathy-style][karpathy-llm-wiki] "LLM Wiki" — a repo designed around
the idea that the interesting knowledge artifact is a living, compiled
wiki, not the raw source material it was distilled from.

[karpathy-llm-wiki]: https://gist.github.com/karpathy/442a6bf555914893e9891c11519de94f

The repo has two data layers and a small amount of Go tooling that keeps
them honest:

- **`raw/`** holds immutable source material: documents, transcripts,
  notes, captures. Nothing in `raw/` is ever rewritten, summarized in
  place, or deleted.
- **`wiki/`** is the canonical compiled knowledge layer, in
  Obsidian-friendly Markdown. Entities, topics, and sources each get
  their own page. Every substantive claim traces back to `raw/` through
  a `## Sources` section.

Claude Code (and any other coding agent operating in this repo) acts as
the compiler and editor: it reads from `raw/`, compiles knowledge into
`wiki/`, and keeps the wiki internally consistent. The rules the agent
follows are in [`CLAUDE.md`](CLAUDE.md) and [`AGENTS.md`](AGENTS.md).

## How the compiled knowledge is enforced

The Go CLI `wikilint` enforces the structural invariants so the wiki
stays well-formed as it grows. It checks that every page has a single
`# ` title, that substantive pages have a non-empty `## Sources`
section, that Markdown links and Obsidian-style `[[wikilinks]]` all
resolve, that no two pages collide on a normalized title slug, and that
no page is an orphan unreachable from `wiki/index.md`.

## File layout

```
llm-wiki-go/
├── .claude/              # Claude Code skills, agents, shared submodule
│   ├── agents/           #   focused sub-agent definitions
│   ├── shared/           #   git submodule: claude-code-support-tools
│   └── skills/           #   workflow skills (ingest, answer, lint, reconcile)
├── cmd/
│   └── wikilint/         # linter CLI entry point
├── internal/
│   ├── wiki/             # helpers for ingest / answer / reconcile workflows
│   └── wikilint/         # linter implementation, tested in-package
├── raw/                  # immutable source material (read-only for agents)
├── wiki/                 # canonical compiled knowledge layer (Markdown)
│   ├── entities/
│   ├── topics/
│   ├── sources/
│   ├── index.md          # entry point; every page must be reachable from here
│   └── log.md            # append-only record of substantive wiki changes
├── AGENTS.md             # agent guidelines
├── CLAUDE.md             # Claude Code guidelines (imports AGENTS.md)
├── Makefile              # convenience wrapper around go + wikilint
├── README.md
└── go.mod
```

No `go.sum` file ships with the scaffold: the project currently depends
only on the Go standard library, and Go does not emit a `go.sum` until
external modules are required.

## Getting started

`raw/` and `wiki/` are **tracked in git** and ship with the repo. The
default example uses `raw/ocms-go.core/`, a directory of symlinks into a
sibling `../ocms-go.core` checkout. `wiki/` ships as the compiled
Markdown.

For the default example, clone `ocms-go.core` as a sibling of this repo
and you're ready to go — no `make setup` needed:

```bash
cd ..
git clone https://github.com/olegiv/ocms-go.core.git
cd llm-wiki-go
```

`make setup` remains available for bootstrapping against a different
source repo. It creates the directory structure (`raw/`,
`wiki/entities/`, `wiki/topics/`, `wiki/sources/`) and seeds
`wiki/index.md` and `wiki/log.md` if they don't already exist.

Using `llm-wiki-go` for a project other than the bundled
`ocms-go.core` example — or running wikis for several projects in
parallel — is documented in
[`docs/MULTI_PROJECT.md`](docs/MULTI_PROJECT.md).

## Commands

All common tasks go through the Makefile. Run `make help` to list the
available targets:

```bash
make help          # list every target with a one-line description
make setup         # create raw/ and wiki/ directory structure
make build         # compile bin/wikilint
make test          # run the Go test suite
make lint          # run wikilint against ./wiki
make check         # full pre-commit chain: fmt-check + vet + test + lint
make fmt           # format all Go files in-place
make clean         # remove build artifacts (bin/)
make install       # install wikilint into $(GOBIN) or $(GOPATH)/bin
```

Under the hood these invoke the standard Go toolchain. If you prefer,
you can drive it directly:

```bash
go build ./...
go test ./...
go run ./cmd/wikilint -wiki ./wiki
```

On success the linter prints exactly `wikilint: OK`. On failure it
prints one issue per line to stderr in the form
`<relative-path>: <message>` and exits with a non-zero status.

## Claude Code workflow

The intended workflow when Claude Code (or any capable coding agent) is
driving the wiki looks like this:

1. **Ingest** — the `ingest-source` skill reads files in `raw/`, extracts
   entities, topics, and claims, creates or updates pages under `wiki/`,
   refreshes `wiki/index.md`, and appends an entry to `wiki/log.md`.
2. **Answer** — the `answer-from-wiki` skill answers user questions from
   `wiki/` first, consulting `raw/` only when the wiki does not already
   contain the answer. Anything new learned from `raw/` is compiled
   back into `wiki/`.
3. **Reconcile** — the `reconcile-conflicts` skill records disagreements
   between sources in a page's `## Contradictions` section instead of
   silently merging them.
4. **Lint** — the `lint-wiki` skill runs `wikilint` and fixes whatever it
   reports, in the wiki rather than in the linter.

Focused sub-agents back each workflow: a read-only `source-analyst` for
extraction from `raw/`, a `wiki-editor` that creates and refines wiki
pages, and a `contradiction-reviewer` that proposes explicit
contradiction notes instead of papering over disagreements.

## Submodule: `.claude/shared`

The `.claude/shared/` directory is a Git submodule pointing at
[`olegiv/claude-code-support-tools`](https://github.com/olegiv/claude-code-support-tools).
It contains shared Claude Code tooling — stacks, hooks, and helpers —
that is reused across multiple projects. After cloning the repo, run:

```bash
git submodule update --init --recursive
```

to populate `.claude/shared/`. (`raw/` and `wiki/` are already tracked
in git and do not need to be created.)

## License

llm-wiki-go is released under the GNU General Public License v3.0.
See [`LICENSE`](LICENSE) for the full text.

Copyright (C) 2026 Oleg Ivanchenko
