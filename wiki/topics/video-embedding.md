# Video embedding

## Summary

Pages can embed a single video between the page header and body. The feature shipped in **0.16.0** and is exposed through the admin page editor, the REST API (`video_url`, `video_title`), and JSON/ZIP import-export. All three core themes (Default, Developer, Starter) ship responsive 16:9 container CSS.

## Provider support

Per [sources/readme.md](../sources/readme.md), [sources/changelog.md](../sources/changelog.md) 0.16.0, and [sources/wiki-home.md](../sources/wiki-home.md): **YouTube, Vimeo, and Dailymotion** are all supported.

Per [sources/docs-video-embedding.md](../sources/docs-video-embedding.md): only YouTube is "Supported"; Vimeo is "Planned"; Dailymotion is absent.

See [Contradictions](#contradictions) below.

Accepted YouTube URL shapes:

```
youtube.com/watch?v=ID
youtu.be/ID
youtube.com/embed/ID
youtube.com/shorts/ID
m.youtube.com/watch?v=ID
```

## Authoring flow

1. Open the page editor and expand the **Video** section.
2. Paste a video URL. Optional: `video_title` (max 255 characters), displayed above the embed.
3. Save. URL parses server-side against registered providers.
4. Empty string clears the embed.

## Security

- **Server-side URL parsing.** Raw URLs never flow into iframes.
- **YouTube privacy mode.** Embeds use `youtube-nocookie.com`.
- **ID pattern validator.** YouTube IDs must match `[a-zA-Z0-9_-]{11}`; malicious IDs rejected.
- **CSP `frame-src`.** Includes `https://www.youtube-nocookie.com`.
- **XSS-safe.** Pattern validation rejects HTML/JS in the ID.

## REST API

Fields `video_url` and `video_title` appear on all pages endpoints. Empty string removes the embed.

## Theme integration

```html
{{if .Page.VideoEmbedHTML}}
<div class="page-video">
    {{if .Page.VideoTitle}}<h3 class="page-video-title">{{.Page.VideoTitle}}</h3>{{end}}
    {{.Page.VideoEmbedHTML}}
</div>
{{end}}
```

CSS: `padding-bottom: 56.25%` on `.page-video` creates the 16:9 responsive container.

## Adding a provider

1. Implement `video.Provider` in `internal/video/video.go`.
2. Add URL matching, ID extraction, embed HTML generation.
3. Register in `NewRegistry()`.
4. Add the provider's embed domain to CSP `frame-src` in `internal/middleware/security.go`.
5. Update validation hint translations.

## Contradictions

### Which providers are actually supported?

- [sources/docs-video-embedding.md](../sources/docs-video-embedding.md) — table lists only YouTube as Supported; Vimeo as Planned; no Dailymotion.
- [sources/readme.md](../sources/readme.md) — README feature list says "Embed YouTube, Vimeo, and Dailymotion videos".
- [sources/changelog.md](../sources/changelog.md) — 0.16.0 entry: "Add video embedding widget for pages with YouTube, Vimeo, and Dailymotion support".
- [sources/wiki-home.md](../sources/wiki-home.md) — via the Content Management callout to video embedding, consistent with README.

The balance of sources says all three are supported; `docs/video-embedding.md` is stale (likely a snapshot taken while Vimeo support was in flight). Code under `internal/video/video.go` would resolve this definitively — to be verified in Phase 2 when Go sources are ingested.

## Open questions

- Which provider implementations actually exist in `internal/video/`? Verify post-Phase-2.

## Sources

- [sources/docs-video-embedding.md](../sources/docs-video-embedding.md)
- [sources/readme.md](../sources/readme.md)
- [sources/changelog.md](../sources/changelog.md)
- [sources/wiki-home.md](../sources/wiki-home.md)
