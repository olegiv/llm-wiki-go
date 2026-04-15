# Modules (built-in)

## Summary

This page consolidates the five documented built-in modules: **DB Manager**, **Developer Tools**, **Informer**, **Sentinel**, and the **Demo Mode** middleware. For the pluggable-module *framework* (module interface, context, lifecycle, template injection), see [topics/module-system.md](module-system.md). For custom-module authoring, see the same page. Demo mode is documented separately at [topics/demo-mode.md](demo-mode.md).

## DB Manager (`dbmanager`)

Admin-only direct SQL runner at `/admin/dbmanager`.

- **Query support:** SELECT, INSERT, UPDATE, DELETE, PRAGMA, EXPLAIN, WITH.
- **UX:** tabular results for reads, row-count + duration for writes, Ctrl+Enter shortcut.
- **History table** `dbmanager_query_history (id, query, user_id, executed_at, rows_affected, execution_time_ms, error)` — full audit trail.
- **Routes:** `GET /admin/dbmanager`, `POST /admin/dbmanager/execute`.
- **Security:** admin-auth middleware, CSRF on POST, demo mode blocks execution entirely.

Query classification is prefix-based: `SELECT / PRAGMA / EXPLAIN / WITH` → read; anything else → write.

## Developer Tools (`developer`)

Test-data generator for development at `/admin/developer`.

- **Generates:** 5–20 tags, 5–20 nested categories (~40% root / ~40% children / ~20% grandchildren), 5–20 placeholder images, 5–20 pages, 5–10 menu items. All translated into every active language.
- **Images:** stdlib-generated 800×600 JPEGs, 10 fixed colors, with `thumbnail`/`medium`/`large` variants. Does not produce `og`, `small`, or `grid` variants (scope limitation, not a bug).
- **Tracking table:** `developer_generated_items (id, entity_type, entity_id, created_at)` — deletion only removes module-created items.
- **Routes:** `GET /admin/developer`, `POST /admin/developer/generate`, `POST /admin/developer/delete`.
- **Policy:** development use only; admin-only in 0.18.1 hardening.

## Informer (`informer`)

Dismissible notification bar at `/admin/informer`.

- **Settings** (DB table `informer_settings`, single row): `Enabled`, `Notification Text` (HTML allowed), `Background Color` (default `#1e40af`), `Text Color` (default `#ffffff`).
- **Dismissal** via cookie `ocms_informer_dismissed` (365-day expiry); version counter increments on save → all cookies invalidated.
- **Demo mode** auto-enables with admin-access info.
- **Theme integration:** `{{informerBar}}` template function invoked after `<body>`.
- **CSS:** `@keyframes informer-spin` animation for the spinner.

## Sentinel (`sentinel`)

IP banning + auto-ban + whitelist at `/admin/sentinel`. Slot 3 in the global middleware chain (after `RequestID` + `TrustedRealIP`).

### Per-request flow

1. Skip if module inactive.
2. Whitelist match → bypass all checks.
3. Banned IP → HTTP 403.
4. Matching auto-ban path → ban + 403.

Honeypot auto-ban runs separately via the `security.honeypot_triggered` hook.

### Pattern syntax

Exact (`192.168.1.100`), wildcard (`192.168.1.*`, `192.168.*.*`), partial (`10*`).

### Default auto-ban paths

`/wp-admin*`, `/wp-login*`, `*/.env`, `*/xmlrpc.php`, `/wp-includes*`, `*/phpmyadmin*`, `*/wp-content/plugins*`.

### Honeypot integration

Bot fills hidden `_website` field on a [Form](../entities/form.md) → forms handler logs WARNING, fires `security.honeypot_triggered` with `{ip, form_slug, form_id, request_url}` → Sentinel hook handler inserts ban. Bot receives fake-success; subsequent requests → 403.

### GeoIP

Country tagging via MaxMind GeoLite2-Country. Requires `OCMS_GEOIP_DB_PATH`. See [topics/geoip.md](geoip.md).

### Database tables

`sentinel_banned_ips`, `sentinel_autoban_paths`, `sentinel_whitelist`, `sentinel_settings`.

### Template functions

`sentinelVersion`, `sentinelIsActive`, `sentinelIsIPBanned(ip)`, `sentinelIsIPWhitelisted(ip)`.

### Safety rails

- Admins cannot ban their own IP.
- Authenticated admin/editor users skip path auto-ban.
- Auto-ban notes / URL truncated to 255 chars (0.17.0).
- Ban URLs validated before rendering (0.18.1 XSS fix).

## Demo Mode

See [topics/demo-mode.md](demo-mode.md). Implemented as middleware (`internal/middleware/demo.go`) — not a pluggable module despite its `Module:` name on the GitHub wiki.

## Modules not yet documented in `docs/` or `wiki/`

- `analytics_ext` — external analytics (Google Analytics, etc.).
- `analytics_int` — internal view/read analytics (since 0.18.0, Medium-style engagement metrics).
- `embed` — content embed / iframe proxy (Dify integration).
- `example` — built-in reference module that uses the core renderer.
- `hcaptcha` — bot protection; see [topics/hcaptcha.md](hcaptcha.md).
- `migrator` — import from other CMS systems; webpage source added in 0.16.0.
- `privacy` — Klaro consent, GDPR, Google Consent Mode v2 (since 0.5.0).

These lack dedicated docs pages; authoritative source is their `modules/<name>/` directories under `ocms-go.core/` (not ingested in Phase 1).

## Sources

- [sources/docs-dbmanager-module.md](../sources/docs-dbmanager-module.md)
- [sources/wiki-module-db-manager.md](../sources/wiki-module-db-manager.md)
- [sources/docs-developer-module.md](../sources/docs-developer-module.md)
- [sources/wiki-module-developer.md](../sources/wiki-module-developer.md)
- [sources/docs-informer-module.md](../sources/docs-informer-module.md)
- [sources/wiki-module-informer.md](../sources/wiki-module-informer.md)
- [sources/docs-sentinel-module.md](../sources/docs-sentinel-module.md)
- [sources/wiki-module-sentinel.md](../sources/wiki-module-sentinel.md)
- [sources/changelog.md](../sources/changelog.md)
