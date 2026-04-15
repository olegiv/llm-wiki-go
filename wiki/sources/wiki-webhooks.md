# wiki/Webhooks.md (oCMS)

## Summary

Wiki version of the webhooks reference. Near-parity with `docs/webhooks.md`: same setup flow at **Admin > Config > Webhooks**, same list of **10 events**, same payload `{type, timestamp, data}` shape, same event-specific data schemas (page, media, form submission), same HMAC-SHA256 `X-Webhook-Signature` verification, same 5-header request set, same retry policy (1, 2, 4, 8 minutes, max 5 attempts, 30-day dead retention), same 1-second debouncing + 5-second cap.

Differences from `docs/webhooks.md`:

- Form submission payload data modes (`OCMS_WEBHOOK_FORM_DATA_MODE`: `redacted`/`none`/`full`) are called out in a note; docs has it on forms page only.
- Destination host allowlisting (`OCMS_WEBHOOK_ALLOWED_HOSTS`) is mentioned explicitly here; docs does not in the webhooks file.
- Shorter webhook handler example (omits the full `package main` scaffolding of the docs version).
- Omits the "HTTPS requirement" prose in favor of pointing at Configuration.
- Uses Obsidian wikilinks to `[[Configuration]]` and `[[Security]]`.

No factual disagreements.

## Sources

- Origin: `raw/ocms-go.core/wiki/Webhooks.md`
