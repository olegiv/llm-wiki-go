# Getting started

## Summary

End-to-end onboarding for a fresh oCMS clone.

## Prerequisites

- **Go** 1.26 or later. Do not downgrade `go.mod` — see [sources/claude.md](../sources/claude.md).
- **Node.js (npm)** for frontend dependencies (HTMX, Alpine.js, TinyMCE assets).
- **sqlc** for SQL code generation.
- **templ** for type-safe HTML templates.
- **goose** for database migrations.
- **Dart Sass** for SCSS compilation.
- **libvips** for image processing (required for the media library).

libvips install paths:

| Platform | Command |
|----------|---------|
| macOS | `brew install vips` |
| Ubuntu/Debian | `sudo apt-get install libvips-dev` |
| Fedora | `sudo dnf install vips-devel` |

## Installation

1. Clone `github.com/olegiv/ocms-go` (recurse submodules if you want the wiki).
2. `go mod download`.
3. `go install` the three generators:
   - `github.com/sqlc-dev/sqlc/cmd/sqlc@latest`
   - `github.com/a-h/templ/cmd/templ@latest`
   - `github.com/pressly/goose/v3/cmd/goose@latest`
4. `sqlc generate && templ generate`.
5. `make assets` — installs npm deps and compiles SCSS.

## First run

```bash
export OCMS_SESSION_SECRET="your-secret-key-at-least-32-bytes"
make install-hooks   # once per clone
make dev             # asset rebuild + dev server
# or
make run             # server only
```

The server starts at `http://localhost:8080`. Set `OCMS_DO_SEED=true` on first run to seed the default admin user.

## Default credentials

| Role | Email | Password |
|------|-------|----------|
| Admin (default seed) | `admin@example.com` | `changeme1234` |
| Demo admin (demo mode) | `demo@example.com` | `demo1234demo` |
| Demo editor (demo mode) | `editor@example.com` | `demo1234demo` |

Rotate immediately after first login. Production startup is blocked on default admin credentials (since 0.9.0).

## Make command surface

See [sources/readme.md](../sources/readme.md) and [sources/wiki-getting-started.md](../sources/wiki-getting-started.md) for the full list. Key targets:

- `make dev` / `make run` / `make stop` / `make restart`.
- `make build` / `make build-prod` / `make build-linux-amd64` / `make build-darwin-arm64`.
- `make test` — requires `OCMS_SESSION_SECRET`.
- `make migrate-up` / `-down` / `-status` / `-create`.
- `make sqlc` / `make templ` / `make assets`.

## Testing

```bash
OCMS_SESSION_SECRET=test-secret-key-32-bytes-long!! go test ./...
govulncheck ./...
```

## Next

- [topics/configuration.md](configuration.md) — environment variable reference.
- [topics/architecture-layers.md](architecture-layers.md) — request flow and `internal/` package layout.
- [topics/admin-interface.md](admin-interface.md) — admin panel tour.

## Sources

- [sources/wiki-getting-started.md](../sources/wiki-getting-started.md)
- [sources/readme.md](../sources/readme.md)
- [sources/claude.md](../sources/claude.md)
