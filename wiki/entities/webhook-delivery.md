# WebhookDelivery

## Summary

`WebhookDelivery` records a single dispatch attempt of an event to a [Webhook](webhook.md). Each delivery carries attempt count, HTTP response code, truncated response body, attempt timestamps, and a status — `pending`, `delivered`, `failed`, or `dead`. Viewable under **Admin > Config > Webhooks > *webhook* > Deliveries**.

## Success criteria

HTTP **2xx** from the destination.

## Retry policy

| Attempt | Delay before next attempt |
|---------|--------------------------|
| 1 | Immediate |
| 2 | 1 minute |
| 3 | 2 minutes |
| 4 | 4 minutes |
| 5 | 8 minutes |

Maximum attempts: **5**. Then status → **dead**.

Retry-eligible failure modes:

- HTTP **5xx** — retried with exponential backoff.
- **Timeout** — retried with exponential backoff.
- **HTTP 429** — retried with exponential backoff.
- HTTP **4xx** (except 408, 429) — **NOT** retried (client error, no point).

## Dead-letter queue

Dead deliveries are retained **30 days** and can be manually retried from the admin UI (resets attempt counter). This is implemented via the dead-letter path visible on the webhook deliveries page.

## Monitoring

Per-delivery display: status, response code, response body (truncated), attempt count, all attempt timestamps. Per-webhook summary in the list view: last-24h deliveries, success-rate percentage, health indicator (green / yellow / red).

Manual **Retry** button resets the delivery's attempt counter and re-dispatches.

## Bounded queue

Since 0.18.1, the debouncer's pending-event queue is bounded to prevent unbounded growth during upstream outages.

## Sources

- [sources/docs-webhooks.md](../sources/docs-webhooks.md)
- [sources/wiki-webhooks.md](../sources/wiki-webhooks.md)
- [sources/changelog.md](../sources/changelog.md)
