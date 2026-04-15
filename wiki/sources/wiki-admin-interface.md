# wiki/Admin-Interface.md (oCMS)

## Summary

Reference for the admin panel at `/admin/`.

- **Dashboard** — statistics, recent submissions, quick actions.
- **Editor routes** — `/admin/` (dashboard), `/admin/events`, `/admin/pages`, `/admin/tags`, `/admin/categories`, `/admin/media`, `/admin/menus`, `/admin/forms`, `/admin/widgets`, `/admin/themes/settings`.
- **Admin-only routes** — `/admin/users`, `/admin/languages`, `/admin/config`, `/admin/themes`, `/admin/modules`, `/admin/api-keys`, `/admin/webhooks`, `/admin/cache`, `/admin/scheduler`, `/admin/export`, `/admin/import`, `/admin/docs`.
- **Bulk actions** — multi-select and bulk delete/revoke on delete-capable lists: pages, tags, users, API keys, media, form submissions. Selection is per-page (current page only), reset on navigation, confirmed by dialog. Partial-success response shape: `{"success": true, "deleted": N, "failed": [{"id": X, "reason": "…"}]}`. Endpoints under `/admin/<view>/bulk-delete` (or `bulk-delete` under form submissions) accept `{"ids": [...]}`.
- **List sorting** — URL-driven via `sort=<field>&dir=asc|desc`. Whitelist-based per view; invalid values fall back to per-view defaults (pages default to `updated_at desc`). Changing sort resets to page 1. User input never interpolated into SQL.
- **Items-per-page selector** — options `10, 20, 50, 100` (media adds `24`); URL state via `per_page`; preserves sort on change.
- **Event logging** — all admin actions recorded at `/admin/events`: content changes, user management, configuration, scheduler events, login attempts.

## Sources

- Origin: `raw/ocms-go.core/wiki/Admin-Interface.md`
