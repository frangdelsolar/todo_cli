package contractor_test

import (
	"os"

	"github.com/frangdelsolar/todo_cli/pkg/auth"
	"github.com/frangdelsolar/todo_cli/pkg/config"
	c "github.com/frangdelsolar/todo_cli/pkg/contractor"
	"github.com/frangdelsolar/todo_cli/pkg/data"
	"github.com/frangdelsolar/todo_cli/pkg/logger"
)

var PKG_NAME = "Contractor Test PKG"
var PKG_VERSION = "1.0.0"

var cfg *config.Config
var log *logger.Logger
var db *data.Database


func init(){
    // Test Bed
    os.Setenv("APP_ENV", "test")

    cfg, err := config.Load()
    if err != nil {
        panic(err)
    }

    log = logger.NewLogger(PKG_NAME, PKG_VERSION)
    log.Info().Msgf("Running %s v%s", PKG_NAME, PKG_VERSION)
    log.Trace().Interface("Config", cfg).Msg("Loaded Config")

	db, err := data.LoadDB()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
		panic(err)
	}
	log.Debug().Msgf("Loaded Database: %s", db.Name())

    auth.InitAuth()
    c.InitContractor()

}

func RunContractorTests(){
    log.Info().Msg("Running Contractor Tests")

    TestCreateContractor()
    TestUpdateContractorName()
}
