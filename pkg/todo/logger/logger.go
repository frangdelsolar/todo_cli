package logger

import "github.com/rs/zerolog"

var log Logger

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