# Caching

## Summary

oCMS uses a multi-layer cache for configuration, sitemap, pages, menus, languages, and translations. Two backends: in-memory (default) and optional Redis (for multi-instance deployments). Cache state is observable and clearable at `/admin/cache`.

## Cache types

| Cache | Purpose | Invalidation |
|-------|---------|--------------|
| Config | Site settings (`site_name`, etc.) | On config update. |
| Sitemap | Generated `sitemap.xml`. | On page / category / tag changes. |
| Pages | Published page content. | On page create / update / delete. |
| Menus | Menu structure + items. | On menu or item change. |
| Languages | Active languages list. | On language change. |
| Translations | Per-entity translations. | On translation change. |

## Context-aware page cache

Page cache keys are shaped `{language}:{role}:{slug}`, so content is cached per language and per access level.

### Role levels

| Role | Level | Description |
|------|-------|-------------|
| `anonymous` | 0 | Not logged in. |
| `public` | 0 | Logged in, no special permissions. |
| `editor` | 1 | Can edit content. |
| `admin` | 2 | Full system access. |

`anonymous` and `public` share level 0 but have distinct cache keys (so logged-in-but-role-less traffic is not served anonymous-cached content).

## Admin bypass

Admin routes (`/admin/*`) bypass the cache for site-configuration lookups — admins always see fresh data. Per [sources/docs-caching.md](../sources/docs-caching.md), this is wired explicitly in `cmd/ocms/main.go`:

```go
// Admin: no cache
r.Use(middleware.LoadSiteConfig(db, nil))

// Frontend: cached
r.Use(middleware.LoadSiteConfig(db, cacheManager))
```

## Backends

- **In-memory (default).** Single-instance; no external dependency; lost on restart. Governed by `OCMS_CACHE_MAX_SIZE` (default `10000`).
- **Redis.** Multi-instance distributed cache; persists across restarts. Activated by setting `OCMS_REDIS_URL`. Keys are prefixed by `OCMS_CACHE_PREFIX` (default `ocms:`).

Default TTL for entries: `OCMS_CACHE_TTL` (default `3600` seconds).

## Statistics and manual clearing

- `/admin/cache` surfaces hit rate, item count, and backend status.
- Programmatic API (from [sources/docs-caching.md](../sources/docs-caching.md)):
  ```go
  cacheManager.ClearAll()
  cacheManager.InvalidateConfig()
  cacheManager.InvalidateSitemap()
  cacheManager.InvalidatePages()
  cacheManager.InvalidatePage(id)
  cacheManager.InvalidateMenus()
  cacheManager.InvalidateLanguages()
  cacheManager.InvalidateTranslations()
  ```

## Automatic invalidation

| Change | Invalidates |
|--------|-------------|
| Page created/updated/deleted | Page cache + Sitemap cache. |
| Menu or menu item changed | Menu cache. |
| Config changed | Config cache. |
| Language changed | Language cache. |
| Translation changed | Translation cache for that entity. |

Publish-toggle cache invalidation was a 0.18.1 fix — the page cache now clears when publishing or unpublishing a page.

## Preload on startup

`cacheManager.Preload(ctx, siteURL)` loads:

- Site configuration.
- All menus.
- Active languages.
- Sitemap (only when `siteURL` is configured — this requirement tightened in 0.18.1: sitemap generation now requires `site_url` in site config and the sitemap cache is keyed by `site_url` rather than request `Host`).

## Sources

- [sources/docs-caching.md](../sources/docs-caching.md)
- [sources/wiki-caching.md](../sources/wiki-caching.md)
- [sources/changelog.md](../sources/changelog.md)
- [sources/readme.md](../sources/readme.md)
