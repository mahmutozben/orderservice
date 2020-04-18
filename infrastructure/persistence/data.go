package persistence

import (
	database "order-service/database"
	domain "order-service/domain"
	"order-service/domain/repository"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Repositories struct {
	User repository.UserRepository
	db   *gorm.DB
}

func NewRepositories() (*Repositories, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}
	return &Repositories{
		User: NewUserRepository(db),
		db:   db,
	}, nil
}

//closes the  database connection
func (s *Repositories) Close() error {
	return s.db.Close()
}

//This migrate all tables
func (s *Repositories) Automigrate() error {
	return s.db.AutoMigrate(&domain.User{}).Error
}
