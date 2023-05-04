package database

import (
	"encoding/json"
	"time"

	"github.com/go-redis/redis"
)

type CacheItf interface {
	SetMovieListCache(key string, data interface{}, expiration time.Duration) error
	GetMovieListCache(key string) ([]byte, error)
}

type RedisCache struct {
	client *redis.Client
}

func (r *RedisCache) SetMovieListCache(key string, data interface{}, expiration time.Duration) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return r.client.Set(key, b, expiration).Err()
}

func (r *RedisCache) GetMovieListCache(key string) ([]byte, error) {
	result, err := r.client.Get(key).Bytes()
	if err == redis.Nil {
		return nil, nil
	}

	return result, err
}
