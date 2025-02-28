package scalenet

import (
	"context"
	"github.com/redis/go-redis/v9"
)

type RedisConnector struct {
	client *redis.Client
	ctx    context.Context
}

func ConnectRedis() *RedisConnector {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	return &RedisConnector{
		client: rdb,
		ctx:    context.Background(),
	}
}

func (rc *RedisConnector) Set(key string, value string) error {
	return rc.client.Set(rc.ctx, key, value, 0).Err()
}

func (rc *RedisConnector) GetValues(key string) ([]string, error) {
	return rc.client.SMembers(rc.ctx, key).Result()
}
