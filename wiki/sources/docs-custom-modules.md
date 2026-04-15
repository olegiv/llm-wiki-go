# docs/custom-modules.md (oCMS)

## Summary

Guide to creating custom modules. Modules self-register via Go's `init()` pattern (same approach as `database/sql` drivers) and are wired up by blank-importing their packages from `custom/modules/imports.go`.

### Three-step setup

1. **Create the module package** under `custom/modules/mymodule/` with `module.go`, `handlers.go`, `register.go`, embedded `templates/` and `locales/`.
2. **`register.go`** calls `module.RegisterCustomModule(New())` inside `init()`.
3. **Blank import** `_ "github.com/olegiv/ocms-go/custom/modules/mymodule"` in `custom/modules/imports.go`.

Build and run — module appears at **Admin > Modules**.

### Auto-registration flow

1. Each custom module's `init()` calls `RegisterCustomModule(New())`.
2. `custom/modules/imports.go` blank-imports every module package.
3. `cmd/ocms/main.go` blank-imports `custom/modules`, loading all custom modules.
4. `module.CustomModules()` returns the registered list at startup.
5. The main registry initializes customs alongside built-ins.

### `module.Module` interface

`Name()`, `Version()`, `Description()`, `Dependencies()`, `Init(*Context) error`, `Shutdown() error`, `RegisterRoutes(chi.Router)`, `RegisterAdminRoutes(chi.Router)`, `TemplateFuncs() template.FuncMap`, `Migrations() []Migration`, `AdminURL() string`, `SidebarLabel() string`, `TranslationsFS() embed.FS`.

Embed `module.BaseModule` (via `module.NewBaseModule(name, version, description)`) for default no-ops.

### `module.Context` services

`DB` (`*sql.DB`), `Store` (`*store.Queries` sqlc queries), `Logger` (`*slog.Logger`), `Config` (`*config.Config`), `Render` (`*render.Renderer`), `Events` (`*service.EventService`), `Hooks` (`*HookRegistry`).

Custom modules typically only need `DB`, `Logger`, and `Hooks`.

### Route middleware

- Public routes get module active-status check (404 when inactive).
- Admin routes get auth + active-status (redirect to modules list when inactive).
- Admin routes are auto-prefixed with `/admin/`.

### Migrations

`Migrations() []module.Migration` returns versioned up/down pairs. Tracked in `module_migrations` table (goose migration `20260213100001`). Version numbers start at 1. **Always use parameterized queries** (`?` placeholders) — never `fmt.Sprintf` for SQL.

### Hooks

Register via `m.ctx.Hooks.Register(module.HookPageAfterSave, module.HookHandler{Name, Module, Priority, Fn})`. Priorities: lower runs first. Available hooks: `HookPageAfterSave`, `HookPageBeforeRender`. Inactive-module hooks are skipped automatically.

### Template functions

`TemplateFuncs() template.FuncMap` provides functions accessible in all themes.

#### Templ layout injection

Modules can inject HTML into templ-based frontend layouts using convention-named template functions with signature `func(...any) template.HTML`. Position table:

- Before `</head>`: `privacyHead`, `analyticsExtHead`, `embedHead`.
- After `<body>`: `informerBar`, `analyticsExtBody`.
- Before `</body>`: `embedBody`.

First argument is the CSP nonce. Templ-based themes receive aggregated output via `BaseTemplateData.ModuleHeadHTML`, `ModuleBodyTopHTML`, `ModuleBodyEndHTML`.

### Embedded admin templates

Use `//go:embed templates/admin.html` + parse in `Init()`. Keeps module self-contained.

### I18n

Standard JSON format at `locales/{lang}/messages.json`. Embed via `//go:embed locales` + expose via `TranslationsFS() embed.FS`.

### Testing helpers

`testutil.TestDB(t)`, `moduleutil.RunMigrations(t, db, m.Migrations())`, `moduleutil.TestModuleContext(t, db)`.

### Active status

Admin > Modules toggle. Active: routes work, sidebar entry, hooks execute. Inactive: public routes 404, admin routes redirect, hooks skipped. Status persists in DB.

### Environment restrictions

Implement `EnvironmentChecker.AllowedEnvs() []string`. Module starts inactive if current env is not allowed.

### Reference

- `custom/modules/bookmarks/` — full example with CRUD, JSON API, admin template, template functions (`bookmarkCount`, `bookmarkFavorites`), hooks, EN+RU translations, tests.
- `modules/example/` — built-in module using the core renderer.

## Sources

- Origin: `raw/ocms-go.core/docs/custom-modules.md`
