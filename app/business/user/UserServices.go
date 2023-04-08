package user

type UserService struct {
	repository IUserRepository
}

func NewUserService(repository IUserRepository) *UserService {
	return &UserService{
		repository,
	}
}
