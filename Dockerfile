# Below is the base image
FROM golang:latest

# Set the working directory to /go/src/app
WORKDIR /go/src/app


# Copy the current directory contents into the container at /go/src/app
COPY . .

# Build the Go app
RUN go build -o main .

# Ensure main has execute permissions
RUN chmod +x main

# Expose port 8080 to the outside world
EXPOSE 8080

# Run the executable
CMD ["./main"]
