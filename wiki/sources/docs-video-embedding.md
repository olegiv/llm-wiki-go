# docs/video-embedding.md (oCMS)

## Summary

Feature reference for page-level video embedding, rendered between the page header and body.

### Supported providers table

| Provider | Status |
|----------|--------|
| YouTube | Supported |
| Vimeo | Planned |

URL formats accepted for YouTube: `youtube.com/watch?v=ID`, `youtu.be/ID`, `youtube.com/embed/ID`, `youtube.com/shorts/ID`, `m.youtube.com/watch?v=ID`.

This list **conflicts with `README.md`, `CHANGELOG.md` 0.16.0 entry, and `wiki/Home.md`**, all of which state YouTube, Vimeo, **and Dailymotion** are supported.

### Authoring flow

Admin page editor → Video section → paste URL, optional `video_title` (max 255 chars) → server-side URL parse → responsive iframe embed with 16:9 aspect ratio in all three core themes (Default, Developer, Starter).

### Security

- Server-side URL parsing; raw URLs never placed directly in iframes.
- YouTube embeds use `youtube-nocookie.com` for privacy.
- Strict 11-character alphanumeric ID validator (`[a-zA-Z0-9_-]{11}`).
- CSP `frame-src` includes `https://www.youtube-nocookie.com`.
- XSS-safe: malicious video IDs rejected by the pattern validator.

### REST API

`video_url` and `video_title` are on the Pages API: `POST /api/v1/pages`, `PUT /api/v1/pages/{id}`. Empty string removes the embed. Included in JSON/ZIP import-export.

### Theme integration

Template snippet:

```html
{{if .Page.VideoEmbedHTML}}
<div class="page-video">
    {{if .Page.VideoTitle}}<h3 class="page-video-title">{{.Page.VideoTitle}}</h3>{{end}}
    {{.Page.VideoEmbedHTML}}
</div>
{{end}}
```

CSS uses padding-bottom: 56.25% for 16:9 responsive containers.

### Adding a provider

Implement `video.Provider` in `internal/video/video.go` → URL matching + ID extraction + embed HTML → register in `NewRegistry()` → add embed domain to CSP `frame-src` in `internal/middleware/security.go` → update translations.

## Sources

- Origin: `raw/ocms-go.core/docs/video-embedding.md`
