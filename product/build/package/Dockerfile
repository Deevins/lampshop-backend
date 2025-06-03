FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o product-service ./cmd/product/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/product-service .

EXPOSE 8081
ENTRYPOINT ["./product-service"]