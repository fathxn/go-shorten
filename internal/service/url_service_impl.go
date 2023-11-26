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

type URLServiceImpl struct {
	URLRepository repository.URLRepository
}

func NewURLService(urlRepository repository.URLRepository) URLService {
	return &URLServiceImpl{URLRepository: urlRepository}
}

func (URLService *URLServiceImpl) Create(ctx context.Context, longURL string) (*domain.URL, error) {
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

func (URLService *URLServiceImpl) GetLongURL(ctx context.Context, shortCode string) (*domain.URL, error) {
	shortURL, err := URLService.URLRepository.FindByShortCode(ctx, shortCode)
	if err != nil {
		return nil, err
	}
	return shortURL, nil
}

func (URLService *URLServiceImpl) GetById(ctx context.Context, id int) (*domain.URL, error) {
	shortURL, err := URLService.URLRepository.FindById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("no data found with id: %v", id)
	}
	return shortURL, nil
}

func (URLService *URLServiceImpl) Delete(ctx context.Context, id int) error {
	err := URLService.URLRepository.Delete(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		return err
	}

	return nil
}

func (URLService *URLServiceImpl) isShortCodeUnique(shortCode string) bool {
	_, err := URLService.URLRepository.FindByShortCode(context.Background(), shortCode)
	return err != nil
}
