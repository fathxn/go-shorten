package repository

import (
	"context"
	"go-short-url/internal/model/domain"

	"gorm.io/gorm"
)

type urlRepository struct {
	db *gorm.DB
}

func NewURLRepository(db *gorm.DB) domain.URLRepository {
	return &urlRepository{db: db}
}

func (r *urlRepository) Insert(ctx context.Context, url *domain.URL) error {
	if err := r.db.WithContext(ctx).Create(url).Error; err != nil {
		return err
	}
	return nil
}

func (r *urlRepository) FindByShortCode(ctx context.Context, shortCode string) (*domain.URL, error) {
	var shortURL domain.URL
	if err := r.db.WithContext(ctx).Where("short_code = ?", shortCode).First(&shortURL).Error; err != nil {
		return nil, err
	}
	return &shortURL, nil
}

func (r *urlRepository) FindById(ctx context.Context, id int) (*domain.URL, error) {
	var URL domain.URL
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&URL).Error; err != nil {
		return nil, err
	}
	return &URL, nil
}

func (r *urlRepository) FindByUserId(ctx context.Context, userId string) ([]domain.URL, error) {
	var URL []domain.URL
	if err := r.db.WithContext(ctx).Where("user_id = ?", userId).Find(&URL).Error; err != nil {
		return nil, err
	}
	return URL, nil
}

func (r *urlRepository) Delete(ctx context.Context, id int) error {
	var URL domain.URL
	if err := r.db.WithContext(ctx).Where("id = ?", id).Delete(&URL).Error; err != nil {
		return err
	}
	return nil
}
