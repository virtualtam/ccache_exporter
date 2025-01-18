// Copyright (c) VirtualTam
// SPDX-License-Identifier: MIT

package ccache

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/alecthomas/units"
)

type pre37TestCase struct {
	tname         string
	inputFilename string
	wantStats     Statistics
	wantErr       error
}

type pre37TestSession struct {
	osAndVersion     string
	osAndVersionCode string
	ccacheVersion    string
	testCases        []pre37TestCase
}

func TestParsePre37Statistics(t *testing.T) {
	sessions := []pre37TestSession{
		{
			osAndVersion:     "Arch Linux",
			osAndVersionCode: "arch-rolling",
			ccacheVersion:    "3.4.3",

			testCases: []pre37TestCase{
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
			osAndVersionCode: "arch-rolling",
			ccacheVersion:    "3.5",

			testCases: []pre37TestCase{
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

			testCases: []pre37TestCase{
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

			testCases: []pre37TestCase{
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

			testCases: []pre37TestCase{
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
					fmt.Sprintf("%s-ccache-%s", ts.osAndVersionCode, ts.ccacheVersion),
					tc.inputFilename,
				)
				input, err := os.ReadFile(inputFilepath)
				if err != nil {
					t.Fatalf("failed to open test input: %q", err)
				}

				s, err := ParsePre37Statistics(string(input))

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

				assertStatisticsEqual(t, s, &tc.wantStats)
			})
		}
	}
}

func TestParsePre37StatisticsEdgeCases(t *testing.T) {
	cases := []struct {
		tname     string
		input     string
		wantStats Statistics
		wantErr   error
	}{
		// ccache cache size units
		{
			tname: "cache size in kB",
			input: `cache size                          16.7 kB
max cache size                      57.0 kB
`,
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
			s, err := ParsePre37Statistics(tc.input)

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

			assertStatisticsEqual(t, s, &tc.wantStats)
		})
	}
}
