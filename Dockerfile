# Step 1: Build Go binaries
FROM golang:1.23-bookworm AS builder

ARG CGO_ENABLED=1

WORKDIR /app
COPY go.mod go.sum ./
RUN --mount=type=cache,sharing=locked,target=/go/pkg/mod go mod download

ADD . .
RUN make common-build

# Step 2: build the actual image
FROM debian:bookworm-slim

RUN --mount=type=cache,sharing=locked,target=/var/lib/apt/lists \
    --mount=type=cache,sharing=locked,target=/var/cache/apt \
    rm -f /etc/apt/apt.conf.d/docker-clean \
    && echo 'Binary::apt::APT::Keep-Downloaded-Packages "true";' > /etc/apt/apt.conf.d/keep-cache \
    && apt-get update \
    && apt-get install -y ccache

RUN groupadd \
        --gid 1000 \
        exporter \
    && useradd \
        --create-home \
        --home-dir /var/lib/exporter \
        --shell /bin/bash \
        --uid 1000 \
        --gid exporter \
        exporter

COPY --from=builder /app/build/ccache_exporter /usr/local/bin/ccache_exporter

USER exporter
WORKDIR /var/lib/exporter

EXPOSE 9508

VOLUME /var/lib/exporter/.ccache

CMD ["ccache_exporter"]
