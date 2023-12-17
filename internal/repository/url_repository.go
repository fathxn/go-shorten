package repository

import (
	"context"
	"go-short-url/internal/model/domain"
	"gorm.io/gorm"
)

type urlRepository struct {
	DB *gorm.DB
}

type URLRepository interface {
	Insert(ctx context.Context, url *domain.URL) error
	FindByShortCode(ctx context.Context, shortCode string) (*domain.URL, error)
	FindById(ctx context.Context, id int) (*domain.URL, error)
	FindByUserId(ctx context.Context, userId string) ([]domain.URL, error)
	Delete(ctx context.Context, id int) error
}

func NewURLRepository(db *gorm.DB) URLRepository {
	return &urlRepository{DB: db}
}

func (URLRepository *urlRepository) Insert(ctx context.Context, url *domain.URL) error {
	err := URLRepository.DB.WithContext(ctx).Create(url).Error
	if err != nil {
		return err
	}
	return nil
}

func (URLRepository *urlRepository) FindByShortCode(ctx context.Context, shortCode string) (*domain.URL, error) {
	var shortURL domain.URL
	err := URLRepository.DB.WithContext(ctx).Where("short_code = ?", shortCode).First(&shortURL).Error
	if err != nil {
		return nil, err
	}
	return &shortURL, nil
}

func (URLRepository *urlRepository) FindById(ctx context.Context, id int) (*domain.URL, error) {
	var URL domain.URL
	err := URLRepository.DB.WithContext(ctx).Where("id = ?", id).First(&URL).Error
	if err != nil {
		return nil, err
	}
	return &URL, err
}

func (URLRepository *urlRepository) FindByUserId(ctx context.Context, userId string) ([]domain.URL, error) {
	var URL []domain.URL
	err := URLRepository.DB.WithContext(ctx).Where("user_id = ?", userId).Find(&URL).Error
	if err != nil {
		return nil, err
	}
	return URL, nil
}

func (URLRepository *urlRepository) Delete(ctx context.Context, id int) error {
	var URL domain.URL
	err := URLRepository.DB.WithContext(ctx).Where("id = ?", id).Delete(&URL).Error
	if err != nil {
		return err
	}

	return nil
}
