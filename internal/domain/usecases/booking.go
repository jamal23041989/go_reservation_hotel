package usecases

import (
	"context"
	"github.com/jamal23041989/go_reservation_hotel/internal/domain"
	"github.com/jamal23041989/go_reservation_hotel/internal/domain/models"
)

type BookingUsecase interface {
	GetBookings(context.Context, domain.Map) ([]*models.Booking, error)
	GetBookingByID(context.Context, string) (*models.Booking, error)
	CreateBooking(context.Context, *models.Booking) (*models.Booking, error)
	UpdateBooking(context.Context, string, domain.Map) error
}
