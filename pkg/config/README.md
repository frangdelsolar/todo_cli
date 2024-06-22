# Config PKG v1.0.5

This package provides a centralized and structured approach to managing environment variables in your Go applications. It offers functionality for loading environment variables from files and accessing them throughout your code, and now includes features for managing session data.

## Features

-   Loads environment variables from .env files.
-   Supports different .env files for various environments (e.g., dev, prod).
-   Provides a Config struct to hold loaded environment variables.
-   Offers functions to load and access the configuration.
-   Introduces functionalities for managing session variables.

## Installation

1. Using `go get`

```bash
    go get -u github.com/frangdelsolar/todo_cli/pkg/config
```

2. Using `go.mod`

-   Add the following line to your `go.mod` file, replacing `<VERSION>` with the desired version.

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
3. Update `SetSession()` method:
```go
    // Dump config and session to .env file
    storedKey := fmt.Sprintf("%s%s", sessionVariablesPrefix, key)
    err := godotenv.Write(map[string]string{
        "PORT": c.Port,
        "NEW_VAR": c.NewVar, <- Add this
        // ... other environment variables
        storedKey: value,
    }, c.envFile)
```
4. Update `Load()` method:
```go 
	// Load environment variables from the file
	err = godotenv.Load(filePath)
	if err != nil {
		fmt.Errorf("error loading %s file", filePath)
		return nil, err
	}

    config.Port = os.Getenv("PORT")
	config.NewVar = os.Getenv("NEW_VAR") <- Add this
    // ... other environment variables
```

## Adding New Session Variables
The `Config` struct now includes a `Session` field, which is a `map[string][string]`. You can use this map to store and retrieve session specific key-value pairs.
1. **Set Session Variables**:
- Use the `SetSession` method of the `Config` struct to set a session variable.
```go
    err := cfg.SetSession("user_id", "12345")
    if err != nil {
        // handle error
    }
```
2. **Get Session Variables**:
- Use the `GetSession` method of the `Config` struct to retrieve a session variable by its key.
```go
    userId, err := cfg.GetSession("user_id")
    if err != nil {
        // handle error (e.g., key not found)
    } else {
        fmt.Println("User ID:", userId)
    }
```
**Note**: This variables will be stored in your `.env` file with the set prefix in the `config.go` file.
```go
    const sessionVariablesPrefix = "TODO_APP_SESSION_"
```
The variable created as `user_id` will be stored as:
```
    TODO_APP_SESSION_USER_ID=1
```

### Optional: Environment-Specific Files

The package can be configured to load different `.env` files based on the `APP_ENV` environment variable. Here's how:

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
