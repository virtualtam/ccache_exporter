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
	}

	for _, tt := range cases {
		t.Run(tt.tname, func(t *testing.T) {
			s := Statistics{}
			s.Parse(tt.input)

			assertStatisticsEqual(t, tt.want, s)
		})
	}
}

func assertStatisticsEqual(t *testing.T, want Statistics, got Statistics) {
	t.Helper()

	if got.CacheDirectory != got.CacheDirectory {
		t.Errorf("CacheDirectory: want  %q, got %q", want.CacheDirectory, got.CacheDirectory)
	}
	if got.PrimaryConfig != got.PrimaryConfig {
		t.Errorf("PrimaryConfig: want %q, got %q", want.PrimaryConfig, got.PrimaryConfig)
	}
	if got.SecondaryConfigReadonly != got.SecondaryConfigReadonly {
		t.Errorf("SecondaryConfigReadonly: want %q, got %q", want.SecondaryConfigReadonly, got.SecondaryConfigReadonly)
	}
	if got.CacheHitDirect != want.CacheHitDirect {
		t.Errorf("CacheHitDirect: want %d, got %d", want.CacheHitDirect, got.CacheHitDirect)
	}
	if got.CacheHitPreprocessed != want.CacheHitPreprocessed {
		t.Errorf("CacheHitPreprocessed: want %d, got %d", want.CacheHitPreprocessed, got.CacheHitPreprocessed)
	}
	if got.CacheMiss != want.CacheMiss {
		t.Errorf("CacheMiss: want %d, got %d", want.CacheMiss, got.CacheMiss)
	}
	if got.CacheHitRate != want.CacheHitRate {
		t.Errorf("CacheHitRate: want %f, got %f", want.CacheHitRate, got.CacheHitRate)
	}
	if got.CacheHitRatio != want.CacheHitRatio {
		t.Errorf("CacheHitRatio: want %f, got %f", want.CacheHitRatio, got.CacheHitRatio)
	}
	if got.CalledForLink != want.CalledForLink {
		t.Errorf("CalledForLink: want %d, got %d", want.CalledForLink, got.CalledForLink)
	}
	if got.CalledForPreprocessing != want.CalledForPreprocessing {
		t.Errorf("CalledForPreprocessing: want %d, got %d", want.CalledForPreprocessing, got.CalledForPreprocessing)
	}
	if got.UnsupportedCodeDirective != want.UnsupportedCodeDirective {
		t.Errorf("UnsupportedCodeDirective: want %d, got %d", want.UnsupportedCodeDirective, got.UnsupportedCodeDirective)
	}
	if got.NoInputFile != want.NoInputFile {
		t.Errorf("NoInputFile: want %d, got %d", want.NoInputFile, got.NoInputFile)
	}
	if got.CleanupsPerformed != want.CleanupsPerformed {
		t.Errorf("CleanupsPerformed: want %d, got %d", want.CleanupsPerformed, got.CleanupsPerformed)
	}
	if got.FilesInCache != want.FilesInCache {
		t.Errorf("FilesInCache: want %d, got %d", want.FilesInCache, got.FilesInCache)
	}
	if got.CacheSize != want.CacheSize {
		t.Errorf("CacheSize: want %q, got %q", want.CacheSize, got.CacheSize)
	}
	if got.CacheSizeBytes != want.CacheSizeBytes {
		t.Errorf("CacheSizeBytes: want %d, got %d", want.CacheSizeBytes, got.CacheSizeBytes)
	}
	if got.MaxCacheSize != want.MaxCacheSize {
		t.Errorf("MaxCacheSize: want %q, got %q", want.MaxCacheSize, got.MaxCacheSize)
	}
	if got.MaxCacheSizeBytes != want.MaxCacheSizeBytes {
		t.Errorf("MaxCacheSizeBytes: want %d, got %d", want.MaxCacheSizeBytes, got.MaxCacheSizeBytes)
	}
}
