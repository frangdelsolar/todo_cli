package task

import (
	"todo_cli/internal/cli/task/period"

	"github.com/spf13/cobra"
)

var TaskCmd = &cobra.Command{
	Use:   "task",
	Short: "Task actions",
}

func init() {
	TaskCmd.AddCommand(CreateTaskCmd)
	TaskCmd.AddCommand(ListTaskCmd)
	TaskCmd.AddCommand(UpdateTaskCmd)
	TaskCmd.AddCommand(DeleteTaskCmd)
	TaskCmd.AddCommand(period.EffectivePeriodCmd)
}
