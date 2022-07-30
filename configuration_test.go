package ccache

import "testing"

func assertConfigurationsEqual(t *testing.T, got, want *Configuration) {
	t.Helper()

	assertStringFieldEquals(t, "CacheDirectory", got.CacheDirectory, want.CacheDirectory)
	assertStringFieldEquals(t, "PrimaryConfig", got.PrimaryConfig, want.PrimaryConfig)
	assertStringFieldEquals(t, "SecondaryConfigReadonly", got.SecondaryConfigReadonly, want.SecondaryConfigReadonly)
	assertStringFieldEquals(t, "MaxCacheSize", got.MaxCacheSize, want.MaxCacheSize)
	assertMetricByteFieldEquals(t, "MaxCacheSizeBytes", got.MaxCacheSizeBytes, want.MaxCacheSizeBytes)
}
