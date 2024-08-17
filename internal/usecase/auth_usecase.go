package usecase

import (
	"context"
	"go-shorten/internal/model/domain"
	"go-shorten/internal/model/dto"

	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase interface {
	RegisterUser(ctx context.Context, registerInput *dto.UserRegisterInput) error
	LoginUser(ctx context.Context, loginInput *dto.UserLoginInput) (*domain.User, error)
}

type authService struct {
	UserRepository domain.UserRepository
}

func NewAuthService(UserRepository domain.UserRepository) AuthUsecase {
	return &authService{UserRepository: UserRepository}
}

func (s *authService) RegisterUser(ctx context.Context, registerInput *dto.UserRegisterInput) error {
	// generateUUID := uuid.New()
	HashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerInput.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user := &domain.User{
		Name:         registerInput.Name,
		Email:        registerInput.Email,
		PasswordHash: string(HashedPassword),
	}
	err = s.UserRepository.Create(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (s *authService) LoginUser(ctx context.Context, loginInput *dto.UserLoginInput) (*domain.User, error) {
	email := loginInput.Email
	password := loginInput.Password
	user, err := s.UserRepository.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return nil, err
	}
	return user, nil
}
