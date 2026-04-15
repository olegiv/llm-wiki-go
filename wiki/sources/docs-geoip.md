# docs/geoip.md (oCMS)

## Summary

GeoIP-based country detection for the internal analytics module, using the MaxMind **GeoLite2-Country** database (MMDB format).

### Setup

1. Create a free MaxMind account and download `GeoLite2-Country.mmdb`.
2. Set `OCMS_GEOIP_DB_PATH=/path/to/GeoLite2-Country.mmdb` (env var or systemd unit).
3. Restart oCMS — the DB is loaded once at startup.

### Runtime behavior

- Memory-mapped load; lookups are thread-safe and <1 ms.
- Memory footprint ~5 MB for GeoLite2-Country.
- DB file changes are auto-detected and reloaded within an hour; restart for immediate reload.

### Graceful degradation

If the DB is missing or unreadable, analytics continues to run with country detection silently disabled, all visitors report "Unknown", and a warning is logged at startup.

### Troubleshooting

- "Country data not appearing": verify env var is set, file exists, startup logs show one of "GeoIP database loaded" or "GeoIP not configured".
- All "Unknown": confirm you downloaded the Country DB (not City or ASN); check for non-private IP ranges; verify file not corrupted.
- All "Local Network": private IP ranges — configure reverse-proxy headers (`X-Real-IP`, `X-Forwarded-For`) and `OCMS_TRUSTED_PROXIES`.

### Weekly updates

MaxMind updates GeoLite2 databases on Tuesdays. Replace the file or restart to pick up new data.

## Sources

- Origin: `raw/ocms-go.core/docs/geoip.md`
