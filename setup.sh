#!/bin/bash
set -e

echo "Setting up Go CRUD API project..."

# Initialize go.mod if it doesn't exist
if [ ! -f go.mod ]; then
    echo "Initializing Go module..."
    go mod init github.com/yourusername/go-crud-api
fi

# Download dependencies
echo "Downloading dependencies..."
go get -u github.com/gin-gonic/gin
go get -u github.com/google/uuid
go get -u github.com/joho/godotenv

# Create directory structure if it doesn't exist
mkdir -p cmd/api
mkdir -p internal/{handler,middleware,model,repository,service}
mkdir -p pkg/utils
mkdir -p config
mkdir -p docs/swagger

echo "Setup complete!"
echo "Run 'go run cmd/api/main.go' to start the server." 