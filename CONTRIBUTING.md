# Contributing

## Getting the source code

Get the sources:

```shell
$ git clone https://github.com/virtualtam/ccache_exporter.git
$ cd ccache_exporter
```

## Building

Run linters:
```shell
$ make lint
```

Build the parser and exporter:

```shell
$ make build
```

Build platform-specific binaries with [Promu](https://github.com/prometheus/promu):

```shell
$ promu crossbuild
```

Build and archive platform-specific binaries:

```shell
$ promu crossbuild
$ promu crossbuild tarballs
```
