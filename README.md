## nginx-to-librato

This is a small application that polls the [HttpStubStatusModule](http://wiki.nginx.org/HttpStubStatusModule)
in nginx, then formats and publishes the resulting data to Librato.

## Usage

    nginx-to-librato -c /etc/nginx-to-librato.conf

## Configuration

    [settings]
    token: your_librato_token
    user: your_librato_email
    source: load-balancer-001
    url: 127.0.0.1:8000/nginx_status
    flush_interval: 10s


## Why?

I didn't need collectd, which librato integrates with. This is easy
to deploy and just generally makes more sense to me.

## Upstart Example

TODO

## nginx Configuration

This assumes you've configured the HttpStubStatusModule like this:

```
location /nginx_status {
  stub_status on;
  allow 127.0.0.1;
  deny all;
}
```

This only allows requests from wherever nginx is located. That is
where you should install nginx-to-librato.

## License

Please see [the license file](License.md).
