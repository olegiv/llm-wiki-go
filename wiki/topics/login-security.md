# Login security

## Summary

Two defense layers guard the login form.

## IP rate limiting

Token bucket per source IP:

| Parameter | Value |
|-----------|-------|
| Burst size | 5 tokens |
| Regen rate | 0.5 tokens/second (i.e. one attempt per 2 seconds sustained) |
| Exhaustion response | HTTP 429 with body `Too many requests. Please try again later.` |

Each login attempt consumes one token.

## Account lockout

Tracked per **email address** (even for non-existent accounts, preventing user enumeration).

| Parameter | Value |
|-----------|-------|
| Max failed attempts | 5 |
| Attempt window | 15 minutes |
| Base lockout | 15 minutes |
| Max lockout | 24 hours |
| Backoff | Doubling each subsequent lockout |

Progression: 15 m → 30 m → 1 h → 2 h → … → 24 h cap. A successful login clears the failed-attempt counter.

### User-facing messages

| Attempt | Message |
|---------|---------|
| 1–2 | "Invalid email or password" |
| 3–4 | "Invalid email or password. N attempts remaining." |
| 5 | "Too many failed attempts. Account locked for 15 minutes." |
| During lockout | "Account is temporarily locked. Try again in N minutes." |

## Event logging

| Event | Level |
|-------|-------|
| Login failed: user not found | WARNING |
| Login failed: invalid password | WARNING |
| Login attempt on locked account | WARNING |
| Account locked due to failed attempts | WARNING |
| User logged in | INFO |

Viewable at `/admin/events`.

## Configuration

Login-protection parameters are defined in `cmd/ocms/main.go` via `middleware.NewLoginProtection(middleware.LoginProtectionConfig{…})`. They are **not** settable via environment variables. To customize, edit the struct literal:

```go
middleware.LoginProtectionConfig{
    IPRateLimit:       0.5,
    IPBurst:           5,
    MaxFailedAttempts: 5,
    LockoutDuration:   15 * time.Minute,
    AttemptWindow:     15 * time.Minute,
}
```

## Reverse proxy and client IP

The middleware checks headers in order: `X-Real-IP`, `X-Forwarded-For`, `RemoteAddr`. Configure reverse proxies to forward the real client IP (e.g. nginx `proxy_set_header X-Real-IP $remote_addr;`). Set `OCMS_TRUSTED_PROXIES` so the trusted-proxy-aware IP resolution (since 0.17.0) accepts these headers; the chi `RealIP` replacement rejects forwarding headers from untrusted peers.

## Recovery

Legitimate user locked out:

1. Wait for the lockout to expire, **or**
2. Restart the application (in-memory lockout state is wiped).

Admin-initiated account unlock is described in [sources/docs-login-security.md](../sources/docs-login-security.md) as "a future version" — not yet implemented.

## Third-layer defense

For bot protection before the lockout trigger, enable [hCaptcha](hcaptcha.md).

## Sources

- [sources/docs-login-security.md](../sources/docs-login-security.md)
- [sources/wiki-login-security.md](../sources/wiki-login-security.md)
- [sources/security.md](../sources/security.md)
- [sources/wiki-security.md](../sources/wiki-security.md)
- [sources/changelog.md](../sources/changelog.md)
