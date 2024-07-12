package auth

import (
	"fmt"

	"github.com/frangdelsolar/todo_cli/pkg/config"
	"github.com/frangdelsolar/todo_cli/pkg/data"
	"github.com/frangdelsolar/todo_cli/pkg/logger"
)

var PKG_NAME = "Auth Test PKG"
var PKG_VERSION = "1.0.4"

func init() {

	cfg, err := config.Load()
	if err != nil {
		fmt.Errorf("error loading config: %v", err)
		return
	}

	log := logger.NewLogger(logger.LoggerConfig{
		PackageName:    PKG_NAME,
		PackageVersion: PKG_VERSION,
	})

	log.Trace().Interface("Config", cfg).Msg("Loaded Config")

	data.LoadDB()
}
