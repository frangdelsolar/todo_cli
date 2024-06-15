package main

import (
	"time"

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
	c.InitCurrency()

	db := data.GetDB()
	log := logger.GetLogger()

	log.Debug().Interface("DB", db).Msg("main.go")

	date := time.Date(2023, 12, 12, 0, 0, 0, 0, time.UTC)

	ars1, _ := d.CreateCurrency("ARS", "1000", date.Format(time.DateOnly))
	log.Info().Interface("ars1", ars1).Msg("ars1")

	ars2, _ := d.CreateCurrency("ARS", "2000", date.Format(time.DateOnly))
	log.Info().Interface("ars2", ars2).Msg("ars2")

	ars3, _ := d.AddCurrency(ars1, ars2, date)
	log.Info().Interface("ars3", ars3).Msg("ars3")

	usd1, _ := d.CreateCurrency("USD", "1", date.Format(time.DateOnly))
	log.Info().Interface("usd1", usd1).Msg("usd1")

	usd2, _ := d.CreateCurrency("USD", "2", date.Format(time.DateOnly))
	log.Info().Interface("usd2", usd2).Msg("usd2")

	usd3, _ := d.AddCurrency(usd1, usd2, date)
	log.Info().Interface("usd3", usd3).Msg("usd3")

	conv1, _ := d.AddCurrency(usd3, ars3, date)
	log.Info().Interface("conv1", conv1).Msg("conv1")

	conv2, _ := d.AddCurrency(ars3, usd3, date)
	log.Info().Interface("conv2", conv2).Msg("conv2")

	// acc, err := d.CreateAccount("test", "0", "ARS", false)
	// if err != nil {
	// 	log.Err(err).Msg("Error creating account")
	// 	return
	// }
	// log.Info().Interface("Account", acc).Msg("Account created")
}
