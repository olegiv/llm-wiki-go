# docs/admin-bulk-actions.md (oCMS)

## Summary

Reference for bulk actions on admin lists. Enumerates supported views (pages, tags, users, API keys as bulk revoke, media, form submissions) and excluded read-only pagers (events, page versions, webhook deliveries, scheduler runs). Documents behavior: current-page selection scope, reset on navigation, partial-success model, confirmation dialog, post-action page reload, items-per-page selector (URL-only via `per_page`, options `10, 20, 50, 100`, media adds `24`), and URL-driven sorting preserved across bulk actions. Includes JSON request (`{"ids":[1,2,3]}`) and response shapes (`{"success":true,"deleted":2,"failed":[{"id":3,"reason":"..."}]}`) plus endpoint list.

Overlaps heavily with the bulk-actions sections of `wiki/Admin-Interface.md`.

## Sources

- Origin: `raw/ocms-go.core/docs/admin-bulk-actions.md`
