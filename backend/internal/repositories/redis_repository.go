package repositories

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type CacheRepository interface {
	Set(ctx context.Context, key string, value interface{}, EX time.Duration) error
	Get(ctx context.Context, key string) (interface{}, error)
}

type RedisRepository struct {
	redis *redis.Client
}

func CreateRedisRepository(redis *redis.Client) *RedisRepository {
	return &RedisRepository{redis: redis}
}

func (r *RedisRepository) Set(ctx context.Context, key string, value interface{}, EX time.Duration) error {
	err := r.redis.Set(ctx, key, value, EX).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisRepository) Get(ctx context.Context, key string) (interface{}, error) {
	value, err := r.redis.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	return value, nil
}
