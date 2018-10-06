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
	cacheHitDirect           *prometheus.Desc
	cacheHitPreprocessed     *prometheus.Desc
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
		cacheHitDirect: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "cache_hit_direct"),
			"Cache hit (direct)",
			nil,
			nil,
		),
		cacheHitPreprocessed: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "cache_hit_preprocessed"),
			"Cache hit (preprocessed)",
			nil,
			nil,
		),
		cacheMiss: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "cache_miss"),
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
			prometheus.BuildFQName(namespace, "", "called_for_link"),
			"Called for link",
			nil,
			nil,
		),
		calledForPreprocessing: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "called_for_preprocessing"),
			"Called for preprocessing",
			nil,
			nil,
		),
		unsupportedCodeDirective: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "unsupported_code_directive"),
			"Unsupported code directive",
			nil,
			nil,
		),
		noInputFile: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "no_input_file"),
			"No input file",
			nil,
			nil,
		),
		cleanupsPerformed: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "cleanups_performed"),
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
	ch <- c.cacheHitDirect
	ch <- c.cacheHitPreprocessed
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

	ch <- prometheus.MustNewConstMetric(c.cacheHitDirect, prometheus.CounterValue, float64(stats.CacheHitDirect))
	ch <- prometheus.MustNewConstMetric(c.cacheHitPreprocessed, prometheus.CounterValue, float64(stats.CacheHitPreprocessed))
	ch <- prometheus.MustNewConstMetric(c.cacheMiss, prometheus.CounterValue, float64(stats.CacheMiss))
	ch <- prometheus.MustNewConstMetric(c.cacheHitRatio, prometheus.GaugeValue, float64(stats.CacheHitRatio))
	ch <- prometheus.MustNewConstMetric(c.calledForLink, prometheus.CounterValue, float64(stats.CalledForLink))
	ch <- prometheus.MustNewConstMetric(c.calledForPreprocessing, prometheus.CounterValue, float64(stats.CalledForPreprocessing))
	ch <- prometheus.MustNewConstMetric(c.unsupportedCodeDirective, prometheus.CounterValue, float64(stats.UnsupportedCodeDirective))
	ch <- prometheus.MustNewConstMetric(c.noInputFile, prometheus.CounterValue, float64(stats.NoInputFile))
	ch <- prometheus.MustNewConstMetric(c.cleanupsPerformed, prometheus.CounterValue, float64(stats.CleanupsPerformed))
	ch <- prometheus.MustNewConstMetric(c.filesInCache, prometheus.CounterValue, float64(stats.FilesInCache))
	ch <- prometheus.MustNewConstMetric(c.cacheSizeBytes, prometheus.CounterValue, float64(stats.CacheSizeBytes))
	ch <- prometheus.MustNewConstMetric(c.maxCacheSizeBytes, prometheus.CounterValue, float64(stats.MaxCacheSizeBytes))
}
