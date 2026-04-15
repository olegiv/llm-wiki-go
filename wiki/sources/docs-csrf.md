# docs/csrf.md (oCMS)

## Summary

Reference for CSRF protection. Implementation: `filippo.io/csrf/gorilla` library, using **Fetch metadata headers** (`Sec-Fetch-Site`) rather than cookie-based tokens. Applies to POST/PUT/DELETE/PATCH. Legacy fallback: Origin/Referer header checks for browsers without Fetch metadata.

`CSRFConfig` struct (`internal/middleware/csrf.go`) has `AuthKey` (32-byte encryption key), `ErrorHandler`, and `TrustedOrigins`. Cookie-related options (Secure, Domain, Path, MaxAge, SameSite) are no longer used.

**`TrustedOrigins` must be `host:port` format** — NOT full URLs. Full URLs cause "origin invalid" errors.

Development (`OCMS_ENV=development`) trusts `localhost:8080` and `127.0.0.1:8080`. Production has no default trusted origins — same-origin only.

Template variables `{{.CSRFField}}` and `{{.CSRFToken}}` are retained for backward compatibility but now output empty strings; they are no longer required.

AJAX requests from the same origin need no CSRF token — Fetch metadata handles it. REST API routes skip CSRF entirely (`middleware.SkipCSRF("/api/v1/...")`); API uses token auth instead.

Troubleshooting table covers `Sec-Fetch-Site cross-site`, `origin invalid`, `referer not supplied` error reasons. Common mistakes: full URLs in `TrustedOrigins`, missing origins in list, HTTPS mismatches.

The CSRF auth key is expected to be the same as `OCMS_SESSION_SECRET`; 32-byte minimum. Production requires HTTPS.

## Sources

- Origin: `raw/ocms-go.core/docs/csrf.md`
