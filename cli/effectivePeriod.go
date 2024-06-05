package cli

import (
	"github.com/urfave/cli/v2"
)

// CreateEffectivePeriodCmd creates a new CLI command for creating an effective period for a task.
//
// This function returns a pointer to a cli.Command struct that represents the
// "create-effective-period" command. The command has the following properties:
// - Name: "create-effective-period"
// - Usage: "Create an effective period for a task"
// - Aliases: []string{"ctp"}
//
// The command's Action field is a function that is executed when the command is
// run. It takes a pointer to a cli.Context struct as a parameter. Inside the
// function, it retrieves the values of the "taskId", "startDate", and "endDate"
// flags from the context and calls the CreateEffectivePeriod method of the CLI's
// DB with the retrieved values.
//
// Returns:
// - *cli.Command: A pointer to the created cli.Command struct.
func (c *CLI) CreateEffectivePeriodCmd() *cli.Command {
	return &cli.Command{
		Name:    "create-effective-period",
		Usage:   "Create an effective period for a task",
		Aliases: []string{"cep"},
		Action: func(cCtx *cli.Context) error {
			taskId := cCtx.String("taskId")
			startDate := cCtx.String("startDate")
			endDate := cCtx.String("endDate")
			frequency := cCtx.String("frequency")
			c.DB.CreateEffectivePeriod(taskId, startDate, endDate, frequency)
			return nil
		},

		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "taskId",
				Usage:    "The id of the task",
				Required: true,
				Aliases:  []string{"t"},
			},

			&cli.StringFlag{
				Name:     "startDate",
				Usage:    "The start date of the effective period (yyyy-mm-dd)",
				Required: true,
				Aliases:  []string{"sd"},
			},

			&cli.StringFlag{
				Name:     "endDate",
				Usage:    "The end date of the end period (yyyy-mm-dd)",
				Required: false,
				Aliases:  []string{"ed"},
			},

			&cli.StringFlag{
				Name:     "frequency",
				Usage:    "The frequency of the effective period [daily, weekly, monthly, yearly]",
				Required: true,
				Aliases:  []string{"f"},
			},
		},
	}
}
