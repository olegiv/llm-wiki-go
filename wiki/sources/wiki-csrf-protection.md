# wiki/CSRF-Protection.md (oCMS)

## Summary

Wiki version of the CSRF protection reference. Covers the same Fetch-metadata-header approach (`Sec-Fetch-Site`) with Origin/Referer fallback, same `host:port` format requirement for `TrustedOrigins`, same development-vs-production defaults (`localhost:8080` + `127.0.0.1:8080` in dev; same-origin only in prod), same AJAX behavior (no tokens needed), same API-route exclusion, and same troubleshooting table.

Differences from `docs/csrf.md`:

- Omits the `filippo.io/csrf/gorilla` library name (treats CSRF as a generic feature).
- Omits the `CSRFConfig` struct / `internal/middleware/csrf.go` Go details.
- Omits the legacy `{{.CSRFField}}`/`{{.CSRFToken}}` template variable note.
- Omits the `middleware.SkipCSRF("/api/v1/...")` code snippet.
- Omits the session-secret / 32-byte key length security consideration.
- Uses Obsidian wikilinks (`[[REST API]]`, `[[Login Security]]`, `[[hCaptcha]]`, `[[Security]]`).

No factual disagreements.

## Sources

- Origin: `raw/ocms-go.core/wiki/CSRF-Protection.md`
