package task

import (
	"fmt"
	"time"

	"github.com/frangdelsolar/todo_cli/pkg/todo/cli/prompt"
	"github.com/frangdelsolar/todo_cli/pkg/todo/cli/task/complete"
	"github.com/frangdelsolar/todo_cli/pkg/todo/cli/task/goal"
	"github.com/frangdelsolar/todo_cli/pkg/todo/data"
	"github.com/frangdelsolar/todo_cli/pkg/todo/logger"
	"github.com/spf13/cobra"
)

var log = logger.GetLogger()


var taskCmdActions = []prompt.SelectableItem{
	{Key: "create", Label: "Add a task"},
	{Key: "update", Label: "Update a task"},
	{Key: "delete", Label: "Delete a task"},
	{Key: "list", Label: "List tasks"},
	{Key: "task_goal", Label: "Task goals"},
	{Key: "complete", Label: "Complete a task"},
	{Key: "exit", Label: "Exit"},
}

var TaskCmd = &cobra.Command{
	Use:   "task",
	Short: "Task actions",
	Run: func(cmd *cobra.Command, args []string) {

		// Show pending tasks
		tasks := data.GetPendingTaskCompletionLogs(time.Now())
		if len(tasks) > 0 {
			fmt.Print("These are your pending tasks: \n")
			for ix, task := range tasks {
				fmt.Printf("%d | %s \n", ix, task.String())
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
		case "task_goal":
			goal.TaskGoalCmd.Run(cmd, args)
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
	TaskCmd.AddCommand(goal.TaskGoalCmd)
	TaskCmd.AddCommand(complete.CompleteTaskCmd)
}
