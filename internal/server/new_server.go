package server

import (
	"github.com/serdarozerr/vectordb-abac/config"
	"github.com/serdarozerr/vectordb-abac/internal/handler"
	"github.com/serdarozerr/vectordb-abac/internal/service"
	"log"
	"net/http"
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
	mux.Handle("/api/v1/collection/query", lm(handler.QueryCollection(ds, logger, llm, config)))

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
