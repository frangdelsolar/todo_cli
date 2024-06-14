package models

import (
	"fmt"
	"strconv"
	"time"

	"gorm.io/gorm"
)


type CurrencyUnit string

const (
	ARS CurrencyUnit = "ARS"
	USD CurrencyUnit = "USD"
)

type Currency struct {
	gorm.Model
	ID          uint      `json:"id" gorm:"primaryKey"`
	Currency    CurrencyUnit    `json:"currency"`
	Amount      float64   `json:"amount"`
	ExchangeRate float64   `json:"exchangeRate"`
	Conversion   float64   `json:"conversion"`
	ExchangeDate time.Time `json:"exchangeDate"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func (c *Currency) String() string {
	return fmt.Sprintf("Currency: %s, Amount: %f", c.Currency, c.Amount)
}

func AddCurrency(a *Currency, b *Currency, date time.Time) (*Currency, error) {
	output := &Currency{}

	if a.Currency == b.Currency {
		output, _ = NewCurrency(
			string(a.Currency),
			fmt.Sprintf("%f", a.Amount + b.Amount),
			date.Format(time.DateOnly),
		)
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
func NewCurrency (currencyCode string, amount string, exchangeDate string) (*Currency, error) {
	
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
		conversion = 1
	} else if cc == ARS {
		rates, err := GetRatesByDate(eDate)
		if err != nil {
			log.Err(err).Msg("Error getting rates")
			return nil, err
		}

		er = (rates.Blue.ValueBuy + rates.Blue.ValueSell) / 2
		conversion = amountFloat / er
	}

	return &Currency{
		Currency: cc,
		Amount: amountFloat,
		ExchangeDate: eDate,
		ExchangeRate: er,
		Conversion: conversion,
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
	a, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return err
	}

	if a < 0 {
		return fmt.Errorf("amount must be greater than 0")
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
	if currency == string(USD) || currency == string(ARS){
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
