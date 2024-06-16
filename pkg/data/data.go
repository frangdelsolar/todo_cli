package data

import (
	"github.com/frangdelsolar/todo_cli/pkg/logger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const PKG_NAME = "Data PKG"
const PKG_VERSION = "1.0.2"

var log *logger.Logger
var logLevel = "debug"

var DB Database

type Database struct {
	*gorm.DB
}

// InitDB initializes a new SQLite database connection and returns a pointer to the Database struct and an error.
//
// It takes a filepath string as a parameter, which is the path to the SQLite database file.
// If the filepath is empty, it defaults to "./data.db".
//
// The function returns a pointer to the Database struct and an error.
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


// GetDB returns the database connection for the current application.
//
// It checks if the database connection has been initialized and if not, it
// calls the InitDB function to establish a connection to the SQLite database.
// If the connection is successful, it returns the database connection.
// If the connection fails, it logs an error message and returns the error.
// If the database connection has already been initialized, it returns the
// existing connection.
//
// Returns:
// - *Database: the database connection.
// - error: an error if the database connection fails.
func GetDB() (*Database, error) {

	if DB == (Database{}) {
		db, err := InitDB("")
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to connect to database")
			return nil, err
		}
		log.Warn().Msg("Database was not initialized. Using default.")
		return db, nil
	}

	return &DB, nil
}