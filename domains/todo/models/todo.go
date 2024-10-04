package models

import "time"

type Todo struct {
	ID          int       `gorm:"primaryKey;autoIncrement;column:id"`
	Title       string    `gorm:"type:varchar(255);not null;column:title"`
	Description string    `gorm:"type:text;column:description"`
	IsCompleted bool      `gorm:"default:false;column:is_completed"`
	CreatedAt   time.Time `gorm:"type:timestamp with time zone;default:current_timestamp;column:created_at"`
	UpdatedAt   time.Time `gorm:"type:timestamp with time zone;default:current_timestamp;column:updated_at"`
}
