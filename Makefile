BUILD_DIR := build

build: build-ccache_exporter
.PHONY: build

build-%:
	go build -o $(BUILD_DIR)/$* ./cmd/$*

distclean:
	rm -rf build
