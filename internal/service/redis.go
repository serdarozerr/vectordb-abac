package service

import (
	"context"
	"github.com/serdarozerr/vectordb-abac/internal/repository"
	"time"
)

const EXP = time.Duration(time.Minute * 10)

// SaveRefreshToken, save refresh token to the MEMDB
func SaveRefreshToken(ctx context.Context, c repository.Cache, user string, token string) error {

	err := c.Set(ctx, user, token, EXP)
	return err
}

// RefreshToken, get refresh token from the MEMDB
func RefreshToken(ctx context.Context, c repository.Cache, code string) (string, error) {
	v, err := c.Get(ctx, code)
	return v.(string), err
}

func SavePublicKey(ctx context.Context, c repository.Cache, user string, pk string) error {
	err := c.Set(ctx, user, pk, EXP)
	return err
}

func PublicKey(ctx context.Context, c repository.Cache, user string) (string, error) {
	v, err := c.Get(ctx, user)
	return v.(string), err
}
