# wiki/Reverse-Proxy.md (oCMS)

## Summary

Wiki version of the reverse-proxy guide. Substantially overlaps `docs/reverse-proxy.md` — same three-backend coverage (Nginx, Apache, Nginx Proxy Manager), same SSL via Let's Encrypt, same troubleshooting table (502 / 504 / 413 / mixed-content), same performance tips.

Notable **differences from `docs/reverse-proxy.md`**:

- **Security headers:** wiki omits the deprecated `X-XSS-Protection` header that docs includes. (Modern browsers ignore this header; the wiki is the more accurate source.)
- **Duplication warning:** wiki explicitly notes "oCMS already sets security headers (CSP, HSTS, etc.). Avoid duplicating headers between the proxy and application." Docs does not.
- Shorter Nginx config (drops session ticket off, session cache size, commented-out HSTS).
- Docker-network configuration section in docs is absent from wiki.
- Nginx Proxy Manager details table omitted; just a short YAML and four-step config list.
- Uses `[[Caching|Redis for distributed caching]]` Obsidian-style cross-link; docs uses plain prose.

## Sources

- Origin: `raw/ocms-go.core/wiki/Reverse-Proxy.md`
