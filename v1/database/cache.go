package database

import (
	"github.com/go-redis/redis"
)

var Cache CacheItf

func InitRedisCache() {
	Cache = &RedisCache{
		client: redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "", // no password set
			DB:       0,  // use default DB
		}),
	}

}
