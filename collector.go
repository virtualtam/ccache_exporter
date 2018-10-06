// Package ccache_exporter implements a Prometheus exporter for ccache metrics
//
// See:
// - https://ccache.samba.org/
// - https://prometheus.io/
package ccache_exporter

import (
	"log"
	"os/exec"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/virtualtam/ccacheparser"
)

const (
	namespace = "ccache"
)

type ccacheCollector struct {
	cacheHit                 *prometheus.Desc
	cacheMiss                *prometheus.Desc
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

func NewCcacheCollector() *ccacheCollector {
	return &ccacheCollector{
		cacheHit: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "cache_hit_total"),
			"Cache hit",
			[]string{"mode"},
			nil,
		),
		cacheMiss: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "cache_miss_total"),
			"Cache miss",
			nil,
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
			prometheus.BuildFQName(namespace, "", "files_in_cache"),
			"Files in cache",
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
			prometheus.BuildFQName(namespace, "", "max_cache_size_bytes"),
			"Maximum cache size (bytes)",
			nil,
			nil,
		),
	}
}

func (c *ccacheCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.cacheHit
	ch <- c.cacheMiss
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

func (c *ccacheCollector) Collect(ch chan<- prometheus.Metric) {
	out, err := exec.Command("ccache", "-s").Output()
	if err != nil {
		log.Fatal(err)
	}

	stats := ccacheparser.Statistics{}
	stats.Parse(string(out[:]))

	// counters
	ch <- prometheus.MustNewConstMetric(c.cacheHit, prometheus.CounterValue, float64(stats.CacheHitDirect), "direct")
	ch <- prometheus.MustNewConstMetric(c.cacheHit, prometheus.CounterValue, float64(stats.CacheHitPreprocessed), "preprocessed")
	ch <- prometheus.MustNewConstMetric(c.cacheMiss, prometheus.CounterValue, float64(stats.CacheMiss))
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
