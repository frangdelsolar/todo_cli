package data

import (
	"github.com/frangdelsolar/todo_cli/pkg/logger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const PKG_NAME = "Data PKG"
const PKG_VERSION = "1.0.1"

var log *logger.Logger
var logLevel = "debug"

var DB Database

type Database struct {
	*gorm.DB
}

func InitDB(filepath string) (*Database, error) {

	log = logger.NewLogger(logLevel, PKG_NAME, PKG_VERSION)
	log.Info().Msgf("Running %s v%s", PKG_NAME, PKG_VERSION)

	if filepath == "" {
		filepath = "./data.db"
	}

	log.Info().Msgf("Connecting to SQLite DB: %s", filepath)
	
	db, err := gorm.Open(sqlite.Open(filepath), &gorm.Config{})
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
	}

	DB = Database{db}
	
	return &DB, err
}

func GetDB() (*Database, error) {

	if DB == (Database{}) {
		db, err := InitDB("")
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to connect to database")
		}
		log.Warn().Msg("Database was not initialized. Using default.")
		return db, nil
	}

	return &DB, nil
}