package persistence

import (
	"errors"
	domain "order-service/domain"
	"order-service/domain/repository"
	"strings"

	"github.com/jinzhu/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepo {
	return &UserRepo{db}
}

//UserRepo implements the repository.UserRepository interface
var _ repository.UserRepository = &UserRepo{}

func (r *UserRepo) SaveUser(user *domain.User) (*domain.User, error) {
	err := r.db.Debug().Create(&user).Error
	if err != nil {
		//If the email is already taken
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "Duplicate") {
			return nil, errors.New("email already taken")
		}
		return nil, errors.New("database error")
	}
	return user, nil
}

func (r *UserRepo) GetUser(id uint64) (*domain.User, error) {
	var user domain.User
	err := r.db.Debug().Where("id = ?", id).Take(&user).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("user not found")
	}
	return &user, nil
}

func (r *UserRepo) GetUsers() ([]domain.User, error) {
	var users []domain.User
	err := r.db.Debug().Find(&users).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("user not found")
	}
	return users, nil
}
