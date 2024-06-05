package taskcmd

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var UpdateTaskCmd = &cobra.Command{
	Use:    "update",
	Short:  "Update a task",
	Run: func(cmd *cobra.Command, args []string) {
		taskId := cmd.Flag("taskId").Value.String()
		title := cmd.Flag("title").Value.String()
		task, err := DB.UpdateTask(taskId, title)
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