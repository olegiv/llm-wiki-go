# docs/scheduler.md (oCMS)

## Summary

Reference for the cron scheduler (`internal/scheduler/`).

### Two job types

- **Core jobs** — built-in (e.g. publish scheduled pages).
- **Module jobs** — registered by modules via `ModuleCron() (string, func())` lifecycle hook.

Each job carries default schedule, optional override (persisted to `scheduler_overrides` table), last/next run timing, and optional manual-trigger capability.

### Registration

Core scheduler owns a single `*cron.Cron` instance. Startup:

```go
entryID, _ := cronInst.AddFunc(effectiveSchedule, jobFunc)
registry.Register("core", "publish_scheduled", "Publish scheduled pages", "0 * * * *", cronInst, entryID, jobFunc, triggerFunc)
```

Modules:

```go
func (m *Module) ModuleCron() (string, func()) {
    return "0 1 * * *", func() { /* work */ }
}
```

Core collects module cron funcs, loads effective schedules from DB, registers in the scheduler registry, and manages execution.

### Admin UI

`/admin/scheduler` — table with columns Job, Source (blue badge = core, cyan = module), Schedule, Last Run, Next Run, Actions (Edit / Reset / Trigger Now).

### Edit schedule

Edit button → modal → new cron expression → save. System validates the expression first; failure restores previous schedule.

### Cron format

Standard 5-field (minute 0–59, hour 0–23, DoM 1–31, month 1–12, DoW 0–6 with 0 = Sunday) + `@every <duration>` shorthand (`@every 1h`, `@every 30m`, etc.).

### Manual triggers

**Trigger Now** runs the job immediately, logs the trigger event (with user), does NOT reset schedule/timing counters. Jobs can opt out of manual triggers (e.g. those requiring specific timing conditions).

### Reset

When an override is active, **Reset** removes it from the DB, restores default schedule, and hides the button.

### Schedule overrides table

`CREATE TABLE scheduler_overrides (source TEXT, name TEXT, override_schedule TEXT, updated_at DATETIME DEFAULT CURRENT_TIMESTAMP, PRIMARY KEY (source, name))`. At most one override per job. Created by goose migration `20260213100000_create_scheduler_overrides.sql`; a runtime `CREATE TABLE IF NOT EXISTS` remains as a safety net. On startup, overrides are loaded and applied.

### Demo mode restrictions

`OCMS_DEMO_MODE=true` disables edit/reset/trigger in the scheduler UI; viewing still works. Attempts show error and redirect.

### Audit logging

Three event types at `/admin/events`: Schedule updated, Schedule reset, Job manually triggered. Each event includes user, job source+name, new schedule (where applicable), client IP, request URL.

### Use cases

Change publish-scheduled-pages from `0 * * * *` to `*/5 * * * *` for near-real-time. Disable by setting to `0 0 31 2 *` (Feb 31 never). Shift timing around backup windows.

### Troubleshooting

Job never runs: validate cron expression, check module active flag, verify app is running, check logs. Manual trigger error: job may not support triggers, check logs, verify module loaded. Override not persisting: DB write permissions, restart.

## Sources

- Origin: `raw/ocms-go.core/docs/scheduler.md`
