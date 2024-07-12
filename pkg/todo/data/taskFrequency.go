package data

import (
	"github.com/frangdelsolar/todo_cli/pkg/todo/models"
)

func CreateTaskFrequency(
	in_frequency_name string,
	in_frequency_day string,
	in_frequency_month string,
	in_frequency_day_of_week string,
) (*models.TaskFrequency, error) {

	tf, err := models.NewTaskFrequency(
		in_frequency_name,
		in_frequency_day,
		in_frequency_month,
		in_frequency_day_of_week,
	)

	if err != nil {
		log.Err(err).Msg("Error creating new Task Frequency")
		return nil, err
	}
	db.Create(&tf)
	return tf, nil
}
