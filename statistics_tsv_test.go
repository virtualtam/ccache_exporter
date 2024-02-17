// Copyright 2018-2022 VirtualTam.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package ccache

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/alecthomas/units"
)

type tsvTestCase struct {
	tname         string
	inputFilename string
	wantStats     Statistics
	wantErr       error
}

type tsvTestSession struct {
	osAndVersion     string
	osAndVersionCode string
	ccacheVersion    string
	redisVersion     string
	testCases        []tsvTestCase
}

func TestParseTSVStatistics(t *testing.T) {
	sessions := []tsvTestSession{
		{
			osAndVersion:     "Arch Linux",
			osAndVersionCode: "arch-rolling",
			ccacheVersion:    "4.6.1",

			testCases: []tsvTestCase{
				{
					tname:         "empty cache",
					inputFilename: "empty.tsv",
					wantStats: Statistics{
						CacheSize: "0B",
					},
				},
				{
					tname:         "first build",
					inputFilename: "firstbuild.tsv",
					wantStats: Statistics{
						CacheMiss:             150,
						CacheMissDirect:       156,
						CacheMissPreprocessed: 151,
						CalledForLink:         44,
						CompilationFailed:     1,
						NoInputFile:           5,
						PreprocessingFailed:   5,
						FilesInCache:          298,
						CacheSize:             "4MB",
						CacheSizeBytes:        units.MetricBytes(4263936),
					},
				},
				{
					tname:         "second build",
					inputFilename: "secondbuild.tsv",
					wantStats: Statistics{
						CacheHitDirect:        148,
						CacheHitPreprocessed:  2,
						CacheMiss:             150,
						CacheMissDirect:       164,
						CacheMissPreprocessed: 152,
						CacheHitRate:          24.350649,
						CacheHitRatio:         0.2435049,
						CalledForLink:         88,
						CompilationFailed:     2,
						NoInputFile:           9,
						PreprocessingFailed:   10,
						FilesInCache:          298,
						CacheSize:             "4MB",
						CacheSizeBytes:        units.MetricBytes(4263936),
					},
				},
			},
		},
		{
			osAndVersion:     "Debian 11",
			osAndVersionCode: "debian-11",
			ccacheVersion:    "4.2",

			testCases: []tsvTestCase{
				{
					tname:         "empty cache",
					inputFilename: "empty.tsv",
					wantStats: Statistics{
						CacheSize: "0B",
					},
				},
				{
					tname:         "first build",
					inputFilename: "firstbuild.tsv",
					wantStats: Statistics{
						CacheMiss:              197,
						CalledForLink:          45,
						CalledForPreprocessing: 74,
						CompilationFailed:      1,
						NoInputFile:            5,
						PreprocessingFailed:    4,
						FilesInCache:           390,
						CacheSize:              "38MB",
						CacheSizeBytes:         units.MetricBytes(38199296),
					},
				},
				{
					tname:         "second build",
					inputFilename: "secondbuild.tsv",
					wantStats: Statistics{
						CacheHitDirect:         121,
						CacheHitPreprocessed:   2,
						CacheMiss:              197,
						CacheHitRate:           38.437500,
						CacheHitRatio:          0.384375,
						CalledForLink:          90,
						CalledForPreprocessing: 76,
						CompilationFailed:      2,
						NoInputFile:            10,
						PreprocessingFailed:    8,
						FilesInCache:           390,
						CacheSize:              "38MB",
						CacheSizeBytes:         units.MetricBytes(38199296),
					},
				},
			},
		},
		{
			osAndVersion:     "Debian 12",
			osAndVersionCode: "debian-12",
			ccacheVersion:    "4.7.5",

			testCases: []tsvTestCase{
				{
					tname:         "empty cache",
					inputFilename: "empty.tsv",
					wantStats: Statistics{
						CacheSize: "0B",
					},
				},
				{
					tname:         "first build",
					inputFilename: "firstbuild.tsv",
					wantStats: Statistics{
						CacheMiss:              146,
						CacheMissDirect:        154,
						CacheMissPreprocessed:  150,
						CalledForLink:          41,
						CalledForPreprocessing: 0,
						CompilationFailed:      4,
						NoInputFile:            5,
						PreprocessingFailed:    4,
						FilesInCache:           290,
						CacheSize:              "24MB",
						CacheSizeBytes:         units.MetricBytes(24883200),
					},
				},
				{
					tname:         "second build",
					inputFilename: "secondbuild.tsv",
					wantStats: Statistics{
						CacheHitDirect:         109,
						CacheHitPreprocessed:   2,
						CacheMiss:              181,
						CacheMissDirect:        199,
						CacheMissPreprocessed:  189,
						CacheHitRate:           16.323529,
						CacheHitRatio:          0.163235,
						CalledForLink:          82,
						CalledForPreprocessing: 0,
						CompilationFailed:      8,
						NoInputFile:            10,
						PreprocessingFailed:    8,
						FilesInCache:           360,
						CacheSize:              "25MB",
						CacheSizeBytes:         units.MetricBytes(25260032),
					},
				},
			},
		},
		{
			osAndVersion:     "Ubuntu 20.04",
			osAndVersionCode: "ubuntu-20.04",
			ccacheVersion:    "3.7.7",

			testCases: []tsvTestCase{
				{
					tname:         "empty cache",
					inputFilename: "empty.tsv",
					wantStats: Statistics{
						CacheSize: "0B",
					},
				},
				{
					tname:         "first build",
					inputFilename: "firstbuild.tsv",
					wantStats: Statistics{
						CacheMiss:              69,
						CalledForLink:          2,
						CalledForPreprocessing: 11,
						CompilationFailed:      5,
						NoInputFile:            4,
						PreprocessingFailed:    1,
						FilesInCache:           121,
						CacheSize:              "2MB",
						CacheSizeBytes:         units.MetricBytes(2371584),
					},
				},
				{
					tname:         "second build",
					inputFilename: "secondbuild.tsv",
					wantStats: Statistics{
						CacheHitDirect:         51,
						CacheHitPreprocessed:   18,
						CacheMiss:              69,
						CacheHitRate:           50,
						CacheHitRatio:          0.5,
						CalledForLink:          4,
						CalledForPreprocessing: 22,
						CompilationFailed:      10,
						NoInputFile:            8,
						PreprocessingFailed:    2,
						FilesInCache:           121,
						CacheSize:              "2MB",
						CacheSizeBytes:         units.MetricBytes(2371584),
					},
				},
			},
		},
		{
			osAndVersion:     "Ubuntu 22.04",
			osAndVersionCode: "ubuntu-22.04",
			ccacheVersion:    "4.5.1",

			testCases: []tsvTestCase{
				{
					tname:         "empty cache",
					inputFilename: "empty.tsv",
					wantStats: Statistics{
						CacheSize: "0B",
					},
				},
				{
					tname:         "first build",
					inputFilename: "firstbuild.tsv",
					wantStats: Statistics{
						CacheMiss:              255,
						CacheMissDirect:        260,
						CacheMissPreprocessed:  256,
						CalledForLink:          45,
						CalledForPreprocessing: 103,
						CompilationFailed:      1,
						NoInputFile:            6,
						PreprocessingFailed:    4,
						FilesInCache:           506,
						CacheSize:              "134MB",
						CacheSizeBytes:         units.MetricBytes(134987776),
					},
				},
				{
					tname:         "second build",
					inputFilename: "secondbuild.tsv",
					wantStats: Statistics{
						CacheHitDirect:         150,
						CacheHitPreprocessed:   2,
						CacheMiss:              255,
						CacheMissDirect:        267,
						CacheMissPreprocessed:  257,
						CacheHitRate:           16.326531,
						CacheHitRatio:          0.163265,
						CalledForLink:          90,
						CalledForPreprocessing: 105,
						CompilationFailed:      2,
						NoInputFile:            12,
						PreprocessingFailed:    8,
						FilesInCache:           506,
						CacheSize:              "134MB",
						CacheSizeBytes:         units.MetricBytes(134987776),
					},
				},
			},
		},
		{
			osAndVersion:     "Ubuntu 24.04",
			osAndVersionCode: "ubuntu-24.04",
			ccacheVersion:    "4.9.1",

			testCases: []tsvTestCase{
				{
					tname:         "empty cache",
					inputFilename: "empty.tsv",
					wantStats: Statistics{
						CacheSize: "0B",
					},
				},
				{
					tname:         "first build",
					inputFilename: "firstbuild.tsv",
					wantStats: Statistics{
						CacheMiss:             147,
						CacheMissDirect:       157,
						CacheMissPreprocessed: 154,
						CalledForLink:         35,
						CompilationFailed:     7,
						NoInputFile:           15,
						PreprocessingFailed:   3,
						FilesInCache:          292,
						CacheSize:             "79MB",
						CacheSizeBytes:        units.MetricBytes(79470592),
					},
				},
				{
					tname:         "second build",
					inputFilename: "secondbuild.tsv",
					wantStats: Statistics{
						CacheHitDirect:        116,
						CacheHitPreprocessed:  2,
						CacheMiss:             176,
						CacheMissDirect:       198,
						CacheMissPreprocessed: 190,
						CacheHitRate:          17.302053,
						CacheHitRatio:         0.173020,
						CalledForLink:         70,
						CompilationFailed:     14,
						NoInputFile:           30,
						PreprocessingFailed:   6,
						FilesInCache:          350,
						CacheSize:             "79MB",
						CacheSizeBytes:        units.MetricBytes(79892480),
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

				s, err := ParseTSVStatistics(string(input))

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

func TestParseTSVStatisticsWithRemoteStorage(t *testing.T) {
	sessions := []tsvTestSession{
		{
			osAndVersion:     "Ubuntu 24.04",
			osAndVersionCode: "ubuntu-24.04",
			ccacheVersion:    "4.9.1",
			redisVersion:     "7",

			testCases: []tsvTestCase{
				{
					tname:         "empty cache",
					inputFilename: "empty.tsv",
					wantStats: Statistics{
						CacheSize: "0B",
					},
				},
				{
					tname:         "first build",
					inputFilename: "firstbuild.tsv",
					wantStats: Statistics{
						CacheMiss:             147,
						CacheMissDirect:       157,
						CacheMissPreprocessed: 154,
						CacheSize:             "0B",
						CalledForLink:         35,
						CompilationFailed:     7,
						NoInputFile:           15,
						PreprocessingFailed:   3,
						RemoteStorageMiss:     147,
						RemoteStorageReadMiss: 309,
						RemoteStorageWrite:    292,
					},
				},
				{
					tname:         "second build",
					inputFilename: "secondbuild.tsv",
					wantStats: Statistics{
						CacheHitDirect:        116,
						CacheHitPreprocessed:  2,
						CacheMiss:             176,
						CacheMissDirect:       198,
						CacheMissPreprocessed: 190,
						CacheHitRate:          17.302053,
						CacheHitRatio:         0.173020,
						CalledForLink:         70,
						CompilationFailed:     14,
						NoInputFile:           30,
						PreprocessingFailed:   6,
						FilesInCache:          350,
						CacheSize:             "79MB",
						CacheSizeBytes:        units.MetricBytes(79892480),
					},
				},
			},
		},
	}

	for _, ts := range sessions {
		for _, tc := range ts.testCases {
			t.Run(fmt.Sprintf("ccache %s on %s using Redis %s (%s)", ts.ccacheVersion, ts.osAndVersion, ts.redisVersion, tc.tname), func(t *testing.T) {

				inputFilepath := filepath.Join(
					"testdata",
					fmt.Sprintf("%s-ccache-%s-redis-%s", ts.osAndVersionCode, ts.ccacheVersion, ts.redisVersion),
					tc.inputFilename,
				)
				input, err := os.ReadFile(inputFilepath)
				if err != nil {
					t.Fatalf("failed to open test input: %q", err)
				}

				s, err := ParseTSVStatistics(string(input))

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
