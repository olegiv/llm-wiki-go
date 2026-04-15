# docs/demo-deployment.md (oCMS)

## Summary

Combined guide for demo mode + Fly.io deployment.

### Demo mode

Requires **both** `OCMS_DO_SEED=true` and `OCMS_DEMO_MODE=true`. Demo seeding is gated on the DO_SEED flag.

### Seeded demo content

| Content | Count | Details |
|---------|-------|---------|
| Users | 2 | Admin + editor. |
| Categories | 4 | Blog, Portfolio, Services, Resources. |
| Tags | 7 | Tutorial, News, Featured, Go, Web Development, Design, Open Source. |
| Pages | 9 | Home, About, Contact, 3 blog posts, 2 portfolio items, 1 services page. |
| Media | 10 | 2400×1600 placeholder images with all variants. |
| Menu items | 6 | Home, Blog, Portfolio, Services, About, Contact. |

### Credentials

- Admin: `demo@example.com` / `demo1234demo`.
- Editor: `editor@example.com` / `demo1234demo`.
- Default admin: `admin@example.com` / `changeme1234`.

### Restrictions table

Content (pages, tags, categories, menus, media, forms, widgets): View only — **Create, edit, delete, unpublish all blocked**. Themes: view + switch, but theme settings cannot be edited. Users / site config / languages / API keys / webhooks / modules: view only. Import/export: disabled. Cache: view stats only. Media uploads: all blocked. DB Manager SQL: blocked. Form submissions: CSV export blocked.

**Note:** This restriction table disagrees with `docs/demo-mode.md` and `wiki/Module:-Demo-Mode.md`, both of which say Create/Edit of pages is **Allowed**. See [topics/demo-mode.md](../topics/demo-mode.md) for the reconciliation.

### Idempotent seeding

Seed functions check for existing data before inserting. Re-seed by deleting the database (`rm ./data/ocms.db*`) and restarting.

### Uploads directory

`OCMS_UPLOADS_DIR` configures the media directory (default `./uploads`).

### Fly.io deployment

Prerequisites: flyctl, Fly.io account (free tier).

Quick start: `fly auth login`, `fly launch --no-deploy --copy-config`, `fly volumes create ocms_data --size 1 --region fra`, set `OCMS_SESSION_SECRET` via `fly secrets set`, `fly deploy`.

`fly.toml`: region `fra`, shared-cpu-1x 256 MB, 1 GB volume at `/app/data`, health probes, auto-scale stops-when-idle.

Environment: `OCMS_ENV=production`, `OCMS_DO_SEED=true`, `OCMS_DEMO_MODE=true`, `OCMS_DB_PATH=/app/data/ocms.db`, `OCMS_UPLOADS_DIR=/app/data/uploads`.

Secrets: `OCMS_SESSION_SECRET` (required).

### Deploy scripts

- `./.fly/scripts/deploy.sh` — standard.
- `./.fly/scripts/deploy.sh --reset` — reset DB.

### Demo reset

Reset script path: `fly ssh console -C "/app/scripts/reset-demo.sh"` + `fly machines restart -a ocms-demo`. Or `rm /app/data/ocms.db*` + restart.

### Monitoring

`fly logs`, `curl https://ocms-demo.fly.dev/health`, `fly ssh console`.

### Cost

Free tier: 3 VMs, 3 GB volume, 160 GB bandwidth. Typical usage 10–50 GB/month.

### Docker alternative

Multi-stage Dockerfile; `docker run -p 8080:8080 -e OCMS_SESSION_SECRET=… -e OCMS_DO_SEED=true -e OCMS_DEMO_MODE=true -v ocms_data:/app/data ocms-demo`.

### Files

```
.fly/README.md
.fly/scripts/deploy.sh
.fly/scripts/reset-demo.sh
fly.toml
Dockerfile
internal/store/seed_demo.go
```

## Sources

- Origin: `raw/ocms-go.core/docs/demo-deployment.md`
