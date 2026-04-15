# docs/caching.md (oCMS)

## Summary

Reference for oCMS's multi-layer cache. Six cache types (Config, Sitemap, Pages, Menus, Languages, Translations) with documented invalidation triggers. Context-aware page-cache keys of shape `{language}:{role}:{slug}` — e.g. `en:anonymous:about-us`. Role levels: `anonymous` 0, `public` 0, `editor` 1, `admin` 2.

Admin routes (`/admin/*`) bypass the cache for site-config lookups. Frontend routes use the cache. Both behaviours wired explicitly in `cmd/ocms/main.go` using `middleware.LoadSiteConfig(db, nil)` (admin) vs `middleware.LoadSiteConfig(db, cacheManager)` (frontend).

Cache backends: in-memory (default, single-instance, lost on restart) and Redis (multi-instance, persists across restarts). Configuration via `OCMS_REDIS_URL`, `OCMS_CACHE_PREFIX`, `OCMS_CACHE_TTL`.

Cache stats UI at `/admin/cache` — hit rate per type, cached item count, backend status. Programmatic manual invalidation via `cacheManager.Clear*` / `cacheManager.Invalidate*` methods (9 methods documented). On startup, `cacheManager.Preload(ctx, siteURL)` loads site config, menus, active languages, and sitemap (if siteURL given).

## Sources

- Origin: `raw/ocms-go.core/docs/caching.md`
