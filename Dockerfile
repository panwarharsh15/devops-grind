FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY go.mod .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o server main.go



FROM alpine:3.20
RUN addgroup -S -g 10001 appgroup && adduser -S -D -H -u 10001 -G appgroup appuser
WORKDIR /app
COPY --from=builder --chown=appuser:appgroup /app/server /server
USER appuser
EXPOSE 3000
ENTRYPOINT ["/server"]
