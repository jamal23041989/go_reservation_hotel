package usecase

import (
	"context"
	"github.com/jamal23041989/go_reservation_hotel/internal/domain"
	"github.com/jamal23041989/go_reservation_hotel/internal/domain/models"
	"github.com/jamal23041989/go_reservation_hotel/internal/domain/repositories"
)

type HotelUsecase struct {
	repo repositories.HotelRepository
}

func NewHotelUsecase(repo repositories.HotelRepository) *HotelUsecase {
	return &HotelUsecase{
		repo: repo,
	}
}

func (uc *HotelUsecase) GetAllHotels(ctx context.Context, filter domain.Map, pag *domain.Pagination) ([]*models.Hotel, error) {
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

func (uc *HotelUsecase) CreateHotel(ctx context.Context, hotel *models.Hotel) (*models.Hotel, error) {
	createHotel, err := uc.repo.CreateHotel(ctx, hotel)
	if err != nil {
		return nil, err
	}
	return createHotel, nil
}

func (uc *HotelUsecase) UpdateHotel(ctx context.Context, filter domain.Map, update domain.Map) error {
	if err := uc.repo.UpdateHotel(ctx, filter, update); err != nil {
		return err
	}
	return nil
}

func (uc *HotelUsecase) GetByIDHotel(ctx context.Context, id string) (*models.Hotel, error) {
	hotel, err := uc.repo.GetByIDHotel(ctx, id)
	if err != nil {
		return nil, err
	}
	return hotel, nil
}
