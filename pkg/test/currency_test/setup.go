package currency_test

import (
	"os"

	"github.com/frangdelsolar/todo_cli/pkg/auth"
	"github.com/frangdelsolar/todo_cli/pkg/config"
	c "github.com/frangdelsolar/todo_cli/pkg/currency"
	"github.com/frangdelsolar/todo_cli/pkg/data"
	"github.com/frangdelsolar/todo_cli/pkg/logger"
)

var PKG_NAME = "Currency Test PKG"
var PKG_VERSION = "0.0.1"

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

    log.Info().Interface("Config", cfg).Msg("Loaded Config")

	db, err := data.LoadDB()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
		panic(err)
	}
	log.Debug().Msgf("Loaded Database: %s", db.Name())

    auth.InitAuth()
}

func RunCurrencyTests(){
    log.Info().Msg("Running Currency Tests")

    TestInitCurrency()
    TestAddCurrencySameCode()
    TestAddCurrencyDifferentCode()
    TestSubCurrencySameCode()
    TestSubCurrencyDifferentCode()
    TestCreateAccount()
    // TestCreateTransaction()
}

func TestInitCurrency(){
    log.Info().Msg("Testing InitCurrency()")
    c.InitCurrency()
    log.Debug().Msg("Applied Currency pkg migrations to database")
}
