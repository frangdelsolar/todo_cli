package main

import (
	"os"
	"todo_app/data"
	"todo_app/models"

	"github.com/rs/zerolog"

	t "github.com/frangdelsolar/todo_cli/pkg/todo"
)

var APP_VERSION = "1.0.0"




func main() {

	log := models.Logger{zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).With().Timestamp().Logger()}

	db, err := data.ConnectDB(log)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
	}

	todoConfig := t.TodoConfig{
		Logger: log.Logger,
		DB: db,
	}
	cli := t.Todo(todoConfig)
	cli.Execute()

	log.Info().Msg("Running TODO APP v" + APP_VERSION)

}
