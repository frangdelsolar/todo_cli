package prompts

import (
	"fmt"
	"strconv"
	"time"

	"todo_cli/data"
	"todo_cli/models"
	"todo_cli/pkg/prompt"
)

type ItemSelection struct {
	Label string
	Items []prompt.SelectableItem
}



func SelectTaskFromAll() (string, error) {
	tasks := data.GetAllTasks()

	if len(tasks) == 0 {
		return "", fmt.Errorf("no tasks found")
	}

	items := convertTasksToSelectableItems(tasks)

	pc := prompt.PromptContent{
		Label: "Select Task",
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

func SelectTaskFromPending() (data.PendingTCLContract, error) {
	pending := data.GetPendingTaskCompletionLogs(time.Now())

	if len(pending) == 0 {
		return data.PendingTCLContract{}, fmt.Errorf("no pending tasks found")
	}

	items := convertPendingsToSelectableItems(pending)

	pc := prompt.PromptContent{
		Label: "Select Task",
		Items: items,
	}

	selection := prompt.GetSelectInput(pc)

	ix, _ := strconv.Atoi(selection.Key)

	return pending[ix], nil
}

func convertPendingsToSelectableItems(pending []data.PendingTCLContract) []prompt.SelectableItem {
	var items []prompt.SelectableItem
	for index, item := range pending {
		items = append(items, prompt.SelectableItem{
			Label: fmt.Sprintf("%d | %s", index, item.String()),
			Key:   fmt.Sprint(index),
		})
	}
	return items
}
