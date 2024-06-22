package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/frangdelsolar/todo_cli/pkg/logger"
	at "github.com/frangdelsolar/todo_cli/pkg/test/auth_test"
	cli "github.com/frangdelsolar/todo_cli/pkg/test/cli_test"
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
    cli.RunCliTests()

    // Find errors in log files
    logsPattern := "*.log"
    logFiles, err := filepath.Glob(logsPattern)
    if err != nil {
        log.Error().Err(err).Msgf("Error finding log files: %v", err)
        return
    }
    errorPattern := "\"level\":\"error\""
    for _, fileName := range logFiles {
        // Use grep command to search for "error"
        cmd := exec.Command("grep", "-i", errorPattern, fileName)
        output, err := cmd.CombinedOutput()
        if err != nil {
            fmt.Errorf("error running grep command: %v", err)
            continue
        }

        if len(output) > 0 {
            log.Warn().Msgf("Found errors in %s:\n%s", fileName, string(output))
            os.Exit(1)
        }
    }

}
