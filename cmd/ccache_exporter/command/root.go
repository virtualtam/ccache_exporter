// Copyright (c) VirtualTam
// SPDX-License-Identifier: MIT

package command

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/virtualtam/venom"

	"github.com/virtualtam/ccache_exporter/cmd/ccache_exporter/config"
	"github.com/virtualtam/ccache_exporter/internal/version"
	"github.com/virtualtam/ccache_exporter/pkg/ccache"
)

const (
	rootCmdName string = "ccache_exporter"
)

var (
	logFormat            string
	defaultLogLevelValue string = zerolog.LevelInfoValue
	logLevelValue        string

	versionDetails *version.Details

	ccacheBinaryPath string
	ccacheWrapper    *ccache.Wrapper
)

// NewRootCommand initializes the exporter's CLI entrypoint and global command flags.
func NewRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   rootCmdName,
		Short: "Prometheus exporter for ccache metrics",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// Configuration file lookup paths
			home, err := os.UserHomeDir()
			if err != nil {
				return err
			}
			homeConfigPath := filepath.Join(home, ".config")

			configPaths := []string{config.DefaultConfigPath, homeConfigPath, "."}

			// Inject global configuration as a pre-run hook
			//
			// This is required to let Viper load environment variables and
			// configuration entries before invoking nested commands.
			v := viper.New()
			if err := venom.InjectTo(v, cmd, config.EnvPrefix, configPaths, config.ConfigName, false); err != nil {
				return err
			}

			// Setup ccache wrapper
			ccacheCommand, err := ccache.NewLocalCommand(ccacheBinaryPath)
			if err != nil {
				log.Fatal().Err(err).Msg("ccache: failed to instantiate command wrapper")
			}

			ccacheWrapper = ccache.NewWrapper(ccacheCommand)

			// Retrieve exporter and ccache versions
			versionDetails = version.NewDetails(ccacheWrapper.Version())

			if cmd.Name() == versionCmdName {
				// Do not setup the service stack for these commands
				return nil
			}

			// Global logger configuration
			if err := config.SetupGlobalLogger(logFormat, logLevelValue); err != nil {
				return err
			}

			if configFileUsed := v.ConfigFileUsed(); configFileUsed != "" {
				log.Info().Str("config_file", v.ConfigFileUsed()).Msg("configuration: using file")
			} else {
				log.Info().Strs("config_paths", configPaths).Msg("configuration: no file found")
			}

			log.Info().
				Str("ccache_binary", ccacheBinaryPath).
				Str("ccache_version", ccacheWrapper.Version()).
				Msg("ccache: command wrapper created")

			return nil
		},
	}

	cmd.PersistentFlags().StringVar(
		&logFormat,
		"log-format",
		config.LogFormatConsole,
		fmt.Sprintf("Log format (%s, %s)", config.LogFormatJSON, config.LogFormatConsole),
	)
	cmd.PersistentFlags().StringVar(
		&logLevelValue,
		"log-level",
		defaultLogLevelValue,
		fmt.Sprintf(
			"Log level (%s)",
			strings.Join(config.LogLevelValues, ", "),
		),
	)

	cmd.PersistentFlags().StringVar(
		&ccacheBinaryPath,
		"ccache-binary-path",
		ccache.DefaultBinaryPath,
		"Path to the ccache binary",
	)

	return cmd
}
