# Use an official Golang runtime as a parent image
FROM golang:1.22.0

# Set the Current Working Directory inside the container
WORKDIR /cmd/ratelimiter

# Copy the go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source from the current directory to the working directory inside the container
COPY . .

# Command to run the tests
CMD ["go", "test", "-v", "./..."]

