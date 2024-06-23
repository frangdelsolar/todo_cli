package main_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/frangdelsolar/todo_cli/pkg/auth"
	"github.com/frangdelsolar/todo_cli/pkg/config"
	"github.com/frangdelsolar/todo_cli/pkg/data"
	"github.com/frangdelsolar/todo_cli/pkg/logger"
	"github.com/stretchr/testify/assert"
	// at "github.com/frangdelsolar/todo_cli/pkg/test/auth_test"
	// cli "github.com/frangdelsolar/todo_cli/pkg/test/cli_test"
	// co "github.com/frangdelsolar/todo_cli/pkg/test/contractor_test"
	// ct "github.com/frangdelsolar/todo_cli/pkg/test/currency_test"
)

var PKG_NAME = "Test PKG"
var PKG_VERSION = "1.0.3"

var log *logger.Logger
var cfg *config.Config
var db *data.Database


func init(){
    os.Setenv("APP_ENV", "test")
    var err error
    cfg, err = config.Load()
    if err != nil {
        fmt.Errorf("Failed to load config: %v", err)
    }
    
    log = logger.NewLogger(logger.LoggerConfig{
        PackageName: PKG_NAME,
        PackageVersion: PKG_VERSION,
    })

    log.Info().Msgf("Running %s v%s", PKG_NAME, PKG_VERSION)
    log.Debug().Interface("Config", cfg).Msg("Loaded Config")

	db, err := data.LoadDB()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
	}
	log.Debug().Msgf("Loaded Database: %s", db.Name())
    
    auth.InitAuth()
}

func TestCreateUser(t *testing.T){
    name := "pepe"
    email := "pepe@admin.com"
    password := "test123"

    u, err := auth.CreateUser(name, email, password)
    if err != nil {
        log.Warn().Msg("Failed to create user")
    }

    assert.Equal(t, u.Name, name, "Expected name to be %s, but got %s", name, u.Name)
    assert.Equal(t, u.Email, email, "Expected email to be %s, but got %s", email, u.Email)

}
