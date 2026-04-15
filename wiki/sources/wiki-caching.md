# wiki/Caching.md (oCMS)

## Summary

GitHub wiki page for the cache system. Near-parity with `docs/caching.md`: same six cache types with identical invalidation triggers, same context-aware page-cache key shape `{language}:{role}:{slug}`, same role levels (`anonymous`/`public`/`editor`/`admin` at levels 0/0/1/2), same "Admin routes bypass cache" behavior, same in-memory-vs-Redis backends, same environment-variable trio (`OCMS_REDIS_URL`, `OCMS_CACHE_PREFIX`, `OCMS_CACHE_TTL`), same stats UI at `/admin/cache`.

Differences from `docs/caching.md` are entirely presentation, not facts:

- Omits the `cmd/ocms/main.go` Go snippet that shows the admin vs frontend middleware wiring.
- Omits the `cacheManager.Clear*`/`Invalidate*` programmatic API list; just points at the admin UI.
- Omits the `cacheManager.Preload(ctx, siteURL)` Go snippet; summarizes what is preloaded instead.
- Uses Obsidian-style `[[Configuration]]` wikilink.

No factual disagreements.

## Sources

- Origin: `raw/ocms-go.core/wiki/Caching.md`
