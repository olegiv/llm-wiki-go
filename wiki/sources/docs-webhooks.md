# docs/webhooks.md (oCMS)

## Summary

Detailed reference for oCMS webhooks.

### Setup

Admin path: **Admin > Config > Webhooks**. Fields: Name, URL (HTTPS in production), Secret (auto-generated or custom), Events list, custom Headers, Active toggle. Testing via in-UI **Test** button; local dev via webhook.site or ngrok tunnels.

### Events

Ten events: `page.created`, `page.updated`, `page.deleted`, `page.published`, `page.unpublished`, `media.uploaded`, `media.deleted`, `form.submitted`, `user.created`, `user.deleted`.

### Payload shape

Top level: `{type, timestamp, data}`. Data varies by event:

- **Page events:** `id`, `title`, `slug`, `status`, `author_id`, `author_email`, `language_id`.
- **Media events:** `id`, `uuid`, `filename`, `mime_type`, `size`, `uploader_id`.
- **Form submission:** `form_id`, `form_name`, `form_slug`, `submission_id`, `data` (field values), `submitted_at`.

### Security

- **HMAC-SHA256** signature in `X-Webhook-Signature` header; verify against shared secret. Includes Go sample verifier.
- Request headers: `Content-Type: application/json`, `User-Agent: oCMS/1.0`, `X-Webhook-Signature`, `X-Webhook-Event`, `X-Webhook-Delivery-ID`, plus configured custom headers.
- **HTTPS requirement** in production (recommended, not enforced by default; `OCMS_REQUIRE_HTTPS_OUTBOUND` adds enforcement).

### Delivery process

1. Event occurs.
2. Delivery record created.
3. HTTP POST to endpoint.
4. Response recorded.

Success = HTTP 2xx.

### Retry

- **5xx / timeout / 429** → exponential-backoff retry.
- **4xx except 408 and 429** → not retried (client error).
- Schedule: attempt 1 immediate, then 1/2/4/8 minutes, 5 attempts max, then marked **dead**.

Dead deliveries retained 30 days; can be manually retried (resets attempt count).

### Monitoring

Admin **Deliveries** tab per webhook shows status, response code (truncated body), attempt count, timestamps. Webhook list shows last 24-hour delivery count, success rate %, health indicator (green/yellow/red). Manual **Retry** available.

### Event debouncing

Events within a 1-second window coalesce; only the latest is sent. Maximum 5-second wait before dispatch. Ten updates in 2 seconds → 1 webhook with final state.

### Best practices

Verify signatures, 2xx quickly (process async), handle duplicates, idempotent handlers, log payloads, monitor failures, secrets stay secret.

Includes a complete Go webhook handler example (reads body, verifies signature, parses, dispatches).

## Sources

- Origin: `raw/ocms-go.core/docs/webhooks.md`
