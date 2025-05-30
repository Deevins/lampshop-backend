# Stage 1: build
FROM golang:1.23-alpine AS builder
WORKDIR /app

COPY product/go.mod product/go.sum ./product/
WORKDIR /app/product
RUN go mod download

WORKDIR /app
COPY product/ ./product/

RUN CGO_ENABLED=0 GOOS=linux go build -o product-service ./product/cmd/product/main.go

# Stage 2: runtime
FROM alpine:latest
WORKDIR /app

# Копируем скомпилированный бинарник
COPY --from=builder /app/product-service ./product-service

EXPOSE 8081

ENTRYPOINT ["./product-service"]
