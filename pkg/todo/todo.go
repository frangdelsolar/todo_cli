package todo

import (
	"fmt"
	"os"

	"github.com/frangdelsolar/todo_cli/pkg/logger"
	"github.com/frangdelsolar/todo_cli/pkg/todo/cli"
	db "github.com/frangdelsolar/todo_cli/pkg/todo/data"
	"github.com/frangdelsolar/todo_cli/pkg/todo/models"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)


const PKG_VERSION = "1.0.0"

var log *logger.Logger
var logLevel = zerolog.DebugLevel

type TodoConfig struct {
	DB     *gorm.DB
}

func Todo(config TodoConfig) *cli.CLI {
	configLogger()

	log.Info().Msg("Running TODO PKG v" + PKG_VERSION)

	// Migrate the schema
    config.DB.AutoMigrate(
		&models.Task{},
		&models.TaskGoal{},
		&models.TaskCompletionLog{},
		&models.TaskFrequency{},
	)

	db.InitDB(config.DB)
	
	return cli.NewCLI(PKG_VERSION)
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
			Str("app", fmt.Sprintf("TODO PKG v%s", PKG_VERSION)).
			Logger()

	log = logger.NewLogger(&zl)
}