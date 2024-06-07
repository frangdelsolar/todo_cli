package db

import (
	"fmt"
	"os"
	"todo_cli/models"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

func InitDB() (Database, error) {

	envFile := "../.env"
	env := os.Getenv("APP_ENV")

	switch env {
	case "test":
		envFile = "../.env.test"
	}
	err := godotenv.Load(envFile)
	if err != nil {
		fmt.Println("Error loading .env file", err)
	}

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	dataBaseFile := os.Getenv("DATA_BASE_FILE")
	if dataBaseFile == "" {
		dataBaseFile = "../data.db"
	}
	log.Info().Msgf("Connecting to SQLite DB: %s", dataBaseFile)
	db, err := gorm.Open(sqlite.Open(dataBaseFile), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	log.Debug().Msg("Initialized SQLite DB")
	db.AutoMigrate(&models.Task{}, &models.EffectivePeriod{}, &models.TaskCompletionLog{})

	return Database{db}, nil
}
