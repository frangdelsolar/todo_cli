package currency

import (
	"fmt"
	"time"

	"github.com/frangdelsolar/todo_cli/pkg/auth"
	d "github.com/frangdelsolar/todo_cli/pkg/data/models"
)

type Account struct {
	d.SystemData
	Name        string    `json:"name" gorm:"not null"`
	TotalID    uint      `json:"totalId" gorm:"not null"`
	Total   *Currency    `json:"total" gorm:"foreignKey:TotalID"`
	DefaultAccount bool `json:"defaultAccount"`
	CurrencyCode CurrencyUnit `json:"currencyCode" gorm:"not null"` 
}

func (a *Account) String() string {
	return fmt.Sprintf("%s (%s)", a.Name, a.CurrencyCode)
}

// UpdateName updates the name of the account.
//
// Parameters:
// - name: the new name for the account.
//
// Returns:
// - error: an error if there was a problem updating the account name.
func (a *Account) UpdateName(name string, requestedBy *auth.User) error {
	if err := AccountNameValidator(name); err != nil {
		log.Err(err).Msg("Error validating account name")
		return err
	}
	a.Name = name
	a.SystemData.UpdatedBy = requestedBy
	a.SystemData.UpdatedAt = time.Now()
	return nil
}

// NewAccount creates a new account with the given name, total Currency, and defaultAccount flag.
//
// Parameters:
// - name: the name of the account.
// - total: the total Currency of the account.
// - defaultAccount: a boolean indicating if it is the default account.
//
// Returns:
// - *Account: the newly created Account.
// - error: an error if there was a problem during creation.
func NewAccount (name string, total *Currency, defaultAccount bool, requestedBy *auth.User) (*Account, error) {

	if err := AccountNameValidator(name); err != nil {
		log.Err(err).Msg("Error validating account name")
		return nil, err
	}

	return &Account{
		Name: name,
		CurrencyCode: total.CurrencyCode,
		Total: total,
		DefaultAccount: defaultAccount,
		SystemData: d.SystemData{
			CreatedBy: requestedBy,
			UpdatedBy: requestedBy,
		},
	}, nil
}

// AccountNameValidator validates the name of an account.
//
// Parameters:
// - name: the name of the account to be validated.
//
// Returns:
// - error: an error if the account name is empty, otherwise nil.
func AccountNameValidator(name string) error {
	if name == "" {
		return fmt.Errorf("account name cannot be empty")
	}
	return nil
}
