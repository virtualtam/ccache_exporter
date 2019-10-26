// Prometheus metrics collection.
//
// See:
// - https://ccache.samba.org/
// - https://prometheus.io/
package ccache

import (
	"log"
	"os/exec"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	namespace = "ccache"
)

type collector struct {
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
}

func NewCollector() *collector {
	return &collector{
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
	}
}

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

func (c *collector) Collect(ch chan<- prometheus.Metric) {
	out, err := exec.Command("ccache", "-s").Output()
	if err != nil {
		log.Fatal(err)
	}

	stats := Statistics{}
	stats.Parse(string(out[:]))

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
	ch <- prometheus.MustNewConstMetric(c.unsupportedCodeDirective, prometheus.CounterValue, float64(stats.UnsupportedCodeDirective))
	ch <- prometheus.MustNewConstMetric(c.noInputFile, prometheus.CounterValue, float64(stats.NoInputFile))
	ch <- prometheus.MustNewConstMetric(c.cleanupsPerformed, prometheus.CounterValue, float64(stats.CleanupsPerformed))

	// gauges
	ch <- prometheus.MustNewConstMetric(c.cacheHitRatio, prometheus.GaugeValue, stats.CacheHitRatio)
	ch <- prometheus.MustNewConstMetric(c.filesInCache, prometheus.GaugeValue, float64(stats.FilesInCache))
	ch <- prometheus.MustNewConstMetric(c.cacheSizeBytes, prometheus.GaugeValue, float64(stats.CacheSizeBytes))
	ch <- prometheus.MustNewConstMetric(c.maxCacheSizeBytes, prometheus.GaugeValue, float64(stats.MaxCacheSizeBytes))
}
