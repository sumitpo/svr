# Start with a base image that includes Go installed
FROM golang:1.22-alpine as builder

ENV GOPROXY=https://goproxy.cn,direct

# Set the current working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum to the working directory
COPY ./app ./

# Download dependencies
RUN go mod download

# Build the Go application
RUN go build -o myapp .

# Start a new stage from scratch
FROM alpine:latest

# Set the current working directory inside the container
WORKDIR /app

# Copy the built executable from the previous stage
COPY --from=builder /app/myapp .

# Expose port 8080 (adjust as needed)
# EXPOSE 8080

# Command to run the executable
CMD ["./myapp"]
