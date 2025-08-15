# Stage 1: Build the Go binary
FROM golang:1.23 AS builder

WORKDIR /app

# Copy go.mod and go.sum first for caching
COPY go.mod go.sum ./
RUN go mod download

# Copy all source code
COPY . .

# Build the binary explicitly, output to /app/main
RUN go build -o main ./stdlib-basic.go

# Stage 2: Small runtime image
FROM debian:bookworm-slim

WORKDIR /app

# Copy binary from builder stage
COPY --from=builder /app/main /app/main

# Expose API port
EXPOSE 8080

# Run the binary
CMD ["./main"]
