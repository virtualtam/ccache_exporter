// Copyright VirtualTam 2018, 2026
// SPDX-License-Identifier: MIT

package metrics

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/rs/zerolog/log"

	"github.com/virtualtam/ccache_exporter/v4/pkg/ccache"
)

const (
	namespace   = "ccache"
	reasonLabel = "reason"
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
	ctx     context.Context
	wrapper *ccache.Wrapper

	// ccache metrics
	call                  *prometheus.Desc
	callHit               *prometheus.Desc
	cacheHitRatio         *prometheus.Desc
	uncacheableCalls      *prometheus.Desc
	errorCalls            *prometheus.Desc
	cleanupsPerformed     *prometheus.Desc
	filesInCache          *prometheus.Desc
	cacheSizeBytes        *prometheus.Desc
	maxCacheSizeBytes     *prometheus.Desc
	remoteStorageError    *prometheus.Desc
	remoteStorageHit      *prometheus.Desc
	remoteStorageMiss     *prometheus.Desc
	remoteStorageReadHit  *prometheus.Desc
	remoteStorageReadMiss *prometheus.Desc
	remoteStorageTimeout  *prometheus.Desc
	remoteStorageWrite    *prometheus.Desc
	version               *prometheus.Desc
}

// newCcacheCollector initializes and returns a Prometheus collector for ccache
// metrics.
func newCcacheCollector(ctx context.Context, ccacheWrapper *ccache.Wrapper) *collector {
	return &collector{
		ctx:     ctx,
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
			"Cache hit ratio: hits / (hits + misses), for cacheable calls",
			nil,
			nil,
		),
		uncacheableCalls: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "uncacheable_calls_total"),
			"Calls ccache could not cache (total), by reason",
			[]string{reasonLabel},
			nil,
		),
		errorCalls: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "error_calls_total"),
			"Calls ccache could not cache due to an error (total), by reason",
			[]string{reasonLabel},
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
	ch <- c.uncacheableCalls
	ch <- c.errorCalls
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
	config, err := c.wrapper.Configuration(c.ctx)
	if err != nil {
		log.Error().Err(err).Msg("ccache: failed to collect configuration")
		parsingErrors.Inc()
		return
	}

	stats, err := c.wrapper.Statistics(c.ctx)
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
	for reason, value := range uncacheableCallsByReason(stats.UncacheableCalls) {
		ch <- prometheus.MustNewConstMetric(c.uncacheableCalls, prometheus.CounterValue, float64(value), reason)
	}
	for reason, value := range errorCallsByReason(stats.ErrorCalls) {
		ch <- prometheus.MustNewConstMetric(c.errorCalls, prometheus.CounterValue, float64(value), reason)
	}
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

// uncacheableCallsByReason maps uncacheable call counters to ccache statistic names.
func uncacheableCallsByReason(calls ccache.UncacheableCalls) map[string]int {
	return map[string]int{
		"autoconf_test":                    calls.AutoconfTest,
		"bad_compiler_arguments":           calls.BadCompilerArguments,
		"called_for_link":                  calls.CalledForLink,
		"called_for_preprocessing":         calls.CalledForPreprocessing,
		"compile_failed":                   calls.CompileFailed,
		"compiler_produced_empty_output":   calls.CompilerProducedEmptyOutput,
		"compiler_produced_no_output":      calls.CompilerProducedNoOutput,
		"compiler_produced_stdout":         calls.CompilerProducedStdout,
		"could_not_use_modules":            calls.CouldNotUseModules,
		"could_not_use_precompiled_header": calls.CouldNotUsePrecompiledHeader,
		"disabled":                         calls.Disabled,
		"multiple_source_files":            calls.MultipleSourceFiles,
		"no_input_file":                    calls.NoInputFile,
		"output_to_stdout":                 calls.OutputToStdout,
		"preprocessor_error":               calls.PreprocessorError,
		"recache":                          calls.Recache,
		"unsupported_code_directive":       calls.UnsupportedCodeDirective,
		"unsupported_compiler_option":      calls.UnsupportedCompilerOption,
		"unsupported_environment_variable": calls.UnsupportedEnvironmentVariable,
		"unsupported_source_encoding":      calls.UnsupportedSourceEncoding,
		"unsupported_source_language":      calls.UnsupportedSourceLanguage,
	}
}

// errorCallsByReason maps error call counters to ccache statistic names.
func errorCallsByReason(calls ccache.ErrorCalls) map[string]int {
	return map[string]int{
		"bad_input_file":           calls.BadInputFile,
		"bad_output_file":          calls.BadOutputFile,
		"compiler_check_failed":    calls.CompilerCheckFailed,
		"could_not_find_compiler":  calls.CouldNotFindCompiler,
		"error_hashing_extra_file": calls.ErrorHashingExtraFile,
		"internal_error":           calls.InternalError,
		"missing_cache_file":       calls.MissingCacheFile,
		"modified_input_file":      calls.ModifiedInputFile,
	}
}
