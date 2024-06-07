package task

import (
	"todo_cli/data"
	"todo_cli/internal/cli/prompts"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var DeleteTaskCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a task",
	Run: func(cmd *cobra.Command, args []string) {
		taskId, err := prompts.SelectTaskFromAll()
		if err != nil {
			log.Err(err).Msg("Error selecting task")
			return
		}

		err = data.DeleteTask(taskId)
		if err != nil {
			log.Err(err).Msg("Error deleting task")
			return
		}
		log.Info().Msg("Task deleted")

	},
}
