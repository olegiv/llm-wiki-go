# docs/sentinel-module.md (oCMS)

## Summary

Reference for the Sentinel IP-security module — bans, auto-bans, whitelist, and honeypot integration.

### Scope

- Manual IP banning with wildcard patterns.
- Path-based auto-ban for suspicious URLs (`/wp-admin*`, `*/phpmyadmin*`, etc.).
- Honeypot-triggered auto-ban via the `security.honeypot_triggered` hook.
- IP whitelist bypassing all checks.
- Optional GeoIP country tagging (MaxMind GeoLite2).
- In-memory caching for fast lookup.
- Event logging of auto-ban actions.
- Safety rails: admins can't ban own IP; authenticated admin/editor users skip path auto-ban.

### Middleware position

Slot 3 in the global chain, after `RequestID` and `TrustedRealIP`. Per-request behavior:

1. If module inactive → skip.
2. Whitelist match → bypass all checks.
3. Banned IP → HTTP 403.
4. Matching auto-ban path → ban + 403.

Honeypot auto-banning fires separately through the hook system.

### Admin UI

At `/admin/sentinel` (a.k.a. **Admin > Modules > Sentinel**).

### Settings (all toggleable)

| Setting | Default |
|---------|---------|
| IP Ban Check | Enabled |
| Auto-Ban by Path | Enabled |
| Auto-Ban on Honeypot | Enabled |

### Banned-IP patterns

Exact (`192.168.1.100`), CIDR-style wildcard (`192.168.1.*`, `192.168.*.*`), partial match (`10*`).

### Default auto-ban paths

Seeded on first install:

- `/wp-admin*`
- `/wp-login*`
- `*/.env`
- `*/xmlrpc.php`
- `/wp-includes*`
- `*/phpmyadmin*`
- `*/wp-content/plugins*`

### Honeypot flow

1. Bot fills hidden `_website` field.
2. Forms handler detects, logs WARNING event.
3. Fires `security.honeypot_triggered` hook with `{ip, form_slug, form_id, request_url}`.
4. Sentinel hook handler checks settings + whitelist.
5. Inserts ban into `sentinel_banned_ips`.
6. Bot gets a fake-success response.
7. All subsequent requests from the IP → 403.

Ban record stores the form URL + `honeypot:<form_slug>` as the matched pattern.

### GeoIP integration

Resolves country codes for banned IPs via MaxMind GeoLite2-Country. Configure via `OCMS_GEOIP_DB_PATH`. Without it, country columns stay empty.

### Database tables

- `sentinel_banned_ips` — pattern, country, notes, URL, timestamp.
- `sentinel_autoban_paths` — trigger patterns.
- `sentinel_whitelist` — whitelist patterns.
- `sentinel_settings` — key-value config.

### Admin routes

9 routes under `/admin/sentinel/*` covering CRUD for bans, paths, whitelist, and settings, plus a `POST /admin/sentinel/ban` AJAX endpoint for banning from the events page.

### Template functions

- `sentinelVersion` — string.
- `sentinelIsActive` — bool.
- `sentinelIsIPBanned(ip)` — bool.
- `sentinelIsIPWhitelisted(ip)` — bool.

### Hook integration

Listens on `security.honeypot_triggered` (fired by the forms handler). Hook payload is `map[string]any` with `ip`, `form_slug`, `form_id`, `request_url`.

### Hardening (0.5.0 → 0.18.1)

- 0.5.0 — initial release.
- 0.5.0 — self-ban prevention for administrators.
- 0.7.0 — session-cookie pre-check to avoid panic-recovery control flow.
- 0.17.0 — honeypot auto-ban and admin/editor exemption; auto-ban notes/URL truncation to 255 chars to prevent unbounded DB storage.
- 0.18.1 — validate Sentinel ban URLs before rendering (XSS hardening).

## Sources

- Origin: `raw/ocms-go.core/docs/sentinel-module.md`
