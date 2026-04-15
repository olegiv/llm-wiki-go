# Change Log

## 2026-04-15 — Batch 10: import/export, Contributing, _Sidebar

Ingested 4 files: `docs/import-export.md`, `wiki/Import-Export.md`, `wiki/Contributing.md`, `wiki/_Sidebar.md`.

### Created pages

- Topic: `import-export.md`.
- Sources: `docs-import-export.md`, `wiki-import-export.md`, `wiki-contributing.md`, `wiki-sidebar.md`.

### Contradictions surfaced this batch

- **Page import/export schema — video fields.** `docs/import-export.md` page export includes `video_url` and `video_title` (added 0.16.0); `wiki/Import-Export.md` omits both. Recorded on `topics/import-export.md`.
- **Media import/export schema — `path` field.** `docs/import-export.md` includes `path`; `wiki/Import-Export.md` omits it. Same drift direction.

### Parity noted (no contradiction)

- `wiki/Contributing.md` matches `CONTRIBUTING.md` on every substantive point; wiki version drops SPDX header templates for JS/CSS/HTML (docs shows all four languages, wiki shows only Go).
- `wiki/_Sidebar.md` is navigation-only (no H1 in the original). Its "Built-in Modules" section matches five of the eleven on-disk `modules/` directories — same discrepancy already captured on `topics/module-system.md`.

## 2026-04-15 — Batch 9: individual modules (DB Manager, Developer, Informer, Sentinel)

Ingested 8 files: `docs/dbmanager-module.md`, `wiki/Module:-DB-Manager.md`, `docs/developer-module.md`, `wiki/Module:-Developer.md`, `docs/informer-module.md`, `wiki/Module:-Informer.md`, `docs/sentinel-module.md`, `wiki/Module:-Sentinel.md`.

### Created pages

- Topic: `modules.md` (consolidates the four documented built-in modules plus pointers to Demo Mode and undocumented modules).
- Sources: `docs-dbmanager-module.md`, `wiki-module-db-manager.md`, `docs-developer-module.md`, `wiki-module-developer.md`, `docs-informer-module.md`, `wiki-module-informer.md`, `docs-sentinel-module.md`, `wiki-module-sentinel.md`.

### Parity noted (no contradiction)

- All four docs/wiki pairs in this batch are in factual parity. Wiki versions drop DDL snippets, i18n key catalogs, module-structure trees, and operational notes.
- Developer module's generated image set (`thumbnail`, `medium`, `large`) is narrower than the full seven media variants in `docs/media.md`. Scope limitation of the test-data generator, not a documentation disagreement.

## 2026-04-15 — Batch 8: themes, modules, demo mode

Ingested 8 files: `docs/custom-themes.md`, `wiki/Theme-System.md`, `docs/custom-modules.md`, `wiki/Module-System.md`, `docs/embed-themes-migration.md`, `docs/demo-deployment.md`, `docs/demo-mode.md`, `wiki/Module:-Demo-Mode.md`.

### Created pages

- Topics: `theme-system.md`, `module-system.md`, `demo-mode.md`.
- Sources: `docs-custom-themes.md`, `wiki-theme-system.md`, `docs-custom-modules.md`, `wiki-module-system.md`, `docs-embed-themes-migration.md`, `docs-demo-deployment.md`, `docs-demo-mode.md`, `wiki-module-demo-mode.md`.

### Contradictions surfaced this batch

- **Demo-mode content-write policy.** `docs/demo-deployment.md` blocks create/edit/delete/unpublish of content; `docs/demo-mode.md` and `wiki/Module:-Demo-Mode.md` allow create/edit and only block delete. Recorded on `topics/demo-mode.md`.
- **Built-in module inventory.** `wiki/Module-System.md` lists 10 modules; on-disk `modules/` contains 11 entries with `example` and `migrator` unlisted and "Demo Mode" present in the wiki but implemented as middleware, not as a module. Recorded on `topics/module-system.md`.

### Parity noted (no contradiction)

- `docs/custom-themes.md` vs `wiki/Theme-System.md` — factually identical.
- `docs/custom-modules.md` vs `wiki/Module-System.md` — wiki omits templ layout injection, embedded admin templates, and testing helpers (scope differences, not conflicts).
- `docs/embed-themes-migration.md` is a **historical design plan** from before 0.2.0; the migration has since shipped.

## 2026-04-15 — Batch 7: deployment (Docker, Fly.io, Ubuntu+Plesk, reverse proxy)

Ingested 7 files: `docs/deploy-ubuntu-plesk.md`, `wiki/Deploy:-Ubuntu-Plesk.md`, `wiki/Deployment.md`, `wiki/Deploy:-Fly.io.md`, `wiki/Docker.md`, `docs/reverse-proxy.md`, `wiki/Reverse-Proxy.md`.

### Created pages

- Topics: `deployment.md`, `reverse-proxy.md`.
- Sources: `docs-deploy-ubuntu-plesk.md`, `wiki-deploy-ubuntu-plesk.md`, `wiki-deployment.md`, `wiki-deploy-fly-io.md`, `wiki-docker.md`, `docs-reverse-proxy.md`, `wiki-reverse-proxy.md`.

### Contradictions surfaced this batch

- **`X-XSS-Protection` header.** `docs/reverse-proxy.md` recommends adding it to Nginx and Apache security-header blocks; `wiki/Reverse-Proxy.md` omits it and explicitly warns against duplicating oCMS-set headers at the proxy. The header is deprecated in modern browsers; wiki is the better reference. Recorded on `topics/reverse-proxy.md`.

### Parity noted (no contradiction)

- `docs/deploy-ubuntu-plesk.md` is a **stub** pointing at `scripts/deploy/README.md` (consolidated in 0.10.2); the real content lives there and is mirrored by the wiki page. Not a conflict — a deliberate pointer. Captured on `topics/deployment.md`.

## 2026-04-15 — Batch 6: REST API, webhooks, scheduler, GeoIP

Ingested 7 files: `wiki/REST-API.md`, `docs/webhooks.md`, `wiki/Webhooks.md`, `docs/scheduler.md`, `wiki/Scheduler.md`, `docs/geoip.md`, `wiki/GeoIP.md`.

### Created pages

- Entities: `api-key.md`, `webhook.md`, `webhook-delivery.md`.
- Topics: `rest-api.md`, `webhooks.md`, `scheduler.md`, `geoip.md`.
- Sources: `wiki-rest-api.md`, `docs-webhooks.md`, `wiki-webhooks.md`, `docs-scheduler.md`, `wiki-scheduler.md`, `docs-geoip.md`, `wiki-geoip.md`.

### Parity noted (no contradiction)

- `docs/webhooks.md` vs `wiki/Webhooks.md` — facts identical; wiki version surfaces `OCMS_WEBHOOK_FORM_DATA_MODE` and `OCMS_WEBHOOK_ALLOWED_HOSTS` inline, docs relegates them to Forms doc.
- `docs/scheduler.md` vs `wiki/Scheduler.md` — identical behavior; wiki omits the Go `AddFunc`/`ModuleCron` snippets and the SQL DDL for `scheduler_overrides`.
- `docs/geoip.md` vs `wiki/GeoIP.md` — identical procedure and troubleshooting.

No new contradictions surfaced in this batch.

## 2026-04-15 — Batch 5: i18n, multi-language, SEO

Ingested 5 files: `docs/i18n.md`, `wiki/Internationalization.md`, `docs/multi-language.md`, `wiki/Multi-Language.md`, `wiki/SEO.md`.

### Created pages

- Entities: `language.md`, `translation.md`.
- Topics: `i18n.md`, `multi-language.md`, `seo.md`.
- Sources: `docs-i18n.md`, `wiki-internationalization.md`, `docs-multi-language.md`, `wiki-multi-language.md`, `wiki-seo.md`.

### Contradictions surfaced this batch

- **OG variant — three-way reinforcement.** `wiki/SEO.md` also claims `og:image` uses the `large` variant (matching `wiki/Media-Library.md`), but `docs/media.md` documents the dedicated `og` variant added in 0.18.0 with `og > large > medium > thumbnail` priority. Extension of the Batch 3 contradiction. Recorded on `topics/seo.md`.
- **`Content-Language` header coverage.** `docs/multi-language.md` documents the HTTP response header; `wiki/Multi-Language.md` omits the "Language Headers" section entirely. Recorded on `topics/multi-language.md`.

### Parity noted (no contradiction)

- `docs/i18n.md` vs `wiki/Internationalization.md` — facts identical (same file format, same helpers, same supported languages). Wiki is condensed.

## 2026-04-15 — Batch 4: caching, CSRF, login, hCaptcha, Security

Ingested 9 files: `docs/caching.md`, `wiki/Caching.md`, `docs/csrf.md`, `wiki/CSRF-Protection.md`, `docs/login-security.md`, `wiki/Login-Security.md`, `wiki/Security.md`, `docs/hcaptcha.md`, `wiki/hCaptcha.md`.

### Created pages

- Topics: `caching.md`, `csrf.md`, `login-security.md`, `hcaptcha.md`.
- Sources: `docs-caching.md`, `wiki-caching.md`, `docs-csrf.md`, `wiki-csrf-protection.md`, `docs-login-security.md`, `wiki-login-security.md`, `wiki-security.md`, `docs-hcaptcha.md`, `wiki-hcaptcha.md`.

### Updated pages

- `topics/security-overview.md` — added `wiki-security.md` as source; added links to new per-topic pages (CSRF, login, hCaptcha); added Production Checklist from `wiki/Security.md`; added open-redirect CHANGELOG citation.

### Contradictions surfaced this batch

- **hCaptcha test keys as defaults.** `docs/hcaptcha.md` and `wiki/hCaptcha.md` both state that hCaptcha test keys (site `10000000-ffff-ffff-ffff-000000000001`, secret `0x0000...`) are "used as defaults when no keys are configured" and that "the widget always passes automatically". `CHANGELOG.md` 0.18.1 explicitly removes these insecure defaults and scrubs persisted copies on upgrade. Both docs files are stale. Recorded on `topics/hcaptcha.md`.
- **`OCMS_HCAPTCHA_DISABLED` env var coverage.** Documented in `docs/hcaptcha.md` and `README.md`; absent from `wiki/hCaptcha.md` and `wiki/Configuration.md`. Recorded on `topics/hcaptcha.md` (extends the Batch 2 observation).

### Parity noted (no contradiction)

- `docs/caching.md` vs `wiki/Caching.md` — facts are identical (6 cache types, same role levels, same env var trio, same admin-bypass behavior). Wiki omits Go snippets; docs includes them. Canonical contradiction test case turned out clean.
- `docs/csrf.md` vs `wiki/CSRF-Protection.md` — facts identical; wiki omits library name (`filippo.io/csrf/gorilla`), `CSRFConfig` Go struct, and session-secret reuse note. No factual disagreement.
- `docs/login-security.md` vs `wiki/Login-Security.md` — facts identical (same thresholds, same messages, same log events). Wiki omits the `middleware.NewLoginProtection` snippet and the "future admin unlock" forward-looking note; wiki adds a `OCMS_TRUSTED_PROXIES` best-practice bullet that docs omits.
- `wiki/Security.md` vs `SECURITY.md` — content parity; wiki adds a 10-item Production Security Checklist.

## 2026-04-15 — Batch 3: admin, media, forms, video

Ingested 7 files: `docs/admin-bulk-actions.md`, `docs/admin-list-sorting.md`, `docs/media.md`, `wiki/Media-Library.md`, `docs/forms.md`, `wiki/Forms.md`, `docs/video-embedding.md`.

### Created pages

- Entities: `media.md`, `form.md`.
- Topics: `video-embedding.md`.
- Sources: `docs-admin-bulk-actions.md`, `docs-admin-list-sorting.md`, `docs-media.md`, `wiki-media-library.md`, `docs-forms.md`, `wiki-forms.md`, `docs-video-embedding.md`.

### Updated pages

- `entities/page.md` — now references `media.md` entity and `video-embedding.md` topic; adds media and video sources.
- `topics/admin-interface.md` — adds `docs-admin-bulk-actions.md` and `docs-admin-list-sorting.md` as sources; adds error response shape and sort pill-highlight detail.
- `index.md` — sources split into Top-level, `docs/`, `wiki/` groups.

### Contradictions surfaced this batch

- **Media variant count.** `docs/media.md` documents **seven** variants (`originals`, `large`, `og`, `small`, `grid`, `thumbnail`, plus `medium`); `wiki/Media-Library.md` documents **four** (`originals`, `large`, `medium`, `thumbnail`). Corroborating `CHANGELOG.md` entries (0.4.0 small, 0.10.0 grid, 0.18.0 og) confirm docs is current; wiki is stale. Recorded on `entities/media.md`.
- **OG image source.** `docs/media.md` selects the dedicated `og` variant; `wiki/Media-Library.md` uses `large` for OG and claims "social platforms resize anyway". Recorded on `entities/media.md`.
- **Upload directory configurability.** `docs/media.md` states "not configurable via environment variables"; `wiki/Media-Library.md`, `README.md`, and `CHANGELOG.md` 0.16.0 all document `OCMS_UPLOADS_DIR`. `docs/media.md` is stale. Recorded on `entities/media.md`.
- **Video provider support.** `docs/video-embedding.md` lists YouTube Supported and Vimeo Planned (no Dailymotion); `README.md`, `CHANGELOG.md` 0.16.0, and `wiki/Home.md` all say YouTube + Vimeo + Dailymotion are supported. Recorded on `topics/video-embedding.md`.

### Parity noted (no contradiction)

- `docs/forms.md` and `wiki/Forms.md` are near-identical — only link style differs (markdown vs Obsidian wikilinks). Cited on `entities/form.md` as two sources in parity.

## 2026-04-15 — Batch 2: getting started, config, content, admin, taxonomy

Ingested 7 GitHub wiki pages: `Home.md`, `Getting-Started.md`, `Configuration.md`, `Content-Management.md`, `Admin-Interface.md`, `Menu-Builder.md`, `Taxonomy.md`.

### Created pages

- Entities: `page.md`, `menu.md`, `category.md`, `tag.md`.
- Topics: `admin-interface.md`, `getting-started.md`, `content-management.md`.
- Sources: `wiki-home.md`, `wiki-getting-started.md`, `wiki-configuration.md`, `wiki-content-management.md`, `wiki-admin-interface.md`, `wiki-menu-builder.md`, `wiki-taxonomy.md`.

### Updated pages

- `entities/ocms.md` — now cites `wiki-home.md`; links to the new entity and topic pages.
- `topics/configuration.md` — three-way source reconciliation. `OCMS_API_RATE_LIMIT` promoted into the main *Security and API* table (documented in CLAUDE + wiki, missing from README); `OCMS_HCAPTCHA_DISABLED` drift noted in the opposite direction.
- `index.md` — entities and topics lists populated; sources split into *Top-level docs* and *Wiki pages*.

### Contradictions surfaced this batch

- **`OCMS_API_RATE_LIMIT` coverage.** Present in `CLAUDE.md` and `wiki/Configuration.md`, absent from `README.md`. Recorded on `topics/configuration.md`.
- **`OCMS_HCAPTCHA_DISABLED` coverage.** Present in `README.md`, absent from `wiki/Configuration.md`. Recorded on `topics/configuration.md`.
- **Variable-catalog scope.** README ≈ 55+, wiki ≈ 40, CLAUDE ≈ 25. Three-way scope disagreement, recorded on `topics/configuration.md`.

### Deviation from plan

- Initially the plan listed `menu-item` as a standalone entity. With only wiki-level prose as source, menu items are documented as nested behavior inside [entities/menu.md](entities/menu.md). A standalone page can be introduced once `internal/model/menu.go` is ingested in Phase 2.
- The plan did not call out `category` and `tag` as separate entities; adding both surfaced naturally from `wiki/Taxonomy.md`. No topic page was created for taxonomy — the two entity pages cover it.

## 2026-04-15 — Batch 1: foundation docs

Ingested top-level meta-documentation from `ocms-go.core`: `README.md`, `CHANGELOG.md`, `CLAUDE.md`, `AGENTS.md`, `SECURITY.md`, `CONTRIBUTING.md` (6 files).

### Created pages

- `entities/ocms.md` — the project as an entity.
- `topics/architecture-layers.md` — request flow and `internal/` package layout.
- `topics/security-overview.md` — authentication, request security, webhook and embed hardening.
- `topics/configuration.md` — full `OCMS_*` environment variable reference.
- `topics/release-history.md` — version table 0.0.0 through 0.18.1.
- `sources/readme.md`, `sources/changelog.md`, `sources/claude.md`, `sources/agents.md`, `sources/security.md`, `sources/contributing.md`.
- `index.md` replaced the seed with populated Entities / Topics / Sources / Log sections.

### Contradictions surfaced

- **Supported versions.** `SECURITY.md` lists 0.14.x and 0.12.x; `CHANGELOG.md` is on 0.18.1. Recorded on `entities/ocms.md` and `topics/release-history.md`.
- **Role count.** `README.md` says "admin/editor" (two); `SECURITY.md` says "Admin, Editor, Public" (three). Recorded on `entities/ocms.md`.
- **Environment variable catalog.** `README.md` ≈ 55 vars; `CLAUDE.md` ≈ 25 vars. `OCMS_API_RATE_LIMIT` appears only in `CLAUDE.md`. Recorded on `topics/configuration.md`.
