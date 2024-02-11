package models

import "time"

type Item struct {
	ItemId      uint   `gorm:"primaryKey"`
	ItemCode    string `gorm:"not null;type:varchar(255)"`
	Description string `gorm:"not null;type:varchar(255)"`
	Quantity    uint   `gorm:"not null;type:uint"`
	OrderId     uint   `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
