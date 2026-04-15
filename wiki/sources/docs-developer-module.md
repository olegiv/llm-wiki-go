# docs/developer-module.md (oCMS)

## Summary

Reference for the Developer module — a test-data generator for development.

### Scope

Generates 5–20 tags, 5–20 categories (nested ~40% root / ~40% children / ~20% grandchildren), 5–20 placeholder images, 5–20 pages with tags/categories/featured images, 5–10 menu items in Main Menu. Every item is translated into each active language (suffixed with `(<lang>)`). Everything is tracked in a `developer_generated_items` table so bulk deletion removes only module-created content.

### Admin UI

At `/admin/developer` (a.k.a. **Admin > Modules > Developer Tools**). Dashboard shows per-type counts; **Generate Test Data** creates the batch; **Delete All Generated Data** reverses it.

### Deletion order

Menu items → translation records → pages + associations (tags/categories) → media (disk + DB) → categories + tags → tracking table. Main Menu (ID=1) itself is preserved; only its generated items are removed.

### Image generation

Standard library `image` + `image/jpeg`. 800×600 original at 85% JPEG. 10 predefined colors. Variants written: `thumbnail` (150×150), `medium` (800×600), `large` (1920×1080). Stored at `./uploads/{variant}/{UUID}/placeholder-N.jpg`.

Note: the variant list here is **incomplete** compared with [sources/docs-media.md](docs-media.md), which lists seven variants (originals, large, og, medium, small, grid, thumbnail). The developer module apparently does not produce `og`, `small`, or `grid` variants for its placeholders. This is a scope limitation of the test-data generator, not a documentation contradiction.

### Routes

- `GET /admin/developer` — dashboard.
- `POST /admin/developer/generate` — generate test data.
- `POST /admin/developer/delete` — delete generated data.

### Word lists

25 adjectives (Amazing, Beautiful, …, Zesty) + 25 nouns (Technology, Science, …, Services) used to compose tag/category/page titles.

### Security and scope

Admin auth + CSRF + active-status toggle. **For development environments only** — production should keep this module inactive. Admin-only restriction hardened in 0.18.1.

### Module structure

```
modules/developer/
├── module.go
├── handlers.go
├── generator.go
└── locales/{en,ru}/messages.json
```

### i18n keys

Prefixed `developer.*`: `title`, `description`, `warning`, `stats.*`, `btn.*`, `confirm.*`, `flash.*`.

### Tracking table schema

`developer_generated_items (id, entity_type, entity_id, created_at)` with `entity_type` ∈ `{tag, category, media, page, menu_item, translation}`.

## Sources

- Origin: `raw/ocms-go.core/docs/developer-module.md`
