// Copyright VirtualTam 2018, 2026
// SPDX-License-Identifier: MIT

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
	assertFloatFieldAlmostEquals(t, "CacheHitRate", got.CacheHitRate, want.CacheHitRate)
	assertFloatFieldAlmostEquals(t, "CacheHitRatio", got.CacheHitRatio, want.CacheHitRatio)

	assertIntFieldEquals(t, "CacheMiss", got.CacheMiss, want.CacheMiss)
	assertIntFieldEquals(t, "CacheMissDirect", got.CacheMissDirect, want.CacheMissDirect)
	assertIntFieldEquals(t, "CacheMissPreprocessed", got.CacheMissPreprocessed, want.CacheMissPreprocessed)

	assertIntFieldEquals(t, "CleanupsPerformed", got.CleanupsPerformed, want.CleanupsPerformed)
	assertIntFieldEquals(t, "FilesInCache", got.FilesInCache, want.FilesInCache)

	assertStringFieldEquals(t, "CacheSize", got.CacheSize, want.CacheSize)
	assertMetricByteFieldEquals(t, "CacheSizeBytes", got.CacheSizeBytes, want.CacheSizeBytes)

	assertIntFieldEquals(t, "RemoteStorageError", got.RemoteStorageError, want.RemoteStorageError)
	assertIntFieldEquals(t, "RemoteStorageHit", got.RemoteStorageHit, want.RemoteStorageHit)
	assertIntFieldEquals(t, "RemoteStorageMiss", got.RemoteStorageMiss, want.RemoteStorageMiss)
	assertIntFieldEquals(t, "RemoteStorageReadHit", got.RemoteStorageReadHit, want.RemoteStorageReadHit)
	assertIntFieldEquals(t, "RemoteStorageReadMiss", got.RemoteStorageReadMiss, want.RemoteStorageReadMiss)
	assertIntFieldEquals(t, "RemoteStorageTimeout", got.RemoteStorageTimeout, want.RemoteStorageTimeout)
	assertIntFieldEquals(t, "RemoteStorageWrite", got.RemoteStorageWrite, want.RemoteStorageWrite)

	assertUncacheableCallsEqual(t, got.UncacheableCalls, want.UncacheableCalls)
	assertErrorCallsEqual(t, got.ErrorCalls, want.ErrorCalls)
}

func assertUncacheableCallsEqual(t *testing.T, got, want UncacheableCalls) {
	t.Helper()

	assertIntFieldEquals(t, "UncacheableCalls.AutoconfTest", got.AutoconfTest, want.AutoconfTest)
	assertIntFieldEquals(t, "UncacheableCalls.BadCompilerArguments", got.BadCompilerArguments, want.BadCompilerArguments)
	assertIntFieldEquals(t, "UncacheableCalls.CalledForLink", got.CalledForLink, want.CalledForLink)
	assertIntFieldEquals(t, "UncacheableCalls.CalledForPreprocessing", got.CalledForPreprocessing, want.CalledForPreprocessing)
	assertIntFieldEquals(t, "UncacheableCalls.CompileFailed", got.CompileFailed, want.CompileFailed)
	assertIntFieldEquals(t, "UncacheableCalls.CompilerProducedEmptyOutput", got.CompilerProducedEmptyOutput, want.CompilerProducedEmptyOutput)
	assertIntFieldEquals(t, "UncacheableCalls.CompilerProducedNoOutput", got.CompilerProducedNoOutput, want.CompilerProducedNoOutput)
	assertIntFieldEquals(t, "UncacheableCalls.CompilerProducedStdout", got.CompilerProducedStdout, want.CompilerProducedStdout)
	assertIntFieldEquals(t, "UncacheableCalls.CouldNotUseModules", got.CouldNotUseModules, want.CouldNotUseModules)
	assertIntFieldEquals(t, "UncacheableCalls.CouldNotUsePrecompiledHeader", got.CouldNotUsePrecompiledHeader, want.CouldNotUsePrecompiledHeader)
	assertIntFieldEquals(t, "UncacheableCalls.Disabled", got.Disabled, want.Disabled)
	assertIntFieldEquals(t, "UncacheableCalls.MultipleSourceFiles", got.MultipleSourceFiles, want.MultipleSourceFiles)
	assertIntFieldEquals(t, "UncacheableCalls.NoInputFile", got.NoInputFile, want.NoInputFile)
	assertIntFieldEquals(t, "UncacheableCalls.OutputToStdout", got.OutputToStdout, want.OutputToStdout)
	assertIntFieldEquals(t, "UncacheableCalls.PreprocessorError", got.PreprocessorError, want.PreprocessorError)
	assertIntFieldEquals(t, "UncacheableCalls.Recache", got.Recache, want.Recache)
	assertIntFieldEquals(t, "UncacheableCalls.UnsupportedCodeDirective", got.UnsupportedCodeDirective, want.UnsupportedCodeDirective)
	assertIntFieldEquals(t, "UncacheableCalls.UnsupportedCompilerOption", got.UnsupportedCompilerOption, want.UnsupportedCompilerOption)
	assertIntFieldEquals(t, "UncacheableCalls.UnsupportedEnvironmentVariable", got.UnsupportedEnvironmentVariable, want.UnsupportedEnvironmentVariable)
	assertIntFieldEquals(t, "UncacheableCalls.UnsupportedSourceEncoding", got.UnsupportedSourceEncoding, want.UnsupportedSourceEncoding)
	assertIntFieldEquals(t, "UncacheableCalls.UnsupportedSourceLanguage", got.UnsupportedSourceLanguage, want.UnsupportedSourceLanguage)
}

func assertErrorCallsEqual(t *testing.T, got, want ErrorCalls) {
	t.Helper()

	assertIntFieldEquals(t, "ErrorCalls.BadInputFile", got.BadInputFile, want.BadInputFile)
	assertIntFieldEquals(t, "ErrorCalls.BadOutputFile", got.BadOutputFile, want.BadOutputFile)
	assertIntFieldEquals(t, "ErrorCalls.CompilerCheckFailed", got.CompilerCheckFailed, want.CompilerCheckFailed)
	assertIntFieldEquals(t, "ErrorCalls.CouldNotFindCompiler", got.CouldNotFindCompiler, want.CouldNotFindCompiler)
	assertIntFieldEquals(t, "ErrorCalls.ErrorHashingExtraFile", got.ErrorHashingExtraFile, want.ErrorHashingExtraFile)
	assertIntFieldEquals(t, "ErrorCalls.InternalError", got.InternalError, want.InternalError)
	assertIntFieldEquals(t, "ErrorCalls.MissingCacheFile", got.MissingCacheFile, want.MissingCacheFile)
	assertIntFieldEquals(t, "ErrorCalls.ModifiedInputFile", got.ModifiedInputFile, want.ModifiedInputFile)
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
