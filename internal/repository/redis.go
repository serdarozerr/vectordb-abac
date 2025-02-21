package repository

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type RedisRepository struct {
	Client *redis.Client
}

func (r *RedisRepository) Get(ctx context.Context, key interface{}) (interface{}, error) {
	v, err := r.Client.Get(ctx, key.(string)).Result()
	if err != nil {
		return "", err
	}
	return v, nil
}

func (r *RedisRepository) Set(ctx context.Context, key interface{}, value interface{}, exp time.Duration) error {
	_, err := r.Client.Set(ctx, key.(string), value, exp).Result()
	return err
}

func (r *RedisRepository) Delete(ctx context.Context, key interface{}) error {
	_, err := r.Client.Del(ctx, key.(string)).Result()
	return err
}
