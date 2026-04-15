# wiki/Multi-Language.md (oCMS)

## Summary

Wiki version of the multi-language reference. Near-parity with `docs/multi-language.md`: same two-part architecture (content translation + admin UI localization), same language-setup flow under **Admin > Config > Languages**, same translation linking model for pages/categories/tags, same URL-prefix convention, same four-step language detection order (URL → cookie → `Accept-Language` → default), same RTL handling, same admin UI localization flow, same theme integration (`{{.CurrentLanguage}}`, `{{.LanguageDirection}}`, `{{range .Translations}}…`), same hreflang emission, same menu-per-language pattern.

Notable difference: **omits the `Content-Language` HTTP header** mention present in `docs/multi-language.md` under "Language Headers". Minor documentation drift; not a factual contradiction (the header is still sent in practice — this is an omission).

Wikilinks replace markdown paths: `[[Internationalization]]`, `[[SEO]]`, `[[Menu Builder]]`, `[[REST API]]`.

## Sources

- Origin: `raw/ocms-go.core/wiki/Multi-Language.md`
