package usecase

import (
	"context"
	"errors"
	"go-shorten/internal/model/domain"
	"go-shorten/internal/model/dto"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase interface {
	VerifyEmail(ctx context.Context, token string) error
	RegisterUser(ctx context.Context, registerInput *dto.UserRegisterInput) error
	LoginUser(ctx context.Context, loginInput *dto.UserLoginInput) (*domain.User, error)
}

type authUsecase struct {
	UserRepository domain.UserRepository
}

func NewAuthUsecase(UserRepository domain.UserRepository) AuthUsecase {
	return &authUsecase{UserRepository: UserRepository}
}

func (u *authUsecase) VerifyEmail(ctx context.Context, token string) error {
	user, err := u.UserRepository.GetByVerificationToken(ctx, token)
	if err != nil {
		return err
	}

	if user == nil {
		return errors.New("verification token has expired")
	}

	return u.UserRepository.UpdateVerificationStatus(ctx, user.Id, true)
}

func (u *authUsecase) RegisterUser(ctx context.Context, registerInput *dto.UserRegisterInput) error {
	// check email is used/registered/stored in database
	emailIsExist, err := u.UserRepository.GetByEmail(ctx, registerInput.Email)
	if err != nil {
		return err
	}
	// check if email is not nil or already stored in database
	if emailIsExist != nil {
		return errors.New("email is already used")
	}

	// hash/encrypt the password from input request
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerInput.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &domain.User{
		Id:           uuid.New(),
		Name:         registerInput.Name,
		Email:        registerInput.Email,
		PasswordHash: string(hashedPassword),
	}

	// create the user/store user data in repository database
	err = u.UserRepository.Create(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (u *authUsecase) LoginUser(ctx context.Context, loginInput *dto.UserLoginInput) (*domain.User, error) {
	email := loginInput.Email
	password := loginInput.Password
	user, err := u.UserRepository.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return nil, err
	}
	return user, nil
}
