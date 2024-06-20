package currency

import (
	"fmt"
	"time"

	"github.com/frangdelsolar/todo_cli/pkg/auth"
)


func CreateTransaction(
    tType string, 
    accountId string, 
    amount string, 
    currencyCode string, 
    date string, 
    concept string, 
    requestedBy string,
) (*Transaction, error) {
    var err error
    var t *Transaction

    user, err := auth.GetUserById(requestedBy)
    if err != nil {
        log.Err(err).Msg("Error getting user")
        return t, err
    }

    total, err := CreateCurrency(currencyCode, amount, date, requestedBy)
    if err != nil {
        log.Err(err).Msg("Error creating currency")
        return t, err
    }

    account, err := GetAccountById(accountId, requestedBy)
    if err != nil {
        log.Err(err).Msg("Error getting account")
        return t, err
    }

    formattedDate, err := time.Parse(time.DateOnly, date)
    if err != nil {
        log.Err(err).Msg("Error parsing date")
        return t, err
    }
    
    transaction, err := NewTransaction(tType, account, total, formattedDate, concept, user)
    if err != nil {
        log.Err(err).Msg("Error creating transaction")
        return t, err
    }

    db.Create(&transaction)

    log.Info().Interface("Transaction", transaction).Msg("Transaction created successfully")

    return transaction, nil
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
func RegisterTransaction(
    account *Account, 
    amount *Currency, 
    date time.Time, 
    concept string, 
    strTType string, 
    requestedBy *auth.User,
) (*Transaction, error) {
	updatedBalance := &Currency{}
	var err error

	if err = TransactionTypeValidator(strTType); err != nil {
		log.Err(err).Msg("Error validating transaction type")
		return nil, err
	}
	tType := TransactionType(strTType)

    currentBalance, err := GetCurrencyById(fmt.Sprint(account.TotalID), fmt.Sprint(requestedBy.ID))
    if err != nil {
        log.Err(err).Msg("Error getting current balance")
        return nil, err
    }

	if tType == Credit {
		updatedBalance, err = AddCurrency(amount, &currentBalance, date)
		if err != nil {
			log.Err(err).Msg("Error crediting account")
			return nil, err
		}
	} else if tType == Debit {
		updatedBalance, err = SubCurrency(amount, &currentBalance, date)
		if err != nil {
			log.Err(err).Msg("Error debiting account")
			return nil, err
		}
	} else {
		log.Err(err).Msg("Invalid transaction type")
		return nil, fmt.Errorf("invalid transaction type")
	}

	account.Total = updatedBalance

	transaction, err := NewTransaction(strTType, account, amount, date, concept, requestedBy)
	if err != nil {
		log.Err(err).Msg("Error creating transaction")
		return nil, err
	}

	return transaction, nil
}
