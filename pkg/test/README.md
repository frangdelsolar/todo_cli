# Test PKG v2.0.0

This test suite verifies the core operations of the TODO app. 

## Usage

### 1. Setup environment variables

Create a `.env.test` file on `test` like so:

```
   DB_PATH="test.db"
   FIREBASE_SECRET="secret"
   LOG_LEVEL="trace"
```

### 2. Run Tests

1. On the test root folder
2. Run

```bash
    make test
```

This will copy your `.env.test` file to the folders of your packages.
All the tests will be run, and you'll be redirected to the coverage html.
Additionally, every package will create logs within their folders, so you can check all the progress.
Once you have finished, you can clear your files by running:
```bash
    make clean
```
