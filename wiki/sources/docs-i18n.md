# docs/i18n.md (oCMS)

## Summary

Reference for the internationalization (i18n) system used by the admin interface, modules, and themes.

### Scope

- Admin UI translations in `internal/i18n/locales/{lang}/messages.json`.
- Module-embedded translations via `//go:embed locales` + `TranslationsFS() embed.FS`.
- **Theme-level translations** under `themes/{theme_name}/locales/`.
- Default languages: English and Russian.
- JSON Schema validation via `.schema/i18n-schema.json`.
- Dynamic language switching from the admin UI.

### Translation file format

```json
{
    "$schema": "../../../../.schema/i18n-schema.json",
    "language": "en",
    "messages": [
        {"id": "module.key", "message": "Original text", "translation": "Translated text"}
    ]
}
```

Required fields: `language`, `messages`, `messages[].id`, `messages[].translation`. Optional: `$schema`, `messages[].message`.

### ID naming

`{namespace}.{category}.{key}` — dot notation. Conventions: `module.title`, `module.description`, `module.field_name`, `module.field_name_help`, `module.error_*`, `module.success_*`, `module.button_*`.

### Template helpers

| Function | Use case | Fallback |
|----------|----------|----------|
| `TTheme` | Theme (frontend) templates | Theme → Global |
| `T` | Admin and module templates | Global only |

### Go usage

```go
import "ocms-go/internal/i18n"
message := i18n.T(lang, "mymodule.success_save")
```

### Adding a language

1. Create `internal/i18n/locales/{lang}/messages.json` and each module's `locales/{lang}/messages.json`.
2. Copy English as template, update `language` field, translate entries.
3. Register language in DB or configuration.

### VS Code schema setup

```json
{
    "json.schemas": [
        {"fileMatch": ["**/locales/*/messages.json"], "url": "./.schema/i18n-schema.json"}
    ]
}
```

### Loading

Translations load automatically during module init: `TranslationsFS()` → per-language directory read → merge with core → available via `i18n.T()`.

### Supported languages

`en` (English), `ru` (Russian).

## Sources

- Origin: `raw/ocms-go.core/docs/i18n.md`
