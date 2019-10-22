BUILD_DIR := build

build: build-ccache_exporter build-ccacheparser
.PHONY: build

build-%:
	go build -o $(BUILD_DIR)/$* ./cmd/$*

distclean:
	rm -rf build

test:
	go test ./...

cover:
	go test -cover ./...
