# Currency PKG v1.0.2

This package provides functionalities for managing currencies and their exchange rates, accounts, and transactions within your Go application. It offers a set of models and methods for:

-   **Currencies**: Represents different currencies with associated exchange rates.
-   **Accounts**: Stores details about financial accounts, including their balance and associated currency.
-   **Transactions**: Tracks financial transactions with amounts, dates, and types (credit/debit).

## Installation:

Use the `go get` command to install the package:

```bash
    go get github.com/frangdelsolar/todo_cli/pkg/currency
```

## Initialization:

**Important**: The currency package requires explicit initialization before use. This step establishes connections, performs database migrations for schema management, and configures the logger.

1. Import the InitCurrency function:

```go
    import (
        "github.com/frangdelsolar/todo_cli/pkg/currency"
    )
```

Call `InitCurrency` at the beginning of your application's main function or a dedicated initialization function:

```go
    func main() {
        currency.InitCurrency()
        // ... your application code using currency package methods
    }
```

## API Documentation:

The detailed API documentation for each function is available within the package source code. Comments within the code provide a quick overview of functionalities.

## Example Usage:

### Creating a Currency:

```go
    import (
        "fmt"

        "github.com/frangdelsolar/todo_cli/pkg/currency"
        "github.com/frangdelsolar/todo_cli/pkg/currency/models"
    )

    func main() {
        currency.InitCurrency()

        newCurrency, err := currency.CreateCurrency("USD", "1000.00", time.Now().Format(time.DateOnly))
        if err != nil {
            fmt.Println("Error creating currency:", err)
            return
        }

        fmt.Printf("Successfully created currency: %s\n", newCurrency.Code)
    }
```

### Creating an Account:

```go
    import (
        "fmt"

        "github.com/frangdelsolar/todo_cli/pkg/currency"
        "github.com/frangdelsolar/todo_cli/pkg/currency/models"
    )

    func main() {
        currency.InitCurrency()

        newAccount, err := currency.CreateAccount("My Savings Account", "1000.00", "USD", true)
        if err != nil {
            fmt.Println("Error creating account:", err)
            return
        }

        fmt.Printf("Successfully created account: %s (ID: %s)\n", newAccount.Name, newAccount.ID)
    }
```

### Updating Account Balance:

```go
    import (
        "fmt"

        "github.com/frangdelsolar/todo_cli/pkg/currency"
        "github.com/frangdelsolar/todo_cli/pkg/currency/models"
    )

    func main() {
        currency.InitCurrency()

        transaction, err := currency.UpdateAccountBalance("acc_id", "USD", "100.00", time.Now().Format(time.DateOnly), "Deposit", "credit")
        if err != nil {
            fmt.Println("Error updating account balance:", err)
            return
        }

        fmt.Printf("Successfully updated account balance. Transaction ID: %s\n", transaction.ID)
    }
```
