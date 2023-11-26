package repository

import (
	"context"
	"go-short-url/internal/model/domain"
	"gorm.io/gorm"
)

type URLRepositoryImpl struct {
	DB *gorm.DB
}

func NewURLRepository(db *gorm.DB) URLRepository {
	return &URLRepositoryImpl{DB: db}
}

func (URLRepository *URLRepositoryImpl) Insert(ctx context.Context, url *domain.URL) error {
	err := URLRepository.DB.WithContext(ctx).Create(url).Error
	if err != nil {
		return err
	}

	return nil
}

func (URLRepository *URLRepositoryImpl) FindByShortCode(ctx context.Context, shortCode string) (*domain.URL, error) {
	var shortURL domain.URL
	err := URLRepository.DB.WithContext(ctx).Where("short_code = ?", shortCode).First(&shortURL).Error
	if err != nil {
		return nil, err
	}
	return &shortURL, nil
}

func (URLRepository *URLRepositoryImpl) FindById(ctx context.Context, id int) (*domain.URL, error) {
	var URL domain.URL
	err := URLRepository.DB.WithContext(ctx).Where("id = ?", id).First(&URL).Error
	if err != nil {
		return nil, err
	}
	return &URL, err
}

func (URLRepository *URLRepositoryImpl) Delete(ctx context.Context, id int) error {
	var URL domain.URL
	err := URLRepository.DB.WithContext(ctx).Where("id = ?", id).Delete(&URL).First(&URL).Error
	if err != nil {
		return err
	}

	return nil
}
