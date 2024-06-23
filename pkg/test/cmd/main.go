package main

import (
	"fmt"
	"os"

	"github.com/frangdelsolar/todo_cli/pkg/config"
	"github.com/frangdelsolar/todo_cli/pkg/logger"
	"github.com/frangdelsolar/todo_cli/pkg/test"

	cli "github.com/frangdelsolar/todo_cli/pkg/test/cli_test"
	co "github.com/frangdelsolar/todo_cli/pkg/test/contractor_test"
	ct "github.com/frangdelsolar/todo_cli/pkg/test/currency_test"
)

var PKG_NAME = "Test PKG"
var PKG_VERSION = "1.0.3"

var log *logger.Logger
var cfg *config.Config

var t *test.TestRunner

func main(){
    var err error
    cfg, err = config.Load()
    if err != nil {
        fmt.Errorf("Failed to load config: %v", err)
    }
    
    log = logger.NewLogger(logger.LoggerConfig{
        PackageName: PKG_NAME,
        PackageVersion: PKG_VERSION,
    })

    log.Info().Msgf("Running %s v%s", PKG_NAME, PKG_VERSION)
    log.Debug().Interface("Config", cfg).Msg("Loaded Config")

    t = test.NewTestRunner()

    ct.RunCurrencyTests(t)
    co.RunContractorTests(t)
    cli.RunCliTests()

    if len(t.Errors) > 0 {
        log.Warn().Msgf("Found %d errors", len(t.Errors))
        os.Exit(1)
    }

}
