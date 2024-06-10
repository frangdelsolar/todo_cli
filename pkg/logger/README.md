# Logger v1.0.2

The `logger` package provides a wrapper around the popular zerolog logging library (https://github.com/rs/zerolog) with additional context specific to your TODO application. It offers a simple interface for setting and retrieving a configured zerolog logger instance.

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

### 1. Configure Zerolog (Outside the `logger` package):

-   Configure your zerolog instance with desired output, fields, and log level. Here's an example:

```go
    package main

    import (
        "fmt"
        "os"

        "github.com/frangdelsolar/todo_cli/pkg/logger"
        "github.com/rs/zerolog"
    )

    var APP_VERSION = "1.0.0"

    func main() {
        var logLevel = zerolog.InfoLevel // Adjust log level as needed

        zerolog.SetGlobalLevel(logLevel)
        yourConfiguredZerologLogger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).
            With().
            Timestamp().
            Str("app", fmt.Sprintf("TODO APP v%s", APP_VERSION)).
            Logger()

        // Use this configured logger instance throughout your application
        // (don't configure within the logger package itself)
    }
```

### 2. Set the Logger (Optional):

-   If you have a pre-configured zerolog instance, you can set it using the `SetLogger` function:

```go
    customLogger := logger.SetLogger(&yourConfiguredZerologLogger)
```

### 3. Get the Logger:

-   Use the `GetLogger` function to retrieve the currently set logger instance:

```go
    appLogger := logger.GetLogger()
    appLogger.Info().Msg("Logging with retrieved logger instance")
```

## Documentation

-   Refer to the [zerolog documentation](https://github.com/rs/zerolog) for more advanced logging features.
