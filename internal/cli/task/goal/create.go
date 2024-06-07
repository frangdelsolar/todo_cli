package goal

import (
	"todo_cli/data"
	"todo_cli/internal/cli/prompts"
	"todo_cli/pkg/prompt"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var CreateTaskGoalCmd = &cobra.Command{
	Use:   "add",
	Short: "Add Efective Periods",
	Run: func(cmd *cobra.Command, args []string) {
		taskId, err := prompts.SelectTaskFromAll()
		if err != nil {
			log.Err(err).Msg("Error selecting task")
			return
		}

		startDate := prompt.PromptGetInput(prompt.PromptContent{Label: "Start Date"})
		endDate := prompt.PromptGetInput(prompt.PromptContent{Label: "End Date"})
		frequency := prompt.PromptGetInput(prompt.PromptContent{Label: "Frequency"})
		category := prompt.PromptGetInput(prompt.PromptContent{Label: "Category"})

		ep, err := data.CreateTaskGoal(taskId, startDate, endDate, frequency, category)
		if err != nil {
			log.Err(err).Msg("Error creating Task Goal")
			return
		}
		log.Info().Interface("ep", ep).Msg("Task Goal added")
	},
}
