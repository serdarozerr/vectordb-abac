package service

import (
	"context"
	"github.com/serdarozerr/vectordb-abac/internal/model"
	"github.com/serdarozerr/vectordb-abac/internal/repository"
	"log"
)

type DBService struct {
	Repository repository.VectorRepository
}

func NewDBService(r repository.VectorRepository) DBServicer {
	s := &DBService{Repository: r}
	return s
}

func (ds *DBService) CreateCollection(ctx context.Context, logger *log.Logger, data model.VectorDBCreate, vd int) error {
	// Business logic to here

	err := ds.Repository.CreateCollection(ctx, data.Name, vd)
	if err != nil {
		return err
	}
	return nil
}

func (ds *DBService) Insert(ctx context.Context, logger *log.Logger, llm *LLM, data model.VectorDBInsert, vd int) error {

	texts, err := Chunk(data.Text)
	if err != nil {
		return err
	}

	emb, err := llm.EmbedText(ctx, texts)
	if err != nil {
		return err
	}

	var vectors = make([][]float32, len(texts))
	var indexes = make([]uint64, len(texts))
	for i := range texts {
		vectors[i] = make([]float32, vd)
	}

	for i, e := range emb {
		indexes[i] = uint64(e.Index)
		for j, v := range e.Embedding {
			vectors[i][j] = float32(v)
		}

	}

	err = ds.Repository.Upsert(ctx, data.CollectionName, indexes, vectors, texts)
	if err != nil {
		return err
	}

	return nil

}

func (ds *DBService) UpdateCollection(logger *log.Logger) error {
	// Business logic to here
	logger.Println("Update collection")
	ds.Repository.UpdateCollection()
	return nil
}

func (ds *DBService) DeleteCollection(logger *log.Logger) error {
	// Business logic to here
	logger.Println("Delete collection")
	ds.Repository.DeleteCollection()
	return nil
}

func (ds *DBService) QueryCollection(ctx context.Context, logger *log.Logger, llm *LLM, data model.VectorDBQuery, vd int) (string, error) {
	// Business logic to here

	query := make([]string, 1)
	query[0] = data.Query

	emb, err := llm.EmbedText(ctx, query)
	if err != nil {
		return "", err
	}

	vector32 := make([]float32, vd)
	for _, e := range emb {
		for i, v := range e.Embedding {
			vector32[i] = float32(v)
		}
	}

	s, err := ds.Repository.Query(ctx, data.CollectionName, vector32)

	res, err := llm.CompleteText(ctx, s)

	if err != nil {
		return "", err
	}
	return res, nil
}
