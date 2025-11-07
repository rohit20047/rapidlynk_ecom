package handlers

import (
	"sticker-store-backend/internal/database"
	"sticker-store-backend/internal/models"

	"github.com/gofiber/fiber/v2"
)

func CreateOrder(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	var body struct {
		AddressID uint `json:"address_id"`
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "faile to aprse the address "})
	}

	var cartItems []models.CartItem

	if err := database.DB.Where("user_id = ?", userID).Find(&cartItems).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to fetch cart"})
	}
	if len(cartItems) == 0 {
		return c.Status(400).JSON(fiber.Map{"error": "cart is empty"})
	}

	var total float64

	for _, item := range cartItems {
		var sticker models.Sticker
		database.DB.First(&sticker, item.StickerID)
		total += float64(item.Quantity) * sticker.Price
	}

	order := models.Order{
		UserID:     userID,
		AddressID:  body.AddressID,
		TotalPrice: total,
		Status:     "pending",
	}

	database.DB.Create(&order)

	for _, item := range cartItems {
		var sticker models.Sticker
		database.DB.First(&sticker, item.StickerID)

		orderItem := models.OrderItem{
			OrderID:   order.ID,
			StickerID: item.StickerID,
			Quantity:  item.Quantity,
			Price:     sticker.Price,
		}
		database.DB.Create(&orderItem)
	}

	// Empty the cart after checkout
	database.DB.Where("user_id = ?", userID).Delete(&models.CartItem{})

	return c.JSON(order)
}
