# A common Makefile that includes rules to be reused in different prometheus projects.
# https://github.com/prometheus/prometheus/blob/master/Makefile.common
include Makefile.common

all: lint cover build
.PHONY: all

clean:
	rm -rf .build .tarballs build
.PHONY: clean

lint: lint-go lint-yaml
.PHONY: lint

lint-go:
	golangci-lint run ./...
.PHONY: lint-go

lint-yaml:
	yamllint .
	yamllint -c .github/.yamllint.yml .github/workflows
.PHONY: lint-yaml

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

# Tools
install-tools: promu install-govulncheck
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v2.15.0
	go install github.com/hashicorp/copywrite@latest
.PHONY: install-tools

# Licence headers
copywrite:
	copywrite headers
.PHONY: copywrite

# Live development server
live:
	@echo "== Watching for changes... (hit Ctrl+C when done)"
	@watchexec --restart --exts go -- go run ./cmd/ccache_exporter/ run
.PHONY: live
