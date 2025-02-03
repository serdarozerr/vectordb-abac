package instance

import (
	"errors"
	"github.com/qdrant/go-client/qdrant"
	"github.com/serdarozerr/vectordb-abac/config"
	"github.com/serdarozerr/vectordb-abac/internal/repository"
)

func NewRepository(dbType string, cfg *config.Config) (repository.VectorRepository, error) {

	switch dbType {
	case "qdrant":
		cl, err := qdrant.NewClient(
			&qdrant.Config{
				Host: cfg.VectorDB.URL,
				Port: cfg.VectorDB.Port,
			})
		if err != nil {
			panic(err)
		}
		r := &repository.QdrantRepository{Client: cl}
		return r, nil

	default:
		err := errors.New("unsupported repository type")
		return nil, err

	}

}
