package services

type IUserService interface {
	GetUser() string
	GetEmail() string
}

type UserService struct {
}

func (u *UserService) GetUser() string {
	return "User"
}

func (u *UserService) GetEmail() string {
	return "Email"
}

func NewUserService() IUserService {
	return &UserService{}
}

var _ IUserService = (*UserService)(nil)
