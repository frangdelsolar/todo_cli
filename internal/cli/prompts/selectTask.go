package prompts

import (
	"fmt"
	"strconv"
	"time"

	"todo_cli/data"
	"todo_cli/models"
	"todo_cli/pkg/prompt"

	"github.com/rs/zerolog/log"
)

type TaskSelection struct {
	Label string
	Items []prompt.SelectableItem
}

func SelectTask(title string, taskGetter func() []models.Task) (uint, error) {
	tasks := taskGetter()

	if len(tasks) == 0 {
		return 0, fmt.Errorf("no tasks found")
	}

	items := convertTasksToSelectableItems(tasks)

	pc := prompt.PromptContent{
		Label: title,
		Items: items,
	}

	selection := prompt.GetSelectInput(pc)

	taskIdInt, err := strconv.Atoi(selection.Key)
	if err != nil {
		return 0, fmt.Errorf("error parsing taskId: %w", err)
	}

	return uint(taskIdInt), nil
}

func convertTasksToSelectableItems(tasks []models.Task) []prompt.SelectableItem {
	var items []prompt.SelectableItem
	for _, task := range tasks {
		items = append(items, prompt.SelectableItem{
			Label: task.String(),
			Key:   fmt.Sprint(task.ID),
		})
	}
	return items
}

func SelectTaskFromAll() (uint, error) {
	return SelectTask("Select Task", data.GetAllTasks)
}

func SelectTaskFromPending() (uint, error) {
	log.Debug().Msg("Selecting task from pending")
	return SelectTask("Select Pending Task", func() []models.Task {
		return data.GetPendingTasksTodoMonthly(time.Now())
	})
}
