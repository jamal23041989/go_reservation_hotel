package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jamal23041989/go_reservation_hotel/internal/domain/models"
	"github.com/jamal23041989/go_reservation_hotel/pkg/auth"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Get("Authorization")
		if token == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "missing token"})
		}

		claims, err := auth.ParseToken(token)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid token"})
		}

		userIDStr, ok := claims["user_id"].(string)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid token payload"})
		}

		userID, err := primitive.ObjectIDFromHex(userIDStr)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid user id"})
		}

		// Можно добавить больше полей при необходимости
		user := models.User{
			ID:    userID,
			Email: claims["email"].(string),
		}

		c.Locals("user", user)
		return c.Next()
	}
}

func GetAuthUser(c *fiber.Ctx) (*models.User, error) {
	user, ok := c.Locals("user").(models.User)
	if !ok {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "unauthorized")
	}
	return &user, nil
}
