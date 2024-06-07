package task

import (
	"todo_cli/data"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var CreateTaskCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a task",
	Run: func(cmd *cobra.Command, args []string) {
		log.Info().Msg("Creating task")

		title := cmd.Flag("title").Value.String()
		task, err := data.CreateTask(title)
		if err != nil {
			log.Err(err).Msg("Error creating task")
			return
		}
		log.Info().Msg("Task created")
		log.Debug().Interface("task", task)
	},
}

func init() {
	CreateTaskCmd.Flags().StringP("title", "t", "", "The title of the task")
}
