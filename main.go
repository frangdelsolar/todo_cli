package main

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"todo_cli/cli"
)

var APP_VERSION = "0.0.1"

func main(){
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Info().Msg("Running TODO App v" + APP_VERSION)

	c := cli.NewCLI()
	c.Run()


}
