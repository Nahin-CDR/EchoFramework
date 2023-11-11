# Use the official Golang image as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the necessary files into the container
COPY . .

# Build the Go application
RUN go build main.go

# Expose the port on which the application will run
EXPOSE 8080

# Command to run the application
CMD ["./main"]
