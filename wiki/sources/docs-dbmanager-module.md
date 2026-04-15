# docs/dbmanager-module.md (oCMS)

## Summary

Reference for the DB Manager module — a built-in admin-only SQL runner.

### Features

SQL execution (SELECT, INSERT, UPDATE, DELETE, PRAGMA, EXPLAIN, WITH), tabular results for SELECT, row-count / duration metrics for DML, query history with click-to-reload, `Ctrl+Enter` shortcut.

### Registration

Registered in `cmd/ocms/main.go` via `moduleRegistry.Register(dbmanager.New())`.

### Admin UI

At `/admin/dbmanager`. Enter SQL → **Execute Query** or Ctrl+Enter → results table or affected-row count.

### Query classification

Prefix-based: `SELECT` / `PRAGMA` / `EXPLAIN` / `WITH` → "read" queries returning rows; everything else → "write" with affected-row count.

### History table

`dbmanager_query_history (id, query, user_id, executed_at, rows_affected, execution_time_ms, error)`.

Logged per execution for audit.

### Routes

- `GET /admin/dbmanager` — dashboard.
- `POST /admin/dbmanager/execute` — execute query.

### Security

Admin-auth middleware + CSRF on POST + full query history for audit + active-status toggle. Demo mode blocks SQL execution entirely.

### i18n

Translation keys prefixed `dbmanager.*`; English + Russian.

### Module structure

```
modules/dbmanager/
├── module.go
├── handlers.go
├── dbmanager_test.go
└── locales/{en,ru}/messages.json
```

### Best practices

Backup before modifications; test with SELECT first; use transactions for complex changes; review history.

## Sources

- Origin: `raw/ocms-go.core/docs/dbmanager-module.md`
