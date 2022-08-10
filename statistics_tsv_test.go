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
	testCases        []tsvTestCase
}

func TestParseTSVStatistics(t *testing.T) {
	sessions := []tsvTestSession{
		{
			osAndVersion:     "Arch Linux",
			osAndVersionCode: "arch",
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
						CacheHitRate:          32.188841,
						CacheHitRatio:         0.321888,
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
	}

	for _, ts := range sessions {
		for _, tc := range ts.testCases {
			t.Run(fmt.Sprintf("ccache %s on %s (%s)", ts.ccacheVersion, ts.osAndVersion, tc.tname), func(t *testing.T) {
				inputFilepath := filepath.Join(
					"testdata",
					fmt.Sprintf("%s-%s", ts.osAndVersionCode, ts.ccacheVersion),
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
