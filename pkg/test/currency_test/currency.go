package currency_test

import (
	"fmt"
	"time"

	"github.com/frangdelsolar/todo_cli/pkg/auth"
	c "github.com/frangdelsolar/todo_cli/pkg/currency"
)

// TestAddCurrencySameCode tests the AddCurrency function when both currencies have the same code.
//
// This function performs the following steps:
// 1. Logs a message indicating that the AddCurrencySameCode test is being run.
// 2. Prepares the necessary data for the test by creating two currencies with the same code and date.
// 3. Performs the test by calling the AddCurrency function with the two currencies and a specific date.
// 4. Logs the added currency and a success message.
// 5. Asserts that the added currency has the expected amount, code, conversion, and exchange rate.
// 6. If any assertion fails, logs an error message with the expected and actual values.
//
// This function does not return any value.
func TestAddCurrencySameCode(){
    log.Info().Msg("Testing AddCurrencySameCode()")

    // data prep
    owner, _ := auth.CreateUser("owenr", "owenr@admin.com", "test123")
    ownerId := fmt.Sprint(owner.ID)
    
    aAmount := "100.00"
    aCode := "USD"
    aDate := "2024-01-01"

    a, err := c.CreateCurrency(aCode, aAmount, aDate, ownerId)
    if err != nil {
        log.Err(err).Msg("Failed to create currency")
    }
    log.Trace().Interface("Currency", a).Msg("Created A Currency")


    bAmount := "200.00"
    bCode := "USD"
    bDate := "2024-01-01"

    b, err := c.CreateCurrency(bCode, bAmount, bDate, ownerId)
    if err != nil {
        log.Err(err).Msg("Failed to create currency")
    }
    log.Trace().Interface("Currency", b).Msg("Created B Currency")

    // Perform test
    cDate := time.Date(2025, 5, 1, 0, 0, 0, 0, time.UTC)
    ccy, err := c.AddCurrency(a, b, cDate)
    if err != nil {
        log.Err(err).Msg("Failed to add currency")
    }

    log.Trace().Interface("Currency", ccy).Msg("Added Currency")
    log.Info().Msg("Added Currency Successfully")

    // assertions
    expectedAmount := 300.00
    expectedCode := c.CurrencyUnit("USD")
    expectedConversion := 300.00
    expectedRate := 1.00

    if ccy.Amount != expectedAmount {
        err = fmt.Errorf("expected amount %f, got %f", expectedAmount, ccy.Amount)
        log.Err(err).Msg("TestAddCurrencySameCode()")
    } else {
        log.Debug().Msgf("Expected amount %f, got %f", expectedAmount, ccy.Amount)
    }

    if ccy.CurrencyCode != expectedCode {
        err = fmt.Errorf("expected code %s, got %s", expectedCode, ccy.CurrencyCode)
        log.Err(err).Msg("TestAddCurrencySameCode()")
    } else {
        log.Debug().Msgf("Expected code %s, got %s", expectedCode, ccy.CurrencyCode)
    }

    if ccy.Conversion != expectedConversion {
        err = fmt.Errorf("expected conversion %f, got %f", expectedConversion, ccy.Conversion)
        log.Err(err).Msg("TestAddCurrencySameCode()")
    } else {
        log.Debug().Msgf("Expected conversion %f, got %f", expectedConversion, ccy.Conversion)
    }

    if ccy.ExchangeRate != expectedRate {
        err = fmt.Errorf("expected rate %f, got %f", expectedRate, ccy.ExchangeRate)
        log.Err(err).Msg("TestAddCurrencySameCode()")
    } else {
        log.Debug().Msgf("Expected rate %f, got %f", expectedRate, ccy.ExchangeRate)
    }
}

// TestAddCurrencyDifferentCode is a test function that tests the AddCurrency function when the currencies have different codes.
//
// It performs the following steps:
// - Creates two currencies with different codes and amounts.
// - Tests the exchange rates of the created currencies.
// - Adds the currencies together using the AddCurrency function.
// - Performs assertions on the resulting currency's amount, code, conversion, and exchange rate.
//
// The function does not take any parameters and does not return any values.
func TestAddCurrencyDifferentCode(){
    log.Info().Msg("Testing AddCurrencyDifferentCode()")

    // data prep
    owner, _ := auth.CreateUser("owenr", "owenr@admin.com", "test123")
    ownerId := fmt.Sprint(owner.ID)
    
    aAmount := "100.00"
    aCode := "USD"
    aDate := "2024-01-01"

    a, err := c.CreateCurrency(aCode, aAmount, aDate, ownerId)
    if err != nil {
        log.Err(err).Msg("Failed to create currency")
    }
    log.Trace().Interface("Currency", a).Msg("Created A Currency")


    bAmount := "8.00"
    bCode := "ARS"
    bDate := "2011-01-03"

    b, err := c.CreateCurrency(bCode, bAmount, bDate, ownerId)
    if err != nil {
        log.Err(err).Msg("Failed to create currency")
    }
    log.Trace().Interface("Currency", b).Msg("Created B Currency")

    // Test Rates
    bExpectedRate := 4.00
    bExpectedConversion := 2.00

    if b.ExchangeRate != bExpectedRate {
        err = fmt.Errorf("expected rate %f, got %f", bExpectedRate, b.ExchangeRate)
        log.Err(err).Msg("AddCurrencyDifferentCode()")
    } else {
        log.Debug().Msgf("Expected rate %f, got %f", bExpectedRate, b.ExchangeRate)
    }

    if b.Conversion != bExpectedConversion {
        err = fmt.Errorf("expected conversion %f, got %f", bExpectedConversion, b.Conversion)
        log.Err(err).Msg("AddCurrencyDifferentCode()")
    } else {
        log.Debug().Msgf("Expected conversion %f, got %f", bExpectedConversion, b.Conversion)
    }
    // End test rates

    // Perform test
    cDate := time.Date(2025, 5, 1, 0, 0, 0, 0, time.UTC)
    ccy, err := c.AddCurrency(a, b, cDate)
    if err != nil {
        log.Err(err).Msg("Failed to add currency")
    }

    log.Trace().Interface("Currency", ccy).Msg("Added Currency")
    log.Info().Msg("Added Currency Successfully")

    // assertions
    expectedAmount := 102.00
    expectedCode := c.CurrencyUnit("USD")
    expectedConversion := 102.00
    expectedRate := 1.00

    if ccy.Amount != expectedAmount {
        err = fmt.Errorf("expected amount %f, got %f", expectedAmount, ccy.Amount)
        log.Err(err).Msg("AddCurrencyDifferentCode()")
    } else {
        log.Debug().Msgf("Expected amount %f, got %f", expectedAmount, ccy.Amount)
    }

    if ccy.CurrencyCode != expectedCode {
        err = fmt.Errorf("expected code %s, got %s", expectedCode, ccy.CurrencyCode)
        log.Err(err).Msg("AddCurrencyDifferentCode()")
    } else {
        log.Debug().Msgf("Expected code %s, got %s", expectedCode, ccy.CurrencyCode)
    }

    if ccy.Conversion != expectedConversion {
        err = fmt.Errorf("expected conversion %f, got %f", expectedConversion, ccy.Conversion)
        log.Err(err).Msg("AddCurrencyDifferentCode()")
    } else {
        log.Debug().Msgf("Expected conversion %f, got %f", expectedConversion, ccy.Conversion)
    }

    if ccy.ExchangeRate != expectedRate {
        err = fmt.Errorf("expected rate %f, got %f", expectedRate, ccy.ExchangeRate)
        log.Err(err).Msg("AddCurrencyDifferentCode()")
    } else {
        log.Debug().Msgf("Expected rate %f, got %f", expectedRate, ccy.ExchangeRate)
    }
}

// TestSubCurrencySameCode is a test function that checks the subtraction of two currencies with the same code.
//
// It prepares the test data by creating two currencies with the same code and different amounts.
// It then performs the subtraction of the two currencies and checks the expected values.
// The expected values are the subtraction of the amounts, the same code, the conversion rate, and the exchange rate.
//
// It uses the auth.CreateUser function to create a user.
// It uses the c.CreateCurrency function to create the currencies.
// It uses the c.SubCurrency function to subtract the currencies.
//
// It logs the progress of the test using the log.Info, log.Debug, and log.Err functions.
//
// It returns nothing.
func TestSubCurrencySameCode(){
    log.Info().Msg("Testing SubCurrencySameCode()")

    // data prep
    owner, _ := auth.CreateUser("owner", "owner@admin.com", "test123")
    ownerId := fmt.Sprint(owner.ID)
    
    aAmount := "100.00"
    aCode := "ARS"
    aDate := "2024-01-01"

    a, err := c.CreateCurrency(aCode, aAmount, aDate, ownerId)
    if err != nil {
        log.Err(err).Msg("Failed to create currency")
    }
    log.Trace().Interface("Currency", a).Msg("Created A Currency")

    bAmount := "25.00"
    bCode := "ARS"
    bDate := "2011-01-03"

    b, err := c.CreateCurrency(bCode, bAmount, bDate, ownerId)
    if err != nil {
        log.Err(err).Msg("Failed to create currency")
    }
    log.Trace().Interface("Currency", b).Msg("Created B Currency")

    // Perform test
    cDate := time.Date(2024, 5, 1, 0, 0, 0, 0, time.UTC)
    ccy, err := c.SubCurrency(a, b, cDate)
    if err != nil {
        log.Err(err).Msg("Failed to sub currency")
    }

    log.Trace().Interface("Currency", ccy).Msg("Substracted Currency")
    log.Info().Msg("Substracted Currency Successfully")

    // assertions
    expectedAmount := 75.00
    expectedCode := c.CurrencyUnit("ARS")
    expectedRate := 1040.00
    expectedConversion := fmt.Sprintf("%.2f", expectedAmount / expectedRate) 

    if ccy.Amount != expectedAmount {
        err = fmt.Errorf("expected amount %f, got %f", expectedAmount, ccy.Amount)
        log.Err(err).Msg("TestSubCurrencySameCode()")
    } else {
        log.Debug().Msgf("Expected amount %f, got %f", expectedAmount, ccy.Amount)
    }

    if ccy.CurrencyCode != expectedCode {
        err = fmt.Errorf("expected code %s, got %s", expectedCode, ccy.CurrencyCode)
        log.Err(err).Msg("TestSubCurrencySameCode()")
    } else {
        log.Debug().Msgf("Expected code %s, got %s", expectedCode, ccy.CurrencyCode)
    }
    
    if fmt.Sprintf("%.2f", ccy.Conversion) != expectedConversion {
        err = fmt.Errorf("expected conversion %f to be equal to %f", ccy.Conversion, expectedAmount / expectedRate)
        log.Err(err).Msg("TestSubCurrencySameCode()")
    } else {
        log.Debug().Msgf("Expected conversion %f to be equal to %f", ccy.Conversion, expectedAmount / expectedRate)
    }

    if ccy.ExchangeRate != expectedRate {
        err = fmt.Errorf("expected rate %f, got %f", expectedRate, ccy.ExchangeRate)
        log.Err(err).Msg("TestSubCurrencySameCode()")
    } else {
        log.Debug().Msgf("Expected rate %f, got %f", expectedRate, ccy.ExchangeRate)
    }
}

// TestSubCurrencyDifferentCode is a test function that tests the SubCurrency function when the currencies have different codes.
//
// It performs the following steps:
// - Creates two currencies with different codes and amounts.
// - Tests the exchange rates of the created currencies.
// - Adds the currencies together using the SubCurrency function.
// - Performs assertions on the resulting currency's amount, code, conversion, and exchange rate.
//
// It uses the auth.CreateUser function to create a user.
// It uses the c.CreateCurrency function to create the currencies.
// It uses the c.SubCurrency function to subtract the currencies.
//
// It logs the progress of the test using the log.Info, log.Debug, and log.Err functions.
//
// It returns nothing.
func TestSubCurrencyDifferentCode(){
    log.Info().Msg("Testing SubCurrencyDifferentCode()")

    // data prep
    owner, _ := auth.CreateUser("owner", "owner@admin.com", "test123")
    ownerId := fmt.Sprint(owner.ID)
    
    aAmount := "20000.00"
    aCode := "ARS"
    aDate := "2024-01-01"

    a, err := c.CreateCurrency(aCode, aAmount, aDate, ownerId)
    if err != nil {
        log.Err(err).Msg("Failed to create currency")
    }
    log.Trace().Interface("Currency", a).Msg("Created A Currency")

    // Test A Rates
    aExpectedRate := 1000.00
    aExpectedConversion := 20.00
    aExpectedConversionStr := fmt.Sprintf("%.2f", aExpectedConversion)

    if a.ExchangeRate != aExpectedRate {
        err = fmt.Errorf("expected rate %f, got %f", aExpectedRate, a.ExchangeRate)
        log.Err(err).Msg("TestSubCurrencyDifferentCode()")
    } else {
        log.Debug().Msgf("Expected rate %f, got %f", aExpectedRate, a.ExchangeRate)
    }

    if fmt.Sprintf("%.2f", a.Conversion) != aExpectedConversionStr {
        err = fmt.Errorf("expected conversion %f to be equal to %f", aExpectedConversion, a.Conversion)
        log.Err(err).Msg("TestSubCurrencyDifferentCode()")
    } else {
        log.Debug().Msgf("Expected conversion %f to be equal to %f", aExpectedConversion, a.Conversion)
    }
    // End test rates

    bAmount := "4.00"
    bCode := "USD"
    bDate := "2011-01-03"

    b, err := c.CreateCurrency(bCode, bAmount, bDate, ownerId)
    if err != nil {
        log.Err(err).Msg("Failed to create currency")
    }
    log.Trace().Interface("Currency", b).Msg("Created B Currency")

    // Perform test
    cDate := time.Date(2024, 5, 1, 0, 0, 0, 0, time.UTC)
    ccy, err := c.SubCurrency(a, b, cDate)
    if err != nil {
        log.Err(err).Msg("Failed to sub currency")
    }

    log.Trace().Interface("Currency", ccy).Msg("Substracted Currency")
    log.Info().Msg("Substracted Currency Successfully")

    // assertions
    expectedAmount := 16640.00
    expectedCode := c.CurrencyUnit("ARS")
    expectedRate := 1040.00
    expectedConversion := fmt.Sprintf("%.2f", expectedAmount / expectedRate)

    if ccy.Amount != expectedAmount {
        err = fmt.Errorf("expected amount %f, got %f", expectedAmount, ccy.Amount)
        log.Err(err).Msg("TestSubCurrencyDifferentCode()")
    } else {
        log.Debug().Msgf("Expected amount %f, got %f", expectedAmount, ccy.Amount)
    }

    if ccy.CurrencyCode != expectedCode {
        err = fmt.Errorf("expected code %s, got %s", expectedCode, ccy.CurrencyCode)
        log.Err(err).Msg("TestSubCurrencyDifferentCode()")
    } else {
        log.Debug().Msgf("Expected code %s, got %s", expectedCode, ccy.CurrencyCode)
    }
    
    if fmt.Sprintf("%.2f", ccy.Conversion) != expectedConversion {
        err = fmt.Errorf("expected conversion %f to be equal to %f", ccy.Conversion, expectedAmount / expectedRate)
        log.Err(err).Msg("TestSubCurrencyDifferentCode()")
    } else {
        log.Debug().Msgf("Expected conversion %f to be equal to %f", ccy.Conversion, expectedAmount / expectedRate)
    }

    if ccy.ExchangeRate != expectedRate {
        err = fmt.Errorf("expected rate %f, got %f", expectedRate, ccy.ExchangeRate)
        log.Err(err).Msg("TestSubCurrencyDifferentCode()")
    } else {
        log.Debug().Msgf("Expected rate %f, got %f", expectedRate, ccy.ExchangeRate)
    }
}

