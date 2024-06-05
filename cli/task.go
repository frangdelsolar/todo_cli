package cli

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// CreateTaskCmd creates a new CLI command for creating a task.
//
// This function returns a pointer to a cli.Command struct that represents the
// "create-task" command. The command has the following properties:
// - Name: "create-task"
// - Usage: "Create a task"
// - Aliases: []string{"ct"}
//
// The command's Action field is a function that is executed when the command is
// run. It takes a pointer to a cli.Context struct as a parameter. Inside the
// function, it retrieves the value of the "title" flag from the context and
// creates a new task using the retrieved title and the CLI's DB.
//
// The command has the following flag:
// - cli.StringFlag:
//   - Name: "title"
//   - Usage: "The title of the task"
//   - Required: true
//   - Aliases: []string{"t"}
//
// Returns:
// - *cli.Command: A pointer to the created cli.Command struct.
func (c *CLI) CreateTaskCmd() *cli.Command {
	return &cli.Command{
		Name:    "create-task",
		Usage:   "Create a task",
		Aliases: []string{"ct"},
		Action: func(cCtx *cli.Context) error {
			title := cCtx.String("title") 
			task, err := c.DB.CreateTask(title)
			if err != nil {
				return err
			}
			fmt.Println(task.String())
			return nil
		},

		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "title",
				Usage:    "The title of the task",
				Required: true,
				Aliases:  []string{"t"},
			},
		},
	}
}

// ListTasksCmd returns a cli.Command that lists all tasks.
//
// It creates a new cli.Command with the following properties:
// - Name: "list-tasks"
// - Usage: "List all tasks"
// - Aliases: []string{"lt"}
//
// The Action field is a function that is executed when the command is run.
// It retrieves all tasks from the CLI's DB and prints them to the console.
//
// Returns:
// - *cli.Command: A pointer to the created cli.Command struct.
func (c *CLI) ListTasksCmd() *cli.Command {
	return &cli.Command{
		Name:    "list-tasks",
		Usage:   "List all tasks",
		Aliases: []string{"lt"},
		Action: func(cCtx *cli.Context) error {
			tasks := c.DB.GetAllTasks()
			for _, task := range tasks {
				fmt.Println(task.String())
			}
			return nil
		},
	}
}

// UpdateTaskCmd returns a pointer to a cli.Command that represents the "edit-task" command.
//
// The command has the following properties:
// - Name: "edit-task"
// - Usage: "Edit a task"
// - Aliases: []string{"et"}
//
// The command's Action field is a function that is executed when the command is run.
// It takes a pointer to a cli.Context struct as a parameter. Inside the function,
// it retrieves the values of the "id" and "title" flags from the context and
// calls the UpdateTask method of the CLI's DB with the retrieved values.
//
// The command has the following flags:
// - cli.StringFlag:
//   - Name: "id"
//   - Usage: "The id of the task"
//   - Required: true
//   - Aliases: []string{"i"}
//   - Name: "title"
//   - Usage: "The title of the task"
//   - Required: true
//   - Aliases: []string{"t"}
//
// Returns:
// - *cli.Command: A pointer to the created cli.Command struct.
func (c *CLI) UpdateTaskCmd() *cli.Command {
	return &cli.Command{			
		Name:    "edit-task",
		Usage:   "Edit a task",
		Aliases: []string{"et"},
		Action: func(cCtx *cli.Context) error {
			id := cCtx.String("id")
			title := cCtx.String("title")
			task, err := c.DB.UpdateTask(id, title)
			if err != nil {
				return err
			}
			fmt.Println(task.String())
			return nil
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "id",
				Usage:    "The id of the task",
				Required: true,
				Aliases:  []string{"i"},
			},
			&cli.StringFlag{
				Name:     "title",
				Usage:    "The title of the task",
				Required: true,
				Aliases:  []string{"t"},
			},
		},
	}
}

// DeleteTaskCmd returns a pointer to a cli.Command that represents the "delete-task" command.
//
// The command has the following properties:
// - Name: "delete-task"
// - Usage: "Delete a task"
// - Aliases: []string{"dt"}
//
// The command's Action field is a function that is executed when the command is run.
// It takes a pointer to a cli.Context struct as a parameter. Inside the function,
// it retrieves the value of the "id" flag from the context and calls the DeleteTask
// method of the CLI's DB with the retrieved value.
//
// The command has the following flag:
// - cli.StringFlag:
//   - Name: "id"
//   - Usage: "The id of the task"
//   - Required: true
//   - Aliases: []string{"i"}
//
// Returns:
// - *cli.Command: A pointer to the created cli.Command struct.
func (c *CLI) DeleteTaskCmd() *cli.Command {
	return &cli.Command{
		Name:    "delete-task",
		Usage:   "Delete a task",
		Aliases: []string{"dt"},
		Action: func(cCtx *cli.Context) error {
			id := cCtx.String("id")
			c.DB.DeleteTask(id)
			fmt.Println("Task deleted successfully")
			return nil
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "id",
				Usage:    "The id of the task",
				Required: true,
				Aliases:  []string{"i"},
			},
		},
	}
}

// ListActiveTasksCmd returns a cli.Command that lists all active tasks.
//
// It creates a new cli.Command with the following properties:
// - Name: "list-active-tasks"
// - Usage: "List all active tasks"
// - Aliases: []string{"lat"}
//
// The Action field is a function that is executed when the command is run.
// It retrieves all active tasks from the CLI's DB and prints them to the console.
//
// Returns:
// - *cli.Command: A pointer to the created cli.Command struct.
func (c *CLI) ListActiveTasksCmd() *cli.Command {
	return &cli.Command{
		Name:    "list-active-tasks",
		Usage:   "List all active tasks",
		Aliases: []string{"lat"},
		Action: func(cCtx *cli.Context) error {
			tasks := c.DB.GetActiveTasks()
			for _, task := range tasks {
				fmt.Println(task.String())
			}
			return nil
		},
	}
}