package currency

import (
	"time"

	"github.com/frangdelsolar/todo_cli/pkg/data"
	"github.com/frangdelsolar/todo_cli/pkg/logger"
	"gorm.io/gorm"
)


var PKG_NAME = "Currency PKG"
var PKG_VERSION = "1.0.0"

var log *logger.Logger
var logLevel = "debug"

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

func InitCurrency() {
	log = logger.NewLogger(logLevel, PKG_NAME, PKG_VERSION)
	log.Info().Msgf("Running %s v%s", PKG_NAME, PKG_VERSION)

	// db := data.GetDB()
	db, err := data.InitDB("./data.db")
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
	}
	
	db.AutoMigrate(&Currency{})
	
}