package service

import (
	"crypto/sha1"
	"encoding/hex"
	"time"

	"url-shortener/config"
	"url-shortener/models"
	"url-shortener/repository"

	"github.com/redis/go-redis/v9"
)

type URLService struct {
	Repo  *repository.URLRepository
	Redis *redis.Client
}

func NewURLService(repo *repository.URLRepository, redisClient *redis.Client) *URLService {
	return &URLService{
		Repo:  repo,
		Redis: redisClient,
	}
}

func (s *URLService) GetOriginalURL(code string) (*models.URL, error) {
	cached, err := s.Redis.Get(config.Ctx, code).Result()
	if err == nil {
		return &models.URL{ShortCode: code, OriginalURL: cached}, nil
	}
	if err != redis.Nil {
		return nil, err
	}

	url, err := s.Repo.FindByShortCode(code)
	if err != nil {
		return nil, err
	}

	_ = s.Redis.Set(config.Ctx, code, url.OriginalURL, 0).Err()
	return url, nil
}

func (s *URLService) CreateShortURL(original string) (*models.URL, error) {
	h := sha1.New()
	h.Write([]byte(original + time.Now().String()))
	code := hex.EncodeToString(h.Sum(nil))[:8]

	url := &models.URL{
		ShortCode:   code,
		OriginalURL: original,
	}

	if err := s.Repo.Save(url); err != nil {
		return nil, err
	}

	_ = s.Redis.Set(config.Ctx, code, original, 0).Err()

	return url, nil
}
