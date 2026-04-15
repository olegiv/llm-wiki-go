# Security overview

## Summary

oCMS layers defense-in-depth across authentication, transport, request handling, and data paths. Detail documentation lives in the project's `docs/` tree, referenced from `SECURITY.md` and mirrored on the GitHub wiki as `wiki/Security.md`.

### Authentication

- Password hashing: **Argon2id** (`internal/auth`).
- Sessions: SCS library with SQLite persistence; 24-hour session lifetime.
- API keys: bearer token with **Argon2id hashing** (since 0.1.0).
- Role-based access control: Admin, Editor, Public.

### Login protection

Three-layer defense:

1. [hCaptcha](hcaptcha.md) — bot filter before credential check.
2. IP rate limiting — 0.5 req/s with burst 5 per source IP.
3. Account lockout — 5 failures in 15 min → 15 m lockout, doubling to 24 h max.

See [login security](login-security.md).

### API security (since 0.9.0)

- Per-key and global source CIDR allowlists (`OCMS_API_ALLOWED_CIDRS`, per-key CIDRs).
- Expiry policy and maximum key lifetime (`OCMS_API_KEY_MAX_TTL_DAYS`, default 90 days in production).
- Source-IP anomaly detection with auto-revocation (`OCMS_REVOKE_API_KEY_ON_SOURCE_IP_CHANGE`).
- Per-key token-bucket rate limiting.
- Fail-closed on malformed `X-Forwarded-For` chains.

### Request security

- **CSRF** via Fetch Metadata headers with Origin/Referer fallback. See [topic/csrf.md](csrf.md).
- **SQL injection:** 100% sqlc-generated parameterized queries.
- **XSS:** Go `html/template` auto-escaping + bluemonday sanitization. Write-path gates `OCMS_SANITIZE_PAGE_HTML` (sanitize on render) and `OCMS_BLOCK_SUSPICIOUS_PAGE_HTML` (reject suspicious patterns on write).
- **CSP:** per-request nonces across the render pipeline; `frame-ancestors` directive to prevent clickjacking.
- **Open redirect:** validated redirect targets (CWE-601 fix in 0.5.0; scheme-relative fix in 0.18.1).

### HTTP security headers

CSP, HSTS, `X-Frame-Options: SAMEORIGIN`, `X-Content-Type-Options: nosniff`, `Referrer-Policy: strict-origin-when-cross-origin`.

### Webhooks

- **HMAC-SHA256** signed payloads.
- Destination host allowlisting (`OCMS_WEBHOOK_ALLOWED_HOSTS`).
- HTTPS enforcement on outbound URLs (`OCMS_REQUIRE_HTTPS_OUTBOUND`).
- Form-data minimization modes: `redacted` (default) / `none` / `full` via `OCMS_WEBHOOK_FORM_DATA_MODE`.

### Embed proxy (Dify integration, since 0.9.0)

- Browser-origin and upstream-host allowlists (`OCMS_EMBED_ALLOWED_ORIGINS`, `OCMS_EMBED_ALLOWED_UPSTREAM_HOSTS`).
- Short-lived signed proxy tokens (`OCMS_EMBED_PROXY_TOKEN`), minted at render time since 0.18.1.

### Trusted-proxy hardening (since 0.17.0)

- Replaced chi's `RealIP` middleware with a trusted-proxy-aware implementation. chi's `RealIP` blindly trusts `X-Real-IP`, `X-Forwarded-For`, and `True-Client-IP` from any peer, which allowed external attackers to spoof `127.0.0.1` and evade IP-based banning.
- Fail-closed on malformed `X-Forwarded-For` chains.
- `OCMS_REQUIRE_TRUSTED_PROXIES` forces startup failure in production when unset.

### File uploads

- 20 MB size limit (2 MB in demo mode).
- MIME validation and magic-number checking.
- UUID-based storage paths.

### Production startup gates (since 0.9.0)

- Block startup on default admin credentials.
- Require embed-origin / trusted-proxy / API-CIDR / captcha configuration in production.
- Security signal summary exposed via the `/health` endpoint (authenticated view).

### Production checklist (from `wiki/Security.md`)

- [ ] Set strong `OCMS_SESSION_SECRET` (32+ bytes).
- [ ] Set `OCMS_ENV=production`.
- [ ] Configure `OCMS_TRUSTED_PROXIES`.
- [ ] Enable SSL/TLS with a valid certificate.
- [ ] Configure `OCMS_EMBED_PROXY_TOKEN` (if embed proxy is used).
- [ ] Set `OCMS_WEBHOOK_ALLOWED_HOSTS` (if webhooks are used).
- [ ] Set `OCMS_API_ALLOWED_CIDRS` (if API is used).
- [ ] Enable hCaptcha on login form.
- [ ] Change default admin password.
- [ ] Configure firewall (only 80/443 open).

### Vulnerability reporting

Preferred: GitHub Security Advisories. Acknowledgment within 48–72 hours; initial assessment within 1 week; fix timeline driven by severity.

### Dependency scanning

- Go: `govulncheck ./...`
- JavaScript: `npm audit`

## Sources

- [sources/security.md](../sources/security.md)
- [sources/wiki-security.md](../sources/wiki-security.md)
- [sources/changelog.md](../sources/changelog.md)
- [sources/claude.md](../sources/claude.md)
- [sources/readme.md](../sources/readme.md)
