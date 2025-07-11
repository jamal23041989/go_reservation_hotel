package repositories

import (
	"context"
	"github.com/jamal23041989/go_reservation_hotel/internal/domain"
	"github.com/jamal23041989/go_reservation_hotel/internal/domain/models"
)

type UserRepository interface {
	GetUserByID(context.Context, string) (*models.User, error)
	GetUsers(context.Context) ([]*models.User, error)
	GetByEmailUser(context.Context, string) (*models.User, error)
	CreateUser(context.Context, *models.User) (*models.User, error)
	UpdateUser(context.Context, string, domain.Map) error
	DeleteUser(context.Context, string) error
	DropUser(context.Context) error
}
