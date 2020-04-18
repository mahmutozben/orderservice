package interfaces

import (
	r "order-service/infrastructure/persistence"
)

type HandlerRepository struct {
	UserHandler UserHandler
}

func CreateHandlers(services *r.Repositories) *HandlerRepository {
	return &HandlerRepository{
		UserHandler: NewUserHandler(services.User),
	}
}
