package main

import (
	"github.com/frangdelsolar/todo_cli/pkg/config"
	"github.com/frangdelsolar/todo_cli/pkg/logger"
)

var APP_NAME = "TODO APP"
var APP_VERSION = "1.5.0"

func main() {

	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	log := logger.NewLogger(APP_NAME, APP_VERSION)
	log.Info().Msgf("Running %s v%s", APP_NAME, APP_VERSION)
	log.Debug().Interface("Config", cfg).Msg("Loaded Config")
}
