package prompts

import (
	"fmt"
	"time"

	"todo_cli/data"
	"todo_cli/models"
	"todo_cli/pkg/prompt"
)

type TaskSelection struct {
	Label string
	Items []prompt.SelectableItem
}

func SelectTask(title string, taskGetter func() []models.Task) (string, error) {
	tasks := taskGetter()

	if len(tasks) == 0 {
		return "", fmt.Errorf("no tasks found")
	}

	items := convertTasksToSelectableItems(tasks)

	pc := prompt.PromptContent{
		Label: title,
		Items: items,
	}

	selection := prompt.GetSelectInput(pc)

	return selection.Key, nil
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

func SelectTaskFromAll() (string, error) {
	return SelectTask("Select Task", data.GetAllTasks)
}

func SelectTaskFromPending() (string, error) {
	return SelectTask("Select Pending Task", func() []models.Task {
		return data.GetPendingTasksTodoMonthly(time.Now())
	})
}
