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

// Insert implements domain.UserRepository.
func (r *userRepository) Create(ctx context.Context, user *domain.User) error {
	query := `
		INSERT INTO users (name, email, password_hash, is_verified, verification_token, verification_token_expires_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, created_at, updated_at
	`
	return r.db.QueryRowxContext(ctx, query,
		user.Name,
		user.Email,
		user.PasswordHash,
		user.IsVerified,
		user.VerificationToken,
		user.VerificationTokenExpiresAt,
	).Scan(&user.Id, &user.CreatedAt, &user.UpdatedAt)
}

// FindById implements domain.UserRepository.
func (r *userRepository) GetById(ctx context.Context, id string) (*domain.User, error) {
	var user domain.User
	query := `
		SELECT name, email
		WHERE id = $1
		LIMIT 1;
	`

	if err := r.db.GetContext(ctx, &user, query, id); err != nil {
		return nil, err
	}

	return &user, nil
}

// FindByEmail implements domain.UserRepository.
func (r *userRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	var user domain.User
	query := `
		SELECT name, email
		WHERE email = $1
		LIMIT 1;
	`

	if err := r.db.GetContext(ctx, &user, query, email); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) GetByVerificationToken(ctx context.Context, token string) (*domain.User, error) {
	var user domain.User
	query := `SELECT name, email, is_verified WHERE verification_token = $1 LIMIT 1;`
	if err := r.db.GetContext(ctx, &user, query, token); err != nil {
		return nil, err
	}

	return &user, nil
}
func (r *userRepository) UpdateVerificationStatus(ctx context.Context, userId string, isVerified bool) error {
	query := `
		UPDATE users
		SET is_verified = $1, updated_at = NOW()
		WHERE id = $2
	`
	_, err := r.db.ExecContext(ctx, query, isVerified, userId)
	if err != nil {
		return err
	}
	return nil
}

// Delete implements domain.UserRepository.
func (r *userRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}
