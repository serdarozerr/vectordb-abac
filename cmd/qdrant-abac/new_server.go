package main

import (
	"log"
	"net/http"
	"qdrant-abac/config"
	"qdrant-abac/internal/handler"
	"qdrant-abac/internal/service"
)

func addRoutes(
	mux *http.ServeMux,
	logger *log.Logger,
	ds service.DBServicer,
	llm *service.LLM,
	config *config.Config,
) {

	lm := handler.LoggerMiddleware(logger)
	mux.Handle("/api/v1/collection/create", lm(handler.CreateCollection(ds, logger, config)))
	mux.Handle("/api/v1/collection/insert", lm(handler.InsertFileToVectorDB(ds, logger, llm, config)))

}

func NewServer(
	logger *log.Logger,
	config *config.Config,
	ds service.DBServicer,
	llm *service.LLM,
) http.Handler {

	mux := http.NewServeMux()

	// TODO add the middlewares/handlers at here
	addRoutes(mux, logger, ds, llm, config)

	return mux

}
