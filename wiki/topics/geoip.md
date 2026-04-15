# GeoIP

## Summary

The internal analytics module resolves visitor countries via **MaxMind GeoLite2-Country** (MMDB). Enable by setting `OCMS_GEOIP_DB_PATH` to the `.mmdb` file. Country data appears in the analytics dashboard's country breakdown, geo-distribution charts, and filtering controls.

## Setup

1. Create a free MaxMind account and download `GeoLite2-Country.mmdb`.
2. Place it somewhere reachable by the oCMS process.
3. Set:
   ```bash
   export OCMS_GEOIP_DB_PATH=/path/to/GeoLite2-Country.mmdb
   ```
   Or in a systemd unit:
   ```ini
   [Service]
   Environment="OCMS_GEOIP_DB_PATH=/opt/geoip/GeoLite2-Country.mmdb"
   ```
4. Restart oCMS.

## Runtime behavior

- Memory-mapped load at startup.
- Lookups are thread-safe and <1 ms.
- Memory footprint roughly 5 MB for GeoLite2-Country.
- File changes are auto-detected and reloaded within ~1 hour. Restart for immediate reload.

## Update cadence

MaxMind publishes GeoLite2 updates weekly (Tuesdays). Replace the file or restart to pick up new data.

## Graceful degradation

If the DB is missing or unreadable:

- Analytics continues to run.
- Country detection silently disables.
- Dashboards report "Unknown" for all visitors.
- A startup warning is logged.

## Troubleshooting

- **No country data:** confirm env var, confirm file is readable, check startup log for `GeoIP database loaded` or `GeoIP not configured`.
- **All "Unknown":** ensure Country variant (not City or ASN); verify non-private IPs; check file integrity.
- **All "Local Network":** traffic from private ranges (10/8, 192.168/16, 172.16/12) — configure reverse-proxy forwarding (`X-Real-IP`, `X-Forwarded-For`) and `OCMS_TRUSTED_PROXIES` so the trusted-proxy-aware resolution (since 0.17.0) accepts them.

## Related entities and topics

- Country inference feeds the analytics module — see Module System and the individual analytics modules to be ingested in later batches.
- Reverse-proxy setup for correct client IP is a prerequisite when oCMS runs behind nginx/Apache.

## Sources

- [sources/docs-geoip.md](../sources/docs-geoip.md)
- [sources/wiki-geoip.md](../sources/wiki-geoip.md)
- [sources/readme.md](../sources/readme.md)
