package test

import (
	"github.com/frangdelsolar/todo_cli/pkg/data"
	"github.com/frangdelsolar/todo_cli/pkg/logger"
)


var log *logger.Logger
var db *data.Database
func PrepareTestingSuite(){
    log = logger.NewLogger("Test PKG", "0.0.1")
    db, _ = data.GetDB()
    
    log.Info().Msg("Testing Suite Initialized")
}
