# wiki/Media-Library.md (oCMS)

## Summary

GitHub wiki page for the media library. Lists **four** image variants: `originals`, `large` (1920×1080 max), `medium` (800×600 max), `thumbnail` (150×150). Uses `large` for both single-page featured images **and** OG/social-sharing contexts ("social platforms resize anyway"). Storage shape and URL format match `docs/media.md`.

REST API response returns all four variant URLs. Featured-image validation: 1200×800 minimum. Supported formats: JPEG/PNG/GIF/WebP (with variants), PDF (originals only). Uploads served with `Cache-Control: public, max-age=604800`.

**States that the upload directory is configured via `OCMS_UPLOADS_DIR` (default `./uploads`)** — directly contradicts `docs/media.md`.

Admin features: bulk actions, folder organization, sorting (name/date/size), pagination.

## Sources

- Origin: `raw/ocms-go.core/wiki/Media-Library.md`
