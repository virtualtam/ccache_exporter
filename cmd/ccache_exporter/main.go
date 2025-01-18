// Copyright 2018-2022 VirtualTam.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package main

import (
	"flag"
	"net/http"
	"time"

	"github.com/justinas/alice"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/hlog"
	"github.com/rs/zerolog/log"

	"github.com/virtualtam/ccache_exporter/pkg/ccache"
)

const (
	DefaultListenAddr = "0.0.0.0:9508"

	webroot = `<html>
<head><title>ccache exporter</title></head>
<body>
<h1>ccache exporter</h1>
<p><a href='/metrics'>Metrics</a></p>
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
		Msg("Request")
}

func main() {
	listenAddr := flag.String("listenAddr", DefaultListenAddr, "Listen on this address")
	ccacheBinaryPath := flag.String("ccacheBinaryPath", ccache.DefaultBinaryPath, "Path to the ccache binary")
	flag.Parse()

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	wrapper, err := ccache.NewLocalCommand(*ccacheBinaryPath)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to instantiate ccache wrapper")
	}
	collector := NewCollector(wrapper)

	prometheus.MustRegister(collector)

	router := http.NewServeMux()

	router.Handle("/metrics", promhttp.Handler())
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte(webroot))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	// Structured logging
	chain := alice.New(hlog.NewHandler(log.Logger), hlog.AccessHandler(accessLogger))

	server := &http.Server{
		Addr:         *listenAddr,
		Handler:      chain.Then(router),
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	log.Info().Msgf("Listening to http://%s", *listenAddr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal().Err(err).Msg("ListenAndServe")
	}
}
