package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

var PKG_NAME = "Config PKG"
var PKG_VERSION = "1.0.2"

var config *Config

// Config represents the configuration loaded from environment variables
type Config struct {
	AppEnv   string
	LogLevel string
	DBPath   string
}

// Load reads configuration from environment variables based on the APP_ENV
func Load() (*Config, error) {

	config = &Config{}
	var err error

	env := os.Getenv("APP_ENV")

	// Determine the appropriate environment file based on APP_ENV
	envFile := ".env"
	if env == "dev" {
		envFile = ".env.dev"
	} else if env == "prod" {
		envFile = ".env.prod"
	} else if env == "test" {
		envFile = ".env.test"
	}

	// Get the current working directory
	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	filePath := filepath.Join(currentDir, envFile)

	// Load environment variables from the file
	err = godotenv.Load(filePath)
	if err != nil {
		fmt.Errorf("error loading %s file", filePath)
		return nil, err
	}

	// Set configuration values from environment variables
	config.AppEnv = env
	config.LogLevel = os.Getenv("LOG_LEVEL")
	config.DBPath = os.Getenv("DB_PATH")

	return config, nil
}

// GetConfig returns the loaded configuration
func GetConfig() *Config {
	return config
}
