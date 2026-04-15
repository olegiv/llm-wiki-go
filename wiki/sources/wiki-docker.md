# wiki/Docker.md (oCMS)

## Summary

Wiki Docker guide. Near-identical to the Docker section in `README.md`, with the same quick-start command, Docker command table, volume mounts table, build argument set, production `.env` example (with embed proxy + webhook allowlist + Redis section), and `docker compose --profile redis up -d` command.

### Volume mounts

| Volume | Container path | Purpose |
|--------|----------------|---------|
| `ocms_data` | `/app/data` | SQLite database. |
| `ocms_uploads` | `/app/uploads` | Media uploads. |
| `./custom` | `/app/custom` | Custom themes and modules. |

### Build arguments

`VERSION` (from `git describe`), `GIT_COMMIT` (short HEAD), `BUILD_TIME` (ISO UTC). Passed to Fly.io deploys since 0.18.0.

### Health checks

`GET /health`, `/health/live`, `/health/ready`.

## Sources

- Origin: `raw/ocms-go.core/wiki/Docker.md`
