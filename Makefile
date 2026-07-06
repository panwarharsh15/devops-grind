APP_NAME=devops-grind
BINARY_NAME=server
CONTAINER_NAME=go-http-container
PORT=3000
IMAGE=devops-grind

.PHONY: help all run build test test-v fmt fmt-check clean docker-build docker-run docker-stop docker-logs docker-clean compose-up compose-down compose-logs compose-ps compose-rebuild

help:
	@echo "Available commands:"
	@echo "  make all              - Format check, test, and build app"
	@echo "  make run              - Run Go app locally"
	@echo "  make build            - Build Go binary"
	@echo "  make test             - Run Go tests"
	@echo "  make test-v           - Run Go tests with verbose output"
	@echo "  make fmt              - Format Go code"
	@echo "  make fmt-check        - Check Go formatting"
	@echo "  make clean            - Remove local binary"
	@echo "  make docker-build     - Run tests and build Docker image"
	@echo "  make docker-run       - Build and run Docker container"
	@echo "  make docker-stop      - Stop and remove Docker container"
	@echo "  make docker-logs      - Show Docker container logs"
	@echo "  make docker-clean     - Remove Docker image"
	@echo "  make compose-up       - Start app using Docker Compose"
	@echo "  make compose-rebuild  - Rebuild and start app using Docker Compose"
	@echo "  make compose-down     - Stop Docker Compose services"
	@echo "  make compose-logs     - Show Docker Compose logs"
	@echo "  make compose-ps       - Show Docker Compose services"

all: fmt-check test build

run:
	go run main.go

build:
	go build -o $(BINARY_NAME) main.go

test:
	go test ./...

test-v:
	go test -v ./...

fmt:
	gofmt -w .

fmt-check:
	test -z "$$(gofmt -l .)"

clean:
	rm -f $(BINARY_NAME)

docker-build: test
	docker build -t $(IMAGE) .

docker-run: docker-build docker-stop
	docker run -d -p $(PORT):$(PORT) --name $(CONTAINER_NAME) $(IMAGE)

docker-stop:
	docker stop $(CONTAINER_NAME) || true
	docker rm $(CONTAINER_NAME) || true

docker-logs:
	docker logs -f $(CONTAINER_NAME)

docker-clean:
	docker rmi $(IMAGE) || true

compose-up:
	docker compose up -d

compose-rebuild:
	docker compose up -d --build

compose-down:
	docker compose down

compose-logs:
	docker compose logs -f

compose-ps:
	docker compose ps