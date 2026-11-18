// Copyright VirtualTam 2018, 2026
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
		case "autoconf_test":
			stats.UncacheableCalls.AutoconfTest, err = strconv.Atoi(row[1])
			if err != nil {
				return &Statistics{}, err
			}

		case "bad_compiler_arguments":
			stats.UncacheableCalls.BadCompilerArguments, err = strconv.Atoi(row[1])
			if err != nil {
				return &Statistics{}, err
			}

		case "bad_input_file":
			stats.ErrorCalls.BadInputFile, err = strconv.Atoi(row[1])
			if err != nil {
				return &Statistics{}, err
			}

		case "bad_output_file":
			stats.ErrorCalls.BadOutputFile, err = strconv.Atoi(row[1])
			if err != nil {
				return &Statistics{}, err
			}

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
			stats.UncacheableCalls.CalledForLink, err = strconv.Atoi(row[1])
			if err != nil {
				return &Statistics{}, err
			}

		case "called_for_preprocessing":
			stats.UncacheableCalls.CalledForPreprocessing, err = strconv.Atoi(row[1])
			if err != nil {
				return &Statistics{}, err
			}

		case "cleanups_performed":
			stats.CleanupsPerformed, err = strconv.Atoi(row[1])
			if err != nil {
				return &Statistics{}, err
			}

		case "compile_failed":
			stats.UncacheableCalls.CompileFailed, err = strconv.Atoi(row[1])
			if err != nil {
				return &Statistics{}, err
			}

		case "compiler_check_failed":
			stats.ErrorCalls.CompilerCheckFailed, err = strconv.Atoi(row[1])
			if err != nil {
				return &Statistics{}, err
			}

		case "compiler_produced_empty_output":
			stats.UncacheableCalls.CompilerProducedEmptyOutput, err = strconv.Atoi(row[1])
			if err != nil {
				return &Statistics{}, err
			}

		case "compiler_produced_no_output":
			stats.UncacheableCalls.CompilerProducedNoOutput, err = strconv.Atoi(row[1])
			if err != nil {
				return &Statistics{}, err
			}

		case "compiler_produced_stdout":
			stats.UncacheableCalls.CompilerProducedStdout, err = strconv.Atoi(row[1])
			if err != nil {
				return &Statistics{}, err
			}

		case "could_not_find_compiler":
			stats.ErrorCalls.CouldNotFindCompiler, err = strconv.Atoi(row[1])
			if err != nil {
				return &Statistics{}, err
			}

		case "could_not_use_modules":
			stats.UncacheableCalls.CouldNotUseModules, err = strconv.Atoi(row[1])
			if err != nil {
				return &Statistics{}, err
			}

		case "could_not_use_precompiled_header":
			stats.UncacheableCalls.CouldNotUsePrecompiledHeader, err = strconv.Atoi(row[1])
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

		case "disabled":
			stats.UncacheableCalls.Disabled, err = strconv.Atoi(row[1])
			if err != nil {
				return &Statistics{}, err
			}

		case "error_hashing_extra_file":
			stats.ErrorCalls.ErrorHashingExtraFile, err = strconv.Atoi(row[1])
			if err != nil {
				return &Statistics{}, err
			}

		case "files_in_cache":
			stats.FilesInCache, err = strconv.Atoi(row[1])
			if err != nil {
				return &Statistics{}, err
			}

		case "internal_error":
			stats.ErrorCalls.InternalError, err = strconv.Atoi(row[1])
			if err != nil {
				return &Statistics{}, err
			}

		case "missing_cache_file":
			stats.ErrorCalls.MissingCacheFile, err = strconv.Atoi(row[1])
			if err != nil {
				return &Statistics{}, err
			}

		case "modified_input_file":
			stats.ErrorCalls.ModifiedInputFile, err = strconv.Atoi(row[1])
			if err != nil {
				return &Statistics{}, err
			}

		case "multiple_source_files":
			stats.UncacheableCalls.MultipleSourceFiles, err = strconv.Atoi(row[1])
			if err != nil {
				return &Statistics{}, err
			}

		case "no_input_file":
			stats.UncacheableCalls.NoInputFile, err = strconv.Atoi(row[1])
			if err != nil {
				return &Statistics{}, err
			}

		case "output_to_stdout":
			stats.UncacheableCalls.OutputToStdout, err = strconv.Atoi(row[1])
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
			stats.UncacheableCalls.PreprocessorError, err = strconv.Atoi(row[1])
			if err != nil {
				return &Statistics{}, err
			}

		case "recache":
			stats.UncacheableCalls.Recache, err = strconv.Atoi(row[1])
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
			stats.UncacheableCalls.UnsupportedCodeDirective, err = strconv.Atoi(row[1])
			if err != nil {
				return &Statistics{}, err
			}

		case "unsupported_compiler_option":
			stats.UncacheableCalls.UnsupportedCompilerOption, err = strconv.Atoi(row[1])
			if err != nil {
				return &Statistics{}, err
			}

		case "unsupported_environment_variable":
			stats.UncacheableCalls.UnsupportedEnvironmentVariable, err = strconv.Atoi(row[1])
			if err != nil {
				return &Statistics{}, err
			}

		case "unsupported_source_encoding":
			stats.UncacheableCalls.UnsupportedSourceEncoding, err = strconv.Atoi(row[1])
			if err != nil {
				return &Statistics{}, err
			}

		case "unsupported_source_language":
			stats.UncacheableCalls.UnsupportedSourceLanguage, err = strconv.Atoi(row[1])
			if err != nil {
				return &Statistics{}, err
			}
		}
	}

	cacheHitTotal := stats.CacheHitDirect + stats.CacheHitPreprocessed
	cacheCallTotal := cacheHitTotal + stats.CacheMiss

	if cacheCallTotal > 0 {
		// Compute the hit ratio like `ccache --show-stats`: hits / (hits + cache_miss)
		//
		// Direct and preprocessed misses count lookup attempts (not calls); a direct mode miss can
		// still be a preprocessed hit.
		stats.CacheHitRatio = float64(cacheHitTotal) / float64(cacheCallTotal)
		stats.CacheHitRate = 100 * stats.CacheHitRatio
	}
	stats.CacheSize = stats.CacheSizeBytes.Floor().String()

	return stats, nil
}
