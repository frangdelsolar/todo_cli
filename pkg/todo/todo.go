package todo

import (
	data "github.com/frangdelsolar/todo_cli/pkg/data"
	"github.com/frangdelsolar/todo_cli/pkg/logger"
	"github.com/frangdelsolar/todo_cli/pkg/todo/cli"
	"github.com/frangdelsolar/todo_cli/pkg/todo/models"
)


const PKG_NAME = "TODO PKG"
const PKG_VERSION = "1.0.3"

var log *logger.Logger
var logLevel = "debug"

func Todo() *cli.CLI {
	log = logger.NewLogger(logLevel, PKG_NAME, PKG_VERSION)
	log.Info().Msgf("Running %s v%s", PKG_NAME, PKG_VERSION)

	db := data.GetDB()

	// Migrate the schema
    db.AutoMigrate(
		&models.Task{},
		&models.TaskGoal{},
		&models.TaskCompletionLog{},
		&models.TaskFrequency{},
	)

	return cli.NewCLI(PKG_VERSION)
}
