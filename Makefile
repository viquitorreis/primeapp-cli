.PHONY:
test:
	@go test -v ./...

.PHONY:
cover:
	@go test -cover ./...

.PHONY:
test-cover:
	@go test -v -cover ./...

.PHONY:
coverage:
	@go test -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out
	@rm coverage.out