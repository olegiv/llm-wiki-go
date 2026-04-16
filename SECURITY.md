# Security Policy

## Reporting a Vulnerability

Please **do not** open a public GitHub issue for security vulnerabilities.

Instead, report them privately through GitHub's private vulnerability
reporting:

> [Report a vulnerability](https://github.com/olegiv/llm-wiki-go/security/advisories/new)

Include as much of the following as you can:

- A description of the issue and its impact.
- Steps to reproduce, a proof of concept, or a minimal test case.
- The commit or release you observed the issue on.
- Any suggested mitigation, if you have one.

## What to Expect

- **Acknowledgement:** within a few days of the report.
- **Triage:** an initial assessment of severity and scope, and a decision
  on whether to accept the report.
- **Fix and disclosure:** coordinated through the GitHub advisory. Once a
  fix is available, the advisory is published and credit is given to the
  reporter unless they prefer to remain anonymous.

## Scope

In scope:

- The `wikilint` CLI and the Go packages under `cmd/` and `internal/`.
- The `Makefile` targets and any scripts checked into this repository.
- The `.claude/` skills and agents defined in this repository.

Out of scope:

- Content-level issues in `raw/` or `wiki/` — `raw/` is immutable source
  material and `wiki/` is compiled output. Both ship with the repo, but
  content disagreements or factual errors are not security
  vulnerabilities.
- Upstream dependencies (Go toolchain, standard library, third-party
  submodules). Report those to their respective maintainers.
- Social-engineering or physical-access attacks on individual maintainers.

## Supported Versions

This project does not yet cut numbered releases. Security fixes land on
the default branch (`main`). Users are expected to track `main` or pin
to a specific commit.
