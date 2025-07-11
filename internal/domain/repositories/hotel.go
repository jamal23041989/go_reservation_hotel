package repositories

import (
	"context"
	"github.com/jamal23041989/go_reservation_hotel/internal/domain"
	"github.com/jamal23041989/go_reservation_hotel/internal/domain/models"
)

type HotelRepository interface {
	GetAllHotels(context.Context, domain.Map, *domain.Pagination) ([]*models.Hotel, error)
	CreateHotel(context.Context, *models.Hotel) (*models.Hotel, error)
	UpdateHotel(context.Context, domain.Map, domain.Map) error
	GetByIDHotel(context.Context, string) (*models.Hotel, error)
}
