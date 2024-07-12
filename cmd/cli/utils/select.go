package utils

import (
	"github.com/manifoldco/promptui"
)

type SelectableItem struct {
	Key   string
	Label string
}

type SelectableList []SelectableItem

// GetLabels returns a slice of strings containing the labels of each SelectableItem in the SelectableList.
//
// No parameters.
// Returns a slice of strings.
func (sl SelectableList) GetLabels() []string {
	items := make([]string, len(sl))

	for i, item := range sl {
		items[i] = item.Label
	}

	return items
}

// SelectPrompt displays a prompt to the user with a list of selectable items and
// returns the selected item. It takes a label string and a SelectableList as
// parameters. The SelectableList is a slice of SelectableItem structs, each
// containing a Key and a Label. The function returns a pointer to the selected
// SelectableItem and an error if the prompt fails.
func SelectPrompt(label string, selectables SelectableList) (*SelectableItem, error) {
	prompt := promptui.Select{
		Label: label,
		Items: selectables.GetLabels(),
	}

	ix, _, err := prompt.Run()

	if err != nil {
		return nil, err
	}

	return &selectables[ix], nil
}
