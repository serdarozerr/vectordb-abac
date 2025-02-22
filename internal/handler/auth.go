package handler

import (
	"encoding/json"
	"github.com/serdarozerr/vectordb-abac/config"
	"github.com/serdarozerr/vectordb-abac/internal/model"
	"github.com/serdarozerr/vectordb-abac/internal/repository"
	"github.com/serdarozerr/vectordb-abac/internal/service/auth"
	"log"
	"net/http"
)

func ConvertCodeToToken(conf *config.Config, logger *log.Logger, c repository.Cache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := model.Decode[model.AuthCode](w, r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return

		}

		token, err := auth.TokenFromAuthCode(r.Context(), conf, c, data.Code)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		err = model.Encode[auth.OnlyAccessToken](w, r, 200, token)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		return

	}
}

func DecodeToken(conf *config.Config, logger *log.Logger, c repository.Cache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := model.Decode[model.AccessToken](w, r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		claims, err := auth.DecodeToken(r.Context(), conf, c, data.AccessToken)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		b, _ := json.Marshal(claims)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Decode Result:"))
		w.Write(b)
		return

	}
}
