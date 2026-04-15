# hCaptcha

## Summary

hCaptcha is an optional bot-protection layer on the login form, implemented as an oCMS module (name: `hcaptcha`, version 1.0.0, DB table: `hcaptcha_settings`, admin URL: `/admin/hcaptcha`). It is privacy-focused and GDPR-compliant.

## Configuration

Three paths, checked in precedence order (env > admin UI > DB):

| Path | Scope |
|------|-------|
| `OCMS_HCAPTCHA_SITE_KEY` / `OCMS_HCAPTCHA_SECRET_KEY` | Runtime override of the DB-stored keys. |
| Admin UI at `/admin/hcaptcha` (**Modules → hCaptcha**) | Database-persisted keys, theme, size, enabled flag. |
| Direct SQL on `hcaptcha_settings` | Emergency changes; requires restart. |

### Settings

| Setting | Options | Description |
|---------|---------|-------------|
| Enabled | On / Off | Toggle the captcha at the feature level. |
| Site Key | String | Public key for the widget. |
| Secret Key | String | Private key for server-side verification. |
| Theme | `light` / `dark` | Widget color. |
| Size | `normal` / `compact` | Widget size. |

## Emergency kill switches

| Method | Effect |
|--------|--------|
| Admin UI toggle off | Disables captcha, module stays loaded. |
| `UPDATE hcaptcha_settings SET enabled = 0` | DB-level setting disable; requires restart. |
| `UPDATE modules SET is_active = 0 WHERE name = 'hcaptcha'` | Unload the module entirely. |
| `OCMS_HCAPTCHA_DISABLED=true` | Runtime override — forces disable regardless of DB state. Admin UI toggle is ignored with a warning notice displayed while set. Useful for CI, staging, and emergency lockout recovery. |

## Verification flow

1. Enabled → widget renders on login form.
2. User completes challenge → token posted with credentials.
3. Server `POST https://api.hcaptcha.com/siteverify` with Secret Key.
4. On success, login proceeds. On failure, re-render with error.

## Reverse proxy IP resolution

Module header order: `X-Forwarded-For` → `X-Real-IP` → `RemoteAddr`.

Note: this order differs from the login-security middleware's order (which is `X-Real-IP` → `X-Forwarded-For` → `RemoteAddr`). Combined with `OCMS_TRUSTED_PROXIES` and the trusted-proxy-aware IP resolution introduced in 0.17.0, the forwarding chain is accepted only from configured trusted peers.

## Error codes

From the hCaptcha API:

- `missing-input-secret`, `invalid-input-secret` — server misconfiguration.
- `missing-input-response`, `invalid-input-response` — client did not complete or token is stale.

## Integration with login protection

Three-layer defense order: **hCaptcha** (bot filter) → IP rate limit → account lockout. See [Login security](login-security.md).

## Contradictions

### Test keys as defaults

- [sources/docs-hcaptcha.md](../sources/docs-hcaptcha.md): "hCaptcha provides test keys for development and testing. These are used as defaults when no keys are configured" — lists site key `10000000-ffff-ffff-ffff-000000000001` and secret `0x0000000000000000000000000000000000000000` with the comment "With test keys, the widget always passes automatically — no challenge is shown."
- [sources/wiki-hcaptcha.md](../sources/wiki-hcaptcha.md): mirrors the same claim.
- [sources/changelog.md](../sources/changelog.md) 0.18.1 Security entry: "**Remove insecure default hCaptcha test keys**" and "Scrub persisted hCaptcha test keys on upgrade."

Both documentation files describe pre-0.18.1 behavior; neither has been updated since the security hardening. Post-0.18.1, unconfigured installs do not automatically pass captcha — this is the correct behavior.

### `OCMS_HCAPTCHA_DISABLED` coverage

- [sources/docs-hcaptcha.md](../sources/docs-hcaptcha.md) documents `OCMS_HCAPTCHA_DISABLED` with rationale and behavior.
- [sources/wiki-hcaptcha.md](../sources/wiki-hcaptcha.md), [sources/wiki-configuration.md](../sources/wiki-configuration.md) — variable absent.
- [sources/readme.md](../sources/readme.md) — documents the variable.

The wiki pages predate the 0.16.0 addition of this variable.

## Sources

- [sources/docs-hcaptcha.md](../sources/docs-hcaptcha.md)
- [sources/wiki-hcaptcha.md](../sources/wiki-hcaptcha.md)
- [sources/readme.md](../sources/readme.md)
- [sources/changelog.md](../sources/changelog.md)
- [sources/security.md](../sources/security.md)
