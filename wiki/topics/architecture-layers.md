# Architecture layers

## Summary

oCMS follows a conventional layered request flow:

```
HTTP request → chi router → middleware chain → handler → service → store (sqlc) → SQLite
```

### Key `internal/` packages

- **`handler/`** — HTTP handlers for the admin UI and REST API (under `handler/api/`). Handler structs receive `*sql.DB`, `*render.Renderer`, and `*scs.SessionManager`, and call `store.New(db)` to get sqlc query access. Admin views use templ components in `internal/views/admin/`.
- **`middleware/`** — standard middleware chain: `SecurityHeaders`, `CSRF`, `Auth`, `LoadUser`, `LoadSiteConfig`. REST API routes use a parallel chain: `APIKeyAuth`, `RequirePermission`, `APIRateLimit`.
- **`service/`** — business logic for media, menus, forms, webhooks, events, and related domain operations.
- **`store/`** — database access via sqlc-generated code from SQL in `store/queries/`. Goose migrations live in `store/migrations/`.
- **`model/`** — domain types (Page, Media, Menu, User, APIKey, Form, Webhook, Config, Language, Translation, Event).
- **`cache/`** — multi-level cache: in-memory primary, optional Redis for distributed deployments (gated by `OCMS_REDIS_URL`).
- **`session/`** — SCS session manager backed by SQLite.
- **`theme/`** and **`themes/`** — runtime theme loader; `themes/` holds the embedded core themes (`default`, `developer`). Custom themes live under `custom/themes/`.
- **`module/`** — module registry and lifecycle hooks. Custom modules self-register via `init()` + blank imports in `custom/modules/imports.go`.
- **`scheduler/`** — cron-based task scheduler with admin registry.
- **`render/`** — template rendering with CSP nonce support.
- **`seo/`** — sitemap, robots.txt, and meta-tag generation.
- **`i18n/`** — admin UI internationalization (English and Russian bundled).
- **`imaging/`** — image processing (thumbnails, variants) via libvips.
- **`transfer/`** — import/export (JSON and ZIP).
- **`webhook/`** — webhook delivery engine with HMAC signing and retries.

### Binary layout

`cmd/ocms/main.go` is thin: it parses configuration, wires dependencies, and delegates to `internal/`. Cross-compile targets exist for Linux AMD64 and macOS ARM64 (`make build-linux-amd64`, `make build-darwin-arm64`).

### Middleware order matters

The chain executes top-to-bottom. `SecurityHeaders` sets CSP (including a per-request nonce), HSTS, and other hardening headers before any other middleware runs, so even short-circuit responses from `CSRF` or `Auth` carry the correct headers.

## Sources

- [sources/claude.md](../sources/claude.md)
- [sources/agents.md](../sources/agents.md)
- [sources/readme.md](../sources/readme.md)
