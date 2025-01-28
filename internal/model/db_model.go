package model

import (
	"context"
	"regexp"
)

type VectorDBCreate struct {
	Name string `json:"name"`
}

var reg, err = regexp.Compile(`^[a-zA-Z]+$`)

func (qc VectorDBCreate) Valid(ctx context.Context) map[string]string {
	reg.MatchString(qc.Name)
	problems := make(map[string]string)
	if !reg.MatchString(qc.Name) {
		problems[qc.Name] = "includes non letter characters"
	}

	return problems
}

type VectorDBInsert struct {
	CollectionName string
	Text           string
}
