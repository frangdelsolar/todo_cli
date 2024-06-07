package task

import (
	"strconv"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"todo_cli/data"
	"todo_cli/internal/cli/utils"
	"todo_cli/models"
	"todo_cli/pkg/prompt"
)

var UpdateTaskCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a task",
	Run: func(cmd *cobra.Command, args []string) {
		taskId := cmd.Flag("taskId").Value.String()
		_, err := strconv.Atoi(taskId)

		if taskId == "" || err != nil {
			taskOptions := utils.GetTaskItems()

			pc := prompt.PromptContent{
				Label: "Task",
				Items: taskOptions,
			}
			selection := prompt.GetSelectInput(pc)
			taskId = selection.Key
		}

		taskIdInt, err := strconv.Atoi(taskId)
		if err != nil {
			log.Err(err).Msg("Error parsing taskId")
			return
		}

		title := cmd.Flag("title").Value.String()
		if title == "" {
			pc := prompt.PromptContent{
				Label:    "Title",
				Validate: models.TaskTitleValidator,
			}
			title = prompt.PromptGetInput(pc)
		}
		task, err := data.UpdateTask(uint(taskIdInt), title)
		if err != nil {
			log.Err(err).Msg("Error updating task")
			return
		}
		log.Info().Msg("Task updated")
		log.Debug().Interface("task", task)
	},
}

func init() {
	UpdateTaskCmd.Flags().StringP("title", "t", "", "The title of the task")
	UpdateTaskCmd.Flags().StringP("taskId", "i", "", "The id of the task")
}
