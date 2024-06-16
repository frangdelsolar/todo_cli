package main

import (
	"fmt"

	c "github.com/frangdelsolar/todo_cli/pkg/currency"
	d "github.com/frangdelsolar/todo_cli/pkg/currency/data"
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
		log := logger.GetLogger()
		log.Err(err).Msg("Error connecting to database")
		return
	}

	log.Debug().Interface("DB", db).Msg("main.go")

	c.InitCurrency()

	acc, err := d.CreateAccount("test", "1000", "ARS", false)
	if err != nil {
		log.Err(err).Msg("Error creating account")
		return
	}
	log.Info().Interface("Account", acc).Msg("Account created")


	t, err:= d.UpdateAccountBalance(fmt.Sprint(acc.ID), "ARS", "100", "2022-01-01", "Debito", "credit")
	if err != nil {
		log.Err(err).Msg("Error updating account balance")
		return
	}
	log.Info().Interface("Transaction", t).Msg("Transaction created")

	log.Info().Interface("Account", acc).Msg("Account created")

}
