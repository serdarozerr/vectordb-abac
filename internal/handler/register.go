package handler

import (
	"log"
	"net/http"
	"qdrant-abac/config"
)

func RegisterEndpoint(logger log.Logger, conf *config.Config) http.Handler {

	// init some things
	//qdb:=db.CreateClient()
	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		//qdb.CreateCollection
	}
	return handler

}
