# Category

## Summary

`Category` is a hierarchical taxonomy type in oCMS. A category has a name, slug, optional parent (for nesting), and per-language translations. Each category has a public page at `/category/{slug}` listing its content. Categories are managed at `/admin/categories` with full CRUD, drag-and-drop reordering, and column sorting.

Categories and [tags](tag.md) together form the oCMS taxonomy system.

## Key fields and behavior

- **Hierarchy.** Tree structure via `parent_id`. Root categories have no parent.
- **Slug.** Used in the public URL `/category/{slug}` and in API responses.
- **Translations.** Categories can be translated across site languages; translation links preserve language prefix (fix shipped in 0.18.1 for the category translation redirect path).
- **API.** `GET /api/v1/categories` returns the full tree (public).
- **Admin listings.** Support bulk delete, column sorting, and configurable pagination.
- **Page cards.** Page cards in themes show all assigned categories as pills (since 0.15.0).

## Sources

- [sources/wiki-taxonomy.md](../sources/wiki-taxonomy.md)
- [sources/readme.md](../sources/readme.md)
- [sources/changelog.md](../sources/changelog.md)
