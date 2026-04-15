# Form

## Summary

`Form` is the public-form entity in oCMS. A form has a name, URL slug (unique per language), title, description, success message, active flag, language code, timestamps, and an ordered list of fields. Public submissions land in a per-form submissions store; admins view and export them via the admin UI. Forms are rendered through the active theme at `/forms/{slug}`.

## Data model

Three tables:

- **`forms`** — `id`, `name`, `slug`, `title`, `description`, `success_message`, `email_to` (reserved), `is_active`, `language_code`, `created_at`, `updated_at`.
- **`form_fields`** — `id`, `form_id`, `type`, `name`, `label`, `placeholder`, `help_text`, `options` (JSON), `validation` (JSON), `is_required`, `position`, `language_code`, timestamps.
- **`form_submissions`** — `id`, `form_id`, `data` (JSON), `ip_address`, `user_agent`, `is_read`, `language_code`, `created_at`.

## Field types

`text`, `email`, `textarea`, `number`, `date`, `select`, `radio`, `checkbox`, `file`, `captcha`.

`select`/`radio`/`checkbox` store options as a JSON array. Validation rules: JSON `{minLength, maxLength, pattern}`. `required` via `is_required`.

## Public submission flow

1. `GET /forms/{slug}` — theme renders the form.
2. `POST /forms/{slug}` — rate-limited at 1 req/sec, burst 5.
3. Server validates payload size, honeypot, captcha, per-field rules.
4. On success: persist submission, fire `form.submitted` webhook, render success message.
5. On failure: re-render with preserved field values.

## Security

- **Honeypot** via hidden `_website` field. Triggered → WARNING event + `security.honeypot_triggered` hook + auto-ban by Sentinel when active. Bot receives fake success.
- **CSRF** via the global `filippo.io/csrf/gorilla` middleware using Fetch metadata headers. No token embedded in HTML.
- **Payload caps:** body 64 KB, per-field 4 KB, total data 16 KB.
- **Captcha** optional via the hCaptcha module. `requireCaptcha` enforces presence; 503 if enforced without captcha infrastructure. Captcha responses never stored.

## Webhook integration

Event: `form.submitted`. Payload data mode controlled by `OCMS_WEBHOOK_FORM_DATA_MODE`:

- `redacted` (default) — sensitive-token fields masked, values truncated to 1024 chars.
- `none` — payload omits form data.
- `full` — full data, 1024-char truncation.

Sensitive-token list: `password`, `passwd`, `token`, `secret`, `api_key`, `apikey`, `authorization`, `auth`, `ssn`, `credit`, `card`, `cvv`. Production startup blocked on `full` mode by `OCMS_REQUIRE_WEBHOOK_FORM_DATA_MINIMIZATION`.

## Admin surface

- Form builder at `/admin/forms`, `/admin/forms/{id}`, field CRUD via JSON, field reorder (`POST /admin/forms/{id}/fields/reorder`).
- Submissions at `/admin/forms/{id}/submissions`, view at `/admin/forms/{id}/submissions/{subId}`.
- Bulk delete submissions: `POST /admin/forms/{id}/submissions/bulk-delete`.
- CSV export: `POST /admin/forms/{id}/submissions/export`. Hardened in 0.18.1 against CSV formula injection.
- Create translation: `POST /admin/forms/{id}/translate`.

## Multi-language

Forms are per-language; same slug may exist across languages. Translation linking via the shared `translations` table. Creating a translation copies form structure to the target language. Admin UI surfaces missing-translation state.

## Sources

- [sources/docs-forms.md](../sources/docs-forms.md)
- [sources/wiki-forms.md](../sources/wiki-forms.md)
- [sources/readme.md](../sources/readme.md)
- [sources/changelog.md](../sources/changelog.md)
