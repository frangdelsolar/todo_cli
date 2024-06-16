package accounts

import (
	"github.com/frangdelsolar/todo_cli/cli/prompt"
	d "github.com/frangdelsolar/todo_cli/pkg/currency/data"
	m "github.com/frangdelsolar/todo_cli/pkg/currency/models"
	"github.com/spf13/cobra"
)


var CreateAccountCmd = &cobra.Command{
	Use:   "create",
	Short: "Create an Account",
	Run: func(cmd *cobra.Command, args []string) {
		var title string
		pc := prompt.PromptContent{
			Label: "Name",
			Validate: m.AccountNameValidator,
		}
		title = prompt.PromptGetInput(pc)

		var amount string
		pc = prompt.PromptContent{
			Label: "Amount",
			Validate: m.CurrencyAmountValidator,
		}
		amount = prompt.PromptGetInput(pc)

		var currencyCode string
		pc = prompt.PromptContent{
			Label: "Currency",
			Validate: m.CurrencyCodeValidator,
		}
		currencyCode = prompt.PromptGetInput(pc)

		task, err := d.CreateAccount(title, amount, currencyCode, false)
		if err != nil {
			log.Err(err).Msg("Error creating task")
			return
		}
		log.Info().Interface("task", task).Msg("Account created")

	},
}
