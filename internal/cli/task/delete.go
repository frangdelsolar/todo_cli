package task

import (
	"strconv"
	"todo_cli/data"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var DeleteTaskCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a task",
	Run: func(cmd *cobra.Command, args []string) {
		taskId := cmd.Flag("taskId").Value.String()
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
	DeleteTaskCmd.Flags().IntP("taskId", "i", 0, "The id of the task")
}
