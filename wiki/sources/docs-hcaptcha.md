# docs/hcaptcha.md (oCMS)

## Summary

Reference for the hCaptcha module — bot protection on the login form.

### Account setup

hCaptcha dashboard account → obtain Site Key + Secret Key.

### Configuration paths

1. **Admin UI** at `/admin/hcaptcha` (Module settings: Enabled, Site Key, Secret Key, Theme light/dark, Size normal/compact).
2. **Environment variables:** `OCMS_HCAPTCHA_SITE_KEY`, `OCMS_HCAPTCHA_SECRET_KEY`. Env vars take precedence over DB.
3. **Force-disable env:** `OCMS_HCAPTCHA_DISABLED=true` overrides database regardless of admin UI state. Emits a warning notice on the settings page. Useful for CI, staging, emergency lockout recovery.

### Module info

- Module name: `hcaptcha`.
- Version: 1.0.0.
- Admin URL: `/admin/hcaptcha`.
- DB table: `hcaptcha_settings`.

### Test keys (claim: used as defaults)

**Docs asserts** that hCaptcha test keys are used as defaults when no keys are configured:

- Site Key: `10000000-ffff-ffff-ffff-000000000001`
- Secret Key: `0x0000000000000000000000000000000000000000`

"With test keys, the widget always passes automatically — no actual challenge is shown. Verification always succeeds." "This means you can enable hCaptcha immediately in development without setting up a real hCaptcha account."

**This claim is contradicted by `CHANGELOG.md` 0.18.1**: "Remove insecure default hCaptcha test keys" and "Scrub persisted hCaptcha test keys on upgrade". The docs page has not been updated.

### Verification flow

User completes widget → server POST to `https://api.hcaptcha.com/siteverify` with Secret Key → success or failure.

### Disable paths

- Admin UI toggle off.
- `UPDATE hcaptcha_settings SET enabled = 0` (keeps module loaded).
- `UPDATE modules SET is_active = 0 WHERE name = 'hcaptcha'` (unload module).
- `OCMS_HCAPTCHA_DISABLED=true` env var (runtime override).

### Reverse proxy

Middleware header order: `X-Forwarded-For`, `X-Real-IP`, `RemoteAddr`.

### Integration with login security

Three-layer login protection: hCaptcha → IP rate limit → account lockout.

### Error codes

`missing-input-secret`, `invalid-input-secret`, `missing-input-response`, `invalid-input-response`.

## Sources

- Origin: `raw/ocms-go.core/docs/hcaptcha.md`
