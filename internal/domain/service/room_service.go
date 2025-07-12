package service

import (
	"context"
	"github.com/jamal23041989/go_reservation_hotel/internal/domain"
	"github.com/jamal23041989/go_reservation_hotel/internal/domain/entity"
)

type RoomService interface {
	GetRooms(context.Context, domain.Map) ([]*entity.Room, error)
	CreateRoom(context.Context, *entity.Room) (*entity.Room, error)
}
