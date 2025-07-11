package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jamal23041989/go_reservation_hotel/pkg"
)

func AdminAuth(c *fiber.Ctx) error {
	user, ok := c.Context().UserValue("user").(*typess.User)
	if !ok {
		return pkg.ErrUnauthorized()
	}
	if !user.IsAdmin {
		return pkg.ErrUnauthorized()
	}
	return c.Next()
}
