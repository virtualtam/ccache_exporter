// Copyright 2018 VirtualTam.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package ccache

import (
	"log"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	namespace = "ccache"
)

type collector struct {
	wrapper Wrapper

	call                     *prometheus.Desc
	callHit                  *prometheus.Desc
	cacheHitRatio            *prometheus.Desc
	calledForLink            *prometheus.Desc
	calledForPreprocessing   *prometheus.Desc
	unsupportedCodeDirective *prometheus.Desc
	noInputFile              *prometheus.Desc
	cleanupsPerformed        *prometheus.Desc
	filesInCache             *prometheus.Desc
	cacheSizeBytes           *prometheus.Desc
	maxCacheSizeBytes        *prometheus.Desc
	cacheDirs []string
}

// NewCollector initializes and returns a Prometheus collector for ccache
// metrics.
func NewCollector(w Wrapper, cacheDirs []string) *collector {
	var (
		labels   = []string{"cachedir"}
	)
	return &collector{
		wrapper: w,
		call: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "call_total"),
			"Cache calls (total)",
			labels,
			nil,
		),
		callHit: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "call_hit_total"),
			"Cache hits",
			append(labels, "mode"),
			nil,
		),
		cacheHitRatio: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "cache_hit_ratio"),
			"Cache hit ratio (direct + preprocessed) / miss",
			labels,
			nil,
		),
		calledForLink: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "called_for_link_total"),
			"Called for link",
			labels,
			nil,
		),
		calledForPreprocessing: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "called_for_preprocessing_total"),
			"Called for preprocessing",
			labels,
			nil,
		),
		unsupportedCodeDirective: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "unsupported_code_directive_total"),
			"Unsupported code directive",
			labels,
			nil,
		),
		noInputFile: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "no_input_file_total"),
			"No input file",
			labels,
			nil,
		),
		cleanupsPerformed: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "cleanups_performed_total"),
			"Cleanups performed",
			labels,
			nil,
		),
		filesInCache: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "cached_files"),
			"Cached files",
			labels,
			nil,
		),
		cacheSizeBytes: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "cache_size_bytes"),
			"Cache size (bytes)",
			labels,
			nil,
		),
		maxCacheSizeBytes: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "cache_size_max_bytes"),
			"Maximum cache size (bytes)",
			labels,
			nil,
		),
		cacheDirs: cacheDirs,
	}
}

// Describe publishes the description of each ccache metric to a metrics
// channel.
func (c *collector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.call
	ch <- c.callHit
	ch <- c.cacheHitRatio
	ch <- c.calledForLink
	ch <- c.calledForPreprocessing
	ch <- c.unsupportedCodeDirective
	ch <- c.noInputFile
	ch <- c.cleanupsPerformed
	ch <- c.filesInCache
	ch <- c.cacheSizeBytes
	ch <- c.maxCacheSizeBytes
}

// Collect gathers metrics from ccache.
func (c *collector) Collect(ch chan<- prometheus.Metric) {
	for _, s := range c.cacheDirs {
		out, err := c.wrapper.ShowStats(s)
		if err != nil {
			log.Fatal(err)
		}

		stats := Statistics{}
		stats.Parse(out)
		
		// counters
		ch <- prometheus.MustNewConstMetric(
			c.call,
			prometheus.CounterValue,
			float64(stats.CacheHitDirect+stats.CacheHitPreprocessed+stats.CacheMiss), stats.CacheDirectory,
		)
		ch <- prometheus.MustNewConstMetric(c.callHit, prometheus.CounterValue, float64(stats.CacheHitDirect), stats.CacheDirectory, "direct")
		ch <- prometheus.MustNewConstMetric(c.callHit, prometheus.CounterValue, float64(stats.CacheHitPreprocessed), stats.CacheDirectory, "preprocessed")
		ch <- prometheus.MustNewConstMetric(c.calledForLink, prometheus.CounterValue, float64(stats.CalledForLink), stats.CacheDirectory)
		ch <- prometheus.MustNewConstMetric(c.calledForPreprocessing, prometheus.CounterValue, float64(stats.CalledForPreprocessing), stats.CacheDirectory)
		ch <- prometheus.MustNewConstMetric(c.unsupportedCodeDirective, prometheus.CounterValue, float64(stats.UnsupportedCodeDirective), stats.CacheDirectory)
		ch <- prometheus.MustNewConstMetric(c.noInputFile, prometheus.CounterValue, float64(stats.NoInputFile), stats.CacheDirectory)
		ch <- prometheus.MustNewConstMetric(c.cleanupsPerformed, prometheus.CounterValue, float64(stats.CleanupsPerformed), stats.CacheDirectory)

		// gauges
		ch <- prometheus.MustNewConstMetric(c.cacheHitRatio, prometheus.GaugeValue, stats.CacheHitRatio, stats.CacheDirectory)
		ch <- prometheus.MustNewConstMetric(c.filesInCache, prometheus.GaugeValue, float64(stats.FilesInCache), stats.CacheDirectory)
		ch <- prometheus.MustNewConstMetric(c.cacheSizeBytes, prometheus.GaugeValue, float64(stats.CacheSizeBytes), stats.CacheDirectory)
		ch <- prometheus.MustNewConstMetric(c.maxCacheSizeBytes, prometheus.GaugeValue, float64(stats.MaxCacheSizeBytes), stats.CacheDirectory)
	}
}
