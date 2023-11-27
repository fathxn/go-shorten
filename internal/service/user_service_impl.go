package service

import (
	"context"
	"go-short-url/internal/model/domain"
	"go-short-url/internal/repository"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &UserServiceImpl{UserRepository: userRepository}
}

func (UserService *UserServiceImpl) Create(ctx context.Context, user *domain.User) error {
	//TODO implement me
	panic("implement me")
}

func (UserService *UserServiceImpl) FindById(ctx context.Context, id string) (*domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (UserService *UserServiceImpl) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}
