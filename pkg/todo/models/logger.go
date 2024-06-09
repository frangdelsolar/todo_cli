package models

import "github.com/rs/zerolog"


type Logger struct{
	*zerolog.Logger
}

func NewLogger(logger *zerolog.Logger) *Logger {
	return &Logger{
		logger,
	}
}