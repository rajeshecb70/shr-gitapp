# Stage 1: Build the Go project
FROM golang:1.23.3-alpine3.20 AS build

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum to download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the Go source code to the container
COPY main.go ./

# Build the Go project and output binary to build/bin/
RUN go build -o build/bin/go-project main.go

# Stage 2: Create the final minimal image
# FROM golang:1.23.3-alpine3.20
FROM alpine:3.20.3

# Set the working directory inside the container
WORKDIR /app

# Copy the built binary from the build stage
COPY --from=build /app/build/bin/go-project /app/go-project

# Expose port 8080
EXPOSE 8080

# Run the binary when the container starts
ENTRYPOINT ["./go-project"]
