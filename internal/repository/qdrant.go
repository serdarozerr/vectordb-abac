package repository

import (
	"context"
	"github.com/qdrant/go-client/qdrant"
)

type QdrantRepository struct {
	Client *qdrant.Client
}

type newUint64 uint64

func (q *QdrantRepository) CreateCollection(ctx context.Context, name string, vd int) error {
	request := &qdrant.CreateCollection{
		CollectionName: name,
		VectorsConfig: qdrant.NewVectorsConfig(&qdrant.VectorParams{
			Size:     uint64(vd),
			Distance: qdrant.Distance_Cosine,
		}),
	}
	err := q.Client.CreateCollection(ctx, request)
	if err != nil {
		return err
	}
	return nil

}

func (q *QdrantRepository) Upsert(ctx context.Context, collectionName string, indexes []uint64, vectors [][]float32, texts []string) error {
	var points = make([]*qdrant.PointStruct, len(indexes))

	for i, _ := range texts {
		point := &qdrant.PointStruct{
			Id:      qdrant.NewIDNum(indexes[i]),
			Vectors: qdrant.NewVectorsDense(vectors[i]),
			Payload: qdrant.NewValueMap(map[string]any{"text": texts[i]}),
		}
		points[i] = point
	}

	u := &qdrant.UpsertPoints{
		CollectionName: collectionName,
		Points:         points,
	}

	_, err := q.Client.Upsert(ctx, u)
	if err != nil {
		return err
	}
	return nil
}

func (q *QdrantRepository) DeleteCollection() {
}

func (q *QdrantRepository) Query(ctx context.Context, collectionName string, query []float32) (string, error) {
	var limit uint64 = 2
	var limit_ptr = &limit

	qp := &qdrant.QueryPoints{CollectionName: collectionName, Query: qdrant.NewQuery(query...), Limit: limit_ptr, WithPayload: qdrant.NewWithPayload(true)}

	result, err := q.Client.Query(ctx, qp)
	if err != nil {
		return "", err
	}
	var similar string = ""

	for _, point := range result {
		similar += point.Payload["text"].GetStringValue()
	}
	return similar, nil

}

func (q *QdrantRepository) UpdateCollection() {

}
