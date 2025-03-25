// Copyright (c) VirtualTam
// SPDX-License-Identifier: MIT

package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/rs/zerolog/log"

	"github.com/virtualtam/ccache_exporter/v4/pkg/ccache"
)

const (
	namespace = "ccache"
)

var (
	parsingErrors = prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: "collector",
			Name:      "parsing_errors_total",
			Help:      "Collector parsing errors (total)",
		},
	)
)

func init() {
	prometheus.MustRegister(parsingErrors)
}

type collector struct {
	wrapper *ccache.Wrapper

	// ccache metrics
	call                     *prometheus.Desc
	callHit                  *prometheus.Desc
	cacheHitRatio            *prometheus.Desc
	calledForLink            *prometheus.Desc
	calledForPreprocessing   *prometheus.Desc
	compilationFailed        *prometheus.Desc
	preprocessingFailed      *prometheus.Desc
	unsupportedCodeDirective *prometheus.Desc
	noInputFile              *prometheus.Desc
	cleanupsPerformed        *prometheus.Desc
	filesInCache             *prometheus.Desc
	cacheSizeBytes           *prometheus.Desc
	maxCacheSizeBytes        *prometheus.Desc
	remoteStorageError       *prometheus.Desc
	remoteStorageHit         *prometheus.Desc
	remoteStorageMiss        *prometheus.Desc
	remoteStorageReadHit     *prometheus.Desc
	remoteStorageReadMiss    *prometheus.Desc
	remoteStorageTimeout     *prometheus.Desc
	remoteStorageWrite       *prometheus.Desc
	version                  *prometheus.Desc
}

// newCcacheCollector initializes and returns a Prometheus collector for ccache
// metrics.
func newCcacheCollector(ccacheWrapper *ccache.Wrapper) *collector {
	return &collector{
		wrapper: ccacheWrapper,
		call: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "call_total"),
			"Cache calls (total)",
			nil,
			nil,
		),
		callHit: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "call_hit_total"),
			"Cache hits",
			[]string{"mode"},
			nil,
		),
		cacheHitRatio: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "cache_hit_ratio"),
			"Cache hit ratio (direct + preprocessed) / miss",
			nil,
			nil,
		),
		calledForLink: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "called_for_link_total"),
			"Called for link",
			nil,
			nil,
		),
		calledForPreprocessing: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "called_for_preprocessing_total"),
			"Called for preprocessing",
			nil,
			nil,
		),
		compilationFailed: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "compilation_failed_total"),
			"Compilation failed",
			nil,
			nil,
		),
		preprocessingFailed: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "preprocessing_failed_total"),
			"Preprocessing failed",
			nil,
			nil,
		),
		unsupportedCodeDirective: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "unsupported_code_directive_total"),
			"Unsupported code directive",
			nil,
			nil,
		),
		noInputFile: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "no_input_file_total"),
			"No input file",
			nil,
			nil,
		),
		cleanupsPerformed: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "cleanups_performed_total"),
			"Cleanups performed",
			nil,
			nil,
		),
		filesInCache: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "cached_files"),
			"Cached files",
			nil,
			nil,
		),
		cacheSizeBytes: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "cache_size_bytes"),
			"Cache size (bytes)",
			nil,
			nil,
		),
		maxCacheSizeBytes: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "cache_size_max_bytes"),
			"Maximum cache size (bytes)",
			nil,
			nil,
		),
		remoteStorageError: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "remote_storage_errors_total"),
			"Remote storage errors",
			nil,
			nil,
		),
		remoteStorageHit: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "remote_storage_hit_total"),
			"Remote storage hits",
			nil,
			nil,
		),
		remoteStorageMiss: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "remote_storage_miss_total"),
			"Remote storage misses",
			nil,
			nil,
		),
		remoteStorageReadHit: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "remote_storage_read_hit_total"),
			"Remote storage read hits",
			nil,
			nil,
		),
		remoteStorageReadMiss: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "remote_storage_read_miss_total"),
			"Remote storage read miss",
			nil,
			nil,
		),
		remoteStorageTimeout: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "remote_storage_timeout_total"),
			"Remote storage timeouts",
			nil,
			nil,
		),
		remoteStorageWrite: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "remote_storage_write_total"),
			"Remote storage writes",
			nil,
			nil,
		),
		version: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "version"),
			"ccache version",
			[]string{"version"},
			nil,
		),
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
	ch <- c.compilationFailed
	ch <- c.preprocessingFailed
	ch <- c.unsupportedCodeDirective
	ch <- c.noInputFile
	ch <- c.cleanupsPerformed
	ch <- c.filesInCache
	ch <- c.cacheSizeBytes
	ch <- c.maxCacheSizeBytes
	ch <- c.remoteStorageError
	ch <- c.remoteStorageHit
	ch <- c.remoteStorageMiss
	ch <- c.remoteStorageReadHit
	ch <- c.remoteStorageReadMiss
	ch <- c.remoteStorageTimeout
	ch <- c.remoteStorageWrite
	ch <- c.version
}

// Collect gathers metrics from ccache.
func (c *collector) Collect(ch chan<- prometheus.Metric) {
	config, err := c.wrapper.Configuration()
	if err != nil {
		log.Error().Err(err).Msg("ccache: failed to collect configuration")
		parsingErrors.Inc()
		return
	}

	stats, err := c.wrapper.Statistics()
	if err != nil {
		log.Error().Err(err).Msg("ccache: failed to collect statistics")
		parsingErrors.Inc()
		return
	}

	// counters
	ch <- prometheus.MustNewConstMetric(
		c.call,
		prometheus.CounterValue,
		float64(stats.CacheHitDirect+stats.CacheHitPreprocessed+stats.CacheMiss),
	)
	ch <- prometheus.MustNewConstMetric(c.callHit, prometheus.CounterValue, float64(stats.CacheHitDirect), "direct")
	ch <- prometheus.MustNewConstMetric(c.callHit, prometheus.CounterValue, float64(stats.CacheHitPreprocessed), "preprocessed")
	ch <- prometheus.MustNewConstMetric(c.calledForLink, prometheus.CounterValue, float64(stats.CalledForLink))
	ch <- prometheus.MustNewConstMetric(c.calledForPreprocessing, prometheus.CounterValue, float64(stats.CalledForPreprocessing))
	ch <- prometheus.MustNewConstMetric(c.compilationFailed, prometheus.CounterValue, float64(stats.CompilationFailed))
	ch <- prometheus.MustNewConstMetric(c.preprocessingFailed, prometheus.CounterValue, float64(stats.PreprocessingFailed))
	ch <- prometheus.MustNewConstMetric(c.unsupportedCodeDirective, prometheus.CounterValue, float64(stats.UnsupportedCodeDirective))
	ch <- prometheus.MustNewConstMetric(c.noInputFile, prometheus.CounterValue, float64(stats.NoInputFile))
	ch <- prometheus.MustNewConstMetric(c.cleanupsPerformed, prometheus.CounterValue, float64(stats.CleanupsPerformed))
	ch <- prometheus.MustNewConstMetric(c.remoteStorageError, prometheus.CounterValue, float64(stats.RemoteStorageError))
	ch <- prometheus.MustNewConstMetric(c.remoteStorageHit, prometheus.CounterValue, float64(stats.RemoteStorageHit))
	ch <- prometheus.MustNewConstMetric(c.remoteStorageMiss, prometheus.CounterValue, float64(stats.RemoteStorageMiss))
	ch <- prometheus.MustNewConstMetric(c.remoteStorageReadHit, prometheus.CounterValue, float64(stats.RemoteStorageReadHit))
	ch <- prometheus.MustNewConstMetric(c.remoteStorageReadMiss, prometheus.CounterValue, float64(stats.RemoteStorageReadMiss))
	ch <- prometheus.MustNewConstMetric(c.remoteStorageTimeout, prometheus.CounterValue, float64(stats.RemoteStorageTimeout))
	ch <- prometheus.MustNewConstMetric(c.remoteStorageWrite, prometheus.CounterValue, float64(stats.RemoteStorageWrite))

	// gauges
	ch <- prometheus.MustNewConstMetric(c.cacheHitRatio, prometheus.GaugeValue, stats.CacheHitRatio)
	ch <- prometheus.MustNewConstMetric(c.filesInCache, prometheus.GaugeValue, float64(stats.FilesInCache))
	ch <- prometheus.MustNewConstMetric(c.cacheSizeBytes, prometheus.GaugeValue, float64(stats.CacheSizeBytes))
	ch <- prometheus.MustNewConstMetric(c.maxCacheSizeBytes, prometheus.GaugeValue, float64(config.MaxCacheSizeBytes))

	// version
	ch <- prometheus.MustNewConstMetric(c.version, prometheus.UntypedValue, 1, c.wrapper.Version())
}
