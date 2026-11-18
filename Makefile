BUILD_DIR ?= build
SRC_FILES := $(shell find . -name "*.go")
PROMU_VERSION ?= v0.21.0

all: lint race cover build
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

cover:
	go test -coverprofile=coverage.out ./...
.PHONY: cover

coverhtml: cover
	go tool cover -html=coverage.out
.PHONY: coverhtml

race:
	go test -race ./...
.PHONY: race

test:
	go test ./...
.PHONY: test

release: clean
	promu crossbuild
	promu crossbuild tarballs
	cd .tarballs; sha256sum * > sha256sums
.PHONY: release

# Install development tools
dev-install-tools: install-promu
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v2.15.0
	go install github.com/hashicorp/copywrite@latest
	go install golang.org/x/tools/gopls@latest
	go install golang.org/x/vuln/cmd/govulncheck@latest
	uv tool install yamllint
.PHONY: dev-install-tools

install-promu:
	go install github.com/prometheus/promu@$(PROMU_VERSION)
.PHONY: install-promu

# Modernize
modernize:
	go run golang.org/x/tools/go/analysis/passes/modernize/cmd/modernize@latest -fix -test ./...
.PHONY: modernize

# Vulnerability check
vulncheck:
	govulncheck -C . ./...
.PHONY: vulncheck

# Licence headers
copywrite:
	copywrite headers
.PHONY: copywrite

# Render the example Grafana dashboard as a PNG image.
DASHBOARD_FROM ?= now-1h
DASHBOARD_TO ?= now

render-example-dashboard:
	curl -sSf -o dashboard.png \
		"http://localhost:13000/render/d/YytQhHJiz/ccache?from=$(DASHBOARD_FROM)&to=$(DASHBOARD_TO)&timezone=browser&width=1600&height=-1&kiosk=true&theme=dark&timeout=60"
.PHONY: render-example-dashboard

# Live development server
live:
	@echo "== Watching for changes... (hit Ctrl+C when done)"
	@watchexec --restart --exts go -- go run ./cmd/ccache_exporter/ run
.PHONY: live
