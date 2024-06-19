package data

import (
	auth "github.com/frangdelsolar/todo_cli/pkg/auth/models"
	"github.com/frangdelsolar/todo_cli/pkg/config"
	"github.com/frangdelsolar/todo_cli/pkg/logger"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const PKG_NAME = "Data PKG"
const PKG_VERSION = "1.1.3"

var DB *Database

type Database struct {
	*gorm.DB
}

type SystemData struct {
	gorm.Model
	CreatedBy   *auth.User
	CreatedByID uint
	UpdatedBy   *auth.User
	UpdatedByID uint
}

// LoadDB initializes a new SQLite database connection and returns a pointer to the Database struct and an error.
//
// It takes a filepath string as a parameter, which is the path to the SQLite database file.
// If the filepath is empty, it defaults to "./data.db".
//
// The function returns a pointer to the Database struct and an error.
func LoadDB() (*Database, error) {

	log := logger.NewLogger(PKG_NAME, PKG_VERSION)

	cfg := config.GetConfig()

	filepath := cfg.DBPath

	log.Info().Msgf("Connecting to database: %s", filepath)

	db, err := gorm.Open(sqlite.Open(filepath), &gorm.Config{})
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
	}

	DB = &Database{db}

	return DB, err
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

	if DB == nil {
		log.Warn().Msg("Database was not initialized. Will initialize now...")
		db, err := LoadDB()
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to connect to database")
			return nil, err
		}
		return db, nil
	}

	return DB, nil
}
