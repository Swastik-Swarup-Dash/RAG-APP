# Start from the official Go image for building
FROM golang:1.21-alpine AS builder

# Install git (required for go mod) and ca-certificates for HTTPS
RUN apk add --no-cache git ca-certificates

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum, then download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go application (replace with your main package path if different)
RUN go build -o rag-app ./cmd/server/main.go

# Start a new, minimal image
FROM alpine:latest

# Set up certificates
RUN apk --no-cache add ca-certificates

# Set working directory
WORKDIR /root/

# Copy the built binary from the builder
COPY --from=builder /app/rag-app .

# Copy any static files if needed (optional)
# COPY ./static ./static

# Expose the port your app runs on
EXPOSE 8080

# Set environment variables (override in docker-compose or at runtime)
# ENV GEMINI_API_KEY=your_gemini_api_key
# ENV PINECONE_API_KEY=your_pinecone_api_key
# ENV PINECONE_INDEX_NAME=rag-index

# Command to run the executable
CMD ["./rag-app"]
