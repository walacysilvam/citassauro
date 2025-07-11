package models

import (
	"time"
)

type Quote struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Author    string    `json:"author" bind:"required"`
	Text      string    `json:"text" bind:"required"`
	Votes     int       `json:"votes"`
	UserID    uint      `json:"user_id"` // autor da citação
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}