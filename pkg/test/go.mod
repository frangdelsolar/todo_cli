module github.com/frangdelsolar/todo_cli/pkg/test

go 1.22.4

replace github.com/frangdelsolar/todo_cli/pkg/config => /Users/frangdelsolar/Desktop/code/todoProject/todo_cli/pkg/config

replace github.com/frangdelsolar/todo_cli/pkg/currency => /Users/frangdelsolar/Desktop/code/todoProject/todo_cli/pkg/currency

require (
	github.com/frangdelsolar/todo_cli/pkg/auth v1.0.4
	github.com/frangdelsolar/todo_cli/pkg/config v1.0.2
	github.com/frangdelsolar/todo_cli/pkg/currency v1.0.3
	github.com/frangdelsolar/todo_cli/pkg/data v1.1.5
	github.com/frangdelsolar/todo_cli/pkg/logger v1.1.0
)

require (
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/mattn/go-sqlite3 v1.14.22 // indirect
	github.com/rs/zerolog v1.33.0 // indirect
	golang.org/x/sys v0.12.0 // indirect
	gorm.io/driver/sqlite v1.5.6 // indirect
	gorm.io/gorm v1.25.10 // indirect
)
