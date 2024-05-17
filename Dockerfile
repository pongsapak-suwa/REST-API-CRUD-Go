# Use the official golang image as base
FROM golang:1.16 AS builder

# Set the current working directory inside the container
WORKDIR /app

# Copy go modules manifests
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/main.go

# Start a new stage from scratch
FROM alpine:latest  

# Set the working directory to /app
WORKDIR /app

# Copy the pre-built binary from the previous stage
COPY --from=builder /app/app .

# Expose port 3000 to the outside world
EXPOSE 3000

# Command to run the executable
CMD ["./app"]
