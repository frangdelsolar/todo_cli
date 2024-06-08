package models

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type FrequencyType string

const (
	Daily   FrequencyType = "daily"
	Weekly  FrequencyType = "weekly"
	Monthly FrequencyType = "monthly"
	Yearly  FrequencyType = "yearly"
)

type TaskGoalCategory string

const (
	Todo TaskGoalCategory = "todo"
)

// TaskGoal represents an task goal associated with a task.
//
// Fields:
// - ID: the ID of the TaskGoal.
// - TaskID: the ID of the task associated with the TaskGoal.
// - StartDate: the start date of the TaskGoal.
// - EndDate: the end date of the TaskGoal.
// - CreatedAt: the timestamp when the TaskGoal was created.
type TaskGoal struct {
	gorm.Model
	ID        string             `json:"id" gorm:"primaryKey"`
	TaskID    string             `json:"taskId"`
	Task      *Task            `json:"task" gorm:"foreignKey:TaskID"`
	StartDate time.Time        `json:"startDate"`
	EndDate   time.Time        `json:"endDate" omitempty:"true"`
	Frequency FrequencyType    `json:"frequency"`
	Category  TaskGoalCategory `json:"category"`
	CreatedAt time.Time        `json:"createdAt"`
	UpdatedAt time.Time        `json:"updatedAt"`
}

// String returns a string representation of the TaskGoal.
//
// Returns:
// - string: a string representation of the TaskGoal.
func (e *TaskGoal) String() string {
	return fmt.Sprintf("TaskGoal %d\nTask ID: %d\nStart Date: %s\nEnd Date: %s\nCreated At: %s\n\n", e.ID, e.TaskID, e.StartDate, e.EndDate, e.CreatedAt)
}

// NewTaskGoal creates a new TaskGoal with the given task ID, start date, and end date.
//
// Parameters:
// - in_taskID: the ID of the task associated with the TaskGoal.
// - in_startDate: the start date of the TaskGoal.
// - in_endDate: the end date of the TaskGoal.
//
// Returns:
// - *TaskGoal: the newly created TaskGoal.
// - error: an error if there was a problem parsing the start or end date, or if the start date is after the end date.
func NewTaskGoal(in_taskId string, in_startDate string, in_endDate string, in_frequency string, in_category string) (*TaskGoal, error) {
	var output *TaskGoal
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

	frequency := FrequencyType(in_frequency)

	err = CategoryValidator(in_category)
	if err != nil {
		log.Warn().Msg("Invalid category. Defaulting to todo.")
		in_category = string(Todo)
	}

	category := TaskGoalCategory(in_category)

	output = &TaskGoal{
		TaskID:    in_taskId,
		StartDate: sd,
		Frequency: frequency,
		Category:  category,
	}

	if in_endDate != "" {
		output.EndDate = ed
	}

	return output, nil
}

// Update updates the start and end dates of an TaskGoal.
//
// Parameters:
// - in_startDate: the new start date in string format.
// - in_endDate: the new end date in string format.
//
// Returns:
// - error: an error if there was a problem parsing the start or end date, or if the start date is after the end date.
func (e *TaskGoal) Update(in_startDate string, in_endDate string) error {
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

// TaskGoalIDValidator validates a task ID.
//
// Parameters:
// - id: the ID of the task goal to validate.
//
// Returns:
// - error: an error if the task ID is 0, otherwise nil.
func TaskGoalIDValidator(id string) error {

	var err error
	if id == "" {
		err = errors.New("task goal ID cannot be empty")
		return err
	}

	_, err = strconv.ParseUint(id, 10, 64)
	err = fmt.Errorf("invalid task goal ID: %w", err)
	return err

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

// CategoryValidator validates the given category string.
//
// It checks if the category is equal to the string representation of the Todo constant.
// If the category is valid, it returns nil. Otherwise, it returns an error with the message "invalid category".
//
// Parameters:
// - category: a string representing the category to be validated.
//
// Returns:
// - error: an error if the category is invalid, otherwise nil.
func CategoryValidator(category string) error {
	if category == string(Todo) {
		return nil
	}
	return errors.New("invalid category")
}
