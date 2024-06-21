package main

import (
	"github.com/frangdelsolar/todo_cli/pkg/logger"
	at "github.com/frangdelsolar/todo_cli/pkg/test/auth_test"
	co "github.com/frangdelsolar/todo_cli/pkg/test/contractor_test"
	ct "github.com/frangdelsolar/todo_cli/pkg/test/currency_test"
)

var PKG_NAME = "Test PKG"
var PKG_VERSION = "1.0.2"

var log *logger.Logger


func main(){
    
    log = logger.NewLogger(logger.LoggerConfig{
        PackageName: PKG_NAME,
        PackageVersion: PKG_VERSION,
    })

    log.Info().Msgf("Running %s v%s", PKG_NAME, PKG_VERSION)

    at.RunAuthTests()
    ct.RunCurrencyTests()
    co.RunContractorTests()
}
