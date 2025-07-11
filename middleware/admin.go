package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jamal23041989/go_reservation_hotel/pkg"
	"github.com/jamal23041989/go_reservation_hotel/types"
)

func AdminAuth(c *fiber.Ctx) error {
	user, ok := c.Context().UserValue("user").(*types.User)
	if !ok {
		return pkg.ErrUnauthorized()
	}
	if !user.IsAdmin {
		return pkg.ErrUnauthorized()
	}
	return c.Next()
}
