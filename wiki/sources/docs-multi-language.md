# docs/multi-language.md (oCMS)

## Summary

Reference for the two-part multi-language system: content translation and admin UI localization.

### Language setup

Admin path: `/admin/config/languages` (referred to as **Admin > Config > Languages**). Language fields: ISO 639-1 **Code**, English **Name**, **Native Name**, **Direction** (LTR/RTL), **Active** toggle, **Position** (for switcher ordering).

One language is marked **default**: fallback when a requested translation is missing, no preference is detected, or the URL omits a language prefix.

### Content translation

Pages, categories, tags, and menus can be translated. The page editor has a **Translations** panel; clicking **Add Translation** for a target language creates a sibling page with shared title and empty body. Translations are **bidirectionally linked**; deleting a page removes the translation links.

### Frontend URL structure

- Default language: `/about-us`.
- Other languages: `/ru/about-us`, `/de/about-us`.

### Language detection order

1. URL prefix (`/ru/…`).
2. Cookie preference (from previous selection).
3. `Accept-Language` request header.
4. Default language (fallback).

### Language switcher

Theme partial `{{template "partials/language_switcher.html" .}}`. Shows current language, available translations for the current page, and links to the homepage in other languages when no translation exists.

### RTL

Setting **Direction = RTL** causes the `<html>` tag to receive `dir="rtl"` automatically. Themes should include RTL-specific CSS.

### Admin UI localization

Users switch language from the admin header. Supported: `en`, `ru`. Adding more requires `internal/i18n/locales/{lang}/messages.json` plus registration in `SupportedLanguages` (`internal/i18n/i18n.go`) and a rebuild.

### Theme integration

Template variables: `{{.CurrentLanguage}}`, `{{.LanguageDirection}}`, `{{range .Translations}}…{{end}}`.

### Hreflang

Automatic hreflang tag emission for SEO:

```html
<link rel="alternate" hreflang="en" href="https://example.com/about-us">
<link rel="alternate" hreflang="ru" href="https://example.com/ru/about-us">
<link rel="alternate" hreflang="x-default" href="https://example.com/about-us">
```

### Menu translation

Create one menu per language; theme chooses the right menu based on current language.

### API

- Language filter: `GET /api/v1/pages?language=ru`.
- Response carries `language_id`, `language_code`, and a `translations[]` array of `{language_code, page_id, slug}`.
- **`Content-Language` HTTP response header** indicates the language of the response content.

### Troubleshooting

Translation missing from switcher → check publication status, link existence, target language active. Wrong language → check URL prefix, cookies, detection order. Missing admin translations → verify `messages.json` exists, keys present, server restarted.

## Sources

- Origin: `raw/ocms-go.core/docs/multi-language.md`
