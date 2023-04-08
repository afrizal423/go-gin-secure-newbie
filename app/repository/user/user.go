package user

import "gorm.io/gorm"

type UserRepository struct {
	db *gorm.DB
}

func NewuserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db,
	}
}
