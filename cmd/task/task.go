package taskcmd

import (
	"todo_cli/db"

	"github.com/spf13/cobra"
)

var DB, _ = db.NewDB()

var TaskCmd = &cobra.Command{
	Use:    "task",
	Short:  "Task actions",
}

func init() {
	TaskCmd.AddCommand(AddEffectivePeriodCmd)
	TaskCmd.AddCommand(CreateTaskCmd)
	TaskCmd.AddCommand(CompleteTaskCmd)
	TaskCmd.AddCommand(ListTaskCmd)
	TaskCmd.AddCommand(UpdateTaskCmd)
	TaskCmd.AddCommand(DeleteTaskCmd)
}