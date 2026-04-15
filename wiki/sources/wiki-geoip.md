# wiki/GeoIP.md (oCMS)

## Summary

Wiki version of the GeoIP reference. Near-parity with `docs/geoip.md`: same MaxMind GeoLite2-Country setup, same `OCMS_GEOIP_DB_PATH` env var, same memory-map load-at-startup behavior, same graceful degradation ("Unknown" + warning log when DB missing), same <1 ms lookup / ~5 MB footprint, same weekly-update cadence (Tuesdays), same troubleshooting steps.

No factual disagreements. Cross-links via Obsidian wikilinks (`[[Reverse Proxy]]`, `[[Configuration]]`).

## Sources

- Origin: `raw/ocms-go.core/wiki/GeoIP.md`
