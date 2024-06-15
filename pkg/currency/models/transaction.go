package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type TransactionType string

const (
	Credit TransactionType = "credit"
	Debit  TransactionType = "debit"
)

type Transaction struct {
	gorm.Model
	ID          uint      `json:"id" gorm:"primaryKey"`
	DateOfTransaction time.Time `json:"dateOfTransaction"`
	TypeOfTrasaction TransactionType    `json:"typeOfTrasaction"`
	Account      *Account  `json:"account" gorm:"foreignKey:AccountID"`
	AccountID    uint      `json:"accountId"`
	Amount       *Currency `json:"amount"`
	AmountID      uint `json:"amountId"`
	Details     string    `json:"details"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func (t *Transaction) String() string {
	return fmt.Sprintf("%s: %s", t.TypeOfTrasaction, t.Details)
}

func NewTransaction(tType string, account *Account, amount *Currency, date time.Time, details string) (*Transaction, error) {
	if err := TransactionTypeValidator(tType); err != nil {
		log.Err(err).Msg("Error validating transaction type")
		return nil, err
	}
	if err := AccountValidator(account); err != nil {
		log.Err(err).Msg("Error validating account")
		return nil, err
	}
	if err := DetailsValidator(details); err != nil {
		log.Err(err).Msg("Error validating details")
		return nil, err
	}

	transaction := &Transaction{
		TypeOfTrasaction: TransactionType(tType),
		Account:           account,
		Amount:            amount,
		Details:           details,
		AccountID:         account.ID,
		DateOfTransaction: date,
	}
	return transaction, nil
}

func TransactionTypeValidator(tType string) error {
	if tType != string(Credit) && tType != string(Debit) {
		return fmt.Errorf("invalid transaction type")
	}
	return nil
}

func AccountValidator(account *Account) error {
	if account == nil {
		return fmt.Errorf("account cannot be nil")
	}
	return nil
}

func DetailsValidator(details string) error {
	if details == "" {
		return fmt.Errorf("details cannot be empty")
	}
	return nil
}

