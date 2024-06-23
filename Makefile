# test:
# 	@echo "Running tests..."
# 	go test -coverprofile=coverage.out -covermode=set $(go list ./...)
# 	go tool cover -html=coverage.out
#     go test -v ./...
#     export APP_ENV=test && go test ./...

.PHONY: test

test:
	cp .env.test pkg/test/auth_test/.env.test 
	# go test -v ./...
	go test -coverprofile=coverage.out -covermode=set $(go list ./...) -v ./...
	go tool cover -html=coverage.out

clean:
	find . -name "*.log" -delete
	find . -name "test.db" -delete
