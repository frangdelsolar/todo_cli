package main

import (
	"github.com/frangdelsolar/todo_cli/pkg/data"
	"github.com/frangdelsolar/todo_cli/pkg/logger"
	t "github.com/frangdelsolar/todo_cli/pkg/todo"
)

var APP_NAME= "TODO APP"
var APP_VERSION = "1.0.3"

var log *logger.Logger
var logLevel = "debug"

func main() {

	log = logger.NewLogger(logLevel, APP_NAME, APP_VERSION)
	log.Info().Msg("Running TODO APP v" + APP_VERSION)
	
	db, err := data.InitDB("data.db")
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
	}

	log.Info().Interface("db", db).Msg("Database connected")

	cli := t.Todo()
	cli.Execute()
	log.Debug().Interface("cli", cli).Msg("CLI initialized")
}
