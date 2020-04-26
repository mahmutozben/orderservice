package repository

import "order-service/domain"

type OrderRepository interface {
	AddOrder(*domain.Order) error
	UpdateOrder() (*domain.Order, error)
	GetOrder(orderId int) (*domain.Order, error)
}
