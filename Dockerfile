# Use the official Golang image as a builder
FROM golang:1.23.0-alpine as builder

# Set the current working directory inside the container
WORKDIR /app

# Copy the Go modules manifests
COPY go.mod ./
COPY go.sum ./

# Download the Go modules
RUN go mod download

# Copy the source code into the container
COPY ./src ./src

# Build the Go application
RUN go build -o main ./src/main.go

# Use a smaller base image to run the application
FROM alpine:latest

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Command to run the application
CMD ["./main"]
