package user

import (
	"errors"

	"github.com/afrizal423/go-gin-secure-newbie/app/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewuserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db,
	}
}

func (u *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	// if err := u.db.Select("email", "password").Where("email = ?", email).Take(&user).Error; err == nil {
	// 	return nil, errors.New("tidak ada ada")
	// }
	if err := u.db.Select("id", "email", "password", "role").First(&user, "email = ?", email).Error; err != nil {
		return nil, errors.New("tidak ada data")
	}
	return &user, nil
}

func (u *UserRepository) Register(data models.User) (models.User, error) {
	var user models.User
	if err := u.db.Where("email = ?", data.Email).First(&user).Error; err == nil {
		return user, errors.New("email sudah ada")
	}
	if err := u.db.Create(&data).Error; err != nil {
		return user, err
	}
	return data, nil
}
