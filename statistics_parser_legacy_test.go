// Copyright 2018-2022 VirtualTam.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package ccache

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/alecthomas/units"
)

type legacyTestCase struct {
	tname         string
	inputFilename string
	wantStats     Statistics
	wantErr       error
}

type legacyTestSession struct {
	osAndVersion     string
	osAndVersionCode string
	ccacheVersion    string
	wantConfig       Configuration
	testCases        []legacyTestCase
}

func TestLegacyParserParseShowStats(t *testing.T) {
	sessions := []legacyTestSession{
		{
			osAndVersion:     "Arch Linux",
			osAndVersionCode: "arch",
			ccacheVersion:    "3.4.3",
			wantConfig: Configuration{
				CacheDirectory:          "/home/virtualtam/.ccache",
				PrimaryConfig:           "/home/virtualtam/.ccache/ccache.conf",
				SecondaryConfigReadonly: "/etc/ccache.conf",
				MaxCacheSize:            "15.0 GB",
				MaxCacheSizeBytes:       units.MetricBytes(15000000000),
			},

			testCases: []legacyTestCase{
				{
					tname:         "empty cache",
					inputFilename: "empty",
					wantStats: Statistics{
						CacheSize: "0.0 kB",
					},
				},
				{
					tname:         "first build",
					inputFilename: "firstbuild",
					wantStats: Statistics{
						CacheMiss:                116,
						CalledForLink:            14,
						CalledForPreprocessing:   85,
						UnsupportedCodeDirective: 2,
						NoInputFile:              29,
						FilesInCache:             361,
						CacheSize:                "6.4 MB",
						CacheSizeBytes:           units.MetricBytes(6400000),
					},
				},
				{
					tname:         "second build",
					inputFilename: "secondbuild",
					wantStats: Statistics{
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
					},
				},
			},
		},

		{
			osAndVersion:     "Arch Linux",
			osAndVersionCode: "arch",
			ccacheVersion:    "3.5",
			wantConfig: Configuration{
				CacheDirectory:          "/home/virtualtam/.ccache",
				PrimaryConfig:           "/home/virtualtam/.ccache/ccache.conf",
				SecondaryConfigReadonly: "/etc/ccache.conf",
				MaxCacheSize:            "5.0 GB",
				MaxCacheSizeBytes:       units.MetricBytes(5000000000),
			},

			testCases: []legacyTestCase{
				{
					tname:         "empty cache",
					inputFilename: "empty",
					wantStats: Statistics{
						CacheSize: "0.0 kB",
					},
				},
				{
					tname:         "first build",
					inputFilename: "firstbuild",
					wantStats: Statistics{
						CacheHitDirect:         1,
						CacheHitPreprocessed:   15,
						CacheMiss:              342,
						CacheHitRate:           4.47,
						CacheHitRatio:          0.0447,
						CalledForLink:          14,
						CalledForPreprocessing: 1,
						FilesInCache:           867,
						CacheSize:              "44.5 MB",
						CacheSizeBytes:         units.MetricBytes(44500000),
					},
				},
				{
					tname:         "second build",
					inputFilename: "secondbuild",
					wantStats: Statistics{
						CacheHitDirect:         349,
						CacheHitPreprocessed:   10,
						CacheMiss:              28,
						CacheHitRate:           92.76,
						CacheHitRatio:          0.9276000000000001,
						CalledForLink:          14,
						CalledForPreprocessing: 1,
						FilesInCache:           943,
						CacheSize:              "46.7 MB",
						CacheSizeBytes:         units.MetricBytes(46700000),
					},
				},
			},
		},

		{
			osAndVersion:     "Debian 9",
			osAndVersionCode: "debian-9",
			ccacheVersion:    "3.3.4",
			wantConfig: Configuration{
				CacheDirectory:          "/home/cached/.ccache",
				PrimaryConfig:           "/home/cached/.ccache/ccache.conf",
				SecondaryConfigReadonly: "/etc/ccache.conf",
				MaxCacheSize:            "5.0 GB",
				MaxCacheSizeBytes:       units.MetricBytes(5000000000),
			},

			testCases: []legacyTestCase{
				{
					tname:         "empty cache",
					inputFilename: "empty",
					wantStats: Statistics{
						CacheSize: "0.0 kB",
					},
				},
				{
					tname:         "first build",
					inputFilename: "firstbuild",
					wantStats: Statistics{
						CacheMiss:                52,
						CalledForLink:            1,
						CalledForPreprocessing:   11,
						NoInputFile:              7,
						UnsupportedCodeDirective: 1,
						FilesInCache:             103,
						CacheSize:                "1.2 MB",
						CacheSizeBytes:           units.MetricBytes(1200000),
					},
				},
				{
					tname:         "second build",
					inputFilename: "secondbuild",
					wantStats: Statistics{
						CacheHitDirect:           50,
						CacheHitPreprocessed:     2,
						CacheMiss:                52,
						CacheHitRate:             50.0,
						CacheHitRatio:            0.5,
						CalledForLink:            2,
						CalledForPreprocessing:   22,
						NoInputFile:              14,
						UnsupportedCodeDirective: 2,
						FilesInCache:             103,
						CacheSize:                "1.2 MB",
						CacheSizeBytes:           units.MetricBytes(1200000),
					},
				},
			},
		},
		{
			osAndVersion:     "Debian 10",
			osAndVersionCode: "debian-10",
			ccacheVersion:    "3.6",
			wantConfig: Configuration{
				CacheDirectory:          "/home/cached/.ccache",
				PrimaryConfig:           "/home/cached/.ccache/ccache.conf",
				SecondaryConfigReadonly: "/etc/ccache.conf",
				MaxCacheSize:            "5.0 GB",
				MaxCacheSizeBytes:       units.MetricBytes(5000000000),
			},

			testCases: []legacyTestCase{
				{
					tname:         "empty cache",
					inputFilename: "empty",
					wantStats: Statistics{
						CacheSize: "0.0 kB",
					},
				},
				{
					tname:         "first build",
					inputFilename: "firstbuild",
					wantStats: Statistics{
						CacheMiss:              57,
						CalledForLink:          1,
						CalledForPreprocessing: 11,
						NoInputFile:            4,
						FilesInCache:           110,
						CacheSize:              "1.7 MB",
						CacheSizeBytes:         units.MetricBytes(1700000),
					},
				},
				{
					tname:         "second build",
					inputFilename: "secondbuild",
					wantStats: Statistics{
						CacheHitDirect:         52,
						CacheHitPreprocessed:   5,
						CacheMiss:              57,
						CacheHitRate:           50.0,
						CacheHitRatio:          0.5,
						CalledForLink:          2,
						CalledForPreprocessing: 22,
						NoInputFile:            8,
						FilesInCache:           110,
						CacheSize:              "1.7 MB",
						CacheSizeBytes:         units.MetricBytes(1700000),
					},
				},
			},
		},

		{
			osAndVersion:     "Ubuntu 18.04",
			osAndVersionCode: "ubuntu-18.04",
			ccacheVersion:    "3.4.1",
			wantConfig: Configuration{
				CacheDirectory:          "/home/cached/.ccache",
				PrimaryConfig:           "/home/cached/.ccache/ccache.conf",
				SecondaryConfigReadonly: "/etc/ccache.conf",
				MaxCacheSize:            "5.0 GB",
				MaxCacheSizeBytes:       units.MetricBytes(5000000000),
			},

			testCases: []legacyTestCase{
				{
					tname:         "empty cache",
					inputFilename: "empty",
					wantStats: Statistics{
						CacheSize: "0.0 kB",
					},
				},
				{
					tname:         "first build",
					inputFilename: "firstbuild",
					wantStats: Statistics{
						CacheMiss:              53,
						CalledForLink:          1,
						CalledForPreprocessing: 11,
						NoInputFile:            4,
						FilesInCache:           104,
						CacheSize:              "1.6 MB",
						CacheSizeBytes:         units.MetricBytes(1600000),
					},
				},
				{
					tname:         "second build",
					inputFilename: "secondbuild",
					wantStats: Statistics{
						CacheHitDirect:         50,
						CacheHitPreprocessed:   3,
						CacheMiss:              53,
						CacheHitRate:           50.0,
						CacheHitRatio:          0.5,
						CalledForLink:          2,
						CalledForPreprocessing: 22,
						NoInputFile:            8,
						FilesInCache:           104,
						CacheSize:              "1.6 MB",
						CacheSizeBytes:         units.MetricBytes(1600000),
					},
				},
			},
		},
	}

	for _, ts := range sessions {
		for _, tc := range ts.testCases {
			t.Run(fmt.Sprintf("ccache %s on %s (%s)", ts.ccacheVersion, ts.osAndVersion, tc.tname), func(t *testing.T) {
				inputFilepath := filepath.Join(
					"testdata",
					fmt.Sprintf("%s-%s", ts.osAndVersionCode, ts.ccacheVersion),
					tc.inputFilename,
				)
				input, err := ioutil.ReadFile(inputFilepath)
				if err != nil {
					t.Fatalf("failed to open test input: %q", err)
				}

				parser := NewLegacyParser()
				c, s, err := parser.ParseShowStats(string(input))

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

				assertConfigurationsEqual(t, c, &ts.wantConfig)
				assertStatisticsEqual(t, s, &tc.wantStats)
			})
		}
	}
}

func TestLegacyParserParseShowStatsEdgeCases(t *testing.T) {
	cases := []struct {
		tname      string
		input      string
		wantConfig Configuration
		wantStats  Statistics
		wantErr    error
	}{
		// ccache cache size units
		{
			tname: "cache size in kB",
			input: `cache size                          16.7 kB
max cache size                      57.0 kB
`,
			wantConfig: Configuration{

				MaxCacheSize:      "57.0 kB",
				MaxCacheSizeBytes: units.MetricBytes(57000),
			},
			wantStats: Statistics{
				CacheSize:      "16.7 kB",
				CacheSizeBytes: units.MetricBytes(16700),
			},
		},

		// error cases
		{
			tname: "unexpected date format",
			input: `stats zeroed                        not a date
`,
			wantErr: errors.New("parsing time \"not a date\" as \"Mon Jan 2 15:04:05 2006\": cannot parse \"not a date\" as \"Mon\""),
		},
		{
			tname: "unexpected cache size unit",
			input: `cache size                         655.4 zB
`,
			wantErr: errors.New("units: unknown unit ZB in 655.4ZB"),
		},
		{
			tname: "unexpected max cache size unit",
			input: `max cache size                      10.7 dB
`,
			wantErr: errors.New("units: unknown unit DB in 10.7DB"),
		},
	}

	for _, tc := range cases {
		t.Run(tc.tname, func(t *testing.T) {
			parser := NewLegacyParser()
			c, s, err := parser.ParseShowStats(tc.input)

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

			assertConfigurationsEqual(t, c, &tc.wantConfig)
			assertStatisticsEqual(t, s, &tc.wantStats)
		})
	}
}
