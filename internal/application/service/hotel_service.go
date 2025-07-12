package service

import (
	"context"
	"github.com/jamal23041989/go_reservation_hotel/internal/domain"
	"github.com/jamal23041989/go_reservation_hotel/internal/domain/entity"
	"github.com/jamal23041989/go_reservation_hotel/internal/domain/repository"
)

type HotelService struct {
	repo repository.HotelRepository
}

func NewHotelService(repo repository.HotelRepository) *HotelService {
	return &HotelService{
		repo: repo,
	}
}

func (uc *HotelService) GetAllHotels(ctx context.Context, filter domain.Map, pag *domain.Pagination) ([]*entity.Hotel, error) {
	if pag.Limit <= 0 {
		pag.Limit = 10
	}
	if pag.Page <= 0 {
		pag.Page = 1
	}

	hotels, err := uc.repo.GetAllHotels(ctx, filter, pag)
	if err != nil {
		return nil, err
	}
	return hotels, nil
}

func (uc *HotelService) CreateHotel(ctx context.Context, hotel *entity.Hotel) (*entity.Hotel, error) {
	createHotel, err := uc.repo.CreateHotel(ctx, hotel)
	if err != nil {
		return nil, err
	}
	return createHotel, nil
}

func (uc *HotelService) UpdateHotel(ctx context.Context, filter domain.Map, update domain.Map) error {
	if err := uc.repo.UpdateHotel(ctx, filter, update); err != nil {
		return err
	}
	return nil
}

func (uc *HotelService) GetByIDHotel(ctx context.Context, id string) (*entity.Hotel, error) {
	hotel, err := uc.repo.GetByIDHotel(ctx, id)
	if err != nil {
		return nil, err
	}
	return hotel, nil
}
