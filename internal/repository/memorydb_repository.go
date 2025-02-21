package repository

import (
	"context"
	"time"
)

type Cache interface {
	Set(ctx context.Context, key, value interface{}, exp time.Duration) error
	Get(ctx context.Context, key interface{}) (interface{}, error)
	Delete(ctx context.Context, key interface{}) error
}
