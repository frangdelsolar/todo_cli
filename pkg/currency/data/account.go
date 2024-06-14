package data

import (
	"fmt"
	"time"

	m "github.com/frangdelsolar/todo_cli/pkg/currency/models"
)


func GetAccountById(id string) (*m.Account, error) {
	var acc m.Account
	db.First(&acc, "id = ?", id)
	if acc == (m.Account{}) {
		return nil, fmt.Errorf("account with ID %s not found", fmt.Sprint(id))
	}
	return &acc, nil
}

func GetAllAccounts() []m.Account {
	var accs []m.Account
	
	db.Find(&accs)

	if len(accs) == 0 {
		log.Warn().Msg("No accounts found")
	}

	return accs
}


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

func UpdateAccount(id string, accountName string) (*m.Account, error) {

	acc, err := GetAccountById(id)
	if err != nil {
		return nil, err
	}

	err = acc.Update(accountName)
	if err != nil {
		return nil, err
	}

	db.Save(&acc)

	return acc, nil
}

func DeleteAccount(id string) error {
	acc, err := GetAccountById(id)
	if err != nil {
		return err
	}

	// Delete total for account as well
	db.Delete(&acc.Total)

	db.Delete(&acc)
	
	return nil
}