package service

import (
	"context"
	"go-shorten/internal/model/domain"
)

type userService struct {
	UserRepository domain.UserRepository
	URLRepository  domain.URLRepository
}

func NewUserService(userRepository domain.UserRepository, urlRepository domain.URLRepository) domain.UserService {
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
