package postgres

import (
	"context"
	"go-shorten/internal/model/domain"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) domain.UserRepository {
	return &userRepository{db: db}
}

// Create implements domain.UserRepository.
func (r *userRepository) Create(ctx context.Context, user *domain.User) error {
	query := `
		INSERT INTO users (name, email, password_hash, is_verified, verified_at, verification_token, verification_token_expires_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, created_at, updated_at
	`

	err := r.db.QueryRowContext(
		ctx,
		query,
		user.Name,
		user.Email,
		user.PasswordHash,
		user.IsVerified,
		user.VerifiedAt,
		user.VerificationToken,
		user.VerificationTokenExpiresAt,
	).Scan(&user.Id, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return err
	}

	return nil
}

// Delete implements domain.UserRepository.
func (r *userRepository) Delete(ctx context.Context, id uuid.UUID) error {
	panic("unimplemented")
}

// GetByEmail implements domain.UserRepository.
func (r *userRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	panic("unimplemented")
}

// GetById implements domain.UserRepository.
func (r *userRepository) GetById(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	panic("unimplemented")
}

// GetByVerificationToken implements domain.UserRepository.
func (r *userRepository) GetByVerificationToken(ctx context.Context, token string) (*domain.User, error) {
	panic("unimplemented")
}

// UpdateVerificationStatus implements domain.UserRepository.
func (r *userRepository) UpdateVerificationStatus(ctx context.Context, userId uuid.UUID, isVerified bool) error {
	panic("unimplemented")
}
