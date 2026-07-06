FROM golang:1.22-alpine As Builder
WORKDIR /app
COPY go.mod .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o server main.go



FROM alpine:3.20
WORKDIR /app
COPY --from=Builder /app/server /server
EXPOSE 3000
CMD ["/server"]
