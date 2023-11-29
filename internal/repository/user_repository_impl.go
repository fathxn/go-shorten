package repository

import (
	"context"
	"go-short-url/internal/model/domain"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{DB: db}
}

func (UserRepository *UserRepositoryImpl) Insert(ctx context.Context, user *domain.User) error {
	err := UserRepository.DB.WithContext(ctx).Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (UserRepository *UserRepositoryImpl) FindById(ctx context.Context, id string) (*domain.User, error) {
	var User domain.User
	err := UserRepository.DB.WithContext(ctx).Where("id = ?", id).First(&User).Error
	if err != nil {
		return nil, err
	}
	return &User, nil
}

func (UserRepository *UserRepositoryImpl) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	var User domain.User
	err := UserRepository.DB.WithContext(ctx).Where("email = ?", email).First(&User).Error
	if err != nil {
		return nil, err
	}
	return &User, nil
}

func (UserRepository *UserRepositoryImpl) Delete(ctx context.Context, id string) error {
	var User domain.User
	err := UserRepository.DB.WithContext(ctx).Where("id = ?", id).Delete(&User).Error
	if err != nil {
		return err
	}
	return nil
}
