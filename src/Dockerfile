# Use the official Go image to build the application
FROM golang:1.23

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules manifests (fixed)
COPY go.mod go.sum .

# Download Go modules
RUN go mod download

# Copy the source code
COPY . .

# Build the Go application
RUN go build -o main .

# Expose the port the app runs on
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
