# Stage 1: сборка
FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY ../admin/go.mod ../admin/go.sum ./
RUN go mod download
COPY ../admin .
RUN go build -o admin cmd/main.go

# Stage 2: запуск
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/admin .
EXPOSE 8083
ENTRYPOINT ["./admin"]