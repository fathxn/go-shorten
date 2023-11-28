package service

import (
	"context"
	"github.com/google/uuid"
	"go-short-url/internal/model/api"
	"go-short-url/internal/model/domain"
	"go-short-url/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &UserServiceImpl{UserRepository: userRepository}
}

func (UserService *UserServiceImpl) Create(ctx context.Context, userInput *api.UserInput) error {
	generateUUID := uuid.New()
	HashedPassword, err := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user := &domain.User{
		Id:           generateUUID.String(),
		Name:         userInput.Name,
		Email:        userInput.Email,
		PasswordHash: string(HashedPassword),
	}
	err = UserService.UserRepository.Insert(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (UserService *UserServiceImpl) FindById(ctx context.Context, id string) (*domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (UserService *UserServiceImpl) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}
