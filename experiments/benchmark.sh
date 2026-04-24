#!/usr/bin/env bash
set -u

REPO=~/Desktop/Projects/Go/ocms-go
WIKI=~/Desktop/Projects/Go/llm-wiki-go
OUT=$WIKI/experiments/quick-results
mkdir -p "$OUT"

read -r -d '' BASELINE_PROMPT <<'EOF' || true
Inventory every cache layer in oCMS. For each layer report:
  1. Package / file that owns it
  2. Env var(s) that configure it and defaults
  3. TTL default
  4. What invalidates it
  5. Which subsystems consume it
Cite every claim with `path:line-range`.

Hard constraints:
- Read only files inside the current `ocms-go` repository.
- Do NOT read or grep anything under `../llm-wiki-go/`. If any tool output
  contains a `llm-wiki-go` path, stop and report contamination instead.
- Do NOT fetch the web.
- Do NOT run `/cost`, report tokens, or measure wall-clock — the harness
  captures that from outside.
EOF

read -r -d '' TREATMENT_PROMPT <<'EOF' || true
Inventory every cache layer in oCMS. For each layer report:
  1. Package / file that owns it
  2. Env var(s) that configure it and defaults
  3. TTL default
  4. What invalidates it
  5. Which subsystems consume it

Hard constraints:
- Start from `../llm-wiki-go/wiki/index.md`. Read `wiki/topics/caching.md`
  and related entity / source pages first.
- Only fall back to `ocms-go` source when the wiki genuinely lacks a fact.
  List such fallbacks in a `## Wiki gaps` section at the end.
- Quote any `## Contradictions` block verbatim — do not silently pick a winner.
- Cite each claim by wiki page path (and raw path:line only for gaps).
- Do NOT run `/cost`, report tokens, or measure wall-clock.
EOF

run() {
  local arm=$1 n=$2 prompt extra
  if [[ $arm == baseline ]]; then
    prompt=$BASELINE_PROMPT
    extra=()
  else
    prompt=$TREATMENT_PROMPT
    extra=(--add-dir "$WIKI")
  fi

  cd "$REPO"
  local t0 t1
  t0=$(date +%s)
  if ! claude -p "$prompt" \
       --dangerously-skip-permissions \
       --output-format=json \
       --model=claude-sonnet-4-6 \
       "${extra[@]}" \
       > "$OUT/$arm-$n.json" 2> "$OUT/$arm-$n.err"; then
    echo "FAILED: $arm-$n (see $OUT/$arm-$n.err)"
  fi
  t1=$(date +%s)
  echo $((t1-t0)) > "$OUT/$arm-$n.wallclock"
  echo "$arm-$n: $((t1-t0))s"
}

for i in 1 2 3; do
  run treatment $i
  run baseline $i
done

cd "$OUT"
{
  echo "arm,run,wall_s,cost_usd,in_tok,out_tok,cache_r_tok,cache_w_tok"
  for f in *.json; do
    base=${f%.json}; arm=${base%-*}; n=${base##*-}
    [[ -s $base.wallclock ]] || continue
    wall=$(cat "$base.wallclock")
    jq -r --arg a "$arm" --arg n "$n" --arg w "$wall" \
      '[$a, $n, $w, (.total_cost_usd // 0),
        (.usage.input_tokens // 0), (.usage.output_tokens // 0),
        (.usage.cache_read_input_tokens // 0),
        (.usage.cache_creation_input_tokens // 0)] | @csv' "$f" 2>/dev/null
  done
} > summary.csv

echo ""
echo "=== summary ==="
column -s, -t < summary.csv
