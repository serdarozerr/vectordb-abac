package service

import (
	"context"
	"github.com/serdarozerr/vectordb-abac/internal/model"
	"log"
)

type Creator interface {
	CreateCollection(ctx context.Context, logger *log.Logger, d model.VectorDBCreate, vd int) error
}

type Inserter interface {
	Insert(ctx context.Context, logger *log.Logger, llm *LLM, data model.VectorDBInsert, vd int) error
}

type Updater interface {
	UpdateCollection(logger *log.Logger) error
}

type Deleter interface {
	DeleteCollection(logger *log.Logger) error
}

type MakeQuery interface {
	QueryCollection(ctx context.Context, logger *log.Logger, llm *LLM, data model.VectorDBQuery, vd int) (string, error)
}

type DBServicer interface {
	Creator
	Updater
	Deleter
	MakeQuery
	Inserter
}
