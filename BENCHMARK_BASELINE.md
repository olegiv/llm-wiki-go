# oCMS feature ship — baseline (no wiki)

Claude Code prompt for the **control arm** of a wiki-vs-no-wiki experiment.
Treatment arm lives in `./BENCHMARK_TREATMENT.md`.

Run headless via the `benchmark.sh` harness. Wall-clock, tokens, and cost
are captured by the harness — this prompt does not self-measure.

---

You are measuring how Claude Code performs on a realistic feature-ship
workflow **without** the LLM Wiki.

## Hard constraints

- Read only files inside the current `ocms-go` repository.
- Do **NOT** read, `Grep`, `Glob`, `cat`, `ls`, or otherwise access anything
  under `../llm-wiki-go/`. If any tool output contains a `llm-wiki-go` path,
  stop and report the contamination instead of continuing.
- Do **NOT** fetch the web.
- You are already on a clean benchmark branch prepared by the harness. Work
  on that branch; do not switch branches.
- Follow every rule in this repo's `CLAUDE.md` and `AGENTS.md` (license
  headers, error wrapping, slog fields, naming, validation functions,
  resource cleanup, etc.). These are not optional — deviations count
  against correctness.

## The feature

Add a **per-page cache TTL override**. Global cache TTL (`OCMS_CACHE_TTL`)
is unchanged; each page may optionally override it.

### Acceptance criteria

1. **Schema.** New goose migration adds a nullable `cache_ttl_seconds`
   INTEGER column to the `pages` table. `down` removes it cleanly.
2. **Semantics.**
   - `NULL` → use global TTL (current behavior, no change for existing rows).
   - `0` → never cache this page (cache middleware MUST bypass storage).
   - `1..2592000` (30 days) → override TTL in seconds.
   - Anything else → validation error at write time.
3. **Store.** `internal/store/queries/*.sql` updated; `sqlc generate` run;
   generated code committed. CRUD paths (`CreatePage`, `UpdatePage`, `GetPage*`)
   carry the new field.
4. **Handler.** `internal/handler/pages*.go` parses the form input, runs a
   `validateCacheTTL` helper returning `""`/error-string per the existing
   validation-function convention, and surfaces errors via flash + re-render.
5. **Admin view (templ).** The page edit form gains a numeric input,
   a label, and help text. Use an existing templUI component if one fits;
   if not, justify the choice in code comments.
6. **Cache middleware.** `internal/middleware/` page-cache middleware reads
   the per-page override (via context or a lookup hook — pick the cheaper
   option and justify it) and applies it for SET; `0` bypasses both SET and
   GET for that page's cache key.
7. **Translations.** Add keys for label, placeholder, help text, and
   validation error in `en` and `ru` locale files. Keys must follow the
   repo's existing naming convention. Do not add French.
8. **Docs.** Update `docs/caching.md` with a new subsection. Per `CLAUDE.md`'s
   wiki-sync requirement, also update `wiki/Caching.md` (the GitHub Wiki
   submodule). Do not run `/update-wiki`.
9. **Tests.** Add unit tests for:
   - Migration up/down round-trip.
   - Store CRUD with the new field (nil, 0, positive, out-of-range rejected).
   - `validateCacheTTL` table-driven test.
   - Handler form-parse test.
   - Cache middleware: per-page TTL honored, `0` bypasses, `NULL` falls back
     to global.
   - A drift test in the style of `internal/api/v2/drift_test.go` that fails
     if `cache_ttl_seconds` disappears from the pages store struct.
10. **Quality gate.** Run `/code-quality` (or `make code-quality-local`);
    zero findings — no `//nolint` comments added, no duplicate code, no
    unhandled errors, no shadowed variables, no package-name collisions.
11. **Security gate.** Run `/security-audit`; zero new findings. Explicitly
    reason about cache poisoning (can an unauthenticated user set a 0 TTL?),
    input bounds, CSRF coverage on the new form field, and whether a
    per-page `0` creates a DoS vector against the origin.
12. **Performance.** No new N+1 queries (confirm by reading the sqlc query
    plan / generated Go). Cache-key generation stays constant-time. If a
    microbenchmark exists for the cache package, run it before and after;
    report deltas.
13. **Final verification.** Run `make check` (or the closest equivalent in
    this repo). gofmt, vet, tests, and any project linter must pass. The
    server must still boot: `OCMS_SESSION_SECRET=test-secret-key-32-bytes-long!!!
    go run ./cmd/ocms &` → `curl -s -o /dev/null -w "%{http_code}"
    http://localhost:8080/` returns 200 → kill the server.

## Required output

Post-run summary at the top of the final message, using this structure:

```
# Per-page cache TTL override — baseline run

## Files changed
<grouped: migration, sqlc, store, handler, view, middleware, i18n, docs, tests>

## Test results
<pasted `make check` tail, including test count and duration>

## Quality gate
<`/code-quality` summary — 0 findings or the findings themselves>

## Security gate
<`/security-audit` summary — 0 findings or the findings themselves>

## Performance notes
<N+1 check result, any benchmark deltas>
```

Do NOT run `/cost`, report token counts, or try to measure wall-clock
yourself. Those three numbers are captured by the shell harness from
`claude -p --output-format=json` and the surrounding `date +%s`. Any
self-measurement wastes tokens and muddies the comparison.

## What "done" looks like

All 13 acceptance criteria met, `make check` clean, `/code-quality` clean,
`/security-audit` clean, server boots. No PR opened — stop at "ready to PR".
