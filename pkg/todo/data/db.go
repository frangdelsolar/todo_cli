package data

import (
	"github.com/frangdelsolar/todo_cli/pkg/todo/logger"

	"gorm.io/gorm"
)

var log *logger.Logger

var DB *Database
type Database struct{ 
	*gorm.DB
}



func InitDB(db *gorm.DB) {
	DB = &Database{db}
}

func init(){
	log = logger.GetLogger()
}