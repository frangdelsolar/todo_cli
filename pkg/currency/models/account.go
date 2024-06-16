package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name"`
	TotalID    uint      `json:"totalId"`
	Total   *Currency    `json:"total" gorm:"foreignKey:TotalID"`
	DefaultAccount bool `json:"defaultAccount"`
	Currency CurrencyUnit `json:"currency"` 
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
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


