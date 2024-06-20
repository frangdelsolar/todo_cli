package currency_test

import (
	"fmt"
	"strconv"

	"github.com/frangdelsolar/todo_cli/pkg/auth"
	c "github.com/frangdelsolar/todo_cli/pkg/currency"
)

// TestCreateAccount tests the CreateAccount function in the currency package.
//
// It performs the following steps:
// 1. Prepares the necessary data, including creating a user and an account.
// 2. Calls the CreateAccount function to create an account.
// 3. Asserts that the created account has the expected properties, such as name, currency code, total amount, default account status, and created by user.
//
// The function does not take any parameters and does not return any values.
func TestCreateAccount(){
    log.Info().Msg("Testing CreateAccount()")
    
    // Data prep
    owner, _ := auth.CreateUser("owner", "owner@admin.com")
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
    
    log.Debug().Interface("Account", account).Msg("Created Account")
    log.Info().Msg("Created Account Successfully")

    // assertions

    if account.Name != accountName {
        err = fmt.Errorf("expected name %s, got %s", accountName, account.Name)
        log.Err(err).Msg("TestCreateAccount()")
    } else {
        log.Debug().Msgf("Expected name %s, got %s", accountName, account.Name)
    }

    if account.CurrencyCode != cCode {
        err = fmt.Errorf("expected code %s, got %s", fmt.Sprint(cCode), fmt.Sprint(account.CurrencyCode))
        log.Err(err).Msg("TestCreateAccount()")
    } else {
        log.Debug().Msgf("Expected code %s, got %s", fmt.Sprint(cCode), fmt.Sprint(account.CurrencyCode))
    }
    
    if account.Total.Amount != amountFloat {
        err = fmt.Errorf("expected amount %f, got %f", amountFloat, account.Total.Amount)
        log.Err(err).Msg("TestCreateAccount()")
    } else {
        log.Debug().Msgf("Expected amount %f, got %f", amountFloat, account.Total.Amount)
    }

    if account.DefaultAccount != defaultAccount {
        err = fmt.Errorf("expected defaultAccount %t, got %t", defaultAccount, account.DefaultAccount)
        log.Err(err).Msg("TestCreateAccount()")
    } else {
        log.Debug().Msgf("Expected defaultAccount %t, got %t", defaultAccount, account.DefaultAccount)
    }

    if account.CreatedBy.ID != owner.ID {
        err = fmt.Errorf("expected createdBy %s, got %s", owner.Name, account.CreatedBy.Name)
        log.Err(err).Msg("TestCreateAccount()")
    } else {
        log.Debug().Msgf("Expected createdBy %s, got %s", owner.Name, account.CreatedBy.Name)
    }
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
func TestUpdateAccountCredit(){
    log.Info().Msg("Testing UpdateAccountCredit()")

    // Data prep
    owner, _ := auth.CreateUser("owner", "owner@admin.com")
    userId := fmt.Sprint(owner.ID)

    accountName:= "Savings"
    code := "USD"
    amount := "100.00"
    defaultAccount := true

    account, err := c.CreateAccount(accountName, amount, code, defaultAccount, userId)
    if err != nil {
        log.Err(err).Msg("Failed to create account")
    }
    log.Debug().Interface("Account", account).Msg("Created Account")
    
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

    log.Debug().Interface("Transaction", transaction).Msg("Created Transaction")
    log.Info().Msg("Created Transaction Successfully")

    // Assertions
    log.Debug().Msg("Validating created transaction...")
    if transaction.AccountID != account.ID {
        err = fmt.Errorf("expected accountId %s, got %d", accountId, transaction.AccountID)
        log.Err(err).Msg("TestUpdateAccountCredit()")
    } else {
        log.Debug().Msgf("Expected accountId %s, got %d", accountId, transaction.AccountID)
    }

    if transaction.TypeOfTrasaction != c.Credit {
        err = fmt.Errorf("expected typeOfTrasaction %s, got %s", c.Credit, transaction.TypeOfTrasaction)
        log.Err(err).Msg("TestUpdateAccountCredit()")
    } else {
        log.Debug().Msgf("Expected typeOfTrasaction %s, got %s", c.Credit, transaction.TypeOfTrasaction)
    }

    if fmt.Sprintf("%.2f", transaction.Amount.Amount) != creditAmount {
        err = fmt.Errorf("expected amount %s, got %s", creditAmount, fmt.Sprintf("%.2f", transaction.Amount.Amount))
        log.Err(err).Msg("TestUpdateAccountCredit()")
    } else {
        log.Debug().Msgf("Expected amount %s, got %s", creditAmount, fmt.Sprintf("%.2f", transaction.Amount.Amount))
    }

    if transaction.Amount.CurrencyCode != c.CurrencyUnit(creditCode) {
        err = fmt.Errorf("expected code %s, got %s", creditCode, transaction.Amount.CurrencyCode)
        log.Err(err).Msg("TestUpdateAccountCredit()")
    } else {
        log.Debug().Msgf("Expected code %s, got %s", creditCode, transaction.Amount.CurrencyCode)
    }

    if transaction.Details != creditConcept {
        err = fmt.Errorf("expected concept %s, got %s", creditConcept, transaction.Details)
        log.Err(err).Msg("TestUpdateAccountCredit()")
    } else {
        log.Debug().Msgf("Expected concept %s, got %s", creditConcept, transaction.Details)
    }

    // Balance should be updated
    account, err = c.GetAccountById(accountId, userId)
    if err != nil {
        log.Err(err).Msg("Failed to get account")
    }
    log.Debug().Interface("Account", account).Msg("Updated Account")


    newTotal, err := c.GetCurrencyById(fmt.Sprint(account.TotalID), userId)
    if err != nil {
        log.Err(err).Msg("Failed to get currency")
    }
    log.Debug().Interface("Balance", newTotal).Msg("New Account Balance")


    // assertions on balance
    log.Debug().Msg("Validating updated balance on account...")
    balanceExpectedAmount := 1100.00
    balanceExpectedCode := c.CurrencyUnit("USD")
    balanceExpectedConversion := 1100.00
    balanceExpectedRate := 1.00

    if newTotal.Amount != balanceExpectedAmount {
        err = fmt.Errorf("expected amount %f, got %f", balanceExpectedAmount, newTotal.Amount)
        log.Err(err).Msg("TestUpdateAccountCredit()")
    } else {
        log.Debug().Msgf("Expected amount %f, got %f", balanceExpectedAmount, newTotal.Amount)
    }

    if newTotal.CurrencyCode != balanceExpectedCode {
        err = fmt.Errorf("expected code %s, got %s", balanceExpectedCode, newTotal.CurrencyCode)
        log.Err(err).Msg("TestUpdateAccountCredit()")
    } else {
        log.Debug().Msgf("Expected code %s, got %s", balanceExpectedCode, newTotal.CurrencyCode)
    }

    if newTotal.Conversion != balanceExpectedConversion {
        err = fmt.Errorf("expected conversion %f, got %f", balanceExpectedConversion, newTotal.Conversion)
        log.Err(err).Msg("TestUpdateAccountCredit()")
    } else {
        log.Debug().Msgf("Expected conversion %f, got %f", balanceExpectedConversion, newTotal.Conversion)
    }

    if newTotal.ExchangeRate != balanceExpectedRate {
        err = fmt.Errorf("expected rate %f, got %f", balanceExpectedRate, newTotal.ExchangeRate)
        log.Err(err).Msg("TestUpdateAccountCredit()")
    } else {
        log.Debug().Msgf("Expected rate %f, got %f", balanceExpectedRate, newTotal.ExchangeRate)
    }
}


