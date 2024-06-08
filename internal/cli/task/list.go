package task

import (
	"fmt"
	"time"
	"todo_cli/data"
	"todo_cli/models"
	"todo_cli/pkg/prompt"

	"github.com/spf13/cobra"
)

var ListTaskCmdActions = []prompt.SelectableItem{
	{Key: "all", Label: "All"},
	{Key: "active", Label: "Active"},
	// {Key: "pending", Label: "Pending"},
}

var ListTaskCmd = &cobra.Command{
	Use:   "list",
	Short: "List Tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks := []models.Task{}
		fmt.Print("Select a list")

		pc := prompt.PromptContent{
			Label: "Available lists",
			Items: ListTaskCmdActions,
		}
		selection := prompt.GetSelectInput(pc)

		all := selection.Key == "all"
		active := selection.Key == "active"
		// pending := selection.Key == "pending"

		if all {
			tasks = data.GetAllTasks()
		} else if active {
			tasks = data.GetActiveTasks(time.Now())
		} //else if pending {
		// 	tasks = data.GetPendingTasksTodoMonthly(time.Now())
		// }

		if len(tasks) == 0 {
			fmt.Println("No tasks found")
			return
		}

		fmt.Printf("%s Tasks:\n", selection.Label)
		for _, task := range tasks {
			fmt.Println(task.String())
		}

	},
}

func init() {
	ListTaskCmd.Flags().BoolP("all", "a", false, "List All tasks")
	ListTaskCmd.Flags().BoolP("active", "c", false, "List Active tasks")
}
