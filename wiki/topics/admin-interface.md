# Admin interface

## Summary

The admin panel is served under `/admin/`. It is built with templ components, HTMX for partial updates, and Alpine.js for client-side state, with templUI primitives (buttons, cards, dialogs, selectbox, etc.) since 0.10.0.

## Route map

### Editor + Admin routes

| Route | Purpose |
|-------|---------|
| `/admin/` | Dashboard. |
| `/admin/events` | Event log (read-only pager). |
| `/admin/pages` | Page management (list, new, edit). |
| `/admin/tags` | Tag management. |
| `/admin/categories` | Category management with drag-and-drop reordering. |
| `/admin/media` | Media library (list, upload, edit). |
| `/admin/menus` | Menu management. |
| `/admin/forms` | Form builder and submissions. |
| `/admin/widgets` | Widget management. |
| `/admin/themes/settings` | Theme settings for the active theme. |

### Admin-only routes

| Route | Purpose |
|-------|---------|
| `/admin/users` | User management. |
| `/admin/languages` | Language management. |
| `/admin/config` | Site configuration. |
| `/admin/themes` | Theme list and activation. |
| `/admin/modules` | Module list with active-status toggle. |
| `/admin/api-keys` | API key management. |
| `/admin/webhooks` | Webhook management. |
| `/admin/cache` | Cache statistics and clear. |
| `/admin/scheduler` | Scheduled jobs management. |
| `/admin/export` | Content export. |
| `/admin/import` | Content import. |
| `/admin/docs` | Site documentation. |

## Bulk actions

Supported on delete-capable lists: pages, tags, users, API keys (bulk revoke), media, form submissions. Read-only pagers (events, page versions, webhook deliveries, scheduler runs) have no bulk actions.

- Selection scope is **current page only**; reset on navigation.
- Confirmation dialog before delete/revoke.
- Partial-success model: valid IDs are processed, failures returned in `failed` array with reasons.
- Endpoints: `POST /admin/<view>/bulk-delete` accepting `{"ids": [...]}`.
- Success response shape: `{"success": true, "deleted": N, "failed": [{"id": X, "reason": "…"}]}`.
- Error response: `{"success": false, "error": "..."}`.

## List sorting

- URL-driven: `sort=<field>&dir=asc|desc`. State preserved across pagination and per-page changes.
- **Whitelist-based per view** — user input is never interpolated into SQL; invalid values fall back to per-view defaults.
- Pages default to `sort=updated_at&dir=desc`.
- Clicking the active column toggles direction. Changing sort field resets to page 1.
- Table views use sortable column headers; media grid exposes sort controls in its filter bar.
- Active sort highlighted with pill-style header state.

## Items per page

- Options: `10, 20, 50, 100`. Media also includes `24` for its legacy default.
- URL state via `per_page` query parameter.
- Changing items-per-page resets to page 1 and preserves active sort.

## Event log

`/admin/events` records:

- Content changes (create/update/delete/publish).
- User management actions.
- Configuration changes.
- Scheduler events.
- Login attempts (see [topics/security-overview.md](security-overview.md)).
- Server start/stop events (since 0.10.1).
- 403/404 errors (since 0.4.0), 404 logging restricted to authenticated users in 0.18.1 to limit noise.

Includes a "Ban IP" button for directly banning an IP from an event entry (since 0.6.0), with event URL preserved as audit context.

## Admin UI milestones

- **0.2.0** — Admin UI i18n (English, Russian).
- **0.10.0** — templUI component migration: buttons, badges, cards, alerts, tables, inputs, labels, pagination, page headers, selectbox, dropdowns. Bulk pager actions and per-page selector.
- **0.16.0** — Collapsible admin sidebar with persistent state; page type filter; editable created/published dates in page editor.
- **0.17.0** — Structured event logging added to all REST API handlers.

## Sources

- [sources/wiki-admin-interface.md](../sources/wiki-admin-interface.md)
- [sources/docs-admin-bulk-actions.md](../sources/docs-admin-bulk-actions.md)
- [sources/docs-admin-list-sorting.md](../sources/docs-admin-list-sorting.md)
- [sources/readme.md](../sources/readme.md)
- [sources/claude.md](../sources/claude.md)
- [sources/changelog.md](../sources/changelog.md)
