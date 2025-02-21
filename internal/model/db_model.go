package model

import (
	"context"
	"regexp"
)

// VectorDBCreate create vector db collection with given name
type VectorDBCreate struct {
	Name string `json:"name"`
}

var reg, err = regexp.Compile(`^[a-zA-Z]+$`)

func (qc VectorDBCreate) Valid(ctx context.Context) map[string]string {
	reg.MatchString(qc.Name)
	problems := make(map[string]string)
	if !reg.MatchString(qc.Name) {
		problems[qc.Name] = "Includes non letter characters"
	}

	return problems
}

// VectorDBInsert params for the inserting the text from giving file
type VectorDBInsert struct {
	CollectionName string
	Text           string
}

// VectorDBQuery params for the Querying vector DB
type VectorDBQuery struct {
	CollectionName string
	Query          string
}

func (qc VectorDBQuery) Valid(ctx context.Context) map[string]string {
	problems := make(map[string]string)

	if qc.Query == "" {
		problems[qc.Query] = "The query is empty"
	}
	if qc.CollectionName == "" {
		problems[qc.Query] = "The collection name is empty"
	}
	return problems

}
