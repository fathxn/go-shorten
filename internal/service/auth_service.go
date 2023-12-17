package service

import (
	"context"
	"github.com/google/uuid"
	"go-short-url/internal/model/domain"
	"go-short-url/internal/model/dto"
	"go-short-url/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	UserRepository repository.UserRepository
}

type AuthService interface {
	RegisterUser(ctx context.Context, registerInput *dto.UserRegisterInput) error
	LoginUser(ctx context.Context, loginInput *dto.UserLoginInput) (*domain.User, error)
}

func NewAuthService(UserRepository repository.UserRepository) AuthService {
	return &authService{UserRepository: UserRepository}
}

func (AuthService *authService) RegisterUser(ctx context.Context, registerInput *dto.UserRegisterInput) error {
	generateUUID := uuid.New()
	HashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerInput.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user := &domain.User{
		Id:           generateUUID.String(),
		Name:         registerInput.Name,
		Email:        registerInput.Email,
		PasswordHash: string(HashedPassword),
	}
	err = AuthService.UserRepository.Insert(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (AuthService *authService) LoginUser(ctx context.Context, loginInput *dto.UserLoginInput) (*domain.User, error) {
	email := loginInput.Email
	password := loginInput.Password
	user, err := AuthService.UserRepository.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return nil, err
	}
	return user, nil
}
