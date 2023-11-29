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

func (UserService *UserServiceImpl) RegisterUser(ctx context.Context, registerInput *api.UserRegisterInput) error {
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
	err = UserService.UserRepository.Insert(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (UserService *UserServiceImpl) LoginUser(ctx context.Context, loginInput *api.UserLoginInput) (*domain.User, error) {
	email := loginInput.Email
	password := loginInput.Password
	user, err := UserService.UserRepository.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (UserService *UserServiceImpl) FindById(ctx context.Context, id string) (*domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (UserService *UserServiceImpl) Delete(ctx context.Context, id string) error {
	err := UserService.UserRepository.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
