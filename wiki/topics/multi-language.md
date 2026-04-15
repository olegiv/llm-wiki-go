# Multi-language

## Summary

The content-translation system in oCMS covers pages, categories, tags, menus, and forms. It layers on top of a configurable [Language](../entities/language.md) registry and exposes sibling relationships as [Translation](../entities/translation.md) edges.

See [topics/i18n.md](i18n.md) for the separate UI-string translation system.

## Language setup

At **Admin > Languages** (`/admin/languages`):

- Add each language with ISO 639-1 code, English name, native name, LTR/RTL direction, active flag, and switcher position.
- Mark exactly one as **default**. Default is fallback when:
  - A requested translation is missing.
  - No language preference is detected.
  - The URL lacks a language prefix.

## URL structure

- Default language: `/about-us`.
- Non-default: `/ru/about-us`, `/de/about-us`.

## Language detection (priority order)

1. **URL prefix** (`/ru/page-slug` → Russian).
2. **Cookie** preference from previous selection.
3. **`Accept-Language`** header.
4. **Default** language.

## Translating content

In a page/category/tag/menu/form editor, the **Translations** panel lists target languages. Clicking **Add Translation**:

- Creates a sibling entity bound to the target language.
- Copies the title (editable) and leaves body empty (for pages).
- Adds a **bidirectional** translation link.

Deleting an entity removes its outgoing translation links but leaves siblings intact.

## Frontend integration

Theme template variables:

```html
{{.CurrentLanguage}}
{{.LanguageDirection}}
{{range .Translations}}
    <a href="{{.URL}}">{{.NativeName}}</a>
{{end}}
```

Language switcher partial:

```html
{{template "partials/language_switcher.html" .}}
```

Shows current language, translations for the current page, and home-page links for languages without a translation.

## Hreflang (SEO)

Auto-emitted for each language of the current page, plus `x-default`:

```html
<link rel="alternate" hreflang="en" href="https://example.com/about-us">
<link rel="alternate" hreflang="ru" href="https://example.com/ru/about-us">
<link rel="alternate" hreflang="x-default" href="https://example.com/about-us">
```

## RTL

Setting a language's Direction to RTL causes `<html dir="rtl">` automatically. Themes must provide RTL CSS.

## Menus per language

Create one menu per language (e.g. "Main Menu - EN", "Main Menu - RU"). Theme loads the menu matching the active request language. See [Menu entity](../entities/menu.md).

## Forms per language

Forms carry a `language_code`. Same slug may exist per language; translations are linked via the shared translations table. See [Form entity](../entities/form.md).

## REST API

- `GET /api/v1/pages?language=ru` — filter by language.
- Page responses include `language_id`, `language_code`, and `translations[]` siblings.
- Response includes a `Content-Language` header per [sources/docs-multi-language.md](../sources/docs-multi-language.md) (omitted from the wiki version — see contradictions).

## Contradictions

### `Content-Language` header coverage

- [sources/docs-multi-language.md](../sources/docs-multi-language.md): API responses include a `Content-Language` response header indicating the language of the response content.
- [sources/wiki-multi-language.md](../sources/wiki-multi-language.md): omits the `Content-Language` header — its "Language Headers" section is absent.

Documentation drift, not a factual disagreement. The header behavior is described only in the docs tree.

## Troubleshooting

- **Translation missing from switcher** — ensure target page is published, the link exists in the Translations panel, target language is active.
- **Wrong language displayed** — check URL prefix, clear cookies, verify detection order.
- **Missing admin translations** — check `messages.json` exists, keys present, restart server.

## Sources

- [sources/docs-multi-language.md](../sources/docs-multi-language.md)
- [sources/wiki-multi-language.md](../sources/wiki-multi-language.md)
- [sources/wiki-seo.md](../sources/wiki-seo.md)
- [sources/readme.md](../sources/readme.md)
