// Copyright VirtualTam 2018, 2026
// SPDX-License-Identifier: MIT

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
						CacheHitRate:          50,
						CacheHitRatio:         0.5,
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
						CacheHitRate:           38.013699,
						CacheHitRatio:          0.380137,
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
			osAndVersion:     "Debian 13",
			osAndVersionCode: "debian-13",
			ccacheVersion:    "4.11.2",

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
						CacheMiss:              137,
						CacheMissDirect:        156,
						CacheMissPreprocessed:  144,
						CalledForLink:          36,
						CalledForPreprocessing: 0,
						CompilationFailed:      7,
						NoInputFile:            5,
						PreprocessingFailed:    12,
						FilesInCache:           272,
						CacheSize:              "33MB",
						CacheSizeBytes:         units.MetricBytes(33853440),
					},
				},
				{
					tname:         "second build",
					inputFilename: "secondbuild.tsv",
					wantStats: Statistics{
						CacheHitDirect:         104,
						CacheHitPreprocessed:   2,
						CacheMiss:              168,
						CacheMissDirect:        208,
						CacheMissPreprocessed:  182,
						CacheHitRate:           38.686131,
						CacheHitRatio:          0.386861,
						CalledForLink:          72,
						CalledForPreprocessing: 0,
						CompilationFailed:      14,
						NoInputFile:            10,
						PreprocessingFailed:    24,
						FilesInCache:           334,
						CacheSize:              "34MB",
						CacheSizeBytes:         units.MetricBytes(34193408),
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
						CacheHitRate:           37.346437,
						CacheHitRatio:          0.373464,
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
						CacheHitRate:          40.136054,
						CacheHitRatio:         0.401361,
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
		{
			osAndVersion:     "Ubuntu 26.04",
			osAndVersionCode: "ubuntu-26.04",
			ccacheVersion:    "4.12.3",

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
						CacheMiss:              140,
						CacheMissDirect:        159,
						CacheMissPreprocessed:  147,
						CalledForLink:          35,
						CalledForPreprocessing: 0,
						CompilationFailed:      7,
						NoInputFile:            6,
						PreprocessingFailed:    12,
						FilesInCache:           277,
						CacheSize:              "82MB",
						CacheSizeBytes:         units.MetricBytes(82989056),
					},
				},
				{
					tname:         "second build",
					inputFilename: "secondbuild.tsv",
					wantStats: Statistics{
						CacheHitDirect:         107,
						CacheHitPreprocessed:   3,
						CacheMiss:              170,
						CacheMissDirect:        211,
						CacheMissPreprocessed:  184,
						CacheHitRate:           39.285714,
						CacheHitRatio:          0.392857,
						CalledForLink:          70,
						CalledForPreprocessing: 0,
						CompilationFailed:      14,
						NoInputFile:            12,
						PreprocessingFailed:    24,
						FilesInCache:           337,
						CacheSize:              "83MB",
						CacheSizeBytes:         units.MetricBytes(83460096),
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
			osAndVersion:     "Debian 13",
			osAndVersionCode: "debian-13",
			ccacheVersion:    "4.11.2",
			redisVersion:     "8",

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
						CacheMiss:             137,
						CacheMissDirect:       156,
						CacheMissPreprocessed: 144,
						CacheSize:             "0B",
						CalledForLink:         36,
						CompilationFailed:     7,
						NoInputFile:           5,
						PreprocessingFailed:   12,
						RemoteStorageMiss:     137,
						RemoteStorageReadMiss: 298,
						RemoteStorageWrite:    272,
					},
				},
				{
					tname:         "second build",
					inputFilename: "secondbuild.tsv",
					wantStats: Statistics{
						CacheHitDirect:        104,
						CacheHitPreprocessed:  2,
						CacheMiss:             168,
						CacheMissDirect:       208,
						CacheMissPreprocessed: 182,
						CacheHitRate:          38.686131,
						CacheHitRatio:         0.386861,
						CalledForLink:         72,
						CacheSize:             "0B",
						CompilationFailed:     14,
						NoInputFile:           10,
						PreprocessingFailed:   24,
						RemoteStorageHit:      106,
						RemoteStorageMiss:     168,
						RemoteStorageReadHit:  210,
						RemoteStorageReadMiss: 386,
						RemoteStorageWrite:    334,
					},
				},
			},
		},
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
						CacheHitRate:          40.136054,
						CacheHitRatio:         0.401361,
						CalledForLink:         70,
						CacheSize:             "0B",
						CompilationFailed:     14,
						NoInputFile:           30,
						PreprocessingFailed:   6,
						RemoteStorageHit:      118,
						RemoteStorageMiss:     176,
						RemoteStorageReadHit:  234,
						RemoteStorageReadMiss: 384,
						RemoteStorageWrite:    350,
					},
				},
			},
		},
		{
			osAndVersion:     "Ubuntu 26.04",
			osAndVersionCode: "ubuntu-26.04",
			ccacheVersion:    "4.12.3",
			redisVersion:     "8",

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
						CacheMiss:             140,
						CacheMissDirect:       159,
						CacheMissPreprocessed: 147,
						CacheSize:             "0B",
						CalledForLink:         35,
						CompilationFailed:     7,
						NoInputFile:           6,
						PreprocessingFailed:   12,
						RemoteStorageMiss:     140,
						RemoteStorageReadMiss: 304,
						RemoteStorageWrite:    277,
					},
				},
				{
					tname:         "second build",
					inputFilename: "secondbuild.tsv",
					wantStats: Statistics{
						CacheHitDirect:        107,
						CacheHitPreprocessed:  3,
						CacheMiss:             170,
						CacheMissDirect:       211,
						CacheMissPreprocessed: 184,
						CacheHitRate:          39.285714,
						CacheHitRatio:         0.392857,
						CalledForLink:         70,
						CacheSize:             "0B",
						CompilationFailed:     14,
						NoInputFile:           12,
						PreprocessingFailed:   24,
						RemoteStorageHit:      110,
						RemoteStorageMiss:     170,
						RemoteStorageReadHit:  217,
						RemoteStorageReadMiss: 391,
						RemoteStorageWrite:    337,
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

func TestParseTSVStatisticsEdgeCases(t *testing.T) {
	cases := []struct {
		tname     string
		input     string
		wantStats Statistics
		wantErr   error
	}{
		// cache eviction
		{
			tname: "cleanups performed",
			input: "cleanups_performed\t3\n",
			wantStats: Statistics{
				CleanupsPerformed: 3,
				CacheSize:         "0B",
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.tname, func(t *testing.T) {
			s, err := ParseTSVStatistics(tc.input)

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
