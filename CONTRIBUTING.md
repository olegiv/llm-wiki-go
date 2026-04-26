# Contributing

Thanks for your interest in `llm-wiki-go`. This guide covers the
practical bits for human contributors. Automated agents working in this
repo should read [`AGENTS.md`](AGENTS.md) (and, for Claude Code
specifically, [`CLAUDE.md`](CLAUDE.md)) instead — it has the full rules.

## Getting set up

```bash
git clone --recurse-submodules https://github.com/olegiv/llm-wiki-go.git
cd llm-wiki-go
make setup
```

`raw/` and `wiki/` are tracked in git and ship with the repo. For the
default `ocms-go.core` example, just clone `ocms-go.core` as a sibling.
`make setup` is available for bootstrapping against a different source
repo — it creates the directory structure and seeds `wiki/index.md` and
`wiki/log.md`.

## Development loop

1. Make your change.
2. Run the full verification chain:

   ```bash
   make check
   ```

   This runs `gofumpt -l .`, `go vet ./...`,
   `golangci-lint run ./...`, `wikilint`, and `go test ./...`.
   All steps must succeed and the wiki linter must print exactly
   `wikilint: OK`.
3. Commit and push.

Use `make help` to see every available target.

## Commit style

- One logical change per commit. If you touch unrelated files, split the
  work into separate commits.
- Imperative subject line (≤ 72 chars): "Add X", not "Added X" or
  "Adding X".
- Use the body to explain *why*, not *what*. The diff already shows
  what.
- No drive-by refactors, reformatting, or unrequested API changes.

## Wiki edits

- **Never modify files under `raw/`.** If raw content looks wrong, flag
  it in the PR or issue instead of fixing it in place.
- Every substantive wiki page has exactly one H1 title, a `## Summary`
  section, and a `## Sources` section (`wiki/index.md` and
  `wiki/log.md` are exempt).
- When sources disagree, record each position under a
  `## Contradictions` section — do not silently merge.
- Every wiki change updates `wiki/index.md` (so new pages stay
  reachable) and appends a short entry to `wiki/log.md` describing what
  changed.

## Reporting security issues

Do not open a public issue for security vulnerabilities. Follow the
process in [`SECURITY.md`](SECURITY.md).

## License

By contributing, you agree that your contributions are released under
the [GNU General Public License v3.0](LICENSE).
