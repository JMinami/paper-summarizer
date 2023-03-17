.PHONY: test

test:
	@echo "Running unit tests..."
	@cd ./src && go test -v ./...
	@echo "All tests passed successfully!"