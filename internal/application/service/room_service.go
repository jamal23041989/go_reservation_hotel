package service

import (
	"context"
	"github.com/jamal23041989/go_reservation_hotel/internal/domain"
	"github.com/jamal23041989/go_reservation_hotel/internal/domain/entity"
	"github.com/jamal23041989/go_reservation_hotel/internal/domain/repository"
)

type RoomService struct {
	repo repository.RoomRepository
}

func NewRoomService(repo repository.RoomRepository) *RoomService {
	return &RoomService{
		repo: repo,
	}
}

func (uc *RoomService) GetRooms(ctx context.Context, filter domain.Map) ([]*entity.Room, error) {
	rooms, err := uc.repo.GetRooms(ctx, filter)
	if err != nil {
		return nil, err
	}
	return rooms, nil
}

func (uc *RoomService) CreateRoom(ctx context.Context, room *entity.Room) (*entity.Room, error) {
	createRoom, err := uc.repo.CreateRoom(ctx, room)
	if err != nil {
		return nil, err
	}
	return createRoom, nil
}
