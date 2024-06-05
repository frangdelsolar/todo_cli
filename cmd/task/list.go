package taskcmd

import (
	"fmt"
	"todo_cli/models"

	"github.com/spf13/cobra"
)

var ListTaskCmd = &cobra.Command{
	Use:    "list",
	Short:  "List Tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks := []models.Task{}

		all := cmd.Flag("all").Changed
		active := cmd.Flag("active").Changed
		// due := cmd.Flag("due").Changed
		if all {
			tasks = DB.GetAllTasks()
		} else if active {
			tasks = DB.GetActiveTasks()
		} 

		if len(tasks) == 0 {
			fmt.Println("No tasks found")
			return
		}

		for _, task := range tasks {
			fmt.Println(task.String())
		}

	},
}

func init() {
	ListTaskCmd.Flags().BoolP("all", "a", false, "List All tasks")
	ListTaskCmd.Flags().BoolP("active", "c", false, "List Active tasks")
	// ListTaskCmd.Flags().BoolP("due", "d", false, "List Due tasks")
}