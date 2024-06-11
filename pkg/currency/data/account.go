package data

import (
	"time"

	m "github.com/frangdelsolar/todo_cli/pkg/currency/models"
)


func CreateAccount(accountName string, amount string, currencyCode string, defaultAccount bool) (*m.Account, error) {

	c, err := CreateCurrency(currencyCode, amount, time.Now().Format(time.DateOnly))
	if err != nil {
		return nil, err
	}

	acc, err := m.NewAccount(accountName, c, defaultAccount)
	if err != nil {
		return nil, err
	}

	db.Create(&acc)

	return acc, nil
}