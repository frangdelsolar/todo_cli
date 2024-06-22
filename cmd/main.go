package main

import (
	"github.com/frangdelsolar/todo_cli/cmd/cli"
	"github.com/frangdelsolar/todo_cli/pkg/auth"
	"github.com/frangdelsolar/todo_cli/pkg/config"
	c "github.com/frangdelsolar/todo_cli/pkg/contractor"
	"github.com/frangdelsolar/todo_cli/pkg/data"
	"github.com/frangdelsolar/todo_cli/pkg/logger"
)

var PKG_NAME = "TODO CLI APP"
var PKG_VERSION = "1.0.6"

var cfg *config.Config
var log *logger.Logger
var db *data.Database


func main(){

    cfg, err := config.Load()
    if err != nil {
        panic(err)
    }
    
    log = logger.NewLogger(logger.LoggerConfig{
        PackageName: PKG_NAME,
        PackageVersion: PKG_VERSION,
    })

    log.Info().Msgf("Running %s v%s", PKG_NAME, PKG_VERSION)
    log.Info().Interface("Config", cfg).Msg("Loaded Config")

	db, err := data.LoadDB()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
		panic(err)
	}
	log.Debug().Msgf("Loaded Database: %s", db.Name())
    
    auth.InitAuth()
    c.InitContractor()

    cli.Execute()
}
