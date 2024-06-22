package main

import (
	"github.com/frangdelsolar/todo_cli/pkg/logger"
)

var PKG_NAME = "GH Test PKG"
var PKG_VERSION = "1.0.2"

var log *logger.Logger


func main(){
    
    log = logger.NewLogger(logger.LoggerConfig{
        PackageName: PKG_NAME,
        PackageVersion: PKG_VERSION,
    })

    log.Info().Msgf("Running %s v%s", PKG_NAME, PKG_VERSION)

    // err := errors.New("Some random shit")
    // log.Err(err).Msg("More shit here")

    // if err != nil {
        // log.Fatal().Msg("Exiting with error")
        // os.Exit(1) // Exit with non-zero code
    // }
}
