package logger

import (
	"os"

	"github.com/rs/zerolog"
)

var log Logger

var APP_VERSION = "1.0.0"

type Logger struct{
	*zerolog.Logger
}



func SetLogger(logger *zerolog.Logger) Logger{
	log = Logger{logger}
	return log
}

func GetLogger() *Logger {
	return &log
}

func init(){
	lg := zerolog.New(os.Stderr).With().Timestamp().Logger()
	lg.Info().Msgf("Running TODO APP v%s", APP_VERSION)
}