package models

import "time"

type User struct {
	ID        uint       `gorm:"primaryKey;" json:"id"`
	Email     string     `gorm:"not null" json:"email"`
	Password  string     `gorm:"not null" json:"password"`
	Products  []Product  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"products"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}
