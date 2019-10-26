// Copyright 2018 VirtualTam.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package ccache

import (
	"time"

	"github.com/alecthomas/units"
)

// Statistics represents information about ccache configuration and usage.
type Statistics struct {
	CacheDirectory           string            `json:"cache_directory"`
	PrimaryConfig            string            `json:"primary_config"`
	SecondaryConfigReadonly  string            `json:"secondary_config_readonly"`
	StatsTime                time.Time         `json:"stats_time"`
	StatsZeroTime            time.Time         `json:"stats_zero_time"`
	CacheHitDirect           int               `json:"cache_hit_direct"`
	CacheHitPreprocessed     int               `json:"cache_hit_preprocessed"`
	CacheMiss                int               `json:"cache_miss"`
	CacheHitRate             float64           `json:"cache_hit_rate"`
	CacheHitRatio            float64           `json:"cache_hit_ratio"`
	CalledForLink            int               `json:"called_for_link"`
	CalledForPreprocessing   int               `json:"called_for_preprocessing"`
	UnsupportedCodeDirective int               `json:"unsupported_code_directive"`
	NoInputFile              int               `json:"no_input_file"`
	CleanupsPerformed        int               `json:"cleanups_performed"`
	FilesInCache             int               `json:"files_in_cache"`
	CacheSize                string            `json:"cache_size"`
	CacheSizeBytes           units.MetricBytes `json:"cache_size_bytes"`
	MaxCacheSize             string            `json:"max_cache_size"`
	MaxCacheSizeBytes        units.MetricBytes `json:"max_cache_size_bytes"`
}
