package taskcmd

import (
	"strconv"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"todo_cli/data"
)

var UpdateTaskCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a task",
	Run: func(cmd *cobra.Command, args []string) {
		taskId := cmd.Flag("taskId").Value.String()
		taskIdInt, err := strconv.Atoi(taskId)
		if err != nil {
			log.Err(err).Msg("Error parsing taskId")
			return
		}
		title := cmd.Flag("title").Value.String()
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
	UpdateTaskCmd.Flags().IntP("taskId", "i", 0, "The id of the task")
}
