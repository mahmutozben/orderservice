package domain

import "time"

type Order struct {
	Id          int         `gorm:"primary_key;auto_increment" json:"id"`
	Address     string      `gorm:"size:1000;not null;" json:"address"`
	OrderDate   time.Time   `gorm:"default:CURRENT_TIMESTAMP" json:"orderDate"`
	Description string      `gorm:"size:500;not null;" json:"description"`
	Status      OrderStatus `gorm:"size:500;not null;" json:"status"`
	OrderItems  []OrderItem
}

type OrderStatus int

const (
	Initial   OrderStatus = 1
	Paid      OrderStatus = 2
	Shipped   OrderStatus = 3
	Cancelled OrderStatus = 4
)

func NewOrder(address, description string) Order {
	order := Order{}
	order.Address = address
	order.Description = description
	order.Status = Initial
	order.OrderDate = time.Now()
	order.OrderItems = []OrderItem{}

	return order
}

func (order *Order) AddOrderItem(productId int, productName string, unitPrice float64, discount int) {
	orderItem := OrderItem{ProductId: productId, ProductName: productName, UnitPrice: unitPrice, Discount: discount}
	order.OrderItems = append(order.OrderItems, orderItem)
}
