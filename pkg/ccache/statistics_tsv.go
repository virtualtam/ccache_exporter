// Copyright (c) VirtualTam
// SPDX-License-Identifier: MIT

package ccache

import (
	"encoding/csv"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/alecthomas/units"
)

// ParseTSVStatistics reads ccache statistics as formatted by the `ccache --print-stats` command.
//
// It relies upon the `ccache --print-stats` command to output machine-readable
// statistics.
func ParseTSVStatistics(text string) (*Statistics, error) {
	r := strings.NewReader(text)

	tsvReader := csv.NewReader(r)
	tsvReader.Comma = '\t'

	tsvData, err := tsvReader.ReadAll()
	if err != nil {
		return &Statistics{}, err
	}

	stats := &Statistics{}

	for _, row := range tsvData {
		if len(row) != 2 {
			// for each row, we expect a key and a value
			continue
		}

		switch row[0] {
		case "cache_miss":
			stats.CacheMiss, err = strconv.Atoi(row[1])
			if err != nil {
				return &Statistics{}, err
			}
		case "cache_size_kibibyte":
			cacheSizeBytes, err := units.ParseBase2Bytes(fmt.Sprintf("%sKiB", row[1]))
			if err != nil {
				return &Statistics{}, err
			}
			stats.CacheSizeBytes = units.MetricBytes(cacheSizeBytes)

		case "called_for_link":
			stats.CalledForLink, err = strconv.Atoi(row[1])
			if err != nil {
				return &Statistics{}, err
			}

		case "called_for_preprocessing":
			stats.CalledForPreprocessing, err = strconv.Atoi(row[1])
			if err != nil {
				return &Statistics{}, err
			}

		case "compile_failed":
			stats.CompilationFailed, err = strconv.Atoi(row[1])
			if err != nil {
				return &Statistics{}, err
			}

		case "direct_cache_hit":
			stats.CacheHitDirect, err = strconv.Atoi(row[1])
			if err != nil {
				return &Statistics{}, err
			}

		case "direct_cache_miss":
			stats.CacheMissDirect, err = strconv.Atoi(row[1])
			if err != nil {
				return &Statistics{}, err
			}

		case "files_in_cache":
			stats.FilesInCache, err = strconv.Atoi(row[1])
			if err != nil {
				return &Statistics{}, err
			}

		case "no_input_file":
			stats.NoInputFile, err = strconv.Atoi(row[1])
			if err != nil {
				return &Statistics{}, err
			}

		case "preprocessed_cache_hit":
			stats.CacheHitPreprocessed, err = strconv.Atoi(row[1])
			if err != nil {
				return &Statistics{}, err
			}

		case "preprocessed_cache_miss":
			stats.CacheMissPreprocessed, err = strconv.Atoi(row[1])
			if err != nil {
				return &Statistics{}, err
			}

		case "preprocessor_error":
			stats.PreprocessingFailed, err = strconv.Atoi(row[1])
			if err != nil {
				return &Statistics{}, err
			}

		case "remote_storage_error":
			stats.RemoteStorageError, err = strconv.Atoi(row[1])
			if err != nil {
				return &Statistics{}, err
			}

		case "remote_storage_hit":
			stats.RemoteStorageHit, err = strconv.Atoi(row[1])
			if err != nil {
				return &Statistics{}, err
			}

		case "remote_storage_miss":
			stats.RemoteStorageMiss, err = strconv.Atoi(row[1])
			if err != nil {
				return &Statistics{}, err
			}

		case "remote_storage_read_hit":
			stats.RemoteStorageReadHit, err = strconv.Atoi(row[1])
			if err != nil {
				return &Statistics{}, err
			}

		case "remote_storage_read_miss":
			stats.RemoteStorageReadMiss, err = strconv.Atoi(row[1])
			if err != nil {
				return &Statistics{}, err
			}

		case "remote_storage_timeout":
			stats.RemoteStorageTimeout, err = strconv.Atoi(row[1])
			if err != nil {
				return &Statistics{}, err
			}

		case "remote_storage_write":
			stats.RemoteStorageWrite, err = strconv.Atoi(row[1])
			if err != nil {
				return &Statistics{}, err
			}

		case "stats_updated_timestamp":
			unixTime, err := strconv.ParseInt(row[1], 10, 64)
			if err != nil {
				return &Statistics{}, err
			}
			stats.StatsTime = time.Unix(unixTime, 0).UTC()

		case "stats_zeroed_timestamp":
			unixTime, err := strconv.ParseInt(row[1], 10, 64)
			if err != nil {
				return &Statistics{}, err
			}
			stats.StatsZeroTime = time.Unix(unixTime, 0).UTC()

		case "unsupported_code_directive":
			stats.UnsupportedCodeDirective, err = strconv.Atoi(row[1])
			if err != nil {
				return &Statistics{}, err
			}
		}
	}

	// Compute fields for compatibility
	//
	// FIXME: "ccache --show-stats" returns seemingly incoherent values
	//        regarding cache hits and misses, making it hard to tell how ratios
	//        should be computed
	cacheHitTotal := stats.CacheHitDirect + stats.CacheHitPreprocessed
	cacheCallTotal := stats.CacheHitDirect + stats.CacheHitPreprocessed + stats.CacheMiss + stats.CacheMissDirect + stats.CacheMissPreprocessed

	if cacheCallTotal > 0 {
		stats.CacheHitRatio = float64(cacheHitTotal) / float64(cacheCallTotal)
		stats.CacheHitRate = 100 * stats.CacheHitRatio
	}
	stats.CacheSize = stats.CacheSizeBytes.Floor().String()

	return stats, nil
}
