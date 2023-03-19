.PHONY: test, tffmt

test:
	@echo "Running unit tests..."
	@cd ./src && go test -v ./...
	@echo "All tests passed successfully!"

tffmt:
	@echo "Run tf fmt"
	@terraform fmt -recursive