// Copyright VirtualTam 2018, 2026
// SPDX-License-Identifier: MIT

package ccache

import (
	"time"

	"github.com/alecthomas/units"
)

// Statistics represents information about ccache usage.
type Statistics struct {
	// Cache status
	CleanupsPerformed int               `json:"cleanups_performed"`
	FilesInCache      int               `json:"files_in_cache"`
	CacheSize         string            `json:"cache_size"`
	CacheSizeBytes    units.MetricBytes `json:"cache_size_bytes"`

	// Timestamps
	StatsTime     time.Time `json:"stats_time"`
	StatsZeroTime time.Time `json:"stats_zero_time"`

	// Cache usage
	CacheHitDirect        int     `json:"cache_hit_direct"`
	CacheHitPreprocessed  int     `json:"cache_hit_preprocessed"`
	CacheMiss             int     `json:"cache_miss"`
	CacheMissDirect       int     `json:"cache_miss_direct"`
	CacheMissPreprocessed int     `json:"cache_miss_preprocessed"`
	CacheHitRate          float64 `json:"cache_hit_rate"`
	CacheHitRatio         float64 `json:"cache_hit_ratio"`

	// Calls ccache could not cache
	UncacheableCalls UncacheableCalls `json:"uncacheable_calls"`
	ErrorCalls       ErrorCalls       `json:"error_calls"`

	// Remote storage
	RemoteStorageError    int `json:"remote_storage_error"`
	RemoteStorageHit      int `json:"remote_storage_hit"`
	RemoteStorageMiss     int `json:"remote_storage_miss"`
	RemoteStorageReadHit  int `json:"remote_storage_read_hit"`
	RemoteStorageReadMiss int `json:"remote_storage_read_miss"`
	RemoteStorageTimeout  int `json:"remote_storage_timeout"`
	RemoteStorageWrite    int `json:"remote_storage_write"`
}

// UncacheableCalls counts the calls ccache could not cache, by reason.
type UncacheableCalls struct {
	AutoconfTest                   int `json:"autoconf_test"`
	BadCompilerArguments           int `json:"bad_compiler_arguments"`
	CalledForLink                  int `json:"called_for_link"`
	CalledForPreprocessing         int `json:"called_for_preprocessing"`
	CompileFailed                  int `json:"compile_failed"`
	CompilerProducedEmptyOutput    int `json:"compiler_produced_empty_output"`
	CompilerProducedNoOutput       int `json:"compiler_produced_no_output"`
	CompilerProducedStdout         int `json:"compiler_produced_stdout"`
	CouldNotUseModules             int `json:"could_not_use_modules"`
	CouldNotUsePrecompiledHeader   int `json:"could_not_use_precompiled_header"`
	Disabled                       int `json:"disabled"`
	MultipleSourceFiles            int `json:"multiple_source_files"`
	NoInputFile                    int `json:"no_input_file"`
	OutputToStdout                 int `json:"output_to_stdout"`
	PreprocessorError              int `json:"preprocessor_error"`
	Recache                        int `json:"recache"`
	UnsupportedCodeDirective       int `json:"unsupported_code_directive"`
	UnsupportedCompilerOption      int `json:"unsupported_compiler_option"`
	UnsupportedEnvironmentVariable int `json:"unsupported_environment_variable"`
	UnsupportedSourceEncoding      int `json:"unsupported_source_encoding"`
	UnsupportedSourceLanguage      int `json:"unsupported_source_language"`
}

// ErrorCalls counts the calls ccache could not cache due to an error, by reason.
type ErrorCalls struct {
	BadInputFile          int `json:"bad_input_file"`
	BadOutputFile         int `json:"bad_output_file"`
	CompilerCheckFailed   int `json:"compiler_check_failed"`
	CouldNotFindCompiler  int `json:"could_not_find_compiler"`
	ErrorHashingExtraFile int `json:"error_hashing_extra_file"`
	InternalError         int `json:"internal_error"`
	MissingCacheFile      int `json:"missing_cache_file"`
	ModifiedInputFile     int `json:"modified_input_file"`
}
