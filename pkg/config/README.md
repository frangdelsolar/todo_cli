# Config PKG v1.0.0
This package provides a centralized and structured approach to managing environment variables in your Go applications. It offers functionality for loading environment variables from files and accessing them throughout your code.

## Features
- Loads environment variables from .env files.
- Supports different .env files for various environments (e.g., dev, prod).
- Provides a Config struct to hold loaded environment variables.
- Offers functions to load and access the configuration.

## Installation
1. Using `go get`
```bash
    go get -u github.com/frangdelsolar/todo_cli/pkg/config
```

2. Using `go.mod`
- Add the following line to your `go.mod` file, replacing `<VERSION>` with the desired version.
```
    require github.com/frangdelsolar/todo_cli/pkg/config <VERSION>
```

## Usage 
1. **Import the package**
```go
    import (
        "fmt"

        "github.com/your-username/your-project/pkg/config"
    )
```

2. **Load Configuration**
```go
    cfg, err := config.Load()
    if err != nil {
        panic(err)
    }
```

3. Access Environment Variables"
```go
    appEnv := cfg.AppEnv
    port := cfg.Port
    // ... access other environment variables defined in the Config struct
```

## Adding New Environment Variables

1. **Modify the Config Struct**: Add a field for the new environment variable
```go
    type Config struct {
        AppEnv     string `env:"APP_ENV"`
        Port       string `env:"PORT"`
        NewVar     string `env:"NEW_VAR"` // Adjust "NEW_VAR" to your desired name
        // ... other environment variables
    }
```
2. Update `.env` Files: Add the new variable to the appropriate `.env` files (e.g., `.env`, `.env.dev`, `.env.prod`) with its corresponding value.

### Optional: Environment-Specific Files
The package can be configured to load different `.env` files based on the `APP_ENV` environment variable.  Here's how:

1. Create separate .env files for each environment (e.g., `.env.dev`, `.env.prod`).
2. Set the `APP_ENV` variable appropiately (e.g., in a shell script):
```bash
    export APP_ENV=dev
```
3. Modify the `Load` function (within the `config` package) to conditionally load the appropriate `.env` file based on the `APP_ENV` value:
```go
    func Load() (*Config, error) {
        // ... existing logic

        env := os.Getenv("APP_ENV")

        envFile := ".env"
        if env == "dev" {
            envFile = ".env.dev"
        } else if env == "prod" {
            envFile = ".env.prod"
        }

        // ... continue with loading the appropriate environment file
    }
```