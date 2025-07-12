package service

import (
	"context"
	"github.com/jamal23041989/go_reservation_hotel/internal/domain"
	"github.com/jamal23041989/go_reservation_hotel/internal/domain/entity"
)

type UserService interface {
	GetUserByID(context.Context, string) (*entity.User, error)
	GetUsers(context.Context) ([]*entity.User, error)
	GetByEmailUser(context.Context, string) (*entity.User, error)
	CreateUser(context.Context, *entity.User) (*entity.User, error)
	UpdateUser(context.Context, string, domain.Map) error
	DeleteUser(context.Context, string) error
}
