# wiki/Contributing.md (oCMS)

## Summary

Wiki version of the contributor guide. Essentially the same content as `CONTRIBUTING.md`:

- **How to contribute:** report issues (duplicate check, descriptive title, repro steps, system info, logs); submit PRs (fork → feature branch from `master` → tests → commit → PR).
- **Dev setup:** clone → `make install-hooks` → `make assets` → `OCMS_SESSION_SECRET=… make dev`. Points at Getting Started for full instructions.
- **Coding standards (Go):** `make fmt`, `make check`, meaningful names, exported-function comments.
- **Testing:** tests for new functionality, existing tests pass, meaningful coverage.
- **Commit messages:** imperative verb, ≤50 char subject line, optional body. Example format given.
- **Database changes:** `make migrate-create name=…`, both up and down, regenerate `sqlc` if needed.
- **License headers:** SPDX short-form required on all new source files — `SPDX-License-Identifier: GPL-3.0-or-later` + copyright line.
- **Contributor license:** GPL-3.0.

Differences from `CONTRIBUTING.md`:

- Wiki version is slightly condensed; omits the exhaustive SPDX-header templates for JS, CSS, and HTML (docs version has all four languages; wiki shows only Go).
- Wiki uses `[[Getting Started]]` wikilink for the dev-setup redirect.

No factual disagreements.

## Sources

- Origin: `raw/ocms-go.core/wiki/Contributing.md`
