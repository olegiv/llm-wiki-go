# docs/media.md (oCMS)

## Summary

Detailed reference for the media library. Enumerates **seven** image variants generated on upload: `originals` (archive, 95% quality), `large` (1920×1080 max, 90%, fit), `og` (1200×630 max, 85%, fit — for social sharing, target < 200 KB), `medium` (800×600 max, 85%, fit), `small` (400×300 max, 85%, fit), `grid` (256×256, 85%, crop — for admin grid), `thumbnail` (150×150, 80%, crop). Storage path shape: `uploads/{variant}/{UUID}/{filename}`. Public URL form: `/uploads/{variant}/{UUID}/{filename}`. Frontend usage table maps context (single page, OG, listing, archives, search, admin picker) to optimal variant. OG image selection priority: `og > large > medium > thumbnail`.

REST API response returns URLs for all variants under a `urls` object. Featured-image validation requires minimum 1200×800 on assignment; pre-existing smaller images are grandfathered. Supported formats: JPEG/PNG/GIF/WebP (with variants), PDF (originals only). Uploads served with `Cache-Control: public, max-age=604800` (1 week).

**States that the upload directory is hardcoded at `./uploads` and "not configurable via environment variables."** This contradicts the wiki page and README.

## Sources

- Origin: `raw/ocms-go.core/docs/media.md`
