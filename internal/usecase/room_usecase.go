package usecase

import (
	"context"
	"github.com/jamal23041989/go_reservation_hotel/internal/domain"
	"github.com/jamal23041989/go_reservation_hotel/internal/domain/models"
	"github.com/jamal23041989/go_reservation_hotel/internal/domain/repositories"
)

type RoomUsecase struct {
	repo repositories.RoomRepository
}

func NewRoomUsecase(repo repositories.RoomRepository) *RoomUsecase {
	return &RoomUsecase{
		repo: repo,
	}
}

func (uc *RoomUsecase) GetRooms(ctx context.Context, filter domain.Map) ([]*models.Room, error) {
	rooms, err := uc.repo.GetRooms(ctx, filter)
	if err != nil {
		return nil, err
	}
	return rooms, nil
}

func (uc *RoomUsecase) CreateRoom(ctx context.Context, room *models.Room) (*models.Room, error) {
	createRoom, err := uc.repo.CreateRoom(ctx, room)
	if err != nil {
		return nil, err
	}
	return createRoom, nil
}
