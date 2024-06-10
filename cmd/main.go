package main

import (
	"os"

	"github.com/frangdelsolar/todo_cli/data"
	"github.com/frangdelsolar/todo_cli/pkg/logger"

	"github.com/rs/zerolog"

	t "github.com/frangdelsolar/todo_cli/pkg/todo"
)

var APP_VERSION = "1.0.0"

func main() {

	zlog := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).With().Timestamp().Logger()

	log := logger.SetLogger(&zlog)
	log.Info().Msg("Running TODO APP v" + APP_VERSION)

	db, err := data.ConnectDB()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
	}

	todoConfig := t.TodoConfig{
		Logger: *log.Logger,
		DB: db,
	}
	cli := t.Todo(todoConfig)
	cli.Execute()


}
