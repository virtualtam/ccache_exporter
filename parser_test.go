// Copyright 2018 VirtualTam.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package ccache

import (
	"errors"
	"io/ioutil"
	"testing"

	"github.com/alecthomas/units"
)

func TestParseReference(t *testing.T) {
	cases := []struct {
		tname     string
		inputFile string
		want      *Statistics
		wantErr   error
	}{
		// ccache 3.4.3
		{
			tname:     "3.4.3 empty cache",
			inputFile: "testdata/3.4.3/empty",
			want: &Statistics{
				CacheDirectory:          "/home/virtualtam/.ccache",
				PrimaryConfig:           "/home/virtualtam/.ccache/ccache.conf",
				SecondaryConfigReadonly: "/etc/ccache.conf",
				CacheSize:               "0.0 kB",
				MaxCacheSize:            "15.0 GB",
				MaxCacheSizeBytes:       units.MetricBytes(15000000000),
			},
		},
		{
			tname:     "3.4.3 first build",
			inputFile: "testdata/3.4.3/firstbuild",
			want: &Statistics{
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
			tname:     "3.4.3 second build",
			inputFile: "testdata/3.4.3/secondbuild",
			want: &Statistics{
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
			tname:     "3.5 empty cache",
			inputFile: "testdata/3.5/empty",
			want: &Statistics{
				CacheDirectory:          "/home/virtualtam/.ccache",
				PrimaryConfig:           "/home/virtualtam/.ccache/ccache.conf",
				SecondaryConfigReadonly: "/etc/ccache.conf",
				CacheSize:               "0.0 kB",
				MaxCacheSize:            "5.0 GB",
				MaxCacheSizeBytes:       units.MetricBytes(5000000000),
			},
		},
		{
			tname:     "3.5 first build",
			inputFile: "testdata/3.5/firstbuild",
			want: &Statistics{
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
			tname:     "3.5 second build",
			inputFile: "testdata/3.5/secondbuild",
			want: &Statistics{
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

	for _, tc := range cases {
		t.Run(tc.tname, func(t *testing.T) {
			input, err := ioutil.ReadFile(tc.inputFile)
			if err != nil {
				t.Fatalf("failed to open test input: %q", err)
			}

			s, err := Parse(string(input))

			if tc.wantErr != nil {
				if err == nil {
					t.Fatal("expected an error, got none")
				} else if err.Error() != tc.wantErr.Error() {
					t.Fatalf("want error %q, got %q", tc.wantErr, err)
				}

				return
			}

			if err != nil {
				t.Fatalf("expected no error, got %q", err)
			}

			assertStatisticsEqual(t, s, tc.want)
		})
	}
}

func TestParse(t *testing.T) {
	cases := []struct {
		tname   string
		input   string
		want    *Statistics
		wantErr error
	}{
		// ccache cache size units
		{
			tname: "cache size in kB",
			input: `cache size                          16.7 kB
max cache size                      57.0 kB
`,
			want: &Statistics{
				CacheSize:         "16.7 kB",
				CacheSizeBytes:    units.MetricBytes(16700),
				MaxCacheSize:      "57.0 kB",
				MaxCacheSizeBytes: units.MetricBytes(57000),
			},
		},

		// error cases
		{
			tname: "unexpected date format",
			input: `stats zeroed                        not a date
`,
			want:    &Statistics{},
			wantErr: errors.New("parsing time \"not a date\" as \"Mon Jan 2 15:04:05 2006\": cannot parse \"not a date\" as \"Mon\""),
		},
		{
			tname: "unexpected cache size unit",
			input: `cache size                         655.4 zB
`,
			want:    &Statistics{},
			wantErr: errors.New("units: unknown unit ZB in 655.4ZB"),
		},
		{
			tname: "unexpected max cache size unit",
			input: `max cache size                      10.7 dB
`,
			want:    &Statistics{},
			wantErr: errors.New("units: unknown unit DB in 10.7DB"),
		},
	}

	for _, tc := range cases {
		t.Run(tc.tname, func(t *testing.T) {
			s, err := Parse(tc.input)

			if tc.wantErr != nil {
				if err == nil {
					t.Fatal("expected an error, got none")
				} else if err.Error() != tc.wantErr.Error() {
					t.Fatalf("want error %q, got %q", tc.wantErr, err)
				}

				return
			}

			if err != nil {
				t.Fatalf("expected no error, got %q", err)
			}

			assertStatisticsEqual(t, s, tc.want)
		})
	}
}

func assertStatisticsEqual(t *testing.T, got, want *Statistics) {
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
