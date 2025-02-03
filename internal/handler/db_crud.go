package handler

import (
	"github.com/serdarozerr/vectordb-abac/config"
	"github.com/serdarozerr/vectordb-abac/internal/model"
	"github.com/serdarozerr/vectordb-abac/internal/service"
	"log"
	"net/http"
)

func CreateCollection(ds service.DBServicer, logger *log.Logger, config *config.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		data, err := model.Decode[model.VectorDBCreate](w, r)

		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		err = ds.CreateCollection(r.Context(), logger, data, config.LLM.VectorDimension)
		if err != nil {
			logger.Println(err)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("created"))

	})

}

func InsertFileToVectorDB(ds service.DBServicer, logger *log.Logger, llm *service.LLM, config *config.Config) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		err := r.ParseMultipartForm(10 << 20)
		if err != nil {
			logger.Println(err)
			w.Write([]byte(err.Error()))
		}

		file, fileHeader, err := r.FormFile("file")
		if err != nil {
			logger.Println(err)
			w.Write([]byte(err.Error()))
		}
		defer file.Close()
		textBytes := make([]byte, fileHeader.Size)
		file.Read(textBytes)

		data := model.VectorDBInsert{
			CollectionName: r.FormValue("CollectionName"),
			Text:           string(textBytes),
		}

		err = ds.Insert(r.Context(), logger, llm, data, config.LLM.VectorDimension)
		if err != nil {
			logger.Println(err)
			w.Write([]byte(err.Error()))
		}

	})

}

func QueryCollection(ds service.DBServicer, logger *log.Logger, llm *service.LLM, config *config.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data, err := model.Decode[model.VectorDBQuery](w, r)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		res, err := ds.QueryCollection(r.Context(), logger, llm, data, config.LLM.VectorDimension)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(res))

	})
}
