# wiki/hCaptcha.md (oCMS)

## Summary

Wiki reference for the hCaptcha module. Covers same account setup, same Admin UI path, same environment variables (`OCMS_HCAPTCHA_SITE_KEY`, `OCMS_HCAPTCHA_SECRET_KEY`), same five configuration options (Enabled, Site Key, Secret Key, Theme, Size), same three-layer login protection integration, same error messages, same error codes, same reverse-proxy header order, same disable paths (admin UI, DB setting, DB module).

**Does not document `OCMS_HCAPTCHA_DISABLED`** — the env-var kill switch covered in `docs/hcaptcha.md` is absent from the wiki page. Consistent with `wiki/Configuration.md` also omitting this variable.

**Inherits the stale test-key-defaults claim** from `docs/hcaptcha.md`: lists test site/secret keys as "used as defaults when no keys are configured" — contradicted by `CHANGELOG.md` 0.18.1 which documents their removal.

Omits the "Module Information" table (module name, version, DB table) present in the docs version.

## Sources

- Origin: `raw/ocms-go.core/wiki/hCaptcha.md`
