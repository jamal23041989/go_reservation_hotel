package repositories

import (
	"context"
	"github.com/jamal23041989/go_reservation_hotel/internal/domain"
	"github.com/jamal23041989/go_reservation_hotel/internal/domain/models"
)

type BookingRepository interface {
	CreateBooking(context.Context, *models.Booking) (*models.Booking, error)
	GetBookings(context.Context, domain.Map) ([]*models.Booking, error)
	GetBookingByID(context.Context, string) (*models.Booking, error)
	UpdateBooking(context.Context, string, domain.Map) error
}
