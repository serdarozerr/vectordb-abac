package handler

import (
	"errors"
	"github.com/serdarozerr/vectordb-abac/config"
	"github.com/serdarozerr/vectordb-abac/internal/repository"
	"github.com/serdarozerr/vectordb-abac/internal/service/auth"
	"golang.org/x/net/context"
	"log"
	"net/http"
	"strings"
)

// it is gonna take some common dependencies and gonna return handler
type Middleware func(handler http.Handler) http.Handler

func AdminMiddleware(logger *log.Logger, db *repository.VectorRepository) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger.Println("Admin middleware called...")
			//TODO check the current user is admin
			next.ServeHTTP(w, r)
		})

	}
}

func LoggerMiddleware(logger *log.Logger) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger.Println("Logger middleware called on the endpoint %s", r.RequestURI)
			next.ServeHTTP(w, r)
		})

	}

}

func ExtractBearerToken(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("authorization header missing")
	}

	// Expected format: "Bearer <token>"
	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return "", errors.New("invalid authorization header format")
	}

	return parts[1], nil
}

func AuthenticationMiddleware(logger *log.Logger, conf *config.Config, c repository.Cache) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			access_token, err := ExtractBearerToken(r)
			if err != nil {
				logger.Printf("Error decoding token: %v", err)
				return
			}
			claims, err := auth.DecodeToken(r.Context(), conf, c, access_token)
			if err != nil {
				logger.Printf("Error decoding token: %v", err)
				return
			}

			ctx := context.WithValue(r.Context(), "claims", claims)

			next.ServeHTTP(w, r.WithContext(ctx))

		})
	}

}
