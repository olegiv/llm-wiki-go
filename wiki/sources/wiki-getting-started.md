# wiki/Getting-Started.md (oCMS)

## Summary

Onboarding guide for new oCMS developers. Covers:

- **Prerequisites** — Go 1.26+, Node.js/npm, sqlc, templ, goose, Dart Sass, libvips; per-OS libvips install instructions (macOS Homebrew, Ubuntu/Debian apt, Fedora dnf).
- **Installation** — five-step sequence: clone, `go mod download`, `go install` the tools (sqlc/templ/goose), regenerate code, `make assets`.
- **Quick start** — export `OCMS_SESSION_SECRET`, `make install-hooks`, then `make dev` (full) or `make run` (skip asset rebuild). App starts at `localhost:8080`.
- **Default admin credentials** — `admin@example.com` / `changeme1234` (requires `OCMS_DO_SEED=true`).
- **Demo mode** — `OCMS_DEMO_MODE=true` (requires seeding) creates demo users `demo@example.com` / `demo1234demo` (admin) and `editor@example.com` / `demo1234demo` (editor), plus sample pages/categories/tags/media/menu items.
- **Make command reference** — 19 commands including build, run, cross-compile, migrations, assets, `sqlc`/`templ` regeneration, `make install-hooks`, `make deploy-binary`.
- **Testing** — run with required `OCMS_SESSION_SECRET=test-secret-key-32-bytes-long!!` and optional `-v` / package-scoped commands. `govulncheck ./...` for dependency scanning.
- **Project structure** — ASCII tree of the repo layout (`cmd/`, `internal/`, `modules/`, `custom/`, `web/`, `uploads/`, etc.).
- **Next steps** — wikilinks to Configuration, Docker, REST API, Theme System, Module System pages.

Overlaps substantially with the "Installation", "Environment Variables", and "Development" sections of `README.md`.

## Sources

- Origin: `raw/ocms-go.core/wiki/Getting-Started.md`
