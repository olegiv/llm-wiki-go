# wiki/Security.md (oCMS)

## Summary

Wiki rewrite of `SECURITY.md` for GitHub-wiki navigation. Same supported-versions table (0.14.x and 0.12.x), same vulnerability-reporting guidance (GitHub Security Advisories preferred, 48–72 h acknowledgment). Same feature catalog: Argon2id + SCS sessions + 24 h lifetime; three-layer login protection (hCaptcha → IP rate limit → account lockout) with pointer to Login Security; request security (CSRF, sqlc, XSS, CSP nonces + `frame-ancestors`, open-redirect, HTML sanitization, suspicious-markup blocking); HTTP security headers; file upload security (20 MB, 2 MB demo, MIME + magic number + UUID storage); API security (per-key CIDR, expiry, rate limit, 90-day default TTL in prod, source-IP anomaly, global API CIDR policy, fail-closed forwarding); webhook security (HMAC-SHA256, destination allowlist, HTTPS outbound, form data minimization); embed proxy security; trusted-proxy hardening; content security (page sanitization, suspicious markup blocking, form captcha requirement).

Adds a **"Production Security Checklist"** not present in `SECURITY.md` — 10-item checklist covering strong session secret, production env, trusted proxies, SSL/TLS, embed token, webhook allowlist, API CIDRs, hCaptcha, default admin password, firewall.

Inherits the stale "0.14.x and 0.12.x" supported-versions claim from `SECURITY.md`.

## Sources

- Origin: `raw/ocms-go.core/wiki/Security.md`
