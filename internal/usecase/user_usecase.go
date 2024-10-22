package usecase

import (
	"context"
	"go-shorten/internal/domain"
	"go-shorten/internal/domain/errors"

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

func (u *userUsecase) RegisterUser(ctx context.Context, registerInput *domain.UserRegisterInput) (string, error) {
	// check email is used/registered/stored in database
	emailIsExist, err := u.UserRepository.GetByEmail(ctx, registerInput.Email)
	if err != nil {
		return "", err
	}
	// check if email is not nil or already stored in database
	if emailIsExist != nil {
		return "", err
	}

	// hash/encrypt the password from input request
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerInput.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", nil
	}

	user := &domain.User{
		Id:           uuid.New().String(),
		Name:         registerInput.Name,
		Email:        registerInput.Email,
		PasswordHash: string(hashedPassword),
	}

	// create the user/store user data in repository database
	err = u.UserRepository.Create(ctx, user)
	if err != nil {
		return "", err
	}
	return user.VerificationToken, nil
}

func (u *userUsecase) VerifyEmail(ctx context.Context, token string) error {
	user, err := u.UserRepository.GetByVerificationToken(ctx, token)
	if err != nil {
		return err
	}

	if user == nil {
		return err
	}

	return nil
}

func (u *userUsecase) LoginUser(ctx context.Context, loginInput *domain.UserLoginInput) (*domain.User, error) {
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

func (u *userUsecase) GetById(ctx context.Context, id string) (*domain.User, error) {
	if id == "" {
		return nil, errors.NewInvalidInputError("user ID cannot be empty", nil)
	}

	user, err := u.UserRepository.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userUsecase) GetURLsByUserId(ctx context.Context, userId string) (*[]domain.URL, error) {
	return nil, nil
}

func (s *userUsecase) Delete(ctx context.Context, id string) error {
	err := s.UserRepository.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
