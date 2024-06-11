package models

import (
	"time"

	"gorm.io/gorm"
)

type TransactionType string

const (
	Credit TransactionType = "Credit"
	Debit  TransactionType = "Debit"
)

type Transaction struct {
	gorm.Model
	ID          uint      `json:"id" gorm:"primaryKey"`
	Currency    Currency    `json:"currency"`
	DateOfTransaction time.Time `json:"dateOfTransaction"`
	TypeOfTrasaction TransactionType    `json:"typeOfTrasaction"`
	Account      *Account  `json:"account" gorm:"foreignKey:AccountID"`
	AccountID    uint      `json:"accountId"`
	Details     string    `json:"details"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}