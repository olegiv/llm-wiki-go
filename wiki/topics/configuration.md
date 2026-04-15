# Configuration

## Summary

oCMS is configured exclusively via `OCMS_*` environment variables. The only strictly required variable is `OCMS_SESSION_SECRET` (minimum 32 bytes). Production environments have additional gates that can force startup failure when unset (the `OCMS_REQUIRE_*` family).

## Required

| Variable | Purpose |
|----------|---------|
| `OCMS_SESSION_SECRET` | Session encryption key. Minimum 32 bytes. |

## Core runtime

| Variable | Default | Purpose |
|----------|---------|---------|
| `OCMS_DB_PATH` | `./data/ocms.db` | SQLite database file. |
| `OCMS_SERVER_HOST` | `localhost` | Listen host. |
| `OCMS_SERVER_PORT` | `8080` | Listen port. |
| `OCMS_ENV` | `development` | `development` or `production`; `production` enables strict gates. |
| `OCMS_LOG_LEVEL` | `info` | `debug` / `info` / `warn` / `error`. |
| `OCMS_ERROR_LOG_PATH` | — | Separate error log file (errors still go to stdout). |
| `OCMS_UPLOADS_DIR` | `./uploads` | Media upload root. |
| `OCMS_CUSTOM_DIR` | `./custom` | Custom themes and modules root. |
| `OCMS_ACTIVE_THEME` | `default` | Active theme; overrides DB/admin setting when set. |

## Caching

| Variable | Default | Purpose |
|----------|---------|---------|
| `OCMS_CACHE_TTL` | `3600` | Default cache TTL (seconds). |
| `OCMS_CACHE_MAX_SIZE` | `10000` | Max entries for the in-memory cache. |
| `OCMS_REDIS_URL` | — | Redis URL for distributed cache. |
| `OCMS_CACHE_PREFIX` | `ocms:` | Redis key prefix. |

## Security and API

| Variable | Default | Purpose |
|----------|---------|---------|
| `OCMS_TRUSTED_PROXIES` | — | Trusted reverse-proxy CIDRs/IPs. |
| `OCMS_REQUIRE_TRUSTED_PROXIES` | prod: `true` | Fail startup in production if unset. |
| `OCMS_API_ALLOWED_CIDRS` | — | Global source CIDRs/IPs allowed to use API keys. |
| `OCMS_REQUIRE_API_ALLOWED_CIDRS` | prod: `true` | Fail API key auth when unset. |
| `OCMS_API_KEY_MAX_TTL_DAYS` | prod: `90` | Maximum API key lifetime (days, `0` disables, max 365). |
| `OCMS_REQUIRE_API_KEY_EXPIRY` | prod: `true` | Require API keys to have expiration timestamps. |
| `OCMS_REQUIRE_API_KEY_SOURCE_CIDRS` | prod: `true` | Require per-key source CIDR restrictions. |
| `OCMS_REVOKE_API_KEY_ON_SOURCE_IP_CHANGE` | prod: `true` | Auto-revoke keys on source IP anomaly when no per-key CIDRs. |
| `OCMS_SANITIZE_PAGE_HTML` | prod: `true` | Sanitize page HTML before rendering. |
| `OCMS_REQUIRE_SANITIZE_PAGE_HTML` | prod: `true` | Fail startup if sanitization disabled. |
| `OCMS_BLOCK_SUSPICIOUS_PAGE_HTML` | prod: `true` | Reject page writes containing suspicious HTML patterns. |
| `OCMS_REQUIRE_BLOCK_SUSPICIOUS_PAGE_HTML` | prod: `true` | Fail startup if blocking disabled or existing pages contain suspicious markers. |
| `OCMS_REQUIRE_FORM_CAPTCHA` | prod: `true` | Require captcha on public form submissions. |
| `OCMS_REQUIRE_HTTPS_OUTBOUND` | prod: `true` | Require HTTPS on outbound integration URLs. |
| `OCMS_API_RATE_LIMIT` | `100` | API requests per minute per key. |

## hCaptcha and GeoIP

| Variable | Purpose |
|----------|---------|
| `OCMS_HCAPTCHA_SITE_KEY` / `OCMS_HCAPTCHA_SECRET_KEY` | hCaptcha credentials for login protection. |
| `OCMS_HCAPTCHA_DISABLED` | Force-disable hCaptcha regardless of database settings. |
| `OCMS_GEOIP_DB_PATH` | Path to `GeoLite2-Country.mmdb` for country detection. |

## Embed proxy (Dify)

| Variable | Purpose |
|----------|---------|
| `OCMS_EMBED_ALLOWED_ORIGINS` | Browser origins permitted to use the embed proxy (exact scheme + host match). |
| `OCMS_EMBED_ALLOWED_UPSTREAM_HOSTS` | Upstream API hosts the proxy may reach. |
| `OCMS_EMBED_PROXY_TOKEN` | Static secret for minting short-lived signed proxy tokens. |
| `OCMS_REQUIRE_EMBED_ALLOWED_ORIGINS` / `_UPSTREAM_HOSTS` / `_TOKEN` | Production startup gates. |

## Webhooks

| Variable | Default | Purpose |
|----------|---------|---------|
| `OCMS_WEBHOOK_ALLOWED_HOSTS` | — | Destination allowlist for active webhooks (exact hostname match). |
| `OCMS_WEBHOOK_FORM_DATA_MODE` | `redacted` | `form.submitted` payload data mode: `redacted` / `none` / `full`. |
| `OCMS_REQUIRE_WEBHOOK_ALLOWED_HOSTS` | prod: `true` | Fail startup with active webhooks and no destination allowlist. |
| `OCMS_REQUIRE_WEBHOOK_FORM_DATA_MINIMIZATION` | prod: `true` | Fail startup when form webhook mode is `full`. |

## Seeding and demo

| Variable | Default | Purpose |
|----------|---------|---------|
| `OCMS_DO_SEED` | `false` | Seed default admin (`admin@example.com` / `changeme1234`) and base config on first run. |
| `OCMS_DEMO_MODE` | `false` | Seed demo users, pages, media; activates demo-mode read-only restrictions. Requires `OCMS_DO_SEED=true`. |

## Contradictions

### Variable catalog completeness

Three sources catalog `OCMS_*` variables with different scope:

- [sources/readme.md](../sources/readme.md) — the full reference, ~55+ variables including every `OCMS_REQUIRE_*` production gate, embed proxy, and webhook policies. Omits `OCMS_API_RATE_LIMIT`.
- [sources/wiki-configuration.md](../sources/wiki-configuration.md) — ~40 variables grouped by concern. Includes `OCMS_API_RATE_LIMIT` (default `100`). Omits `OCMS_HCAPTCHA_DISABLED` and `OCMS_REQUIRE_EMBED_PROXY_TOKEN`.
- [sources/claude.md](../sources/claude.md) — working subset of ~25 variables. Includes `OCMS_API_RATE_LIMIT` (default `100`). Omits most `OCMS_REQUIRE_*` gates.

`OCMS_API_RATE_LIMIT` appears in two of three sources but not in the README — that is drift rather than a semantic conflict, and this wiki page documents it under *Security and API* above.

`OCMS_HCAPTCHA_DISABLED` appears in README but is missing from the wiki `Configuration.md`. Likely drift in the opposite direction.

## Sources

- [sources/readme.md](../sources/readme.md)
- [sources/wiki-configuration.md](../sources/wiki-configuration.md)
- [sources/claude.md](../sources/claude.md)
- [sources/security.md](../sources/security.md)
