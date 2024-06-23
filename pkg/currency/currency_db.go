package currency

import (
	"fmt"

	"github.com/frangdelsolar/todo_cli/pkg/auth"
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
func CreateCurrency(currencyCode string, amount string, exchangeDate string, requestedBy string) (*Currency, error) {
	var c *Currency

	user, err := auth.GetUserById(requestedBy)
	if err != nil {
		return c, err
	} 

	c, err = NewCurrency(currencyCode, amount, exchangeDate, user)
	if err != nil {
		return c, err
	}

	db.Create(&c)

    log.Trace().Interface("currency", c).Msg("Created currency")
    log.Info().Msg("Created currency")

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
func GetCurrencyById(id string, requestedBy string) (Currency, error) {
	var c Currency

	db.
        First(&c, "id = ?", id).
        Where("created_by = ?", requestedBy)
        
	if c == (Currency{}) {
		return c, fmt.Errorf("currency with ID %s not found", fmt.Sprint(id))
	}
	return c, nil
}

// GetAllCurrencies retrieves all the currencies from the database.
//
// Returns:
// - []Currency: a slice of Currency objects representing all the currencies.
func GetAllCurrencies(requestedBy string) []Currency {
	var cs []Currency

	db.
        Find(&cs).
        Where("created_by = ?", requestedBy)

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
func DeleteCurrency(id string, requestedBy string) error {
	c, err := GetCurrencyById(id, requestedBy)
	if err != nil {
		log.Err(err).Msg("Error getting currency")
		return err
	}

	db.Delete(&c)

    log.Info().Msgf("Deleted currency with ID %s", id)

	return nil
}


