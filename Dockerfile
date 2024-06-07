# Define the base image with a specific Go version
FROM golang:1.22

# Set the working directory inside the container
WORKDIR /app

# Download Go dependencies (assuming go modules)
COPY go.mod go.sum ./

# Copy your Go source code
COPY . .

# Build your Go application
RUN go build -o main ./cmd/main.go

# Set the command to run your application
CMD ["./main"]