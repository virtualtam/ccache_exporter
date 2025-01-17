# A common Makefile that includes rules to be reused in different prometheus projects.
# https://github.com/prometheus/prometheus/blob/master/Makefile.common
include Makefile.common

all: lint cover build
.PHONY: all

clean:
	rm -rf .build .tarballs ccache_exporter ccacheparser
.PHONY: clean

lint:
	golangci-lint run ./...
.PHONY: lint

test:
	go test -race ./...
.PHONY: test

cover:
	go test -cover -race ./...
.PHONY: cover

release: clean
	promu crossbuild
	promu crossbuild tarballs
	cd .tarballs; sha256sum * > sha256sums
.PHONY: release

# Live development server
live:
	@echo "== Watching for changes... (hit Ctrl+C when done)"
	@watchexec --restart --exts go -- go run ./cmd/ccache_exporter/
.PHONY: live
