package main

import (
	"os"
	cmd "todo_cli/internal/cli"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var APP_VERSION = "0.0.1"

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Info().Msg("Running TODO App v" + APP_VERSION)

	cmd.Execute()
}
