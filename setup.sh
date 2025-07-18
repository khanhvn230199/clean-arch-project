#!/bin/bash

# Setup script for Clean Architecture Go Project
# This script will help you set up the project quickly

set -e

echo "🚀 Setting up Clean Architecture Go Project..."

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Function to print colored output
print_status() {
    echo -e "${GREEN}[INFO]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Check if Go is installed
if ! command -v go &> /dev/null; then
    print_error "Go is not installed. Please install Go 1.21 or later."
    exit 1
fi

# Check Go version
GO_VERSION=$(go version | awk '{print $3}' | sed 's/go//')
REQUIRED_VERSION="1.21"

if [ "$(printf '%s\n' "$REQUIRED_VERSION" "$GO_VERSION" | sort -V | head -n1)" != "$REQUIRED_VERSION" ]; then
    print_error "Go version $GO_VERSION is too old. Please install Go $REQUIRED_VERSION or later."
    exit 1
fi

print_status "Go version $GO_VERSION is compatible ✓"

# Create project directory structure
print_status "Creating project directory structure..."

mkdir -p cmd/server
mkdir -p internal/domain/entity
mkdir -p internal/domain/repository
mkdir -p internal/domain/service
mkdir -p internal/usecase
mkdir -p infrastructure/repository
mkdir -p infrastructure/database
mkdir -p api/handler
mkdir -p api/route
mkdir -p config
mkdir -p migrations
mkdir -p bin
mkdir -p tmp

print_status "Directory structure created ✓"

# Download dependencies
print_status "Downloading Go dependencies..."
go mod download

print_status "Dependencies downloaded ✓"

# Check if PostgreSQL is installed
if command -v psql &> /dev/null; then
    print_status "PostgreSQL is installed ✓"
    
    # Ask user if they want to create database
    read -p "Do you want to create a new database? (y/n): " -n 1 -r
    echo
    if [[ $REPLY =~ ^[Yy]$ ]]; then
        read -p "Enter database name: " DB_NAME
        read -p "Enter database user (default: postgres): " DB_USER
        DB_USER=${DB_USER:-postgres}
        
        # Create database
        print_status "Creating database '$DB_NAME'..."
        createdb -U $DB_USER $DB_NAME || print_warning "Database might already exist"
        
        # Run migrations
        print_status "Running database migrations..."
        psql -U $DB_USER -d $DB_NAME -f migrations/001_create_users_table.sql
        
        print_status "Database setup completed ✓"
        
        # Update .env file
        print_status "Updating .env file..."
        echo "DATABASE_URL=postgres://$DB_USER:password@localhost/$DB_NAME?sslmode=disable" > .env
        echo "PORT=8080" >> .env
        echo "ENVIRONMENT=development" >> .env
        
        print_status ".env file updated ✓"
    fi
else
    print_warning "PostgreSQL is not installed. Please install PostgreSQL to use the database features."
    print_status "You can also use Docker: docker-compose up -d db"
fi

# Install development tools
print_status "Installing development tools..."

# Check and install Air for hot reload
if ! command -v air &> /dev/null; then
    print_status "Installing Air for hot reload..."
    go install github.com/cosmtrek/air@latest
    print_status "Air installed ✓"
else
    print_status "Air is already installed ✓"
fi

# Check if Docker is installed
if command -v docker &> /dev/null; then
    print_status "Docker is installed ✓"
    
    read -p "Do you want to start the application with Docker? (y/n): " -n 1 -r
    echo
    if [[ $REPLY =~ ^[Yy]$ ]]; then
        print_status "Starting application with Docker Compose..."
        docker-compose up -d
        print_status "Application started! Visit http://localhost:8080/health"
        print_status "Adminer (DB admin) available at http://localhost:8081"
    fi
else
    print_warning "Docker is not installed. You can install it for easier development."
fi

print_status "Setup completed! 🎉"
echo
echo "Next steps:"
echo "1. Update the .env file with your database credentials"
echo "2. Run 'make run' to start the development server"
echo "3. Or run 'make dev' for hot reload development"
echo "4. Visit http://localhost:8080/health to check if the server is running"
echo
echo "Available commands:"
echo "  make run      - Run the application"
echo "  make dev      - Run with hot reload (requires Air)"
echo "  make test     - Run tests"
echo "  make build    - Build the application"
echo "  make clean    - Clean build artifacts"
echo
echo "API endpoints:"
echo "  GET    /health           - Health check"
echo "  POST   /api/v1/users     - Create user"
echo "  GET    /api/v1/users     - Get all users"
echo "  GET    /api/v1/users/:id - Get user by ID"
echo "  PUT    /api/v1/users/:id - Update user"
echo "  DELETE /api/v1/users/:id - Delete user"
echo
print_status "Happy coding! 🚀"