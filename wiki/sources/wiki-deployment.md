# wiki/Deployment.md (oCMS)

## Summary

Overview of deployment methods with pointers into detail pages.

### Method matrix

| Method | Best for |
|--------|----------|
| Docker | Quick setup, containerized environments. |
| Fly.io | Public demos, small projects, free tier. |
| Ubuntu + Plesk | Multi-instance production. |
| Binary | Simple single-server setups. |

### Binary build

`make build-linux-amd64`, `make build-darwin-arm64`, or `make build-all-platforms`. The binary **embeds the core themes** (`default`, `developer`) and all static assets — no extra files needed unless using custom themes.

### Reverse proxy

Always run behind a reverse proxy in production for SSL, caching, and security. Nginx / Apache / Nginx Proxy Manager configs live on the Reverse Proxy page.

### Health checks

Monitor `/health`, `/health/live`, `/health/ready`.

### Production essentials

```
OCMS_SESSION_SECRET=your-secure-secret-key-at-least-32-bytes
OCMS_ENV=production
OCMS_TRUSTED_PROXIES=127.0.0.1/32
```

### Database

SQLite by default at `OCMS_DB_PATH` (default `./data/ocms.db`). Backup with `sqlite3 … .backup '…'` and `tar -czf uploads-backup.tar.gz uploads/`.

### Custom themes / modules layout

```
/path/to/deployment/
├── ocms
├── data/
├── uploads/
└── custom/
    ├── themes/
    └── modules/
```

Set `OCMS_CUSTOM_DIR=./custom` (default).

## Sources

- Origin: `raw/ocms-go.core/wiki/Deployment.md`
