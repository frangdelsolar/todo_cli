package main

import (
	"os"
	"todo_cli/data"
	"todo_cli/internal/cli"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var APP_VERSION = "0.0.1"

func main() {

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Info().Msg("Running TODO App v" + APP_VERSION)

	err := data.ConnectDB()
	if err != nil {
		log.Err(err).Msg("Error connecting to database")
		return
	}
	cli.Execute()
}
