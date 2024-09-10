package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	OrderID      uint      `gorm:"primaryKey" json:"order_id"`
	CustomerName string    `json:"customer_name"`
	OrderedAt    time.Time `json:"ordered_at"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Items        []Item    `gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE;" json:"items"`
}

type OrderRepository interface {
	CreateOrder(order *Order) error
	GetOrder() ([]Order, error)
	UpdateOrder(*Order) error
	DeleteOrder(id int) error
}

type connectOrder struct {
	DB *gorm.DB
}

func new_connectOrder(db *gorm.DB) OrderRepository {
	return &connectOrder{db}
}

func (o *connectOrder) CreateOrder(order *Order) error {
	return o.DB.Create(order).Error
}
func (o *connectOrder) GetOrder() ([]Order, error) {
	var orders []Order
	err := o.DB.Find(&orders).Error
	return orders, err
}

func (o *connectOrder) UpdateOrder(order *Order) error {
	return o.DB.Save(order).Error
}

func (o *connectOrder) DeleteOrder(id int) error {
	return o.DB.Delete(&Order{}, id).Error
}
