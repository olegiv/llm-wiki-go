# SEO

## Summary

oCMS emits search-engine and social-sharing metadata directly from the theme engine and surfaces per-page controls for overrides. Sitemap, robots.txt, hreflang, Open Graph, and Twitter Card all come for free.

## Per-page controls

- **Title override** (falls back to page title if unset).
- **Description** (SERP snippet).
- **Keywords** (comma-separated).
- **Canonical URL** (duplicate-content control).
- **NoIndex** / **NoFollow** (per-page crawler directives).

## Open Graph and Twitter Cards

Emitted automatically from page metadata:

- `og:title`, `og:description`.
- `og:image` plus `og:image:width`, `og:image:height`, `og:image:type`.
- Twitter Card tags.
- Since 0.10.2: article-specific OG tags (`article:published_time`, `article:modified_time`, `article:author`, `article:section`, `article:tag`); homepage canonical link.

### Which image variant provides `og:image`?

See [entities/media.md](../entities/media.md) contradiction. Short version: `docs/media.md` uses a dedicated `og` variant (1200×630, added in 0.18.0) with fallback priority `og > large > medium > thumbnail`. `wiki/Media-Library.md` and `wiki/SEO.md` still say `large`. The post-0.18.0 behavior is `og`-first.

## Sitemap

- Auto-generated XML at `/sitemap.xml`.
- Includes all published pages; updated on content change.
- Cached; see [topics/caching.md](caching.md).
- Since 0.18.1: cache keyed by configured `site_url` (not request `Host`); generation requires `site_url` in site config.
- Since 0.10.2: fix for double slashes in sitemap URLs when site URL has a trailing slash.

## Robots.txt

- Served at `/robots.txt`.
- References the sitemap URL.
- Configurable via site settings.

## Hreflang

Auto-emitted for each language translation of the current page, plus `x-default` to the default-language variant. See [topics/multi-language.md](multi-language.md) for the full detection and linking model.

## HTTP security headers (SEO-relevant)

| Header | Purpose |
|--------|---------|
| `Strict-Transport-Security` | HSTS for HTTPS-only policy. |
| `X-Frame-Options: SAMEORIGIN` | Prevents embedding in foreign frames. |
| `X-Content-Type-Options: nosniff` | MIME type enforcement. |
| `Referrer-Policy: strict-origin-when-cross-origin` | Referrer minimization. |
| `Content-Security-Policy` | Per-request nonce-based CSP; includes `frame-ancestors` directive. |

These headers are sent unconditionally and are relevant for SEO by signalling trust and avoiding mixed-content penalties.

## Open-redirect fixes (historical)

- 0.5.0: fixed CWE-601 open-redirect in the legacy `/blog/tag/{slug}` redirect.
- 0.18.1: prevented scheme-relative open redirects in the trailing-slash middleware.

Neither is a current SEO control, but both matter for signal integrity on link-based migrations.

## Sources

- [sources/wiki-seo.md](../sources/wiki-seo.md)
- [sources/readme.md](../sources/readme.md)
- [sources/changelog.md](../sources/changelog.md)
- [sources/security.md](../sources/security.md)
