package main

import (
	"fmt"
	"time"

	c "github.com/frangdelsolar/todo_cli/pkg/currency"
	"github.com/frangdelsolar/todo_cli/pkg/data"
	"github.com/frangdelsolar/todo_cli/pkg/logger"
)

func main(){
	/*
		This is just a playground to test the currency package.
		Feel free to use it as a template.
	*/
	c.InitCurrency()

	db:= data.GetDB()
	log:= logger.GetLogger()

	log.Debug().Interface("DB", db).Msg("main.go")

	date:= time.Date(2023, 12, 12, 0, 0, 0, 0, time.UTC)

	a, _ := c.CreateCurrency("ARS", "1", date.Format(time.DateOnly))
	fmt.Println(a)

	acc, err := c.NewAccount("test", "ARS", *a, false)
	if err != nil {
		log.Err(err).Msg("Error creating account")
		return
	}
	log.Info().Interface("Account", acc).Msg("Account created")
}





