# Demo mode

## Summary

Demo mode (`OCMS_DEMO_MODE=true`, requires `OCMS_DO_SEED=true`) seeds sample content on first start and locks down destructive admin operations so visitors can explore the admin UI without damaging data. It is implemented as middleware (`internal/middleware/demo.go`) and handler guards (`internal/handler/demo.go`), **not** as a pluggable module, despite being described as "Module: Demo Mode" on the GitHub wiki.

Shipped in 0.7.0 ("Read-only content protection"). Hardened across 0.7.0 (form CSV export blocked, SQL execution blocked, IP masking in events), 0.8.0 (bookmarks + theme settings read-only), 0.18.1 (IP leak fix in events ban UI, demo password rotated on startup, stale demo credentials scrubbed from seeded homepage, disallow production seeding in demo mode).

## Seeded content

On first start with both seeding flags set:

| Content | Count | Details |
|---------|-------|---------|
| Users | 2 | Admin + editor with demo credentials. |
| Categories | 4 | Blog, Portfolio, Services, Resources. |
| Tags | 7 | Tutorial, News, Featured, Go, Web Development, Design, Open Source. |
| Pages | 9 | Home, About, Contact, 3 blog posts, 2 portfolio items, 1 services page. |
| Media | 10 | 2400×1600 placeholder images with all variants. |
| Menu items | 6 | Home, Blog, Portfolio, Services, About, Contact. |

## Credentials

| Role | Email | Password |
|------|-------|----------|
| Admin | `demo@example.com` | `demo1234demo` |
| Editor | `editor@example.com` | `demo1234demo` |
| Default admin | `admin@example.com` | `changeme1234` (rotated on startup in demo mode since 0.18.1) |

## What is blocked

Per the reconciled view (see [Contradictions](#contradictions) below):

- **Content writes:** Pages, tags, categories, menus, media, forms, widgets — **delete** is blocked; create/edit policy differs between sources.
- **User management** — blocked.
- **Site configuration** — blocked.
- **Language management** — blocked.
- **API key management** — blocked.
- **Webhook management** — blocked.
- **Module management / toggle / settings** — blocked.
- **Cache clearing** — blocked (view stats still allowed).
- **Import / export** — blocked.
- **DB Manager SQL execution** — blocked.
- **Form submission CSV export** — blocked.
- **Scheduler edit / reset / trigger** — blocked (see [scheduler topic](scheduler.md)).
- **Media uploads** — 2 MB cap (vs 20 MB normal).

### Still allowed

- Browse the admin UI.
- Theme activation (swap between themes).
- REST API **reads**.

## Implementation

- `internal/middleware/demo.go` — middleware + `IsDemoMode()` helper + `DemoRestriction` constants + `DemoModeMessageDetailed()`.
- `internal/handler/demo.go` — `demoGuard(w, r, renderer, restriction, redirectURL)` helper for Web UI; JSON/HTMX handlers use `middleware.IsDemoMode()` + `writeJSONError(…, http.StatusForbidden, …)`.

### Adding a new restriction

1. Define `const RestrictionMyAction DemoRestriction = "my_action"` in `internal/middleware/demo.go`.
2. Add the message to `DemoModeMessageDetailed`.
3. Wire the guard in the handler.

### User experience

- Web UI: redirect back with a flash message ("X is disabled in demo mode").
- API: HTTP 403 with explanation body.

## Scheduled reset

Recommended for public demos: cron a daily DB reset (`reset-demo.sh` + `fly machines restart`) so spam content does not accumulate.

## Contradictions

### What content operations are blocked?

- [sources/docs-demo-deployment.md](../sources/docs-demo-deployment.md) restrictions table: Content "Create, edit, delete, unpublish" are **all blocked** (aligns with 0.7.0 changelog: "Read-only content protection (blocks create/edit/delete/unpublish)").
- [sources/docs-demo-mode.md](../sources/docs-demo-mode.md) and [sources/wiki-module-demo-mode.md](../sources/wiki-module-demo-mode.md): Create/Edit pages is **Allowed**; only **Delete** is blocked.

These two descriptions disagree on Create/Edit behavior. The restrictions table in `docs/demo-mode.md` itself allows create/edit of pages, tags, categories, menus, forms — but blocks deletion of same.

Likely resolution: the behavior changed between 0.7.0 and a later release — relaxing the create/edit block to only block deletions while preserving the "no abuse" guarantee via scheduled reset. `docs/demo-deployment.md` was not updated to reflect the relaxation.

**For current behavior**, trust `internal/middleware/demo.go` as the authoritative reference (to be ingested in Phase 2).

## Open questions

- When exactly did create/edit become allowed? Needs `git log internal/middleware/demo.go` or a spot-check in Phase 2 source ingest.

## Sources

- [sources/docs-demo-mode.md](../sources/docs-demo-mode.md)
- [sources/docs-demo-deployment.md](../sources/docs-demo-deployment.md)
- [sources/wiki-module-demo-mode.md](../sources/wiki-module-demo-mode.md)
- [sources/wiki-deploy-fly-io.md](../sources/wiki-deploy-fly-io.md)
- [sources/changelog.md](../sources/changelog.md)
