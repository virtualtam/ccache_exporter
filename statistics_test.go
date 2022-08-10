// Copyright 2018-2022 VirtualTam.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package ccache

import (
	"math"
	"testing"

	"github.com/alecthomas/units"
)

func assertStatisticsEqual(t *testing.T, got, want *Statistics) {
	t.Helper()

	assertIntFieldEquals(t, "CacheHitDirect", got.CacheHitDirect, want.CacheHitDirect)
	assertIntFieldEquals(t, "CacheHitPreprocessed", got.CacheHitPreprocessed, want.CacheHitPreprocessed)
	assertIntFieldEquals(t, "CacheMiss", got.CacheMiss, want.CacheMiss)
	assertIntFieldEquals(t, "CacheMissDirect", got.CacheMissDirect, want.CacheMissDirect)
	assertIntFieldEquals(t, "CacheMissPreprocessed", got.CacheMissPreprocessed, want.CacheMissPreprocessed)
	assertFloatFieldAlmostEquals(t, "CacheHitRate", got.CacheHitRate, want.CacheHitRate)
	assertFloatFieldAlmostEquals(t, "CacheHitRatio", got.CacheHitRatio, want.CacheHitRatio)
	assertIntFieldEquals(t, "CalledForLink", got.CalledForLink, want.CalledForLink)
	assertIntFieldEquals(t, "CalledForPreprocessing", got.CalledForPreprocessing, want.CalledForPreprocessing)
	assertIntFieldEquals(t, "CompilationFailed", got.CompilationFailed, want.CompilationFailed)
	assertIntFieldEquals(t, "PreprocessingFailed", got.PreprocessingFailed, want.PreprocessingFailed)
	assertIntFieldEquals(t, "UnsupportedCodeDirective", got.UnsupportedCodeDirective, want.UnsupportedCodeDirective)
	assertIntFieldEquals(t, "NoInputFile", got.NoInputFile, want.NoInputFile)
	assertIntFieldEquals(t, "CleanupsPerformed", got.CleanupsPerformed, want.CleanupsPerformed)
	assertIntFieldEquals(t, "FilesInCache", got.FilesInCache, want.FilesInCache)
	assertStringFieldEquals(t, "CacheSize", got.CacheSize, want.CacheSize)
	assertMetricByteFieldEquals(t, "CacheSizeBytes", got.CacheSizeBytes, want.CacheSizeBytes)
}

func assertFloatFieldAlmostEquals(t *testing.T, fieldName string, got, want float64) {
	t.Helper()
	if math.Abs(got-want) > 0.01 {
		t.Errorf("%s: want %f, got %f", fieldName, want, got)
	}
}

func assertIntFieldEquals(t *testing.T, fieldName string, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("%s: want %d, got %d", fieldName, want, got)
	}
}

func assertMetricByteFieldEquals(t *testing.T, fieldName string, got, want units.MetricBytes) {
	t.Helper()
	if got != want {
		t.Errorf("%s: want %d, got %d", fieldName, want, got)
	}
}

func assertStringFieldEquals(t *testing.T, fieldName, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("%s: want %q, got %q", fieldName, want, got)
	}
}
