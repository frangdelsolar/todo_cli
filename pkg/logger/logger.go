package logger

import (
	"github.com/rs/zerolog"
)

var PKG_VERSION = "1.0.0"

var log Logger

type Logger struct{
	*zerolog.Logger
}

// NewLogger creates a new Logger instance with the provided zerolog.Logger.
//
// Parameters:
// - logger: A pointer to a zerolog.Logger instance.
//
// Returns:
// - A pointer to the newly created Logger instance.
func NewLogger(logger *zerolog.Logger) *Logger{
	log = Logger{logger}
	return &log
}

// GetLogger returns a pointer to the Logger instance.
//
// Returns:
// - *Logger: A pointer to the Logger instance.
func GetLogger() *Logger {
	return &log
}
