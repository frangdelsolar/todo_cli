package currency

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name"`
	Total    Currency    `json:"total"`
	DefaultAccount bool `json:"defaultAccount"`
	Currency CurrencyUnit `json:"currency"` 
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func NewAccount (name string, currencyCode string, total Currency, defaultAccount bool) (*Account, error) {

	if err := AccountNameValidator(name); err != nil {
		log.Err(err).Msg("Error validating account name")
		return nil, err
	}
	
	if err := CurrencyCodeValidator(currencyCode); err != nil {
		log.Err(err).Msg("Error validating currency code")
		return nil, err
	}

	return &Account{
		Name: name,
		Currency: CurrencyUnit(currencyCode),
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