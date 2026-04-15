# REST API

## Summary

oCMS exposes a versioned REST API at `/api/v1/*` for programmatic access to content. Authentication is via bearer [API Key](../entities/api-key.md). Most list endpoints accept pagination, language filtering, and return an envelope with `data` + `meta`. Built-in interactive documentation is available at `/api/v1/docs` and `/admin/docs`.

## Authentication

```bash
curl -H "Authorization: Bearer your-api-key" http://localhost:8080/api/v1/pages
```

API middleware chain (admin middleware is parallel): `APIKeyAuth → RequirePermission → APIRateLimit`. CSRF does not apply — token auth covers state-changing requests.

## Endpoints

| Method | Endpoint | Description | Auth |
|--------|----------|-------------|------|
| GET | `/api/v1/pages` | List published pages | Optional |
| GET | `/api/v1/pages/{id}` | Get page by ID | Optional |
| GET | `/api/v1/pages/slug/{slug}` | Get page by slug | Optional |
| POST | `/api/v1/pages` | Create page | Required (`pages:write`) |
| PUT | `/api/v1/pages/{id}` | Update page | Required |
| DELETE | `/api/v1/pages/{id}` | Delete page | Required |
| GET | `/api/v1/media` | List media | Optional |
| POST | `/api/v1/media` | Upload media | Required (`media:write`) |
| GET | `/api/v1/tags` | List tags | Public |
| GET | `/api/v1/categories` | List categories (tree) | Public |
| GET | `/api/v1/docs` | API documentation | Public |

Draft-page reads require `pages:read` (since 0.18.1).

## Response envelope

```json
{
  "data": { },
  "meta": { "total": 100, "page": 1, "per_page": 20 }
}
```

Error shape:

```json
{
  "error": {
    "code": "validation_error",
    "message": "Validation failed",
    "details": { "title": "Title is required" }
  }
}
```

Validation errors on tag creation return **HTTP 422** (since 0.15.0 for API tag auto-creation).

## Pagination

`?page=N` (default `1`), `?per_page=N` (default `20`).

## Filtering

- `?language=<code>` — e.g. `?language=ru`. See [multi-language](multi-language.md).

## Rate limiting

Per-key token-bucket, default **100 requests/minute per key**. Configurable via `OCMS_API_RATE_LIMIT`. Exhaustion → HTTP 429.

## Health endpoints

Separate from `/api/v1`:

| Endpoint | Purpose |
|----------|---------|
| `GET /health` | Overall status. |
| `GET /health/live` | Always `{"status":"alive"}`. |
| `GET /health/ready` | DB connectivity check. |

**Public** (unauthenticated): minimal `{"status":"healthy"}` or `{"status":"degraded"}`.

**Authenticated** (admin session or API key): full details — uptime, version, DB latency, disk space. `?verbose=true` adds Go runtime info (goroutines, memory, CPU count).

## Structured logging (since 0.17.0)

REST API handlers emit structured events through `apiLogger`. API key names were intentionally scrubbed from event metadata in 0.17.0 to prevent topology leakage.

## Sources

- [sources/wiki-rest-api.md](../sources/wiki-rest-api.md)
- [sources/readme.md](../sources/readme.md)
- [sources/claude.md](../sources/claude.md)
- [sources/changelog.md](../sources/changelog.md)
