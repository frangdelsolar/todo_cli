package currency

import (
	"fmt"
	"time"

	"github.com/frangdelsolar/todo_cli/pkg/auth"
)

// UpdateAccountBalance updates the account balance based on the provided parameters.
//
// Parameters:
// - accountId: the ID of the account to update.
// - currencyCode: the currency code for the transaction.
// - amount: the amount to update the balance.
// - date: the date of the transaction.
// - concept: the concept of the transaction.
// - tType: the type of transaction (credit or debit).
// Returns a pointer to the updated transaction and an error if any.
func UpdateAccountBalance(accountId, currencyCode string, amount string, date string, concept string, tType string, requestedBy string) (*Transaction, error) {
	user, err := auth.GetUserById(requestedBy)
	if err != nil {
		log.Err(err).Msg("Error getting user")
		return nil, err
	}
	
	acc, err := GetAccountById(accountId, requestedBy)
	if err != nil {
		log.Err(err).Msg("Error getting account")
		return nil, err
	}

	c, err := NewCurrency(currencyCode, amount, date, user)
	if err != nil {
		log.Err(err).Msg("Error creating currency")
		return nil, err
	}

	exDate, err := time.Parse(time.DateOnly, date)
	if err != nil {
		log.Err(err).Msg("Error parsing date")
		return nil, err
	}

	transaction, err := RegisterTransaction(acc, c, exDate, concept, tType, user)
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
// - *Account: a pointer to the retrieved account, or nil if not found.
// - error: an error if there was a problem retrieving the account.
func GetAccountById(id string, requestedBy string) (*Account, error) {
	var acc Account
	
	db.First(&acc, "id = ?", id).Where("created_by = ?", requestedBy)
	
	if acc == (Account{}) {
		return nil, fmt.Errorf("account with ID %s not found", fmt.Sprint(id))
	}
	
	return &acc, nil
}

// GetAllAccounts retrieves all accounts from the database.
//
// It initializes an empty slice of Account and populates it with the results of a database query.
// If no accounts are found, it logs a warning message.
//
// Returns:
// - []Account: a slice of Account containing all the accounts retrieved from the database.
func GetAllAccounts(requestedBy string) []Account {
	var accs []Account
	
	db.Find(&accs).Where("created_by = ?", requestedBy)

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
// - *Account: a pointer to the newly created account, or nil if there was an error.
// - error: an error if there was a problem creating the account.
func CreateAccount(accountName string, amount string, currencyCode string, defaultAccount bool, requestedBy string) (*Account, error) {

	u, err := auth.GetUserById(requestedBy)
	if err != nil {
		return nil, err
	}

	c, err := CreateCurrency(currencyCode, amount, time.Now().Format(time.DateOnly), requestedBy)
	if err != nil {
		return nil, err
	}

	acc, err := NewAccount(accountName, c, defaultAccount, u)
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
// - *Account: a pointer to the updated account, or nil if there was an error.
// - error: an error if there was a problem updating the account.
func UpdateAccountName(id string, accountName string, requestedBy string) (*Account, error) {

	acc, err := GetAccountById(id, requestedBy)
	if err != nil {
		return nil, err
	}

	user, err := auth.GetUserById(requestedBy)
	if err != nil {
		return nil, err
	}

	err = acc.UpdateName(accountName, user)
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
func DeleteAccount(id string, requestedBy string) error {

	acc, err := GetAccountById(id, requestedBy)
	if err != nil {
		return err
	}

	// Delete total for account as well
	db.Delete(&acc.Total)

	db.Delete(&acc)
	
	return nil
}
