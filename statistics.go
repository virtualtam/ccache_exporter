// Copyright 2018-2022 VirtualTam.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

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
	CacheHitDirect         int     `json:"cache_hit_direct"`
	CacheHitPreprocessed   int     `json:"cache_hit_preprocessed"`
	CacheMiss              int     `json:"cache_miss"`
	CacheMissDirect        int     `json:"cache_miss_direct"`
	CacheMissPreprocessed  int     `json:"cache_miss_preprocessed"`
	CacheHitRate           float64 `json:"cache_hit_rate"`
	CacheHitRatio          float64 `json:"cache_hit_ratio"`
	CalledForLink          int     `json:"called_for_link"`
	CalledForPreprocessing int     `json:"called_for_preprocessing"`

	// Uncacheable
	CompilationFailed        int `json:"compilation_failed"`
	PreprocessingFailed      int `json:"preprocessing_failed"`
	UnsupportedCodeDirective int `json:"unsupported_code_directive"`
	NoInputFile              int `json:"no_input_file"`

	// Remote storage
	RemoteStorageError    int `json:"remote_storage_error"`
	RemoteStorageHit      int `json:"remote_storage_hit"`
	RemoteStorageMiss     int `json:"remote_storage_miss"`
	RemoteStorageReadHit  int `json:"remote_storage_read_hit"`
	RemoteStorageReadMiss int `json:"remote_storage_read_miss"`
	RemoteStorageTimeout  int `json:"remote_storage_timeout"`
	RemoteStorageWrite    int `json:"remote_storage_write"`
}
