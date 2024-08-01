package usecase

import (
	"context"
	"errors"
	"fmt"
	"go-shorten/internal/model/domain"
	"go-shorten/util"

	"gorm.io/gorm"
)

type urlService struct {
	URLRepository domain.URLRepository
}

func NewURLService(urlRepository domain.URLRepository) domain.URLService {
	return &urlService{URLRepository: urlRepository}
}

func (s *urlService) Create(ctx context.Context, longURL string, userId string) (*domain.URL, error) {
	shortCode, err := util.GenerateUniqueCode(s.isShortCodeUnique)
	if err != nil {
		return nil, fmt.Errorf("failed to generate short code: %v", err)
	}
	shortUrl := &domain.URL{
		UserId:    userId,
		LongURL:   longURL,
		ShortCode: shortCode,
	}
	err = s.URLRepository.Insert(ctx, shortUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to create short URL: %v", err)
	}
	return shortUrl, nil
}

func (s *urlService) GetLongURL(ctx context.Context, shortCode string) (*domain.URL, error) {
	shortURL, err := s.URLRepository.FindByShortCode(ctx, shortCode)
	if err != nil {
		return nil, err
	}
	return shortURL, nil
}

func (s *urlService) GetById(ctx context.Context, id int) (*domain.URL, error) {
	shortURL, err := s.URLRepository.FindById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("no data found with id: %v", id)
	}
	return shortURL, nil
}

func (s *urlService) Delete(ctx context.Context, id int) error {
	err := s.URLRepository.Delete(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		return err
	}
	return nil
}

func (s *urlService) isShortCodeUnique(shortCode string) bool {
	_, err := s.URLRepository.FindByShortCode(context.Background(), shortCode)
	return err != nil
}
