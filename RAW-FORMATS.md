# Raw source formats

Guidance for what can go under `raw/` and what the ingesting agent can
actually do with it. This is tool meta-documentation for
`llm-wiki-go`; it is intentionally kept out of `wiki/` because `wiki/`
is the compiled knowledge layer about the sample project (currently
`ocms-go.core`), not about the scaffolding.

## TL;DR

`raw/` is **format-agnostic at the code level** — the Go helpers do
not filter by extension. The real constraint is what the ingesting
agent can read. Plain-text-like formats (Markdown, `.txt`, source
code, JSON/YAML/XML, `.log`, `.csv`), PDFs, images, and Jupyter
notebooks work directly. Other formats — notably MS Office — need to
be converted to a text-like format before they become useful. A
hard 10 MiB per-file cap applies.

## What the code accepts

- `internal/wiki/ingest.go` — `ListRawFiles` (line 22) walks every
  regular file under the raw directory. The only file it skips is
  `.gitkeep` (line 35). There is no extension allow-list or
  deny-list.
- `internal/wiki/ingest.go` — `ReadRawFile` (line 53) reads raw bytes
  up to `maxFileBytes = 10 << 20` (10 MiB, line 17). It does not
  parse or transform the content.
- `internal/wiki/ingest_test.go` — the test fixtures at lines 14–16
  deliberately mix `alpha.txt`, `nested/beta.log`, and
  `nested/deep/gamma.md` to demonstrate that the walker treats
  extensions as opaque.
- `README.md` — describes the payload generically as "documents,
  transcripts, notes, captures". No format is privileged.

So: if a file lives under `raw/` and is 10 MiB or smaller, the Go
code will list and read it without complaint.

## What the agent can actually ingest

The `ingest-source` skill is executed by a coding agent (Claude
Code or similar) reading the file through its own file-reading
contract. The agent's reader is the binding constraint, not the Go
code.

**Readable as-is (direct ingest works):**

- Markdown (`.md`), plain text (`.txt`), `.log`, `.csv`
- Source code in any language (the file is just text to the reader)
- Structured text: JSON, YAML, TOML, XML, HTML
- PDFs (up to ~20 pages per read for large files)
- Images: PNG, JPG, and other common raster formats — ingested
  visually, useful when the content is diagrammatic
- Jupyter notebooks (`.ipynb`) — cells and outputs are surfaced

**Readable with pre-conversion:**

- MS Office files (see below)
- Other proprietary document formats (Keynote, Pages, older
  WordPerfect, etc.)
- Audio and video — require transcription before they are useful as
  raw source

**Not meaningfully ingestable:**

- Executables, compiled object files, stripped binaries
- Compressed archives without a text payload the agent can unpack

## MS Office files

Can MS Office files sit directly under `raw/`?

**Physically: yes.** Nothing in `ListRawFiles` or `ReadRawFile`
rejects `.docx`, `.xlsx`, `.pptx`, `.doc`, `.xls`, or `.ppt`.

**Practically: no — not without conversion.**

- `.docx`, `.xlsx`, `.pptx` are ZIP containers wrapping XML parts.
  To a general-purpose file reader they appear as binary blobs. The
  ingesting agent's `Read` tool handles text, PDFs, images, and
  `.ipynb`; it does not natively extract Office documents.
- Legacy `.doc`, `.xls`, `.ppt` are proprietary OLE compound binary
  files and are even less tractable for direct ingest.
- The 10 MiB `maxFileBytes` cap covers most `.docx` and `.xlsx`
  files but is commonly exceeded by media-heavy `.pptx` decks.

Recommended handling (in order of preference):

1. **Convert to Markdown or plain text before dropping into
   `raw/`.** Tools: `pandoc -f docx -t gfm`, `docx2txt`, or
   `libreoffice --headless --convert-to txt`. Name the converted
   file descriptively, e.g. `raw/specs/product-brief.docx.md`, so
   the origin is obvious.
2. **Export to PDF** and commit the PDF. PDFs are directly
   ingestable.
3. **Archive the Office original outside `raw/`** (e.g. in object
   storage or a sibling repo) and only place the converted,
   agent-readable version under `raw/`.

The guiding rule: every file under `raw/` should be something the
ingesting agent can open and understand without additional tooling.
If a source arrives in Office format, convert it first rather than
committing the binary and hoping the pipeline will do something
useful with it.

## 10 MiB size cap

`ReadRawFile` hard-caps input at `maxFileBytes = 10 << 20` bytes.
Oversized files return an error that includes the actual and
maximum size; they are not silently truncated.

If a legitimate source exceeds 10 MiB:

- Split it into logical chunks (per chapter, per section, per day
  for long transcripts) and commit each chunk as its own file.
- If the file is inherently large (big PDFs, high-resolution image
  dumps), consider whether it belongs under `raw/` at all, or
  whether a pointer + external storage is more appropriate.

Raising `maxFileBytes` is a code change, not a workflow change, and
should be discussed before doing it — the cap exists to keep
ingest predictable and to avoid accidentally pulling gigabyte-scale
artifacts into the agent's context.
