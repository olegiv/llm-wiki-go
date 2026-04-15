# oCMS

## Summary

oCMS is a lightweight Go content management system by Oleg Ivanchenko. It bundles a templ-based admin UI (HTMX + Alpine.js), SQLite storage with [goose](https://github.com/pressly/goose) migrations and [sqlc](https://sqlc.dev) generated queries, multi-level caching (in-memory + optional Redis), SCS session authentication with Argon2id password hashing, extensible theme and module systems, a REST API with bearer-token authentication, webhooks with HMAC-SHA256 signing, JSON+ZIP import/export, URL-prefixed multi-language support, and integrated SEO (sitemap, Open Graph, canonical URLs).

## Key facts

- **Current version:** 0.18.1 (2026-04-14).
- **License:** GPL-3.0-or-later.
- **Go toolchain:** 1.26 or later (enforced — see [sources/claude.md](../sources/claude.md)).
- **Storage:** SQLite via sqlc and goose. All database access goes through generated code in `internal/store/`.
- **Templates:** [templ](https://templ.guide) for type-safe Go HTML.
- **Frontend:** HTMX + Alpine.js with TinyMCE 8 rich editor and a custom SCSS framework.
- **Password hashing:** Argon2id (also used for API key hashing since 0.1.0).
- **Session store:** SCS with SQLite persistence, 24-hour session lifetime.
- **Image processing:** libvips (required for the media library).
- **Containerization:** Multi-stage Dockerfile; `docker compose` Redis profile.
- **Repository:** `github.com/olegiv/ocms-go` with wiki as a submodule at `wiki/`.

## Role model

Per [sources/security.md](../sources/security.md), three roles exist: **Admin**, **Editor**, **Public**. The admin UI only manages the first two; Public is the unauthenticated default.

## Core entities

- [Page](page.md) — primary content entity (versions, drafts, scheduling, translations).
- [Menu](menu.md) — hierarchical, per-language navigation.
- [Category](category.md) — hierarchical taxonomy.
- [Tag](tag.md) — flat taxonomy.

## Main topic pages

- [Architecture layers](../topics/architecture-layers.md)
- [Configuration](../topics/configuration.md)
- [Content management](../topics/content-management.md)
- [Admin interface](../topics/admin-interface.md)
- [Getting started](../topics/getting-started.md)
- [Release history](../topics/release-history.md)
- [Security overview](../topics/security-overview.md)

## Contradictions

### Supported versions
- [sources/security.md](../sources/security.md): "0.14.x and 0.12.x" are supported; "< 0.12" is unsupported.
- [sources/changelog.md](../sources/changelog.md): Current release is 0.18.1 (2026-04-14); 0.15.0, 0.16.0, 0.17.0, and 0.18.0 all shipped between 2026-02-05 and 2026-04-09.

The `SECURITY.md` supported-versions table is stale — it predates the 0.15.x–0.18.x series.

### Role count phrasing
- [sources/readme.md](../sources/readme.md) describes user management as "admin/editor" (two).
- [sources/security.md](../sources/security.md) defines RBAC as "Admin, Editor, and Public" (three).

Probable resolution: README describes the admin UI's user management screen (which only creates Admin or Editor accounts), while SECURITY describes the full authorization model including unauthenticated Public access. Not factually contradictory, but the phrasing is inconsistent.

## Open questions

- Which versions are actually supported today? The `SECURITY.md` table likely needs an update to cover 0.17.x and 0.18.x.

## Sources

- [sources/readme.md](../sources/readme.md)
- [sources/changelog.md](../sources/changelog.md)
- [sources/claude.md](../sources/claude.md)
- [sources/security.md](../sources/security.md)
- [sources/agents.md](../sources/agents.md)
- [sources/contributing.md](../sources/contributing.md)
- [sources/wiki-home.md](../sources/wiki-home.md)
