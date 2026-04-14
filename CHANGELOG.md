# Changelog

All notable changes to `llm-wiki-go` are documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

The first tagged release will promote `[Unreleased]` to
`[0.1.0] - YYYY-MM-DD` once the tag is cut.

## [Unreleased]

### Added

- `wikilint` CLI with `-wiki` and `-version` flags. Structural checks
  cover single H1, required `## Sources` section, Markdown and Obsidian
  `[[wikilink]]` resolution, unique normalized title slugs, and
  orphan-page detection (every page must be reachable from
  `wiki/index.md`).
- `internal/wiki` helpers supporting the ingest / answer / reconcile
  workflows used by Claude Code skills.
- `internal/version` package that derives a one-line build descriptor
  from `runtime/debug.BuildInfo`.
- Makefile targets: `setup`, `build`, `test`, `fmt`, `fmt-check`,
  `vet`, `lint`, `check`, `tidy`, `clean`, `install`, `help`.
- Claude Code skills (`ingest-source`, `answer-from-wiki`,
  `reconcile-conflicts`, `lint-wiki`) and sub-agents
  (`source-analyst`, `wiki-editor`, `contradiction-reviewer`) under
  `.claude/`.
- CI workflows: `go.yml` (build, vet, test, wikilint) and
  `dependency-review.yml`.
- Dependabot config for `gomod`, `github-actions`, and `gitsubmodule`.
- `LICENSE` (GPL-3.0), `SECURITY.md`, `CONTRIBUTING.md`, and Go /
  CodeQL / Dependency-review badges in `README.md`.
