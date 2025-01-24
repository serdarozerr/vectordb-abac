package main

import (
	"log"
	"net/http"
	"qdrant-abac/config"
	"qdrant-abac/internal/db"
	"qdrant-abac/internal/handler"
	"qdrant-abac/internal/service"
)

func addRoutes(
	mux *http.ServeMux,
	logger *log.Logger,
	config *config.Config,
	dbClient db.DBClient,
) {

	lm := handler.LoggerMiddleware(logger)
	mux.Handle("/api/v1/create", lm(service.Create(logger, dbClient)))

}

func NewServer(
	logger *log.Logger,
	config *config.Config,
	dbClient db.DBClient,
) http.Handler {

	mux := http.NewServeMux()

	// TODO add the middlewares/handlers at here
	addRoutes(mux, logger, config, dbClient)

	return mux

}
