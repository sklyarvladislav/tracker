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

# Build and run everything with docker-compose
up:
    @echo "Building and starting all services with docker-compose..."
    docker-compose up --build -d
    @echo ""
    @echo "✅ Services started successfully!"
    @echo ""
    @echo "🌐 Access the application:"
    @echo "   Frontend: http://localhost:5173"
    @echo "   Backend API: http://localhost:8080/api"
    @echo ""
    @echo "📊 View logs: docker-compose logs -f"
    @echo "🛑 Stop services: just down"

# Stop docker-compose services
down:
    @echo "Stopping docker-compose services..."
    docker-compose down

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
    @echo "Starting backend server on http://localhost:8080..."
    go run .

# Run the frontend dev server
dev-frontend:
    @echo "Starting frontend dev server on http://localhost:5173..."
    cd frontend && npm run dev

# Run both backend and frontend in development mode
dev:
    @echo "Starting development servers..."
    @echo "Backend: http://localhost:8080"
    @echo "Frontend: http://localhost:5173"
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

# Run with docker-compose (alias for up)
compose-up: up

# Stop docker-compose (alias for down)
compose-down: down

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

# View logs from docker-compose
logs:
    docker-compose logs -f
