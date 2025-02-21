package repository

import (
	"context"
)

type Creator interface {
	CreateCollection(ctx context.Context, name string, vd int) error
}

type Upserter interface {
	Upsert(ctx context.Context, collectionName string, indexes []uint64, vectors [][]float32, texts []string) error
}
type Updater interface {
	UpdateCollection(ctx context.Context, collectionName string, indexes []uint64, vectors [][]float32) error
}

type Deleter interface {
	DeleteCollection()
}

type QueryMaker interface {
	Query(ctx context.Context, collectionName string, query []float32) (string, error)
}

type VectorRepository interface {
	Creator
	Updater
	Deleter
	QueryMaker
	Upserter
}
