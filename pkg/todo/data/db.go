package data

import (
	"github.com/frangdelsolar/todo_cli/pkg/data"
	"github.com/frangdelsolar/todo_cli/pkg/logger"
)

var log *logger.Logger
var db *data.Database

func init(){
	log = logger.GetLogger()
	db = data.GetDB()
}