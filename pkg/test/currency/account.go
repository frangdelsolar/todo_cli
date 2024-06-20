package currency

import (
	"fmt"
	"strconv"

	"github.com/frangdelsolar/todo_cli/pkg/auth"
	c "github.com/frangdelsolar/todo_cli/pkg/currency"
)


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


