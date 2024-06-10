package currency

import "github.com/frangdelsolar/todo_cli/pkg/logger"


var log = logger.GetLogger()

var APP_VERSION = "1.0.0"

func main(){
	log.Info().Msgf("Running CURRENCY PKG v%s", APP_VERSION)
}