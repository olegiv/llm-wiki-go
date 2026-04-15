# docs/admin-list-sorting.md (oCMS)

## Summary

Reference for column sorting on admin pagers. Lists supported views (pages, tags, users, API keys, media, form submissions) and unchanged read-only pagers. Documents query parameters (`sort=<field>`, `dir=asc|desc`) and UI behavior (sortable column headers on table views, grid-bar sort controls on media, active-sort pill highlighting, direction-toggle on re-click, page reset on sort field change, fallback to per-view defaults on invalid inputs). Pages default to `sort=updated_at&dir=desc`. Sort fields are whitelisted per view; user input is never interpolated into SQL ORDER BY.

Overlaps with the sorting section of `wiki/Admin-Interface.md`.

## Sources

- Origin: `raw/ocms-go.core/docs/admin-list-sorting.md`
