# docs/embed-themes-migration.md (oCMS)

## Summary

Historical migration plan for embedding core themes into the binary and establishing a unified `custom/` directory. This is a design document — **the migration has been completed** (see [sources/changelog.md](changelog.md) 0.2.0 "Embedded Themes Architecture") and the doc is retained in the repo as a record.

### Motivation

Before 0.2.0, `themes/` was external and required at runtime; deployments needed the themes directory alongside the binary. Problems: binary couldn't run standalone, no clear place for user themes, multiple env vars for multiple directories.

### Target architecture

- `internal/themes/` with `//go:embed all:default all:developer` — core themes embedded.
- `custom/` directory with `themes/` and `modules/` subdirectories — user content, gitignored.
- Loading priority: custom overrides embedded when names match.

### Nine-phase plan

1. Prepare embedded themes structure — move `themes/*` into `internal/themes/`, add `embed.go`.
2. Update theme manager — dual-source loading (embedded + filesystem).
3. Update static file handler — serve embedded or filesystem assets based on theme source.
4. Create `custom/` directory structure with `.gitkeep` and README.
5. Configuration — add `OCMS_CUSTOM_DIR` (default `./custom`); deprecate `OCMS_THEMES_DIR`.
6. Application bootstrap — pass embedded FS to theme manager; create `custom/` on startup if missing.
7. Documentation updates — CLAUDE.md env var table, new `docs/custom-content.md`.
8. Deployment and sync procedures — `deploy.sh` becomes custom-theme aware; `setup-site.sh` drops theme copying; `sync-prod-to-dev.sh` gets `--sync-custom`; `backup-multi.sh` includes `custom/`; README rewritten.
9. Testing — unit tests for dual-source loading, override priority; integration tests for static serving from both sources; manual and deployment tests.

### Backward compatibility

- `OCMS_THEMES_DIR`: kept but deprecation-warned; maps to `$OCMS_CUSTOM_DIR/themes`.
- Existing external themes: move `themes/mytheme/` to `custom/themes/mytheme/`.
- Docker volume mounts: old `-v ./themes:/app/themes` → new `-v ./custom:/app/custom`.
- Server deployments: binary update works immediately; custom themes migrate via `mv themes custom/themes`; update `.env`.

### Rollback plan

Revert theme manager + move themes back + remove embed.go + restore old config handling + revert deploy scripts + on server restore `/opt/ocms/themes/`.

### Open questions (historical)

1. Theme settings persistence across embed migration.
2. Theme screenshots accessibility post-embed.
3. Template hot-reload in development (filesystem build tag?).
4. Module custom directory — deferred.
5. Server migration automation (script to move themes + update `.env`).
6. `deploy.sh` vhost requirement — optional?

These were all resolved or deferred by the actual 0.2.0 implementation.

## Relation to current state

This doc describes the plan, not the implementation. For the realized design, see `topics/theme-system.md` and `topics/deployment.md`.

## Sources

- Origin: `raw/ocms-go.core/docs/embed-themes-migration.md`
