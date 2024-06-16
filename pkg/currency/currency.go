package currency

import (
	m "github.com/frangdelsolar/todo_cli/pkg/currency/models"
	"github.com/frangdelsolar/todo_cli/pkg/data"
	"github.com/frangdelsolar/todo_cli/pkg/logger"
)


var PKG_NAME = "Currency PKG"
var PKG_VERSION = "1.0.1"

var log *logger.Logger
var logLevel = "debug"

var db *data.Database = &data.Database{}


// InitCurrency initializes the currency package.
//
// It initializes the logger with the package name and version.
// It attempts to connect to the database using the data package's GetDB() function.
// If an error occurs during the connection, it logs the error and exits.
// It then auto-migrates the database tables for the Currency, Account, and Transaction models.
// Finally, it logs the initialized database.
//
// No parameters.
// No return values.
func InitCurrency() {
	var err error
	log = logger.NewLogger(logLevel, PKG_NAME, PKG_VERSION)
	log.Info().Msgf("Running %s v%s", PKG_NAME, PKG_VERSION)

	db, err = data.GetDB()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
		return
	}
	log.Debug().Interface("Database", db).Msg("Initialized Database")

	db.AutoMigrate(
		&m.Currency{}, 
		&m.Account{}, 
		&m.Transaction{},
	)

	log.Debug().Interface("Database", db).Msg("Initialized Database")
	
}