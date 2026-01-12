package models

import (
	"time"

	"gorm.io/gorm"
)

// Habit represents a habit that the user wants to track
type Habit struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	Name        string         `gorm:"not null" json:"name"`
	Description string         `json:"description"`
	Color       string         `gorm:"default:#10b981" json:"color"` // Default green color
	Entries     []HabitEntry   `json:"entries,omitempty"`
}

// HabitEntry represents a single day's entry for a habit
type HabitEntry struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	HabitID   uint           `gorm:"not null;index" json:"habit_id"`
	Date      time.Time      `gorm:"not null;index" json:"date"`
	Completed bool           `gorm:"default:false" json:"completed"`
	Notes     string         `json:"notes"`
}
