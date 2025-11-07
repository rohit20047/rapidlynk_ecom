package handlers

import (
	"sticker-store-backend/internal/database"
	"sticker-store-backend/internal/models"

	"github.com/gofiber/fiber/v2"
)

func ADDToCart(c *fiber.Ctx) error {
	userId := c.Locals("user_id").(uint)

	var body struct {
		StickerID uint `json:"sticker_id"`
		Quantity  int  `json:"quantity"`
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request"})
	}

	if body.Quantity <= 0 {
		body.Quantity = 1
	}

	item := models.CartItem{
		UserID:    userId,
		StickerID: body.StickerID,
		Quantity:  body.Quantity,
	}

	if err := database.DB.Create(&item).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "could not add ton cart"})
	}

	return c.Status(201).JSON(item)

}

func GetCart(c *fiber.Ctx) error {
	userId := c.Locals("user_id").(uint)

	var items []models.CartItem

	if err := database.DB.Where("user_id = ?", userId).Find(&items).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "could not fetch cart "})
	}

	return c.JSON(items)
}

func RemoveFromCart(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := database.DB.Delete(&models.CartItem{}, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "could not delete item"})
	}

	return c.JSON(fiber.Map{"message": "item removed"})
}
