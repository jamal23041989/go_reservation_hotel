package service

import (
	"context"
	"github.com/jamal23041989/go_reservation_hotel/internal/domain"
	"github.com/jamal23041989/go_reservation_hotel/internal/domain/entity"
)

type BookingService interface {
	GetBookings(context.Context, domain.Map) ([]*entity.Booking, error)
	GetBookingByID(context.Context, string) (*entity.Booking, error)
	CreateBooking(context.Context, *entity.Booking) (*entity.Booking, error)
	UpdateBooking(context.Context, string, domain.Map) error
}
