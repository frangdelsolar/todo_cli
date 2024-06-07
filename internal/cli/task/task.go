package task

import (
	"fmt"
	"time"
	"todo_cli/data"
	"todo_cli/internal/cli/task/complete"
	"todo_cli/internal/cli/task/period"
	"todo_cli/pkg/prompt"

	"github.com/spf13/cobra"
)

var taskCmdActions = []prompt.SelectableItem{
	{Key: "create", Label: "Add a task"},
	{Key: "update", Label: "Update a task"},
	{Key: "delete", Label: "Delete a task"},
	{Key: "list", Label: "List tasks"},
	{Key: "period", Label: "Effective periods"},
	{Key: "complete", Label: "Complete a task"},
	{Key: "exit", Label: "Exit"},
}

var TaskCmd = &cobra.Command{
	Use:   "task",
	Short: "Task actions",
	Run: func(cmd *cobra.Command, args []string) {

		// Show pending tasks
		tasks := data.GetPendingTasksTodoMonthly(time.Now())
		if len(tasks) > 0 {
			fmt.Print("These are your pending tasks:")
			for _, task := range tasks {
				fmt.Println(task.String())
			}
		} else {
			fmt.Println("No pending tasks today!!!")
		}

		fmt.Print("Select an action")
		pc := prompt.PromptContent{
			Label: "Action",
			Items: taskCmdActions,
		}
		action := prompt.GetSelectInput(pc)
		switch action.Key {
		case "create":
			CreateTaskCmd.Run(cmd, args)
		case "update":
			UpdateTaskCmd.Run(cmd, args)
		case "delete":
			DeleteTaskCmd.Run(cmd, args)
		case "list":
			ListTaskCmd.Run(cmd, args)
		case "period":
			period.EffectivePeriodCmd.Run(cmd, args)
		case "complete":
			complete.CompleteTaskCmd.Run(cmd, args)
		}
	},
}

func init() {
	TaskCmd.AddCommand(CreateTaskCmd)
	TaskCmd.AddCommand(ListTaskCmd)
	TaskCmd.AddCommand(UpdateTaskCmd)
	TaskCmd.AddCommand(DeleteTaskCmd)
	TaskCmd.AddCommand(period.EffectivePeriodCmd)
	TaskCmd.AddCommand(complete.CompleteTaskCmd)
}
