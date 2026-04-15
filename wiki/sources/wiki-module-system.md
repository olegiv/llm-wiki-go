# wiki/Module-System.md (oCMS)

## Summary

Wiki overview of the module system. Covers the same four-step creation flow, same `module.Module` interface, same `module.Context` field set (sans `Render`), same public / admin route patterns, same migration shape, same hook registration, same template-function mechanism, same i18n conventions, same active-status toggle, same environment restrictions.

### Built-in modules inventory

Lists ten modules: **Developer**, **DB Manager**, **Informer**, **Demo Mode**, **analytics_ext**, **analytics_int**, **embed**, **hcaptcha**, **privacy**, **Sentinel**. On disk, the project's `modules/` directory also contains `example` (a built-in reference module for contributors), plus `migrator` and `privacy` — see [Contradictions below](#contradictions) for discrepancies with the actual on-disk inventory.

Differences from `docs/custom-modules.md`:

- Adds a **Built-in Modules** table as an introductory inventory.
- Omits the `Render` field from the `module.Context` table.
- Adds `module.HookSecurityHoneypotTriggered` to the hooks list (docs omits this one).
- Omits the "Templ layout injection" section — prose and convention-named functions are not described.
- Omits the "Embedded admin templates" section.
- Omits the "Testing helpers" section.
- Uses Obsidian-style wikilinks to individual module pages.

No factual disagreements on the core interface. The on-disk inventory versus the wiki inventory is a separate discrepancy — see Contradictions.

## Contradictions

### Built-in module list vs on-disk `modules/`

The wiki enumerates ten modules. On disk, `ocms-go.core/modules/` contains eleven directories: `analytics_ext`, `analytics_int`, `dbmanager`, `developer`, `embed`, `example`, `hcaptcha`, `informer`, `migrator`, `privacy`, `sentinel`. Differences:

- `example` and `migrator` exist on disk but are not in the wiki's Built-in Modules table.
- "Demo Mode" is listed in the wiki table but has no dedicated directory under `modules/` — it is implemented inline in `internal/middleware/demo.go` and `internal/handler/demo.go` rather than as a pluggable module.

Captured fully on [topics/module-system.md](../topics/module-system.md).

## Sources

- Origin: `raw/ocms-go.core/wiki/Module-System.md`
