package period

import (
	"todo_cli/data"
	"todo_cli/internal/cli/prompts"
	"todo_cli/pkg/prompt"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var CreateEffectivePeriodCmd = &cobra.Command{
	Use:   "add",
	Short: "Add Efective Periods",
	Run: func(cmd *cobra.Command, args []string) {
		log.Info().Msg("Adding Effective Period to task")

		taskId, err := prompts.SelectTaskFromAll()
		if err != nil {
			log.Err(err).Msg("Error selecting task")
			return
		}

		startDate := prompt.PromptGetInput(prompt.PromptContent{Label: "Start Date"})
		endDate := prompt.PromptGetInput(prompt.PromptContent{Label: "End Date"})
		frequency := prompt.PromptGetInput(prompt.PromptContent{Label: "Frequency"})
		category := prompt.PromptGetInput(prompt.PromptContent{Label: "Category"})

		_, err = data.CreateEffectivePeriod(taskId, startDate, endDate, frequency, category)
		if err != nil {
			log.Err(err).Msg("Error creating Effective Period")
			return
		}
		log.Info().Msg("Effective Period created")

	},
}
