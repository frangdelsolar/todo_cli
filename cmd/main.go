package main

import (
	"fmt"

	"github.com/frangdelsolar/todo_cli/pkg/config"
)

var APP_NAME = "TODO APP"
var APP_VERSION = "1.5.0"

func main() {

	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	fmt.Println(cfg.DBPath)

}
