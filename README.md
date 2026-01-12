# Habit Tracker

A beautiful single-page web application for tracking daily habits with a GitHub-style contribution graph.

## Features

- 🎨 **Beautiful Dark Theme UI** - Default dark theme with Zinc/Slate color palette
- 📊 **GitHub-Style Contribution Graph** - Visualize your habit completion over time
- ✨ **Interactive Dashboard** - Add, edit, and delete habits with a user-friendly interface
- 📱 **Responsive Design** - Works seamlessly on desktop and mobile devices
- 🎯 **Streak Tracking** - Track your current streak for each habit
- 🎨 **Customizable Colors** - Choose from 8 different colors for each habit
- 💾 **Persistent Storage** - SQLite database for reliable data storage

## Tech Stack

### Frontend
- **Vue 3** (Composition API)
- **Tailwind CSS** (Dark theme with Zinc variants)
- **Vite** (Build tool)

### Backend
- **Go (Golang)** with Gin framework
- **GORM** ORM
- **SQLite** database

### DevOps
- **Docker** and **Docker Compose**
- **Justfile** for task automation

## Quick Start

### Option 1: Using Docker Compose (Recommended)

The easiest way to run the entire application:

```bash
# Build and start both frontend and backend
just up

# Or manually with docker-compose
docker-compose up --build
```

**Access the application:**
- 🌐 **Frontend**: [http://localhost:5173](http://localhost:5173)
- 🔌 **Backend API**: [http://localhost:8080/api](http://localhost:8080/api)

**Stop the services:**
```bash
just down
```

### Option 2: Local Development

#### Prerequisites

- Go 1.21 or higher
- Node.js 20 or higher
- Just command runner (optional, but recommended)

#### Installation

1. Clone the repository:
```bash
git clone https://github.com/sklyarvladislav/tracker.git
cd tracker
```

2. Install dependencies:
```bash
# Using Just
just install

# Or manually
go mod download
cd frontend && npm install
```

### Development

#### Using Just (Recommended)

```bash
# Run both backend and frontend
just dev

# Run backend only
just dev-backend
# Access at: http://localhost:8080

# Run frontend only (in another terminal)
just dev-frontend
# Access at: http://localhost:5173

# Build the application
just build

# Run migrations
just migrate
```

#### Manual Commands

**Backend:**
```bash
# Run backend server
go run .
# Access at: http://localhost:8080

# Build backend
go build -o tracker-backend .
```

**Frontend:**
```bash
# Run frontend dev server
cd frontend
npm run dev
# Access at: http://localhost:5173

# Build frontend
npm run build
```

### Production Deployment

#### Using Docker

```bash
# Build Docker image
docker build -t habit-tracker .

# Run container
docker run -p 8080:8080 -v $(pwd)/data:/data habit-tracker
```

#### Using Docker Compose

```bash
# Start the application
docker-compose up -d

# Stop the application
docker-compose down
```

#### Using Just

```bash
# Build Docker image
just docker-build

# Run Docker container
just docker-run

# Or use docker-compose
just compose-up
```

## API Documentation

### Habits

- `GET /api/habits` - Get all habits
- `POST /api/habits` - Create a new habit
- `GET /api/habits/:id` - Get a specific habit
- `PUT /api/habits/:id` - Update a habit
- `DELETE /api/habits/:id` - Delete a habit

### Habit Entries

- `GET /api/habits/:id/entries` - Get all entries for a habit
- `POST /api/habits/:id/entries` - Create a new entry
- `POST /api/habits/:id/toggle?date=YYYY-MM-DD` - Toggle entry for a specific date

### Request/Response Examples

**Create Habit:**
```json
POST /api/habits
{
  "name": "Morning Exercise",
  "description": "30 minutes of exercise",
  "color": "#10b981"
}
```

**Toggle Entry:**
```bash
POST /api/habits/1/toggle?date=2024-01-15
```

## Environment Variables

### Backend
- `PORT` - Server port (default: 8080)
- `DATABASE_PATH` - SQLite database file path (default: ./tracker.db)
- `GIN_MODE` - Gin mode (debug/release, default: debug)

### Frontend
- `VITE_API_URL` - Backend API URL (default: http://localhost:8080/api)

## Project Structure

```
tracker/
├── backend/
│   ├── database/       # Database initialization and migrations
│   ├── handlers/       # HTTP request handlers
│   ├── middleware/     # Middleware functions
│   └── models/         # Data models
├── frontend/
│   ├── src/
│   │   ├── components/ # Vue components
│   │   ├── services/   # API service layer
│   │   ├── App.vue     # Main application component
│   │   └── main.js     # Application entry point
│   └── public/         # Static assets
├── main.go             # Backend entry point
├── Dockerfile          # Docker configuration
├── docker-compose.yml  # Docker Compose configuration
├── Justfile            # Task automation
└── README.md           # This file
```

## Available Just Commands

```bash
just install        # Install all dependencies
just migrate        # Run database migrations
just build          # Build frontend and backend
just build-backend  # Build backend only
just build-frontend # Build frontend only
just dev-backend    # Run backend in development mode
just dev-frontend   # Run frontend in development mode
just docker-build   # Build Docker image
just docker-run     # Run Docker container
just compose-up     # Start with docker-compose
just compose-down   # Stop docker-compose
just clean          # Clean build artifacts
just test           # Run tests
just fmt            # Format code
just lint           # Lint code
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

MIT License

## Author

Vladislav Sklyar
