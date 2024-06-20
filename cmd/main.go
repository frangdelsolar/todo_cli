package main

import (
	"fmt"
	"time"

	"github.com/frangdelsolar/todo_cli/pkg/auth"
	"github.com/frangdelsolar/todo_cli/pkg/config"
	c "github.com/frangdelsolar/todo_cli/pkg/currency"
	"github.com/frangdelsolar/todo_cli/pkg/data"
	"github.com/frangdelsolar/todo_cli/pkg/logger"
)

var APP_NAME = "TODO APP"
var APP_VERSION = "1.5.0"

func main() {

	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	log := logger.NewLogger(APP_NAME, APP_VERSION)
	log.Info().Msgf("Running %s v%s", APP_NAME, APP_VERSION)
	log.Debug().Interface("Config", cfg).Msg("Loaded Config")

	db, err := data.LoadDB()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
		panic(err)
	}
	log.Debug().Msgf("Loaded Database: %s", db.Name())

	auth.InitAuth()

	u, err := auth.CreateUser("pepe", "pepe@admin.com")
	if err != nil {
		log.Err(err).Msg("Failed to create user")
	}
	log.Info().Interface("User", u).Msg("Created User")

	c.InitCurrency()

	acc, err := c.CreateAccount("pepe", "90", "USD", true, fmt.Sprint(u.ID))
	if err != nil {
		log.Err(err).Msg("Failed to create account")
	}
	log.Info().Interface("Account", acc).Msg("Created Account")

	cur, err := c.NewCurrency("USD", "100", "2022-01-01", u)
	if err != nil {
		log.Err(err).Msg("Failed to create currency")
	}
	log.Info().Interface("Currency", cur).Msg("Created Currency")

	t, err := c.RegisterTransaction(acc, cur, time.Now(), "test", "credit", u)
	if err != nil {
		log.Err(err).Msg("Failed to create transaction")
	}
	log.Info().Interface("Transaction", t).Msg("Created Transaction")

}
