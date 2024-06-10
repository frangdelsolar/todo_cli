package data

import (
	"os"

	"github.com/frangdelsolar/todo_cli/pkg/todo/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// ConnectDB establishes a connection to the SQLite database specified by the DATA_BASE_FILE environment variable.
// If the environment variable is not set, it defaults to "../data.db". It logs the connection details and migrates
// the schema for the models.Task, models.TaskGoal, and models.TaskCompletionLog. It returns an error if the
// database connection fails.
//
// Returns:
// - error: an error if the database connection fails.
func ConnectDB(log *models.Logger) (*gorm.DB, error) {

	// Get the DATA_BASE_FILE environment variable
	dataBaseFile := os.Getenv("DATA_BASE_FILE")
	if dataBaseFile == "" {
		dataBaseFile = "../data.db"
	}

	log.Info().Msgf("Connecting to SQLite DB: %s", dataBaseFile)

	// Connect to the SQLite database
	db, err := gorm.Open(sqlite.Open(dataBaseFile), &gorm.Config{})

	return db, err
}
