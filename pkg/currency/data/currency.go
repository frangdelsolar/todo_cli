package data

import (
	"fmt"

	m "github.com/frangdelsolar/todo_cli/pkg/currency/models"
)

// CreateCurrency creates a new Currency object in the database.
//
// Parameters:
// - currencyCode: the code of the currency (e.g. "USD", "EUR").
// - amount: the amount of the currency.
// - exchangeDate: the date of the exchange rate.
//
// Returns:
// - *Currency: the created Currency object.
// - error: an error if the creation failed.
func CreateCurrency(currencyCode string, amount string, exchangeDate string) (*m.Currency, error) {
	var c *m.Currency

	c, err := m.NewCurrency(currencyCode, amount, exchangeDate)
	if err != nil {
		return c, err
	}

	db.Create(&c)

	return c, nil
}

// GetCurrencyById retrieves a Currency object from the database by its ID.
//
// Parameters:
// - id: the ID of the Currency to retrieve.
//
// Returns:
// - Currency: the retrieved Currency object, or an empty Currency object if not found.
// - error: an error if the Currency retrieval fails.
func GetCurrencyById(id string) (m.Currency, error) {
	var c m.Currency

	db.First(&c, "id = ?", id)
	if c == (m.Currency{}) {
		return c, fmt.Errorf("currency with ID %s not found", fmt.Sprint(id))
	}
	return c, nil
}

// GetAllCurrencies retrieves all the currencies from the database.
//
// Returns:
// - []Currency: a slice of Currency objects representing all the currencies.
func GetAllCurrencies() []m.Currency {
	var cs []m.Currency

	db.Find(&cs)

	if len(cs) == 0 {
		log.Warn().Msg("No currencies found")
		return cs
	}

	return cs
}

// DeleteCurrency deletes a currency from the database by its ID.
//
// Parameters:
// - id: the ID of the currency to delete.
//
// Returns:
// - error: an error if the currency retrieval or deletion fails.
func DeleteCurrency(id string) error {
	c, err := GetCurrencyById(id)
	if err != nil {
		log.Err(err).Msg("Error getting currency")
		return err
	}

	db.Delete(&c)

	return nil
}
