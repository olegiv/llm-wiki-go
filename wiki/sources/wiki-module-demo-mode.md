# wiki/Module:-Demo-Mode.md (oCMS)

## Summary

Wiki version of the demo mode reference. Matches `docs/demo-mode.md`: **Create/Edit of pages is Allowed**, only Delete is Blocked. Same user-experience behavior (flash messages on web UI; HTTP 403 on API). Same 2 MB upload cap. Same security scope.

Differences from `docs/demo-mode.md`:

- Omits the implementation section (no middleware file paths or `demoGuard` helper code).
- Omits the "Adding New Restrictions" section.
- Adds a **Demo credentials** table (demo@example.com / editor@example.com) directly in the reference.
- Adds a **Scheduler Restrictions** note — demo mode also disables scheduler edits, resets, and manual triggers (confirming the interaction from [sources/docs-scheduler.md](docs-scheduler.md)).
- Wikilinks `[[Deploy: Fly.io]]` and `[[Scheduler]]`.

Inherits the **Create/Edit Allowed** stance, which disagrees with `docs/demo-deployment.md`'s restrictions table. See [topics/demo-mode.md](../topics/demo-mode.md) for reconciliation.

## Sources

- Origin: `raw/ocms-go.core/wiki/Module:-Demo-Mode.md`
