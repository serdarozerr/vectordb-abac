package repository

import (
	"fmt"
)

type MilvusRepository struct {
	//add milvus variable
}

func (m *MilvusRepository) CreateCollection() {
	fmt.Println("Creating collection")
}

func (m *MilvusRepository) DeleteCollection() {
}

func (m *MilvusRepository) Query() {

}

func (m *MilvusRepository) UpdateCollection() {

}
