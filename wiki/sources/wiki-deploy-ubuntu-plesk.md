# wiki/Deploy:-Ubuntu-Plesk.md (oCMS)

## Summary

Full deployment guide for multi-instance oCMS on Ubuntu with Plesk (one instance per domain/vhost). Unlike `docs/deploy-ubuntu-plesk.md` (a stub pointing at `scripts/deploy/README.md`), this wiki page holds the real content.

### Architecture

Shared binary at `/opt/ocms/bin/ocms` (embeds core themes). Each site runs on its own port behind Nginx reverse proxy. Per-site layout:

```
/var/www/vhosts/example.com/ocms/
├── data/ocms.db
├── uploads/
├── custom/      (optional)
├── backups/
├── .env
└── ocms.pid
```

Site registry at `/etc/ocms/sites.conf`. Systemd template at `/etc/systemd/system/ocms@.service`.

### Six-step quick-start

1. Build binary (`make build-linux-amd64`), `scp` to server, install deploy scripts (`setup-site.sh`, `deploy-multi.sh`, `backup-multi.sh`, `healthcheck-multi.sh`, `helper.sh`), `ocmsctl`, and `ocms@.service`. `systemctl daemon-reload`.
2. `setup-site.sh <domain> <system-user> [<port>] [<group>]` — generates directories, session secret, systemd drop-in, Nginx snippet.
3. Plesk SSL: Let's Encrypt via Plesk's Websites & Domains panel. Enable 301 HTTP→HTTPS.
4. Paste Nginx snippet into Plesk's **Additional nginx directives**.
5. `systemctl enable --now ocms@<site-id>`.
6. Login at `/admin/` with `admin@example.com` / `changeme1234`; rotate immediately.

### `ocmsctl` CLI

`start`, `stop`, `restart`, `status`, `logs`, `list`, `health`.

### Deploy flows

- **Binary-only, multi-instance:** `deploy-binary.sh server.example.com site_a site_b site_c` with `-u USER`, `--skip-build`, `--dry-run`.
- **With custom content:** `deploy.sh server.example.com my_site -v /var/www/vhosts/example.com -o hosting` with `--skip-binary` option.
- **On-server multi-instance refresh:** `/opt/ocms/deploy-multi.sh /tmp/ocms`.
- **Prod → dev sync:** `sync-prod-to-dev.sh` with `--no-db`, `--no-uploads`, `--no-logs`, `--sync-custom`, `--dry-run`.

### Backups

`backup-multi.sh [site-id]`: includes DB, uploads, custom content. Stored per-site under `backups/` with 30-day retention.

### Health checks

`healthcheck-multi.sh` auto-restarts failed systemd instances (max 3 attempts, 5-min cooldown). Slack / email alerting optional.

### Cron snippets

```
0 3 * * * root /opt/ocms/backup-multi.sh >> /var/log/ocms-backup.log 2>&1
*/5 * * * * root /opt/ocms/healthcheck-multi.sh 2>&1 | grep -v "^$"
```

### Logrotate

`ocms-logrotate.conf`: daily, 30 days, gzip compression.

### File permissions table

Ocms binary `root:root 755`. Per-site `{user}:psaserv 750`. `.env` `600`. Data and uploads `755`.

### Sites registry

`# SITE_ID VHOST_PATH SYSTEM_USER PORT` lines. Comment a line to disable without removing data.

### Removing a site

`systemctl disable --now`, remove systemd drop-in, `daemon-reload`, `rm -rf {vhost}/ocms`, edit `sites.conf`.

### Troubleshooting

Service start: missing session secret, perms, port in use. `502 Bad Gateway`: check `127.0.0.1:8081/health`, `nginx -t`. Plesk duplicate `location /`: use regex `location ~ ^/(.*)$`. DB corruption: `sqlite3 … PRAGMA integrity_check`.

### Breaking change in 0.16.0

`sites.conf` column 2 is now the **full instance directory** (not just the vhost path). Deploy scripts no longer append `/ocms` automatically. Old deployments must migrate before updating scripts.

## Sources

- Origin: `raw/ocms-go.core/wiki/Deploy:-Ubuntu-Plesk.md`
