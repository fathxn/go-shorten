package postgres

import (
	"context"
	"go-shorten/internal/model/domain"

	"github.com/jmoiron/sqlx"
)

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) domain.UserRepository {
	return &userRepository{db: db}
}

// Delete implements domain.UserRepository.
func (u *userRepository) Delete(ctx context.Context, id string) error {
	panic("unimplemented")
}

// FindByEmail implements domain.UserRepository.
func (u *userRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	panic("unimplemented")
}

// FindById implements domain.UserRepository.
func (u *userRepository) FindById(ctx context.Context, id string) (*domain.User, error) {
	panic("unimplemented")
}

// Insert implements domain.UserRepository.
func (u *userRepository) Insert(ctx context.Context, user *domain.User) error {
	panic("unimplemented")
}
