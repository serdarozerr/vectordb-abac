package server

import (
	"github.com/serdarozerr/vectordb-abac/config"
	"github.com/serdarozerr/vectordb-abac/internal/handler"
	"github.com/serdarozerr/vectordb-abac/internal/repository"
	"github.com/serdarozerr/vectordb-abac/internal/service"
	"log"
	"net/http"
)

func addRoutesCollection(
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

func addRoutesAuth(mux *http.ServeMux, config *config.Config, logger *log.Logger, c repository.Cache) {
	mux.Handle("/api/v1/auth/token", handler.ConvertCodeToToken(config, logger, c))
	mux.Handle("/api/v1/auth/token-decode", handler.DecodeToken(config, logger, c))
}

func NewServer(
	logger *log.Logger,
	config *config.Config,
	ds service.DBServicer,
	llm *service.LLM,
	c repository.Cache,
) http.Handler {

	mux := http.NewServeMux()

	// add the middlewares/handlers at here
	addRoutesCollection(mux, logger, ds, llm, config)
	addRoutesAuth(mux, config, logger, c)

	return mux

}
