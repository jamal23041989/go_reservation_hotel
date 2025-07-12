package service

import (
	"context"
	"github.com/jamal23041989/go_reservation_hotel/internal/domain"
	"github.com/jamal23041989/go_reservation_hotel/internal/domain/entity"
)

type HotelService interface {
	GetAllHotels(context.Context, domain.Map, *domain.Pagination) ([]*entity.Hotel, error)
	CreateHotel(context.Context, *entity.Hotel) (*entity.Hotel, error)
	UpdateHotel(context.Context, domain.Map, domain.Map) error
	GetByIDHotel(context.Context, string) (*entity.Hotel, error)
}
