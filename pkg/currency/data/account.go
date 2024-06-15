package data

import (
	"fmt"
	"time"

	m "github.com/frangdelsolar/todo_cli/pkg/currency/models"
)

func UpdateAccountBalance(accountId, currencyCode string, amount string, date string, concept string, tType string) (*m.Transaction, error) {
	acc, err := GetAccountById(accountId)
	if err != nil {
		log.Err(err).Msg("Error getting account")
		return nil, err
	}

	c, err := m.NewCurrency(currencyCode, amount, date)
	if err != nil {
		log.Err(err).Msg("Error creating currency")
		return nil, err
	}

	exDate, err := time.Parse(time.DateOnly, date)
	if err != nil {
		log.Err(err).Msg("Error parsing date")
		return nil, err
	}

	balance, err := GetCurrencyById(fmt.Sprint(acc.TotalID))
	if err != nil {
		log.Err(err).Msg("Error getting account balance")
		return nil, err
	}

	transaction, err := acc.RegisterTransaction(&balance, c, exDate, concept, tType)
	if err != nil {
		log.Err(err).Msg("Error creating transaction")
		return nil, err
	}

	db.Save(&transaction)
	db.Save(&acc)

	return transaction, nil
}


// GetAccountById retrieves an account by its ID.
//
// Parameters:
// - id: the ID of the account to retrieve.
//
// Returns:
// - *m.Account: a pointer to the retrieved account, or nil if not found.
// - error: an error if there was a problem retrieving the account.
func GetAccountById(id string) (*m.Account, error) {
	var acc m.Account
	db.First(&acc, "id = ?", id)
	if acc == (m.Account{}) {
		return nil, fmt.Errorf("account with ID %s not found", fmt.Sprint(id))
	}
	return &acc, nil
}

// GetAllAccounts retrieves all accounts from the database.
//
// It initializes an empty slice of m.Account and populates it with the results of a database query.
// If no accounts are found, it logs a warning message.
//
// Returns:
// - []m.Account: a slice of m.Account containing all the accounts retrieved from the database.
func GetAllAccounts() []m.Account {
	var accs []m.Account
	
	db.Find(&accs)

	if len(accs) == 0 {
		log.Warn().Msg("No accounts found")
	}

	return accs
}

// CreateAccount creates a new account with the given accountName, amount, currencyCode, and defaultAccount.
//
// Parameters:
// - accountName: the name of the account to be created.
// - amount: the initial amount of the account.
// - currencyCode: the currency code of the account.
// - defaultAccount: a boolean indicating whether the account is the default account.
//
// Returns:
// - *m.Account: a pointer to the newly created account, or nil if there was an error.
// - error: an error if there was a problem creating the account.
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

// UpdateAccountName updates the name of an account with the given ID.
//
// Parameters:
// - id: the ID of the account to be updated.
// - accountName: the new name for the account.
//
// Returns:
// - *m.Account: a pointer to the updated account, or nil if there was an error.
// - error: an error if there was a problem updating the account.
func UpdateAccountName(id string, accountName string) (*m.Account, error) {

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

// DeleteAccount deletes an account by ID.
//
// Parameters:
// - id: the ID of the account to be deleted.
// Return type:
// - error: an error if there was a problem deleting the account.
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