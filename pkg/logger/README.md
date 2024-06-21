# Logger v1.1.2

The `logger` package provides a wrapper around the popular zerolog logging library (https://github.com/rs/zerolog) with additional context specific to your TODO application. It provides a wrapper around zerolog, adding context specific to your application and a simple interface for managing the logger instance.

## Features

-   Leverages zerolog for structured logging.
-   Adds context for your TODO application (e.g., app name, version).
-   Provides functions to create, configure, and retrieve the logger instance.
-   Integrates with your configuration package (`config`) to retrieve the log level from environment variables.
-   Writes logs to a file named `<pkgName>.log` or `default.log` in a dedicated logs directory.
-   Offers a console-like format for log files, including timestamps and caller information.

## Installation

### 1. Using `go get`:

```bash
    go get -u github.com/frangdelsolar/todo_cli/pkg/logger
```

### 2. Using `go mod`:

-   Add the following line to your `go.mod` file, replacing `<VERSION>` with the desired version:

```
require github.com/frangdelsolar/todo_cli/pkg/logger <VERSION>
```

## Usage

### 1. Configure Logging Level on your environment

-   Set the desired log level (e.g., "trace", "debug", "info", "warn", "error") using an environment variable.
-   Refer to `config` package (documentation)[https://github.com/frangdelsolar/todo_cli/blob/master/pkg/config/README.md] to get instructions on how to approach this.

### 2. Create the logger instance

```go
    import (
        "github.com/frangdelsolar/todo_cli/pkg/logger"
    )

    var APP_NAME = "TODO APP"
    var APP_VERSION = "1.1.0"

    func main() {

        cfg, err := config.Load()
        if err != nil {
            panic(err)
        }

        log := logger.NewLogger(logger.LoggerConfig{
            PackageName: APP_NAME,
            PackageVersion: APP_VERSION,
            LogLevel: cfg.LogLevel, // this can either be define here or in the env file.
        })

        log.Info().Msgf("Running %s environment", cfg.AppEnv)
        log.Info().Msg("Running TODO APP v" + APP_VERSION)
    }
```

**Output example**

```shell
    17:31:37 INF cmd/main.go:23 > Running local environment app="TODO APP v1.0.4"
    17:31:38 INF cmd/main.go:24 > Running TODO APP v1.0.4 app="TODO APP v1.0.4"
```

### 3. Get the Logger:

-   Use the `GetLogger` function to retrieve the currently set logger instance:

```go
    appLogger := logger.GetLogger()
    appLogger.Info().Msg("Logging with retrieved logger instance")
```
Note: if there is no logger already initialized, it will return a default instance.

## Documentation

-   Refer to the [zerolog documentation](https://github.com/rs/zerolog) for more advanced logging features.
