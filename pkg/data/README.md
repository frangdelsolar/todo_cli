# DATA PKG v1.1.0

Data PKG is a Go package that simplifies database interaction for your application. It provides a lightweight wrapper around the popular GORM library (https://gorm.io/docs/query.html) specifically for SQLite databases.

## Features

-   Leverages the power of GORM for efficient database operations.
-   Centralized database connection management.
-   Logging integration for connection status and errors.
-   Easy database initialization with optional file path configuration.

## Installation

1. Use `go get` to install the package

```bash
    go get -u github.com/frangdelsolar/todo_cli/pkg/data
```

2. Import the package in your Go project:

```go
    import (
        "github.com/frangdelsolar/todo_cli/pkg/data"
    )
```

## Usage

### 1. Initialize Database

Call the `data.InitDB` function to establish a connection to your SQLite database. You can optionally specify a custom file path for the database file:

```go
    db, err := data.InitDB("path/to/your/database.db")
    if err != nil {
        // Handle error
    }
```

If no file path is provided, it will use the default ../data.db location.

### 2. Access Database Connection:

Use the `data.GetDB` function to retrieve the underlying GORM database instance for further data manipulation. If database has not been initialized, then it will create a new instance with default values.

```go
    db, err := data.GetDB()
    if err != nil {
        // Handle error
    }
```

### 3. Perform Database Operations:

Use the standard GORM API methods on the retrieved gormDB instance to interact with your database tables. Refer to the GORM documentation for details: https://gorm.io/docs/index.html

### 4. 

```go
    type Model struct {
        ID        uint `json:"id" gorm:"primaryKey"`
        CreatedAt time.Time
        CreatedBy auth.User
        UpdatedAt time.Time
        UpdatedBy auth.User
    }
```