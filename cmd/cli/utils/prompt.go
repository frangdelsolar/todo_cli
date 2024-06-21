package utils

import (
	"github.com/manifoldco/promptui"
)

type PromptConfig struct {
    Label    string
    Validate func(input string) error
}

// Prompt displays a prompt to the user with the provided configuration.
//
// The function takes a PromptConfig struct as a parameter, which contains the label
// for the prompt and a validation function for the user input. It returns the user's
// input as a string and an error if any occurred during the prompt.
func Prompt(config PromptConfig) (string, error) {
	prompt := promptui.Prompt{
		Label:    config.Label,
		Validate: config.Validate,
	}

	return prompt.Run()
}
