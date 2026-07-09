FROM golang:1.26.4-alpine AS builder
WORKDIR /app
COPY go.mod .
RUN go mod download
COPY . .
ARG VERSION=dev
ARG COMMIT=local
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-X main.version=${VERSION} -X main.commit=${COMMIT}" -o server main.go



FROM alpine:3.20
RUN addgroup -S -g 10001 appgroup && adduser -S -D -H -u 10001 -G appgroup appuser
WORKDIR /app
COPY --from=builder --chown=appuser:appgroup /app/server /server
USER appuser
EXPOSE 3000
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD wget -qO- http://localhost:3000/ready || exit 1
ENTRYPOINT ["/server"]
