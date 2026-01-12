package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sklyarvladislav/tracker/backend/database"
	"github.com/sklyarvladislav/tracker/backend/models"
)

// GetHabits returns all habits
func GetHabits(c *gin.Context) {
	var habits []models.Habit
	if err := database.DB.Preload("Entries").Find(&habits).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, habits)
}

// GetHabit returns a single habit by ID
func GetHabit(c *gin.Context) {
	id := c.Param("id")
	var habit models.Habit
	if err := database.DB.Preload("Entries").First(&habit, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Habit not found"})
		return
	}
	c.JSON(http.StatusOK, habit)
}

// CreateHabit creates a new habit
func CreateHabit(c *gin.Context) {
	var habit models.Habit
	if err := c.ShouldBindJSON(&habit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if habit.Color == "" {
		habit.Color = "#10b981"
	}

	if err := database.DB.Create(&habit).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, habit)
}

// UpdateHabit updates an existing habit
func UpdateHabit(c *gin.Context) {
	id := c.Param("id")
	var habit models.Habit
	if err := database.DB.First(&habit, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Habit not found"})
		return
	}

	var input models.Habit
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update fields
	habit.Name = input.Name
	habit.Description = input.Description
	if input.Color != "" {
		habit.Color = input.Color
	}

	if err := database.DB.Save(&habit).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, habit)
}

// DeleteHabit deletes a habit
func DeleteHabit(c *gin.Context) {
	id := c.Param("id")
	if err := database.DB.Delete(&models.Habit{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Habit deleted successfully"})
}

// GetHabitEntries returns all entries for a habit
func GetHabitEntries(c *gin.Context) {
	habitID := c.Param("id")
	var entries []models.HabitEntry
	if err := database.DB.Where("habit_id = ?", habitID).Find(&entries).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, entries)
}

// ToggleHabitEntry toggles a habit entry for a specific date
func ToggleHabitEntry(c *gin.Context) {
	habitID := c.Param("id")
	dateStr := c.Query("date")
	
	if dateStr == "" {
		dateStr = time.Now().Format("2006-01-02")
	}

	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format. Use YYYY-MM-DD"})
		return
	}

	// Parse habitID to uint
	var habitIDUint uint
	database.DB.Raw("SELECT ? + 0", habitID).Scan(&habitIDUint)

	// Check if entry exists for this date
	var entry models.HabitEntry
	result := database.DB.Where("habit_id = ? AND DATE(date) = DATE(?)", habitIDUint, date).First(&entry)

	if result.Error != nil {
		// Create new entry
		entry = models.HabitEntry{
			HabitID:   habitIDUint,
			Date:      date,
			Completed: true,
		}

		if err := database.DB.Create(&entry).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	} else {
		// Toggle existing entry
		entry.Completed = !entry.Completed
		if err := database.DB.Save(&entry).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, entry)
}

// CreateHabitEntry creates or updates a habit entry
func CreateHabitEntry(c *gin.Context) {
	habitID := c.Param("id")
	
	var input struct {
		Date      string `json:"date"`
		Completed bool   `json:"completed"`
		Notes     string `json:"notes"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	date, err := time.Parse("2006-01-02", input.Date)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format. Use YYYY-MM-DD"})
		return
	}

	// Parse habitID to uint
	var habitIDUint uint
	database.DB.Raw("SELECT ? + 0", habitID).Scan(&habitIDUint)

	// Check if entry exists
	var entry models.HabitEntry
	result := database.DB.Where("habit_id = ? AND DATE(date) = DATE(?)", habitIDUint, date).First(&entry)

	if result.Error != nil {
		// Create new entry
		entry = models.HabitEntry{
			HabitID:   habitIDUint,
			Date:      date,
			Completed: input.Completed,
			Notes:     input.Notes,
		}
		if err := database.DB.Create(&entry).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, entry)
	} else {
		// Update existing entry
		entry.Completed = input.Completed
		entry.Notes = input.Notes
		if err := database.DB.Save(&entry).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, entry)
	}
}
