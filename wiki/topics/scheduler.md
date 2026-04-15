# Scheduler

## Summary

The `internal/scheduler/` package provides a single in-process cron scheduler that runs both core jobs and module-contributed jobs. The admin UI at `/admin/scheduler` lists jobs, exposes edit / reset / trigger actions, and records schedule changes and manual triggers in the audit log.

## Job sources

- **Core** — built-in system jobs (e.g. `core.publish_scheduled` at `0 * * * *`).
- **Module** — registered via the `ModuleCron() (string, func())` lifecycle hook.

Each job carries:

- Default cron schedule.
- Optional override (persisted to the `scheduler_overrides` table).
- Timing metadata: last run, next run.
- Optional manual-trigger capability.

## Registration

Core (from [sources/docs-scheduler.md](../sources/docs-scheduler.md)):

```go
entryID, _ := cronInst.AddFunc(effectiveSchedule, jobFunc)
registry.Register("core", "publish_scheduled", "Publish scheduled pages", "0 * * * *",
                  cronInst, entryID, jobFunc, triggerFunc)
```

Module:

```go
func (m *Module) ModuleCron() (string, func()) {
    return "0 1 * * *", func() { /* work */ }
}
```

On startup, the core system collects cron funcs from active modules, loads effective schedules from the database, and registers each job in the scheduler registry.

## Cron syntax

5-field (minute, hour, day-of-month, month, day-of-week; 0 = Sunday) + `@every <duration>` shorthand.

| Expression | Meaning |
|------------|---------|
| `*/5 * * * *` | Every 5 minutes |
| `0 * * * *` | Every hour at :00 |
| `0 1 * * *` | Daily at 01:00 |
| `0 1 * * 0` | Weekly (Sundays 01:00) |
| `0 0 1 * *` | Monthly (1st of month) |
| `@every 30m` | Every 30 minutes |
| `@every 6h` | Every 6 hours |

## Overrides

Editing a schedule persists a row to `scheduler_overrides` (`(source, name)` primary key). On startup, overrides are loaded and applied. **Reset** removes the override and restores the default. Overrides survive restarts; migration `20260213100000_create_scheduler_overrides.sql` creates the table, with a runtime `CREATE TABLE IF NOT EXISTS` as a safety net.

## Manual triggers

- **Trigger Now** runs the job immediately, logs the trigger (with user), does **not** reset schedule timing.
- Not all jobs support manual triggers — those requiring specific timing conditions (e.g. "publish at exact moment") may opt out.

## Demo-mode restrictions

When `OCMS_DEMO_MODE=true`, the scheduler UI reverts to read-only: no edit, reset, or trigger actions. Attempts produce an error and redirect. This is part of the broader demo-mode hardening introduced in 0.7.0.

## Audit logging

Three event types land in `/admin/events`:

- Schedule updated.
- Schedule reset.
- Job manually triggered.

Each records: user, job source + name, new schedule (where applicable), client IP, request URL.

## Disabling a job

Set its schedule to an impossible expression such as `0 0 31 2 *` (February 31 never). Or coordinate with the module maintainer to remove the module.

## Sources

- [sources/docs-scheduler.md](../sources/docs-scheduler.md)
- [sources/wiki-scheduler.md](../sources/wiki-scheduler.md)
- [sources/readme.md](../sources/readme.md)
- [sources/changelog.md](../sources/changelog.md)
