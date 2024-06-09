package task

import (
	"github.com/frangdelsolar/todo_cli/pkg/todo/cli/prompt"
	"github.com/frangdelsolar/todo_cli/pkg/todo/cli/prompts"
	"github.com/frangdelsolar/todo_cli/pkg/todo/data"
	"github.com/frangdelsolar/todo_cli/pkg/todo/models"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
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

		log.Info().Interface("task", task).Msg("Task updated")
	},
}
