package service

import (
	"context"
	"log"
	"net/http"
	"qdrant-abac/internal/db"
	"regexp"
)

type QdrantCreate struct {
	Name string `json:"name"`
}

var reg, err = regexp.Compile(`^[a-zA-Z]+$`)

func (qc *QdrantCreate) Valid(ctx context.Context) map[string]string {
	reg.MatchString(qc.Name)
	problems := make(map[string]string)
	if !reg.MatchString(qc.Name) {
		problems[qc.Name] = "includes non letter characters"
	}

	return problems
}

func Create(logger *log.Logger, db db.DBClient) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Println("Creating QDRant Collection...")

		w.WriteHeader(200)
		w.Write([]byte("Created"))

	})
}
