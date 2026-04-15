# wiki/Import-Export.md (oCMS)

## Summary

Wiki version of the import/export reference. Matches `docs/import-export.md` on most facts: same JSON / ZIP formats, same 10 export options, same `Skip` / `Overwrite` / `Rename` conflict strategies, same 8-step import order (Languages → Users → Categories → Tags → Pages → Media → Menus → Forms), same validation output shape, same API examples, same migration walkthrough.

### Differences from `docs/import-export.md` (drift — predates 0.16.0)

- **Page export shape omits `video_url` and `video_title`.** Docs version includes these fields (added in 0.16.0 per `CHANGELOG.md`).
- **Media schema omits `path`.** Docs version includes it.

Both omissions look like drift from before the 0.16.0 / video-embedding release.

### Other differences (not contradictions)

- Omits the "Export Schema" intermediate subsection in favor of the one consolidated "Full Export Schema Reference" at the bottom.
- Omits the manual-backup bash snippet.
- Omits the CSV-formula-injection / CSV-export hardening note.

## Sources

- Origin: `raw/ocms-go.core/wiki/Import-Export.md`
