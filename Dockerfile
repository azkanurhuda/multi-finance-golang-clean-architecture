# Use the official Golang image as the base image
FROM golang:1.20

# Set the working directory inside the container
WORKDIR /app

# Copy the local package files to the container's workspace
COPY . .

# Download and install any dependencies
RUN go mod tidy && go mod download

# Expose port 3000 to the outside world
EXPOSE 3000

# Command to run the application using 'go run'
CMD ["go", "run", "cmd/web/main.go"]
