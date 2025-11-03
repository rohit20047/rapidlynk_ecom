package handlers

import (
	"log"
	"os"
	"time"

	"sticker-store-backend/internal/database"
	"sticker-store-backend/internal/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// POST /api/register
func Register(c *fiber.Ctx) error {

	jwtSecret := []byte(os.Getenv("JWT_SECRET"))
	log.Println(jwtSecret)
	if len(jwtSecret) == 0 {
		log.Fatal("JWT_SECRET not set in environment")
	}
	var body struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request"})
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "could not hash password"})
	}

	// Create user
	user := models.User{
		Name:     body.Name,
		Email:    body.Email,
		Password: string(hashedPassword),
	}

	result := database.DB.Create(&user)
	if result.Error != nil {
		return c.Status(400).JSON(fiber.Map{"error": "email already exists"})
	}

	// Generate JWT token (auto-login)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(72 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "could not create token"})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "user registered successfully",
		"token":   tokenString,
		"user": fiber.Map{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
		},
	})
}
