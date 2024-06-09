package todo

import (
	"github.com/frangdelsolar/todo_cli/pkg/todo/cli"
	db "github.com/frangdelsolar/todo_cli/pkg/todo/data"
	"github.com/frangdelsolar/todo_cli/pkg/todo/models"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)


const APP_VERSION = "1.0.0"

var log zerolog.Logger

type TodoConfig struct {
	Logger zerolog.Logger
	DB     *gorm.DB
}

func Todo(config TodoConfig) *cli.CLI {
	log = config.Logger
	log.Info().Msg("Running TODO PKG v" + APP_VERSION)

	// Migrate the schema
    config.DB.AutoMigrate(
		&models.Task{},
		&models.TaskGoal{},
		&models.TaskCompletionLog{},
		&models.TaskFrequency{},
	)

	db.InitDB(config.DB, &log)
	
	return cli.NewCLI(&log)
}