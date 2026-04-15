# Release history

## Summary

oCMS follows [Semantic Versioning](https://semver.org) and records every release in [Keep-a-Changelog](https://keepachangelog.com) format. Development progressed from an initial release on 2026-01-31 (0.0.0) through 0.18.1 on 2026-04-14 — roughly 19 releases in ~10 weeks. Version 0.13.0 is not present in the changelog.

## Version arc

| Version | Date | Highlights |
|---------|------|-----------|
| 0.0.0 | 2026-01-31 | Initial release: pages, media, forms, themes, modules, REST API, i18n, webhooks, import/export, caching. |
| 0.1.0 | 2026-02-01 | Argon2id API key hashing, hCaptcha, GeoIP analytics, Ubuntu/Plesk deploy. |
| 0.2.0 | 2026-02-01 | Embedded themes architecture, module active-status toggle, admin UI i18n (EN + RU). |
| 0.3.0 | 2026-02-03 | DB Manager module, context-aware page caching, page types, SECURITY.md introduced. |
| 0.4.0 | 2026-02-04 | Docker multi-stage build, TinyMCE 8 replaces TipTap editor. |
| 0.5.0 | 2026-02-05 | Privacy (Klaro) and Sentinel (IP-ban) modules; open-redirect CWE-601 fix. |
| 0.6.0 | 2026-02-06 | Fly.io deployment, event-log ban-IP button. |
| 0.7.0 | 2026-02-09 | Informer module; comprehensive demo-mode read-only protections. |
| 0.8.0 | 2026-02-12 | Custom bookmarks module + Starter theme; scheduler admin UI. |
| 0.9.0 | 2026-02-18 | **Major security push:** embed proxy, API-key CIDR/expiry/anomaly policies, page HTML sanitization, form hardening, CSP nonces, production startup gates. |
| 0.10.0 | 2026-03-22 | templUI component migration; bulk pager actions; per-page selector. |
| 0.10.1 | 2026-03-25 | Binary-only deploy script; server start/stop event logging. |
| 0.10.2 | 2026-03-27 | OG article meta tags, canonical link on homepage; brace-expansion CVE-2026-33750 fix. |
| 0.11.0 | 2026-03-31 | Page summary field, admin draft preview with noindex protection. |
| 0.12.0 | 2026-04-02 | Author profile fields (avatar, bio, website, LinkedIn, GitHub). |
| 0.14.0 | 2026-04-04 | Analytics IP/CIDR exclusions, Prism.js code syntax highlighting, frame-ancestors CSP. |
| 0.15.0 | 2026-04-04 | Dify knowledge-base exports, wiki submodule, ~15.6k new test lines. |
| 0.16.0 | 2026-04-06 | Video embedding (YouTube/Vimeo/Dailymotion); **breaking:** `sites.conf` format change. |
| 0.17.0 | 2026-04-08 | Structured API event logging; Sentinel honeypot auto-ban; trusted-proxy-aware IP resolution replacing chi `RealIP`. |
| 0.18.0 | 2026-04-09 | Medium-style read/engagement analytics with retention reports. |
| 0.18.1 | 2026-04-14 | Dependency updates; embed-proxy token fixes; extensive XSS and data-protection hardenings. |

## Supported versions

Per [sources/security.md](../sources/security.md), the officially supported lines are **0.14.x** and **0.12.x**. This table has not been updated to cover 0.15.x through 0.18.x — see the contradiction note on [entities/ocms.md](../entities/ocms.md).

## Notable breaking changes

- **0.16.0** — `sites.conf` column 2 is now the full instance directory; deploy scripts no longer append `/ocms` automatically. Existing deployments must migrate before updating scripts.

## Sources

- [sources/changelog.md](../sources/changelog.md)
- [sources/security.md](../sources/security.md)
