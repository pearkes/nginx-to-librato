## nginx-to-librato

This is a small application that polls the [HttpStubStatusModule](http://wiki.nginx.org/HttpStubStatusModule)
in nginx, then formats and publishes the resulting data to Librato.

## Usage

    nginx-to-librato [-debug] -config /etc/nginx-to-librato.conf

## Configuration

    [settings]
    token: your_librato_token
    user: your_librato_email
    source: load-balancer-001
    url: 127.0.0.1:8000/nginx_status
    flush_interval: 10s

## nginx Configuration

This assumes you've configured the HttpStubStatusModule like this:

    location /nginx_status {
      stub_status on;
      allow 127.0.0.1;
      deny all;
    }

This only allows requests from wherever nginx is located. That is
where you should install nginx-to-librato. Alternatively, you
could put the application on a monitoring server and only allow
that IP to access the `/nginx_status` page.

## Upstart Example

    description "nginx-to-librato"

    start on runlevel [23]

    env TARGET=/usr/local/bin/nginx-to-librato
    env LOG=/usr/local/bin/nginx-to-librato
    env CONF=/etc/nginx-to-librato.conf

    respawn

    exec su -m -l -c "$TARGET -config $CONF >> $LOG"

## License

Please see [the license file](LICENSE.md).
