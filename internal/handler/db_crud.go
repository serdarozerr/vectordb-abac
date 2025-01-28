package handler

import (
	"log"
	"net/http"
	"qdrant-abac/config"
	"qdrant-abac/internal/model"
	"qdrant-abac/internal/service"
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

//curl -X POST \
//-H "Content-type: application/json" \
//-H "Accept: application/json" \
//-d '{"name":"testcollection"}' \
//"http://localhost:8000/api/v1/collection/create"

//curl -X POST -F "file=@test.txt" -F "CollectionName=testcollection"  http://localhost:8000/api/v1/collection/insert
