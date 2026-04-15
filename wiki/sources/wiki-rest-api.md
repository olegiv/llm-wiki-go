# wiki/REST-API.md (oCMS)

## Summary

Wiki-only reference page for the oCMS REST API.

### Authentication

Bearer token via `Authorization: Bearer <api-key>` header. Keys created at **Admin > API Keys**. Features: fine-grained permissions (read/write per resource), per-key token-bucket rate limiting, CIDR restrictions, expiration, revocation.

### Endpoint table

- `GET /api/v1/pages` (optional auth) — list published pages.
- `GET /api/v1/pages/{id}` / `GET /api/v1/pages/slug/{slug}` — read by id/slug.
- `POST` / `PUT` / `DELETE /api/v1/pages[/{id}]` — require auth.
- `GET /api/v1/media` (optional), `POST /api/v1/media` (required).
- `GET /api/v1/tags`, `GET /api/v1/categories` (tree) — public.
- `GET /api/v1/docs` — interactive API documentation. Also mirrored at `/admin/docs`.

### Page fields

Includes `video_url` (embeddable URL for YouTube/Vimeo/Dailymotion) and `video_title` (optional title above the embed).

### Response shape

Success wraps payload: `{"data": …, "meta": {"total", "page", "per_page"}}`. Error wraps `{"error": {"code", "message", "details"}}`.

### Pagination

`?page=N&per_page=N` (defaults 1 / 20).

### Filtering

`?language=ru` — filter by language.

### Rate limit

Default **100 requests/minute per key**, configurable via `OCMS_API_RATE_LIMIT`. Exhaustion → HTTP 429.

### Health endpoints

- `GET /health` — overall status.
- `GET /health/live` — liveness, always `{"status":"alive"}`.
- `GET /health/ready` — readiness (DB check).

Public returns minimal status; authenticated (admin session or API key) returns full details (uptime, version, DB latency, disk space). `?verbose=true` adds Go runtime info.

### Security

- API routes are excluded from CSRF Protection — token auth covers them.
- All queries parameterized via sqlc.

## Sources

- Origin: `raw/ocms-go.core/wiki/REST-API.md`
