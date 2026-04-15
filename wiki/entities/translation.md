# Translation

## Summary

`Translation` links content entities across [languages](language.md). Pages, categories, tags, menus, and forms are all translatable. Each entity belongs to exactly one language; translation links form a bidirectional graph across a language family.

Translation is distinct from the admin UI's i18n system (for UI strings). This entity covers **content** translation — the user-facing data — while [topics/i18n.md](../topics/i18n.md) covers **UI** translation.

## Translatable entities

- [Page](page.md)
- [Category](category.md)
- [Tag](tag.md)
- [Menu](menu.md) — one menu per language
- [Form](form.md) — slug unique per language

## Creating a translation

In the admin editor of the source entity, the **Translations** panel lists each configured target language. Clicking **Add Translation** creates a sibling entity in that language:

- Page: same title, empty body, bidirectional link.
- Category / Tag: same structure, language-bound slug, bidirectional link.
- Menu / Form: copy the shape into the target language.

Deleting an entity removes its translation links; the siblings remain.

## REST API shape

A page resource exposes its siblings:

```json
{
  "data": {
    "id": 1,
    "language_id": 1,
    "language_code": "en",
    "translations": [
      {"language_code": "ru", "page_id": 5, "slug": "o-nas"}
    ]
  }
}
```

## Hreflang

The theme engine emits hreflang link tags for each known translation plus an `x-default` pointing to the default-language variant.

```html
<link rel="alternate" hreflang="en" href="https://example.com/about-us">
<link rel="alternate" hreflang="ru" href="https://example.com/ru/about-us">
<link rel="alternate" hreflang="x-default" href="https://example.com/about-us">
```

## Cache

The translations cache (one of six in [topics/caching.md](../topics/caching.md)) stores per-entity translation metadata and is invalidated on any translation change.

## Sources

- [sources/docs-multi-language.md](../sources/docs-multi-language.md)
- [sources/wiki-multi-language.md](../sources/wiki-multi-language.md)
- [sources/docs-forms.md](../sources/docs-forms.md)
- [sources/wiki-taxonomy.md](../sources/wiki-taxonomy.md)
- [sources/readme.md](../sources/readme.md)
