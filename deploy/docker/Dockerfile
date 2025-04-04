# syntax=docker/dockerfile:1.4

# Build Stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Install git for module fetching
RUN apk add --no-cache git

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o alerts ./cmd/api/main.go

# Runtime Stage
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/alerts .

EXPOSE 8080

CMD ["./alerts"]
