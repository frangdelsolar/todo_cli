# Test PKG v1.0.1

This test suite verifies the core operations of the TODO app. So far supports testing for:

-   Logger
    -   Create
    -   Get
-   Data
    -   Create
    -   Get
    -   Migrations
-   Auth
    -   User creation
-   Currency
    -   Operations with currencies
    -   Account creation and balance update

## Verbosity

Supported values for log level: trace, debug, info

## Installation

### 1. Using `go get`:

```bash
    go get -u github.com/frangdelsolar/todo_cli/pkg/test
```

### 2. Using `go mod`:

-   Add the following line to your `go.mod` file, replacing `<VERSION>` with the desired version:

```
    require github.com/frangdelsolar/todo_cli/pkg/test <VERSION>
```

## Usage

### 1. Setup environment variables

Create a `.env.test` file on `test/cmd` like so:

```
   DB_PATH="test.db"
   LOG_LEVEL="trace"
```

### 2. Run Tests

1. cd `test/cmd`
2. Run

```bash
    make test
```

This will trigger the tests that you have chosen.
NOTE: if the test suite has not been yet run in your env, you will need to comment the line to remove the existing test database in Makefile. Like this:

```
    test:
        # test test.db && rm test.db
        APP_ENV=test go run .
```

You can uncomment once you have created the `test.db` for the first time.

### 3. Test suite

You're welcome to choose the tests that you want to run like this:

#### 1. Select what packages you want to run

-   Open `test/cmd/main.go`
-   Modify `main()` as you please.

```go
    func main(){
        ...
        at.RunAuthTests()
        // ct.RunCurrencyTests()
        ...
    }
```

#### 2. Select what tests you want to run for each package

-   Open `test/<package>_test/setup.go`
-   Modify `Run<Package>Tests()` function.

```go
    func RunCurrencyTests(){
        log.Info().Msg("Running Currency Tests")

        // TestAddCurrencySameCode()
        // TestAddCurrencyDifferentCode()
        // TestSubCurrencySameCode()
        // TestSubCurrencyDifferentCode()
        TestCreateAccount()
        TestUpdateAccountCredit()
    }
```
