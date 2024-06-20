package contractor

import (
	"time"

	c "github.com/frangdelsolar/todo_cli/pkg/currency"
	data "github.com/frangdelsolar/todo_cli/pkg/data/models"
)

type AgreementType string

const (
    AgreementTypeOneOff AgreementType = "one-off"
    AgreementTypeRecurring AgreementType = "recurring"
    AgreementTypeFixedAmount AgreementType = "fixed-amount"
    AgreementTypeInstallment AgreementType = "installment" 
)

type FrequencyType string

const (
    FrequencyTypeDaily FrequencyType = "daily"
    FrequencyTypeWeekly FrequencyType = "weekly"
    FrequencyTypeMonthly FrequencyType = "monthly"
    FrequencyTypeYearly FrequencyType = "yearly"
)
type Frequency struct {
    data.SystemData
    Type FrequencyType `gorm:"not null"`
    Day int
    Month int
    DayOfWeek int
}

type EffectivePeriod struct {
    data.SystemData
    StartDate time.Time
    EndDate time.Time
}

type Agreement struct {
    data.SystemData
    Contractor *Contractor
    ContractorId string `gorm:"not null"`
    Type AgreementType `gorm:"not null"`
    Concept string `gorm:"not null"`
    Frequency *Frequency
    FrequencyId string `gorm:"not null"` 
    EffectivePeriod *EffectivePeriod
    EffectivePeriodId string `gorm:"not null"`
    Repetitions int // only if type is installement
    FixedAmount *c.Currency // only if type is fixed-amount
    CurrencyCode  *c.CurrencyUnit `gorm:"not null"`
}




