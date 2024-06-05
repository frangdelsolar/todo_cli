package cli

import (
	"fmt"
	"todo_cli/db"

	"github.com/urfave/cli/v2"
)

func CreateTaskCmd() *cli.Command {
	return &cli.Command{
		Name:    "create-task",
		Usage:   "Create a task",
		Aliases: []string{"ct"},
		Action: func(cCtx *cli.Context) error {
			title := cCtx.String("title") 
			db.NewTask(title)
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

func ListTasksCmd() *cli.Command {
	return &cli.Command{
		Name:    "list-tasks",
		Usage:   "List all tasks",
		Aliases: []string{"lt"},
		Action: func(cCtx *cli.Context) error {
			tasks := db.GetAllTasks()
			for _, task := range tasks {
				fmt.Println(task)
			}
			return nil
		},
	}
}

func UpdateTaskCmd() *cli.Command {
	return &cli.Command{			
		Name:    "edit-task",
		Usage:   "Edit a task",
		Aliases: []string{"et"},
		Action: func(cCtx *cli.Context) error {
			id := cCtx.String("id")
			title := cCtx.String("title")
			db.UpdateTask(id, title)
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

func DeleteTaskCmd() *cli.Command {
	return &cli.Command{
		Name:    "delete-task",
		Usage:   "Delete a task",
		Aliases: []string{"dt"},
		Action: func(cCtx *cli.Context) error {
			id := cCtx.String("id")
			db.DeleteTask(id)
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