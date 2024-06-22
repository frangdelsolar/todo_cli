package config

import (
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
)

const (
    PKG_NAME = "Config PKG"
    PKG_VERSION = "1.0.5"

    sessionVariablesPrefix = "TODO_APP_SESSION_"
)

var config *Config

// Config represents the configuration loaded from environment variables
type Config struct {
	AppEnv   string `env:"APP_ENV" default:"dev"`
	LogLevel string `env:"LOG_LEVEL" default:"debug"`
	DBPath   string `env:"DB_PATH" default:"./data.db"`
    FirebaseSecret string `env:"FIREBASE_SECRET" default:""`
    Session map[string]string
    envFile string
}

// SetEnvFile sets the environment file based on the AppEnv field of the Config struct.
//
// The AppEnv field is used to determine the value of the envFile field.
// If the AppEnv field is "dev", the envFile field is set to ".env.dev".
// If the AppEnv field is "prod", the envFile field is set to ".env.prod".
// If the AppEnv field is "test", the envFile field is set to ".env.test".
// Otherwise, the envFile field is set to ".env".
func (c *Config) SetEnvFile() {
	c.envFile = ".env"
	if c.AppEnv == "dev" {
		c.envFile = ".env.dev"
	} else if c.AppEnv == "prod" {
		c.envFile = ".env.prod"
	} else if c.AppEnv == "test" {
		c.envFile = ".env.test"
	}
}

// SetSession sets a session variable in the Config struct and writes the updated configuration and session to the .env file.
//
// Parameters:
// - key: the key of the session variable to be set.
// - value: the value of the session variable to be set.
//
// Returns:
// - error: an error if there was a problem writing the updated configuration and session to the .env file.
func (c *Config) SetSession(key string, value string) error {
    c.Session[key] = value
    os.Setenv(key, value)

    // Dump config and session to .env file
    storedKey := fmt.Sprintf("%s%s", sessionVariablesPrefix, key)
    err := godotenv.Write(map[string]string{
        "LOG_LEVEL": c.LogLevel,
        "DB_PATH": c.DBPath,
        "FIREBASE_SECRET": c.FirebaseSecret,
        storedKey: value,
    }, c.envFile)
    
    if err != nil {
        return err
    }

    return nil
}

// GetSession retrieves a session variable from the Config struct using the provided key.
//
// Parameters:
// - key: the key of the session variable to be retrieved.
// Return types:
// - string: the value of the session variable.
// - error: an error if the key is not found in the session.
func (c *Config) GetSession(key string) (string, error) {
    // storedKey := fmt.Sprintf("%s%s", sessionVariablesPrefix, key)
    value, ok := c.Session[key]
    if !ok {
        return "", fmt.Errorf("key %s not found in session", key)
    }
    return value, nil
}

// Load loads the configuration values from environment variables and returns a Config struct and an error.
//
// It sets the configuration values from environment variables, including APP_ENV, LOG_LEVEL, DB_PATH, and FIREBASE_ADMIN_SDK.
// It also sets session variables from environment variables that start with sessionVariablesPrefix.
//
// Parameters:
//   None.
//
// Return types:
//   *Config: A pointer to the Config struct with the loaded configuration values.
//   error: An error if there was a problem loading the configuration values.
func Load() (*Config, error) {

	config = &Config{}
	var err error

    // Set configuration values from environment variables
	config.AppEnv = os.Getenv("APP_ENV")

    if config.AppEnv != "cicd" {
        config.SetEnvFile()

        // Get the current working directory
        currentDir, err := os.Getwd()
        if err != nil {
            fmt.Errorf("error getting current directory")
            return nil, err
        }

        filePath := filepath.Join(currentDir, config.envFile)

        // Load environment variables from the file
        err = godotenv.Load(filePath)
        if err != nil {
            fmt.Errorf("error loading %s file", filePath)
            return nil, err
        }
    }

    config.LogLevel = os.Getenv("LOG_LEVEL")
	config.DBPath = os.Getenv("DB_PATH")

    // decode firebase secret
    fbSecret := os.Getenv("FIREBASE_SECRET")
    decoded, err := base64.StdEncoding.DecodeString(fbSecret)
    if err != nil {
        return nil, err
    }
    config.FirebaseSecret = string(decoded)

    // Set session variables
    config.Session = make(map[string]string)
    for _, value := range os.Environ() {
        if strings.HasPrefix(value, sessionVariablesPrefix) {
            keyValue := strings.TrimPrefix(value, sessionVariablesPrefix)

            key := strings.Split(keyValue, "=")[0]
            value := strings.Split(keyValue, "=")[1]

            config.Session[key] = value
        }
    }

	return config, nil
}

// GetConfig returns the loaded configuration
func GetConfig() *Config {
	return config
}
