package data

import (
	"github.com/frangdelsolar/todo_cli/pkg/todo/models"

	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

var log models.Logger

var DB *Database
type Database struct{ 
	*gorm.DB
}

func InitDB(db *gorm.DB, lg *zerolog.Logger) {
	DB = &Database{db}
	log = models.Logger{lg}
}