package db

import (
	"os"
	"todo_cli/models"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

func InitDB() (Database, error) {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	db, err := gorm.Open(sqlite.Open("../test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	log.Debug().Msg("Initialized SQLite DB")
	db.AutoMigrate(&models.Task{}, &models.EffectivePeriod{}, &models.TaskCompletionLog{})

	return Database{db}, nil
}
