package currency

import (
	"fmt"
	"time"

	"github.com/frangdelsolar/todo_cli/pkg/auth"
	d "github.com/frangdelsolar/todo_cli/pkg/data/models"
)

type TransactionType string

const (
	Credit TransactionType = "credit"
	Debit  TransactionType = "debit"
)

type Transaction struct {
	d.SystemData
	DateOfTransaction time.Time `json:"dateOfTransaction"`
	TypeOfTrasaction TransactionType    `json:"typeOfTrasaction"`
	Account      *Account  `json:"account" gorm:"foreignKey:AccountID"`
	AccountID    uint      `json:"accountId"`
	Amount       *Currency `json:"amount"`
	AmountID      uint `json:"amountId"`
	Details     string    `json:"details"`
}

// String returns a string representation of the Transaction.
//
// It formats the Transaction by combining the type of transaction and the details.
// The type of transaction is obtained from the `TypeOfTrasaction` field,
// and the details are obtained from the `Details` field.
//
// Returns:
// - string: a string representation of the Transaction.
func (t *Transaction) String() string {
	return fmt.Sprintf("%s: %s", t.TypeOfTrasaction, t.Details)
}

// NewTransaction creates a new transaction.
//
// Parameters:
// - tType: the type of the transaction as a string.
// - account: the account associated with the transaction.
// - amount: the amount of the transaction.
// - date: the date of the transaction.
// - details: the details of the transaction.
//
// Returns:
// - *Transaction: the newly created transaction.
// - error: an error if there was a validation issue.
func NewTransaction(tType string, account *Account, amount *Currency, date time.Time, details string, requestedBy *auth.User) (*Transaction, error) {
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
		SystemData:        d.SystemData{
			CreatedBy: requestedBy, 
			UpdatedBy: requestedBy, 
		},
	}
	return transaction, nil
}

// TransactionTypeValidator validates the transaction type.
//
// Parameters:
// - tType: the type of the transaction as a string.
// Return type: error.
func TransactionTypeValidator(tType string) error {
	if tType != string(Credit) && tType != string(Debit) {
		return fmt.Errorf("invalid transaction type")
	}
	return nil
}

// AccountValidator validates an account.
//
// It checks if the provided account is nil and returns an error if it is.
//
// Parameters:
// - account (*Account): The account to be validated.
//
// Returns:
// - error: An error if the account is nil, otherwise nil.
func AccountValidator(account *Account) error {
	if account == nil {
		return fmt.Errorf("account cannot be nil")
	}
	return nil
}

// DetailsValidator validates the details of a transaction.
//
// Parameters:
// - details: the details of the transaction as a string.
// Return type: error.
func DetailsValidator(details string) error {
	if details == "" {
		return fmt.Errorf("details cannot be empty")
	}
	return nil
}

