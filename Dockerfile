# Use an official Golang image as the base image
FROM golang:1.22.4

# Set environment variables for Go
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Add CA certificates
RUN apt-get update && apt-get install -y ca-certificates && \
    update-ca-certificates && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules manifest files first (to leverage caching)
COPY go.mod go.sum ./

# Download Go module dependencies
RUN go mod download

# Copy the entire application code
COPY . .

# Copy the .env file into the container
COPY .env .env
# Change directory to the location of the `main` package before building
WORKDIR /app/cmd/api

# Build the application
RUN go build -o /app/app .

# Expose the application's port
EXPOSE 8080

# Command to run the application
CMD ["/app/app"]
