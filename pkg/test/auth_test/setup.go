package auth_test

import (
	"os"

	"github.com/frangdelsolar/todo_cli/pkg/auth"
	"github.com/frangdelsolar/todo_cli/pkg/config"
	"github.com/frangdelsolar/todo_cli/pkg/data"
	"github.com/frangdelsolar/todo_cli/pkg/logger"
)

var PKG_NAME = "Auth Test PKG"
var PKG_VERSION = "0.0.1"

var cfg *config.Config
var log *logger.Logger
var db *data.Database


func init(){
    // Set Env
    os.Setenv("APP_ENV", "test")

    cfg, err := config.Load()
    if err != nil {
        panic(err)
    }
    
    log = logger.NewLogger(PKG_NAME, PKG_VERSION)
    log.Info().Msgf("Running %s v%s", PKG_NAME, PKG_VERSION)

    log.Info().Interface("Config", cfg).Msg("Loaded Config")

	db, err := data.LoadDB()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
		panic(err)
	}
	log.Debug().Msgf("Loaded Database: %s", db.Name())
}

func RunAuthTests(){
    log.Info().Msg("Running Auth Tests")

    TestInitAuth()
    TestCreateUser()
}


func TestInitAuth(){
    log.Info().Msg("Testing InitAuth()")
    auth.InitAuth()
    log.Debug().Msg("Applied Auth migrations to database")
}
