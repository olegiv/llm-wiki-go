# wiki/SEO.md (oCMS)

## Summary

Wiki page covering oCMS's built-in SEO tools.

### Per-page meta fields

- **Title** (falls back to page title).
- **Description** (meta description for SERP snippets).
- **Keywords** (comma-separated).
- **Canonical URL** (override for duplicate-content control).
- **NoIndex** / **NoFollow** (per-page crawler directives).

### Open Graph + Twitter Cards

- `og:title`, `og:description`.
- **`og:image` — uses the `large` variant** from the media library (per this page).
- Twitter Card meta tags emitted automatically.

**Claim about the OG variant contradicts `docs/media.md`**, which documents a dedicated `og` variant preferred for social sharing with priority `og > large > medium > thumbnail`. The `wiki/SEO.md` page inherits the stale "large" choice from `wiki/Media-Library.md`.

### Sitemap

- Auto-generated XML sitemap at `/sitemap.xml`.
- Includes all published pages.
- Updated automatically on content change.
- Cached for performance.
- Since 0.18.1: cache keyed by configured `site_url` (not request `Host`) and generation requires `site_url` in site config.

### Robots.txt

- Served at `/robots.txt`.
- References sitemap URL.
- Configurable via site settings.

### Hreflang

Auto-emitted for multi-language sites (see [[Multi-Language]]):

```html
<link rel="alternate" hreflang="en" href="https://example.com/about-us">
<link rel="alternate" hreflang="ru" href="https://example.com/ru/about-us">
<link rel="alternate" hreflang="x-default" href="https://example.com/about-us">
```

### Security headers (repeated here)

CSP, HSTS, `X-Frame-Options: SAMEORIGIN`, `X-Content-Type-Options: nosniff`, `Referrer-Policy: strict-origin-when-cross-origin`.

## Sources

- Origin: `raw/ocms-go.core/wiki/SEO.md`
