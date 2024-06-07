package utils

import (
	"fmt"
	"todo_cli/data"
	"todo_cli/pkg/prompt"
)

func GetTaskItems() []prompt.SelectableItem {
	tasks := data.GetAllTasks()

	var items []prompt.SelectableItem
	for _, task := range tasks {
		items = append(items, prompt.SelectableItem{
			Label: task.String(),
			Key:   fmt.Sprint(task.ID),
		})
	}
	return items
}
