# Groxy

Groxy works as CGI and proxies requests from clients to upstream.
It helps to expose server processes which listen to HTTP directly on shared hosts.

## Upstream

Supported upstreams which listen to below:

- Unix domain socket

Supported protocols are below

- http

## Configuring

Configure everything with environment variables.

- `GROXY_UPSTREAM_ADDRESS` (Required) - Full path to the .sock file to be listened to by upstream.
- `GROXY_UPSTREAM_TIMEOUT` (Default: 180) - Timeout seconds to wait for response from upstream. Set 0 to disable the timeout.
- `HTTP_*` (Optional) - Custom header to add or override request headers.

For example, create the following `~/public_html/.htaccess`

```apacheconf
SetEnv UPSTREAM_SOCKET_PATH unix:///home/user/.local/state/app/app.sock
SetEnv UPSTREAM_TIMEOUT 300
SetEnv HTTP_USER_AGENT "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:86.0)"

<IfModule mod_dir.c>
    DirectoryIndex  Groxy
</IfModule>

<IfModule mod_rewrite.c>
    RewriteEngine On
    RewriteBase /
    RewriteRule ^(.*)$ Groxy [QSA,PT,L]
</IfModule>
```

## Contributing
