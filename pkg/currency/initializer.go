package currency

import (
	"github.com/frangdelsolar/todo_cli/pkg/data"
	"github.com/frangdelsolar/todo_cli/pkg/logger"
)


var PKG_NAME = "Currency PKG"
var PKG_VERSION = "1.0.3"

var log *logger.Logger

var db *data.Database

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
	log = logger.NewLogger(PKG_NAME, PKG_VERSION)

	db, err = data.GetDB()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
		return
	}

	db.AutoMigrate(
		&Account{}, 
		&Currency{}, 
		&Transaction{},
	)

	log.Debug().Msg("Applied Currency pkg migrations to database")
	
}