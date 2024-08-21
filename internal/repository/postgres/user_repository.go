package postgres

import (
	"context"
	"database/sql"
	"go-shorten/internal/model/domain"
	"time"

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
		INSERT INTO users (id, name, email, password_hash, is_verified, verification_token, verification_token_expires_at, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`

	_, err := r.db.ExecContext(
		ctx,
		query,
		user.Id,
		user.Name,
		user.Email,
		user.PasswordHash,
		user.IsVerified,
		user.VerificationToken,
		user.VerificationTokenExpiresAt,
		user.CreatedAt,
		user.UpdatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}

// UpdateVerificationStatus implements domain.UserRepository.
func (r *userRepository) UpdateVerificationStatus(ctx context.Context, userId uuid.UUID, verifiedAt *time.Time) error {
	query := `
		UPDATE users
		SET is_verified = true, verified_at = $2
		WHERE id = $1
	`
	_, err := r.db.ExecContext(ctx, query, userId, verifiedAt)

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
	var user domain.User
	query := `
		SELECT id, name, email, is_verified
		FROM users
		WHERE email = $1
	`
	err := r.db.GetContext(ctx, &user, query, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

// GetById implements domain.UserRepository.
func (r *userRepository) GetById(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	panic("unimplemented")
}

// GetByVerificationToken implements domain.UserRepository.
func (r *userRepository) GetByVerificationToken(ctx context.Context, token string) (*domain.User, error) {
	panic("unimplemented")
}
