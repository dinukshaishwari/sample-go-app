# Use an official Go runtime as the base image
FROM golang:1.17

# Set the working directory inside the container
WORKDIR /go/src/app

# Copy the Go module files
COPY go.mod .
COPY go.sum .

# Download and install the application dependencies
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application
RUN go build -o app

# Expose the port used by the REST API
EXPOSE 8000:8000ss

# Run the Go REST API
CMD ["./app"]