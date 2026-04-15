# CSRF protection

## Summary

oCMS validates state-changing requests (POST / PUT / DELETE / PATCH) via **Fetch metadata headers** (`Sec-Fetch-Site`) rather than cookie-based CSRF tokens. The implementation is `filippo.io/csrf/gorilla`. Legacy browsers fall back to Origin / Referer validation.

The design is cookie-independent: it does not rely on `SameSite`, `Secure`, or `Domain` cookie attributes. Browsers supply Fetch metadata automatically, so AJAX and HTMX calls from the same origin need no token.

## Trusted origins

`CSRFConfig.TrustedOrigins` takes a list of **`host:port`** strings — not full URLs. Full URLs (`http://localhost:8080`) produce `origin invalid` errors.

### Defaults

| Environment | Defaults |
|-------------|----------|
| `OCMS_ENV=development` | `localhost:8080`, `127.0.0.1:8080` |
| `OCMS_ENV=production` | none — same-origin only |

Production deployments behind a distinct public hostname must add that hostname to `TrustedOrigins`.

## Auth key and session secret

`CSRFConfig.AuthKey` is the 32-byte encryption key. It is expected to be the same as `OCMS_SESSION_SECRET`. Minimum key length: 32 bytes.

## Legacy template variables

`{{.CSRFField}}` and `{{.CSRFToken}}` are retained as empty-string no-ops for backward compatibility. Templates that still include `{{.CSRFField}}` continue to render without errors but produce nothing.

## API routes

REST API endpoints use bearer-token API-key authentication instead of CSRF. They are excluded from CSRF validation via `middleware.SkipCSRF("/api/v1/...")`.

## Troubleshooting

| Error reason | Cause | Fix |
|--------------|-------|-----|
| `Sec-Fetch-Site cross-site` | Request from a different site. | Ensure request origin matches; add to `TrustedOrigins` if legitimate. |
| `origin invalid` | `Origin` header doesn't match trusted origins. | Use `host:port` (not full URLs) in `TrustedOrigins`. |
| `referer not supplied` | No Referer header on legacy browser. | Use a modern browser; Fetch metadata will cover it. |

## curl vs browser testing

curl does not send `Sec-Fetch-Site`. The library falls back to `Origin` and `Referer` for it:

```bash
curl -X POST http://localhost:8080/login \
  -H "Origin: http://localhost:8080" \
  -H "Referer: http://localhost:8080/login" \
  -d "email=admin@example.com&password=changeme"
```

## Related

- [Login security](login-security.md) — rate limiting and account lockout.
- [hCaptcha](hcaptcha.md) — optional bot protection on login.
- [Security overview](security-overview.md) — the broader security posture.

## Sources

- [sources/docs-csrf.md](../sources/docs-csrf.md)
- [sources/wiki-csrf-protection.md](../sources/wiki-csrf-protection.md)
- [sources/security.md](../sources/security.md)
