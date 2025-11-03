package main

import (
	"fmt"
	"log"
	"os"

	"sticker-store-backend/internal/database"
	"sticker-store-backend/internal/handlers"
	"sticker-store-backend/internal/models"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.Connect()

	// Auto-migrate your models
	database.DB.AutoMigrate(&models.Sticker{})

	app := fiber.New()

	// Simple route test
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("🚀 Sticker Store API is running!")
	})

	app.Get("/api/stickers", handlers.GetStickers)
	app.Post("/api/stickers", handlers.CreateSticker)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("🚀 Server running on port", port)
	log.Fatal(app.Listen(":" + port))
}
