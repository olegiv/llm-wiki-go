# wiki/Menu-Builder.md (oCMS)

## Summary

Describes the oCMS menu builder:

- Drag-and-drop ordering of items.
- Hierarchical (parent-child nested) menu structures.
- Link types: pages, categories, tags, external URLs.
- Multiple menu locations per theme (header, footer, sidebar) configured in `theme.json` widget areas.
- Per-language menus — each language gets its own menu, selected automatically based on current request language.
- Admin workflow: Admin > Menus → **New Menu** → set name and optional location → add items with link type → drag to reorder/nest → Save.
- Theme integration via `{{template "partials/menu.html" .}}`.

Points to Theme System for template integration details and Multi-Language for the per-language workflow.

## Sources

- Origin: `raw/ocms-go.core/wiki/Menu-Builder.md`
