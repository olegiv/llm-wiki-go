# SECURITY.md (oCMS)

## Summary

Security policy and feature inventory for oCMS.

### Policy

- **Supported versions:** 0.14.x and 0.12.x. Versions below 0.12 are unsupported.
- **Reporting:** GitHub Security Advisories (preferred) or direct email. Acknowledgment within 48–72 hours; initial assessment within 1 week; fix timeline depends on severity.

### Security feature inventory

- **Authentication & authorization.** Argon2id password hashing; SCS session management with SQLite persistence and 24-hour lifetime; role-based access control (Admin, Editor, Public).
- **Login protection.** IP rate limiting; exponential-backoff account lockout; optional hCaptcha.
- **API security.** Bearer tokens with Argon2id hashing; granular per-key permissions; per-key token-bucket rate limiting; per-key source CIDR allowlists; configurable maximum key lifetime (default 90 days in production); source-IP anomaly detection with auto-revocation; global API CIDR policy (`OCMS_API_ALLOWED_CIDRS`); fail-closed on malformed forwarding headers.
- **Request security.** CSRF via Fetch Metadata headers with Origin/Referer fallback; 100% sqlc-parameterized SQL; Go `html/template` auto-escaping + bluemonday sanitization; per-request CSP nonces; `frame-ancestors` directive; open-redirect validation; configurable page HTML sanitization (`OCMS_SANITIZE_PAGE_HTML`) and suspicious-markup blocking (`OCMS_BLOCK_SUSPICIOUS_PAGE_HTML`).
- **HTTP security headers.** CSP, HSTS, `X-Frame-Options: SAMEORIGIN`, `X-Content-Type-Options: nosniff`, `Referrer-Policy: strict-origin-when-cross-origin`.
- **Webhook security.** HMAC-SHA256 signatures; destination host allowlisting (`OCMS_WEBHOOK_ALLOWED_HOSTS`); HTTPS outbound enforcement (`OCMS_REQUIRE_HTTPS_OUTBOUND`); form-data minimization modes (`OCMS_WEBHOOK_FORM_DATA_MODE`).
- **Embed proxy.** Browser-origin allowlisting (`OCMS_EMBED_ALLOWED_ORIGINS`); upstream-host allowlisting (`OCMS_EMBED_ALLOWED_UPSTREAM_HOSTS`); short-lived signed tokens (`OCMS_EMBED_PROXY_TOKEN`).
- **Trusted-proxy hardening.** Custom trusted-proxy-aware IP resolution replacing chi's `RealIP` (which blindly trusts `True-Client-IP`, `X-Real-IP`, `X-Forwarded-For` from any peer); configurable startup failure when unset.
- **Content security.** Page HTML sanitization on render; suspicious-markup blocking on writes; optional captcha requirement on public forms.
- **File upload security.** Size limits (20 MB, 2 MB in demo mode); MIME validation; magic-number checks; UUID-based storage.

### Detail documentation

- Login protection → `docs/login-security.md`.
- CSRF config → `docs/csrf.md`.
- hCaptcha → `docs/hcaptcha.md`.

### Dependency scanning

- Go: `govulncheck ./...`
- JavaScript: `npm audit`

## Sources

- Origin: `raw/ocms-go.core/top-level/SECURITY.md`
