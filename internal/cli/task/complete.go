package taskcmd

import (
	"strconv"
	"todo_cli/data"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var CompleteTaskCmd = &cobra.Command{
	Use:   "complete",
	Short: "Mark a task as completed",
	Run: func(cmd *cobra.Command, args []string) {
		log.Info().Msg("Marking task as completed")
		taskId := cmd.Flag("taskId").Value.String()
		taskIdInt, err := strconv.Atoi(taskId)
		if err != nil {
			log.Err(err).Msg("Error parsing taskId")
			return
		}
		completedAt := cmd.Flag("completedAt").Value.String()
		_, err = data.CreateTaskCompletionLog(uint(taskIdInt), completedAt)
		if err != nil {
			log.Err(err).Msg("Error marking task as completed")
			return
		}
		log.Info().Msg("Task marked as completed")
	},
}

func init() {
	CompleteTaskCmd.Flags().IntP("taskId", "i", 0, "The ID of the task to mark as completed")
	CompleteTaskCmd.MarkFlagRequired("taskId")

	CompleteTaskCmd.Flags().StringP("completedAt", "c", "", "The date and time when the task was completed (Optional) Format: YYYY-MM-DD")
}
