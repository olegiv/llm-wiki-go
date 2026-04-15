# CLAUDE.md (oCMS)

## Summary

Claude Code workflow and architecture guide for oCMS contributors. Documents:

- **Build / dev commands.** `make dev` (assets + server), `make run` (server only), `make build`, `make test` with required `OCMS_SESSION_SECRET`, database migrations via goose, code generation via `sqlc generate` / `templ generate`.
- **Go toolchain rules.** Local Go must match `go.mod`; never downgrade; never set `GOTOOLCHAIN=local` as a workaround.
- **Pre-commit hook protocol.** `.git/hooks/pre-commit` blocks non-interactive commits; user-initiated commits use `--no-verify`.
- **Request flow.** HTTP → chi router → middleware chain → handler → store (sqlc) → SQLite.
- **Middleware chain.** Admin: `SecurityHeaders → CSRF → Auth → LoadUser → LoadSiteConfig`. API: `APIKeyAuth → RequirePermission → APIRateLimit`.
- **Package dependency tree** rooted at `cmd/ocms/main.go`.
- **25 environment variables** (subset of the full README reference).
- **Admin routes, REST API endpoints, health-check routes, SEO routes.**
- **Post-change testing workflow:** rebuild assets, restart server, curl endpoints.
- **Code style rules:** SPDX headers, error wrapping with `%w`, structured slog logging, naming conventions, typed context keys, resource-leak defense via `defer rows.Close()` in callers.
- **Code quality rules:** forbid `//nolint:dupl`, forbid empty slice literals (`[]T{}`), forbid variable names colliding with imported packages.
- **Wiki sync requirement.** Editing any `docs/*.md`, `README.md`, `SECURITY.md`, or `CONTRIBUTING.md` requires updating the corresponding `wiki/*.md` page and running `/update-wiki`.
- **Seven sub-agents** (`test-runner`, `db-manager`, `api-developer`, `module-developer`, `security-auditor`, `code-quality-auditor`, `frontend-developer`) and **14 slash commands** (`/test`, `/build`, `/migrate`, `/sqlc-generate`, `/dev-server`, `/api-test`, `/security-audit`, `/code-quality`, `/commit-prepare`, `/templui-add`, `/templui-list`, `/update-wiki`, `/clean`, `/fly-deploy`).

## Sources

- Origin: `raw/ocms-go.core/top-level/CLAUDE.md`
