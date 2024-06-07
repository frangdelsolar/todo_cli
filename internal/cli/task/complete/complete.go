package complete

import (
	"todo_cli/data"
	"todo_cli/internal/cli/prompts"
	"todo_cli/pkg/prompt"

	"github.com/spf13/cobra"
)

var CompleteTaskCmd = &cobra.Command{
	Use:   "complete",
	Short: "Mark a task as completed",
	Run: func(cmd *cobra.Command, args []string) {
		taskId := prompts.SelectTask()

		pc := prompt.PromptContent{
			Label: "Completed At",
		}
		completedAt := prompt.PromptGetInput(pc)
		data.CreateTaskCompletionLog(taskId, completedAt)
	},
}

func init() {
	CompleteTaskCmd.Flags().IntP("taskId", "i", 0, "The ID of the task to mark as completed")
	CompleteTaskCmd.MarkFlagRequired("taskId")

	CompleteTaskCmd.Flags().StringP("completedAt", "c", "", "The date and time when the task was completed (Optional) Format: YYYY-MM-DD")
}
