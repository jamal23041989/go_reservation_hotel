package service

import (
	"context"
	"github.com/jamal23041989/go_reservation_hotel/internal/domain"
	"github.com/jamal23041989/go_reservation_hotel/internal/domain/entity"
	"github.com/jamal23041989/go_reservation_hotel/internal/domain/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (uc *UserService) GetUserByID(ctx context.Context, id string) (*entity.User, error) {
	user, err := uc.repo.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (uc *UserService) GetUsers(ctx context.Context) ([]*entity.User, error) {
	users, err := uc.repo.GetUsers(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (uc *UserService) GetByEmailUser(ctx context.Context, email string) (*entity.User, error) {
	user, err := uc.repo.GetByEmailUser(ctx, email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (uc *UserService) CreateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	createUser, err := uc.repo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}
	return createUser, nil
}

func (uc *UserService) UpdateUser(ctx context.Context, id string, update domain.Map) error {
	return uc.repo.UpdateUser(ctx, id, update)
}

func (uc *UserService) DeleteUser(ctx context.Context, id string) error {
	return uc.repo.DeleteUser(ctx, id)
}

func (uc *UserService) DropUser(ctx context.Context) error {
	return uc.repo.DropUser(ctx)
}
