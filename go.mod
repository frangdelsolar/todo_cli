module github.com/frangdelsolar/todo_cli

go 1.22.4

replace github.com/frangdelsolar/todo_cli/pkg/logger => /Users/frangdelsolar/Desktop/code/todoProject/todo_cli/pkg/logger

require (
	github.com/frangdelsolar/todo_cli/pkg/auth v1.0.4
	github.com/frangdelsolar/todo_cli/pkg/config v1.0.3
	github.com/frangdelsolar/todo_cli/pkg/data v1.1.5
	github.com/frangdelsolar/todo_cli/pkg/logger v1.1.1
	github.com/manifoldco/promptui v0.9.0
	github.com/spf13/cobra v1.8.1
)

require (
	github.com/chzyer/readline v0.0.0-20180603132655-2972be24d48e // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/mattn/go-sqlite3 v1.14.22 // indirect
	github.com/rs/zerolog v1.33.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/sys v0.12.0 // indirect
	gorm.io/driver/sqlite v1.5.6 // indirect
	gorm.io/gorm v1.25.10 // indirect
)
