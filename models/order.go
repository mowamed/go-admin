package models

import "gorm.io/gorm"

type Order struct {
	Id         uint        `json:"id"`
	FirstName  string      `json:"-"`
	LastName   string      `json:"-"`
	Name       string      `json:"name" gorm:"-"`
	Total      float32     `json:"total" gorm:"-"`
	Email      string      `json:"email"`
	UpdatedAt  string      `json:"updated_at"`
	CreatedAt  string      `json:"created_at"`
	OrderItems []OrderItem `json:"order_items" gorm:"foreignKey:OrderId"`
}

type OrderItem struct {
	Id           uint
	OrderId      uint
	ProductTitle string
	Price        float32
	Quantity     int
}

func (order *Order) Count(db *gorm.DB) int64 {
	var total int64
	db.Model(&Order{}).Count(&total)
	return total
}

func (order *Order) Take(db *gorm.DB, limit int, offset int) interface{} {
	var orders []Order
	db.Preload("OrderItems").Offset(offset).Limit(limit).Find(&orders)

	for i, _ := range orders {
		var total float32 = 0

		for _, orderItem := range orders[i].OrderItems {
			total += orderItem.Price * float32(orderItem.Quantity)
		}

		orders[i].Name = orders[i].FirstName + " " + orders[i].LastName
		orders[i].Total = total
	}
	return orders
}
