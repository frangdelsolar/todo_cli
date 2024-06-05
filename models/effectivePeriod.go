package models

import (
	"fmt"
	"time"

	"github.com/gofrs/uuid"
	"github.com/rs/zerolog/log"
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
	ID        uuid.UUID `json:"id"`
	TaskID    uuid.UUID `json:"taskId"`
	StartDate string    `json:"startDate"`
	EndDate   string    `json:"endDate"`
	CreatedAt string    `json:"createdAt"`
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
func NewEffectivePeriod(in_taskID string, in_startDate string, in_endDate string) (*EffectivePeriod, error) {
	var output *EffectivePeriod
	var err error

	st, err := time.Parse(time.DateOnly, in_startDate)
	if err != nil {
		log.Err(err).Msg("Error parsing start date")
		return output, err
	}
	formatedStartDate := st.Format(time.RFC3339)

	formatedEndDate := ""
	if in_endDate != "" {
		parsedEndDate, err := time.Parse(time.DateOnly, in_endDate)
		if err != nil {
			log.Err(err).Msg("Error parsing end date")
			return output, err
		}
		if st.After(parsedEndDate) {
			log.Err(err).Msg("Start date is after end date")
			return output, err
		}
		formatedEndDate = parsedEndDate.Format(time.RFC3339)
	}

	now := time.Now().Format(time.RFC3339)
	output= &EffectivePeriod{
		ID:        uuid.Must(uuid.NewV4()),
		TaskID:    uuid.Must(uuid.FromString(in_taskID)),
		StartDate: formatedStartDate,
		EndDate:   formatedEndDate ,
		CreatedAt: now,
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

	st, err := time.Parse(time.DateOnly, in_startDate)
	if err != nil {
		log.Err(err).Msg("Error parsing start date")
		return err
	}
	formatedStartDate := st.Format(time.RFC3339)

	formatedEndDate := ""
	if in_endDate != "" {
		parsedEndDate, err := time.Parse(time.DateOnly, in_endDate)
		if err != nil {
			log.Err(err).Msg("Error parsing end date")
			return err
		}
		if st.After(parsedEndDate) {
			log.Err(err).Msg("Start date is after end date")
			return err
		}
		formatedEndDate = parsedEndDate.Format(time.RFC3339)
	}

	e.StartDate = formatedStartDate
	e.EndDate = formatedEndDate

	return err
}







