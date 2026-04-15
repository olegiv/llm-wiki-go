# Language

## Summary

`Language` is the configured-language entity in oCMS. Each language has an ISO 639-1 code, English name, native name, direction (LTR or RTL), active flag, and position (for switcher ordering). Exactly one configured language is marked **default** and serves as the fallback when translations are missing or no preference is detected.

Languages are managed at `/admin/languages` (admin-only route) and `/admin/config/languages` (as referenced by the docs).

## Key fields

| Field | Purpose |
|-------|---------|
| `code` | ISO 639-1 (e.g. `en`, `ru`, `de`). |
| `name` | English name (e.g. "Russian"). |
| `native_name` | Name in the language (e.g. "Русский"). |
| `direction` | `ltr` or `rtl`. RTL causes `<html dir="rtl">`. |
| `active` | Whether it is available on the site. |
| `position` | Language-switcher ordering. Drag-and-drop in admin. |
| (default flag) | Exactly one language is marked default. |

## Language detection (public)

Resolution order on each request:

1. URL prefix (`/ru/about-us` → `ru`).
2. Cookie preference set by the switcher.
3. `Accept-Language` header.
4. Default language.

## URL structure

- Default language: `/about-us`.
- Other languages: `/ru/about-us`, `/de/about-us`.

## REST API

- Filter with `?language=<code>`: `GET /api/v1/pages?language=ru`.
- Page responses include `language_id`, `language_code`, and a `translations[]` array of sibling pages.
- Responses carry a `Content-Language` HTTP header (per [sources/docs-multi-language.md](../sources/docs-multi-language.md); not mentioned in the wiki version).

## Admin UI localization

Distinct from content languages: the admin interface itself is currently translated to English and Russian. Adding an admin language requires `internal/i18n/locales/{lang}/messages.json` plus registration in `SupportedLanguages` (`internal/i18n/i18n.go`). See [topics/i18n.md](../topics/i18n.md).

## Cache

The language cache (one of six cache types in [topics/caching.md](../topics/caching.md)) stores the list of active languages and is invalidated on any language change.

## Sources

- [sources/docs-multi-language.md](../sources/docs-multi-language.md)
- [sources/wiki-multi-language.md](../sources/wiki-multi-language.md)
- [sources/docs-i18n.md](../sources/docs-i18n.md)
- [sources/readme.md](../sources/readme.md)
