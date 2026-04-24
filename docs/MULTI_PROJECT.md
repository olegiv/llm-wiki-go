# Using llm-wiki-go for a different project

The repo's default layout compiles a wiki for a single source repo —
the bundled example is `ocms-go.core`, with `raw/ocms-go.core/`
symlinked into `../ocms-go.core`. To compile a wiki for a different
project, or to run wikis for several projects in parallel, clone
`llm-wiki-go` a second time next to your new source repo. This page
documents the layout, the manual steps `make setup` doesn't handle,
and the caveats that come with running more than one wiki at a time.

## On-disk layout

Keep every wiki clone and every source repo as siblings under a single
parent directory. Nesting a wiki clone inside its source repo (or vice
versa) breaks relative symlinks and confuses editor tooling.

```
Projects/Go/
├── ocms-go.core/         # source repo
├── llm-wiki-go/          # wiki for ocms-go.core (the bundled default)
├── logwatch-ai-go/       # second source repo
└── logwatch-ai-wiki/     # second wiki clone; raw/ → ../logwatch-ai-go
```

Name each additional clone `<project>-wiki` so it lines up visually
with `<project>-go`. The first clone stays `llm-wiki-go` both because
that is what the upstream repo is called and because the default
`raw/ocms-go.core/` example already lives there.

## Bootstrapping a new wiki clone

The steps below assume you want a wiki for `logwatch-ai-go`. Replace
the name everywhere if your project is different.

```bash
cd Projects/Go
git clone https://github.com/olegiv/llm-wiki-go.git logwatch-ai-wiki
cd logwatch-ai-wiki
git submodule update --init --recursive

# The fresh clone ships with ocms-go.core's compiled wiki and raw
# symlinks. Clear both before starting the new project's wiki.
rm -rf wiki raw

# Seed empty wiki/ (entities/topics/sources/ + index.md + log.md) and
# an empty raw/. make setup is parameter-free.
make setup

# Create the raw/<project>/ symlink by hand. make setup does NOT do
# this, despite what the "bootstrapping against a different source
# repo" phrasing in the README might suggest.
ln -s ../../logwatch-ai-go raw/logwatch-ai-go

# Sanity-check: the symlink resolves and the linter passes on the
# seeded empty wiki.
ls -la raw/logwatch-ai-go
make lint   # expect: wikilint: OK
```

From here, follow the normal ingest loop described in the main
`README.md` ("Claude Code workflow" section): create `sources/` pages
for meaningful files in `raw/logwatch-ai-go/`, extract entities into
`wiki/entities/`, build topic pages in `wiki/topics/`, and keep
`wiki/index.md` + `wiki/log.md` current on every edit.

## Remote strategy

Each wiki clone is its own working tree and will diverge heavily from
the upstream `llm-wiki-go` repo under `raw/` and `wiki/`. You have
three reasonable options for where to push it:

The cleanest separation is to give each wiki its own remote — e.g.
create a `logwatch-ai-wiki` repo on GitHub and push there. Keep
`olegiv/llm-wiki-go` as an `upstream` remote so tooling updates can
still be pulled in.

```bash
cd logwatch-ai-wiki
git remote rename origin upstream
git remote add origin git@github.com:olegiv/logwatch-ai-wiki.git
git push -u origin main
```

The lightest-weight option is to keep the clone local-only until you
actually need a remote. Nothing in the tooling assumes an `origin`
exists.

The option to avoid is pushing the new wiki back to the shared
`llm-wiki-go` remote under a different branch. Those branches will
pick up ocms-go.core's `wiki/` and `raw/` on every pull and you will
fight merge conflicts on every tooling update.

## Keeping tooling in sync across clones

Every wiki clone carries its own copy of the tooling: `cmd/wikilint/`,
`internal/wikilint/`, `internal/wiki/`, `Makefile`, `AGENTS.md`,
`CLAUDE.md`, `CONTRIBUTING.md`, and the `.claude/` tree (including the
`.claude/shared/` submodule). When you fix a linter rule or update a
skill in one clone, pull it into the others.

Treat one clone as the canonical tooling copy — the original
`llm-wiki-go` is the obvious choice — and propagate changes outward
from there:

```bash
cd logwatch-ai-wiki
git fetch upstream main

# Merge only tooling paths, leaving wiki/ and raw/ untouched. Adjust
# the path list when the repo layout changes.
git checkout upstream/main -- \
    cmd/ internal/ Makefile \
    AGENTS.md CLAUDE.md CONTRIBUTING.md README.md \
    go.mod go.sum .claude/

git submodule update --recursive
git commit -m "Sync tooling from llm-wiki-go upstream"
```

What diverges per clone, by design: `raw/` (one symlink tree per
project) and `wiki/` (entirely different content, different index,
different log). Everything else should stay byte-identical to upstream
so `make check` behaves the same way in every clone.

## Caveats and constraints

`make setup` does not create the `raw/<project>/` symlink. It only
creates empty `raw/`, `wiki/entities/`, `wiki/topics/`, `wiki/sources/`
and seeds `wiki/index.md` + `wiki/log.md`. The symlink is a manual
step and has to be repeated on every new clone.

`wikilint` expects a single `wiki/index.md` at the top of `wiki/`.
That matches the one-wiki-per-clone pattern exactly; it does not
support a layout where `wiki/<project-a>/` and `wiki/<project-b>/` sit
side by side in the same clone. If you need that, see "When to pick a
different layout" below.

The `.claude/shared/` submodule is per-clone state. Run
`git submodule update --init --recursive` after every fresh clone,
and `git submodule update --recursive` after every upstream merge
that touches the submodule pointer.

Relative symlinks in `raw/` are resolved relative to `raw/`, not to
the repo root. `raw/logwatch-ai-go -> ../../logwatch-ai-go` is
correct; `raw/logwatch-ai-go -> ../logwatch-ai-go` is not. Verify
with `readlink raw/<project>` after creating the symlink.

## Version pinning

The wiki's claims are only reproducible against a specific snapshot
of the source repo. Record the source repo's commit SHA in your
wiki's `wiki/index.md` (near the top of `## Summary`) and bump it
with a dedicated `wiki/log.md` entry whenever you re-ingest after
meaningful source churn. Without this, old wiki pages silently drift
out of date as the source evolves.

## When to pick a different layout

The one-wiki-per-clone pattern stops being the right answer if you
genuinely need cross-project knowledge pages — e.g. a single topic
page comparing token-budget strategies across `ocms-go.core` and
`logwatch-ai-go`. At that point, a multi-project layout inside one
clone looks like this:

```
raw/
├── ocms-go.core/
└── logwatch-ai-go/

wiki/
├── ocms-go.core/       # index.md, log.md, entities/, topics/, sources/
├── logwatch-ai-go/     # same structure, different content
└── _cross/             # topics spanning projects
```

This is not supported out of the box. `make setup` needs a project
parameter, and `wikilint` needs to either run once per
`wiki/<project>/` or learn to recognize per-project sub-indexes
(right now it enforces a single top-level `wiki/index.md` as the
reachability root). Neither change is large, but both are real. Defer
this layout until you've actually felt the absence of cross-project
pages; until then, the sibling-clones pattern described above is
simpler and lets `wikilint` stay untouched.
