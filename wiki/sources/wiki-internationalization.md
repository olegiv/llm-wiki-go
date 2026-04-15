# wiki/Internationalization.md (oCMS)

## Summary

Wiki version of the i18n reference. Near-parity with `docs/i18n.md`: same scope (admin, modules, themes), same JSON file format, same `{namespace}.{category}.{key}` ID convention, same `TTheme` vs `T` template helper distinction, same supported languages (`en`, `ru`), same JSON Schema validation setup, same Add-a-Language workflow.

Differences from `docs/i18n.md`:

- Uses the full Go module path `github.com/olegiv/ocms-go/internal/i18n` in the example (docs uses the shorter `ocms-go/internal/i18n`).
- Cross-links with Obsidian wikilinks to `[[Module System]]`, `[[Theme System]]`, `[[Multi-Language]]`.
- Points readers at the content-translation system (Multi-Language) for context.
- Slightly condensed best-practices section.

No factual disagreements.

## Sources

- Origin: `raw/ocms-go.core/wiki/Internationalization.md`
