package complete

import (
	"time"

	"github.com/frangdelsolar/todo_cli/pkg/todo/cli/prompts"
	"github.com/frangdelsolar/todo_cli/pkg/todo/data"
	"github.com/frangdelsolar/todo_cli/pkg/todo/logger"
	"github.com/spf13/cobra"
)

var log = logger.GetLogger()

var CompleteTaskCmd = &cobra.Command{
	Use:   "complete",
	Short: "Mark a task as completed",
	Run: func(cmd *cobra.Command, args []string) {
		tcl, err := prompts.SelectTaskFromPending()
		if err != nil {
			log.Err(err).Msg("Error selecting task")
			return
		}

		data.CreateTaskCompletionLog(tcl.TaskID, tcl.DueDate.Format(time.DateOnly), tcl.TaskGoalID)

		log.Info().Interface("task", tcl).Msg("Task completed")
	},
}
