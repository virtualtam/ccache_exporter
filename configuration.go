package ccache

import (
	"bufio"
	"fmt"
	"path/filepath"
	"strings"
	"unicode"

	"github.com/alecthomas/units"
)

// Configuration represents information about ccache configuration.
type Configuration struct {
	CacheDirectory          string            `json:"cache_directory"`
	PrimaryConfig           string            `json:"primary_config"`
	SecondaryConfigReadonly string            `json:"secondary_config_readonly"`
	MaxCacheSize            string            `json:"max_cache_size"`
	MaxCacheSizeBytes       units.MetricBytes `json:"max_cache_size_bytes"`
}

func splitConfigurationField(c rune) bool {
	return unicode.IsSpace(c) || c == '=' || c == '(' || c == ')'
}

// ParseConfiguration parses ccache configuration returned by the
// `--show-config` (ccache >= 3.7) or the `--print-config` (ccache < 3.7)
// command.
//
// See https://ccache.dev/releasenotes.html#_ccache_3_7
func ParseConfiguration(text string) (*Configuration, error) {
	reader := strings.NewReader(text)
	scanner := bufio.NewScanner(reader)

	configuration := &Configuration{}
	var err error

	for scanner.Scan() {
		// split each configuration line into 3 fields:
		//
		// (<configuration source>) <key> = <value>
		fields := strings.FieldsFunc(scanner.Text(), splitConfigurationField)
		if len(fields) != 3 {
			continue
		}

		switch fields[1] {
		case "cache_dir":
			configuration.CacheDirectory = fields[2]
			configuration.PrimaryConfig = filepath.Join(fields[2], "ccache.conf")
		case "max_size":
			sanitizedMaxCacheSize := fmt.Sprintf(
				"%sB",
				strings.Replace(strings.ToUpper(fields[2]), " ", "", -1),
			)
			configuration.MaxCacheSize = sanitizedMaxCacheSize
			configuration.MaxCacheSizeBytes, err = units.ParseMetricBytes(sanitizedMaxCacheSize)
			if err != nil {
				return &Configuration{}, err
			}
		}
	}

	return configuration, nil
}
