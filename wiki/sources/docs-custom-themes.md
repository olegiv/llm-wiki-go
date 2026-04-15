# docs/custom-themes.md (oCMS)

## Summary

Comprehensive guide to building custom themes. Themes live under `custom/themes/` and require **no Go code** — `theme.json` + HTML templates + static assets + optional `locales/` directory.

### Directory structure

Canonical layout:

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

### `theme.json`

Required fields: `name`, `version`, `author`, `description`, `templates` object (mapping page types to template paths). Optional: `screenshot`, `settings`, `widget_areas`.

### Activation

Set `OCMS_ACTIVE_THEME=mytheme` (env var wins over admin setting) or activate in **Admin > Themes**.

### Loader behavior

1. Load embedded core themes (`default`, `developer`) from binary.
2. Scan `custom/themes/` for external directories.
3. External themes override embedded themes with the same name.
4. Parse all templates with the registered function map.
5. Load theme-specific translations from `locales/`.

### Settings

Types: `color`, `select`, `image`, `text`. Accessible in templates via `.ThemeSettings` map.

### Widget areas

Defined in `theme.json`; rendered via `{{range (index .Widgets "sidebar")}}{{renderWidget . $.ThemeSettings $.LangCode}}{{end}}`.

### Template categories

- **Layouts** — `templates/layouts/`; named by relative path.
- **Partials** — `templates/partials/`; named by filename; included with `{{template "header.html" .}}`.
- **Pages** — `templates/pages/`; each defines `{{define "content"}}` block injected into the base layout.

### Required integration hooks

| Hook | Location | Purpose |
|------|----------|---------|
| `{{privacyHead}}` | `<head>` | Privacy consent. |
| `{{analyticsExtHead}}` | `<head>` | External analytics head scripts. |
| `{{analyticsExtBody}}` | Before `</body>` | External analytics body scripts. |
| `{{embedHead .CSPNonce .PageOrigin}}` | `<head>` | Embed head content. |
| `{{embedBody .CSPNonce .PageOrigin}}` | Before `</body>` | Embed body content. |
| `{{informerBar}}` | In `<body>` | Informer bar widget. |

### Template data fields

`.Title`, `.SiteName`, `.LangCode`, `.LangPrefix`, `.MetaDescription`, `.Canonical`, `.Page`, `.Pages`, `.Categories`, `.Tags`, `.HeaderMenuItems`, `.FooterMenuItems`, `.ThemeSettings`, `.Widgets`, `.Pagination`, `.SearchQuery`, `.Year`.

### Translations

Use `TTheme` helper (theme → global fallback); `T` for global-only. File at `locales/{lang}/messages.json` with the standard schema.

### Common override keys

`frontend.read_more`, `frontend.all_posts`, `frontend.recent_posts`, `frontend.view_all_posts`, `frontend.page_not_found`, `frontend.go_home`, `frontend.related_posts`, `search.title`, `search.placeholder`, `sidebar.categories`, `sidebar.tags`.

### Static assets

Served at `/themes/{name}/static/`. Best practices: unique class prefix (e.g. `mt-`), CSS custom properties in `:root`, IIFE JS, use `DOMContentLoaded`, don't include Alpine.js / HTMX (oCMS loads them).

### Overriding core themes

Create a custom theme with the same name as a core theme (e.g. `custom/themes/default/`). The custom version completely replaces the embedded version.

### Testing

`go test -v ./custom/themes/starter/...` with `OCMS_SESSION_SECRET=…`.

### Reference

Starter theme at `custom/themes/starter/` — full page coverage, card-grid magazine layout, theme settings (accent color, sidebar toggle, hero style, favicon), widget areas (sidebar, footer columns), English + Russian, responsive design, CSS custom properties, comprehensive tests.

Activate with `OCMS_ACTIVE_THEME=starter`.

## Sources

- Origin: `raw/ocms-go.core/docs/custom-themes.md`
