# wiki/Login-Security.md (oCMS)

## Summary

Wiki-version reference for login protection. Near-parity with `docs/login-security.md`: same two layers (IP rate limiting, account lockout), same numeric values (0.5 req/s burst 5; 5 failures in 15 min; 15 m base doubling to 24 h), same non-existent-email enumeration protection, same success-resets-counter behavior, same event-logging table, same troubleshooting guidance.

Differences from `docs/login-security.md`:

- Omits the `middleware.NewLoginProtection` Go config snippet.
- Omits the "In a future version" note about admin-initiated account unlock.
- Wikilinks `[[Reverse Proxy]]` and `[[hCaptcha]]` instead of local markdown paths.
- Adds recommendation to configure `OCMS_TRUSTED_PROXIES` under Best Practices (not present in docs version).

No factual disagreements.

## Sources

- Origin: `raw/ocms-go.core/wiki/Login-Security.md`
