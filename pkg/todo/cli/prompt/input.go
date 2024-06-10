package prompt

import (
	"fmt"
	"os"

	"github.com/frangdelsolar/todo_cli/pkg/todo/logger"
	"github.com/manifoldco/promptui"
)

var log = logger.GetLogger()

type SelectableItem struct {
	Label string
	Key   string
}

type PromptContent struct {
	Label    string
	Items    []SelectableItem
	Validate func(input string) error
}

func PromptGetInput(pc PromptContent) string {

	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }} ",
		Valid:   "{{ . | green }} ",
		Invalid: "{{ . | red }} ",
		Success: "{{ . | bold }} ",
	}

	prompt := promptui.Prompt{
		Label:     pc.Label,
		Templates: templates,
		Validate:  pc.Validate,
	}

	result, err := prompt.Run()
	if err != nil {
		log.Err(err).Msg("Prompt failed")
		return ""
	}

	return result
}

func GetSelectInput(pc PromptContent) SelectableItem {
	index := -1
	var result SelectableItem
	var err error

	labels := make([]string, len(pc.Items))
	for i, item := range pc.Items {
		labels[i] = item.Label
	}

	for index < 0 {
		prompt := promptui.Select{
			Label: pc.Label,
			Items: labels,
		}

		index, _, err = prompt.Run()
	}

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	result = pc.Items[index]

	return result
}
