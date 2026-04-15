# Webhook

## Summary

`Webhook` is the outbound-event subscription entity. Each webhook has a name, destination URL (HTTPS recommended; enforced by `OCMS_REQUIRE_HTTPS_OUTBOUND` in production), HMAC-SHA256 signing secret, subscribed event list, optional custom headers, and an active flag. Webhooks are managed at **Admin > Config > Webhooks**.

Each dispatched event produces a [WebhookDelivery](webhook-delivery.md) record that tracks the retry / success lifecycle.

## Events subscribed

Ten event types:

- `page.created`, `page.updated`, `page.deleted`, `page.published`, `page.unpublished`.
- `media.uploaded`, `media.deleted`.
- `form.submitted`.
- `user.created`, `user.deleted`.

## Payload envelope

```json
{
  "type": "page.published",
  "timestamp": "2024-01-15T10:30:00Z",
  "data": { /* event-specific */ }
}
```

Event-specific `data` schemas:

- **Page events.** `id`, `title`, `slug`, `status`, `author_id`, `author_email`, `language_id`.
- **Media events.** `id`, `uuid`, `filename`, `mime_type`, `size`, `uploader_id`.
- **Form submission.** `form_id`, `form_name`, `form_slug`, `submission_id`, `submitted_at`, plus `data` shaped by `OCMS_WEBHOOK_FORM_DATA_MODE` (`redacted`/`none`/`full`).

## Security

- **Signing.** HMAC-SHA256 over the raw request body, delivered in `X-Webhook-Signature`.
- **Request headers:**
  - `Content-Type: application/json`
  - `User-Agent: oCMS/1.0`
  - `X-Webhook-Signature` — HMAC-SHA256
  - `X-Webhook-Event` — event type
  - `X-Webhook-Delivery-ID` — unique delivery id
- **Destination allowlist.** `OCMS_WEBHOOK_ALLOWED_HOSTS` restricts active deliveries to enumerated hosts (exact hostname match). `OCMS_REQUIRE_WEBHOOK_ALLOWED_HOSTS` blocks production startup without an allowlist when active webhooks exist.
- **HTTPS outbound.** `OCMS_REQUIRE_HTTPS_OUTBOUND` forces HTTPS-only outbound URLs in production.
- **Form data minimization.** `OCMS_REQUIRE_WEBHOOK_FORM_DATA_MINIMIZATION` prevents `full` mode in production.
- **Debouncing bounds.** Pending-event queue is bounded to prevent memory growth under event storms (since 0.18.1).

## Event debouncing

Rapid-fire updates coalesce:

- Events within a 1-second window are merged; only the latest state is sent.
- Maximum wait before dispatch: 5 seconds.
- 10 updates in 2 seconds → 1 delivery carrying final state.

## Best practices

- Verify signatures (reject on mismatch).
- Return 2xx immediately; process work async.
- Assume at-least-once delivery; use idempotent handlers.
- Log payloads for audit.

## Sources

- [sources/docs-webhooks.md](../sources/docs-webhooks.md)
- [sources/wiki-webhooks.md](../sources/wiki-webhooks.md)
- [sources/security.md](../sources/security.md)
- [sources/changelog.md](../sources/changelog.md)
