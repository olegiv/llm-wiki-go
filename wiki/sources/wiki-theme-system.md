# wiki/Theme-System.md (oCMS)

## Summary

Wiki overview of the theme system. Near-parity with `docs/custom-themes.md`: same directory layout, same `theme.json` schema (`name`, `version`, `author`, `description`, `templates`, optional `settings`/`widget_areas`), same activation (env var `OCMS_ACTIVE_THEME` or Admin UI), same setting types (`color`, `select`, `image`, `text`), same widget area API, same template categorization (layouts / partials / pages), same integration hook list, same `TTheme` vs `T` distinction, same override-core-themes mechanism (same-name custom theme replaces embedded), same reference to the Starter theme.

Differences from `docs/custom-themes.md`:

- Opens with a one-line overview ("Core themes embedded, custom themes in `custom/themes/`, custom takes priority, no Go code needed").
- Omits the explicit 5-step loader description.
- Shorter "How It Works" exposition.
- Uses Obsidian wikilinks for cross-references (`[[Internationalization]]`, `[[Module: Informer|Informer bar]]`).
- Shorter "Common Override Keys" table (9 keys vs the 11 in docs).
- Omits the CSS and JavaScript best-practices section.
- Omits the testing section.

No factual disagreements.

## Sources

- Origin: `raw/ocms-go.core/wiki/Theme-System.md`
