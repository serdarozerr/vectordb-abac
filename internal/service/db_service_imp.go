package service

import (
	"context"
	"log"
	"qdrant-abac/internal/model"
	"qdrant-abac/internal/repository"
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

func (ds *DBService) Query(logger *log.Logger) error {
	// Business logic to here
	logger.Println("Query collection")
	ds.Repository.Query()
	return nil
}
