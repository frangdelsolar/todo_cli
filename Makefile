.PHONY: run
run:
	APP_ENV=dev go run cmd/main.go 

.PHONY: serve
serve:
	APP_ENV=dev go run server/server.go

.PHONY: clean
clean:
	find . -name "coverage.out" -delete
	find . -name "*.log" -delete
	find . -name "*.db" -delete
	find . -name "rates.json" -delete

.PHONY: test
test:
	cp .env.test pkg/test/auth/.env.test 
	cp .env.test pkg/test/cli/.env.test 
	cp .env.test pkg/test/contractor/.env.test 
	cp .env.test pkg/test/currency/.env.test 
	export APP_ENV=test
	go test -coverprofile=coverage.out -covermode=set $(go list ./...) -v ./...
	go tool cover -html=coverage.out
