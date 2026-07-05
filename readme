# DevOps Grind - Go HTTP Server

A simple Go-based HTTP server created for learning Docker, multi-stage Docker builds, and GitHub Actions CI/CD.

---

## Overview

This project is a beginner-friendly Go HTTP server.
The main purpose of this repository is to practice:

* Building a basic Go web server
* Creating REST API endpoints
* Writing a Dockerfile
* Using multi-stage Docker builds
* Building and pushing Docker images using GitHub Actions
* Understanding the DevOps workflow from code to container image

---

## Tech Stack

* **Language:** Go
* **Containerization:** Docker
* **CI/CD:** GitHub Actions
* **Container Registry:** GitHub Container Registry, GHCR

---

## Features

* Health check API endpoint
* Hello API endpoint with query parameter
* Echo API endpoint for POST JSON requests
* Dockerized Go application
* GitHub Actions workflow for quality check and Docker image push
* Lightweight final Docker image using Alpine Linux

---

## Project Structure

```bash
.
├── .github/
│   └── workflows/
│       └── docker-publish.yml
├── Dockerfile
├── go.mod
├── main.go
└── readme
```

---

## API Endpoints

| Method | Endpoint            | Description                                  |
| ------ | ------------------- | -------------------------------------------- |
| GET    | `/health`           | Checks if the server is running              |
| GET    | `/hello?name=Harsh` | Returns a hello message                      |
| POST   | `/echo`             | Accepts JSON input and returns the same data |

---

## Prerequisites

Before running this project, make sure these tools are installed:

* Git
* Go 1.22 or above
* Docker

---

## Run Locally

### 1. Clone the repository

```bash
git clone https://github.com/panwarharsh15/devops-grind.git
cd devops-grind
```

### 2. Run the Go application

```bash
go run main.go
```

By default, the application runs on port `3000`.

---

## Test the APIs

### Health check

```bash
curl http://localhost:3000/health
```

Expected output:

```json
{
  "status": "ok",
  "message": "server is alive",
  "timestamp": "2026-07-06T00:00:00Z"
}
```

### Hello API

```bash
curl "http://localhost:3000/hello?name=Harsh"
```

Expected output:

```json
{
  "status": "ok",
  "message": "hello, Harsh!",
  "timestamp": "2026-07-06T00:00:00Z"
}
```

### Echo API

```bash
curl -X POST http://localhost:3000/echo \
  -H "Content-Type: application/json" \
  -d '{"name":"Harsh","role":"DevOps Engineer"}'
```

Expected output:

```json
{
  "status": "ok",
  "message": "echo",
  "data": {
    "name": "Harsh",
    "role": "DevOps Engineer"
  },
  "timestamp": "2026-07-06T00:00:00Z"
}
```

---

## Run with Docker

### 1. Build Docker image

```bash
docker build -t devops-grind .
```

### 2. Run Docker container

```bash
docker run -p 3000:3000 devops-grind
```

### 3. Test container

```bash
curl http://localhost:3000/health
```

---

## Dockerfile Explanation

This project uses a multi-stage Docker build.

### Builder Stage

The first stage uses the official Go Alpine image to download dependencies and build the Go binary.

### Runtime Stage

The second stage uses a lightweight Alpine image and copies only the final binary from the builder stage.

This keeps the final Docker image smaller and cleaner.

---

## GitHub Actions CI/CD

This repository uses GitHub Actions for automation.

The workflow runs when code is pushed to the `main` branch.

Pipeline steps:

1. Checkout source code
2. Setup Go
3. Check Go formatting using `gofmt`
4. Run Go tests
5. Setup Docker Buildx
6. Login to GitHub Container Registry
7. Build and push Docker image to GHCR

Workflow file:

```bash
.github/workflows/docker-publish.yml
```

Docker image format:

```bash
ghcr.io/panwarharsh15/devops-grind:latest
```

---

## Useful Commands

### Check Go formatting

```bash
gofmt -w .
```

### Run tests

```bash
go test ./...
```

### Build Go binary manually

```bash
go build -o server main.go
```

### Run binary

```bash
./server
```

### Check Docker images

```bash
docker images
```

### Check running containers

```bash
docker ps
```

### Check container logs

```bash
docker logs <container_id>
```

---

## Learning Goals

Through this project, I learned:

* How to create a basic HTTP server in Go
* How to expose APIs using Go's `net/http` package
* How to build and run a Go app locally
* Difference between `go run` and `go build`
* How to write a Dockerfile
* How multi-stage Docker builds reduce image size
* How to use GitHub Actions for CI/CD
* How to push Docker images to GitHub Container Registry

---

## Future Improvements

* Add unit tests for all handlers
* Add Kubernetes deployment files
* Add Helm chart
* Add Docker image vulnerability scanning using Trivy
* Add Makefile for common commands
* Add application metrics endpoint
* Add structured logging
* Add GitHub Actions build status badge

---

## Author

**Harsh Panwar**

DevOps engineer/SRE building practical projects around Go, Docker, CI/CD, and cloud-native tools.
