# wiki/Configuration.md (oCMS)

## Summary

Environment variable reference on the GitHub wiki. Organizes variables into groups: Required, Core Settings, Caching, Security, API Security, Embed Proxy, Webhooks, GeoIP, and a Production Configuration Example (`.env` snippet). Covers roughly the same territory as `README.md`'s variable table but with around 40 variables (not the full 55+).

Documented here but missing from `README.md`:

- `OCMS_API_RATE_LIMIT` — API requests per minute per key, default `100`. Also present in `CLAUDE.md`.

Documented in `README.md` but missing from this wiki page:

- `OCMS_HCAPTCHA_DISABLED` — force-disable hCaptcha regardless of database settings.
- `OCMS_DEMO_MODE` is present here but several production gate toggles (e.g. `OCMS_REQUIRE_EMBED_PROXY_TOKEN`) are absent.

The production `.env` example is a near-duplicate of the one in `README.md`.

## Sources

- Origin: `raw/ocms-go.core/wiki/Configuration.md`
