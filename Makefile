# A common Makefile that includes rules to be reused in different prometheus projects.
# https://github.com/prometheus/prometheus/blob/master/Makefile.common
include Makefile.common

BUILD_DIR ?= build
SRC_FILES := $(shell find . -name "*.go")

all: lint cover build
.PHONY: all

clean:
	rm -rf .build $(BUILD_DIR)
.PHONY: clean

$(BUILD_DIR)/%: $(SRC_FILES)
	go build -trimpath -o $@ ./cmd/$*

lint:
	golangci-lint run ./...
.PHONY: lint

test:
	go test -race ./...
.PHONY: test

cover:
	go test -cover -race ./...
.PHONY: cover

release:
	promu crossbuild
	promu crossbuild tarballs
.phony: release
