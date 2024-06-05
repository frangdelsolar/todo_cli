package cli

import (
	"os"

	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
)
type CLI struct {
	App *cli.App
}

func NewCLI() *CLI {
	c := &CLI{}
	c.App = &cli.App{}
	c.init()
	return c
}

func (c *CLI) Run() {
	err := c.App.Run(os.Args)
	if err != nil {
		log.Err(err).Msg("Failed to run CLI")
	}
}

func (c *CLI) RegisterCommand(command *cli.Command) {
	c.App.Commands = append(c.App.Commands, command)
}

func (c *CLI) init () {
	c.RegisterCommand(CreateTaskCmd())
	c.RegisterCommand(ListTasksCmd())
	c.RegisterCommand(UpdateTaskCmd())
	c.RegisterCommand(DeleteTaskCmd())

}