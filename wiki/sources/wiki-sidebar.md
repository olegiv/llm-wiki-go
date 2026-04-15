# wiki/_Sidebar.md (oCMS)

## Summary

Navigation sidebar for the GitHub wiki — no H1 in the original; the first line is bold-wrapped `Home` link. It groups wiki pages into nine navigation sections:

- **Getting Started** — Getting Started, Configuration, Docker.
- **Core Features** — Content Management, Media Library, Taxonomy, Menu Builder, Forms, Multi-Language, SEO.
- **Administration** — Admin Interface, Import-Export, Caching, Scheduler.
- **API & Integrations** — REST API, Webhooks.
- **Security** — Security, Login Security, CSRF Protection, hCaptcha, GeoIP.
- **Theme Development** — Theme System.
- **Module Development** — Module System, Internationalization.
- **Built-in Modules** — Module: Sentinel, Module: Developer, Module: DB Manager, Module: Informer, Module: Demo Mode.
- **Deployment** — Deployment, Reverse Proxy, Deploy: Fly.io, Deploy: Ubuntu Plesk.

Also links Contributing at the bottom.

The "Built-in Modules" sidebar list matches five of the eleven on-disk `modules/` directories. Same discrepancy documented on [topics/module-system.md](../topics/module-system.md) — analytics_ext / analytics_int / embed / example / hcaptcha / migrator / privacy are not represented in the sidebar, and "Demo Mode" is in the sidebar but has no `modules/` directory.

## Sources

- Origin: `raw/ocms-go.core/wiki/_Sidebar.md`
