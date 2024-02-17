# generate-testdata

## Overview
This directory contains utilities to generate `ccache` testdata that can be used as
reference input in Go unitary tests, to ensure `ccache` statistics are properly parsed.

## Resources

The `Makefile` provides useful targets to:

- build Docker images for Debian and Ubuntu, that provide all necessary packages
  to compile software with `ccache` enabled
- compile software with `ccache` and save cache usage statistics for:
    - an empty cache
    - a cache populated by a first build
    - a cache populated and reused by a second build

## Example: Debian 12 testdata
Build the Docker image:

```shell
make docker-debian-12
```

Generate testdata:

```shell
make ccache-testdata-debian-12
```

## Example: Ubuntu 24.04 testdata with Redis remote storage enabled
Build the Docker image:

```shell
make docker-ubuntu-24.04
```

Generate testdata:

```shell
./generate-testdata local/ccache ubuntu-24.04 ccache-testdata-ubuntu-24.04 yes
```
