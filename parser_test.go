// Copyright 2018 VirtualTam.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package ccache

import (
	"testing"

	"github.com/alecthomas/units"
)

func TestParse(t *testing.T) {
	cases := []struct {
		tname string
		input string
		want  Statistics
	}{
		// ccache 3.4.3
		{
			tname: "3.4.3 empty cache",

			input: `cache directory                     /home/virtualtam/.ccache
primary config                      /home/virtualtam/.ccache/ccache.conf
secondary config      (readonly)    /etc/ccache.conf
stats zero time                     Sun Sep 23 01:18:52 2018
cache hit (direct)                     0
cache hit (preprocessed)               0
cache miss                             0
cache hit rate                      0.00 %
cleanups performed                     0
files in cache                         0
cache size                           0.0 kB
max cache size                      15.0 GB
`,

			want: Statistics{
				CacheDirectory:          "/home/virtualtam/.ccache",
				PrimaryConfig:           "/home/virtualtam/.ccache/ccache.conf",
				SecondaryConfigReadonly: "/etc/ccache.conf",
				CacheSize:               "0.0 kB",
				MaxCacheSize:            "15.0 GB",
				MaxCacheSizeBytes:       units.MetricBytes(15000000000),
			},
		},
		{
			tname: "3.4.3 first build",

			input: `cache directory                     /home/virtualtam/.ccache
primary config                      /home/virtualtam/.ccache/ccache.conf
secondary config      (readonly)    /etc/ccache.conf
stats zero time                     Sun Sep 23 01:18:52 2018
cache hit (direct)                     0
cache hit (preprocessed)               0
cache miss                           116
cache hit rate                      0.00 %
called for link                       14
called for preprocessing              85
unsupported code directive             2
no input file                         29
cleanups performed                     0
files in cache                       361
cache size                           6.4 MB
max cache size                      15.0 GB
`,

			want: Statistics{
				CacheDirectory:           "/home/virtualtam/.ccache",
				PrimaryConfig:            "/home/virtualtam/.ccache/ccache.conf",
				SecondaryConfigReadonly:  "/etc/ccache.conf",
				CacheMiss:                116,
				CalledForLink:            14,
				CalledForPreprocessing:   85,
				UnsupportedCodeDirective: 2,
				NoInputFile:              29,
				FilesInCache:             361,
				CacheSize:                "6.4 MB",
				CacheSizeBytes:           units.MetricBytes(6400000),
				MaxCacheSize:             "15.0 GB",
				MaxCacheSizeBytes:        units.MetricBytes(15000000000),
			},
		},
		{
			tname: "3.4.3 second build",

			input: `cache directory                     /home/virtualtam/.ccache
primary config                      /home/virtualtam/.ccache/ccache.conf
secondary config      (readonly)    /etc/ccache.conf
stats zero time                     Sun Sep 23 01:18:52 2018
cache hit (direct)                    73
cache hit (preprocessed)               4
cache miss                           207
cache hit rate                     27.11 %
called for link                       28
called for preprocessing             170
unsupported code directive             4
no input file                         58
cleanups performed                     0
files in cache                       639
cache size                          12.1 MB
max cache size                      15.0 GB
`,

			want: Statistics{
				CacheDirectory:           "/home/virtualtam/.ccache",
				PrimaryConfig:            "/home/virtualtam/.ccache/ccache.conf",
				SecondaryConfigReadonly:  "/etc/ccache.conf",
				CacheHitDirect:           73,
				CacheHitPreprocessed:     4,
				CacheMiss:                207,
				CacheHitRate:             27.11,
				CacheHitRatio:            0.2711,
				CalledForLink:            28,
				CalledForPreprocessing:   170,
				UnsupportedCodeDirective: 4,
				NoInputFile:              58,
				FilesInCache:             639,
				CacheSize:                "12.1 MB",
				CacheSizeBytes:           units.MetricBytes(12100000),
				MaxCacheSize:             "15.0 GB",
				MaxCacheSizeBytes:        units.MetricBytes(15000000000),
			},
		},

		// ccache 3.5
		{
			tname: "3.5 empty cache",

			input: `cache directory                     /home/virtualtam/.ccache
primary config                      /home/virtualtam/.ccache/ccache.conf
secondary config      (readonly)    /etc/ccache.conf
cache hit (direct)                     0
cache hit (preprocessed)               0
cache miss                             0
cache hit rate                      0.00 %
cleanups performed                     0
files in cache                         0
cache size                           0.0 kB
max cache size                       5.0 GB
`,

			want: Statistics{
				CacheDirectory:          "/home/virtualtam/.ccache",
				PrimaryConfig:           "/home/virtualtam/.ccache/ccache.conf",
				SecondaryConfigReadonly: "/etc/ccache.conf",
				CacheSize:               "0.0 kB",
				MaxCacheSize:            "5.0 GB",
				MaxCacheSizeBytes:       units.MetricBytes(5000000000),
			},
		},
		{
			tname: "3.5 first build",

			input: `cache directory                     /home/virtualtam/.ccache
primary config                      /home/virtualtam/.ccache/ccache.conf
secondary config      (readonly)    /etc/ccache.conf
stats updated                       Sat Oct 20 00:49:12 2018
cache hit (direct)                     1
cache hit (preprocessed)              15
cache miss                           342
cache hit rate                      4.47 %
called for link                       14
called for preprocessing               1
preprocessor error                     1
cleanups performed                     0
files in cache                       867
cache size                          44.5 MB
max cache size                       5.0 GB
`,

			want: Statistics{
				CacheDirectory:          "/home/virtualtam/.ccache",
				PrimaryConfig:           "/home/virtualtam/.ccache/ccache.conf",
				SecondaryConfigReadonly: "/etc/ccache.conf",
				CacheHitDirect:          1,
				CacheHitPreprocessed:    15,
				CacheMiss:               342,
				CacheHitRate:            4.47,
				CacheHitRatio:           0.0447,
				CalledForLink:           14,
				CalledForPreprocessing:  1,
				FilesInCache:            867,
				CacheSize:               "44.5 MB",
				CacheSizeBytes:          units.MetricBytes(44500000),
				MaxCacheSize:            "5.0 GB",
				MaxCacheSizeBytes:       units.MetricBytes(5000000000),
			},
		},
		{
			tname: "3.5 second build",

			input: `cache directory                     /home/virtualtam/.ccache
primary config                      /home/virtualtam/.ccache/ccache.conf
secondary config      (readonly)    /etc/ccache.conf
stats updated                       Sat Oct 20 00:50:26 2018
stats zeroed                        Sat Oct 20 00:49:42 2018
cache hit (direct)                   349
cache hit (preprocessed)              10
cache miss                            28
cache hit rate                     92.76 %
called for link                       14
called for preprocessing               1
preprocessor error                     1
cleanups performed                     0
files in cache                       943
cache size                          46.7 MB
max cache size                       5.0 GB
`,

			want: Statistics{
				CacheDirectory:          "/home/virtualtam/.ccache",
				PrimaryConfig:           "/home/virtualtam/.ccache/ccache.conf",
				SecondaryConfigReadonly: "/etc/ccache.conf",
				CacheHitDirect:          349,
				CacheHitPreprocessed:    10,
				CacheMiss:               28,
				CacheHitRate:            92.76,
				CacheHitRatio:           0.9276000000000001,
				CalledForLink:           14,
				CalledForPreprocessing:  1,
				FilesInCache:            943,
				CacheSize:               "46.7 MB",
				CacheSizeBytes:          units.MetricBytes(46700000),
				MaxCacheSize:            "5.0 GB",
				MaxCacheSizeBytes:       units.MetricBytes(5000000000),
			},
		},

		// ccache cache size units
		{
			tname: "cache size in kB",
			input: `cache size                          16.7 kB
max cache size                      57.0 kB
`,
			want: Statistics{
				CacheSize:         "16.7 kB",
				CacheSizeBytes:    units.MetricBytes(16700),
				MaxCacheSize:      "57.0 kB",
				MaxCacheSizeBytes: units.MetricBytes(57000),
			},
		},
	}

	for _, tt := range cases {
		t.Run(tt.tname, func(t *testing.T) {
			s := Statistics{}
			s.Parse(tt.input)

			assertStatisticsEqual(t, s, tt.want)
		})
	}
}

func assertStatisticsEqual(t *testing.T, got, want Statistics) {
	t.Helper()

	assertStringFieldEquals(t, "CacheDirectory", got.CacheDirectory, want.CacheDirectory)
	assertStringFieldEquals(t, "PrimaryConfig", got.PrimaryConfig, want.PrimaryConfig)
	assertStringFieldEquals(t, "SecondaryConfigReadonly", got.SecondaryConfigReadonly, want.SecondaryConfigReadonly)
	assertIntFieldEquals(t, "CacheHitDirect", got.CacheHitDirect, want.CacheHitDirect)
	assertIntFieldEquals(t, "CacheHitPreprocessed", got.CacheHitPreprocessed, want.CacheHitPreprocessed)
	assertIntFieldEquals(t, "CacheMiss", got.CacheMiss, want.CacheMiss)
	assertFloatFieldEquals(t, "CacheHitRate", got.CacheHitRate, want.CacheHitRate)
	assertFloatFieldEquals(t, "CacheHitRatio", got.CacheHitRatio, want.CacheHitRatio)
	assertIntFieldEquals(t, "CalledForLink", got.CalledForLink, want.CalledForLink)
	assertIntFieldEquals(t, "CalledForPreprocessing", got.CalledForPreprocessing, want.CalledForPreprocessing)
	assertIntFieldEquals(t, "UnsupportedCodeDirective", got.UnsupportedCodeDirective, want.UnsupportedCodeDirective)
	assertIntFieldEquals(t, "NoInputFile", got.NoInputFile, want.NoInputFile)
	assertIntFieldEquals(t, "CleanupsPerformed", got.CleanupsPerformed, want.CleanupsPerformed)
	assertIntFieldEquals(t, "FilesInCache", got.FilesInCache, want.FilesInCache)
	assertStringFieldEquals(t, "CacheSize", got.CacheSize, want.CacheSize)
	assertMetricByteFieldEquals(t, "CacheSizeBytes", got.CacheSizeBytes, want.CacheSizeBytes)
	assertStringFieldEquals(t, "MaxCacheSize", got.MaxCacheSize, want.MaxCacheSize)
	assertMetricByteFieldEquals(t, "MaxCacheSizeBytes", got.MaxCacheSizeBytes, want.MaxCacheSizeBytes)
}

func assertFloatFieldEquals(t *testing.T, fieldName string, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("%s: want %f, got %f", fieldName, want, got)
	}
}

func assertIntFieldEquals(t *testing.T, fieldName string, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("%s: want %d, got %d", fieldName, want, got)
	}
}

func assertMetricByteFieldEquals(t *testing.T, fieldName string, got, want units.MetricBytes) {
	t.Helper()
	if got != want {
		t.Errorf("%s: want %d, got %d", fieldName, want, got)
	}
}

func assertStringFieldEquals(t *testing.T, fieldName, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("%s: want %q, got %q", fieldName, want, got)
	}
}
