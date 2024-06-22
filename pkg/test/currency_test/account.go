package currency_test

import (
	"fmt"
	"strconv"

	"github.com/frangdelsolar/todo_cli/pkg/auth"
	c "github.com/frangdelsolar/todo_cli/pkg/currency"
	"github.com/frangdelsolar/todo_cli/pkg/test"
)

// TestCreateAccount tests the CreateAccount function in the currency package.
//
// It performs the following steps:
// 1. Prepares the necessary data, including creating a user and an account.
// 2. Calls the CreateAccount function to create an account.
// 3. Asserts that the created account has the expected properties, such as name, currency code, total amount, default account status, and created by user.
//
// The function does not take any parameters and does not return any values.
func TestCreateAccount(t *test.Test){
    t.Run("TestCreateAccount()")
    
    // Data prep
    owner, _ := auth.CreateUser("owner", "owner@admin.com", "test123")
    userId := fmt.Sprint(owner.ID)

    accountName:= "Savings"
    code := "USD"
    amount := "100.00"
    defaultAccount := true

    cCode := c.CurrencyUnit(code)
	amountFloat, _ := strconv.ParseFloat(amount, 64)
    
    account, err := c.CreateAccount(accountName, amount, code, defaultAccount, userId)
    if err != nil {
        log.Err(err).Msg("Failed to create account")
    }
    
    t.AssertEqual(account.Name, accountName)
    t.AssertEqual(account.CurrencyCode, cCode)
    t.AssertEqual(account.Total.Amount, amountFloat)
    t.AssertEqual(account.DefaultAccount, defaultAccount)
    t.AssertEqual(account.CreatedBy.ID, owner.ID)

    t.Stop()
}

// TestUpdateAccountCredit is a test function that tests the UpdateAccountBalance function in the currency package.
//
// It performs the following steps:
// 1. Prepares the necessary data, including creating a user and an account.
// 2. Calls the UpdateAccountBalance function to update the account balance.
// 3. Asserts that the created transaction has the expected properties, such as account ID, type, amount, currency code, and details.
// 4. Retrieves the updated account balance and asserts that it matches the expected values.
//
// The function takes no parameters and does not return any values.
func TestUpdateAccountCredit(t *test.Test){
    t.Run("TestUpdateAccountCredit()")

    // Data prep
    owner, _ := auth.CreateUser("owner", "owner@admin.com", "test123")
    userId := fmt.Sprint(owner.ID)

    accountName:= "Savings"
    code := "USD"
    amount := "100.00"
    defaultAccount := true

    account, err := c.CreateAccount(accountName, amount, code, defaultAccount, userId)
    if err != nil {
        log.Err(err).Msg("Failed to create account")
    }
    
    accountId := fmt.Sprint(account.ID)
    creditAmount := "1000.00"
    creditCode := "USD"
    creditDate := "2022-01-01"
    creditConcept := "test"
    creditType := "credit"

    // Perform test
    transaction, err := c.UpdateAccountBalance(accountId, creditCode, creditAmount, creditDate, creditConcept, creditType, userId)
    if err != nil {
        log.Err(err).Msg("Failed to update account balance")
    }

    // Assertions
    log.Debug().Msg("Validating created transaction...")

    t.AssertEqual(transaction.AccountID, account.ID)
    t.AssertEqual(transaction.TypeOfTrasaction, c.Credit)
    t.AssertEqual(fmt.Sprintf("%.2f", transaction.Amount.Amount), creditAmount)
    t.AssertEqual(transaction.Amount.CurrencyCode, c.CurrencyUnit(creditCode))
    t.AssertEqual(transaction.Details, creditConcept)
    t.AssertEqual(transaction.CreatedBy.ID, owner.ID)

    // Balance should be updated
    account, err = c.GetAccountById(accountId, userId)
    if err != nil {
        log.Err(err).Msg("Failed to get account")
    }

    newTotal, err := c.GetCurrencyById(fmt.Sprint(account.TotalID), userId)
    if err != nil {
        log.Err(err).Msg("Failed to get currency")
    }

    // assertions on balance
    log.Debug().Msg("Validating updated balance on account...")

    balanceExpectedAmount := 1100.00
    balanceExpectedCode := c.CurrencyUnit("USD")
    balanceExpectedConversion := 1100.00
    balanceExpectedRate := 1.00

    t.AssertEqual(newTotal.Amount, balanceExpectedAmount)
    t.AssertEqual(newTotal.CurrencyCode, balanceExpectedCode)
    t.AssertEqual(newTotal.Conversion, balanceExpectedConversion)
    t.AssertEqual(newTotal.ExchangeRate, balanceExpectedRate)

    t.Stop()
}


