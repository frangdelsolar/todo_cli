package main

import (
	"os"

	"github.com/frangdelsolar/todo_cli/data"

	"github.com/rs/zerolog"

	t "github.com/frangdelsolar/todo_cli/pkg/todo"
	"github.com/frangdelsolar/todo_cli/pkg/todo/models"
)

var APP_VERSION = "1.0.0"

func main() {

	zlog := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).With().Timestamp().Logger()

	log := models.NewLogger(&zlog)

	db, err := data.ConnectDB(log)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
	}

	todoConfig := t.TodoConfig{
		Logger: *log.Logger,
		DB: db,
	}
	cli := t.Todo(todoConfig)
	cli.Execute()

	log.Info().Msg("Running TODO APP v" + APP_VERSION)

}
