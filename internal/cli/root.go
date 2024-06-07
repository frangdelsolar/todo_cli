package cli

import (
	taskcmd "todo_cli/internal/cli/task"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "TODO: A Powerful Command-Line Task Management Tool",
	Long: `Take control of your to-do list with TODO, 
a user-friendly and efficient command-line application.  
With TODO, you can ditch the scattered sticky notes and cluttered reminders, 
and organize your tasks in a centralized and accessible way.
`,
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(taskcmd.TaskCmd)
}
