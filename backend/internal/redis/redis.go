package redis

import (
	"errors"
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

func CreateNewRedisClient() (*redis.Client, error) {
	pass, ok := os.LookupEnv("REDIS_PASSWORD")
	if !ok {
		return nil, errors.New("Can't connect to Redis")
	}

	hostPort, ok := os.LookupEnv("REDIS_HOST_PORT")
	if !ok {
		return nil, errors.New("Can't connect to Redis")
	}

	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("redis:%s", hostPort), //container name:port
		Password: pass,
		DB:       0, //default db
	})

	return client, nil
}
