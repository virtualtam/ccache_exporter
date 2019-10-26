# Ensure that 'all' is the default target otherwise it will be the first target from Makefile.common.
all::

# Needs to be defined before including Makefile.common to auto-generate targets
DOCKER_ARCHS ?= amd64

# A common Makefile that includes rules to be reused in different prometheus projects.
# https://github.com/prometheus/prometheus/blob/master/Makefile.common
include Makefile.common

cover:
	go test -cover -race ./...
.PHONY: cover
