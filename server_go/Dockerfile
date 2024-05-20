# Stage 1: Build the application
FROM golang:1.22.3-bookworm AS builder

# Enable CGO
ENV CGO_ENABLED=1

# Install build dependencies
RUN apt-get update \
 && DEBIAN_FRONTEND=noninteractive \
  apt-get install --no-install-recommends --assume-yes \
    build-essential \
    libsqlite3-dev \
    ca-certificates \
    openssl

# Set the Current Working Directory inside the container
WORKDIR /src

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o ./server main.go

# Stage 2: Create the runtime image
FROM debian:bookworm

# Install runtime dependencies
RUN apt-get update \
 && DEBIAN_FRONTEND=noninteractive \
  apt-get install --no-install-recommends --assume-yes \
    libsqlite3-0 \
    ca-certificates \
    openssl

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /src/server /usr/bin/server

COPY --from=builder /src/server /app

### copy all other files to the /app director from Builder
COPY --from=builder /src /app

# Expose port 3000
EXPOSE 3000

# Command to run the executable
CMD ["./server"]
