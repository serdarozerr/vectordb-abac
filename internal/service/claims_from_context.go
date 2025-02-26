package service

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/net/context"
)

func ValueFromContext[T any](ctx context.Context, key string) (T, bool) {
	claim, ok := ctx.Value(key).(T)
	return claim, ok
}

func RolesFromClaims(claims jwt.MapClaims) ([]interface{}, error) {

	roles, ok := claims["resource_access"].(map[string]interface{})["vector-abac"].(map[string]interface{})["roles"].([]interface{})
	if !ok {
		return nil, errors.New("no roles found in claims")
	}
	return roles, nil
}
