// Copyright (c) VirtualTam
// SPDX-License-Identifier: MIT

package main

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/virtualtam/ccache_exporter/cmd/ccache_exporter/command"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: time.RFC3339,
	})

	rootCommand := command.NewRootCommand()
	rootCommand.AddCommand(
		command.NewRunCommand(),
		command.NewVersionCommand(),
	)

	cobra.CheckErr(rootCommand.Execute())
}
