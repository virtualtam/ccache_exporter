// Copyright (c) VirtualTam
// SPDX-License-Identifier: MIT

package metrics

import (
	"net/http"
	"time"

	"github.com/justinas/alice"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog/hlog"
	"github.com/rs/zerolog/log"

	"github.com/virtualtam/ccache_exporter/pkg/ccache"
)

const (
	webroot = `<html>
  <head><title>ccache exporter</title></head>
  <body>
    <h1>ccache exporter</h1>
    <p><a href="/metrics">Metrics</a></p>
  </body>
</html>`
)

func accessLogger(r *http.Request, status, size int, dur time.Duration) {
	hlog.FromRequest(r).Info().
		Dur("duration_ms", dur).
		Str("host", r.Host).
		Str("path", r.URL.Path).
		Int("size", size).
		Int("status", status).
		Msg("handle request")
}

// NewServer registers metrics collectors and returns a HTTP server to expose them.
func NewServer(wrapper *ccache.LocalCommand, listenAddr string) *http.Server {
	ccacheCollector := NewCcacheCollector(wrapper)
	prometheus.MustRegister(ccacheCollector)

	router := http.NewServeMux()

	router.Handle("/metrics", promhttp.Handler())
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte(webroot))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	// Setup structured logging middleware
	chain := alice.New(hlog.NewHandler(log.Logger), hlog.AccessHandler(accessLogger))

	server := &http.Server{
		Addr:         listenAddr,
		Handler:      chain.Then(router),
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	return server
}
