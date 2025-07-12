package handler

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/jamal23041989/go_reservation_hotel/internal/domain"
	"github.com/jamal23041989/go_reservation_hotel/internal/domain/usecases"
	"github.com/jamal23041989/go_reservation_hotel/internal/middleware"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

type BookingHandler struct {
	bookingUsecase usecases.BookingUsecase
}

func NewBookingHandler(bookingUsecase usecases.BookingUsecase) *BookingHandler {
	return &BookingHandler{
		bookingUsecase: bookingUsecase,
	}
}

func (h *BookingHandler) HandleGetBookings(c *fiber.Ctx) error {
	bookings, err := h.bookingUsecase.GetBookings(c.Context(), domain.Map{})
	if err != nil {
		return err
	}
	return c.JSON(bookings)
}

func (h *BookingHandler) HandleGetBooking(c *fiber.Ctx) error {
	id := c.Params("id")
	booking, err := h.bookingUsecase.GetBookingByID(c.Context(), id)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return c.JSON(map[string]string{"error": "not found"})
		}
		return err
	}

	user, err := middleware.GetAuthUser(c)
	if err != nil {
		return err
	}

	if booking.UserID != user.ID {
		return c.Status(http.StatusUnauthorized).JSON(domain.GenericResp{
			Type: "error",
			Msg:  "not authorized",
		})
	}

	return c.JSON(booking)
}

func (h *BookingHandler) HandleCancelBooking(c *fiber.Ctx) error {
	id := c.Params("id")

	booking, err := h.bookingUsecase.GetBookingByID(c.Context(), id)
	if err != nil {
		return err
	}

	user, err := middleware.GetAuthUser(c)
	if err != nil {
		return err
	}

	if booking.UserID != user.ID {
		return c.Status(http.StatusUnauthorized).JSON(domain.GenericResp{
			Type: "error",
			Msg:  "not authorized",
		})
	}

	if err := h.bookingUsecase.UpdateBooking(c.Context(), c.Params("id"), domain.Map{"canceled": true}); err != nil {
		return err
	}

	return c.JSON(domain.GenericResp{
		Type: "msg",
		Msg:  "updated",
	})
}
