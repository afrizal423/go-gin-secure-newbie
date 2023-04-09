package models

import "time"

type User struct {
	ID       uint   `gorm:"primaryKey;" json:"id"`
	Email    string `gorm:"not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
	// role 1 admin role 2 user
	Role      uint       `gorm:"not null" json:"roles"`
	Products  []Product  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"products"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}
