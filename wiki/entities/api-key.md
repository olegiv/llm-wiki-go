# API Key

## Summary

`APIKey` is the REST-API authentication credential entity. Each key has an Argon2id-hashed bearer token, per-key permission flags (read/write per resource), optional expiration timestamp, optional per-key source-CIDR allowlist, rate-limit settings, active flag, and author/audit metadata. Keys are managed at `/admin/api-keys` and consumed via the `Authorization: Bearer <token>` header on `/api/v1/*` endpoints.

## Key fields and behavior

- **Hash algorithm.** Argon2id on the full key (since 0.1.0, replacing SHA-256). See [sources/changelog.md](../sources/changelog.md).
- **Permissions.** Fine-grained read/write flags per resource (`pages:read`, `pages:write`, `media:write`, etc.). Draft page reads require `pages:read` (enforced since 0.18.1).
- **Rate limit.** Token-bucket per key. Default **100 req/min per key**, tunable via `OCMS_API_RATE_LIMIT` (see [topics/configuration.md](../topics/configuration.md)).
- **Per-key CIDR allowlist.** Restricts the key to specific source IP ranges. `OCMS_REQUIRE_API_KEY_SOURCE_CIDRS` forces a value in production.
- **Source-IP anomaly detection.** When a key with no per-key CIDRs is used from an unexpected source IP, the key is auto-revoked (`OCMS_REVOKE_API_KEY_ON_SOURCE_IP_CHANGE`, production default `true`).
- **Expiration.** Keys can require expiry timestamps (`OCMS_REQUIRE_API_KEY_EXPIRY`). Maximum lifetime capped by `OCMS_API_KEY_MAX_TTL_DAYS` (default 90 days in production). Legacy non-expiring keys receive admin-UI warnings (since 0.9.0).
- **Global source-CIDR policy.** `OCMS_API_ALLOWED_CIDRS` gates all API-key traffic regardless of per-key settings.
- **Revocation.** Bulk revoke on the admin list (`/admin/api-keys` bulk action) deactivates without deletion. API-key bulk-delete endpoint: `POST /admin/api-keys/bulk-delete`.
- **Throttled verification.** Concurrent Argon2 verification is throttled to prevent DoS (since 0.18.1).

## Event logging

Since 0.17.0, all REST API handlers log structured events per request with a category-scoped `apiLogger`. API key names were removed from event metadata to prevent topology leakage via log exports.

## Sources

- [sources/wiki-rest-api.md](../sources/wiki-rest-api.md)
- [sources/security.md](../sources/security.md)
- [sources/wiki-security.md](../sources/wiki-security.md)
- [sources/readme.md](../sources/readme.md)
- [sources/changelog.md](../sources/changelog.md)
