package ccache

import (
	"io/ioutil"
	"testing"

	"github.com/alecthomas/units"
	"github.com/stretchr/testify/assert"
)

func TestParseEmptyCacheStats35(t *testing.T) {
	input, err := ioutil.ReadFile("tests/3.5/empty.txt")
	if err != nil {
		panic(err)
	}

	s := Statistics{}
	s.Parse(string(input))

	assert.Equal(t, "/home/virtualtam/.ccache", s.CacheDirectory)
	assert.Equal(t, "/home/virtualtam/.ccache/ccache.conf", s.PrimaryConfig)
	assert.Equal(t, "/etc/ccache.conf", s.SecondaryConfigReadonly)
	assert.Equal(t, 0, s.CacheHitDirect)
	assert.Equal(t, 0, s.CacheHitPreprocessed)
	assert.Equal(t, 0, s.CacheMiss)
	assert.Equal(t, 0.0, s.CacheHitRate)
	assert.Equal(t, 0.0, s.CacheHitRatio)
	assert.Equal(t, 0, s.CalledForLink)
	assert.Equal(t, 0, s.CalledForPreprocessing)
	assert.Equal(t, 0, s.UnsupportedCodeDirective)
	assert.Equal(t, 0, s.NoInputFile)
	assert.Equal(t, 0, s.CleanupsPerformed)
	assert.Equal(t, 0, s.FilesInCache)
	assert.Equal(t, "0.0 kB", s.CacheSize)
	assert.Equal(t, units.MetricBytes(0), s.CacheSizeBytes)
	assert.Equal(t, "5.0 GB", s.MaxCacheSize)
	assert.Equal(t, units.MetricBytes(5000000000), s.MaxCacheSizeBytes)
}

func TestParseFirstBuildStats35(t *testing.T) {
	input, err := ioutil.ReadFile("tests/3.5/first.txt")
	if err != nil {
		panic(err)
	}

	s := Statistics{}
	s.Parse(string(input))

	assert.Equal(t, "/home/virtualtam/.ccache", s.CacheDirectory)
	assert.Equal(t, "/home/virtualtam/.ccache/ccache.conf", s.PrimaryConfig)
	assert.Equal(t, "/etc/ccache.conf", s.SecondaryConfigReadonly)
	assert.Equal(t, 1, s.CacheHitDirect)
	assert.Equal(t, 15, s.CacheHitPreprocessed)
	assert.Equal(t, 342, s.CacheMiss)
	assert.Equal(t, 4.47, s.CacheHitRate)
	assert.Equal(t, 0.0447, s.CacheHitRatio)
	assert.Equal(t, 14, s.CalledForLink)
	assert.Equal(t, 1, s.CalledForPreprocessing)
	assert.Equal(t, 0, s.UnsupportedCodeDirective)
	assert.Equal(t, 0, s.NoInputFile)
	assert.Equal(t, 0, s.CleanupsPerformed)
	assert.Equal(t, 867, s.FilesInCache)
	assert.Equal(t, "44.5 MB", s.CacheSize)
	assert.Equal(t, units.MetricBytes(44500000), s.CacheSizeBytes)
	assert.Equal(t, "5.0 GB", s.MaxCacheSize)
	assert.Equal(t, units.MetricBytes(5000000000), s.MaxCacheSizeBytes)
}

func TestParseSecondBuildStats35(t *testing.T) {
	input, err := ioutil.ReadFile("tests/3.5/second.txt")
	if err != nil {
		panic(err)
	}

	s := Statistics{}
	s.Parse(string(input))

	assert.Equal(t, "/home/virtualtam/.ccache", s.CacheDirectory)
	assert.Equal(t, "/home/virtualtam/.ccache/ccache.conf", s.PrimaryConfig)
	assert.Equal(t, "/etc/ccache.conf", s.SecondaryConfigReadonly)
	assert.Equal(t, 349, s.CacheHitDirect)
	assert.Equal(t, 10, s.CacheHitPreprocessed)
	assert.Equal(t, 28, s.CacheMiss)
	assert.Equal(t, 92.76, s.CacheHitRate)
	assert.Equal(t, 0.9276000000000001, s.CacheHitRatio)
	assert.Equal(t, 14, s.CalledForLink)
	assert.Equal(t, 1, s.CalledForPreprocessing)
	assert.Equal(t, 0, s.UnsupportedCodeDirective)
	assert.Equal(t, 0, s.NoInputFile)
	assert.Equal(t, 0, s.CleanupsPerformed)
	assert.Equal(t, 943, s.FilesInCache)
	assert.Equal(t, "46.7 MB", s.CacheSize)
	assert.Equal(t, units.MetricBytes(46700000), s.CacheSizeBytes)
	assert.Equal(t, "5.0 GB", s.MaxCacheSize)
	assert.Equal(t, units.MetricBytes(5000000000), s.MaxCacheSizeBytes)
}
