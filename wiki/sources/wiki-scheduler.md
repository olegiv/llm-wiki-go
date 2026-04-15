# wiki/Scheduler.md (oCMS)

## Summary

Wiki version of the scheduler reference. Near-parity with `docs/scheduler.md`: same two job types (core + module), same job metadata (default schedule, override, last/next run, manual-trigger flag), same `/admin/scheduler` UI (same table columns and color-coded badges), same edit workflow with cron validation, same cron syntax support (5-field + `@every`), same manual-trigger semantics, same Reset behavior, same demo-mode restrictions, same three audit event types.

Differences from `docs/scheduler.md`:

- Omits the Go code snippets for `cronInst.AddFunc(...)` registration and `ModuleCron() (string, func())` module hook.
- Omits the `scheduler_overrides` SQL DDL and goose migration filename.
- Demo-mode reference uses wikilink `[[Module: Demo Mode]]` instead of describing `OCMS_DEMO_MODE`.

No factual disagreements.

## Sources

- Origin: `raw/ocms-go.core/wiki/Scheduler.md`
