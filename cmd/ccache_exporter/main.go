// Copyright 2018 VirtualTam.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	ccache "github.com/virtualtam/ccache_exporter"
)

const (
	DefaultListenAddr = ":9508"

	webroot = `<html>
<head><title>ccache exporter</title></head>
<body>
<h1>ccache exporter</h1>
<p><a href='/metrics'>Metrics</a></p>
</body>
</html>`
)

func main() {
	listenAddr := flag.String("listenAddr", DefaultListenAddr, "Listen on this address")
	ccacheBinaryPath := flag.String("ccacheBinaryPath", ccache.DefaultBinaryPath, "Path to the ccache binary")
	flag.Parse()

	wrapper, err := ccache.NewBinaryWrapper(*ccacheBinaryPath)
	if err != nil {
		log.Fatal(err)
	}
	collector := ccache.NewCollector(wrapper)

	prometheus.MustRegister(collector)

	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte(webroot))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	log.Println("Listening on", *listenAddr)
	log.Fatal(http.ListenAndServe(*listenAddr, nil))
}
