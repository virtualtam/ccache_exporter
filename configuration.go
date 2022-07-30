package ccache

import "github.com/alecthomas/units"

// Configuration represents information about ccache configuration.
type Configuration struct {
	CacheDirectory          string            `json:"cache_directory"`
	PrimaryConfig           string            `json:"primary_config"`
	SecondaryConfigReadonly string            `json:"secondary_config_readonly"`
	MaxCacheSize            string            `json:"max_cache_size"`
	MaxCacheSizeBytes       units.MetricBytes `json:"max_cache_size_bytes"`
}
