.PHONY: test
test:
	cp .env.test auth/.env.test 
	cp .env.test cli/.env.test 
	cp .env.test contractor/.env.test 
	cp .env.test currency/.env.test 
	go test -coverprofile=coverage.out -covermode=set $(go list ./...) -v ./...
	go tool cover -html=coverage.out

.PHONY: clean
clean:
	find . -name "coverage.out" -delete
	find . -name "*.log" -delete
	find . -name "test.db" -delete
