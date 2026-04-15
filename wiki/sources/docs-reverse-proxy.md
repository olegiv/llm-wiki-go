# docs/reverse-proxy.md (oCMS)

## Summary

Detailed reverse-proxy configuration guide for Nginx, Apache, and Nginx Proxy Manager.

### Prerequisites

oCMS on `localhost:8080`, domain, SSL cert (Let's Encrypt recommended). Set `OCMS_SERVER_HOST=127.0.0.1`, `OCMS_SERVER_PORT=8080`, `OCMS_ENV=production`.

### Nginx

Full config example with:

- 80 → 443 permanent redirect.
- Modern SSL: TLS 1.2/1.3, ECDHE ciphers, session cache.
- HSTS (commented out; opt-in).
- `client_max_body_size 100M`.
- Proxy headers: `Host`, `X-Real-IP`, `X-Forwarded-For`, `X-Forwarded-Proto`.
- 60-second connect/send/read timeouts.
- WebSocket upgrade headers on `location /`.
- Long-lived caching for `/static/`, `/themes/`, `/uploads/` (1 year, 1 year, 30 days respectively, `Cache-Control: public, immutable` for static/themes).
- `access_log off` on `/health`.
- `location ~ /\\. { deny all; }` to block dotfiles.

### Nginx load balancing

`upstream ocms_backend { least_conn; server 127.0.0.1:8080; server 127.0.0.1:8081; server 127.0.0.1:8082 backup; keepalive 32; }`. Requires `OCMS_REDIS_URL` for distributed caching.

### Apache

- Enable `proxy`, `proxy_http`, `ssl`, `headers`, `rewrite`.
- HTTP → HTTPS rewrite.
- SSL: `SSLProtocol all -SSLv3 -TLSv1 -TLSv1.1`.
- `ProxyPreserveHost On`, `ProxyRequests Off`.
- `RequestHeader set X-Real-IP "%{REMOTE_ADDR}s"`, `X-Forwarded-Proto https`.
- `ProxyPass / http://127.0.0.1:8080/`, `ProxyPassReverse / …`.
- `ProxyTimeout 60`.
- Header-based cache hints on `/static`, `/themes`, `/uploads`.
- `SetEnv nolog` on `/health`; `env=!nolog` on `CustomLog`.

### Apache load balancing

`<Proxy "balancer://ocms_cluster">` with three `BalancerMember` entries (`status=+H` marks hot spare), `lbmethod=bybusyness`. Enable `proxy_balancer` and `lbmethod_bybusyness` modules.

### Nginx Proxy Manager

Docker Compose YAML, web UI at port 81, default `admin@example.com` / `changeme`. Configure Proxy Host: forward 127.0.0.1:8080, enable Block Common Exploits + Websockets Support. SSL tab: Request new cert + Force SSL + HTTP/2 + HSTS. Advanced tab includes the same cache/static-asset snippets.

### SSL via Let's Encrypt

`sudo certbot --nginx -d example.com -d www.example.com` for Nginx. `--apache` for Apache. `systemctl enable certbot.timer` for auto-renewal.

### Security headers

Documents `X-Frame-Options`, `X-Content-Type-Options`, **`X-XSS-Protection`** (deprecated in modern browsers), `Referrer-Policy`, `Permissions-Policy`, and a CSP example.

### Troubleshooting

502 / 504 / 413 / mixed-content Origin diagnosis. "Verify proxy headers" step. Health-check monitoring snippet.

### Next steps

Points at `webhooks.md`, `multi-language.md`, `import-export.md`.

## Sources

- Origin: `raw/ocms-go.core/docs/reverse-proxy.md`
