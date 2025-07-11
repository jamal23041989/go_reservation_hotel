package usecase

import (
	"context"
	"github.com/jamal23041989/go_reservation_hotel/internal/domain"
	"github.com/jamal23041989/go_reservation_hotel/internal/domain/models"
	"github.com/jamal23041989/go_reservation_hotel/internal/domain/repositories"
)

type BookingUsecase struct {
	repo repositories.BookingRepository
}

func NewBookingUsecase(repo repositories.BookingRepository) *BookingUsecase {
	return &BookingUsecase{
		repo: repo,
	}
}

func (uc *BookingUsecase) CreateBooking(ctx context.Context, booking *models.Booking) (*models.Booking, error) {
	createBooking, err := uc.repo.CreateBooking(ctx, booking)
	if err != nil {
		return nil, err
	}
	return createBooking, nil
}

func (uc *BookingUsecase) GetBookings(ctx context.Context, filter domain.Map) ([]*models.Booking, error) {
	bookings, err := uc.repo.GetBookings(ctx, filter)
	if err != nil {
		return nil, err
	}
	return bookings, nil
}

func (uc *BookingUsecase) GetBookingByID(ctx context.Context, id string) (*models.Booking, error) {
	booking, err := uc.repo.GetBookingByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return booking, nil
}

func (uc *BookingUsecase) UpdateBooking(ctx context.Context, id string, update domain.Map) error {
	err := uc.repo.UpdateBooking(ctx, id, update)
	if err != nil {
		return err
	}
	return nil
}
