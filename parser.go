// Copyright 2018 VirtualTam.
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

// Parse reads ccache statistics as formatted by the `ccache -s` command.
func (s *Statistics) Parse(text string) {
	matches := rules["cacheDirectory"].FindStringSubmatch(text)
	if len(matches) == 2 {
		s.CacheDirectory = matches[1]
	}

	matches = rules["primaryConfig"].FindStringSubmatch(text)
	if len(matches) == 2 {
		s.PrimaryConfig = matches[1]
	}

	matches = rules["secondaryConfigReadonly"].FindStringSubmatch(text)
	if len(matches) == 2 {
		s.SecondaryConfigReadonly = matches[1]
	} else if len(matches) == 3 {
		s.SecondaryConfigReadonly = matches[2]
	}

	// now's the time
	s.StatsTime = time.Now()

	// assume stats originate from the local host
	matches = rules["statsZeroTime"].FindStringSubmatch(text)
	if len(matches) == 3 {
		statsZeroTime := rules["statsZeroTime"].FindStringSubmatch(text)[2]
		s.StatsZeroTime, _ = time.ParseInLocation("Mon Jan 2 15:04:05 2006", statsZeroTime, s.StatsTime.Location())
	}

	matches = rules["cacheHitDirect"].FindStringSubmatch(text)
	if len(matches) == 2 {
		s.CacheHitDirect, _ = strconv.Atoi(matches[1])
	}

	matches = rules["cacheHitPreprocessed"].FindStringSubmatch(text)
	if len(matches) == 2 {
		s.CacheHitPreprocessed, _ = strconv.Atoi(matches[1])
	}

	matches = rules["cacheMiss"].FindStringSubmatch(text)
	if len(matches) == 2 {
		s.CacheMiss, _ = strconv.Atoi(matches[1])
	}

	matches = rules["cacheHitRate"].FindStringSubmatch(text)
	if len(matches) == 3 {
		s.CacheHitRate, _ = strconv.ParseFloat(matches[1], 64)
		s.CacheHitRatio = s.CacheHitRate / 100
	}

	matches = rules["calledForLink"].FindStringSubmatch(text)
	if len(matches) == 2 {
		s.CalledForLink, _ = strconv.Atoi(matches[1])
	}

	matches = rules["calledForPreprocessing"].FindStringSubmatch(text)
	if len(matches) == 2 {
		s.CalledForPreprocessing, _ = strconv.Atoi(matches[1])
	}

	matches = rules["unsupportedCodeDirective"].FindStringSubmatch(text)
	if len(matches) == 2 {
		s.UnsupportedCodeDirective, _ = strconv.Atoi(matches[1])
	}

	matches = rules["noInputFile"].FindStringSubmatch(text)
	if len(matches) == 2 {
		s.NoInputFile, _ = strconv.Atoi(matches[1])
	}

	matches = rules["cleanupsPerformed"].FindStringSubmatch(text)
	if len(matches) == 2 {
		s.CleanupsPerformed, _ = strconv.Atoi(matches[1])
	}

	matches = rules["filesInCache"].FindStringSubmatch(text)
	if len(matches) == 2 {
		s.FilesInCache, _ = strconv.Atoi(matches[1])
	}

	matches = rules["cacheSize"].FindStringSubmatch(text)
	if len(matches) == 2 {
		s.CacheSize = matches[1]
		s.CacheSizeBytes, _ = units.ParseMetricBytes(strings.Replace(s.CacheSize, " ", "", -1))
	}

	matches = rules["maxCacheSize"].FindStringSubmatch(text)
	if len(matches) == 2 {
		s.MaxCacheSize = matches[1]
		s.MaxCacheSizeBytes, _ = units.ParseMetricBytes(strings.Replace(s.MaxCacheSize, " ", "", -1))
	}
}
