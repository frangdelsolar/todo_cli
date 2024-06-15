# Currency PKG v1.0.0

This package provides functionalities for working with currencies, including:

-   Downloading currency exchange rates from an external API.
-   Representing exchange rates for different sources (Official, Blue).
-   Converting currency amounts based on exchange rates and dates. It uses an average between buy and sell of Blue rate.

## Installation

```bash
    go get -u github.com/your-username/currency
```

## Usage

### Downloading Rates

```go
    package main

    import (
        "fmt"

        "github.com/frangdelsolar/todo_cli/pkg/currency"
    )

    func main() {
        err := currency.DownloadRates()
        if err != nil {
            fmt.Println("Error downloading rates:", err)
            return
        }
        fmt.Println("Downloaded currency rates successfully!")
    }
```

**Note**: It will store the file `rates.json` in the folder you run this command.

### Getting Rates for a Date

```go
    package main

    import (
        "fmt"
        "time"

        "github.com/your-username/currency"
    )

    func main() {
        date, err := time.Parse("2006-01-02", "2023-06-10") // Replace with desired date
        if err != nil {
            fmt.Println("Error parsing date:", err)
            return
        }
        rates, err := currency.GetRatesByDate(date)
        if err != nil {
            fmt.Println("Error getting rates:", err)
            return
        }
        fmt.Printf("Exchange rates for %s:\n", date.Format("2006-01-02"))
        fmt.Printf("  Official: Buy %.2f, Sell %.2f\n", rates.Official.ValueBuy, rates.Official.ValueSell)
        fmt.Printf("  Blue: Buy %.2f, Sell %.2f\n", rates.Blue.ValueBuy, rates.Blue.ValueSell)
    }
```
