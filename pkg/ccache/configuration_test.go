// Copyright (c) VirtualTam
// SPDX-License-Identifier: MIT

package ccache

import (
	"testing"

	"github.com/alecthomas/units"
)

func TestParseConfiguration(t *testing.T) {
	cases := []struct {
		tname string
		input string
		want  Configuration
	}{
		{
			tname: "configured via environment and file (without size unit)",
			input: `(environment) cache_dir = /home/cached/.ccache
(/home/cached/.ccache/ccache.conf) max_size = 5.0G
`,
			want: Configuration{
				CacheDirectory:    "/home/cached/.ccache",
				PrimaryConfig:     "/home/cached/.ccache/ccache.conf",
				MaxCacheSize:      "5.0GB",
				MaxCacheSizeBytes: units.MetricBytes(5000000000),
			},
		},
		{
			tname: "configured via file (without size unit)",
			input: `(default) cache_dir = /home/cached/.ccache
(/home/cached/.ccache/ccache.conf) max_size = 15.0G
`,
			want: Configuration{
				CacheDirectory:    "/home/cached/.ccache",
				PrimaryConfig:     "/home/cached/.ccache/ccache.conf",
				MaxCacheSize:      "15.0GB",
				MaxCacheSizeBytes: units.MetricBytes(15000000000),
			},
		},
		{
			tname: "configured via file",
			input: `(default) cache_dir = /home/cached/.ccache
(/home/cached/.ccache/ccache.conf) max_size = 17.0 GB
`,
			want: Configuration{
				CacheDirectory:    "/home/cached/.ccache",
				PrimaryConfig:     "/home/cached/.ccache/ccache.conf",
				MaxCacheSize:      "17.0GB",
				MaxCacheSizeBytes: units.MetricBytes(17000000000),
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.tname, func(t *testing.T) {
			got, err := ParseConfiguration(tc.input)
			if err != nil {
				t.Fatalf("want no error, got %q", err)
			}

			assertConfigurationsEqual(t, got, &tc.want)
		})
	}
}

func assertConfigurationsEqual(t *testing.T, got, want *Configuration) {
	t.Helper()

	assertStringFieldEquals(t, "CacheDirectory", got.CacheDirectory, want.CacheDirectory)
	assertStringFieldEquals(t, "PrimaryConfig", got.PrimaryConfig, want.PrimaryConfig)
	assertStringFieldEquals(t, "MaxCacheSize", got.MaxCacheSize, want.MaxCacheSize)
	assertMetricByteFieldEquals(t, "MaxCacheSizeBytes", got.MaxCacheSizeBytes, want.MaxCacheSizeBytes)
}
