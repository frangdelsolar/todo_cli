package auth

import (
	"github.com/frangdelsolar/todo_cli/pkg/auth"
	"github.com/frangdelsolar/todo_cli/pkg/config"
	"github.com/frangdelsolar/todo_cli/pkg/data"
	"github.com/frangdelsolar/todo_cli/pkg/logger"
)

var PKG_NAME = "Auth Test PKG"
var PKG_VERSION = "1.0.4"

func init(){
    config.Load()
    
    logger.NewLogger(logger.LoggerConfig{
        PackageName: PKG_NAME,
        PackageVersion: PKG_VERSION,
    })

	data.LoadDB()
    auth.InitAuth()
}
