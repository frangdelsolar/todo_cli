package complete

import (
	"todo_cli/data"
	"todo_cli/internal/cli/prompts"
	"todo_cli/pkg/prompt"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var CompleteTaskCmd = &cobra.Command{
	Use:   "complete",
	Short: "Mark a task as completed",
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug().Msg("Completing task")

		taskId, err := prompts.SelectTaskFromPending()
		if err != nil || taskId == 0 {
			log.Err(err).Msg("Error selecting task")
			return
		}

		pc := prompt.PromptContent{
			Label: "Completed At",
		}
		completedAt := prompt.PromptGetInput(pc)

		data.CreateTaskCompletionLog(taskId, completedAt)
	},
}
