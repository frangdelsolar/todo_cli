package currency_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/frangdelsolar/todo_cli/pkg/auth"
	c "github.com/frangdelsolar/todo_cli/pkg/currency"
	"github.com/stretchr/testify/assert"
)

// TestCreateAccount tests the CreateAccount function in the currency package.
//
// It performs the following steps:
// 1. Prepares the necessary data, including creating a user and an account.
// 2. Calls the CreateAccount function to create an account.
// 3. Asserts that the created account has the expected properties, such as name, currency code, total amount, default account status, and created by user.
//
// The function does not take any parameters and does not return any values.
func TestCreateAccount(t *testing.T) {

	// Data prep
	owner, _ := auth.CreateUser("owner", "owner@admin.com", "test123")
	userId := fmt.Sprint(owner.ID)

	accountName := "Savings"
	code := "USD"
	amount := "100.00"
	defaultAccount := true

	cCode := c.CurrencyUnit(code)
	amountFloat, _ := strconv.ParseFloat(amount, 64)

	account, err := c.CreateAccount(accountName, amount, code, defaultAccount, userId)
	if err != nil {
		t.Errorf("Failed to create account: %v", err)
	}

	assert.Equal(t, account.Name, accountName, "Expected name to be %s, but got %s", accountName, account.Name)
	assert.Equal(t, account.CurrencyCode, cCode, "Expected currency code to be %s, but got %s", cCode, account.CurrencyCode)
	assert.Equal(t, account.Total.Amount, amountFloat, "Expected total amount to be %f, but got %f", amountFloat, account.Total.Amount)
	assert.Equal(t, account.DefaultAccount, defaultAccount, "Expected default account status to be %t, but got %t", defaultAccount, account.DefaultAccount)
	assert.Equal(t, account.CreatedBy.ID, owner.ID, "Expected created by user id to be %s, but got %s", owner.ID, account.CreatedBy.ID)
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
func TestUpdateAccountCredit(t *testing.T) {

	// Data prep
	owner, _ := auth.CreateUser("owner", "owner@admin.com", "test123")
	userId := fmt.Sprint(owner.ID)

	accountName := "Savings"
	code := "USD"
	amount := "100.00"
	defaultAccount := true

	account, err := c.CreateAccount(accountName, amount, code, defaultAccount, userId)
	if err != nil {
		t.Errorf("Failed to create account: %v", err)
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
		t.Errorf("Failed to update account balance: %v", err)
	}

	// Assertions
	assert.Equal(t, transaction.AccountID, account.ID, "Expected account id to be %s, but got %s", account.ID, transaction.AccountID)
	assert.Equal(t, transaction.TypeOfTrasaction, c.Credit, "Expected transaction type to be %s, but got %s", c.Credit, transaction.TypeOfTrasaction)
	assert.Equal(t, fmt.Sprintf("%.2f", transaction.Amount.Amount), creditAmount, "Expected amount to be %s, but got %s", creditAmount, fmt.Sprintf("%.2f", transaction.Amount.Amount))
	assert.Equal(t, transaction.Amount.CurrencyCode, c.CurrencyUnit(creditCode), "Expected currency code to be %s, but got %s", c.CurrencyUnit(creditCode), transaction.Amount.CurrencyCode)
	assert.Equal(t, transaction.Details, creditConcept, "Expected details to be %s, but got %s", creditConcept, transaction.Details)
	assert.Equal(t, transaction.CreatedBy.ID, owner.ID, "Expected created by user id to be %s, but got %s", owner.ID, transaction.CreatedBy.ID)

	// Balance should be updated
	account, err = c.GetAccountById(accountId, userId)
	if err != nil {
		t.Errorf("Failed to get account: %v", err)
	}

	newTotal, err := c.GetCurrencyById(fmt.Sprint(account.TotalID), userId)
	if err != nil {
		t.Errorf("Failed to get currency: %v", err)
	}

	// assertions on balance
	balanceExpectedAmount := 1100.00
	balanceExpectedCode := c.CurrencyUnit("USD")
	balanceExpectedConversion := 1100.00
	balanceExpectedRate := 1.00

	assert.Equal(t, newTotal.Amount, balanceExpectedAmount, "Expected total amount to be %f, but got %f", balanceExpectedAmount, newTotal.Amount)
	assert.Equal(t, newTotal.CurrencyCode, balanceExpectedCode, "Expected currency code to be %s, but got %s", balanceExpectedCode, newTotal.CurrencyCode)
	assert.Equal(t, newTotal.Conversion, balanceExpectedConversion, "Expected conversion to be %f, but got %f", balanceExpectedConversion, newTotal.Conversion)
	assert.Equal(t, newTotal.ExchangeRate, balanceExpectedRate, "Expected exchange rate to be %f, but got %f", balanceExpectedRate, newTotal.ExchangeRate)
}
