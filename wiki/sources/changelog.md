# CHANGELOG (oCMS)

## Summary

Keep-a-Changelog-formatted history of oCMS releases from **0.0.0** (2026-01-31) through **0.18.1** (2026-04-14). Contains Added / Changed / Fixed / Security sections per version plus `[Unreleased]` header.

### Feature arcs by version

- **0.0.0** — Initial release: pages, media, forms, themes, modules, REST API, security, SEO, i18n, webhooks, import/export, caching.
- **0.1.0** — Argon2id API key hashing, GeoIP analytics, hCaptcha, Ubuntu/Plesk single-instance deploy.
- **0.2.0** — Embedded core themes architecture, module active-status toggle, admin UI i18n (English + Russian).
- **0.3.0** — DB Manager module, context-aware page caching, page types, `SECURITY.md` introduced.
- **0.4.0** — Docker multi-stage build, TinyMCE 8 replaces TipTap editor.
- **0.5.0** — Privacy (Klaro) and Sentinel (IP-ban) modules, global URL redirects, open-redirect CWE-601 fix.
- **0.6.0** — Fly.io deployment configuration, event-log ban-IP button.
- **0.7.0** — Informer notification-bar module, comprehensive demo-mode read-only protections.
- **0.8.0** — Custom bookmarks module example + Starter theme, scheduler admin UI.
- **0.9.0** — Major security push: embed proxy (Dify), API-key CIDR/expiry/anomaly policies, page HTML sanitization, form submission hardening, production startup gates, CSP nonces.
- **0.10.0–0.10.2** — templUI component migration, deploy-docs consolidation, brace-expansion CVE-2026-33750 fix.
- **0.11.0** — Page summary field, admin draft preview with noindex protection.
- **0.12.0** — Author profile fields (avatar, bio, website, LinkedIn, GitHub).
- **0.14.0** — Analytics IP/CIDR exclusions, code block syntax highlighting (Prism.js), frame-ancestors CSP.
- **0.15.0** — Dify knowledge-base exports, wiki submodule, ~15.6k new test lines.
- **0.16.0** — Video embedding (YouTube/Vimeo/Dailymotion); **breaking change** to `sites.conf` format.
- **0.17.0** — Structured API event logging, Sentinel honeypot auto-ban, trusted-proxy-aware IP resolution replacing chi `RealIP`.
- **0.18.0** — Medium-style read/engagement analytics with retention reports.
- **0.18.1** — Dependency bumps, embed-proxy token fixes, extensive XSS and data-protection hardenings.

Note: version 0.13.0 does not appear in the changelog.

## Sources

- Origin: `raw/ocms-go.core/top-level/CHANGELOG.md`
