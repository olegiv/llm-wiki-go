# Import / export

## Summary

oCMS supports full-site content export and import via JSON (single file) or ZIP (JSON + media). Selective export (per content type), conflict resolution (skip / overwrite / rename), and dry-run validation are all first-class.

## Admin surface

- **Admin > Config > Export** (`/admin/export`) — select content types → JSON or ZIP → download.
- **Admin > Config > Import** (`/admin/import`) — upload → validate → configure → import → review summary.

## Exportable content types

Pages (with status filter: all / published / drafts), Categories, Tags, Media (ZIP includes files; JSON is metadata-only), Menus, Forms (optionally with submissions), Users (emails only; no passwords), Languages, Site Config.

## Format comparison

| Format | Contents | Best for |
|--------|----------|----------|
| JSON | Content data only | Small sites, API consumption, version control. |
| ZIP | `export.json` + `media/` | Complete backups, server migration, sites with media. |

## Export schema (top-level)

```json
{
  "version": "1.0",
  "exported_at": "2024-01-15T10:30:00Z",
  "site": {"name", "base_url", "tagline"},
  "languages": [],
  "users": [],
  "categories": [],
  "tags": [],
  "pages": [],
  "media": [],
  "menus": [],
  "forms": [],
  "config": {}
}
```

### Page export (from `docs/import-export.md`)

Fields: `id`, `title`, `slug`, `body`, `status`, `author_email`, `categories[]`, `tags[]`, nested `seo` object, `translations` map, `language_code`, **`video_url`**, **`video_title`** (both since 0.16.0), timestamps.

### Media export (from `docs/import-export.md`)

Fields: `uuid`, `filename`, `original_name`, `mime_type`, `size`, `path`, `alt_text`, `caption`.

## Import flow

1. Upload file.
2. **Validate** — checks JSON shape, required fields, reference integrity, file compatibility. Emits found counts, warnings (duplicates), errors (missing references).
3. Choose conflict strategy:
   | Strategy | Behavior |
   |----------|----------|
   | Skip | Keep existing, ignore imported. |
   | Overwrite | Replace existing with imported. |
   | Rename | Import with suffix (`-1`, `-2`, …). |
4. Optional **Dry Run** — preview without writing.
5. Execute.
6. Summary: Created / Updated / Skipped / Errors per entity type.

## Import order

Enforced to preserve relational integrity:

1. Languages
2. Users
3. Categories
4. Tags
5. Pages
6. Media
7. Menus
8. Forms

## API access

```bash
# Export
curl -X POST http://localhost:8080/admin/export \
  -H "Cookie: session=your-session-cookie" \
  -d "include_pages=true&include_media=true&format=json" \
  -o export.json

# Import
curl -X POST http://localhost:8080/admin/import \
  -H "Cookie: session=your-session-cookie" \
  -F "file=@export.json" \
  -F "conflict_strategy=skip"
```

## Security

- CSV export of form submissions hardened against **formula injection** (since 0.18.1).
- Reported inline alongside uploads subsystem MIME validation and crypto-random passwords for imported users.

## Demo mode

Import and export are **both blocked** in demo mode (see [topics/demo-mode.md](demo-mode.md)).

## Contradictions

### Page schema — video fields

- [sources/docs-import-export.md](../sources/docs-import-export.md) — page export includes `video_url` and `video_title` (added 0.16.0 per `CHANGELOG.md`).
- [sources/wiki-import-export.md](../sources/wiki-import-export.md) — page export omits both fields.

Clear pre-0.16.0 drift in the wiki version.

### Media schema — `path` field

- [sources/docs-import-export.md](../sources/docs-import-export.md) — media fields include `path`.
- [sources/wiki-import-export.md](../sources/wiki-import-export.md) — media fields omit `path`.

Same direction of drift; the wiki has not been refreshed.

## Sources

- [sources/docs-import-export.md](../sources/docs-import-export.md)
- [sources/wiki-import-export.md](../sources/wiki-import-export.md)
- [sources/changelog.md](../sources/changelog.md)
