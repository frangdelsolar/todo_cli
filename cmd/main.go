package main

import (
	"fmt"
	"os"

	"github.com/frangdelsolar/todo_cli/data"
	"github.com/frangdelsolar/todo_cli/pkg/logger"
	t "github.com/frangdelsolar/todo_cli/pkg/todo"

	"github.com/rs/zerolog"
)

var APP_VERSION = "1.0.0"

var log *logger.Logger
var logLevel = zerolog.DebugLevel

func main() {

	configLogger()

	db, err := data.ConnectDB()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
	}

	todoConfig := t.TodoConfig{
		DB: db,
	}
	cli := t.Todo(todoConfig)
	cli.Execute()
}

// configLogs initializes the logger and sets the global log level. It also
// creates a new logger with a console writer and adds additional fields like
// "app" and "version". Finally, it logs an info message indicating that the
// TODO APP is running with the specified version.
//
// No parameters.
// No return value.
func configLogger() {
	zerolog.SetGlobalLevel(logLevel)
	zl := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).
		    With().
			Timestamp().
			Str("app", fmt.Sprintf("TODO APP v%s", APP_VERSION)).
			Logger()

	log = logger.NewLogger(&zl)
	log.Info().Msg("Running TODO APP v" + APP_VERSION)
}