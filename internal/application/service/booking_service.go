package service

import (
	"context"
	"github.com/jamal23041989/go_reservation_hotel/internal/domain"
	"github.com/jamal23041989/go_reservation_hotel/internal/domain/entity"
	"github.com/jamal23041989/go_reservation_hotel/internal/domain/repository"
)

type BookingService struct {
	repo repository.BookingRepository
}

func NewBookingService(repo repository.BookingRepository) *BookingService {
	return &BookingService{
		repo: repo,
	}
}

func (uc *BookingService) CreateBooking(ctx context.Context, booking *entity.Booking) (*entity.Booking, error) {
	createBooking, err := uc.repo.CreateBooking(ctx, booking)
	if err != nil {
		return nil, err
	}
	return createBooking, nil
}

func (uc *BookingService) GetBookings(ctx context.Context, filter domain.Map) ([]*entity.Booking, error) {
	bookings, err := uc.repo.GetBookings(ctx, filter)
	if err != nil {
		return nil, err
	}
	return bookings, nil
}

func (uc *BookingService) GetBookingByID(ctx context.Context, id string) (*entity.Booking, error) {
	booking, err := uc.repo.GetBookingByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return booking, nil
}

func (uc *BookingService) UpdateBooking(ctx context.Context, id string, update domain.Map) error {
	err := uc.repo.UpdateBooking(ctx, id, update)
	if err != nil {
		return err
	}
	return nil
}
