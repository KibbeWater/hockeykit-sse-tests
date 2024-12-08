# Use the official Golang image as base
FROM golang:1.21-alpine

# Set working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o main ./cmd/main.go

# Make the scenarios directory a volume
VOLUME /app/tests/scenarios

# Expose port 8080
EXPOSE 8080

# Run the binary
CMD ["./main"]