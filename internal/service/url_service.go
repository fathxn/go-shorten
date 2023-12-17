package service

import (
	"context"
	"errors"
	"fmt"
	"go-short-url/internal/model/domain"
	"go-short-url/internal/repository"
	"go-short-url/util"
	"gorm.io/gorm"
)

type urlService struct {
	URLRepository repository.URLRepository
}

type URLService interface {
	Create(ctx context.Context, longURL string) (*domain.URL, error)
	GetLongURL(ctx context.Context, shortCode string) (*domain.URL, error)
	GetById(ctx context.Context, id int) (*domain.URL, error)
	// GetByUserId(ctx context.Context, userId string) (*[]domain.URL, error)
	Delete(ctx context.Context, id int) error
}

func NewURLService(urlRepository repository.URLRepository) URLService {
	return &urlService{URLRepository: urlRepository}
}

func (URLService *urlService) Create(ctx context.Context, longURL string) (*domain.URL, error) {
	shortCode, err := util.GenerateUniqueCode(URLService.isShortCodeUnique)
	if err != nil {
		return nil, fmt.Errorf("failed to generate short code: %v", err)
	}
	shortUrl := &domain.URL{
		LongURL:   longURL,
		ShortCode: shortCode,
	}
	err = URLService.URLRepository.Insert(ctx, shortUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to create short URL: %v", err)
	}
	return shortUrl, nil
}

func (URLService *urlService) GetLongURL(ctx context.Context, shortCode string) (*domain.URL, error) {
	shortURL, err := URLService.URLRepository.FindByShortCode(ctx, shortCode)
	if err != nil {
		return nil, err
	}
	return shortURL, nil
}

func (URLService *urlService) GetById(ctx context.Context, id int) (*domain.URL, error) {
	shortURL, err := URLService.URLRepository.FindById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("no data found with id: %v", id)
	}
	return shortURL, nil
}

func (URLService *urlService) Delete(ctx context.Context, id int) error {
	err := URLService.URLRepository.Delete(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		return err
	}
	return nil
}

func (URLService *urlService) isShortCodeUnique(shortCode string) bool {
	_, err := URLService.URLRepository.FindByShortCode(context.Background(), shortCode)
	return err != nil
}
