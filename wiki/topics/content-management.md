# Content management

## Summary

The content pipeline revolves around the [Page](../entities/page.md) entity. Pages are authored in a TinyMCE 8 rich editor, tracked through automatic versioning, organized via [categories](../entities/category.md) and [tags](../entities/tag.md), surfaced through [menus](../entities/menu.md), and delivered to the public frontend through the active theme.

## Editing workflow

1. Admin navigates to `/admin/pages` and creates or opens a page.
2. Editor writes title, body, and optional summary (max 500 characters, since 0.11.0). Slug auto-generates from title but is editable.
3. Optional: assign categories, tags, featured image from the media library, video embed (YouTube/Vimeo/Dailymotion with optional title — since 0.16.0), and SEO overrides.
4. Save creates a new version. Publish flips the draft to published; schedule with `scheduled_at` to auto-publish later.
5. Admins and editors can preview drafts at any time (rendered `noindex, nofollow` since 0.11.0).

## Versioning

- Every save creates a Page version row.
- Authorized users can view and restore prior versions from `/admin/pages`.
- Version preview bodies are properly escaped in the admin modal (hardened in 0.18.1).

## Search

- SQLite FTS5 indexes title and body for published pages.
- Exposed at `/search` on the public frontend.
- Unlisted posts are filtered from tag/category pages and search results (since 0.18.1).

## Organization

- [Taxonomy](../entities/category.md): hierarchical categories (`/category/{slug}`) and flat tags (`/tag/{slug}`).
- [Menus](../entities/menu.md): named, ordered, nestable, per-language.
- **Page types:** content vs. utility. "Exclude from lists" flag removes utility pages from listings (since 0.3.0).

## Cache

Context-aware page caching (since 0.3.0):

- Automatic admin-request bypass.
- Invalidation on content update.
- Invalidation on publish toggle (fix in 0.18.1).
- Sitemap cache keyed by configured `site_url` (since 0.18.1 security fix).

## Translation

Pages link bidirectionally across languages. Translation link UI lives in the page editor; URL-prefixed routing (e.g. `/ru/about-us`) serves the translated variant. See the Multi-Language topic for the full workflow (to be ingested in a later batch).

## Import and export

See the Import-Export topic (to be ingested). Supports JSON and ZIP, conflict resolution (skip/overwrite/rename), and a dry-run preview mode.

## Webhooks

Page lifecycle emits `page.created`, `page.updated`, `page.deleted`, `page.published` events, which trigger HMAC-signed webhook deliveries to configured endpoints.

## Sources

- [sources/wiki-content-management.md](../sources/wiki-content-management.md)
- [sources/readme.md](../sources/readme.md)
- [sources/changelog.md](../sources/changelog.md)
