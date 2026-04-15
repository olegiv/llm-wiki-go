# Tag

## Summary

`Tag` is a flat, non-hierarchical taxonomy type in oCMS. A tag has a name, slug, and per-language translations. Each tag has a public page at `/tag/{slug}` listing tagged content. Tags are managed at `/admin/tags` with full CRUD.

Tags and [categories](category.md) together form the oCMS taxonomy system.

## Key fields and behavior

- **Slug.** Used in the public URL `/tag/{slug}` and in API responses.
- **Translations.** Tags can be translated; admin list supports search and language filters (since 0.10.0).
- **API.** `GET /api/v1/tags` returns a flat list (public). `POST /api/v1/pages` auto-creates tags by name since 0.15.0, returning 422 on tag validation errors.
- **Ordering.** Tag clouds in themes order by usage count (not alphabetically) since 0.3.0.
- **Admin listings.** Support bulk delete, column sorting, and configurable pagination.
- **Legacy URL redirect.** `/blog/tag/{slug}` redirects to `/tag/{slug}` (since 0.5.0; CWE-601 open-redirect fix applied).

## Sources

- [sources/wiki-taxonomy.md](../sources/wiki-taxonomy.md)
- [sources/readme.md](../sources/readme.md)
- [sources/changelog.md](../sources/changelog.md)
