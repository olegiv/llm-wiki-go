# Theme system

## Summary

oCMS ships **two core themes** (`default`, `developer`) embedded in the binary, plus the ability to load **custom themes** from `custom/themes/`. Custom themes are filesystem-based Go templates — no Go code required.

Loading precedence: **custom > embedded** when names match. The same-name override lets you shadow `default` or `developer` without patching core code.

## Project layout

```
custom/themes/mytheme/
├── theme.json
├── templates/
│   ├── layouts/base.html
│   ├── pages/home.html, page.html, list.html, 404.html, category.html, tag.html, search.html, form.html
│   └── partials/header.html, footer.html, sidebar.html, pagination.html, post-card.html, language-switcher.html
├── static/css/theme.css
├── static/js/theme.js
└── locales/en/messages.json, ru/messages.json
```

## `theme.json`

Required: `name`, `version`, `author`, `description`, `templates` (map of page types to template paths).

Optional: `screenshot`, `settings` (array), `widget_areas` (array).

### Setting types

`color`, `select`, `image`, `text`. Values surface in templates via `.ThemeSettings` map.

### Widget areas

```json
{
  "widget_areas": [
    {"id": "sidebar", "name": "Sidebar", "description": "Widgets in the sidebar"},
    {"id": "footer-1", "name": "Footer Column 1"}
  ]
}
```

Render: `{{range (index .Widgets "sidebar")}}{{renderWidget . $.ThemeSettings $.LangCode}}{{end}}`.

## Activation

Set `OCMS_ACTIVE_THEME` (env var wins over admin setting) or activate at **Admin > Themes**. Default value: `default`.

## Loader

On startup the theme manager:

1. Loads embedded `default` and `developer` from `internal/themes.FS`.
2. Scans `custom/themes/` for external theme directories.
3. Custom themes override embedded themes with the same name.
4. Parses all templates with the registered function map.
5. Loads per-theme translations from `locales/`.

## Template categories

- **Layouts** (`templates/layouts/`) — named by relative path. Wraps page content.
- **Partials** (`templates/partials/`) — named by filename. Reused via `{{template "header.html" .}}`.
- **Pages** (`templates/pages/`) — each defines `{{define "content"}}` block. Internally parsed with `content_` prefix.

## Integration hooks (required in base layout)

| Hook | Location | Purpose |
|------|----------|---------|
| `{{privacyHead}}` | `<head>` | Privacy consent. |
| `{{analyticsExtHead}}` | `<head>` | External analytics head. |
| `{{analyticsExtBody}}` | Before `</body>` | External analytics body. |
| `{{embedHead .CSPNonce .PageOrigin}}` | `<head>` | Embed head. |
| `{{embedBody .CSPNonce .PageOrigin}}` | Before `</body>` | Embed body. |
| `{{informerBar}}` | In `<body>` | Informer bar widget. |

## Template data (base layout and pages)

`.Title`, `.SiteName`, `.LangCode`, `.LangPrefix`, `.MetaDescription`, `.Canonical`, `.Page`, `.Pages`, `.Categories`, `.Tags`, `.HeaderMenuItems`, `.FooterMenuItems`, `.ThemeSettings`, `.Widgets`, `.Pagination`, `.SearchQuery`, `.Year`.

## Translations

Theme-level locale files override global translations only for the keys they define. Use `TTheme` (theme → global fallback) for frontend strings; use `T` elsewhere.

Common override keys: `frontend.read_more`, `frontend.all_posts`, `frontend.recent_posts`, `frontend.view_all_posts`, `frontend.page_not_found`, `frontend.go_home`, `frontend.related_posts`, `search.title`, `search.placeholder`, `sidebar.categories`, `sidebar.tags`.

See [topics/i18n.md](i18n.md) for the broader i18n system.

## Static assets

Served at `/themes/{name}/static/`. Alpine.js and HTMX are already loaded by oCMS — don't duplicate. Use a unique CSS class prefix, CSS custom properties on `:root` for customization, wrap JS in IIFEs, initialize on `DOMContentLoaded`.

## Overriding a core theme

Create `custom/themes/default/` to completely replace the embedded `default` theme. Copy `internal/themes/default/` as a starting point.

## Architectural history

The embedded-themes design was introduced in 0.2.0 and formally planned in [sources/docs-embed-themes-migration.md](../sources/docs-embed-themes-migration.md). Subsequent releases added widget settings (0.8.0), template function injection points (0.15.0), templUI component migration (0.10.0), and the module HTML injection contract for templ-based frontends (0.15.0, re-hardened in 0.15.0 INJ-001 fix).

## Reference

Starter theme at `custom/themes/starter/` — full template coverage, card-grid magazine layout, theme settings (accent color, sidebar toggle, hero style, favicon), widget areas, EN + RU, responsive. Activate with `OCMS_ACTIVE_THEME=starter`.

Testing: `go test -v ./custom/themes/starter/...` with `OCMS_SESSION_SECRET=…`.

## Sources

- [sources/docs-custom-themes.md](../sources/docs-custom-themes.md)
- [sources/wiki-theme-system.md](../sources/wiki-theme-system.md)
- [sources/docs-embed-themes-migration.md](../sources/docs-embed-themes-migration.md)
- [sources/readme.md](../sources/readme.md)
- [sources/changelog.md](../sources/changelog.md)
