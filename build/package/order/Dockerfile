# Stage 1: сборка
FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY ../order/go.mod ../order/go.sum ./
RUN go mod download
COPY ../order .
RUN go build -o order cmd/main.go

# Stage 2: запуск
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/order .
EXPOSE 8082
ENTRYPOINT ["./order"]