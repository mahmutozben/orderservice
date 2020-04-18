package database

import (
	"fmt"
	configuration "order-service/configuration"

	"github.com/jinzhu/gorm"
)

func Connect() (*gorm.DB, error) {

	config, err := configuration.LoadConfiguration()
	if err != nil {
		panic(err)
	}

	conn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		config.Database.Host,
		config.Database.Port,
		config.Database.User,
		config.Database.Database,
		config.Database.Password)

	db, err := gorm.Open(config.Database.DbDriver, conn)
	if err != nil {
		return nil, err
	}
	db.LogMode(true)

	return db, nil
}
