package usecase

import (
	"context"
	"github.com/jamal23041989/go_reservation_hotel/internal/domain"
	"github.com/jamal23041989/go_reservation_hotel/internal/domain/models"
	"github.com/jamal23041989/go_reservation_hotel/internal/domain/repositories"
)

type UserUsecase struct {
	repo repositories.UserRepository
}

func NewUserUsecase(repo repositories.UserRepository) *UserUsecase {
	return &UserUsecase{
		repo: repo,
	}
}

func (uc *UserUsecase) GetUserByID(ctx context.Context, id string) (*models.User, error) {
	user, err := uc.repo.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (uc *UserUsecase) GetUsers(ctx context.Context) ([]*models.User, error) {
	users, err := uc.repo.GetUsers(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (uc *UserUsecase) GetByEmailUser(ctx context.Context, email string) (*models.User, error) {
	user, err := uc.repo.GetByEmailUser(ctx, email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (uc *UserUsecase) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	createUser, err := uc.repo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}
	return createUser, nil
}

func (uc *UserUsecase) UpdateUser(ctx context.Context, id string, update domain.Map) error {
	return uc.repo.UpdateUser(ctx, id, update)
}

func (uc *UserUsecase) DeleteUser(ctx context.Context, id string) error {
	return uc.repo.DeleteUser(ctx, id)
}

func (uc *UserUsecase) DropUser(ctx context.Context) error {
	return uc.repo.DropUser(ctx)
}
