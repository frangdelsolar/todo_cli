package logger

import (
	"fmt"
	"os"

	"github.com/frangdelsolar/todo_cli/pkg/config"
	"github.com/rs/zerolog"
)

const PKG_NAME = "Logger PKG"
const PKG_VERSION = "1.1.1"

var log *Logger
var cfg *config.Config

type Logger struct {
	*zerolog.Logger
}

// NewLogger creates a new Logger instance with the provided zerolog.Logger.
//
// Parameters:
// - logger: A pointer to a zerolog.Logger instance.
//
// Returns:
// - A pointer to the newly created Logger instance.
func NewLogger(pkgName string, pkgVersion string) *Logger {

	cfg = config.GetConfig()

	logger := ConfigLogger(cfg.LogLevel, pkgName, pkgVersion)
	log = &Logger{&logger}
	return log
}

func ConfigLogger(logLevel string, pkgName string, pkgVersion string) zerolog.Logger {
	var zlogLevel zerolog.Level
	switch logLevel {
    case "trace":
        zlogLevel = zerolog.TraceLevel
	case "debug":
		zlogLevel = zerolog.DebugLevel
	case "info":
		zlogLevel = zerolog.InfoLevel
	case "warn":
		zlogLevel = zerolog.WarnLevel
	case "error":
		zlogLevel = zerolog.ErrorLevel
	default:
		zlogLevel = zerolog.DebugLevel
	}

	zerolog.SetGlobalLevel(zlogLevel)
	zl := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).
		With().
		Timestamp().
		Str("app", fmt.Sprintf("%s v%s", pkgName, pkgVersion)).
		Logger()

	return zl
}

// GetLogger returns a pointer to the Logger instance.
//
// Returns:
// - *Logger: A pointer to the Logger instance.
func GetLogger() *Logger {
	return log
}
