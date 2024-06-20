# Contractor PKG v1.0.0

## Features

...

## Installation

1. Using `go get`

```bash
    go get -u github.com/frangdelsolar/todo_cli/pkg/contractor
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

        "github.com/your-username/your-project/pkg/contractor"
    )
```

2. **Initialize the package**

```go
    err := contractor.InitContractor()
    if err != nil {
        panic(err)
    }
```

NOTE: Make sure you have configured your database connection before calling `InitAuth`. See (Data PKG Documentation)[https://github.com/frangdelsolar/todo_cli/blob/master/pkg/data/README.md]

3. **Create a new contractor**:

```go
    newContractor, err := contractor.NewContractor("Company 1")
    if err != nil {
        // handle error
        fmt.Println("Error creating contractor:", err)
    } else {
        // user created successfully
        fmt.Println("Contractor created:", newUser.ID)
    }
```
