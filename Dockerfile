# Dockerfile

# Start from the latest Golang base image
FROM golang:1.24.0

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum first to cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go app from the `cmd/main.go` entry point
RUN go build -o server ./cmd

# Expose port (Railway uses PORT env)
EXPOSE 8080

# Run the compiled binary
CMD ["./server"]