package db

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/rs/zerolog/log"
)

const dbFileName = "db.json"
var dbPath = filepath.Join("./db", dbFileName)


type DB struct {
    Tasks map[string]Task `json:"tasks"`
}

func (db *DB) init() {
    log.Debug().Msg("Initializing DB")

    // Create the "db" directory if it doesn't exist
    err := os.MkdirAll(filepath.Dir(dbPath), os.ModePerm) // Create all necessary directories
    if err != nil {
        log.Err(err).Msg("Failed to create directory structure")
        return
    }

    // Create the "db.json" file if it doesn't exist
    _, err = os.Stat(dbPath)
    if os.IsNotExist(err) {
        f, err := os.Create(dbPath)
        if err != nil {
            log.Err(err).Msg("Failed to create DB file")
            return
        }
        defer f.Close()
    }

	// Read the contents of the "db.json" file
	data, err := os.ReadFile(dbPath)
	if err != nil {
		log.Err(err).Msg("Failed to read DB file")
		return
	}

	if len(data) > 0 {
		err = json.Unmarshal(data, &db)
		if err != nil {
			log.Err(err).Msg("Failed to unmarshal DB file")
			return
		}
	} else {
		db.Tasks = make(map[string]Task)
	}

	log.Info().Msg("DB initialized")
}

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

func GetDB() *DB {
	db := DB{}
	db.init()
	return &db
}