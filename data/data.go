package data

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	dbpkg "todo_cli/db"
)

var DB dbpkg.Database

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	dbInstance, err := dbpkg.InitDB()
	if err != nil {
		panic(err)
	}
	DB = dbInstance
}
