#use an official Golang runtime as the base image
FROM golang:1.16

# Set the working directory inside the container
WORKDIR /app

# Copy the local source files to the container's working directory
COPY ./ .

# Build the Go application
RUN go build -o myapp

# Define environment variables for sourceLocation and outputLocation
ENV SOURCE_LOCATION=""
ENV OUTPUT_LOCATION=""

# Define the command to run your application when the container starts
CMD ["./main"]