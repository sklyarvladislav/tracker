# Justfile for Habit Tracker

# Default recipe to display help information
default:
    @just --list

# Install dependencies
install:
    @echo "Installing backend dependencies..."
    go mod download
    @echo "Installing frontend dependencies..."
    cd frontend && npm install

# Run database migrations
migrate:
    @echo "Running database migrations..."
    @DATABASE_PATH=./tracker.db go run . &
    @sleep 2
    @pkill -f "go run"
    @echo "Migrations completed"

# Build the backend
build-backend:
    @echo "Building backend..."
    go build -o tracker-backend .

# Build the frontend
build-frontend:
    @echo "Building frontend..."
    cd frontend && npm run build

# Build both frontend and backend
build: build-frontend build-backend
    @echo "Build completed"

# Run the backend server (development)
dev-backend:
    @echo "Starting backend server..."
    go run .

# Run the frontend dev server
dev-frontend:
    @echo "Starting frontend dev server..."
    cd frontend && npm run dev

# Run both backend and frontend in development mode
dev:
    @echo "Starting development servers..."
    @just dev-backend &
    @just dev-frontend

# Build and run with Docker
docker-build:
    @echo "Building Docker image..."
    docker build -t habit-tracker .

# Run Docker container
docker-run:
    @echo "Running Docker container..."
    docker run -p 8080:8080 -v $(pwd)/data:/data habit-tracker

# Run with docker-compose
compose-up:
    @echo "Starting with docker-compose..."
    docker-compose up -d

# Stop docker-compose
compose-down:
    @echo "Stopping docker-compose..."
    docker-compose down

# Clean build artifacts
clean:
    @echo "Cleaning build artifacts..."
    rm -f tracker-backend
    rm -f tracker.db
    rm -rf frontend/dist
    rm -rf frontend/node_modules

# Run tests
test:
    @echo "Running tests..."
    go test ./...

# Format code
fmt:
    @echo "Formatting Go code..."
    go fmt ./...
    @echo "Formatting frontend code..."
    cd frontend && npm run format || true

# Lint code
lint:
    @echo "Linting Go code..."
    golangci-lint run || true
    @echo "Linting frontend code..."
    cd frontend && npm run lint || true
