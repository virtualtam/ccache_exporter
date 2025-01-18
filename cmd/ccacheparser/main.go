// Copyright (c) VirtualTam
// SPDX-License-Identifier: MIT

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/virtualtam/ccache_exporter/pkg/ccache"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})

	stat, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	if stat.Mode()&os.ModeNamedPipe == 0 {
		// TODO add flags, read from stdin / file(s)
		// TODO add help
		// TODO switch between legacy / TSV parsers
		panic("No data piped to stdin")
	}

	var text string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text += scanner.Text() + "\n"
	}

	stats, err := ccache.ParseTSVStatistics(text)
	if err != nil {
		log.Fatal().Err(err).Msg("Parse")
	}

	statsJSON, err := json.Marshal(stats)
	if err != nil {
		log.Fatal().Err(err).Msg("Marshal")
	}

	fmt.Println(string(statsJSON))
}
