package models

import (
	"errors"
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
	return fmt.Sprintf("EffectivePeriod %d\nTask ID: %d\nStart Date: %s\nEnd Date: %s\nCreated At: %s\n\n", e.ID, e.TaskID, e.StartDate, e.EndDate, e.CreatedAt)
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

	// Run Validations
	err = TaskIDValidator(in_taskId)
	if err != nil {
		log.Err(err).Msg("Error validating task ID")
		return output, err
	}

	now := time.Now()

	var sd time.Time
	err = DateValidator(in_startDate)
	if err != nil {
		log.Warn().Msg("Error validating start date. Defaulting to today.")
		sd = now
	} else {
		sd, _ = time.Parse(time.DateOnly, in_startDate)
	}

	var ed time.Time
	if in_endDate != "" {
		err = DateValidator(in_endDate)
		if err != nil {
			log.Err(err).Msg("Error validating end date")
			return output, err
		}
		ed, _ = time.Parse(time.DateOnly, in_endDate)

		if sd.After(ed) {
			log.Err(err).Msg("Start date is after end date")
			return output, err
		}
	}

	err = FrequencyValidator(in_frequency)
	if err != nil {
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

	return output, nil
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

	now := time.Now()

	var sd time.Time
	err = DateValidator(in_startDate)
	if err != nil {
		log.Warn().Msg("Error validating start date. Defaulting to today.")
		sd = now
	} else {
		sd, _ = time.Parse(time.DateOnly, in_startDate)
	}

	var ed time.Time
	if in_endDate != "" {
		err = DateValidator(in_endDate)
		if err != nil {
			log.Err(err).Msg("Error validating end date")
			return err
		}
		ed, _ = time.Parse(time.DateOnly, in_endDate)

		if sd.After(ed) {
			log.Err(err).Msg("Start date is after end date")
			return err
		}
	}

	e.StartDate = sd
	e.EndDate = ed

	return err
}

// TaskIDValidator validates a task ID.
//
// Parameters:
// - id: the ID of the task to validate.
//
// Returns:
// - error: an error if the task ID is 0, otherwise nil.
func TaskIDValidator(id uint) error {

	if id == 0 {
		return errors.New("task ID cannot be 0")
	}
	return nil
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

// FrequencyValidator validates the given frequency string.
//
// It checks if the frequency is one of the following: Daily, Weekly, Monthly, Yearly.
// If the frequency is not valid, it returns an error with the message "invalid frequency".
// If the frequency is valid, it returns nil.
//
// Parameters:
// - frequency: a string representing the frequency to be validated.
//
// Returns:
// - error: an error if the frequency is invalid, otherwise nil.
func FrequencyValidator(frequency string) error {
	if frequency == string(Daily) || frequency == string(Weekly) || frequency == string(Monthly) || frequency == string(Yearly) {
		return nil
	}
	return errors.New("invalid frequency")
}
