package usecase

import (
	"context"
	"go-shorten/internal/model/domain"
	"go-shorten/internal/model/dto"
)

type userUsecase struct {
	UserRepository domain.UserRepository
	URLRepository  domain.URLRepository
}

func NewUserService(userRepository domain.UserRepository, urlRepository domain.URLRepository) domain.UserService {
	return &userUsecase{UserRepository: userRepository, URLRepository: urlRepository}
}

func (s *userUsecase) RegisterUser(ctx context.Context, registerUser *dto.UserRegisterInput) error {
	return nil
}

func (s *userUsecase) VerifyEmail(ctx context.Context, token string) error {
	return nil
}

func (s *userUsecase) GetById(ctx context.Context, id string) (*domain.User, error) {
	user, err := s.UserRepository.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userUsecase) GetURLsByUserId(ctx context.Context, userId string) (*[]domain.URL, error) {
	user, err := s.URLRepository.FindByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *userUsecase) Delete(ctx context.Context, id string) error {
	err := s.UserRepository.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
