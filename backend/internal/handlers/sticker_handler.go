package handlers

import (
	"sticker-store-backend/internal/database"
	"sticker-store-backend/internal/models"

	"github.com/gofiber/fiber/v2"
)

// GetStickers - fetch all stickers
func GetStickers(c *fiber.Ctx) error {
	var stickers []models.Sticker
	result := database.DB.Find(&stickers)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": result.Error.Error(),
		})
	}

	return c.JSON(stickers)
}

// CreateSticker - add new sticker
func CreateSticker(c *fiber.Ctx) error {
	var sticker models.Sticker

	if err := c.BodyParser(&sticker); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	result := database.DB.Create(&sticker)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": result.Error.Error(),
		})
	}

	return c.Status(201).JSON(sticker)
}
