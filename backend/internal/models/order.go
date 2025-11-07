package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID     uint    `json:"user_id"`
	AddressID  uint    `json:"address_id"`
	TotalPrice float64 `json:"total_price"`
	Status     string  `json:"status" gorm:"default:'pending'"`
}

type OrderItem struct {
	gorm.Model
	OrderID   uint    `json:"order_id"`
	StickerID uint    `json:"sticker_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}
