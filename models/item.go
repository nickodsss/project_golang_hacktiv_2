package models

import (
	"time"

	"gorm.io/gorm"
)

type Item struct {
	ItemID      uint      `gorm:"primaryKey" json:"item_id"`
	ItemCode    string    `json:"item_code"`
	Description string    `json:"description"`
	Quantity    int       `json:"quantity"`
	OrderID     uint      `json:"order_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ItemRepository interface {
	CreateItem(item *Item) error
	GetItem() ([]Item, error)
	UpdateItem(*Item) error
	DeleteItem(id int) error
}

type connectItem struct {
	DB *gorm.DB
}

func new_connectItem(db *gorm.DB) ItemRepository {
	return &connectItem{db}
}

func (o *connectItem) CreateItem(item *Item) error {
	return o.DB.Create(item).Error
}

func (o *connectItem) GetItem() ([]Item, error) {
	var items []Item
	err := o.DB.Find(&items).Error
	return items, err
}

func (o *connectItem) UpdateItem(item *Item) error {
	return o.DB.Save(item).Error
}

func (o *connectItem) DeleteItem(id int) error {
	return o.DB.Delete(&Item{}, id).Error
}
