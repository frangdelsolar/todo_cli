package db

import (
	"encoding/json"
	"os"
	"path/filepath"
	"todo_cli/models"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const dbFileName = "db.json"
var dbPath = filepath.Join("./db", dbFileName)

// init initializes the logger to output to the console.
func init(){
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}

// DB represents the database content.
type DB struct {
    Tasks map[string]models.Task `json:"tasks"`
	EffectivePeriods map[string]models.EffectivePeriod `json:"taskPeriods"`
	TaskCompletionLogs map[string]models.TaskCompletionLog `json:"taskCompletionLogs"`
}

// NewDB initializes a new DB instance.
//
// It creates the necessary directories and the "db.json" file if they don't exist.
// It reads the contents of the "db.json" file and unmarshals it into the DB struct.
// If the file is empty, it initializes the Tasks and EffectivePeriods maps.
// It returns a pointer to the DB instance and an error if any.
func NewDB() (*DB, error){
    log.Debug().Msg("Initializing DB")

	var db DB

    // Create the "db" directory if it doesn't exist
    err := os.MkdirAll(filepath.Dir(dbPath), os.ModePerm) // Create all necessary directories
    if err != nil {
        log.Err(err).Msg("Failed to create directory structure")
        return nil, err
    }

    // Create the "db.json" file if it doesn't exist
    _, err = os.Stat(dbPath)
    if os.IsNotExist(err) {
        f, err := os.Create(dbPath)
        if err != nil {
            log.Err(err).Msg("Failed to create DB file")
            return nil, err
        }
        defer f.Close()
    }

	// Read the contents of the "db.json" file
	data, err := os.ReadFile(dbPath)
	if err != nil {
		log.Err(err).Msg("Failed to read DB file")
		return nil, err
	}

	if len(data) > 0 {
		err = json.Unmarshal(data, &db)
		if err != nil {
			log.Err(err).Msg("Failed to unmarshal DB file")
			return nil, err
		}
	} else {
		db.Tasks = map[string]models.Task{}
		db.EffectivePeriods = map[string]models.EffectivePeriod{}
		db.TaskCompletionLogs = map[string]models.TaskCompletionLog{}
	}

	log.Info().Msg("DB initialized")

	return &db, nil
}

// Save saves the DB instance to a JSON file.
//
// It marshals the DB instance into JSON format with indentation and writes it to
// the file specified by dbPath. If any error occurs during the marshaling or
// writing process, it logs the error and returns. Otherwise, it logs a message
// indicating that the DB has been saved.
func (db *DB) Save() {
	log.Debug().Msg("Saving DB")
	data, err := json.MarshalIndent(db, "", "\t")
	if err != nil {
		log.Err(err).Msg("Failed to marshal DB")
		return
	}
	err = os.WriteFile(dbPath, data, 0644)
	if err != nil {
		log.Err(err).Msg("Failed to write DB file")
		return
	}
	log.Info().Msg("DB saved")
}
