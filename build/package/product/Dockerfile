# Stage 1: сборка
FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY ../product/go.mod ../product/go.sum ./
RUN go mod download
COPY ../product .
RUN go build -o product cmd/main.go

# Stage 2: запуск
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/product .
EXPOSE 8081
ENTRYPOINT ["./product"]