package currency

import (
	m "github.com/frangdelsolar/todo_cli/pkg/currency/models"
	"github.com/frangdelsolar/todo_cli/pkg/data"
	"github.com/frangdelsolar/todo_cli/pkg/logger"
)


var PKG_NAME = "Currency PKG"
var PKG_VERSION = "1.0.0"

var log *logger.Logger
var logLevel = "debug"

var db *data.Database = &data.Database{}

func InitCurrency() {
	var err error
	log = logger.NewLogger(logLevel, PKG_NAME, PKG_VERSION)
	log.Info().Msgf("Running %s v%s", PKG_NAME, PKG_VERSION)

	
	/*
		Comment this to initialize a new database
	*/
	// db = data.GetDB()
	// log.Debug().Interface("Database", db).Msg("Initialized Database")

	/*
		Uncomment this to initialize a new database.
		Comment the previous line
	*/

	db, err = data.InitDB("./data.db")
	if err != nil {
		log.Err(err).Msg("Error initializing database")
		return
	}

	db.AutoMigrate(
		&m.Currency{}, 
		&m.Account{}, 
		&m.Transaction{},
	)

	log.Debug().Interface("Database", db).Msg("Initialized Database")
	
}