package repository

import (
	domain "order-service/domain"
)

type UserRepository interface {
	SaveUser(*domain.User) (*domain.User, error)
	GetUser(uint64) (*domain.User, error)
	GetUsers() ([]domain.User, error)
}
