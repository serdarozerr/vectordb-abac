package db

import (
	"github.com/qdrant/go-client/qdrant"
	"qdrant-abac/config"
)

//
//type Qdrant struct {
//	Creator
//	Updater
//	Deleter
//	MakeQuery
//	CreateClient
//}

type QdrantService struct {
	Client *qdrant.Client
}

func (q *QdrantService) NewClient(cfg *config.Config) {

	client, err := qdrant.NewClient(
		&qdrant.Config{
			Host: cfg.VectorDB.URL,
			Port: cfg.VectorDB.Port,
		})

	if err != nil {
		panic(err)
	}
	q.Client = client
}

func (q *QdrantService) CreateCollection() {

}

func (q *QdrantService) DeleteCollection() {

}

func (q *QdrantService) Query() {

}

func (q *QdrantService) UpdateCollection() {

}
