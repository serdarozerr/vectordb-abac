package handler

import (
	"github.com/serdarozerr/vectordb-abac/internal/repository"
	"log"
	"net/http"
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
