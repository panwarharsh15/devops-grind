APP_NAME=devops-grind
BINARY_NAME=server
CONTAINER_NAME=go-http-container
PORT=3000
IMAGE=devops-grind:latest
run:
	go run main.go
build:
	go build -o $(BINARY_NAME) main.go
docker-build:
	docker build -t $(IMAGE) .
docker-run:	docker-build 
	docker run -d -p $(PORT):$(PORT) --name $(CONTAINER_NAME) $(IMAGE)
