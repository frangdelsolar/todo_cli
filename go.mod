module github.com/frangdelsolar/todo_cli

go 1.22.4

replace github.com/frangdelsolar/todo_cli/pkg/auth => /Users/frangdelsolar/Desktop/code/todoProject/todo_cli/pkg/auth

require (
	github.com/frangdelsolar/todo_cli/pkg/auth v1.0.0
	github.com/frangdelsolar/todo_cli/pkg/currency v1.0.1
	github.com/frangdelsolar/todo_cli/pkg/data v1.0.2
	github.com/frangdelsolar/todo_cli/pkg/logger v1.0.3
	github.com/frangdelsolar/todo_cli/pkg/todo v1.0.41
	github.com/manifoldco/promptui v0.9.0
	github.com/spf13/cobra v1.8.1
)

require (
	github.com/chzyer/readline v1.5.1 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-sqlite3 v1.14.22 // indirect
	github.com/rs/zerolog v1.33.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/sys v0.21.0 // indirect
	gorm.io/driver/sqlite v1.5.6 // indirect
	gorm.io/gorm v1.25.10 // indirect
)
