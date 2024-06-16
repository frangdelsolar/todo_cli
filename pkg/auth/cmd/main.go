package main

import (
	a "github.com/frangdelsolar/todo_cli/pkg/auth"
	b "github.com/frangdelsolar/todo_cli/pkg/auth/data"
	"github.com/frangdelsolar/todo_cli/pkg/data"
	"github.com/frangdelsolar/todo_cli/pkg/logger"
)

func main() {
	/*
		This is just a playground to test the currency package.
		Feel free to use it as a template.
	*/

	log := logger.GetLogger()

	db, err := data.GetDB()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
		return 
	}

	log.Debug().Interface("DB", db).Msg("main.go")

	a.InitAuth()

	u, err := b.CreateUser("frangdelsolar", "jU6wz@example.com")
	log.Debug().Interface("u", u).Err(err).Msg("main.go")

}
