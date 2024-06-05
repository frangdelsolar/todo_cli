package taskcmd

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var DeleteTaskCmd = &cobra.Command{
	Use:    "delete",
	Short:  "Delete a task",
	Run: func(cmd *cobra.Command, args []string) {
		taskId := cmd.Flag("taskId").Value.String()

		err := DB.DeleteTask(taskId)
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