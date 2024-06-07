package prompts

import (
	"fmt"
	"strconv"
	"todo_cli/data"
	"todo_cli/pkg/prompt"

	"github.com/rs/zerolog/log"
)

func SelectTask() uint {
	taskOptions := GetTaskItems()

	pc := prompt.PromptContent{
		Label: "Task",
		Items: taskOptions,
	}
	selection := prompt.GetSelectInput(pc)

	taskIdInt, err := strconv.Atoi(selection.Key)
	if err != nil {
		log.Err(err).Msg("Error parsing taskId")
		panic(err)
	}

	return uint(taskIdInt)
}

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
