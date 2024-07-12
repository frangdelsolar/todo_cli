package todo

import (
	data "github.com/frangdelsolar/todo_cli/pkg/data"
	"github.com/frangdelsolar/todo_cli/pkg/logger"
	"github.com/frangdelsolar/todo_cli/pkg/todo/models"
)

const PKG_NAME = "TODO PKG"
const PKG_VERSION = "1.0.6"

var log *logger.Logger
var logLevel = "debug"

// InitTodo initializes the todo package.
//
// No parameters.
// No return values.
func InitTodo() {
	log = logger.NewLogger(logLevel, PKG_NAME, PKG_VERSION)
	log.Info().Msgf("Running %s v%s", PKG_NAME, PKG_VERSION)

	db, err := data.GetDB()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
		return
	}

	// Migrate the schema
	db.AutoMigrate(
		&models.Task{},
		&models.TaskGoal{},
		&models.TaskCompletionLog{},
		&models.TaskFrequency{},
	)

}
