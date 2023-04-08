package models

import "time"

type Product struct {
	ID          uint   `gorm:"primaryKey;" json:"id"`
	Title       string `json:"title"`
	Description string `json:"Description"`
	UserID      uint
	User        *User
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}
