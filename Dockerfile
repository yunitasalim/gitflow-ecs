#use an official Golang runtime as the base image
FROM golang:1.16

# Set the working directory inside the container
WORKDIR /app

RUN pwd

# Copy the local source files to the container's working directory
COPY . .

# Show the result of file copies
RUN ls

RUN pwd

# Build the Go application
RUN go build main.go

# Define environment variables for sourceLocation and outputLocation
ENV SOURCE_LOCATION=""
ENV OUTPUT_LOCATION=""

# Define the command to run your application when the container starts
CMD ["./main"]
