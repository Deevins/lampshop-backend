# Stage 1: Build
FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# 👇 правильный путь к main.go
RUN CGO_ENABLED=0 GOOS=linux go build -o order-service ./cmd/order/main.go

# Stage 2: Runtime
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/order-service .

EXPOSE 8082
ENTRYPOINT ["./order-service"]