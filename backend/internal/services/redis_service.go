package services

import (
	"context"
	"enroll-tracker/internal/repositories"
	"time"
)

type RedisService struct {
	repository repositories.CacheRepository
}

func CreateRedisService(repo *repositories.RedisRepository) *RedisService {
	return &RedisService{repository: repo}
}

func (s *RedisService) Set(ctx context.Context, key string, value interface{}, EX time.Duration) error {
	err := s.repository.Set(ctx, key, value, EX)
	return err
}

func (s *RedisService) Get(ctx context.Context, key string) (interface{}, error) {
	val, err := s.repository.Get(ctx, key)
	if err != nil {
		return nil, err
	}
	return val, nil
}
