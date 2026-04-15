# docs/informer-module.md (oCMS)

## Summary

Reference for the Informer module — a dismissible notification bar with a spinning indicator rendered above page content.

### Features

- Dismissible (cookie `ocms_informer_dismissed`, 365-day expiry).
- Auto-reset on admin settings change (internal version counter increments, invalidating all cookies).
- HTML support in notification text (rendered as-is; admin-only editing means no escaping needed).
- Customizable background and text colors.
- Translatable via locale files (module UI itself is EN + RU).
- Auto-enabled in demo mode with admin-access hint.

### Admin settings

At `/admin/informer` (a.k.a. **Admin > Modules > Informer**):

| Setting | Default |
|---------|---------|
| Enabled | Off |
| Notification Text | Empty (HTML allowed) |
| Background Color | `#1e40af` (blue) |
| Text Color | `#ffffff` (white) |

Live preview in the admin UI.

### Cookie-based dismissal flow

1. Visitor clicks close button.
2. Bar hides immediately.
3. Cookie set to current settings version (integer), 365-day expiry.
4. Subsequent visits stay hidden while cookie version matches.

Admin settings save increments the version → all existing cookies invalidated → bar reappears for everyone.

### Demo mode

`OCMS_DEMO_MODE=true` auto-enables the bar at startup with admin-access info and demo credentials.

### Theme integration

Template function `informerBar` invoked after `<body>` in the base layout:

```html
<body>
    {{informerBar}}
    ...
</body>
```

Both core themes (`default`, `developer`) include this hook. Custom themes must add the same line.

### Translation

The admin-entered notification text is used as-is. For multi-language sites, admins can define translated `informer.bar_text` entries in theme or global locale files.

### Technical metadata

- Module name: `informer`.
- Database table: `informer_settings` (single row, id=1).
- Cookie name: `ocms_informer_dismissed`.
- Cookie value: integer version counter.
- Admin URL: `/admin/informer`.
- CSS animation: `@keyframes informer-spin` for the spinner; inline styles for colors.

### Hardening (0.7.0 and later)

- 0.7.0 initial release.
- 0.10.1 — informerBar template-error fix when module inactive; module no-op template placeholders added.
- 0.18.1 — Sanitize informer text instead of escaping (from 0.9.0 fix); resilience fixes.

## Sources

- Origin: `raw/ocms-go.core/docs/informer-module.md`
