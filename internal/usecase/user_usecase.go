package usecase

import (
	"context"
	"errors"
	"go-shorten/internal/model/domain"
	"go-shorten/internal/model/dto"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
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

func (u *userUsecase) RegisterUser(ctx context.Context, registerInput *dto.UserRegisterInput) error {
	generateUUID := uuid.New()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerInput.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	verificationToken := generateUUID.String()
	expiresAt := time.Now().Add(1 * time.Hour)

	user := &domain.User{
		Name:                       registerInput.Name,
		Email:                      registerInput.Email,
		PasswordHash:               string(hashedPassword),
		VerificationToken:          verificationToken,
		VerificationTokenExpiresAt: expiresAt,
	}

	err = u.UserRepository.Create(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (u *userUsecase) VerifyEmail(ctx context.Context, token string) error {
	user, err := u.UserRepository.GetByVerificationToken(ctx, token)
	if err != nil {
		return err
	}

	if user == nil {
		return errors.New("verification token has expired")
	}

	return u.UserRepository.UpdateVerificationStatus(ctx, user.Id, true)
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
