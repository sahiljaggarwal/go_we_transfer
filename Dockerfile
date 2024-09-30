# Use the official Golang image as a builder
FROM golang:1.23.0-alpine as builder

# Set the current working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files for dependency management
COPY go.mod go.sum ./

# Download all the dependencies
RUN go mod download

# Copy the entire source code into the container
COPY src/ ./src/

# Build the Go application (no .exe extension, just binary)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./src/main.go

# Start a new stage from scratch
FROM alpine:latest

# Set the working directory
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Expose port 8080 (or any other port your application uses)
EXPOSE 8080

# Command to run the application
CMD ["./main"]