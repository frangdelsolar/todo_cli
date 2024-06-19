# Data PKG v1.1.4

Data PKG is a Go package that simplifies database interaction for your application. It provides a lightweight wrapper around the popular GORM library (https://gorm.io/docs/query.html) specifically for SQLite databases.

## Features

-   Leverages the power of GORM for efficient database operations.
-   Centralized database connection management.
-   Logging integration for connection status and errors.
-   Easy database initialization with optional file path configuration.

## Installation


### 1. Using `go get`:

```bash
    go get -u github.com/frangdelsolar/todo_cli/pkg/data
```

### 2. Using `go mod`:

-   Add the following line to your `go.mod` file, replacing `<VERSION>` with the desired version:

```
require github.com/frangdelsolar/todo_cli/pkg/data <VERSION>
```

## Usage

### 1. Configure Database Path on your environment
- Set the path to your SQlite database in your environment file, like so:
```
    DB_PATH="dev.db"
```
NOTE: Refer to `config` package (documentation)[https://github.com/frangdelsolar/todo_cli/blob/master/pkg/config/README.md] to get instructions on how to approach this.

### 2. Initialize Database

Call the `data.LoadDB` function to establish a connection to your SQLite database.
```go
    db, err := data.LoadDB()
    if err != nil {
        // Handle error
    }
```
This will return a reference to the database object.

### 2. Access Database Connection:

Use the `data.GetDB` function to retrieve the underlying database instance for further data manipulation. 
```go
    db, err := data.GetDB()
    if err != nil {
        // Handle error
    }
```

### 3. Perform Database Operations:

Use the standard GORM API methods on the retrieved gormDB instance to interact with your database tables. Refer to the GORM documentation for details: https://gorm.io/docs/index.html

### 4. Declaring models
This package offers a structure to keep track of system data like: createdBy and updatedBy as well as some usefull info stored by gorm.Model.

```go
    // example model
    type Example struct {
        data.SystemData
        Field1 string
        Field2 string
    }
```

```go
    // System Data
    type SystemData struct {
        gorm.Model
        CreatedBy   *auth.User
        CreatedByID uint
        UpdatedBy   *auth.User
        UpdatedByID uint
    }
```

```go
    // gorm.Model definition
    type Model struct {
        ID        uint           `gorm:"primaryKey"`
        CreatedAt time.Time
        UpdatedAt time.Time
        DeletedAt gorm.DeletedAt `gorm:"index"`
    }
```

Your model will become:
```go
    type Example struct {
        ID        uint           `gorm:"primaryKey"`
        CreatedAt time.Time
        UpdatedAt time.Time
        DeletedAt gorm.DeletedAt `gorm:"index"`
        CreatedBy   *auth.User
        CreatedByID uint
        UpdatedBy   *auth.User
        UpdatedByID uint
        Field1 string
        Field2 string
    }
```

