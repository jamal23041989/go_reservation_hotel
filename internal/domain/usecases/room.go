package usecases

import (
	"context"
	"github.com/jamal23041989/go_reservation_hotel/internal/domain"
	"github.com/jamal23041989/go_reservation_hotel/internal/domain/models"
)

type RoomUsecase interface {
	GetRooms(context.Context, domain.Map) ([]*models.Room, error)
	CreateRoom(context.Context, *models.Room) (*models.Room, error)
}
