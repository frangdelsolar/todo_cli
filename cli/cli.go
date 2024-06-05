package cli

import (
	"os"
	"todo_cli/db"

	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
)

// CLI represents the CLI app.
type CLI struct {
	App *cli.App
	DB  *db.DB
}

// NewCLI creates a new CLI instance.
//
// This function initializes a new CLI instance and sets up the necessary
// configuration for the CLI application. 
//
// The `Commands` field of the `cli.App` instance is set to an array of
// `cli.Command` pointers. Each `cli.Command` pointer represents a command
// that can be executed by the CLI application.
//
// Returns:
// - *CLI: A pointer to the newly created CLI instance.
func NewCLI() (*CLI, error){
	c := &CLI{}

	db, err := db.NewDB()
	if err != nil {
		log.Err(err).Msg("Error creating new DB")
		return nil, err
	}
	c.DB = db

	c.App = &cli.App{
		Commands:   []*cli.Command{
			c.CreateTaskCmd(),
			c.ListTasksCmd(),
			c.UpdateTaskCmd(),
			c.DeleteTaskCmd(),
			c.ListActiveTasksCmd(),
			c.CreateEffectivePeriodCmd(),
		},
	}

	
	return c, nil
}

// Run executes the CLI application.
//
// It runs the CLI application using the provided command-line arguments.
// It returns an error if there was a problem running the application.
// Returns:
// - error: An error if the CLI application failed to run.
func (c *CLI) Run() error {
	err := c.App.Run(os.Args)
	if err != nil {
		log.Err(err).Msg("Failed to run CLI")
		return err
	}
	return nil
}