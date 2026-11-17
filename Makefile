BUILD_DIR ?= build
SRC_FILES := $(shell find . -name "*.go")
PROMU_VERSION ?= v0.21.0

all: lint cover build
.PHONY: all

build: $(BUILD_DIR)/ccache_exporter $(BUILD_DIR)/ccacheparser
.PHONY: build

$(BUILD_DIR)/%: $(SRC_FILES) go.mod go.sum
	go build -trimpath -ldflags=-s -o $@ ./cmd/$*

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
install-tools: install-promu
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v2.15.0
	go install github.com/hashicorp/copywrite@latest
	go install golang.org/x/vuln/cmd/govulncheck@latest
.PHONY: install-tools

install-promu:
	go install github.com/prometheus/promu@$(PROMU_VERSION)
.PHONY: install-promu

# Licence headers
copywrite:
	copywrite headers
.PHONY: copywrite

# Live development server
live:
	@echo "== Watching for changes... (hit Ctrl+C when done)"
	@watchexec --restart --exts go -- go run ./cmd/ccache_exporter/ run
.PHONY: live
