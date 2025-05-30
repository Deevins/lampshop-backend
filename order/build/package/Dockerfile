# Stage 1: build
FROM golang:1.23-alpine AS builder
WORKDIR /app

COPY order/go.mod order/go.sum ./order/
WORKDIR /app/order
RUN go mod download

WORKDIR /app
COPY order/ ./order/

RUN CGO_ENABLED=0 GOOS=linux go build -o order-service ./order/cmd/order/main.go

# Stage 2: runtime
FROM alpine:latest
WORKDIR /app

# Копируем скомпилированный бинарник
COPY --from=builder /app/order-service ./order-service

EXPOSE 8081

ENTRYPOINT ["./order-service"]
