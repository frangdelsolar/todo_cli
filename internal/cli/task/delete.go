package task

import (
	"strconv"
	"todo_cli/data"
	"todo_cli/internal/cli/utils"
	"todo_cli/pkg/prompt"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var DeleteTaskCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a task",
	Run: func(cmd *cobra.Command, args []string) {
		taskId := cmd.Flag("taskId").Value.String()

		if taskId == "" {
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

		err = data.DeleteTask(uint(taskIdInt))
		if err != nil {
			log.Err(err).Msg("Error deleting task")
			return
		}
		log.Info().Msg("Task deleted")

	},
}

func init() {
	DeleteTaskCmd.Flags().StringP("taskId", "i", "", "The id of the task")
}
