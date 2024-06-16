# Todo PKG v1.0.6

This package provides a set of models and methods for managing tasks in your Go application. It can be integrated seamlessly into various environments, including command-line interfaces (CLIs), APIs, and GraphQL servers.

## Features:

-   **Models**: Represent tasks, task goals (schedules for tasks), task completion logs, and task frequencies.
-   **Methods**: Offer functionalities for CRUD operations (Create, Read, Update, Delete) on tasks and task goals, managing task completion logs, and creating different task frequencies.

## Installation:

Use the go get command to install the package:

```bash
    go get github.com/frangdelsolar/todo_cli/pkg/todo
```

## Initialization:

Before using the package's methods, it's crucial to call the InitTodo() function. This initializes the database connection, applies database migrations (schema creation), and sets up the logger:

```go
    import (
        "github.com/frangdelsolar/todo_cli/pkg/todo"
    )

    func main() {
        todo.InitTodo()
        // ... your application code using todo package methods
    }
```

## API Documentation

The detailed API documentation for each method is available within the package source code. For a quick overview of the functionalities, refer to the comments within the code.

## Example Usage

### Creating a Task:

```go
    import (
        "fmt"

        "github.com/frangdelsolar/todo_cli/pkg/todo"
        "github.com/frangdelsolar/todo_cli/pkg/todo/models"
    )

    func main() {
        todo.InitTodo()

        newTask, err := todo.CreateTask("Buy groceries")
        if err != nil {
            fmt.Println("Error creating task:", err)
            return
        }

        fmt.Printf("Successfully created task: %s\n", newTask.Title)
    }
```

### Retrieving All Tasks:

```go
    import (
        "fmt"

        "github.com/frangdelsolar/todo_cli/pkg/todo"
        "github.com/frangdelsolar/todo_cli/pkg/todo/models"
    )

    func main() {
        todo.InitTodo()

        tasks := todo.GetAllTasks()
        if len(tasks) == 0 {
            fmt.Println("No tasks found")
            return
        }

        for _, task := range tasks {
            fmt.Printf("Task: %s (ID: %s)\n", task.Title, task.ID)
        }
    }
```
