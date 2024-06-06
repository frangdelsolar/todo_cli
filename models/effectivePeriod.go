package models

import (
	"fmt"
	"time"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type Frequency string

const (
	Daily   Frequency = "daily"
	Weekly  Frequency = "weekly"
	Monthly Frequency = "monthly"
	Yearly  Frequency = "yearly"
)

// ValidFrequency checks if the given frequency string is valid.
//
// Parameters:
// - freq: the frequency string to be checked.
//
// Returns:
// - bool: true if the frequency string is valid, false otherwise.
func ValidFrequency(freq string) bool {
	switch freq {
	case string(Daily), string(Weekly), string(Monthly), string(Yearly):
		return true
	default:
		return false
	}
}

// EffectivePeriod represents an effective period associated with a task.
//
// Fields:
// - ID: the ID of the EffectivePeriod.
// - TaskID: the ID of the task associated with the EffectivePeriod.
// - StartDate: the start date of the EffectivePeriod.
// - EndDate: the end date of the EffectivePeriod.
// - CreatedAt: the timestamp when the EffectivePeriod was created.
type EffectivePeriod struct {
	gorm.Model
	ID        uint      `json:"id" gorm:"primaryKey"`
	TaskID    uint      `json:"taskId"`
	Task      *Task     `json:"task" gorm:"foreignKey:TaskID"`
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate" omitempty:"true"`
	Frequency Frequency `json:"frequency"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// String returns a string representation of the EffectivePeriod.
//
// Returns:
// - string: a string representation of the EffectivePeriod.
func (e *EffectivePeriod) String() string {
	return fmt.Sprintf("EffectivePeriod %s\nTask ID: %s\nStart Date: %s\nEnd Date: %s\nCreated At: %s\n\n", e.ID, e.TaskID, e.StartDate, e.EndDate, e.CreatedAt)
}

// NewEffectivePeriod creates a new EffectivePeriod with the given task ID, start date, and end date.
//
// Parameters:
// - in_taskID: the ID of the task associated with the EffectivePeriod.
// - in_startDate: the start date of the EffectivePeriod.
// - in_endDate: the end date of the EffectivePeriod.
//
// Returns:
// - *EffectivePeriod: the newly created EffectivePeriod.
// - error: an error if there was a problem parsing the start or end date, or if the start date is after the end date.
func NewEffectivePeriod(in_taskId uint, in_startDate string, in_endDate string, in_frequency string) (*EffectivePeriod, error) {
	var output *EffectivePeriod
	var err error

	if in_taskId == 0 {
		log.Err(err).Msg("Task ID cannot be empty")
		return output, err
	}

	// TODO: Validate in_taskId actually exists
	now := time.Now()

	sd, err := time.Parse(time.DateOnly, in_startDate)
	if err != nil {
		log.Warn().Msgf("Error parsing start date: %s", err)
		sd = now
	}

	var ed time.Time
	if in_endDate != "" {
		ed, err = time.Parse(time.DateOnly, in_endDate)
		if err != nil {
			log.Err(err).Msg("Error parsing end date")
			return output, err
		}
		if sd.After(ed) {
			log.Err(err).Msg("Start date is after end date")
			return output, err
		}
	}

	if !ValidFrequency(in_frequency) {
		log.Warn().Msg("Invalid frequency. Defaulting to monthly.")
		in_frequency = string(Monthly)
	}

	frequency := Frequency(in_frequency)

	output = &EffectivePeriod{
		TaskID:    uint(in_taskId),
		StartDate: sd,
		Frequency: frequency,
	}

	if in_endDate != "" {
		output.EndDate = ed
	}

	return output, err
}

// Update updates the start and end dates of an EffectivePeriod.
//
// Parameters:
// - in_startDate: the new start date in string format.
// - in_endDate: the new end date in string format.
//
// Returns:
// - error: an error if there was a problem parsing the start or end date, or if the start date is after the end date.
func (e *EffectivePeriod) Update(in_startDate string, in_endDate string) error {
	var err error

	sd, err := time.Parse(time.DateOnly, in_startDate)
	if err != nil {
		log.Err(err).Msg("Error parsing start date")
		return err
	}

	var ed time.Time
	if in_endDate != "" {
		ed, err = time.Parse(time.DateOnly, in_endDate)
		if err != nil {
			log.Err(err).Msg("Error parsing end date")
			return err
		}
		if sd.After(ed) {
			log.Err(err).Msg("Start date is after end date")
			return err
		}
	}

	e.StartDate = sd
	e.EndDate = ed

	return err
}
