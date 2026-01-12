package main

import (
"log"
"os"

"github.com/gin-gonic/gin"
"github.com/sklyarvladislav/tracker/backend/database"
"github.com/sklyarvladislav/tracker/backend/handlers"
"github.com/sklyarvladislav/tracker/backend/middleware"
)

func main() {
// Get database path from environment or use default
dbPath := os.Getenv("DATABASE_PATH")
if dbPath == "" {
dbPath = "./tracker.db"
}

// Initialize database
if err := database.Initialize(dbPath); err != nil {
log.Fatal("Failed to connect to database:", err)
}

// Run migrations
if err := database.Migrate(); err != nil {
log.Fatal("Failed to run migrations:", err)
}

// Set up Gin router
r := gin.Default()
r.Use(middleware.CORS())

// API routes
api := r.Group("/api")
{
// Habit routes
api.GET("/habits", handlers.GetHabits)
api.POST("/habits", handlers.CreateHabit)
api.GET("/habits/:id", handlers.GetHabit)
api.PUT("/habits/:id", handlers.UpdateHabit)
api.DELETE("/habits/:id", handlers.DeleteHabit)

// Habit entry routes
api.GET("/habits/:id/entries", handlers.GetHabitEntries)
api.POST("/habits/:id/entries", handlers.CreateHabitEntry)
api.POST("/habits/:id/toggle", handlers.ToggleHabitEntry)
}

// Serve static files in production
r.Static("/assets", "./frontend/dist/assets")
r.StaticFile("/", "./frontend/dist/index.html")
r.NoRoute(func(c *gin.Context) {
c.File("./frontend/dist/index.html")
})

// Get port from environment or use default
port := os.Getenv("PORT")
if port == "" {
port = "8080"
}

log.Printf("Server starting on port %s", port)
if err := r.Run(":" + port); err != nil {
log.Fatal("Failed to start server:", err)
}
}
