package service

import (
	"context"
	"go-shorten/internal/model/domain"
	"go-shorten/internal/repository"
)

type userService struct {
	UserRepository repository.UserRepository
	URLRepository  domain.URLRepository
}

type UserService interface {
	GetById(ctx context.Context, id string) (*domain.User, error)
	GetURLsByUserId(ctx context.Context, userId string) (*[]domain.URL, error)
	Delete(ctx context.Context, id string) error
}

func NewUserService(userRepository repository.UserRepository, urlRepository domain.URLRepository) UserService {
	return &userService{UserRepository: userRepository, URLRepository: urlRepository}
}

func (s *userService) GetById(ctx context.Context, id string) (*domain.User, error) {
	user, err := s.UserRepository.FindById(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) GetURLsByUserId(ctx context.Context, userId string) (*[]domain.URL, error) {
	user, err := s.URLRepository.FindByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *userService) Delete(ctx context.Context, id string) error {
	err := s.UserRepository.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
