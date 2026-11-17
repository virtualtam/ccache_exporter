// Copyright VirtualTam 2018, 2026
// SPDX-License-Identifier: MIT

package config

const (
	// EnvPrefix is the prefix for all environment variables usable to configure the exporter.
	EnvPrefix string = "CCACHE_EXPORTER"

	// DefaultConfigPath is the default directory where the exporter will look for a configuration file.
	DefaultConfigPath string = "/etc"

	// ConfigName is the name (without extension) of the configuration file.
	ConfigName string = "ccache_exporter"
)
