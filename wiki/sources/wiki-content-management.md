# wiki/Content-Management.md (oCMS)

## Summary

Overview page describing the oCMS content pipeline, oriented around pages:

- **Page management** — TinyMCE-based rich editor; automatic versioning with restore; scheduled publishing; draft/published states; featured-image selection from the media library; SEO-friendly slugs with auto-generation from titles; video embedding for YouTube/Vimeo/Dailymotion with responsive rendering; editable `created_at` and `published_at` timestamps.
- **Full-text search** — SQLite FTS5 indexes titles and body content of published pages.
- **Organization** — via categories (hierarchical) and tags (flat), plus menus for navigation.
- **Admin UX** — pagination, sorting, filtering (language, status, page type, text), and multi-select bulk actions.
- **Translation** — bidirectional translation links between language versions of a page (points to Multi-Language page for full workflow).
- **Import/Export** — JSON or ZIP, with conflict resolution.
- **Webhooks** — create/update/delete/publish events trigger webhook deliveries.

## Sources

- Origin: `raw/ocms-go.core/wiki/Content-Management.md`
