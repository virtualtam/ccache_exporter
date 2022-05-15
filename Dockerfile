# Step 1: build Go binaries
FROM golang:1.18-alpine as builder

ARG CGO_ENABLED=1

RUN apk add --update --no-cache \
        ca-certificates \
        gcc \
        git \
        musl-dev

WORKDIR /app
ADD . .
RUN go build -trimpath ./cmd/ccache_exporter 2>&1

# Step 2: build the actual image
FROM alpine:3.14

RUN apk add --update --no-cache ccache \
    && adduser -D exporter

COPY --from=builder /app/ccache_exporter /usr/local/bin

USER exporter

EXPOSE 9508

VOLUME /home/exporter/.ccache

CMD ["ccache_exporter"]
