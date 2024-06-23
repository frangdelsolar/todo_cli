package contractor

import (
	"github.com/frangdelsolar/todo_cli/pkg/auth"
	"github.com/frangdelsolar/todo_cli/pkg/config"
	"github.com/frangdelsolar/todo_cli/pkg/data"
	"github.com/frangdelsolar/todo_cli/pkg/logger"

	c "github.com/frangdelsolar/todo_cli/pkg/contractor"
)

var PKG_NAME = "Contractor Test PKG"
var PKG_VERSION = "1.0.4"

func init(){
    config.Load()
    
    logger.NewLogger(logger.LoggerConfig{
        PackageName: PKG_NAME,
        PackageVersion: PKG_VERSION,
    })

	data.LoadDB()
    auth.InitAuth()
    c.InitContractor()
}


