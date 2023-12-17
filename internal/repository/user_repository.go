package repository

import (
	"context"
	"go-short-url/internal/model/domain"
	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

type UserRepository interface {
	Insert(ctx context.Context, user *domain.User) error
	FindById(ctx context.Context, id string) (*domain.User, error)
	FindByEmail(ctx context.Context, email string) (*domain.User, error)
	Delete(ctx context.Context, id string) error
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{DB: db}
}

func (UserRepository *userRepository) Insert(ctx context.Context, user *domain.User) error {
	err := UserRepository.DB.WithContext(ctx).Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (UserRepository *userRepository) FindById(ctx context.Context, id string) (*domain.User, error) {
	var User domain.User
	err := UserRepository.DB.WithContext(ctx).Where("id = ?", id).First(&User).Error
	if err != nil {
		return nil, err
	}
	return &User, nil
}

func (UserRepository *userRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	var User domain.User
	err := UserRepository.DB.WithContext(ctx).Where("email = ?", email).First(&User).Error
	if err != nil {
		return nil, err
	}
	return &User, nil
}

func (UserRepository *userRepository) Delete(ctx context.Context, id string) error {
	var User domain.User
	err := UserRepository.DB.WithContext(ctx).Where("id = ?", id).Delete(&User).Error
	if err != nil {
		return err
	}
	return nil
}
