package cli_test

import (
	"fmt"

	"github.com/frangdelsolar/todo_cli/pkg/auth"
	"github.com/frangdelsolar/todo_cli/pkg/config"
	"github.com/frangdelsolar/todo_cli/pkg/data"
	"github.com/frangdelsolar/todo_cli/pkg/logger"
)

var PKG_NAME = "CLI Test PKG"
var PKG_VERSION = "1.0.0"

var cfg *config.Config
var log *logger.Logger
var db *data.Database


func init(){

    cfg, err := config.Load()
    if err != nil {
        fmt.Errorf("Failed to load config: %v", err)
    }
    
    log = logger.NewLogger(logger.LoggerConfig{
        PackageName: PKG_NAME,
        PackageVersion: PKG_VERSION,
    })
    log.Info().Msgf("Running %s v%s", PKG_NAME, PKG_VERSION)
    log.Trace().Interface("Config", cfg).Msg("Loaded Config")

	db, err := data.LoadDB()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
		panic(err)
	}
	log.Debug().Msgf("Loaded Database: %s", db.Name())
    
    auth.InitAuth()
}

func RunCliTests(){
    log.Info().Msg("Running Auth Tests")
    
    TestRegisterUser()
    TestLoginUser()

}
