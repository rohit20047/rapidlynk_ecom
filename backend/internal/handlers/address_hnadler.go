package handlers

import (
	"sticker-store-backend/internal/database"
	"sticker-store-backend/internal/models"

	"github.com/gofiber/fiber/v2"
)

func AddAddress(c *fiber.Ctx) error {
	userId := c.Locals("user_id").(uint)

	var body struct {
		User_Id uint   `json:"user_id"`
		HouseNo string `json:"house_no"`
		Street  string `json:"street"`
		Line1   string `json:"line1"`
		City    string `json:"city"`
		State   string `json:"state"`
		ZipCode string `json:"zip_code"`
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request body"})
	}

	address := models.Address{
		User_Id: userId,
		HouseNo: body.HouseNo,
		Street:  body.Street,
		Line1:   body.Line1,
		City:    body.City,
		State:   body.State,
		ZipCode: body.ZipCode,
	}

	if err := database.DB.Create(&address).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "could not save address"})
	}

	return c.Status(201).JSON(address)

}

func GetAddresses(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	var addresses []models.Address
	if err := database.DB.Where("user_id = ?", userID).Find(&addresses).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "could not fetch addresses"})
	}

	return c.JSON(addresses)

}
