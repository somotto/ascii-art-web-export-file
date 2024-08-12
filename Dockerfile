# Use the official Go image as a parent image
FROM golang:1.22-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod ./

# Download the Go module dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Use a minimal alpine image for the final stage
FROM alpine:3.18

# Set the working directory
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Copy the necessary txt files and static assets
COPY *.txt ./
COPY templates ./templates
COPY static ./static

# Metadata labels
LABEL maintainer="somotto <stephenonyango4783@gmail.com>"
LABEL version="1.0"
LABEL description="ASCII Art Web Dockerize"

# Expose the port the app runs on
EXPOSE 8000

# Command to run the executable
CMD ["./main"]