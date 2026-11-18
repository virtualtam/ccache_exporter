// Copyright VirtualTam 2018, 2026
// SPDX-License-Identifier: MIT

package metrics

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/prometheus/client_golang/prometheus/testutil"

	"github.com/virtualtam/ccache_exporter/v4/pkg/ccache"
)

const (
	testdataDir   = "ubuntu-26.04-ccache-4.12.3"
	ccacheVersion = "ccache version 4.12.3"
)

func TestCollectorUncacheableAndErrorCalls(t *testing.T) {
	ctx := t.Context()

	command := &fakeCommand{
		version: ccacheVersion,
		config:  readTestdata(t, "config"),
		stats:   readTestdata(t, "secondbuild.tsv"),
	}
	c := newCcacheCollector(ctx, ccache.NewWrapper(ctx, command))

	want := `
# HELP ccache_error_calls_total Calls ccache could not cache due to an error (total), by reason
# TYPE ccache_error_calls_total counter
ccache_error_calls_total{reason="bad_input_file"} 0
ccache_error_calls_total{reason="bad_output_file"} 0
ccache_error_calls_total{reason="compiler_check_failed"} 0
ccache_error_calls_total{reason="could_not_find_compiler"} 0
ccache_error_calls_total{reason="error_hashing_extra_file"} 0
ccache_error_calls_total{reason="internal_error"} 0
ccache_error_calls_total{reason="missing_cache_file"} 0
ccache_error_calls_total{reason="modified_input_file"} 0
# HELP ccache_uncacheable_calls_total Calls ccache could not cache (total), by reason
# TYPE ccache_uncacheable_calls_total counter
ccache_uncacheable_calls_total{reason="autoconf_test"} 0
ccache_uncacheable_calls_total{reason="bad_compiler_arguments"} 0
ccache_uncacheable_calls_total{reason="called_for_link"} 70
ccache_uncacheable_calls_total{reason="called_for_preprocessing"} 0
ccache_uncacheable_calls_total{reason="compile_failed"} 14
ccache_uncacheable_calls_total{reason="compiler_produced_empty_output"} 0
ccache_uncacheable_calls_total{reason="compiler_produced_no_output"} 0
ccache_uncacheable_calls_total{reason="compiler_produced_stdout"} 0
ccache_uncacheable_calls_total{reason="could_not_use_modules"} 0
ccache_uncacheable_calls_total{reason="could_not_use_precompiled_header"} 0
ccache_uncacheable_calls_total{reason="disabled"} 0
ccache_uncacheable_calls_total{reason="multiple_source_files"} 0
ccache_uncacheable_calls_total{reason="no_input_file"} 12
ccache_uncacheable_calls_total{reason="output_to_stdout"} 2
ccache_uncacheable_calls_total{reason="preprocessor_error"} 24
ccache_uncacheable_calls_total{reason="recache"} 0
ccache_uncacheable_calls_total{reason="unsupported_code_directive"} 0
ccache_uncacheable_calls_total{reason="unsupported_compiler_option"} 0
ccache_uncacheable_calls_total{reason="unsupported_environment_variable"} 0
ccache_uncacheable_calls_total{reason="unsupported_source_encoding"} 0
ccache_uncacheable_calls_total{reason="unsupported_source_language"} 0
`

	err := testutil.CollectAndCompare(c, strings.NewReader(want), "ccache_uncacheable_calls_total", "ccache_error_calls_total")
	if err != nil {
		t.Fatal(err)
	}
}

func TestCollectorRemovedUncacheableMetrics(t *testing.T) {
	ctx := t.Context()

	command := &fakeCommand{
		version: ccacheVersion,
		config:  readTestdata(t, "config"),
		stats:   readTestdata(t, "secondbuild.tsv"),
	}
	c := newCcacheCollector(ctx, ccache.NewWrapper(ctx, command))

	removed := []string{
		"ccache_called_for_link_total",
		"ccache_called_for_preprocessing_total",
		"ccache_compilation_failed_total",
		"ccache_no_input_file_total",
		"ccache_preprocessing_failed_total",
		"ccache_unsupported_code_directive_total",
	}

	for _, name := range removed {
		if n := testutil.CollectAndCount(c, name); n != 0 {
			t.Errorf("%s: want 0 series, got %d", name, n)
		}
	}
}

var _ ccache.Command = &fakeCommand{}

type fakeCommand struct {
	version string
	config  string
	stats   string
}

func (c *fakeCommand) PrintConfig(_ context.Context) (string, error) {
	return c.config, nil
}

func (c *fakeCommand) ShowConfig(_ context.Context) (string, error) {
	return c.config, nil
}

func (c *fakeCommand) PrintStats(_ context.Context) (string, error) {
	return c.stats, nil
}

func (c *fakeCommand) ShowStats(_ context.Context) (string, error) {
	return c.stats, nil
}

func (c *fakeCommand) Version(_ context.Context) (string, error) {
	return c.version, nil
}

func readTestdata(t *testing.T, name string) string {
	t.Helper()

	data, err := os.ReadFile(filepath.Join("..", "..", "..", "pkg", "ccache", "testdata", testdataDir, name))
	if err != nil {
		t.Fatalf("failed to open test input: %q", err)
	}

	return string(data)
}
