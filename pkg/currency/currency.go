package main

import (
	"time"

	"gorm.io/gorm"
)

type CurrencyUnit string

const (
	ARS CurrencyUnit = "ARS"
	USD CurrencyUnit = "USD"
	EUR CurrencyUnit = "EUR"
)

type Currency struct {
	gorm.Model
	ID          uint      `json:"id" gorm:"primaryKey"`
	Currency    CurrencyUnit    `json:"currency"`
	Amount      float64   `json:"amount"`
	ExchangeRate float64   `json:"exchangeRate"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

