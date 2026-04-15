# Deployment

## Summary

oCMS deploys cleanly in four shapes — two targeted at demos / single hosts, two at production:

| Method | Best for |
|--------|----------|
| **Docker / Docker Compose** | Quick setup, containerized environments. |
| **Fly.io** | Public demos, small projects, free tier. |
| **Ubuntu + Plesk** | Multi-instance production with one site per vhost. |
| **Binary** | Simple single-server setups. |

All four methods share the same reverse-proxy recommendation ([topics/reverse-proxy.md](reverse-proxy.md)) and health-check surface (`/health`, `/health/live`, `/health/ready`).

## Binary build

Cross-compile targets:

```bash
make build-linux-amd64
make build-darwin-arm64
make build-all-platforms
```

The binary **embeds the core themes** (`default`, `developer`) and all static assets — no additional files are needed unless custom themes live under `custom/themes/`.

Production strip-and-trim via `make build-prod`.

## Docker

```bash
OCMS_SESSION_SECRET=$(openssl rand -base64 32) OCMS_DO_SEED=true docker compose up -d
```

Volume mounts:

| Volume | Container path | Purpose |
|--------|----------------|---------|
| `ocms_data` | `/app/data` | SQLite database. |
| `ocms_uploads` | `/app/uploads` | Media uploads. |
| `./custom` | `/app/custom` | Custom themes and modules. |

`docker compose --profile redis up -d` starts Redis alongside for distributed caching.

Build-time arguments: `VERSION` (from `git describe`), `GIT_COMMIT` (short HEAD), `BUILD_TIME`. Passed through on Fly.io builds since 0.18.0.

## Fly.io

Canonical public demo (`ocms-demo.fly.dev`).

- **VM:** shared-cpu-1x, 256 MB RAM, region `fra` (Frankfurt).
- **Volume:** 1 GB persistent at `/app/data`.
- **Auto-scale:** stops when idle, auto-starts on HTTP.
- Seeds demo content (2 users, 4 categories, 7 tags, 9 pages, 10 media, 6 menu items).
- Demo admin password is rotated on startup (since 0.18.1); credentials no longer logged.
- Reset demo: `fly ssh console -C "/app/scripts/reset-demo.sh"` or `rm /app/data/ocms.db*` + `fly machines restart`.
- OOM mitigation (0.7.0): Argon2id memory params tuned for 256 MB; `GOMEMLIMIT=200MiB` replaces `GOGC=50`.

Deploy scripts: `./.fly/scripts/deploy.sh` (standard), `./.fly/scripts/deploy.sh --reset` (reset DB).

## Ubuntu + Plesk (multi-instance)

Shared binary at `/opt/ocms/bin/ocms`, `ocmsctl` at `/opt/ocms/bin/ocmsctl`, systemd template `ocms@.service`, registry at `/etc/ocms/sites.conf`.

Per-site layout:

```
/var/www/vhosts/example.com/ocms/
├── data/ocms.db
├── uploads/
├── custom/      (optional)
├── backups/
├── .env
└── ocms.pid
```

Provision with `setup-site.sh <domain> <system-user> [<port>] [<group>]`. SSL via Plesk Let's Encrypt. Nginx snippet pasted into Plesk's **Additional nginx directives**. `systemctl enable --now ocms@<site-id>`.

### `ocmsctl` CLI

`start | stop | restart | status | logs | list | health`.

### Deploy flows

- **Binary-only**, multiple sites: `deploy-binary.sh server.example.com site_a site_b site_c`.
- **With custom content**: `deploy.sh server.example.com my_site -v /var/www/vhosts/example.com -o hosting`.
- **On-server multi-refresh**: `/opt/ocms/deploy-multi.sh /tmp/ocms`.
- **Prod → dev sync**: `sync-prod-to-dev.sh` with `--no-db`, `--no-uploads`, `--no-logs`, `--sync-custom`, `--dry-run`.

### Operations

- Backups: `backup-multi.sh [site-id]` — DB + uploads + custom, 30-day retention. Cron: `0 3 * * *`.
- Health: `healthcheck-multi.sh` auto-restarts failed instances (max 3 tries, 5-min cooldown). Cron: `*/5 * * * *`.
- Logrotate: `ocms-logrotate.conf` — daily, 30-day retention, gzip.

### File permissions

| Path | Owner | Mode |
|------|-------|------|
| `/opt/ocms/bin/ocms` | `root:root` | 755 |
| `{vhost}/ocms/` | `{user}:psaserv` | 750 |
| `{vhost}/ocms/.env` | `{user}:psaserv` | 600 |
| `{vhost}/ocms/data/` | `{user}:psaserv` | 755 |
| `{vhost}/ocms/uploads/` | `{user}:psaserv` | 755 |

### Breaking change (0.16.0)

`sites.conf` column 2 is now the **full instance directory**. Deploy scripts no longer append `/ocms` automatically. Existing deployments must migrate before updating scripts.

### Troubleshooting

Service won't start: missing `OCMS_SESSION_SECRET`, perms, port in use. `502 Bad Gateway`: `curl http://127.0.0.1:<port>/health`, `nginx -t`. Plesk duplicate `location /`: use regex location pattern. DB integrity: `PRAGMA integrity_check;`.

## Canonical source ambiguity (note, not a contradiction)

- `docs/deploy-ubuntu-plesk.md` is now a **stub** pointing to `scripts/deploy/README.md` (consolidated in 0.10.2).
- `wiki/Deploy:-Ubuntu-Plesk.md` holds the **full** guide.
- `scripts/deploy/README.md` is the canonical master per the stub (not yet ingested in Phase 1; lives under the ocms-go.core tree, not under `docs/` or `wiki/`).

No factual conflict: docs defers to the README, wiki mirrors its content.

## Sources

- [sources/wiki-deployment.md](../sources/wiki-deployment.md)
- [sources/wiki-deploy-ubuntu-plesk.md](../sources/wiki-deploy-ubuntu-plesk.md)
- [sources/docs-deploy-ubuntu-plesk.md](../sources/docs-deploy-ubuntu-plesk.md)
- [sources/wiki-deploy-fly-io.md](../sources/wiki-deploy-fly-io.md)
- [sources/wiki-docker.md](../sources/wiki-docker.md)
- [sources/readme.md](../sources/readme.md)
- [sources/changelog.md](../sources/changelog.md)
