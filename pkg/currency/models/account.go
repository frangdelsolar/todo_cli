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

func (a *Account) Update(name string) error {
	if err := AccountNameValidator(name); err != nil {
		log.Err(err).Msg("Error validating account name")
		return err
	}
	a.Name = name
	return nil
}

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

func AccountNameValidator(name string) error {
	if name == "" {
		return fmt.Errorf("account name cannot be empty")
	}
	return nil
}


