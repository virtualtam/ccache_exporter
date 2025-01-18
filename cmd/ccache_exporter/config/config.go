// Copyright (c) VirtualTam
// SPDX-License-Identifier: MIT

package config

const (
	// Prefix for all environment variables
	EnvPrefix string = "CCACHE_EXPORTER"

	// Directory where the exporter will look for a configuration file
	DefaultConfigPath string = "/etc"

	// Name (without extension) of the configuration file
	ConfigName string = "ccache_exporter"
)
