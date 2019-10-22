package ccache

import (
	"io/ioutil"
	"testing"

	"github.com/alecthomas/units"
	"github.com/stretchr/testify/assert"
)

func TestParseEmptyCacheStats343(t *testing.T) {
	input, err := ioutil.ReadFile("tests/3.4.3/empty.txt")
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
	assert.Equal(t, "15.0 GB", s.MaxCacheSize)
	assert.Equal(t, units.MetricBytes(15000000000), s.MaxCacheSizeBytes)
}

func TestParseFirstBuildStats343(t *testing.T) {
	input, err := ioutil.ReadFile("tests/3.4.3/first.txt")
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
	assert.Equal(t, 116, s.CacheMiss)
	assert.Equal(t, 0.0, s.CacheHitRate)
	assert.Equal(t, 0.0, s.CacheHitRatio)
	assert.Equal(t, 14, s.CalledForLink)
	assert.Equal(t, 85, s.CalledForPreprocessing)
	assert.Equal(t, 2, s.UnsupportedCodeDirective)
	assert.Equal(t, 29, s.NoInputFile)
	assert.Equal(t, 0, s.CleanupsPerformed)
	assert.Equal(t, 361, s.FilesInCache)
	assert.Equal(t, "6.4 MB", s.CacheSize)
	assert.Equal(t, units.MetricBytes(6400000), s.CacheSizeBytes)
	assert.Equal(t, "15.0 GB", s.MaxCacheSize)
	assert.Equal(t, units.MetricBytes(15000000000), s.MaxCacheSizeBytes)
}

func TestParseSecondBuildStats343(t *testing.T) {
	input, err := ioutil.ReadFile("tests/3.4.3/second.txt")
	if err != nil {
		panic(err)
	}

	s := Statistics{}
	s.Parse(string(input))

	assert.Equal(t, "/home/virtualtam/.ccache", s.CacheDirectory)
	assert.Equal(t, "/home/virtualtam/.ccache/ccache.conf", s.PrimaryConfig)
	assert.Equal(t, "/etc/ccache.conf", s.SecondaryConfigReadonly)
	assert.Equal(t, 73, s.CacheHitDirect)
	assert.Equal(t, 4, s.CacheHitPreprocessed)
	assert.Equal(t, 207, s.CacheMiss)
	assert.Equal(t, 27.11, s.CacheHitRate)
	assert.Equal(t, 0.2711, s.CacheHitRatio)
	assert.Equal(t, 28, s.CalledForLink)
	assert.Equal(t, 170, s.CalledForPreprocessing)
	assert.Equal(t, 4, s.UnsupportedCodeDirective)
	assert.Equal(t, 58, s.NoInputFile)
	assert.Equal(t, 0, s.CleanupsPerformed)
	assert.Equal(t, 639, s.FilesInCache)
	assert.Equal(t, "12.1 MB", s.CacheSize)
	assert.Equal(t, units.MetricBytes(12100000), s.CacheSizeBytes)
	assert.Equal(t, "15.0 GB", s.MaxCacheSize)
	assert.Equal(t, units.MetricBytes(15000000000), s.MaxCacheSizeBytes)
}
