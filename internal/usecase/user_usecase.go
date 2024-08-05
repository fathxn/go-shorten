package usecase

import (
	"context"
	"errors"
	"fmt"
	"go-shorten/internal/model/domain"
	"go-shorten/internal/model/dto"
	"go-shorten/pkg/email"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	UserRepository domain.UserRepository
	URLRepository  domain.URLRepository
	emailService   email.EmailService
}

func NewUserUsecase(userRepository domain.UserRepository, urlRepository domain.URLRepository) domain.UserUsecase {
	return &userUsecase{UserRepository: userRepository, URLRepository: urlRepository}
}

func (u *userUsecase) RegisterUser(ctx context.Context, registerInput *dto.UserRegisterInput) error {
	generateUUID := uuid.New()
	HashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerInput.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	verificationToken := generateUUID.String()
	expiresAt := time.Now().Add(1 * time.Hour)

	user := &domain.User{
		Id:                         generateUUID.String(),
		Name:                       registerInput.Name,
		Email:                      registerInput.Email,
		PasswordHash:               string(HashedPassword),
		VerificationToken:          verificationToken,
		VerificationTokenExpiresAt: expiresAt,
	}

	err = u.UserRepository.Create(ctx, user)
	if err != nil {
		return err
	}

	verificationLink := "localhost" + "/verify?token=" + verificationToken
	if err := u.emailService.SendVerificationEmail(registerInput.Email, "Email Verification", verificationLink); err != nil {
		return fmt.Errorf("failed to send email verification: %w", err)
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

func (u *userUsecase) GetById(ctx context.Context, id string) (*domain.User, error) {
	user, err := u.UserRepository.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userUsecase) GetURLsByUserId(ctx context.Context, userId string) (*[]domain.URL, error) {
	user, err := u.URLRepository.FindByUserId(ctx, userId)
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
