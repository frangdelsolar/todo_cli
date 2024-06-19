package currency

import (
	"fmt"
	"time"

	d "github.com/frangdelsolar/todo_cli/pkg/data/models"
)

type Account struct {
	d.SystemData
	Name        string    `json:"name"`
	TotalID    uint      `json:"totalId"`
	Total   *Currency    `json:"total" gorm:"foreignKey:TotalID"`
	DefaultAccount bool `json:"defaultAccount"`
	Currency CurrencyUnit `json:"currency"` 
}

func (a *Account) String() string {
	return fmt.Sprintf("%s (%s)", a.Name, a.Currency)
}

// UpdateName updates the name of the account.
//
// Parameters:
// - name: the new name for the account.
//
// Returns:
// - error: an error if there was a problem updating the account name.
func (a *Account) UpdateName(name string) error {
	if err := AccountNameValidator(name); err != nil {
		log.Err(err).Msg("Error validating account name")
		return err
	}
	a.Name = name
	return nil
}

// RegisterTransaction registers a transaction for the account.
//
// Parameters:
// - currentBalance: the current balance of the account.
// - amount: the amount of the transaction.
// - date: the date of the transaction.
// - concept: the concept of the transaction.
// - strTType: the type of the transaction as a string.
//
// Returns:
// - *Transaction: the registered transaction.
// - error: an error if there was a problem registering the transaction.
func (a *Account) RegisterTransaction(currentBalance *Currency, amount *Currency, date time.Time, concept string, strTType string) (*Transaction, error) {
	updatedBalance := &Currency{}
	var err error

	if err = TransactionTypeValidator(strTType); err != nil {
		log.Err(err).Msg("Error validating transaction type")
		return nil, err
	}
	tType := TransactionType(strTType)

	if tType == Credit {
		updatedBalance, err = AddCurrency(amount, currentBalance, date)
		if err != nil {
			log.Err(err).Msg("Error crediting account")
			return nil, err
		}
	} else if tType == Debit {
		updatedBalance, err = SubCurrency(amount, currentBalance, date)
		if err != nil {
			log.Err(err).Msg("Error debiting account")
			return nil, err
		}
	} else {
		log.Err(err).Msg("Invalid transaction type")
		return nil, fmt.Errorf("invalid transaction type")
	}

	a.Total = updatedBalance

	transaction, err := NewTransaction(strTType, a, amount, date, concept)
	if err != nil {
		log.Err(err).Msg("Error creating transaction")
		return nil, err
	}

	return transaction, nil
}

// NewAccount creates a new account with the given name, total Currency, and defaultAccount flag.
//
// Parameters:
// - name: the name of the account.
// - total: the total Currency of the account.
// - defaultAccount: a boolean indicating if it is the default account.
//
// Returns:
// - *Account: the newly created Account.
// - error: an error if there was a problem during creation.
func NewAccount (name string, total *Currency, defaultAccount bool) (*Account, error) {

	if err := AccountNameValidator(name); err != nil {
		log.Err(err).Msg("Error validating account name")
		return nil, err
	}

	return &Account{
		Name: name,
		Currency: total.CurrencyCode,
		Total: total,
		DefaultAccount: defaultAccount,
	}, nil
}

// AccountNameValidator validates the name of an account.
//
// Parameters:
// - name: the name of the account to be validated.
//
// Returns:
// - error: an error if the account name is empty, otherwise nil.
func AccountNameValidator(name string) error {
	if name == "" {
		return fmt.Errorf("account name cannot be empty")
	}
	return nil
}



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
func UpdateAccountBalance(accountId, currencyCode string, amount string, date string, concept string, tType string) (*Transaction, error) {
	acc, err := GetAccountById(accountId)
	if err != nil {
		log.Err(err).Msg("Error getting account")
		return nil, err
	}

	c, err := NewCurrency(currencyCode, amount, date)
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
// - *Account: a pointer to the retrieved account, or nil if not found.
// - error: an error if there was a problem retrieving the account.
func GetAccountById(id string) (*Account, error) {
	var acc Account
	db.First(&acc, "id = ?", id)
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
func GetAllAccounts() []Account {
	var accs []Account
	
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
// - *Account: a pointer to the newly created account, or nil if there was an error.
// - error: an error if there was a problem creating the account.
func CreateAccount(accountName string, amount string, currencyCode string, defaultAccount bool) (*Account, error) {

	c, err := CreateCurrency(currencyCode, amount, time.Now().Format(time.DateOnly))
	if err != nil {
		return nil, err
	}

	acc, err := NewAccount(accountName, c, defaultAccount)
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
func UpdateAccountName(id string, accountName string) (*Account, error) {

	acc, err := GetAccountById(id)
	if err != nil {
		return nil, err
	}

	err = acc.UpdateName(accountName)
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