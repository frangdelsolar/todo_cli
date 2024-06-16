package auth

import (
	m "github.com/frangdelsolar/todo_cli/pkg/auth/models"
	"github.com/frangdelsolar/todo_cli/pkg/data"
	"github.com/frangdelsolar/todo_cli/pkg/logger"
)


var PKG_NAME = "Auth PKG"
var PKG_VERSION = "1.0.3"

var log *logger.Logger
var logLevel = "debug"

var db *data.Database = &data.Database{}


// InitAuth initializes the authentication package by setting up the logger and
// connecting to the database. It also performs the necessary database migrations.
//
// No parameters.
// No return values.
func InitAuth() {
	var err error
	log = logger.NewLogger(logLevel, PKG_NAME, PKG_VERSION)
	log.Info().Msgf("Running %s v%s", PKG_NAME, PKG_VERSION)

	db, err = data.GetDB()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
		return 
	}

	db.AutoMigrate(
		&m.User{}, 
	)

	log.Debug().Interface("Database", db).Msg("Initialized Database")
	
}