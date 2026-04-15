# Webhooks

## Summary

oCMS webhooks dispatch HMAC-signed HTTP POST requests to external endpoints when internal events occur. Each subscription is a [Webhook](../entities/webhook.md) entity; each dispatch produces a [WebhookDelivery](../entities/webhook-delivery.md) record with its own retry lifecycle.

## Event catalog

Ten events are available:

- `page.created`, `page.updated`, `page.deleted`, `page.published`, `page.unpublished`
- `media.uploaded`, `media.deleted`
- `form.submitted`
- `user.created`, `user.deleted`

## Payload

Top-level envelope:

```json
{
  "type": "page.published",
  "timestamp": "2024-01-15T10:30:00Z",
  "data": { /* event-specific */ }
}
```

Form-submission `data` shape is modulated by `OCMS_WEBHOOK_FORM_DATA_MODE`:

- `redacted` (default) — sensitive-token fields masked, values truncated to 1024 chars.
- `none` — form data omitted.
- `full` — full values, truncated to 1024 chars. Blocked in production by `OCMS_REQUIRE_WEBHOOK_FORM_DATA_MINIMIZATION`.

See [Form entity](../entities/form.md) for the sensitive-field token list.

## Security

- **HMAC-SHA256** signature in `X-Webhook-Signature` over the raw request body.
- Request headers: `Content-Type`, `User-Agent: oCMS/1.0`, `X-Webhook-Signature`, `X-Webhook-Event`, `X-Webhook-Delivery-ID`, plus configured custom headers.
- **Host allowlist.** `OCMS_WEBHOOK_ALLOWED_HOSTS` restricts destinations to exact hostnames. `OCMS_REQUIRE_WEBHOOK_ALLOWED_HOSTS` prevents production startup when active webhooks lack an allowlist.
- **HTTPS outbound.** `OCMS_REQUIRE_HTTPS_OUTBOUND` forces HTTPS on all outbound integration URLs in production.
- **Bounded queue.** Debouncer's pending-event queue is bounded (since 0.18.1).

### Go verifier sample

```go
func verifySignature(body []byte, signature, secret string) bool {
    mac := hmac.New(sha256.New, []byte(secret))
    mac.Write(body)
    expected := hex.EncodeToString(mac.Sum(nil))
    return hmac.Equal([]byte(signature), []byte(expected))
}
```

## Retry and dead-letter

- Retry schedule: 1 min, 2 min, 4 min, 8 min (5 attempts max).
- Retry triggers: 5xx, timeout, 429.
- Not retried: 4xx except 408/429.
- After 5 attempts → `dead`, retained 30 days, manually retriable.

## Debouncing

- 1-second coalesce window.
- 5-second maximum wait before dispatch.
- Rapid edits yield one delivery with final state.

## Monitoring

- `/admin/webhooks` list: last-24h delivery count, success rate, health indicator.
- Per-webhook Deliveries tab: per-delivery status, response code, attempt count, timestamps; manual Retry.

## Best practices

1. Verify signatures; reject on mismatch.
2. 2xx quickly, process async.
3. Idempotent handlers (at-least-once semantics).
4. Log payloads for audit.
5. Monitor and alert on repeated failures.
6. Never expose the shared secret.

## Sources

- [sources/docs-webhooks.md](../sources/docs-webhooks.md)
- [sources/wiki-webhooks.md](../sources/wiki-webhooks.md)
- [sources/security.md](../sources/security.md)
- [sources/changelog.md](../sources/changelog.md)
