package models

import "gorm.io/gorm"

type CartItem struct {
	gorm.Model
	UserID    uint `json:"user_id"`
	StickerID uint `json:"sticker_id"`
	Quantity  int  `json:"quntity"`
}
