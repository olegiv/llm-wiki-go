# docs/login-security.md (oCMS)

## Summary

Reference for the login protection system — two layers: IP rate limiting and account lockout.

### IP-based rate limiting

- Token bucket per IP: 5 tokens, regen at 0.5/sec.
- Each login attempt consumes 1 token.
- Exhausted → HTTP 429 "Too many requests. Please try again later."

### Account lockout

- 5 failed attempts within a 15-minute window → account locked.
- Base lockout: 15 minutes. Doubles each subsequent lockout up to 24 hours.
- Lockout progression: 15 m → 30 m → 1 h → 2 h → … → 24 h cap.
- Failed attempts are tracked **even for non-existent emails** to prevent user enumeration.
- A successful login clears the failed-attempt counter.
- Stale tracking entries auto-purged.

### User messages

- Attempts 1–2: "Invalid email or password".
- Attempts 3–4: "Invalid email or password. N attempts remaining."
- Attempt 5: "Too many failed attempts. Account locked for 15 minutes."
- During lockout: "Account is temporarily locked. Try again in N minutes."

### Event logging

Five event types logged at WARNING/INFO levels; viewable at `/admin/events`.

### Configuration

Settings are code-defined in `cmd/ocms/main.go` via `middleware.NewLoginProtection(middleware.LoginProtectionConfig{...})` with fields `IPRateLimit`, `IPBurst`, `MaxFailedAttempts`, `LockoutDuration`, `AttemptWindow`. Not settable via environment variables.

### Reverse proxy headers

Middleware checks in order: `X-Real-IP`, `X-Forwarded-For`, `RemoteAddr`.

### Best practices

Strong passwords, monitor events, HTTPS, CAPTCHA/2FA/IP allowlist for high-security deployments.

### Troubleshooting

Locked-out users: wait or restart (clears in-memory state). Document notes: "In a future version, admin users will be able to unlock accounts from the admin panel" — not yet implemented.

## Sources

- Origin: `raw/ocms-go.core/docs/login-security.md`
