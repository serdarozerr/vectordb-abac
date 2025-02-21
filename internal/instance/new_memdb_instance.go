package instance

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/serdarozerr/vectordb-abac/config"
	"github.com/serdarozerr/vectordb-abac/internal/repository"
)

func NewMemeDbInstance(conf *config.Config, dbType string) (repository.Cache, error) {

	switch dbType {
	case "redis":
		opt := &redis.Options{}
		opt.Addr = fmt.Sprintf(conf.Redis.URL + ":" + conf.Redis.Port)
		opt.Password = conf.Redis.Password
		opt.DB = conf.Redis.DB
		r := redis.NewClient(opt)
		ctx := context.Background()

		_, err := r.Ping(ctx).Result()
		if err != nil {
			return nil, err
		}

		return &repository.RedisRepository{Client: r}, nil
	default:
		return nil, fmt.Errorf("unsupported mem db type: %s", dbType)
	}

}
