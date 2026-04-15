# AGENTS.md (oCMS)

## Summary

Short repo-guidelines document for AI coding agents working on oCMS. Lists:

- **Project structure.** `cmd/ocms/` (entrypoint), `internal/` (core runtime: handlers, middleware, services, store, views, cache, scheduler), `modules/` (built-in modules: analytics, embed, privacy, migrator, etc.), `custom/` (user-defined modules/themes), `web/` (shared templates and frontend assets), `internal/store/migrations/` and `internal/store/queries/` (goose + sqlc sources), `docs/` (feature/deployment/security docs).
- **Build/test/dev commands.** `make dev`, `make run`, `make build` / `make build-prod`, `make test`, `make assets`, `make migrate-up` / `-down` / `-status`, `make install-hooks`.
- **Coding style & naming.** Idiomatic Go, `gofmt`, package-oriented structure, `golangci-lint` configured in `.golangci.yml`. CamelCase exports, mixedCaps internal, concise package names. Test files `_test.go` alongside implementation. Do not hand-edit generated files (`*_templ.go`, `*.sql.go`).
- **Testing.** Go standard `testing` package. Run with `OCMS_SESSION_SECRET=test-secret-key-32-bytes-long!!! go test ./...` or `make test`. Cover new handlers, middleware, store queries, module behavior, and permission/security paths.
- **Commits & PRs.** Imperative subject lines, short and specific (example style: `Add CSP nonce wiring`, `Fix code quality issues`). PRs must include purpose, key changes, test evidence, linked issues; UI/template/theme changes need screenshots. Run `make check-no-absolute-paths` to catch committed local paths (`/Users/…`, `/home/…`, `C:\Users\…`).

AGENTS.md is deliberately terse — a quick-reference overlay on top of the richer `CLAUDE.md`.

## Sources

- Origin: `raw/ocms-go.core/top-level/AGENTS.md`
