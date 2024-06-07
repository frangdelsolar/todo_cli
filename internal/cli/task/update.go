package task

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"todo_cli/data"
	"todo_cli/internal/cli/prompts"
	"todo_cli/models"
	"todo_cli/pkg/prompt"
)

var UpdateTaskCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a task",
	Run: func(cmd *cobra.Command, args []string) {
		taskId, err := prompts.SelectTaskFromAll()
		if err != nil {
			log.Err(err).Msg("Error selecting task")
			return
		}

		var title string
		pc := prompt.PromptContent{
			Label:    "Title",
			Validate: models.TaskTitleValidator,
		}
		title = prompt.PromptGetInput(pc)

		task, err := data.UpdateTask(taskId, title)

		if err != nil {
			log.Err(err).Msg("Error updating task")
			return
		}

		log.Debug().Interface("task", task).Msg("Task updated")
	},
}
