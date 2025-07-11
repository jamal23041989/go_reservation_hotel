package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/jamal23041989/go_reservation_hotel/internal/domain/models"
)

func getAuthUser(c *fiber.Ctx) (*models.User, error) {
	user, ok := c.Context().UserValue("user").(*models.User)
	if !ok {
		return nil, fmt.Errorf("unauthorized")
	}
	return user, nil
}
