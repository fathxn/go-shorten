package repository

import (
	"context"
	"go-shorten/internal/model/domain"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

type UserRepository interface {
	Insert(ctx context.Context, user *domain.User) error
	FindById(ctx context.Context, id string) (*domain.User, error)
	FindByEmail(ctx context.Context, email string) (*domain.User, error)
	Delete(ctx context.Context, id string) error
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Insert(ctx context.Context, user *domain.User) error {
	err := r.db.WithContext(ctx).Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepository) FindById(ctx context.Context, id string) (*domain.User, error) {
	var User domain.User
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&User).Error
	if err != nil {
		return nil, err
	}
	return &User, nil
}

func (r *userRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	var User domain.User
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&User).Error
	if err != nil {
		return nil, err
	}
	return &User, nil
}

func (r *userRepository) Delete(ctx context.Context, id string) error {
	var User domain.User
	err := r.db.WithContext(ctx).Where("id = ?", id).Delete(&User).Error
	if err != nil {
		return err
	}
	return nil
}
