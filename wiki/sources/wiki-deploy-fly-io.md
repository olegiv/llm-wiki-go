# wiki/Deploy:-Fly.io.md (oCMS)

## Summary

Deployment guide for Fly.io — the canonical home for the public demo instance.

### Prerequisites

- flyctl CLI.
- Fly.io account (free tier is sufficient).

### Quick start

```bash
fly auth login
fly launch --no-deploy --copy-config
fly volumes create ocms_data --size 1 --region fra
fly secrets set OCMS_SESSION_SECRET="$(openssl rand -base64 32)" --app ocms-demo
fly deploy
```

### `fly.toml` highlights

- Region: `fra` (Frankfurt; customizable).
- VM: shared-cpu-1x, **256 MB RAM**.
- Volume: 1 GB persistent at `/app/data`.
- Health checks: `/health/live` (liveness) and `/health/ready` (readiness).
- Auto-scaling: stops when idle, auto-starts on HTTP request.

Environment variables baked into `fly.toml`:

```
OCMS_ENV=production
OCMS_DO_SEED=true
OCMS_DEMO_MODE=true
OCMS_DB_PATH=/app/data/ocms.db
OCMS_UPLOADS_DIR=/app/data/uploads
```

Secrets (set via `fly secrets set`): `OCMS_SESSION_SECRET` (required).

### Demo mode content (on fresh deploy)

| Content | Count |
|---------|-------|
| Users | 2 (admin + editor) |
| Categories | 4 |
| Tags | 7 |
| Pages | 9 |
| Media | 10 placeholder images |
| Menu items | 6 |

### Demo credentials

`demo@example.com` / `demo1234demo` (admin) and `editor@example.com` / `demo1234demo` (editor).

Note: since 0.18.1, the demo admin password is rotated on startup; a stale hard-coded default was removed from the seeded homepage.

### Demo reset

```bash
fly ssh console -C "/app/scripts/reset-demo.sh"
fly machines restart -a ocms-demo
# or
fly ssh console -C "rm /app/data/ocms.db*" -a ocms-demo
fly machines restart -a ocms-demo
```

### Deploy scripts

`./.fly/scripts/deploy.sh` (standard) and `./.fly/scripts/deploy.sh --reset` (reset DB).

### Monitoring

`fly logs`, `curl https://ocms-demo.fly.dev/health`, `fly ssh console`.

### Cost (free tier)

| Resource | Usage | Free tier |
|----------|-------|-----------|
| VM | shared-cpu-1x, 256 MB | 3 VMs |
| Volume | 1 GB | 3 GB |
| Bandwidth | ~10–50 GB/mo | 160 GB/mo |

### OOM mitigation

On 256 MB Fly.io VMs, 0.7.0 fixed an OOM crash at login by tuning Argon2id memory parameters, and replaced `GOGC=50` with `GOMEMLIMIT=200MiB` in the Fly config.

## Sources

- Origin: `raw/ocms-go.core/wiki/Deploy:-Fly.io.md`
