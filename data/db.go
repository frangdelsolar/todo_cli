package data

import (
	"os"
	"todo_cli/models"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *Database

// Database represents a connection to the SQLite database.
type Database struct {
	*gorm.DB
}

// ConnectDB establishes a connection to the SQLite database specified by the DATA_BASE_FILE environment variable.
// If the environment variable is not set, it defaults to "../data.db". It logs the connection details and migrates
// the schema for the models.Task, models.EffectivePeriod, and models.TaskCompletionLog. It returns an error if the
// database connection fails.
//
// Returns:
// - error: an error if the database connection fails.
func ConnectDB() error {

	// Get the DATA_BASE_FILE environment variable
	dataBaseFile := os.Getenv("DATA_BASE_FILE")
	if dataBaseFile == "" {
		dataBaseFile = "../data.db"
	}

	log.Info().Msgf("Connecting to SQLite DB: %s", dataBaseFile)

	// Connect to the SQLite database
	db, err := gorm.Open(sqlite.Open(dataBaseFile), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(
		&models.Task{},
		&models.EffectivePeriod{},
		&models.TaskCompletionLog{},
	)

	DB = &Database{db}

	log.Debug().Msg("Initialized SQLite DB")

	return nil
}
