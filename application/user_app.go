package application

import (
	domain "order-service/domain"
	"order-service/domain/repository"
)

type userApp struct {
	us repository.UserRepository
}

//UserApp implements the UserAppInterface
var _ UserAppInterface = &userApp{}

type UserAppInterface interface {
	SaveUser(*domain.User) (*domain.User, error)
	GetUsers() ([]domain.User, error)
	GetUser(uint64) (*domain.User, error)
}

func (u *userApp) SaveUser(user *domain.User) (*domain.User, error) {
	return u.us.SaveUser(user)
}

func (u *userApp) GetUser(userId uint64) (*domain.User, error) {
	return u.us.GetUser(userId)
}

func (u *userApp) GetUsers() ([]domain.User, error) {
	return u.us.GetUsers()
}
