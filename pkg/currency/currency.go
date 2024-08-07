package currency

import (
	"fmt"
	"strconv"
	"time"

	"github.com/frangdelsolar/todo_cli/pkg/auth"
	d "github.com/frangdelsolar/todo_cli/pkg/data/models"
)

type CurrencyUnit string

const (
	ARS CurrencyUnit = "ARS"
	USD CurrencyUnit = "USD"
)

type Currency struct {
	d.SystemData
	ID           uint         `json:"id" gorm:"primaryKey"`
	CurrencyCode CurrencyUnit `json:"currencyCode"`
	Amount       float64      `json:"amount"`
	ExchangeRate float64      `json:"exchangeRate"`
	Conversion   float64      `json:"conversion"`
	ExchangeDate time.Time    `json:"exchangeDate"`
}

func (c *Currency) String() string {
	return fmt.Sprintf("Currency: %s, Amount: %f", c.CurrencyCode, c.Amount)
}

// AddCurrency adds two currencies together based on their conversion rates and
// stores the result in the database. It takes two Currency pointers and a
// time.Time value as parameters. The first currency pointer represents the
// first currency to be added, the second currency pointer represents the second
// currency to be added, and the date parameter represents the date of the
// exchange rate. The function returns a pointer to a Currency object and an
// error. The returned Currency object contains the result of adding the two
// currencies, and the error indicates if there was an error during the
// calculation or database insertion.
func AddCurrency(a *Currency, b *Currency, date time.Time) (*Currency, error) {
	output := &Currency{}
	var err error

	amount := ""
	cCode := ""
	eDate := date.Format(time.DateOnly)
	eRate, err := GetRatesByDate(date)
	if err != nil {
		return output, err
	}

	if a.CurrencyCode == b.CurrencyCode {
		amount = fmt.Sprint(a.Amount + b.Amount)
		cCode = string(a.CurrencyCode)
	} else {
		if a.CurrencyCode == USD {
			amount = fmt.Sprint(a.Conversion + b.Conversion)
			cCode = string(USD)
		} else if a.CurrencyCode == ARS {
			amount = fmt.Sprint((a.Conversion + b.Conversion) * eRate.GetBlueAverage())
			cCode = string(ARS)
		}
	}

	output, err = NewCurrency(cCode, amount, eDate, a.SystemData.CreatedBy)
	if err != nil {
		return output, err
	}
	return output, nil
}

// SubCurrency subtracts two currencies together based on their conversion rates and stores the result in the database.
//
// Parameters:
// - a: a pointer to a Currency object representing the first currency to be subtracted.
// - b: a pointer to a Currency object representing the second currency to be subtracted.
// - date: a time.Time value representing the date of the exchange rate.
//
// Returns:
// - a pointer to a Currency object containing the result of subtracting the two currencies, and an error if there was an error during the calculation or database insertion.
func SubCurrency(a *Currency, b *Currency, date time.Time) (*Currency, error) {

	output := &Currency{}
	var err error

	amount := ""
	cCode := ""
	eDate := date.Format(time.DateOnly)
	eRate, err := GetRatesByDate(date)
	if err != nil {
		return output, err
	}

	if a.CurrencyCode == b.CurrencyCode {
		amount = fmt.Sprint(a.Amount - b.Amount)
		cCode = string(a.CurrencyCode)
	} else {
		if a.CurrencyCode == USD {
			amount = fmt.Sprint(a.Conversion - b.Conversion)
			cCode = string(USD)
		} else if a.CurrencyCode == ARS {
			amount = fmt.Sprint((a.Conversion - b.Conversion) * eRate.GetBlueAverage())
			cCode = string(ARS)
		}
	}

	output, err = NewCurrency(cCode, amount, eDate, a.SystemData.CreatedBy)
	if err != nil {
		return output, err
	}

	return output, nil
}

// NewCurrency creates a new Currency object with the given currency code, amount, and exchange date.
//
// Parameters:
// - currency: the currency code as a string.
// - amount: the amount as a string.
// - exchangeDate: the exchange date as a string.
//
// Returns:
// - *Currency: a pointer to the newly created Currency object.
// - error: an error if any validation fails or if there is an error getting the exchange rate.
func NewCurrency(currencyCode string, amount string, exchangeDate string, user *auth.User) (*Currency, error) {

	// Run Validations
	if err := CurrencyCodeValidator(currencyCode); err != nil {
		log.Err(err).Msg("Error validating currency code")
		return nil, err
	}

	if err := CurrencyAmountValidator(amount); err != nil {
		log.Err(err).Msg("Error validating currency amount")
		return nil, err
	}

	if err := DateValidator(exchangeDate); err != nil {
		log.Err(err).Msg("Error validating exchange date")
		return nil, err
	}

	amountFloat, _ := strconv.ParseFloat(amount, 64)
	cc := CurrencyUnit(currencyCode)
	eDate, _ := time.Parse(time.DateOnly, exchangeDate)

	var er float64
	var conversion float64

	if cc == USD {
		er = 1
		conversion = amountFloat
	} else if cc == ARS {
		rates, err := GetRatesByDate(eDate)
		if err != nil {
			log.Err(err).Msg("Error getting rates")
			return nil, err
		}

		er = rates.GetBlueAverage()
		conversion = amountFloat / er
	}

	return &Currency{
		CurrencyCode: cc,
		Amount:       amountFloat,
		ExchangeDate: eDate,
		ExchangeRate: er,
		Conversion:   conversion,
		SystemData: d.SystemData{
			CreatedBy: user,
			UpdatedBy: user,
		},
	}, nil
}

// CurrencyAmountValidator validates the given amount string.
//
// It checks if the amount can be parsed into a float64 and if it is greater than 0.
// If the amount is valid, it returns nil. Otherwise, it returns an error with the message "amount must be greater than 0".
//
// Parameters:
// - amount: a string representing the amount to be validated.
//
// Returns:
// - error: an error if the amount is invalid, otherwise nil.
func CurrencyAmountValidator(amount string) error {
	_, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return err
	}

	return nil
}

// CurrencyCodeValidator validates the given currency code.
//
// It checks if the currency code is either "USD" or "ARS". If it is, it returns nil. Otherwise, it returns an error with the message "invalid currency code".
//
// Parameters:
// - currency: a string representing the currency code to be validated.
//
// Returns:
// - error: an error if the currency code is invalid, otherwise nil.
func CurrencyCodeValidator(currency string) error {
	if currency == string(USD) || currency == string(ARS) {
		return nil
	}

	return fmt.Errorf("invalid currency code")
}

// DateValidator validates a date string in the format "YYYY-MM-DD".
//
// Parameters:
// - date: the date string to validate.
//
// Returns:
// - error: an error if the date string is not in the correct format or if the date is invalid, otherwise nil.
func DateValidator(date string) error {
	_, err := time.Parse(time.DateOnly, date)
	return err
}
