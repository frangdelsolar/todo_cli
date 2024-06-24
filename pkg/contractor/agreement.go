package contractor

import (
	"errors"
	"strconv"
	"time"

	c "github.com/frangdelsolar/todo_cli/pkg/currency"
	data "github.com/frangdelsolar/todo_cli/pkg/data/models"
)

type AgreementType string

const (
	AgreementTypeOneOff      AgreementType = "one-off"
	AgreementTypeRecurring   AgreementType = "recurring"
	AgreementTypeFixedAmount AgreementType = "fixed-amount"
)

type EffectivePeriod struct {
	data.SystemData
	StartDate time.Time
	EndDate   time.Time
}

type Agreement struct {
	data.SystemData
	Contractor        *Contractor `gorm:"foreignKey:ContractorId"`
	ContractorId      string      `gorm:"not null"`
	Type              AgreementType
	Concept           string
	Frequency         *Frequency       `gorm:"foreignKey:FrequencyId"`
	FrequencyId       string           `gorm:"not null"`
	EffectivePeriod   *EffectivePeriod `gorm:"foreignKey:EffectivePeriodId"`
	EffectivePeriodId string           `gorm:"not null"`
	Installments      int              // only if type is fixed-amount
	FixedAmount       *c.Currency      `gorm:"foreignKey:FixedAmountId"` // only if type is fixed-amount
	FixedAmountId     string           `gorm:"not null"`
	CurrencyCode      *c.CurrencyUnit  `gorm:"not null"`
}

func NewAgreement() (*Agreement, error) {
	return nil, nil
}

// ValidateConcept validates the concept input.
//
// Parameters:
// - concept: the concept to validate.
//
// Returns:
// - error: an error if the concept is empty, otherwise nil.
func ValidateConcept(concept string) error {
	if concept == "" {
		return errors.New("concept cannot be empty")
	}
	return nil
}

// ValidateID validates the ID input.
//
// Parameters:
// - id: the ID to validate.
// Return type:
// - error: an error if the ID is empty, 0, or not an integer.
func ValidateID(id string) error {
	if id == "" {
		return errors.New("ID cannot be empty")
	}

	if id == "0" {
		return errors.New("ID cannot be 0")
	}

	_, err := strconv.Atoi(id)
	if err != nil {
		return errors.New("ID must be an integer")
	}

	return nil
}

// ValidateAgreementType validates the agreement type.
//
// Parameters:
// - at: the agreement type as a string.
// Return type:
// - error: an error if the agreement type is invalid.
func ValidateAgreementType(at string) error {
	switch at {
	case string(AgreementTypeOneOff):
		return nil
	case string(AgreementTypeRecurring):
		return nil
	case string(AgreementTypeFixedAmount):
		return nil
	default:
		return errors.New("invalid agreement type")
	}
}
