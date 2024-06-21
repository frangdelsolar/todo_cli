package auth

import (
	"github.com/frangdelsolar/todo_cli/pkg/data"
	"github.com/frangdelsolar/todo_cli/pkg/logger"
)


var PKG_NAME = "Auth PKG"
var PKG_VERSION = "1.0.5"

var log *logger.Logger
var db *data.Database 

// InitAuth initializes the authentication package by setting up the logger and
// connecting to the database. It also performs the necessary database migrations.
//
// No parameters.
// No return values.
func InitAuth() {
	var err error
	
	log = logger.NewLogger(logger.LoggerConfig{
        PackageName: PKG_NAME,
        PackageVersion: PKG_VERSION,
    })

	db, err = data.GetDB()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
		return 
	}

	db.AutoMigrate(
		&User{}, 
	)

	log.Debug().Msg("Applied Auth migrations to database")
	
}
