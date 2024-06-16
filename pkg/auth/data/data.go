package data

import (
	"github.com/frangdelsolar/todo_cli/pkg/data"
	"github.com/frangdelsolar/todo_cli/pkg/logger"
)

var log = logger.GetLogger()
var db *data.Database

func init(){
	var err error
	db, err = data.GetDB()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
		return 
	}
}