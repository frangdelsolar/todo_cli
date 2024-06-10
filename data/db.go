package data

import (
	"fmt"
	"os"

	"github.com/frangdelsolar/todo_cli/logger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var log = logger.GetLogger()

func init(){
	fmt.Print("Getting Logger")
}

// ConnectDB establishes a connection to the SQLite database specified by the DATA_BASE_FILE environment variable.
// If the environment variable is not set, it defaults to "../data.db". It logs the connection details and migrates
// the schema for the models.Task, models.TaskGoal, and models.TaskCompletionLog. It returns an error if the
// database connection fails.
//
// Returns:
// - error: an error if the database connection fails.
func ConnectDB() (*gorm.DB, error) {

	// Get the DATA_BASE_FILE environment variable
	dataBaseFile := os.Getenv("DATA_BASE_FILE")
	if dataBaseFile == "" {
		dataBaseFile = "../data.db"
	}

	fmt.Printf("About to log something %s", log)
	log.Info().Msgf("Connecting to SQLite DB: %s", dataBaseFile)

	// Connect to the SQLite database
	db, err := gorm.Open(sqlite.Open(dataBaseFile), &gorm.Config{})

	return db, err
}
