.PHONY:
test: test
	@go test -v ./...

.PHONY:
cover: cover
	@go test -cover ./...

.PHONY:
test-cover: test-cover
	@go test -v -cover ./...

.PHONY:
coverage: coverate
	@go test -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out
	@rm coverage.out
