# wiki/Taxonomy.md (oCMS)

## Summary

Describes the two oCMS taxonomy types:

- **Categories** — hierarchical (parent-child), public pages at `/category/{slug}`, per-language translations, full CRUD at `/admin/categories` with drag-and-drop reordering.
- **Tags** — flat (non-hierarchical), public pages at `/tag/{slug}`, per-language translations, full CRUD at `/admin/tags`.

Both are exposed on the REST API: `GET /api/v1/categories` returns a tree, `GET /api/v1/tags` returns a flat list. Both admin lists support bulk actions, column sorting, and configurable pagination.

## Sources

- Origin: `raw/ocms-go.core/wiki/Taxonomy.md`
