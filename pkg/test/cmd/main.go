package main

import (
	"github.com/frangdelsolar/todo_cli/pkg/logger"
	at "github.com/frangdelsolar/todo_cli/pkg/test/auth_test"
	ct "github.com/frangdelsolar/todo_cli/pkg/test/currency_test"
)

var PKG_NAME = "Test PKG"
var PKG_VERSION = "0.0.1"

var log *logger.Logger


func main(){
    
    log = logger.NewLogger(PKG_NAME, PKG_VERSION)
    log.Info().Msgf("Running %s v%s", PKG_NAME, PKG_VERSION)

    at.RunAuthTests()
    ct.RunCurrencyTests()
}
