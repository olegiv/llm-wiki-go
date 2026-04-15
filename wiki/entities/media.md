# Media

## Summary

`Media` is the uploaded-file entity in oCMS. Each item has a UUID, original filename, MIME type, optional folder, timestamps, and (for images) a set of automatically generated size variants produced by libvips on upload. Media is managed at `/admin/media` and exposed through the REST API. Uploads are stored under `uploads/{variant}/{UUID}/{filename}` with a one-week public `Cache-Control` header.

## Image variants

Per [sources/docs-media.md](../sources/docs-media.md), **seven** variants are generated on image upload:

| Variant | Dimensions | Quality | Mode | Purpose |
|---------|------------|---------|------|---------|
| `originals` | Original size | 95% | — | Archive/download |
| `large` | 1920×1080 max | 90% | Fit | Single-page featured images |
| `og` | 1200×630 max | 85% | Fit | Open Graph / social sharing (< 200 KB target) |
| `medium` | 800×600 max | 85% | Fit | Listings (homepage, categories, tags) |
| `small` | 400×300 max | 85% | Fit | Compact listings |
| `grid` | 256×256 | 85% | Crop | Admin grid views |
| `thumbnail` | 150×150 | 80% | Crop | Search results, admin previews |

OG image selection priority: `og > large > medium > thumbnail`.

## Supported formats

- **Images with variants:** JPEG, PNG, GIF, WebP.
- **Documents (no variants):** PDF.
- **Other:** stored in `originals` only, no variants.
- **TIFF:** blocked to mitigate CVE-2023-36308 (since 0.1.0).

## Featured-image validation

- Minimum 1200 × 800 pixels on **new** assignment.
- Existing assignments below the minimum are grandfathered (saving unrelated fields does not re-validate).
- Since 0.18.1, unchanged media skips featured-image revalidation on save.

## REST API

`GET /api/v1/media`, `POST /api/v1/media` (upload, requires `media:write`). Responses include a `urls` object with one URL per variant.

## Events and webhooks

Upload and delete events are captured in the admin event log (since 0.2.0). There is no dedicated `media.*` webhook event in the current webhook schema.

## Admin features

- Bulk delete via `POST /admin/media/bulk-delete`.
- Folder organization.
- Sorting by name, date, size; grid view exposes sort controls in the filter bar.
- Configurable items-per-page (`10, 20, 24, 50, 100`). `24` is preserved as the media legacy default.
- Regenerate-variants button on the media edit page (since 0.3.0).
- Grid image variant used for admin media cards (since 0.10.0).
- Unified media dropzone component across admin screens (since 0.3.0).
- Media picker security hardened in 0.2.0 (CodeQL alert #28: URL sanitization before `img.src`).

## Contradictions

### Variant count
- [sources/docs-media.md](../sources/docs-media.md) documents **seven** variants (adds `og`, `small`, `grid`).
- [sources/wiki-media-library.md](../sources/wiki-media-library.md) documents **four** (`originals`, `large`, `medium`, `thumbnail`).
- [sources/readme.md](../sources/readme.md)'s "Automatic thumbnail and variant generation" bullet is non-specific.

The seven-variant table in `docs/media.md` is consistent with [sources/changelog.md](../sources/changelog.md) (0.18.0 added the `og` variant; 0.4.0 added `small` in the developer-module test-data generator; 0.10.0 added `grid`). The wiki page has not been updated for these additions.

### OG image source
- [sources/docs-media.md](../sources/docs-media.md) selects a dedicated `og` variant (with fallback priority `og > large > medium > thumbnail`).
- [sources/wiki-media-library.md](../sources/wiki-media-library.md) uses `large` for OG and notes "social platforms resize anyway".

Docs reflect the post-0.18.0 behavior; wiki is stale.

### Upload directory configurability
- [sources/docs-media.md](../sources/docs-media.md): "This path is not configurable via environment variables."
- [sources/wiki-media-library.md](../sources/wiki-media-library.md) and [sources/readme.md](../sources/readme.md): configured via `OCMS_UPLOADS_DIR` (default `./uploads`).
- [sources/changelog.md](../sources/changelog.md) 0.16.0: "Add `OCMS_UPLOADS_DIR` support in migrator module".

`OCMS_UPLOADS_DIR` clearly exists. The `docs/media.md` sentence is stale.

## Sources

- [sources/docs-media.md](../sources/docs-media.md)
- [sources/wiki-media-library.md](../sources/wiki-media-library.md)
- [sources/readme.md](../sources/readme.md)
- [sources/changelog.md](../sources/changelog.md)
