# oCMS feature ship — wiki-first treatment

Claude Code prompt for the **treatment arm** of a wiki-vs-no-wiki experiment.
Control arm lives in `./BENCHMARK_BASELINE.md`.

Run headless via `claude -p "$(cat ../llm-wiki-go/BENCHMARK_TREATMENT.md)"
--dangerously-skip-permissions --add-dir ../llm-wiki-go --output-format=json`
from a fresh clean branch in `../ocms-go/`. Wall-clock, tokens, and cost
are captured by the harness — this prompt does not self-measure.

---

You are measuring how Claude Code performs on a realistic feature-ship
workflow **with** the LLM Wiki. The feature and the acceptance criteria
are identical to the baseline. The *method* is different: follow the
answer-from-wiki-first rule from this repo's `CLAUDE.md` for every
discovery step.

## Hard constraints

- Before opening any file in `../ocms-go/`, read `wiki/index.md` in this
  repo to see what pages exist.
- For every subsystem you need to understand (cache, pages entity, handler
  pattern, migrations, i18n, docs layout, validation helpers), read the
  relevant `wiki/topics/*.md` or `wiki/entities/*.md` page **first**.
  Only then touch `../ocms-go/` source.
- Follow `## Sources` links into `raw/` inside this repo for line-level
  detail before jumping to `../ocms-go/` directly.
- Quote any `## Contradictions` block verbatim when you encounter one — do
  not silently pick a winner.
- You may still read `../ocms-go/` files that the wiki sends you to via
  `## Sources`, or to edit (the work happens there). You may NOT use
  `../ocms-go/` as your discovery starting point.
- You are already on a clean benchmark branch in `../ocms-go/` prepared by
  the harness. Work on that branch; do not switch branches.
- Follow every rule in `../ocms-go/CLAUDE.md` and `../ocms-go/AGENTS.md`
  (license headers, error wrapping, slog fields, naming, validation
  functions, resource cleanup). Deviations count against correctness.
- Treat `wiki/` as read-only for the duration of the benchmark. Do not
  compile new facts back during the run (that's a separate experiment).

## The feature

Identical to `./BENCHMARK_BASELINE.md` — do not paraphrase or narrow.

Add a **per-page cache TTL override**. Global cache TTL (`OCMS_CACHE_TTL`)
is unchanged; each page may optionally override it.

### Acceptance criteria

1. **Schema.** New goose migration adds a nullable `cache_ttl_seconds`
   INTEGER column to the `pages` table. `down` removes it cleanly.
2. **Semantics.**
   - `NULL` → use global TTL.
   - `0` → never cache this page.
   - `1..2592000` (30 days) → override TTL in seconds.
   - Anything else → validation error at write time.
3. **Store.** `internal/store/queries/*.sql` updated; `sqlc generate` run;
   generated code committed. CRUD paths carry the new field.
4. **Handler.** Parse form input, `validateCacheTTL` helper returning
   `""`/error-string per the repo's validation-function convention, flash +
   re-render on error.
5. **Admin view (templ).** Numeric input with label and help text. Prefer
   templUI; justify in a comment if you roll your own.
6. **Cache middleware.** Honor per-page override on SET; `0` bypasses both
   SET and GET.
7. **Translations.** `en` and `ru` only. Follow existing key conventions.
8. **Docs.** `docs/caching.md` and `wiki/Caching.md` (oCMS's GitHub Wiki
   submodule — *not* this repo). Do not run `/update-wiki`.
9. **Tests.** Migration round-trip, store CRUD (nil / 0 / positive /
   out-of-range), `validateCacheTTL` table test, handler form-parse test,
   middleware TTL-resolution test, drift test that fails if
   `cache_ttl_seconds` disappears from the store struct.
10. **Quality gate.** `/code-quality` clean, no `//nolint`, no duplicate
    code, no unhandled errors.
11. **Security gate.** `/security-audit` clean. Reason explicitly about
    cache poisoning, input bounds, CSRF on the new field, and whether `0`
    creates a DoS vector.
12. **Performance.** No new N+1s, cache-key generation constant-time,
    benchmark deltas reported if benchmarks exist.
13. **Final verification.** `make check` clean, server boots, branch ready
    to PR but no PR opened.

## Required output

Post-run summary at the top of the final message:

```
# Per-page cache TTL override — wiki-first run

## Wiki pages consulted
<list wiki/ pages in order of first access, with the question each answered>

## Contradictions surfaced
<any `## Contradictions` blocks encountered, quoted verbatim with wiki paths.
 If this section is empty, state that explicitly.>

## Wiki gaps (compile-back targets)
<facts the wiki didn't have that you had to fetch from `../ocms-go/` directly.
 Each: what was missing, which wiki page should have had it, where you found
 it. If empty, state that explicitly — it's the strongest possible result.>

## Files changed
<grouped: migration, sqlc, store, handler, view, middleware, i18n, docs, tests>

## Test results
<pasted `make check` tail>

## Quality gate
<`/code-quality` summary>

## Security gate
<`/security-audit` summary>

## Performance notes
<N+1 check, benchmark deltas>
```

Do NOT run `/cost`, report token counts, or try to measure wall-clock
yourself. Those three numbers are captured by the shell harness from
`claude -p --output-format=json` and the surrounding `date +%s`. Any
self-measurement wastes tokens and muddies the comparison.

## What "done" looks like

All 13 acceptance criteria met; `make check`, `/code-quality`, and
`/security-audit` clean; server boots; branch ready to PR; wiki-pages,
contradictions, and wiki-gaps sections populated honestly.

## After the run — experiment hygiene

Run each arm three or four times in fresh sessions on the same model with
only filesystem tooling, then compare:

- `wall_clock_seconds`, `total_tokens`, `total_cost_usd` — efficiency delta.
- Acceptance-criteria pass rate — correctness delta (any failure collapses
  the whole run, since the gates are pass/fail).
- Treatment's `## Contradictions surfaced` — drift the baseline silently
  resolved; this is the wiki's unique contribution.
- Treatment's `## Wiki gaps` — compile-back backlog for the next wiki pass;
  each run feeds the next improvement to `llm-wiki-go`.

Record per-run numbers in a small table; don't average one run each —
LLM variance will swamp a single-run comparison.
