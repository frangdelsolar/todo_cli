package goal

import (
	"fmt"

	"github.com/frangdelsolar/todo_cli/pkg/todo/cli/prompt"
	"github.com/frangdelsolar/todo_cli/pkg/todo/cli/prompts"
	"github.com/frangdelsolar/todo_cli/pkg/todo/data"
	"github.com/frangdelsolar/todo_cli/pkg/todo/models"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var CreateTaskGoalCmd = &cobra.Command{
	Use:   "add",
	Short: "Add Efective Periods",
	Run: func(cmd *cobra.Command, args []string) {

		var taskId string
		var err error

		if len(args) > 0 {
			taskId= args[0]
		} else {
			taskId, err = prompts.SelectTaskFromAll()
			if err != nil {
				log.Err(err).Msg("Error selecting task")
				return
			}
		}

		startDate := prompt.PromptGetInput(prompt.PromptContent{Label: "Start Date"})
		endDate := prompt.PromptGetInput(prompt.PromptContent{Label: "End Date"})
		frequency := prompt.PromptGetInput(prompt.PromptContent{Label: "Frequency"})
		if frequency == "" {
			frequency = string(models.Monthly)
		}

		day := "1"
		month := "1"
		dayOfWeek := "1"
		switch frequency {
			case string(models.Daily):
				break
			case string(models.Weekly):
				dayOfWeek = prompt.PromptGetInput(prompt.PromptContent{Label: "Day of Week"})
			case string(models.Monthly):
				day = prompt.PromptGetInput(prompt.PromptContent{Label: "Day of Month"})
			case string(models.Yearly):
				day = prompt.PromptGetInput(prompt.PromptContent{Label: "Day of Month"})
				month = prompt.PromptGetInput(prompt.PromptContent{Label: "Month"})
			default:
				log.Error().Msg("Invalid frequency")
				return
		}

		frequencyInstance, err := data.CreateTaskFrequency(
			frequency, 
			day, 
			month, 
			dayOfWeek,
		)
		if err != nil {
			log.Warn().Msgf("Error creating Task Frequency: %s", err.Error())
			frequencyInstance, _ = data.CreateTaskFrequency(
				string(models.Monthly), 
				"1", 
				"1", 
				"1",
			)
		}

		category := prompt.PromptGetInput(prompt.PromptContent{Label: "Category"})

		ep, err := data.CreateTaskGoal(
			taskId, 
			startDate, 
			endDate, 
			fmt.Sprint(frequencyInstance.ID), 
			category,
		)
		if err != nil {
			log.Err(err).Msg("Error creating Task Goal")
			return
		}
		log.Info().Interface("ep", ep).Msg("Task Goal added")
	},
}
