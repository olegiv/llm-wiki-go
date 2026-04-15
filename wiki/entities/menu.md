# Menu

## Summary

`Menu` is a named navigation structure in oCMS. A menu has a name, an optional location identifier, an associated language, and an ordered list of menu items. Menu items can be nested to form hierarchies and can link to pages, categories, tags, or external URLs. Menus are managed at `/admin/menus` with a drag-and-drop editor for reordering and nesting.

## Key fields and behavior

- **Location identifier.** Themes declare named locations (e.g. `header`, `footer`, `sidebar`) in `theme.json` widget areas; a menu's location selects which slot it fills.
- **Per-language.** For multilingual sites, create one menu per language. The theme loads the menu that matches the active request language.
- **Link types per item.** Page (internal), category, tag, or external URL.
- **Nesting.** Items have `parent_id` pointing at another item in the same menu; `parent_id` of root items is normalized (fix shipped in 0.18.1).
- **Position.** Items are sorted by position; drag-and-drop in the admin updates positions.
- **URL scheme validation.** External URL items validate scheme (reject `javascript:`, `data:`); hardened since 0.18.1.
- **Unpublished pages.** Menu items referencing unpublished pages are skipped in Dify knowledge-base generation (since 0.15.0).

## Theme integration

Themes load menus via partials, typically:

```html
{{template "partials/menu.html" .}}
```

## Sources

- [sources/wiki-menu-builder.md](../sources/wiki-menu-builder.md)
- [sources/wiki-home.md](../sources/wiki-home.md)
- [sources/readme.md](../sources/readme.md)
- [sources/changelog.md](../sources/changelog.md)
