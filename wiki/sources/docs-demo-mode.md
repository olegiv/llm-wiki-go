# docs/demo-mode.md (oCMS)

## Summary

Reference for demo mode as a security feature. When `OCMS_DEMO_MODE=true`, destructive admin operations are blocked while non-destructive exploration remains possible.

### Restrictions table (more permissive version)

- **Pages:** Create / Edit **Allowed**; Delete **Blocked**.
- **Media:** Create / Edit **Allowed**; Delete **Blocked**.
- **User management:** **Blocked**.
- **Site configuration:** **Blocked**.
- **Language management:** **Blocked**.
- **API key management:** **Blocked**.
- **Webhook management:** **Blocked**.
- **Module management:** **Blocked**.
- **Cache clearing:** **Blocked**.
- **Data export / import:** **Blocked**.
- **Delete tags / categories:** **Blocked**.
- **Delete menus / items:** **Blocked**.
- **Delete forms / widgets:** **Blocked**.

**Directly contradicts `docs/demo-deployment.md`**, which declares Create/Edit of content "Blocked". See [topics/demo-mode.md](../topics/demo-mode.md).

### Enabling

`OCMS_DEMO_MODE=true` (env var or `fly.toml` `[env]` block).

### User experience

- Web UI: redirect back with flash message.
- API: HTTP 403 with explanation.

Example flash messages: "Deleting pages is disabled in demo mode", "User management is disabled in demo mode", "API key management is disabled in demo mode".

### Allowed actions

Browse admin; create/edit pages, tags, categories, menus, forms; upload media (2 MB cap); view settings; use REST API for reads.

### Upload size limit

2 MB (vs 20 MB normal). Constant: `DemoUploadMaxSize = 2 * 1024 * 1024`.

### Scheduled reset

Recommended for public demos: daily DB reset via `reset-demo.sh` + `fly machines restart`.

### Implementation

- `internal/middleware/demo.go` — core middleware + helpers.
- `internal/handler/demo.go` — handler-level guards.

Adding a new demo guard:

```go
func (h *MyHandler) DangerousAction(w http.ResponseWriter, r *http.Request) {
    if demoGuard(w, r, h.renderer, middleware.RestrictionMyAction, "/admin/myroute") {
        return
    }
    // ...
}
```

For JSON/HTMX:

```go
if middleware.IsDemoMode() {
    writeJSONError(w, http.StatusForbidden, middleware.DemoModeMessageDetailed(middleware.RestrictionMyAction))
    return
}
```

Add new restriction type: define `RestrictionMyAction` constant in `internal/middleware/demo.go`, add message in `DemoModeMessageDetailed`, wire the guard.

### Security scope

Prevents: data deletion, config tampering, user privilege escalation, external service abuse (webhooks), data exfiltration, storage exhaustion.

Does NOT prevent: creating spam content (use scheduled reset), accessing existing admin features (use separate demo credentials), viewing sensitive information (don't use real data).

## Sources

- Origin: `raw/ocms-go.core/docs/demo-mode.md`
