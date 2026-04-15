# Page

## Summary

`Page` is the primary content entity in oCMS. A page has a title, slug, body (rich HTML via TinyMCE), optional featured image, optional summary, created/updated/published timestamps, categories (many), tags (many), per-language translation siblings, and a page type (used to exclude utility pages from lists). Pages support drafts, scheduled publishing, and automatic versioning with restore.

## Key fields and behavior

- **Slug.** SEO-friendly, auto-generated from title on create; user-editable.
- **Body.** Edited in TinyMCE 8 rich editor. Write paths may enforce HTML sanitization (`OCMS_SANITIZE_PAGE_HTML`) and reject suspicious markup (`OCMS_BLOCK_SUSPICIOUS_PAGE_HTML`).
- **Versions.** Every save creates a version; authorized users can view and restore any prior version. Version preview bodies are properly escaped in the admin modal (hardened in 0.18.1).
- **Scheduling.** `scheduled_at` triggers automatic publish via the [scheduler](../topics/architecture-layers.md). Cleared on manual publish/unpublish (fix shipped in 0.18.1).
- **Draft preview.** Admin and editor users can preview drafts; rendered with `noindex, nofollow` headers (since 0.11.0).
- **Featured image.** Selected from [media](media.md); minimum 1200×800 on assignment. Also used as the default Open Graph image source.
- **OG image variant.** A dedicated 1200×630 `og` variant is auto-generated (since 0.18.0), preferred over `large` for social-sharing.
- **Video embedding.** `video_url` + optional `video_title` fields; responsive iframe render (since 0.16.0). See [topics/video-embedding.md](../topics/video-embedding.md).
- **Summary field.** Optional textarea (max 500 chars) that replaces auto-generated excerpt on frontend listings (since 0.11.0).
- **Translations.** Pages link bidirectionally across languages. See [topics/content-management.md](../topics/content-management.md).
- **Cache.** Context-aware page caching (since 0.3.0) automatically bypassed for admin requests; invalidated on publish toggle (fix shipped in 0.18.1).
- **Webhook events.** `page.created`, `page.updated`, `page.deleted`, `page.published` trigger deliveries.

## REST API

- `GET /api/v1/pages` — list published pages.
- `GET /api/v1/pages/{id}` — get by ID.
- `GET /api/v1/pages/slug/{slug}` — get by slug.
- `POST /api/v1/pages` — create (requires `pages:write`).
- `PUT /api/v1/pages/{id}` — update.
- `DELETE /api/v1/pages/{id}` — delete.

Draft access requires `pages:read` permission (enforced since 0.18.1). Author email is gated on API key authentication to prevent enumeration via `?include=author` (since 0.17.0).

## Sources

- [sources/readme.md](../sources/readme.md)
- [sources/wiki-content-management.md](../sources/wiki-content-management.md)
- [sources/wiki-home.md](../sources/wiki-home.md)
- [sources/docs-media.md](../sources/docs-media.md)
- [sources/docs-video-embedding.md](../sources/docs-video-embedding.md)
- [sources/changelog.md](../sources/changelog.md)
