// Copyright 2018-2022 VirtualTam.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package ccache

import (
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/alecthomas/units"
)

var rules = map[string]*regexp.Regexp{
	"cacheDirectory":           regexp.MustCompile(`cache directory\s+(.+)`),
	"primaryConfig":            regexp.MustCompile(`primary config\s+(.+)`),
	"secondaryConfigReadonly":  regexp.MustCompile(`secondary config\s+(\(readonly\)\s+)?(.+)`),
	"statsZeroTime":            regexp.MustCompile(`stats zero( time|ed)\s+(.*)`),
	"cacheHitDirect":           regexp.MustCompile(`cache hit \(direct\)\s+(\d+)`),
	"cacheHitPreprocessed":     regexp.MustCompile(`cache hit \(preprocessed\)\s+(\d+)`),
	"cacheMiss":                regexp.MustCompile(`cache miss\s+(\d+)`),
	"cacheHitRate":             regexp.MustCompile(`cache hit rate\s+(\d+(\.\d+)?) %`),
	"calledForLink":            regexp.MustCompile(`called for link\s+(\d+)`),
	"calledForPreprocessing":   regexp.MustCompile(`called for preprocessing\s+(\d+)`),
	"unsupportedCodeDirective": regexp.MustCompile(`unsupported code directive\s+(\d+)`),
	"noInputFile":              regexp.MustCompile(`no input file\s+(\d+)`),
	"cleanupsPerformed":        regexp.MustCompile(`cleanups performed\s+(\d+)`),
	"filesInCache":             regexp.MustCompile(`files in cache\s+(\d+)`),
	"cacheSize":                regexp.MustCompile(`cache size\s+(.+)`),
	"maxCacheSize":             regexp.MustCompile(`max cache size\s+(.+)`),
}

// LegacyParser provides support for ccache versions 3.3 to 3.6 (included).
//
// It relies upon the `ccache --show-stats` command to ouptut machine-readable
// statistics.
//
// Starting with ccache 3.7, this command was overhauled to print human-readable
// statistics, with `ccache --print-stats` being the new command to get
// machine-readable statistics.
type LegacyParser struct{}

// NewLegacyParser initializes and returns a new LegacyParser.
func NewLegacyParser() *LegacyParser {
	return &LegacyParser{}
}

// ParseShowStats reads ccache configuration and statistics as formatted by the `ccache --show-stats` command.
func (p *LegacyParser) ParseShowStats(text string) (*Configuration, *Statistics, error) {
	config := &Configuration{}
	stats := &Statistics{}
	var err error

	matches := rules["cacheDirectory"].FindStringSubmatch(text)
	if len(matches) == 2 {
		config.CacheDirectory = matches[1]
	}

	matches = rules["primaryConfig"].FindStringSubmatch(text)
	if len(matches) == 2 {
		config.PrimaryConfig = matches[1]
	}

	matches = rules["secondaryConfigReadonly"].FindStringSubmatch(text)
	if len(matches) == 2 {
		config.SecondaryConfigReadonly = matches[1]
	} else if len(matches) == 3 {
		config.SecondaryConfigReadonly = matches[2]
	}

	// now's the time
	stats.StatsTime = time.Now()

	// assume stats originate from the local host
	matches = rules["statsZeroTime"].FindStringSubmatch(text)
	if len(matches) == 3 {
		statsZeroTime := rules["statsZeroTime"].FindStringSubmatch(text)[2]
		stats.StatsZeroTime, err = time.ParseInLocation("Mon Jan 2 15:04:05 2006", statsZeroTime, stats.StatsTime.Location())
		if err != nil {
			return &Configuration{}, &Statistics{}, err
		}
	}

	matches = rules["cacheHitDirect"].FindStringSubmatch(text)
	if len(matches) == 2 {
		stats.CacheHitDirect, err = strconv.Atoi(matches[1])
		if err != nil {
			return &Configuration{}, &Statistics{}, err
		}
	}

	matches = rules["cacheHitPreprocessed"].FindStringSubmatch(text)
	if len(matches) == 2 {
		stats.CacheHitPreprocessed, err = strconv.Atoi(matches[1])
		if err != nil {
			return &Configuration{}, &Statistics{}, err
		}
	}

	matches = rules["cacheMiss"].FindStringSubmatch(text)
	if len(matches) == 2 {
		stats.CacheMiss, err = strconv.Atoi(matches[1])
		if err != nil {
			return &Configuration{}, &Statistics{}, err
		}
	}

	matches = rules["cacheHitRate"].FindStringSubmatch(text)
	if len(matches) == 3 {
		stats.CacheHitRate, err = strconv.ParseFloat(matches[1], 64)
		stats.CacheHitRatio = stats.CacheHitRate / 100
		if err != nil {
			return &Configuration{}, &Statistics{}, err
		}
	}

	matches = rules["calledForLink"].FindStringSubmatch(text)
	if len(matches) == 2 {
		stats.CalledForLink, err = strconv.Atoi(matches[1])
		if err != nil {
			return &Configuration{}, &Statistics{}, err
		}
	}

	matches = rules["calledForPreprocessing"].FindStringSubmatch(text)
	if len(matches) == 2 {
		stats.CalledForPreprocessing, err = strconv.Atoi(matches[1])
		if err != nil {
			return &Configuration{}, &Statistics{}, err
		}
	}

	matches = rules["unsupportedCodeDirective"].FindStringSubmatch(text)
	if len(matches) == 2 {
		stats.UnsupportedCodeDirective, err = strconv.Atoi(matches[1])
		if err != nil {
			return &Configuration{}, &Statistics{}, err
		}
	}

	matches = rules["noInputFile"].FindStringSubmatch(text)
	if len(matches) == 2 {
		stats.NoInputFile, err = strconv.Atoi(matches[1])
		if err != nil {
			return &Configuration{}, &Statistics{}, err
		}
	}

	matches = rules["cleanupsPerformed"].FindStringSubmatch(text)
	if len(matches) == 2 {
		stats.CleanupsPerformed, err = strconv.Atoi(matches[1])
		if err != nil {
			return &Configuration{}, &Statistics{}, err
		}
	}

	matches = rules["filesInCache"].FindStringSubmatch(text)
	if len(matches) == 2 {
		stats.FilesInCache, err = strconv.Atoi(matches[1])
		if err != nil {
			return &Configuration{}, &Statistics{}, err
		}
	}

	matches = rules["cacheSize"].FindStringSubmatch(text)
	if len(matches) == 2 {
		stats.CacheSize = matches[1]
		sanitizedCacheSize := strings.Replace(strings.ToUpper(stats.CacheSize), " ", "", -1)
		stats.CacheSizeBytes, err = units.ParseMetricBytes(sanitizedCacheSize)
		if err != nil {
			return &Configuration{}, &Statistics{}, err
		}
	}

	matches = rules["maxCacheSize"].FindStringSubmatch(text)
	if len(matches) == 2 {
		config.MaxCacheSize = matches[1]
		sanitizedMaxCacheSizeBytes := strings.Replace(strings.ToUpper(config.MaxCacheSize), " ", "", -1)
		config.MaxCacheSizeBytes, err = units.ParseMetricBytes(sanitizedMaxCacheSizeBytes)
		if err != nil {
			return &Configuration{}, &Statistics{}, err
		}
	}

	return config, stats, nil
}
