package auth

import (
	"github.com/frangdelsolar/todo_cli/pkg/config"
	"github.com/frangdelsolar/todo_cli/pkg/data"
	"github.com/frangdelsolar/todo_cli/pkg/logger"
)


var PKG_NAME = "Auth PKG"
var PKG_VERSION = "1.0.8"

var log *logger.Logger
var db *data.Database 
var fa *FirebaseAdmin
var cfg *config.Config

// InitAuth initializes the authentication package by setting up the logger and
// connecting to the database. It also performs the necessary database migrations.
//
// No parameters.
// No return values.
func InitAuth() {
	var err error

    cfg, err = config.Load()
    if err != nil {
        log.Fatal().Err(err).Msg("Failed to load config")
        return
    }
	
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

    // fa, err = NewFirebaseAdmin(&AuthConfig{CredentialsFilePath: cfg.FirebaseAdminSdk})
    // if err != nil {
    //     log.Err(err).Msg("Failed to initialize Firebase Admin")
    // }

    // log.Debug().Interface("Admin", fa.App).Msg("Initialized Firebase Admin")
    log.Info(). Msg("Initialized Auth")
	
}
