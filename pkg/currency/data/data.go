package data

import (
	"github.com/frangdelsolar/todo_cli/pkg/data"
	"github.com/frangdelsolar/todo_cli/pkg/logger"
)

var log *logger.Logger
var db *data.Database


// init initializes the package by setting up the logger and connecting to the database.
//
// It does not take any parameters.
// It does not return any values.
func init(){
	var err error

	log = logger.GetLogger()

	db, err = data.GetDB()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
		return
	}
}