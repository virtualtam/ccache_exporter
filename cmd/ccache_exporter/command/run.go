// Copyright (c) VirtualTam
// SPDX-License-Identifier: MIT

package command

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/virtualtam/ccache_exporter/cmd/ccache_exporter/metrics"
)

const (
	defaultListenAddr = "0.0.0.0:9508"
)

var (
	listenAddr string
)

// NewRunCommand initializes a CLI command to start the exporter's HTTP server.
func NewRunCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "run",
		Short: "Start the exporter's HTTP server",
		RunE: func(cmd *cobra.Command, args []string) error {
			httpServer := metrics.NewServer(ccacheCommand, listenAddr)

			log.Info().Str("addr", listenAddr).Msg("starting HTTP server")
			return httpServer.ListenAndServe()
		},
	}

	cmd.Flags().StringVar(
		&listenAddr,
		"listen-addr",
		defaultListenAddr,
		"Listen to this address (host:port)",
	)

	return cmd
}
