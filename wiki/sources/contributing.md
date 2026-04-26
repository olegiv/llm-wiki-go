# CONTRIBUTING.md (oCMS)

## Summary

Contributor guide for oCMS.

- **Issue reporting.** Check duplicates, use descriptive titles, include reproduction steps and system info (Go version, OS, browser when relevant), attach logs.
- **Pull request flow.** Fork → create feature branch from **`master`** → implement → add/update tests → run `OCMS_SESSION_SECRET=test-secret-key-32-bytes-long!! go test ./...` → commit with clear message → push to fork → open PR.
- **Dev setup.** Quick path: clone, `make install-hooks`, `make assets`, `OCMS_SESSION_SECRET=… make dev`.
- **Coding standards.** Idiomatic Go, `make fmt` and `make check` before committing, meaningful variable/function names, comments on exported functions and complex logic.
- **Testing.** Write tests for new functionality; ensure existing tests pass; aim for meaningful coverage.
- **Commit messages.** Imperative verb (Add / Fix / Update / Remove), first line < 50 characters, optional body.
- **Database changes.** Create goose migration via `make migrate-create name=…`, write both up and down migrations, regenerate sqlc queries if needed.
- **SPDX license headers.** Required on all new Go, JS/TS, CSS/SCSS, and HTML template source files in the form `SPDX-License-Identifier: GPL-3.0-or-later` plus copyright line.

Contributions are licensed under GPL-3.0.

## Sources

- Origin: `raw/ocms-go.core/top-level/CONTRIBUTING.md`
