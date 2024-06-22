# Auth PKG v1.0.8

This package offers a set of functions for managing user authentication in your Go applications. It provides functionalities for creating, retrieving, updating, and deleting users from a database. Additionally, it includes basic user validation for name and email.

## Features

-   Manages user registration, retrieval, update, and deletion.
-   Connects to a database for user storage (requires gorm.io/gorm package).
-   Validates user name and email format.

## Installation

1. Using `go get`

```bash
    go get -u github.com/frangdelsolar/todo_cli/pkg/auth
```

2. Using `go.mod`

-   Add the following line to your `go.mod` file, replacing `<VERSION>` with the desired version.

```
    require github.com/frangdelsolar/todo_cli/pkg/auth <VERSION>
```

## Usage

1. **Import the package**

```go
    import (
        "fmt"

        "github.com/your-username/your-project/pkg/auth"
    )
```

2. **Initialize the package**

```go
    err := auth.InitAuth()
    if err != nil {
        panic(err)
    }
```

NOTE: Make sure you have configured your database connection before calling `InitAuth`. See (Data PKG Documentation)[https://github.com/frangdelsolar/todo_cli/blob/master/pkg/data/README.md]

3. **Create a new user**:

```go
    newUser, err := auth.NewUser("John Doe", "john.doe@email.com")
    if err != nil {
        // handle error
        fmt.Println("Error creating user:", err)
    } else {
        // user created successfully
        fmt.Println("User created:", newUser.ID)
    }
```

Similar for the rest of the methods listed in `data.go`. You'll see documentation there as well.
