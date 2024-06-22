# Todo APP v1.0.6

Todo APP is a command-line tool that empowers you to efficiently manage your tasks. It provides a user-friendly interface to add, view, edit, and complete tasks, making it easy to stay organized and on top of your goals.

## Installation

### 1. Building from Source

-   Clone the Repository:

```bash
    git clone git@github.com:frangdelsolar/todo_cli.git
```

-   Navigate to the Project Directory:

```bash
    cd todo_cli
```

-   Build the Executable:

```bash
    go build .
```

-   Run the Application:

1. Open the `cmd` directory: `cd todo_cli/cmd`
2. Execute `go run .`

### Customization (Optional)

-   For a more convenient experience, consider adding an alias to your shell configuration file (e.g., ~/.zshrc or ~/.bashrc). This allows you to run todo commands directly from any directory:

```bash
    # In your shell configuration file
    todo() {
        cd <path/to/todo_app>  # Replace with the actual path to your Todo APP directory
        APP_ENV=dev go run cmd/main.go "$@"
    }
```

**Note**: Replace <path/to/todo_app> with the actual directory where you cloned the repository or built the executable.

## Usage

- You can login and register as a user :D
