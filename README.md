Easy-to-configure Static Web/Reverse Proxy Server in Go

![](http://i.cloudup.com/i5Tpn00lCc.png)

## Install

```bash
$ go get github.com/azer/boxcars/boxcars
```

## Usage

Create a configuration file *(it'll be auto-loading changes once you start)*  like the below example;

```json
{
  "foo.com": "/home/you/sites/foo.com",
  "*.bar.net": "localhost:8080",
  "qux.org": {
    "/static": "/home/you/qux.org/static",
    "/favicon.ico": "/home/you/qux.org/static/favicon.ico",
    "*": "localhost:3000"
  }
}
```

And start the server:

```bash
$ boxcars config.json
```

To specify the port:

```bash
$ boxcars -port=8001 config.json
```

You can change the configuration anytime during boxcars running. 
It'll be watching your file and reloading it only if it parses with no error.

## Configuration Examples

I use below configuration for a static single-page app that connects to an HTTP API:

```json
{
  "singlepage.com": {
    "/api": "localhost:1337",
    "*": "sites/singlepage.com"
  }
}
```

To catch any domain:

```json
{
  "foo.com": "localhost:1234",
  "*": "/home/you/404.html"
}
```

To set a custom 404 page for a static server:

```json
{
  "foo.com": {
    "/": "/home/you/sites/foo.com",
    "*": "/home/you/404.html"
  }
}
```

## Logging

Boxcars uses [debug](http://github.com/azer/debug) for logging. To enable logging for specific modules: 

```bash
$ DEBUG=server,sites boxcars config.json
```

To silentize:

```bash
$ DEBUG=. boxcars config.json
```

To enable very verbose mode (not recommended):

```bash
$ DEBUG=* boxcars config.json
```

It'll be outputting to stderr.

## Benchmarks

* [Nginx VS Boxcars](https://gist.github.com/azer/5955772)

## TODO

* Add -daemon option.

![](http://i.cloudup.com/rH_0UwNYg1.jpg)
