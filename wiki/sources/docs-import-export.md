# docs/import-export.md (oCMS)

## Summary

Reference for content import / export.

### Formats

- **JSON** — single `.json` file. Best for small sites, API consumption, version control.
- **ZIP** — archive containing `export.json` + `media/` directory. Best for complete backups, server migration, sites with media.

### Export options

Content types selectable independently: Pages (with page-status filter — all / published / drafts), Categories, Tags, Media (ZIP includes files), Menus, Forms (optionally with submissions), Users (emails only, no passwords), Languages, Site Config.

### Export paths

Admin UI at `/admin/export` (under **Admin > Config > Export**). API: `POST /admin/export` with `include_*=true` and `format=json|zip` form fields.

### Export schema (top-level)

```json
{
  "version": "1.0",
  "exported_at": "2024-01-15T10:30:00Z",
  "site": {...},
  "languages": [...],
  "users": [...],
  "categories": [...],
  "tags": [...],
  "pages": [...],
  "media": [...],
  "menus": [...],
  "forms": [...],
  "config": {...}
}
```

### Page export shape

Includes `id`, `title`, `slug`, `body`, `status`, `author_email`, `categories` (by slug), `tags` (by slug), `seo` object (`meta_title`, `meta_description`, `meta_keywords`, `og_image`, `no_index`, `no_follow`, `canonical_url`), `translations` (map of `lang_code → page_id`), `language_code`, **`video_url`**, **`video_title`**, and timestamps.

The `video_url` and `video_title` fields were added in 0.16.0; they appear in this docs file but are **absent from `wiki/Import-Export.md`** — see `topics/import-export.md` for the reconciliation.

### Media export fields

`uuid`, `filename`, `original_name`, `mime_type`, `size`, `path`, `alt_text`, `caption`.

Docs lists `path`; the wiki version omits it.

### Import flow

1. Upload file at `/admin/import`.
2. **Validate** for preview (valid JSON, required fields, reference integrity, file compatibility).
3. Set conflict strategy (**Skip**, **Overwrite**, **Rename** with `-1`, `-2` suffix).
4. Execute import (or use **Dry Run**).
5. Review summary.

### Import order

Languages → Users → Categories → Tags → Pages → Media → Menus → Forms. Preserves relational integrity.

### Validation output

Example structure: counts (pages / categories / tags / media), warnings (already-exists entries), errors (missing references).

### Import summary

Created / Updated / Skipped / Errors counts per entity type.

### API import

`POST /admin/import` with `file=@export.json` and `conflict_strategy=skip|overwrite|rename`.

### Migration guide (server to server)

1. ZIP export from source.
2. Fresh install + migrations + admin user on target.
3. Upload ZIP, validate, conflict = Overwrite, import.
4. Verify pages, media, translations.

### Backup strategies

1. SQLite direct file copy (via `sqlite3 .backup`).
2. Uploads directory tar.
3. Periodic API export.

### Troubleshooting

- Invalid JSON — re-export from source.
- Reference not found — respect import order.
- Media missing — use ZIP instead of JSON.
- Large imports — split by content type or increase timeout.

### CSV formula injection

Form submission CSV export was hardened against CSV formula injection in 0.18.1.

### Schema reference

The doc includes a full schema reference covering languages, users, categories, tags, pages (with nested `seo` and `translations` objects), media, menus (with nested `items`), forms (with nested `fields` and `submissions`), and `config`.

## Sources

- Origin: `raw/ocms-go.core/docs/import-export.md`
