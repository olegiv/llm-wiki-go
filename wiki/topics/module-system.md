# Module system

## Summary

Modules extend oCMS without modifying core code. Built-in modules live under `modules/`; custom modules under `custom/modules/`. Both implement `module.Module` and are wired into the registry at startup. Custom modules self-register via Go's `init()` pattern (same as `database/sql` drivers) and are loaded by blank-importing their packages from `custom/modules/imports.go`.

## Built-in modules (on-disk inventory)

`modules/` under the source tree contains **11** directories:

- `analytics_ext` — external analytics (Google Analytics, etc.) integration.
- `analytics_int` — internal view/read analytics (since 0.18.0).
- `dbmanager` — admin-only direct SQL query execution.
- `developer` — test-data generation for development.
- `embed` — content embed / iframe proxy (Dify integration).
- `example` — reference module using the core renderer.
- `hcaptcha` — bot protection on login; see [hCaptcha topic](hcaptcha.md).
- `informer` — dismissible notification bar.
- `migrator` — import from other CMS systems (including webpage source since 0.16.0).
- `privacy` — Klaro consent, GDPR, Google Consent Mode v2.
- `sentinel` — IP banning, auto-ban paths, honeypot auto-ban, whitelist.

## Custom modules

Three-step setup:

1. Create `custom/modules/mymodule/` with `module.go` (definition), `handlers.go`, `register.go` (self-registration), `templates/` (embedded), `locales/` (embedded).
2. `register.go` calls `module.RegisterCustomModule(New())` in `init()`.
3. Blank-import `_ "github.com/olegiv/ocms-go/custom/modules/mymodule"` in `custom/modules/imports.go`.

Build and run — module appears at **Admin > Modules**.

### Auto-registration flow

1. Each custom module's `init()` calls `RegisterCustomModule(New())`.
2. `custom/modules/imports.go` blank-imports every custom module package.
3. `cmd/ocms/main.go` blank-imports `custom/modules`.
4. `module.CustomModules()` returns the registered list.
5. The main registry initializes customs alongside built-ins.

## `module.Module` interface

```go
type Module interface {
    Name() string
    Version() string
    Description() string
    Dependencies() []string
    Init(ctx *Context) error
    Shutdown() error
    RegisterRoutes(r chi.Router)
    RegisterAdminRoutes(r chi.Router)
    TemplateFuncs() template.FuncMap
    Migrations() []Migration
    AdminURL() string
    SidebarLabel() string
    TranslationsFS() embed.FS
}
```

Embed `module.BaseModule` (via `module.NewBaseModule(name, version, description)`) for no-op defaults.

## `module.Context`

Fields: `DB *sql.DB`, `Store *store.Queries`, `Logger *slog.Logger`, `Config *config.Config`, `Render *render.Renderer`, `Events *service.EventService`, `Hooks *HookRegistry`. Custom modules typically use `DB`, `Logger`, `Hooks`.

## Routes

- `RegisterRoutes(chi.Router)` — public; 404 when module is inactive.
- `RegisterAdminRoutes(chi.Router)` — admin; redirects to modules list when inactive. Auto-prefixed with `/admin/`.

## Migrations

`Migrations() []module.Migration` returns versioned up/down pairs. Tracked in the `module_migrations` table (goose migration `20260213100001`). Version numbers start at 1. Always use parameterized queries — never `fmt.Sprintf` for SQL.

## Hooks

`m.ctx.Hooks.Register(hook, module.HookHandler{Name, Module, Priority, Fn})`. Lower priority runs first. Inactive-module hooks are skipped.

Available hooks:

- `HookPageAfterSave`
- `HookPageBeforeRender`
- `HookSecurityHoneypotTriggered` — see [Form entity](../entities/form.md) honeypot behavior.

## Template functions

`TemplateFuncs() template.FuncMap` provides functions accessible in all themes.

### Templ layout injection (convention names)

| Position | Function names | Typical use |
|----------|---------------|-------------|
| Before `</head>` | `privacyHead`, `analyticsExtHead`, `embedHead` | Consent banners, analytics, embed styles. |
| After `<body>` | `informerBar`, `analyticsExtBody` | Info bars, analytics noscript. |
| Before `</body>` | `embedBody` | Chat widgets, embed scripts. |

Signature: `func(...any) template.HTML`. First argument is the CSP nonce. Templ-based themes receive aggregated output via `BaseTemplateData.ModuleHeadHTML`, `ModuleBodyTopHTML`, `ModuleBodyEndHTML`.

## Embedded admin templates

Custom modules can render their own admin pages via `//go:embed templates/admin.html` parsed in `Init()`. Keeps the module self-contained.

## Internationalization

Standard JSON format at `locales/{lang}/messages.json`. Embed via `//go:embed locales` and expose via `TranslationsFS() embed.FS`. See [topics/i18n.md](i18n.md).

## Active status

Admin > Modules toggle. Active: routes work, sidebar entry, hooks execute, template functions registered. Inactive: public routes 404, admin routes redirect, hooks skipped, template functions not registered. Status persists in DB.

Module active/inactive status was a 0.2.0 feature; 0.10.2 fixed "Fix crash when enabling a module that was inactive at server startup"; 0.15.0 INJ-001 fixed deactivated modules still injecting HTML until server restart.

## Environment restrictions

Implement `EnvironmentChecker.AllowedEnvs() []string`. Module starts inactive if current `OCMS_ENV` is not in the allowed list.

## Testing

```go
db, cleanup := testutil.TestDB(t)
defer cleanup()

m := New()
moduleutil.RunMigrations(t, db, m.Migrations())
ctx, _ := moduleutil.TestModuleContext(t, db)
m.Init(ctx)
```

## Contradictions

### Built-in module inventory

- [sources/wiki-module-system.md](../sources/wiki-module-system.md) lists **10** modules in a **Built-in Modules** table: Developer, DB Manager, Informer, Demo Mode, analytics_ext, analytics_int, embed, hcaptcha, privacy, Sentinel.
- The on-disk `modules/` directory contains **11** entries: `analytics_ext`, `analytics_int`, `dbmanager`, `developer`, `embed`, `example`, `hcaptcha`, `informer`, `migrator`, `privacy`, `sentinel`.
- [sources/wiki-home.md](../sources/wiki-home.md) lists only four modules as "Built-in": Developer, DB Manager, Informer, Demo Mode.
- [sources/agents.md](../sources/agents.md) mentions "analytics, embed, privacy, migrator, etc." — a different subset.

Three-way disagreement. Resolution by observation:

- **"Demo Mode"** appears as a module in some docs but has no `modules/demo-mode` directory — it's implemented as middleware in `internal/middleware/demo.go` and `internal/handler/demo.go`. It is a *mode*, not a pluggable module.
- **`example`** and **`migrator`** exist on disk but are under-documented.
- The "Built-in Modules" inventory in the wiki should include `analytics_int`, `embed`, `privacy`, `migrator`, `example` — and drop "Demo Mode".

## Reference implementation

- `custom/modules/bookmarks/` — CRUD, public JSON API, admin dashboard with embedded template, template functions (`bookmarkCount`, `bookmarkFavorites`), hooks, EN+RU i18n, tests.
- `modules/example/` — built-in reference module using the core renderer.

## Sources

- [sources/docs-custom-modules.md](../sources/docs-custom-modules.md)
- [sources/wiki-module-system.md](../sources/wiki-module-system.md)
- [sources/readme.md](../sources/readme.md)
- [sources/wiki-home.md](../sources/wiki-home.md)
- [sources/agents.md](../sources/agents.md)
- [sources/changelog.md](../sources/changelog.md)
