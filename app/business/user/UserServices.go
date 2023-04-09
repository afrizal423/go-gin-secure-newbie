package user

import (
	"errors"

	"github.com/afrizal423/go-gin-secure-newbie/app/models"
	"github.com/afrizal423/go-gin-secure-newbie/pkg/utils/hash"
)

type UserService struct {
	repository IUserRepository
	hashing    hash.Hashing
}

func NewUserService(repository IUserRepository) *UserService {
	return &UserService{
		repository,
		hash.Hashing{},
	}
}

func (u *UserService) Register(data models.User) (models.User, error) {
	var dt models.User
	// hash password
	data.Password = u.hashing.HashPassword(data.Password)
	// repositoty to insert data
	data, err := u.repository.Register(data)
	if err != nil {
		return dt, err
	}
	return data, nil
}

func (u *UserService) Login(data models.User) (string, error) {
	dt, err := u.repository.FindByEmail(data.Email)
	if err != nil {
		return "", errors.New("tidak ada data")
	}
	// fmt.Println(data.Password, dt.Password)
	// return "", nil
	match, err := u.hashing.VerifikasiPassword(data.Password, dt.Password)
	if err != nil {
		return "", err
	}

	if match {
		// bener
		return "ini jwt token", nil
	} else {
		// salah
		return "", errors.New("password salah")
	}
}
