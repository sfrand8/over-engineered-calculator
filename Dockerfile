# Use the official Go image to build the binary
FROM golang:1.24.1 AS builder

# Set working directory inside the container
WORKDIR /app

# Copy everything to the container
COPY . .

# Ensure we build a statically linked binary (important for Alpine)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server .

# Use a minimal base image
FROM alpine:latest

# Set working directory
WORKDIR /root/

# Copy compiled binary from the builder stage
COPY --from=builder /app/server .

# Ensure the binary has execution permissions
RUN chmod +x /root/server

# Expose the port
EXPOSE 8080

# Run the server
CMD ["./server"]
