package usecase

import (
	"context"
	"go-shorten/internal/model/domain"

	"github.com/google/uuid"
)

type userUsecase struct {
	UserRepository domain.UserRepository
	URLRepository  domain.URLRepository
}

func NewUserUsecase(userRepository domain.UserRepository, urlRepository domain.URLRepository) domain.UserUsecase {
	return &userUsecase{
		UserRepository: userRepository,
		URLRepository:  urlRepository,
	}
}

func (u *userUsecase) GetById(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	user, err := u.UserRepository.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userUsecase) GetURLsByUserId(ctx context.Context, userId string) (*[]domain.URL, error) {
	return nil, nil
}

func (s *userUsecase) Delete(ctx context.Context, id uuid.UUID) error {
	err := s.UserRepository.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
