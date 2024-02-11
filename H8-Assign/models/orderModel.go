package models

import "time"

type Order struct {
	OrderId      uint   `gorm:"primaryKey"`
	CustomerName string `gorm:"not null;type:varchar(255)"`
	OrderList    []Item
	OrderedAt    time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
