package db

import "qdrant-abac/config"

type Creator interface {
	CreateCollection()
}

type Updater interface {
	UpdateCollection()
}

type Deleter interface {
	DeleteCollection()
}

type MakeQuery interface {
	Query()
}

type DBClient interface {
	NewClient(cfg *config.Config)
}
