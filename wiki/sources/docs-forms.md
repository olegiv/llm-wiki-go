# docs/forms.md (oCMS)

## Summary

Comprehensive reference for the oCMS form builder. Three-part architecture: form builder admin UI → public rendering at `/forms/{slug}` → submissions management.

### Supported field types

`text`, `email`, `textarea`, `number`, `date`, `select`, `radio`, `checkbox`, `file`, `captcha` (hCaptcha widget, verified via hook, never stored).

### Validation format

JSON rules: `minLength`, `maxLength`, `pattern` (regex). Server-side. Built-in checks for email, number, date. Required via `is_required` flag.

### Public submission flow

Load → validate payload size / honeypot / captcha / field rules → store + dispatch webhook + success message (configurable, default "Thank you for your submission."), or re-render with errors.

### Security

- **Honeypot:** hidden `_website` field. Bots that fill it get a fake-success response and a WARNING security event; `security.honeypot_triggered` hook fires; Sentinel auto-bans IP when active.
- **CSRF:** global middleware (`filippo.io/csrf/gorilla`) using Fetch metadata; no token in HTML.
- **Rate limit:** 1 request/second, burst 5 per IP → 429 on excess.
- **Payload caps:** total body 64 KB; per-field 4 KB; total data 16 KB.

### Captcha

One `captcha` field per form. Rendered via `form.captcha_widget` hook, verified via `form.captcha_verify`. `requireCaptcha` flag mandates captcha; 503 if enabled without captcha module available. Responses never stored.

### Submissions admin

Sortable by ID/created/read/IP. Default sort newest first. Unread badge. Read status auto-marked on view. Bulk delete. CSV export (columns: ID, Submitted At, IP, Read, + each field label).

### Webhooks

Event: `form.submitted`. Payload modes via `OCMS_WEBHOOK_FORM_DATA_MODE`: `redacted` (default, mask sensitive, truncate 1024 chars), `none`, `full`. Sensitive-field token list: `password`, `passwd`, `token`, `secret`, `api_key`, `apikey`, `authorization`, `auth`, `ssn`, `credit`, `card`, `cvv`. Production: `OCMS_REQUIRE_WEBHOOK_FORM_DATA_MINIMIZATION` blocks `full` mode.

### Multi-language

`language_code` column on forms, form_fields, form_submissions. Slugs unique per language. `POST /admin/forms/{id}/translate` creates a translation copy. Admin UI shows available and missing translations.

### Routes

14 admin routes under `/admin/forms/*` plus two public routes (`GET` render, `POST` submit).

### Database schema

Three tables documented in detail: `forms`, `form_fields`, `form_submissions` (columns, types, FKs).

## Sources

- Origin: `raw/ocms-go.core/docs/forms.md`
