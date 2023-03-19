.PHONY: test, tffmt

test:
	@echo "Running unit tests..."
	@cd ./src && go test -v ./... -count=1
	@echo "All tests passed successfully!"

tffmt:
	@echo "Run tf fmt"
	@terraform fmt -recursive