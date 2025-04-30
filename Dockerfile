# Use the official Golang base image
FROM golang:1.21-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum first to leverage caching
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the Go app
RUN go build main.go

# Use a minimal base image for the final container
FROM alpine:latest

# Install certificates (if your Go app makes HTTPS calls)
RUN apk --no-cache add ca-certificates

# Set the working directory
WORKDIR /root/

# Copy the binary from the builder
COPY --from=builder /app/main .
COPY --from=builder /app/config.yaml .

# Expose the port the app runs on
EXPOSE 8010

# Command to run the executable
CMD ["./main"]
