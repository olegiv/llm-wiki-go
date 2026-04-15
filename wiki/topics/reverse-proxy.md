# Reverse proxy

## Summary

oCMS should always run behind a reverse proxy in production for SSL termination, static-asset caching, WebSocket upgrades, and improved security. This topic covers Nginx, Apache, and Nginx Proxy Manager, and also captures the interaction with oCMS's own trusted-proxy hardening (since 0.17.0).

## Environment configuration

```bash
export OCMS_SERVER_HOST=127.0.0.1
export OCMS_SERVER_PORT=8080
export OCMS_ENV=production
export OCMS_TRUSTED_PROXIES=127.0.0.1/32   # or your load balancer CIDR
```

`OCMS_TRUSTED_PROXIES` is mandatory for correct client-IP resolution; without it, chi's replaced trusted-proxy-aware middleware (0.17.0) rejects forwarding headers. `OCMS_REQUIRE_TRUSTED_PROXIES` forces startup failure in production when unset.

## Nginx

### Baseline server block

- 80 → 443 permanent redirect.
- TLS 1.2/1.3, ECDHE ciphers, session cache.
- `client_max_body_size 100M` for uploads.
- Proxy headers: `Host`, `X-Real-IP`, `X-Forwarded-For`, `X-Forwarded-Proto`.
- Timeouts: 60 seconds connect / send / read.
- `location /` with WebSocket upgrade headers.
- Long-lived caching: `/static/` and `/themes/` for 1 year (immutable), `/uploads/` for 30 days.
- `access_log off` on `/health`.
- `location ~ /\\. { deny all; }` to block dotfiles.

### Load balancing

```nginx
upstream ocms_backend {
    least_conn;
    server 127.0.0.1:8080;
    server 127.0.0.1:8081;
    server 127.0.0.1:8082 backup;
    keepalive 32;
}
```

Distributed caching (Redis) is mandatory for multi-instance deployments — set `OCMS_REDIS_URL`.

## Apache

Enable `proxy`, `proxy_http`, `ssl`, `headers`, `rewrite` (plus `proxy_balancer`, `lbmethod_bybusyness` for load balancing).

```apache
<VirtualHost *:443>
    ServerName example.com
    SSLEngine on
    SSLCertificateFile …
    SSLCertificateKeyFile …
    SSLProtocol all -SSLv3 -TLSv1 -TLSv1.1

    ProxyPreserveHost On
    ProxyRequests Off
    RequestHeader set X-Real-IP "%{REMOTE_ADDR}s"
    RequestHeader set X-Forwarded-Proto "https"

    ProxyPass / http://127.0.0.1:8080/
    ProxyPassReverse / http://127.0.0.1:8080/
    ProxyTimeout 60
</VirtualHost>
```

Load balancer example uses `<Proxy "balancer://ocms_cluster">` with `lbmethod=bybusyness`; hot spares via `status=+H`.

## Nginx Proxy Manager (GUI)

Docker Compose image `jc21/nginx-proxy-manager:latest`, admin at port 81 (default `admin@example.com` / `changeme`). Proxy Host: forward 127.0.0.1:8080, enable Block Common Exploits + Websockets Support. SSL tab: request cert + Force SSL + HTTP/2 + HSTS. Advanced tab holds the static-asset caching blocks.

## Security headers at the proxy layer

The wiki page advises **not** to duplicate headers between proxy and application:

> oCMS already sets security headers (CSP, HSTS, etc.). Avoid duplicating headers between the proxy and application.

If you do add headers at the proxy, the docs page recommends `X-Frame-Options`, `X-Content-Type-Options`, `Referrer-Policy`, `Permissions-Policy`, plus a CSP. It also suggests `X-XSS-Protection`, **which is deprecated in modern browsers** — prefer the set in the wiki page.

## SSL via Let's Encrypt

```bash
sudo apt install certbot python3-certbot-nginx   # or python3-certbot-apache
sudo certbot --nginx -d example.com -d www.example.com
sudo systemctl enable certbot.timer
```

## Troubleshooting

| Symptom | Likely cause | Fix |
|---------|--------------|-----|
| 502 Bad Gateway | oCMS not running | `curl http://127.0.0.1:8080/health` |
| 504 Gateway Timeout | Slow response | Increase proxy timeouts; check resources |
| 413 Request Entity Too Large | Upload too big | Raise `client_max_body_size` |
| Mixed-content warnings | Missing `X-Forwarded-Proto` | Ensure proxy passes `https` scheme |
| IP bans miss real attackers | Trusted-proxy misconfig | Set `OCMS_TRUSTED_PROXIES` to cover the proxy layer |

## Performance tips

1. Enable HTTP/2.
2. Let oCMS handle gzip — don't double-compress at the proxy.
3. Connection keepalive for repeated requests.
4. Long cache times for versioned static assets.
5. Don't poll `/health` too frequently — every 30 seconds is plenty.

## Contradictions

### `X-XSS-Protection` header

- [sources/docs-reverse-proxy.md](../sources/docs-reverse-proxy.md) lists `X-XSS-Protection "1; mode=block"` in both Nginx and Apache security-header sections.
- [sources/wiki-reverse-proxy.md](../sources/wiki-reverse-proxy.md) omits this header and explicitly warns against duplicating headers that oCMS already sets.

Modern browsers ignore `X-XSS-Protection`; the wiki is the better reference.

## Sources

- [sources/docs-reverse-proxy.md](../sources/docs-reverse-proxy.md)
- [sources/wiki-reverse-proxy.md](../sources/wiki-reverse-proxy.md)
- [sources/security.md](../sources/security.md)
- [sources/changelog.md](../sources/changelog.md)
