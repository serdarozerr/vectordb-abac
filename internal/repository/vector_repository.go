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
	UpdateCollection()
}

type Deleter interface {
	DeleteCollection()
}

type QueryMaker interface {
	Query()
}

type VectorRepository interface {
	Creator
	Updater
	Deleter
	QueryMaker
	Upserter
}
