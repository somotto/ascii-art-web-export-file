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

# Metadata labels
LABEL maintainer="stephenonyango4783@gmail.com, wendytbt4@gmail.com, apikojuma94@gmail.com"
LABEL version="1.0"
LABEL description="ASCII Art Web Dockerize"

# Expose the port the app runs on
EXPOSE 8000

# Command to run the executable
CMD ["./main"]