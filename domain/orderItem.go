package domain

type OrderItem struct {
	ProductId   int     `gorm:"size:100;not null;" json:"productId"`
	ProductName string  `gorm:"size:100;not null;" json:"productName"`
	UnitPrice   float64 `gorm:"size:100;not null;" json:"unitPrice"`
	Discount    int     `gorm:"size:100;not null;" json:"discount"`
}
