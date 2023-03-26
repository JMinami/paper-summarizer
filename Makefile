.PHONY: test tffmt docker-build docker-clean

test:
	@echo "Running unit tests..."
	@cd ./src && go test -v ./... -count=1
	@echo "All tests passed successfully!"

tffmt:
	@echo "Run tf fmt"
	@terraform fmt -recursive

include .env

# イメージ名とタグを定義
IMAGE_NAME := golang-web-server
IMAGE_TAG := latest

docker-build:
	docker build -t $(IMAGE_NAME):$(IMAGE_TAG) \
	-f ./docker/Dockerfile.dev \
	--build-arg LINE_MESSAGING_API_CHANNEL_ACCESS_TOKEN=${LINE_MESSAGING_API_CHANNEL_ACCESS_TOKEN} \
	--build-arg LINE_MESSAGING_API_CHANNEL_SECRET=${LINE_MESSAGING_API_CHANNEL_SECRET} \
	.

docker-run:
	docker run -p 8080:8080 $(IMAGE_NAME):$(IMAGE_TAG)

docker-clean:
	docker rmi $(IMAGE_NAME):$(IMAGE_TAG)